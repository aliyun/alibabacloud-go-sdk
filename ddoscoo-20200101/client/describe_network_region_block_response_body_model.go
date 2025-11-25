// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkRegionBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *DescribeNetworkRegionBlockResponseBodyConfig) *DescribeNetworkRegionBlockResponseBody
	GetConfig() *DescribeNetworkRegionBlockResponseBodyConfig
	SetRequestId(v string) *DescribeNetworkRegionBlockResponseBody
	GetRequestId() *string
}

type DescribeNetworkRegionBlockResponseBody struct {
	// The configuration of blocked locations.
	Config *DescribeNetworkRegionBlockResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNetworkRegionBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRegionBlockResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRegionBlockResponseBody) GetConfig() *DescribeNetworkRegionBlockResponseBodyConfig {
	return s.Config
}

func (s *DescribeNetworkRegionBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkRegionBlockResponseBody) SetConfig(v *DescribeNetworkRegionBlockResponseBodyConfig) *DescribeNetworkRegionBlockResponseBody {
	s.Config = v
	return s
}

func (s *DescribeNetworkRegionBlockResponseBody) SetRequestId(v string) *DescribeNetworkRegionBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkRegionBlockResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkRegionBlockResponseBodyConfig struct {
	// The codes of the countries or areas from which the requests are blocked.
	Countries []*int64 `json:"Countries,omitempty" xml:"Countries,omitempty" type:"Repeated"`
	// The codes of the administrative regions in China from which the requests are blocked.
	Provinces []*int64 `json:"Provinces,omitempty" xml:"Provinces,omitempty" type:"Repeated"`
	// The status of the Location Blacklist policy. Valid values:
	//
	// 	- **on**: enabled
	//
	// 	- **off**: disabled
	//
	// example:
	//
	// on
	RegionBlockSwitch *string `json:"RegionBlockSwitch,omitempty" xml:"RegionBlockSwitch,omitempty"`
}

func (s DescribeNetworkRegionBlockResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRegionBlockResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRegionBlockResponseBodyConfig) GetCountries() []*int64 {
	return s.Countries
}

func (s *DescribeNetworkRegionBlockResponseBodyConfig) GetProvinces() []*int64 {
	return s.Provinces
}

func (s *DescribeNetworkRegionBlockResponseBodyConfig) GetRegionBlockSwitch() *string {
	return s.RegionBlockSwitch
}

func (s *DescribeNetworkRegionBlockResponseBodyConfig) SetCountries(v []*int64) *DescribeNetworkRegionBlockResponseBodyConfig {
	s.Countries = v
	return s
}

func (s *DescribeNetworkRegionBlockResponseBodyConfig) SetProvinces(v []*int64) *DescribeNetworkRegionBlockResponseBodyConfig {
	s.Provinces = v
	return s
}

func (s *DescribeNetworkRegionBlockResponseBodyConfig) SetRegionBlockSwitch(v string) *DescribeNetworkRegionBlockResponseBodyConfig {
	s.RegionBlockSwitch = &v
	return s
}

func (s *DescribeNetworkRegionBlockResponseBodyConfig) Validate() error {
	return dara.Validate(s)
}
