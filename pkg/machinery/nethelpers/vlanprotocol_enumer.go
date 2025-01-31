// Code generated by "enumer -type=VLANProtocol -linecomment -text"; DO NOT EDIT.

//
package nethelpers

import (
	"fmt"
)

const (
	_VLANProtocolName_0 = "802.1q"
	_VLANProtocolName_1 = "802.1ad"
)

var (
	_VLANProtocolIndex_0 = [...]uint8{0, 6}
	_VLANProtocolIndex_1 = [...]uint8{0, 7}
)

func (i VLANProtocol) String() string {
	switch {
	case i == 33024:
		return _VLANProtocolName_0
	case i == 34984:
		return _VLANProtocolName_1
	default:
		return fmt.Sprintf("VLANProtocol(%d)", i)
	}
}

var _VLANProtocolValues = []VLANProtocol{33024, 34984}

var _VLANProtocolNameToValueMap = map[string]VLANProtocol{
	_VLANProtocolName_0[0:6]: 33024,
	_VLANProtocolName_1[0:7]: 34984,
}

// VLANProtocolString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func VLANProtocolString(s string) (VLANProtocol, error) {
	if val, ok := _VLANProtocolNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to VLANProtocol values", s)
}

// VLANProtocolValues returns all values of the enum
func VLANProtocolValues() []VLANProtocol {
	return _VLANProtocolValues
}

// IsAVLANProtocol returns "true" if the value is listed in the enum definition. "false" otherwise
func (i VLANProtocol) IsAVLANProtocol() bool {
	for _, v := range _VLANProtocolValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for VLANProtocol
func (i VLANProtocol) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for VLANProtocol
func (i *VLANProtocol) UnmarshalText(text []byte) error {
	var err error
	*i, err = VLANProtocolString(string(text))
	return err
}
