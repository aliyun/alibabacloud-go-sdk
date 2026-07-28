// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicIpAddressPoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListPublicIpAddressPoolsResponseBody
	GetNextToken() *string
	SetPublicIpAddressPoolList(v []*ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) *ListPublicIpAddressPoolsResponseBody
	GetPublicIpAddressPoolList() []*ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList
	SetRequestId(v string) *ListPublicIpAddressPoolsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPublicIpAddressPoolsResponseBody
	GetTotalCount() *int32
}

type ListPublicIpAddressPoolsResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// - If **NextToken*	- is empty, no subsequent request is to be sent.
	//
	// - If **NextToken*	- is returned, the value indicates the token for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of IP address pool instances.
	PublicIpAddressPoolList []*ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList `json:"PublicIpAddressPoolList,omitempty" xml:"PublicIpAddressPoolList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned under the current request conditions.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublicIpAddressPoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublicIpAddressPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicIpAddressPoolsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPublicIpAddressPoolsResponseBody) GetPublicIpAddressPoolList() []*ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	return s.PublicIpAddressPoolList
}

func (s *ListPublicIpAddressPoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublicIpAddressPoolsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPublicIpAddressPoolsResponseBody) SetNextToken(v string) *ListPublicIpAddressPoolsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBody) SetPublicIpAddressPoolList(v []*ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) *ListPublicIpAddressPoolsResponseBody {
	s.PublicIpAddressPoolList = v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBody) SetRequestId(v string) *ListPublicIpAddressPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBody) SetTotalCount(v int32) *ListPublicIpAddressPoolsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBody) Validate() error {
	if s.PublicIpAddressPoolList != nil {
		for _, item := range s.PublicIpAddressPoolList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList struct {
	// The business type of the IP address pool.
	//
	// - **CloudBox**: CloudBox. Only CloudBox users support this type.
	//
	// - **Default*	- (default): default, indicating a non-special type.
	//
	// example:
	//
	// CloudBox
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The business status of the IP address pool instance.
	//
	// - **Normal**: normal.
	//
	// - **FinancialLocked**: locked.
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The creation time, in the format of `YYYY-MM-DDThh:mm:ssZ`.
	//
	// example:
	//
	// 2022-05-10T01:37:38Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the IP address pool instance.
	//
	// example:
	//
	// AddressPoolDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether idle IP addresses are available.
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// true
	IpAddressRemaining *bool `json:"IpAddressRemaining,omitempty" xml:"IpAddressRemaining,omitempty"`
	// The line type.
	//
	// - **BGP**: BGP (multi-ISP) line.
	//
	// - **BGP_PRO**: BGP (multi-ISP) Pro line.
	//
	// For more information about BGP (multi-ISP) lines and BGP (multi-ISP) Pro lines, see [EIP line types](https://help.aliyun.com/document_detail/32321.html).
	//
	// If you are a whitelist user of single-ISP bandwidth, the returned type may also be:
	//
	// - **ChinaTelecom**: China Telecom
	//
	// - **ChinaUnicom**: China Unicom
	//
	// - **ChinaMobile**: China Mobile
	//
	// - **ChinaTelecom_L2**: China Telecom L2
	//
	// - **ChinaUnicom_L2**: China Unicom L2
	//
	// - **ChinaMobile_L2**: China Mobile L2
	//
	// If you are a China (Hangzhou) Finance Cloud user, **BGP_FinanceCloud*	- is returned.
	//
	// example:
	//
	// BGP
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The name of the IP address pool instance.
	//
	// example:
	//
	// AddressPoolName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud account to which the IP address pool belongs.
	//
	// example:
	//
	// 121012345612*****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The instance ID of the IP address pool.
	//
	// example:
	//
	// pippool-6wetvn6fumkgycssx****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// The region ID of the IP address pool.
	//
	// example:
	//
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the IP address pool belongs.
	//
	// example:
	//
	// rg-acfmxazb4pcdvf****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security protection level.
	//
	// - If this parameter is empty, the default value is Anti-DDoS Basic.
	//
	// - If the value is **AntiDDoS_Enhanced**, it indicates Anti-DDoS (Enhanced).
	SecurityProtectionTypes []*string `json:"SecurityProtectionTypes,omitempty" xml:"SecurityProtectionTypes,omitempty" type:"Repeated"`
	// The sharing type of the IP address pool.
	//
	// - **Shared**: The IP address pool is a shared IP address pool.
	//
	// - Empty: The IP address pool is not a shared IP address pool.
	//
	// example:
	//
	// Shared
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The instance status of the IPAM pool.
	//
	// - **Created**: active.
	//
	// - **Deleting**: being deleted.
	//
	// - **Modifying**: being modified.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The total number of available IP addresses in the public IP address pool.
	//
	// example:
	//
	// 100
	TotalIpNum *int32 `json:"TotalIpNum,omitempty" xml:"TotalIpNum,omitempty"`
	// The number of used IP addresses in the public IP address pool.
	//
	// example:
	//
	// 20
	UsedIpNum *int32 `json:"UsedIpNum,omitempty" xml:"UsedIpNum,omitempty"`
	// The type of the user. Valid values:
	//
	// - **admin**: administrator. An administrator can delete, modify, and query IP address pools, and allocate elastic IP addresses (EIPs) from IP address pools.
	//
	// - **user**: regular user. A regular user can only allocate EIPs from IP address pools and query IP address pools, but cannot modify or delete IP address pools.
	//
	// example:
	//
	// admin
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	// The zones of the IP address pool.
	//
	// This parameter is returned only when the business type of the IP address pool is CloudBox.
	Zones []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) String() string {
	return dara.Prettify(s)
}

func (s ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GoString() string {
	return s.String()
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetBizType() *string {
	return s.BizType
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetDescription() *string {
	return s.Description
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetIpAddressRemaining() *bool {
	return s.IpAddressRemaining
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetIsp() *string {
	return s.Isp
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetName() *string {
	return s.Name
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetPublicIpAddressPoolId() *string {
	return s.PublicIpAddressPoolId
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetSecurityProtectionTypes() []*string {
	return s.SecurityProtectionTypes
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetShareType() *string {
	return s.ShareType
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetStatus() *string {
	return s.Status
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetTags() []*ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolListTags {
	return s.Tags
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetTotalIpNum() *int32 {
	return s.TotalIpNum
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetUsedIpNum() *int32 {
	return s.UsedIpNum
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetUserType() *string {
	return s.UserType
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) GetZones() []*string {
	return s.Zones
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetBizType(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.BizType = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetBusinessStatus(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.BusinessStatus = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetCreationTime(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.CreationTime = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetDescription(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.Description = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetIpAddressRemaining(v bool) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.IpAddressRemaining = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetIsp(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.Isp = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetName(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.Name = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetOwnerId(v int64) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.OwnerId = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetPublicIpAddressPoolId(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.PublicIpAddressPoolId = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetRegionId(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.RegionId = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetResourceGroupId(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetSecurityProtectionTypes(v []*string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.SecurityProtectionTypes = v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetShareType(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.ShareType = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetStatus(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.Status = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetTags(v []*ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolListTags) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.Tags = v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetTotalIpNum(v int32) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.TotalIpNum = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetUsedIpNum(v int32) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.UsedIpNum = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetUserType(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.UserType = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) SetZones(v []*string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList {
	s.Zones = v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolList) Validate() error {
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

type ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolListTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceDept
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolListTags) String() string {
	return dara.Prettify(s)
}

func (s ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolListTags) GoString() string {
	return s.String()
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolListTags) GetKey() *string {
	return s.Key
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolListTags) GetValue() *string {
	return s.Value
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolListTags) SetKey(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolListTags {
	s.Key = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolListTags) SetValue(v string) *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolListTags {
	s.Value = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponseBodyPublicIpAddressPoolListTags) Validate() error {
	return dara.Validate(s)
}
