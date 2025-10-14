// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNatGateways(v []*DescribeNatGatewaysResponseBodyNatGateways) *DescribeNatGatewaysResponseBody
	GetNatGateways() []*DescribeNatGatewaysResponseBodyNatGateways
	SetPageNumber(v int32) *DescribeNatGatewaysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNatGatewaysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeNatGatewaysResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeNatGatewaysResponseBody
	GetTotalCount() *int32
}

type DescribeNatGatewaysResponseBody struct {
	// The details of the NAT gateways.
	NatGateways []*DescribeNatGatewaysResponseBodyNatGateways `json:"NatGateways,omitempty" xml:"NatGateways,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2818A8F4-5E2B-5611-8732-5ACF7B677059
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of NAT gateways that are returned.
	//
	// example:
	//
	// 49
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNatGatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBody) GetNatGateways() []*DescribeNatGatewaysResponseBodyNatGateways {
	return s.NatGateways
}

func (s *DescribeNatGatewaysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNatGatewaysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNatGatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatGatewaysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNatGatewaysResponseBody) SetNatGateways(v []*DescribeNatGatewaysResponseBodyNatGateways) *DescribeNatGatewaysResponseBody {
	s.NatGateways = v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetPageNumber(v int32) *DescribeNatGatewaysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetPageSize(v int32) *DescribeNatGatewaysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetRequestId(v string) *DescribeNatGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetTotalCount(v int32) *DescribeNatGatewaysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) Validate() error {
	if s.NatGateways != nil {
		for _, item := range s.NatGateways {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatGatewaysResponseBodyNatGateways struct {
	// The time when the NAT gateway was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-22T07:03:32Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the ENS node.
	//
	// example:
	//
	// cn-xiangyang-5
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The EIPs that are associated with the NAT gateway.
	IpLists []*DescribeNatGatewaysResponseBodyNatGatewaysIpLists `json:"IpLists,omitempty" xml:"IpLists,omitempty" type:"Repeated"`
	// The name of the NAT gateway.
	//
	// example:
	//
	// test0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// nat-5t7nh1cfm6kxiszlttr38****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// n-5s2ml1olj0kzaws9n1yrj****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The type of the NAT gateway.
	//
	// example:
	//
	// enat.default
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the NAT gateway. Valid values:
	//
	// 	- **Creating**: After you send a request to create a NAT gateway, the system creates the NAT gateway in the background. The NAT gateway remains in the Creating state until the operation is completed.
	//
	// 	- **Available**: The NAT gateway is in the Available state after the creation is complete.
	//
	// 	- **Deleting**: After you send a request to delete a NAT gateway, the system deletes the NAT gateway in the background. The NAT gateway remains in the Deleting state until the operation is completed.
	//
	// example:
	//
	// Available
	Status *string                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*DescribeNatGatewaysResponseBodyNatGatewaysTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-5szpp1os9m55myirbflfw****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeNatGatewaysResponseBodyNatGateways) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGateways) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetIpLists() []*DescribeNatGatewaysResponseBodyNatGatewaysIpLists {
	return s.IpLists
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetName() *string {
	return s.Name
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetSpec() *string {
	return s.Spec
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetStatus() *string {
	return s.Status
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetTags() []*DescribeNatGatewaysResponseBodyNatGatewaysTags {
	return s.Tags
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetCreationTime(v string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.CreationTime = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetEnsRegionId(v string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetIpLists(v []*DescribeNatGatewaysResponseBodyNatGatewaysIpLists) *DescribeNatGatewaysResponseBodyNatGateways {
	s.IpLists = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetName(v string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.Name = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetNatGatewayId(v string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetNetworkId(v string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.NetworkId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetSpec(v string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.Spec = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetStatus(v string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.Status = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetTags(v []*DescribeNatGatewaysResponseBodyNatGatewaysTags) *DescribeNatGatewaysResponseBodyNatGateways {
	s.Tags = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetVSwitchId(v string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) Validate() error {
	if s.IpLists != nil {
		for _, item := range s.IpLists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatGatewaysResponseBodyNatGatewaysIpLists struct {
	// The ID of the EIP.
	//
	// example:
	//
	// eip-50g****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The EIP.
	//
	// example:
	//
	// 8.XX.XX.162
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The association between the EIP and the Internet NAT gateway. Valid values:
	//
	// 	- **UsedByForwardTable**: The EIP is specified in a DNAT entry.
	//
	// 	- **UsedBySnatTable**: The EIP is specified in an SNAT entry.
	//
	// 	- **Idle**: The EIP is not specified in an SNAT entry or a DNAT entry.
	//
	// example:
	//
	// Idle
	UsingStatus *string `json:"UsingStatus,omitempty" xml:"UsingStatus,omitempty"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysIpLists) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysIpLists) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) GetUsingStatus() *string {
	return s.UsingStatus
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) SetAllocationId(v string) *DescribeNatGatewaysResponseBodyNatGatewaysIpLists {
	s.AllocationId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) SetIpAddress(v string) *DescribeNatGatewaysResponseBodyNatGatewaysIpLists {
	s.IpAddress = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) SetUsingStatus(v string) *DescribeNatGatewaysResponseBodyNatGatewaysIpLists {
	s.UsingStatus = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewaysResponseBodyNatGatewaysTags struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Deprecated
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// Deprecated
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysTags) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysTags) GetKey() *string {
	return s.Key
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysTags) GetValue() *string {
	return s.Value
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysTags) SetKey(v string) *DescribeNatGatewaysResponseBodyNatGatewaysTags {
	s.Key = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysTags) SetTagKey(v string) *DescribeNatGatewaysResponseBodyNatGatewaysTags {
	s.TagKey = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysTags) SetTagValue(v string) *DescribeNatGatewaysResponseBodyNatGatewaysTags {
	s.TagValue = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysTags) SetValue(v string) *DescribeNatGatewaysResponseBodyNatGatewaysTags {
	s.Value = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysTags) Validate() error {
	return dara.Validate(s)
}
