// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeRegionsResponseBody
	GetCode() *int32
	SetData(v *DescribeRegionsResponseBodyData) *DescribeRegionsResponseBody
	GetData() *DescribeRegionsResponseBodyData
	SetMessage(v string) *DescribeRegionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRegionsResponseBody
	GetSuccess() *bool
}

type DescribeRegionsResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeRegionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeRegionsResponseBody) GetData() *DescribeRegionsResponseBodyData {
	return s.Data
}

func (s *DescribeRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRegionsResponseBody) SetCode(v int32) *DescribeRegionsResponseBody {
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

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionsResponseBodyData struct {
	Regions []*DescribeRegionsResponseBodyDataRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
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
	RegionCnName *string `json:"RegionCnName,omitempty" xml:"RegionCnName,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName   *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeRegionsResponseBodyDataRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyDataRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyDataRegions) GetRegionCnName() *string {
	return s.RegionCnName
}

func (s *DescribeRegionsResponseBodyDataRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyDataRegions) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeRegionsResponseBodyDataRegions) SetRegionCnName(v string) *DescribeRegionsResponseBodyDataRegions {
	s.RegionCnName = &v
	return s
}

func (s *DescribeRegionsResponseBodyDataRegions) SetRegionId(v string) *DescribeRegionsResponseBodyDataRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyDataRegions) SetRegionName(v string) *DescribeRegionsResponseBodyDataRegions {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionsResponseBodyDataRegions) Validate() error {
	return dara.Validate(s)
}
