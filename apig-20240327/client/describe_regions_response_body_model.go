// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeRegionsResponseBody
	GetCode() *string
	SetData(v *DescribeRegionsResponseBodyData) *DescribeRegionsResponseBody
	GetData() *DescribeRegionsResponseBodyData
	SetMessage(v string) *DescribeRegionsResponseBody
	GetMessage() *string
	SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody
	GetRegions() *DescribeRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
}

type DescribeRegionsResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *DescribeRegionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Deprecated
	Regions *DescribeRegionsResponseBodyRegions `json:"regions,omitempty" xml:"regions,omitempty" type:"Struct"`
	// example:
	//
	// E6BD6C79-32B1-5D7E-A89A-93C5A6B7xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeRegionsResponseBody) GetData() *DescribeRegionsResponseBodyData {
	return s.Data
}

func (s *DescribeRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRegionsResponseBody) GetRegions() *DescribeRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) SetCode(v string) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetData(v *DescribeRegionsResponseBodyData) *DescribeRegionsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRegionsResponseBody) SetMessage(v string) *DescribeRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	if s.Regions != nil {
		if err := s.Regions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionsResponseBodyData struct {
	Regions []*DescribeRegionsResponseBodyDataRegions `json:"regions,omitempty" xml:"regions,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyData) GetRegions() []*DescribeRegionsResponseBodyDataRegions {
	return s.Regions
}

func (s *DescribeRegionsResponseBodyData) SetRegions(v []*DescribeRegionsResponseBodyDataRegions) *DescribeRegionsResponseBodyData {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBodyData) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyDataRegions struct {
	LocalName *string `json:"localName,omitempty" xml:"localName,omitempty"`
	RegionId  *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescribeRegionsResponseBodyDataRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyDataRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyDataRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyDataRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyDataRegions) SetLocalName(v string) *DescribeRegionsResponseBodyDataRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyDataRegions) SetRegionId(v string) *DescribeRegionsResponseBodyDataRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyDataRegions) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) GetRegion() []*DescribeRegionsResponseBodyRegionsRegion {
	return s.Region
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) Validate() error {
	if s.Region != nil {
		for _, item := range s.Region {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	// example:
	//
	// 华东1（杭州）
	LocalName *string `json:"localName,omitempty" xml:"localName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) Validate() error {
	return dara.Validate(s)
}
