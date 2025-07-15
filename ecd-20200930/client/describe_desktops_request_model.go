// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *DescribeDesktopsRequest
	GetChargeType() *string
	SetDesktopGroupId(v string) *DescribeDesktopsRequest
	GetDesktopGroupId() *string
	SetDesktopId(v []*string) *DescribeDesktopsRequest
	GetDesktopId() []*string
	SetDesktopName(v string) *DescribeDesktopsRequest
	GetDesktopName() *string
	SetDesktopStatus(v string) *DescribeDesktopsRequest
	GetDesktopStatus() *string
	SetDesktopStatusList(v []*string) *DescribeDesktopsRequest
	GetDesktopStatusList() []*string
	SetDesktopType(v string) *DescribeDesktopsRequest
	GetDesktopType() *string
	SetDirectoryId(v string) *DescribeDesktopsRequest
	GetDirectoryId() *string
	SetEndUserId(v []*string) *DescribeDesktopsRequest
	GetEndUserId() []*string
	SetExcludedEndUserId(v []*string) *DescribeDesktopsRequest
	GetExcludedEndUserId() []*string
	SetExpiredTime(v string) *DescribeDesktopsRequest
	GetExpiredTime() *string
	SetFillResourceGroup(v bool) *DescribeDesktopsRequest
	GetFillResourceGroup() *bool
	SetFilterDesktopGroup(v bool) *DescribeDesktopsRequest
	GetFilterDesktopGroup() *bool
	SetGpuInstanceGroupId(v string) *DescribeDesktopsRequest
	GetGpuInstanceGroupId() *string
	SetGroupId(v string) *DescribeDesktopsRequest
	GetGroupId() *string
	SetImageId(v []*string) *DescribeDesktopsRequest
	GetImageId() []*string
	SetManagementFlag(v string) *DescribeDesktopsRequest
	GetManagementFlag() *string
	SetMaxResults(v int32) *DescribeDesktopsRequest
	GetMaxResults() *int32
	SetMultiResource(v bool) *DescribeDesktopsRequest
	GetMultiResource() *bool
	SetNextToken(v string) *DescribeDesktopsRequest
	GetNextToken() *string
	SetOfficeSiteId(v string) *DescribeDesktopsRequest
	GetOfficeSiteId() *string
	SetOfficeSiteName(v string) *DescribeDesktopsRequest
	GetOfficeSiteName() *string
	SetOnlyDesktopGroup(v bool) *DescribeDesktopsRequest
	GetOnlyDesktopGroup() *bool
	SetOsTypes(v []*string) *DescribeDesktopsRequest
	GetOsTypes() []*string
	SetPageNumber(v int32) *DescribeDesktopsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDesktopsRequest
	GetPageSize() *int32
	SetPolicyGroupId(v string) *DescribeDesktopsRequest
	GetPolicyGroupId() *string
	SetProtocolType(v string) *DescribeDesktopsRequest
	GetProtocolType() *string
	SetQosRuleId(v string) *DescribeDesktopsRequest
	GetQosRuleId() *string
	SetQueryFotaUpdate(v bool) *DescribeDesktopsRequest
	GetQueryFotaUpdate() *bool
	SetRegionId(v string) *DescribeDesktopsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDesktopsRequest
	GetResourceGroupId() *string
	SetSnapshotPolicyId(v string) *DescribeDesktopsRequest
	GetSnapshotPolicyId() *string
	SetSubPayType(v string) *DescribeDesktopsRequest
	GetSubPayType() *string
	SetTag(v []*DescribeDesktopsRequestTag) *DescribeDesktopsRequest
	GetTag() []*DescribeDesktopsRequestTag
	SetUserName(v string) *DescribeDesktopsRequest
	GetUserName() *string
}

type DescribeDesktopsRequest struct {
	// The billing method of the cloud computer.
	//
	// Valid values:
	//
	// 	- Postpaid (default): pay-as-you-go
	//
	// 	- PrePaid: subscription
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the cloud computer pool. If you specify `OnlyDesktopGroup`, ignore `DesktopGroupId`. If you leave `DesktopId` empty, all IDs of the cloud computers in the cloud computer pool are queried.````
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The cloud computer IDs. You can specify the IDs of 1 to 100 cloud computers.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The cloud computer name.
	//
	// example:
	//
	// testDesktopName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The cloud computer status.
	//
	// Valid values:
	//
	// 	- Stopped
	//
	// 	- Starting
	//
	// 	- Rebuilding
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// 	- Expired
	//
	// 	- Deleted
	//
	// 	- Pending
	//
	// example:
	//
	// Running
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	// The list of cloud computer status.
	DesktopStatusList []*string `json:"DesktopStatusList,omitempty" xml:"DesktopStatusList,omitempty" type:"Repeated"`
	// The cloud computer type. You can call the [DescribeDesktopTypes](https://help.aliyun.com/document_detail/188882.html) operation to query the IDs of all supported types.
	//
	// example:
	//
	// eds.general.2c8g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The directory ID, which is the same as the office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The authorized users of the cloud computer. You can specify the IDs of 1 to 100 users.
	//
	// >  During a specific period of time, only one user can connect to and use the cloud computer.
	//
	// example:
	//
	// alice
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// The list of authorized users that you want to exclude from the cloud computer. You can specify the IDs of 1 to 100 users.
	//
	// example:
	//
	// andy
	ExcludedEndUserId []*string `json:"ExcludedEndUserId,omitempty" xml:"ExcludedEndUserId,omitempty" type:"Repeated"`
	// The time when a subscription cloud computer expires.
	//
	// example:
	//
	// 2022-12-31T15:59:59Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// Specifies whether to query the information about the enterprise resource group.
	//
	// example:
	//
	// true
	FillResourceGroup *bool `json:"FillResourceGroup,omitempty" xml:"FillResourceGroup,omitempty"`
	// Specifies whether to exclude pooled cloud computers.
	//
	// Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// false
	FilterDesktopGroup *bool `json:"FilterDesktopGroup,omitempty" xml:"FilterDesktopGroup,omitempty"`
	// The ID of the elastic GPU pool.
	//
	// example:
	//
	// gp-0bm2iz1v6m6nx****
	GpuInstanceGroupId *string `json:"GpuInstanceGroupId,omitempty" xml:"GpuInstanceGroupId,omitempty"`
	// The ID of the cloud computer pool.
	//
	// example:
	//
	// dg-boyczi8enfyc5****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The IDs of the images.
	ImageId []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	// The flag that is used to manage the cloud desktops.
	//
	// example:
	//
	// NoFlag
	ManagementFlag *string `json:"ManagementFlag,omitempty" xml:"ManagementFlag,omitempty"`
	// The number of entries per page.
	//
	// 	- Maximum value: 100
	//
	// 	- Default value: 10
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies whether the shared group is a multi-cloud computer type.
	//
	// Valid values:
	//
	// - true: a multi-cloud computer type.
	//
	// - false: a single-cloud computer type.
	//
	// example:
	//
	// false
	MultiResource *bool `json:"MultiResource,omitempty" xml:"MultiResource,omitempty"`
	// The token that determines the start point of the next query. If this parameter is left empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The office network name.
	//
	// example:
	//
	// testName
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// Specifies whether to query pooled cloud computers.
	//
	// example:
	//
	// true
	OnlyDesktopGroup *bool `json:"OnlyDesktopGroup,omitempty" xml:"OnlyDesktopGroup,omitempty"`
	// The operating systems (OSs).
	OsTypes []*string `json:"OsTypes,omitempty" xml:"OsTypes,omitempty" type:"Repeated"`
	// The page number.
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
	// The ID of the cloud computer policy.
	//
	// example:
	//
	// system-all-enabled-policy
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The protocol.
	//
	// Valid values:
	//
	// 	- HDX: High-definition Experience (HDX) protocol
	//
	// 	- ASP: in-house Adaptive Streaming Protocol (ASP) (recommended)
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The ID of the network throttling rule.
	//
	// example:
	//
	// qos-5605u0gelk200****
	QosRuleId *string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
	// Specifies whether to query the image update information about the cloud computer.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	QueryFotaUpdate *bool `json:"QueryFotaUpdate,omitempty" xml:"QueryFotaUpdate,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by Elastic Desktop Service (EDS).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the enterprise resource group.
	//
	// example:
	//
	// rg-4hsvzbbmqdzu3s****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the snapshot policy.
	//
	// example:
	//
	// sp-hb12mclyne09xw***
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitempty" xml:"SnapshotPolicyId,omitempty"`
	// The billing method of the cloud computer.
	//
	// Valid values:
	//
	// 	- duration: hourly plan (available for users in the whitelist)
	//
	// 	- postPaid: pay-as-you-go
	//
	// 	- monthPackage: monthly subscription (120-hour or 250-hour computing plan)
	//
	// 	- prePaid: monthly subscription (unlimited-hour computing plan)
	//
	// example:
	//
	// monthPackage
	SubPayType *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	// The tags that you want to add to the cloud computer. A tag is a key-value pair that consists of a tag key and a tag value. Tags are used to identify resources. You can use tags to manage cloud computers by group. This facilitates search and batch operations. For more information, see [Use tags to manage cloud computers](https://help.aliyun.com/document_detail/203781.html).
	Tag []*DescribeDesktopsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the end user.
	//
	// example:
	//
	// Alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDesktopsRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *DescribeDesktopsRequest) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeDesktopsRequest) GetDesktopStatus() *string {
	return s.DesktopStatus
}

func (s *DescribeDesktopsRequest) GetDesktopStatusList() []*string {
	return s.DesktopStatusList
}

func (s *DescribeDesktopsRequest) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribeDesktopsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeDesktopsRequest) GetEndUserId() []*string {
	return s.EndUserId
}

func (s *DescribeDesktopsRequest) GetExcludedEndUserId() []*string {
	return s.ExcludedEndUserId
}

func (s *DescribeDesktopsRequest) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDesktopsRequest) GetFillResourceGroup() *bool {
	return s.FillResourceGroup
}

func (s *DescribeDesktopsRequest) GetFilterDesktopGroup() *bool {
	return s.FilterDesktopGroup
}

func (s *DescribeDesktopsRequest) GetGpuInstanceGroupId() *string {
	return s.GpuInstanceGroupId
}

func (s *DescribeDesktopsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDesktopsRequest) GetImageId() []*string {
	return s.ImageId
}

func (s *DescribeDesktopsRequest) GetManagementFlag() *string {
	return s.ManagementFlag
}

func (s *DescribeDesktopsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDesktopsRequest) GetMultiResource() *bool {
	return s.MultiResource
}

func (s *DescribeDesktopsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopsRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeDesktopsRequest) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *DescribeDesktopsRequest) GetOnlyDesktopGroup() *bool {
	return s.OnlyDesktopGroup
}

func (s *DescribeDesktopsRequest) GetOsTypes() []*string {
	return s.OsTypes
}

func (s *DescribeDesktopsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDesktopsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDesktopsRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeDesktopsRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeDesktopsRequest) GetQosRuleId() *string {
	return s.QosRuleId
}

func (s *DescribeDesktopsRequest) GetQueryFotaUpdate() *bool {
	return s.QueryFotaUpdate
}

func (s *DescribeDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDesktopsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDesktopsRequest) GetSnapshotPolicyId() *string {
	return s.SnapshotPolicyId
}

func (s *DescribeDesktopsRequest) GetSubPayType() *string {
	return s.SubPayType
}

func (s *DescribeDesktopsRequest) GetTag() []*DescribeDesktopsRequestTag {
	return s.Tag
}

func (s *DescribeDesktopsRequest) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDesktopsRequest) SetChargeType(v string) *DescribeDesktopsRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopGroupId(v string) *DescribeDesktopsRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopId(v []*string) *DescribeDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopName(v string) *DescribeDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopStatus(v string) *DescribeDesktopsRequest {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopStatusList(v []*string) *DescribeDesktopsRequest {
	s.DesktopStatusList = v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopType(v string) *DescribeDesktopsRequest {
	s.DesktopType = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDirectoryId(v string) *DescribeDesktopsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetEndUserId(v []*string) *DescribeDesktopsRequest {
	s.EndUserId = v
	return s
}

func (s *DescribeDesktopsRequest) SetExcludedEndUserId(v []*string) *DescribeDesktopsRequest {
	s.ExcludedEndUserId = v
	return s
}

func (s *DescribeDesktopsRequest) SetExpiredTime(v string) *DescribeDesktopsRequest {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopsRequest) SetFillResourceGroup(v bool) *DescribeDesktopsRequest {
	s.FillResourceGroup = &v
	return s
}

func (s *DescribeDesktopsRequest) SetFilterDesktopGroup(v bool) *DescribeDesktopsRequest {
	s.FilterDesktopGroup = &v
	return s
}

func (s *DescribeDesktopsRequest) SetGpuInstanceGroupId(v string) *DescribeDesktopsRequest {
	s.GpuInstanceGroupId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetGroupId(v string) *DescribeDesktopsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetImageId(v []*string) *DescribeDesktopsRequest {
	s.ImageId = v
	return s
}

func (s *DescribeDesktopsRequest) SetManagementFlag(v string) *DescribeDesktopsRequest {
	s.ManagementFlag = &v
	return s
}

func (s *DescribeDesktopsRequest) SetMaxResults(v int32) *DescribeDesktopsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopsRequest) SetMultiResource(v bool) *DescribeDesktopsRequest {
	s.MultiResource = &v
	return s
}

func (s *DescribeDesktopsRequest) SetNextToken(v string) *DescribeDesktopsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsRequest) SetOfficeSiteId(v string) *DescribeDesktopsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetOfficeSiteName(v string) *DescribeDesktopsRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeDesktopsRequest) SetOnlyDesktopGroup(v bool) *DescribeDesktopsRequest {
	s.OnlyDesktopGroup = &v
	return s
}

func (s *DescribeDesktopsRequest) SetOsTypes(v []*string) *DescribeDesktopsRequest {
	s.OsTypes = v
	return s
}

func (s *DescribeDesktopsRequest) SetPageNumber(v int32) *DescribeDesktopsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDesktopsRequest) SetPageSize(v int32) *DescribeDesktopsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDesktopsRequest) SetPolicyGroupId(v string) *DescribeDesktopsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetProtocolType(v string) *DescribeDesktopsRequest {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDesktopsRequest) SetQosRuleId(v string) *DescribeDesktopsRequest {
	s.QosRuleId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetQueryFotaUpdate(v bool) *DescribeDesktopsRequest {
	s.QueryFotaUpdate = &v
	return s
}

func (s *DescribeDesktopsRequest) SetRegionId(v string) *DescribeDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetResourceGroupId(v string) *DescribeDesktopsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetSnapshotPolicyId(v string) *DescribeDesktopsRequest {
	s.SnapshotPolicyId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetSubPayType(v string) *DescribeDesktopsRequest {
	s.SubPayType = &v
	return s
}

func (s *DescribeDesktopsRequest) SetTag(v []*DescribeDesktopsRequestTag) *DescribeDesktopsRequest {
	s.Tag = v
	return s
}

func (s *DescribeDesktopsRequest) SetUserName(v string) *DescribeDesktopsRequest {
	s.UserName = &v
	return s
}

func (s *DescribeDesktopsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopsRequestTag struct {
	// The tag key. If you specify the `Tag` parameter, you must also specify the `Key` parameter. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun` and contain only spaces.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDesktopsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDesktopsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDesktopsRequestTag) SetKey(v string) *DescribeDesktopsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDesktopsRequestTag) SetValue(v string) *DescribeDesktopsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDesktopsRequestTag) Validate() error {
	return dara.Validate(s)
}
