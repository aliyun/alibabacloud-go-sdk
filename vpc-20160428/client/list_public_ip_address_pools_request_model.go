// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicIpAddressPoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *ListPublicIpAddressPoolsRequest
	GetDryRun() *bool
	SetIsp(v string) *ListPublicIpAddressPoolsRequest
	GetIsp() *string
	SetMaxResults(v int32) *ListPublicIpAddressPoolsRequest
	GetMaxResults() *int32
	SetName(v string) *ListPublicIpAddressPoolsRequest
	GetName() *string
	SetNextToken(v string) *ListPublicIpAddressPoolsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListPublicIpAddressPoolsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListPublicIpAddressPoolsRequest
	GetOwnerId() *int64
	SetPublicIpAddressPoolIds(v []*string) *ListPublicIpAddressPoolsRequest
	GetPublicIpAddressPoolIds() []*string
	SetRegionId(v string) *ListPublicIpAddressPoolsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListPublicIpAddressPoolsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ListPublicIpAddressPoolsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListPublicIpAddressPoolsRequest
	GetResourceOwnerId() *int64
	SetSecurityProtectionEnabled(v bool) *ListPublicIpAddressPoolsRequest
	GetSecurityProtectionEnabled() *bool
	SetStatus(v string) *ListPublicIpAddressPoolsRequest
	GetStatus() *string
	SetTags(v []*ListPublicIpAddressPoolsRequestTags) *ListPublicIpAddressPoolsRequest
	GetTags() []*ListPublicIpAddressPoolsRequestTags
}

type ListPublicIpAddressPoolsRequest struct {
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**(default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The line type. Valid values:
	//
	// 	- **BGP*	- (default): BGP (Multi-ISP) line
	//
	// 	- **BGP_PRO**: BGP (Multi-ISP) Pro line
	//
	// For more information about the BGP (Multi-ISP) line and BGP (Multi-ISP) Pro line, see the "Line types" section of [What is EIP?](https://help.aliyun.com/document_detail/32321.html)
	//
	// If you are allowed to use single-ISP bandwidth, you can also choose one of the following values:
	//
	// 	- **ChinaTelecom**
	//
	// 	- **ChinaUnicom**
	//
	// 	- **ChinaMobile**
	//
	// 	- **ChinaTelecom_L2**
	//
	// 	- **ChinaUnicom_L2**
	//
	// 	- **ChinaMobile_L2**
	//
	// If your services are deployed in China East 1 Finance, this parameter is required and you must set the parameter to **BGP_FinanceCloud**.
	//
	// example:
	//
	// BGP
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The maximum number of entries to return. Valid values: **10*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the IP address pool.
	//
	// If you enter a name, the name must be 1 to 128 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// AddressPoolName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
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
	// The IDs of the IP address pool.
	//
	// You can enter up to 100 IDs.
	PublicIpAddressPoolIds []*string `json:"PublicIpAddressPoolIds,omitempty" xml:"PublicIpAddressPoolIds,omitempty" type:"Repeated"`
	// The ID of the region in which the IP address pool that you want to query resides.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
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
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable Anti-DDoS Pro/Premium. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	SecurityProtectionEnabled *bool `json:"SecurityProtectionEnabled,omitempty" xml:"SecurityProtectionEnabled,omitempty"`
	// The status of the IP address pool. Valid values:
	//
	// 	- **Created**
	//
	// 	- **Deleting**
	//
	// 	- **Modifying**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags to add to the resource.
	Tags []*ListPublicIpAddressPoolsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListPublicIpAddressPoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublicIpAddressPoolsRequest) GoString() string {
	return s.String()
}

func (s *ListPublicIpAddressPoolsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListPublicIpAddressPoolsRequest) GetIsp() *string {
	return s.Isp
}

func (s *ListPublicIpAddressPoolsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPublicIpAddressPoolsRequest) GetName() *string {
	return s.Name
}

func (s *ListPublicIpAddressPoolsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPublicIpAddressPoolsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListPublicIpAddressPoolsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListPublicIpAddressPoolsRequest) GetPublicIpAddressPoolIds() []*string {
	return s.PublicIpAddressPoolIds
}

func (s *ListPublicIpAddressPoolsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPublicIpAddressPoolsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPublicIpAddressPoolsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListPublicIpAddressPoolsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListPublicIpAddressPoolsRequest) GetSecurityProtectionEnabled() *bool {
	return s.SecurityProtectionEnabled
}

func (s *ListPublicIpAddressPoolsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListPublicIpAddressPoolsRequest) GetTags() []*ListPublicIpAddressPoolsRequestTags {
	return s.Tags
}

func (s *ListPublicIpAddressPoolsRequest) SetDryRun(v bool) *ListPublicIpAddressPoolsRequest {
	s.DryRun = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetIsp(v string) *ListPublicIpAddressPoolsRequest {
	s.Isp = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetMaxResults(v int32) *ListPublicIpAddressPoolsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetName(v string) *ListPublicIpAddressPoolsRequest {
	s.Name = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetNextToken(v string) *ListPublicIpAddressPoolsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetOwnerAccount(v string) *ListPublicIpAddressPoolsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetOwnerId(v int64) *ListPublicIpAddressPoolsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetPublicIpAddressPoolIds(v []*string) *ListPublicIpAddressPoolsRequest {
	s.PublicIpAddressPoolIds = v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetRegionId(v string) *ListPublicIpAddressPoolsRequest {
	s.RegionId = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetResourceGroupId(v string) *ListPublicIpAddressPoolsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetResourceOwnerAccount(v string) *ListPublicIpAddressPoolsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetResourceOwnerId(v int64) *ListPublicIpAddressPoolsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetSecurityProtectionEnabled(v bool) *ListPublicIpAddressPoolsRequest {
	s.SecurityProtectionEnabled = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetStatus(v string) *ListPublicIpAddressPoolsRequest {
	s.Status = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) SetTags(v []*ListPublicIpAddressPoolsRequestTags) *ListPublicIpAddressPoolsRequest {
	s.Tags = v
	return s
}

func (s *ListPublicIpAddressPoolsRequest) Validate() error {
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

type ListPublicIpAddressPoolsRequestTags struct {
	// The tag key to add to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPublicIpAddressPoolsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListPublicIpAddressPoolsRequestTags) GoString() string {
	return s.String()
}

func (s *ListPublicIpAddressPoolsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListPublicIpAddressPoolsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListPublicIpAddressPoolsRequestTags) SetKey(v string) *ListPublicIpAddressPoolsRequestTags {
	s.Key = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequestTags) SetValue(v string) *ListPublicIpAddressPoolsRequestTags {
	s.Value = &v
	return s
}

func (s *ListPublicIpAddressPoolsRequestTags) Validate() error {
	return dara.Validate(s)
}
