// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iABTestExperiment interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ABTestExperiment
	GetName() *string
	SetOnline(v bool) *ABTestExperiment
	GetOnline() *bool
	SetParams(v map[string]*string) *ABTestExperiment
	GetParams() map[string]*string
	SetSerialNumber(v int32) *ABTestExperiment
	GetSerialNumber() *int32
	SetTraffic(v int32) *ABTestExperiment
	GetTraffic() *int32
}

type ABTestExperiment struct {
	Name         *string            `json:"name,omitempty" xml:"name,omitempty"`
	Online       *bool              `json:"online,omitempty" xml:"online,omitempty"`
	Params       map[string]*string `json:"params,omitempty" xml:"params,omitempty"`
	SerialNumber *int32             `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	Traffic      *int32             `json:"traffic,omitempty" xml:"traffic,omitempty"`
}

func (s ABTestExperiment) String() string {
	return dara.Prettify(s)
}

func (s ABTestExperiment) GoString() string {
	return s.String()
}

func (s *ABTestExperiment) GetName() *string {
	return s.Name
}

func (s *ABTestExperiment) GetOnline() *bool {
	return s.Online
}

func (s *ABTestExperiment) GetParams() map[string]*string {
	return s.Params
}

func (s *ABTestExperiment) GetSerialNumber() *int32 {
	return s.SerialNumber
}

func (s *ABTestExperiment) GetTraffic() *int32 {
	return s.Traffic
}

func (s *ABTestExperiment) SetName(v string) *ABTestExperiment {
	s.Name = &v
	return s
}

func (s *ABTestExperiment) SetOnline(v bool) *ABTestExperiment {
	s.Online = &v
	return s
}

func (s *ABTestExperiment) SetParams(v map[string]*string) *ABTestExperiment {
	s.Params = v
	return s
}

func (s *ABTestExperiment) SetSerialNumber(v int32) *ABTestExperiment {
	s.SerialNumber = &v
	return s
}

func (s *ABTestExperiment) SetTraffic(v int32) *ABTestExperiment {
	s.Traffic = &v
	return s
}

func (s *ABTestExperiment) Validate() error {
	return dara.Validate(s)
}
