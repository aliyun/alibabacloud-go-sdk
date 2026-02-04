// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrantVSwitchesToCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListGrantVSwitchesToCenResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGrantVSwitchesToCenResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListGrantVSwitchesToCenResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListGrantVSwitchesToCenResponseBody
	GetTotalCount() *int32
	SetVSwitches(v []*ListGrantVSwitchesToCenResponseBodyVSwitches) *ListGrantVSwitchesToCenResponseBody
	GetVSwitches() []*ListGrantVSwitchesToCenResponseBodyVSwitches
}

type ListGrantVSwitchesToCenResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A9288C78-881A-5D30-A8A9-68E05EE0A086
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of vSwitches.
	VSwitches []*ListGrantVSwitchesToCenResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
}

func (s ListGrantVSwitchesToCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGrantVSwitchesToCenResponseBody) GoString() string {
	return s.String()
}

func (s *ListGrantVSwitchesToCenResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGrantVSwitchesToCenResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGrantVSwitchesToCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGrantVSwitchesToCenResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListGrantVSwitchesToCenResponseBody) GetVSwitches() []*ListGrantVSwitchesToCenResponseBodyVSwitches {
	return s.VSwitches
}

func (s *ListGrantVSwitchesToCenResponseBody) SetPageNumber(v int32) *ListGrantVSwitchesToCenResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBody) SetPageSize(v int32) *ListGrantVSwitchesToCenResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBody) SetRequestId(v string) *ListGrantVSwitchesToCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBody) SetTotalCount(v int32) *ListGrantVSwitchesToCenResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBody) SetVSwitches(v []*ListGrantVSwitchesToCenResponseBodyVSwitches) *ListGrantVSwitchesToCenResponseBody {
	s.VSwitches = v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBody) Validate() error {
	if s.VSwitches != nil {
		for _, item := range s.VSwitches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGrantVSwitchesToCenResponseBodyVSwitches struct {
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1194lh263wx1gsk****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// nametest
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The ID of the VPC to which the vSwitch belongs.
	//
	// example:
	//
	// vpc-bp12ge2tq5gzdc915****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone to which the vSwitch belongs.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListGrantVSwitchesToCenResponseBodyVSwitches) String() string {
	return dara.Prettify(s)
}

func (s ListGrantVSwitchesToCenResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *ListGrantVSwitchesToCenResponseBodyVSwitches) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListGrantVSwitchesToCenResponseBodyVSwitches) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *ListGrantVSwitchesToCenResponseBodyVSwitches) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGrantVSwitchesToCenResponseBodyVSwitches) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListGrantVSwitchesToCenResponseBodyVSwitches) SetVSwitchId(v string) *ListGrantVSwitchesToCenResponseBodyVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBodyVSwitches) SetVSwitchName(v string) *ListGrantVSwitchesToCenResponseBodyVSwitches {
	s.VSwitchName = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBodyVSwitches) SetVpcId(v string) *ListGrantVSwitchesToCenResponseBodyVSwitches {
	s.VpcId = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBodyVSwitches) SetZoneId(v string) *ListGrantVSwitchesToCenResponseBodyVSwitches {
	s.ZoneId = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBodyVSwitches) Validate() error {
	return dara.Validate(s)
}
