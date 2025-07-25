// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLocation interface {
	dara.Model
	String() string
	GoString() string
	SetLocationType(v string) *Location
	GetLocationType() *string
	SetLocationValue(v map[string]interface{}) *Location
	GetLocationValue() map[string]interface{}
}

type Location struct {
	LocationType  *string                `json:"LocationType,omitempty" xml:"LocationType,omitempty"`
	LocationValue map[string]interface{} `json:"LocationValue,omitempty" xml:"LocationValue,omitempty"`
}

func (s Location) String() string {
	return dara.Prettify(s)
}

func (s Location) GoString() string {
	return s.String()
}

func (s *Location) GetLocationType() *string {
	return s.LocationType
}

func (s *Location) GetLocationValue() map[string]interface{} {
	return s.LocationValue
}

func (s *Location) SetLocationType(v string) *Location {
	s.LocationType = &v
	return s
}

func (s *Location) SetLocationValue(v map[string]interface{}) *Location {
	s.LocationValue = v
	return s
}

func (s *Location) Validate() error {
	return dara.Validate(s)
}
