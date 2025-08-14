// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrantVSwitchEnisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGrantVSwitchEnis(v []*ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) *ListGrantVSwitchEnisResponseBody
	GetGrantVSwitchEnis() []*ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis
	SetMaxResults(v int64) *ListGrantVSwitchEnisResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListGrantVSwitchEnisResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListGrantVSwitchEnisResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListGrantVSwitchEnisResponseBody
	GetTotalCount() *string
}

type ListGrantVSwitchEnisResponseBody struct {
	// The information about the ENI.
	GrantVSwitchEnis []*ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis `json:"GrantVSwitchEnis,omitempty" xml:"GrantVSwitchEnis,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DBFE1736-2F33-5309-9954-875B11E9519D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// > If MaxResults and NextToken are sued to query results by page, ignore this parameter.
	//
	// example:
	//
	// 6
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGrantVSwitchEnisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGrantVSwitchEnisResponseBody) GoString() string {
	return s.String()
}

func (s *ListGrantVSwitchEnisResponseBody) GetGrantVSwitchEnis() []*ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis {
	return s.GrantVSwitchEnis
}

func (s *ListGrantVSwitchEnisResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListGrantVSwitchEnisResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGrantVSwitchEnisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGrantVSwitchEnisResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListGrantVSwitchEnisResponseBody) SetGrantVSwitchEnis(v []*ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) *ListGrantVSwitchEnisResponseBody {
	s.GrantVSwitchEnis = v
	return s
}

func (s *ListGrantVSwitchEnisResponseBody) SetMaxResults(v int64) *ListGrantVSwitchEnisResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGrantVSwitchEnisResponseBody) SetNextToken(v string) *ListGrantVSwitchEnisResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGrantVSwitchEnisResponseBody) SetRequestId(v string) *ListGrantVSwitchEnisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGrantVSwitchEnisResponseBody) SetTotalCount(v string) *ListGrantVSwitchEnisResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGrantVSwitchEnisResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis struct {
	// The ENI description.
	//
	// example:
	//
	// created by CBN
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ENI ID.
	//
	// example:
	//
	// eni-p0w172vv82kxzb49****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The ENI name.
	//
	// example:
	//
	// my-eni-name
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The primary private IPv4 address of the ENI.
	//
	// example:
	//
	// 192.168.XX.XX
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// Indicates whether the ENI is created by a transit router. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// ENIs that are created by transit routers cannot be used as multicast sources or members.
	//
	// example:
	//
	// false
	TransitRouterFlag *bool `json:"TransitRouterFlag,omitempty" xml:"TransitRouterFlag,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-p0w9s2ig1jnwgrbzl****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-p0w9alkte4w2htrqe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) String() string {
	return dara.Prettify(s)
}

func (s ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) GoString() string {
	return s.String()
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) GetDescription() *string {
	return s.Description
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) GetTransitRouterFlag() *bool {
	return s.TransitRouterFlag
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) SetDescription(v string) *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis {
	s.Description = &v
	return s
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) SetNetworkInterfaceId(v string) *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) SetNetworkInterfaceName(v string) *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis {
	s.NetworkInterfaceName = &v
	return s
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) SetPrimaryIpAddress(v string) *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis {
	s.PrimaryIpAddress = &v
	return s
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) SetTransitRouterFlag(v bool) *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis {
	s.TransitRouterFlag = &v
	return s
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) SetVSwitchId(v string) *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis {
	s.VSwitchId = &v
	return s
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) SetVpcId(v string) *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis {
	s.VpcId = &v
	return s
}

func (s *ListGrantVSwitchEnisResponseBodyGrantVSwitchEnis) Validate() error {
	return dara.Validate(s)
}
