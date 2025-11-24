// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeVSwitchesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeVSwitchesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeVSwitchesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVSwitchesResponseBody
	GetTotalCount() *int32
	SetVSwitches(v []*DescribeVSwitchesResponseBodyVSwitches) *DescribeVSwitchesResponseBody
	GetVSwitches() []*DescribeVSwitchesResponseBodyVSwitches
}

type DescribeVSwitchesResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// “”
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of vSwitches that are deployed in the VPC in the region. This parameter is optional and is not returned by default.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The available vSwitches.
	VSwitches []*DescribeVSwitchesResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeVSwitchesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVSwitchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVSwitchesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVSwitchesResponseBody) GetVSwitches() []*DescribeVSwitchesResponseBodyVSwitches {
	return s.VSwitches
}

func (s *DescribeVSwitchesResponseBody) SetMaxResults(v int32) *DescribeVSwitchesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetNextToken(v string) *DescribeVSwitchesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetRequestId(v string) *DescribeVSwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetTotalCount(v int32) *DescribeVSwitchesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetVSwitches(v []*DescribeVSwitchesResponseBodyVSwitches) *DescribeVSwitchesResponseBody {
	s.VSwitches = v
	return s
}

func (s *DescribeVSwitchesResponseBody) Validate() error {
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

type DescribeVSwitchesResponseBodyVSwitches struct {
	// Indicates whether the vSwitch is the default vSwitch. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The state of the vSwitch. Valid values:
	//
	// 	- `Pending`: The vSwitch is being configured.
	//
	// 	- `Available`: The vSwitch is available.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1g24p9no0iqir46****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// vsw-test
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The ID of the VPC to which the vSwitch belongs.
	//
	// example:
	//
	// vpc-bp17gig441u0msmd6****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone to which the switch belongs.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchesResponseBodyVSwitches) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitches) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVSwitchesResponseBodyVSwitches) GetStatus() *string {
	return s.Status
}

func (s *DescribeVSwitchesResponseBodyVSwitches) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchesResponseBodyVSwitches) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchesResponseBodyVSwitches) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVSwitchesResponseBodyVSwitches) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetIsDefault(v bool) *DescribeVSwitchesResponseBodyVSwitches {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetStatus(v string) *DescribeVSwitchesResponseBodyVSwitches {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetVSwitchId(v string) *DescribeVSwitchesResponseBodyVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetVSwitchName(v string) *DescribeVSwitchesResponseBodyVSwitches {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetVpcId(v string) *DescribeVSwitchesResponseBodyVSwitches {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetZoneId(v string) *DescribeVSwitchesResponseBodyVSwitches {
	s.ZoneId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) Validate() error {
	return dara.Validate(s)
}
