package pb

import (
	"errors"
	"fmt"
	"sort"

	"database/sql/driver"
)

var (
	languageP2S = map[Language]string{
		Language_VI: "VI",
		Language_EN: "EN",
	}

	languageS2P = inverseMap(languageP2S)
)

func (Language) Values() []string {
	var values []string
	for key := range Language_name {
		x := Language(key)
		if x == Language_UNKNOWN {
			continue
		}
		name, ok := languageP2S[x]
		if !ok {
			panic(fmt.Sprintf("missing database value mapping for enum %s", x.String()))
		}

		values = append(values, name)
	}

	sort.Strings(values)

	return values
}

func (x Language) Value() (driver.Value, error) {
	val, ok := languageP2S[x]
	if !ok {
		return nil, errors.New("invalid enum value")
	}

	return val, nil
}

func (x *Language) Scan(val any) error {
	var s string

	switch v := val.(type) {
	case nil:
		return nil
	case string:
		s = v
	case []byte:
		s = string(v)
	}

	ev, ok := languageS2P[s]
	if !ok {
		*x = Language_UNKNOWN
		return nil
	}

	*x = ev
	return nil
}
