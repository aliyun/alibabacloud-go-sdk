// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountStatus(v string) *DescribeMemberAccountsRequest
	GetAccountStatus() *string
	SetInstanceId(v string) *DescribeMemberAccountsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeMemberAccountsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeMemberAccountsRequest
	GetResourceManagerResourceGroupId() *string
	SetSourceIp(v string) *DescribeMemberAccountsRequest
	GetSourceIp() *string
}

type DescribeMemberAccountsRequest struct {
	// The status of the member that you want to query.
	//
	// 	- **enabled**: managed.
	//
	// 	- **disabled**: not managed.
	//
	// 	- **disabling**: being deleted.
	//
	// example:
	//
	// enabled
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The source IP address of the request. The system specifies this parameter.
	//
	// example:
	//
	// 0.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeMemberAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMemberAccountsRequest) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *DescribeMemberAccountsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMemberAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMemberAccountsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeMemberAccountsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeMemberAccountsRequest) SetAccountStatus(v string) *DescribeMemberAccountsRequest {
	s.AccountStatus = &v
	return s
}

func (s *DescribeMemberAccountsRequest) SetInstanceId(v string) *DescribeMemberAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMemberAccountsRequest) SetRegionId(v string) *DescribeMemberAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMemberAccountsRequest) SetResourceManagerResourceGroupId(v string) *DescribeMemberAccountsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeMemberAccountsRequest) SetSourceIp(v string) *DescribeMemberAccountsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeMemberAccountsRequest) Validate() error {
	return dara.Validate(s)
}
