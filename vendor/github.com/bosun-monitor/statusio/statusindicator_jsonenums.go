// generated by jsonenums -type=StatusIndicator; DO NOT EDIT

package statusio

import (
	"encoding/json"
	"fmt"
)

var (
	_StatusIndicatorNameToValue = map[string]StatusIndicator{
		"None":     None,
		"Minor":    Minor,
		"Major":    Major,
		"Critical": Critical,
	}

	_StatusIndicatorValueToName = map[StatusIndicator]string{
		None:     "None",
		Minor:    "Minor",
		Major:    "Major",
		Critical: "Critical",
	}
)

func init() {
	var v StatusIndicator
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_StatusIndicatorNameToValue = map[string]StatusIndicator{
			interface{}(None).(fmt.Stringer).String():     None,
			interface{}(Minor).(fmt.Stringer).String():    Minor,
			interface{}(Major).(fmt.Stringer).String():    Major,
			interface{}(Critical).(fmt.Stringer).String(): Critical,
		}
	}
}

// MarshalJSON is generated so StatusIndicator satisfies json.Marshaler.
func (r StatusIndicator) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _StatusIndicatorValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid StatusIndicator: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so StatusIndicator satisfies json.Unmarshaler.
func (r *StatusIndicator) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("StatusIndicator should be a string, got %s", data)
	}
	v, ok := _StatusIndicatorNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid StatusIndicator %q", s)
	}
	*r = v
	return nil
}
