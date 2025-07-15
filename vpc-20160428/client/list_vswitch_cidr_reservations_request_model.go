// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVSwitchCidrReservationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpVersion(v string) *ListVSwitchCidrReservationsRequest
	GetIpVersion() *string
	SetMaxResults(v int64) *ListVSwitchCidrReservationsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListVSwitchCidrReservationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListVSwitchCidrReservationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListVSwitchCidrReservationsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListVSwitchCidrReservationsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListVSwitchCidrReservationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListVSwitchCidrReservationsRequest
	GetResourceOwnerId() *int64
	SetTags(v []*ListVSwitchCidrReservationsRequestTags) *ListVSwitchCidrReservationsRequest
	GetTags() []*ListVSwitchCidrReservationsRequestTags
	SetVSwitchCidrReservationIds(v []*string) *ListVSwitchCidrReservationsRequest
	GetVSwitchCidrReservationIds() []*string
	SetVSwitchCidrReservationType(v string) *ListVSwitchCidrReservationsRequest
	GetVSwitchCidrReservationType() *string
	SetVSwitchId(v string) *ListVSwitchCidrReservationsRequest
	GetVSwitchId() *string
}

type ListVSwitchCidrReservationsRequest struct {
	// The IP version of the reserved CIDR block. Valid values:
	//
	// 	- **IPv4*	- (default)
	//
	// 	- **IPv6**
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the vSwitch.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tags []*ListVSwitchCidrReservationsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the reserved CIDR block. You can specify at most 10 IDs.
	VSwitchCidrReservationIds []*string `json:"VSwitchCidrReservationIds,omitempty" xml:"VSwitchCidrReservationIds,omitempty" type:"Repeated"`
	// The type of the reserved CIDR block. Set the value to **prefix**.
	//
	// >  When you allocate CIDR blocks, or enable the service to automatically allocate CIDR blocks to elastic network interfaces (ENIs), the CIDR blocks to allocate must fall into the reserved CIDR block. If the reserved CIDR is exhausted, an error message is returned.
	//
	// example:
	//
	// prefix
	VSwitchCidrReservationType *string `json:"VSwitchCidrReservationType,omitempty" xml:"VSwitchCidrReservationType,omitempty"`
	// The ID of the vSwitch for which you want to query reserved CIDR blocks.
	//
	// example:
	//
	// vsw-25navfgbue4g****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ListVSwitchCidrReservationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVSwitchCidrReservationsRequest) GoString() string {
	return s.String()
}

func (s *ListVSwitchCidrReservationsRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListVSwitchCidrReservationsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListVSwitchCidrReservationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVSwitchCidrReservationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListVSwitchCidrReservationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListVSwitchCidrReservationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVSwitchCidrReservationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListVSwitchCidrReservationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListVSwitchCidrReservationsRequest) GetTags() []*ListVSwitchCidrReservationsRequestTags {
	return s.Tags
}

func (s *ListVSwitchCidrReservationsRequest) GetVSwitchCidrReservationIds() []*string {
	return s.VSwitchCidrReservationIds
}

func (s *ListVSwitchCidrReservationsRequest) GetVSwitchCidrReservationType() *string {
	return s.VSwitchCidrReservationType
}

func (s *ListVSwitchCidrReservationsRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListVSwitchCidrReservationsRequest) SetIpVersion(v string) *ListVSwitchCidrReservationsRequest {
	s.IpVersion = &v
	return s
}

func (s *ListVSwitchCidrReservationsRequest) SetMaxResults(v int64) *ListVSwitchCidrReservationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVSwitchCidrReservationsRequest) SetNextToken(v string) *ListVSwitchCidrReservationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVSwitchCidrReservationsRequest) SetOwnerAccount(v string) *ListVSwitchCidrReservationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListVSwitchCidrReservationsRequest) SetOwnerId(v int64) *ListVSwitchCidrReservationsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVSwitchCidrReservationsRequest) SetRegionId(v string) *ListVSwitchCidrReservationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVSwitchCidrReservationsRequest) SetResourceOwnerAccount(v string) *ListVSwitchCidrReservationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListVSwitchCidrReservationsRequest) SetResourceOwnerId(v int64) *ListVSwitchCidrReservationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListVSwitchCidrReservationsRequest) SetTags(v []*ListVSwitchCidrReservationsRequestTags) *ListVSwitchCidrReservationsRequest {
	s.Tags = v
	return s
}

func (s *ListVSwitchCidrReservationsRequest) SetVSwitchCidrReservationIds(v []*string) *ListVSwitchCidrReservationsRequest {
	s.VSwitchCidrReservationIds = v
	return s
}

func (s *ListVSwitchCidrReservationsRequest) SetVSwitchCidrReservationType(v string) *ListVSwitchCidrReservationsRequest {
	s.VSwitchCidrReservationType = &v
	return s
}

func (s *ListVSwitchCidrReservationsRequest) SetVSwitchId(v string) *ListVSwitchCidrReservationsRequest {
	s.VSwitchId = &v
	return s
}

func (s *ListVSwitchCidrReservationsRequest) Validate() error {
	return dara.Validate(s)
}

type ListVSwitchCidrReservationsRequestTags struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, and cannot start with acs: or aliyun. It cannot contain http:// or https://.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVSwitchCidrReservationsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListVSwitchCidrReservationsRequestTags) GoString() string {
	return s.String()
}

func (s *ListVSwitchCidrReservationsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListVSwitchCidrReservationsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListVSwitchCidrReservationsRequestTags) SetKey(v string) *ListVSwitchCidrReservationsRequestTags {
	s.Key = &v
	return s
}

func (s *ListVSwitchCidrReservationsRequestTags) SetValue(v string) *ListVSwitchCidrReservationsRequestTags {
	s.Value = &v
	return s
}

func (s *ListVSwitchCidrReservationsRequestTags) Validate() error {
	return dara.Validate(s)
}
