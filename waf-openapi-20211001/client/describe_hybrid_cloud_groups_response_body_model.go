// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*DescribeHybridCloudGroupsResponseBodyGroups) *DescribeHybridCloudGroupsResponseBody
	GetGroups() []*DescribeHybridCloudGroupsResponseBodyGroups
	SetRequestId(v string) *DescribeHybridCloudGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeHybridCloudGroupsResponseBody
	GetTotalCount() *int32
}

type DescribeHybridCloudGroupsResponseBody struct {
	// The list of hybrid cloud node groups.
	Groups []*DescribeHybridCloudGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 045660E7-C4C6-5CD7-8182-7B337D95ADF4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of hybrid cloud node groups returned.
	//
	// example:
	//
	// 146
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHybridCloudGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudGroupsResponseBody) GetGroups() []*DescribeHybridCloudGroupsResponseBodyGroups {
	return s.Groups
}

func (s *DescribeHybridCloudGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHybridCloudGroupsResponseBody) SetGroups(v []*DescribeHybridCloudGroupsResponseBodyGroups) *DescribeHybridCloudGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBody) SetRequestId(v string) *DescribeHybridCloudGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBody) SetTotalCount(v int32) *DescribeHybridCloudGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBody) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridCloudGroupsResponseBodyGroups struct {
	// The back-to-origin mark of the protection cluster. The value is in the **{CarrierTag}-{ContinentTag}-{CityTag}-{Identifier}*	- format. The identifier is optional.
	//
	// > For a list of valid values, see Additional information about response parameters.
	//
	// example:
	//
	// aliyun-asiapacific-beijing-56477821
	BackSourceMark *string `json:"BackSourceMark,omitempty" xml:"BackSourceMark,omitempty"`
	// The continent code of the protection cluster.
	//
	// > For a list of valid codes, see Additional information about response parameters.
	//
	// example:
	//
	// 410
	ContinentsValue *int32 `json:"ContinentsValue,omitempty" xml:"ContinentsValue,omitempty"`
	// The ID of the hybrid cloud node group.
	//
	// example:
	//
	// 123
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the hybrid cloud node group.
	//
	// example:
	//
	// StorageGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the hybrid cloud node group. Valid values:
	//
	// - **protect**: protection node group.
	//
	// - **control**: control node group.
	//
	// - **storage**: storage node group.
	//
	// - **controlStorage**: control and storage node group.
	//
	// example:
	//
	// protect
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The IP address of the load balancer that is associated with the hybrid cloud node group.
	//
	// example:
	//
	// 1.1.XX.XX
	LoadBalanceIp *string `json:"LoadBalanceIp,omitempty" xml:"LoadBalanceIp,omitempty"`
	// The ID of the protection node.
	//
	// example:
	//
	// 1312
	LocationId *int64 `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	// The carrier code of the protection cluster.
	//
	// > For a list of valid codes, see Additional information about response parameters.
	//
	// example:
	//
	// 0
	OperatorValue *int32 `json:"OperatorValue,omitempty" xml:"OperatorValue,omitempty"`
	// The ports that are used by the hybrid cloud cluster. Multiple ports are separated by commas (,).
	//
	// example:
	//
	// 80,9200,20018
	Ports *string `json:"Ports,omitempty" xml:"Ports,omitempty"`
	// The city code of the protection cluster.
	//
	// > For a list of valid codes, see Additional information about response parameters.
	//
	// example:
	//
	// 0
	RegionCodeValue *int32 `json:"RegionCodeValue,omitempty" xml:"RegionCodeValue,omitempty"`
	// The description of the hybrid cloud node group.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DescribeHybridCloudGroupsResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) GetBackSourceMark() *string {
	return s.BackSourceMark
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) GetContinentsValue() *int32 {
	return s.ContinentsValue
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) GetGroupId() *int32 {
	return s.GroupId
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) GetLoadBalanceIp() *string {
	return s.LoadBalanceIp
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) GetLocationId() *int64 {
	return s.LocationId
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) GetOperatorValue() *int32 {
	return s.OperatorValue
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) GetPorts() *string {
	return s.Ports
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) GetRegionCodeValue() *int32 {
	return s.RegionCodeValue
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) GetRemark() *string {
	return s.Remark
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetBackSourceMark(v string) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.BackSourceMark = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetContinentsValue(v int32) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.ContinentsValue = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetGroupId(v int32) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetGroupName(v string) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetGroupType(v string) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.GroupType = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetLoadBalanceIp(v string) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.LoadBalanceIp = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetLocationId(v int64) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.LocationId = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetOperatorValue(v int32) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.OperatorValue = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetPorts(v string) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.Ports = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetRegionCodeValue(v int32) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.RegionCodeValue = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetRemark(v string) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.Remark = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) Validate() error {
	return dara.Validate(s)
}
