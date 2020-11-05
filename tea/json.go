package tea

import (
	"encoding/json"
)

// MustJson always returns bytes with json object.
// In case of failure information about error is lost
// and empty json is returned instead.
//
func MustJson(o interface{}) []byte {
	if b, err := json.Marshal(o); err == nil {
		return b
	}

	return []byte("{}")
}
