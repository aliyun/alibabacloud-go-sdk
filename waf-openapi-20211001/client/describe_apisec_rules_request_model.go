// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeApisecRulesRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeApisecRulesRequest
	GetLang() *string
	SetLevel(v string) *DescribeApisecRulesRequest
	GetLevel() *string
	SetName(v string) *DescribeApisecRulesRequest
	GetName() *string
	SetOrigin(v string) *DescribeApisecRulesRequest
	GetOrigin() *string
	SetPageNumber(v int64) *DescribeApisecRulesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeApisecRulesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeApisecRulesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecRulesRequest
	GetResourceManagerResourceGroupId() *string
	SetStatus(v int64) *DescribeApisecRulesRequest
	GetStatus() *int64
	SetType(v string) *DescribeApisecRulesRequest
	GetType() *string
}

type DescribeApisecRulesRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0x***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The level of the policy.
	//
	// If Type is set to risk or event, you can set this parameter to one of the following values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// If Type is set to sensitive_word, you can set this parameter to one of the following values:
	//
	// 	- **S1**
	//
	// 	- **S2**
	//
	// 	- **S3**
	//
	// 	- **S4**
	//
	// example:
	//
	// high
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// Information Leak
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source of the policy. Valid values:
	//
	// 	- **custom**
	//
	// 	- **default**
	//
	// example:
	//
	// custom
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
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
	// The status of the policy. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **risk**: risk detection
	//
	// 	- **event**: security event
	//
	// 	- **sensitive_word**: sensitive data
	//
	// 	- **auth_flag**: authentication credential
	//
	// 	- **api_tag**: business purpose
	//
	// 	- **desensitization**: masking
	//
	// 	- **whitelist**: whitelist
	//
	// 	- **recognition**: API recognition
	//
	// 	- **offline_api**: lifecycle management
	//
	// This parameter is required.
	//
	// example:
	//
	// risk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApisecRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeApisecRulesRequest) GetLevel() *string {
	return s.Level
}

func (s *DescribeApisecRulesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeApisecRulesRequest) GetOrigin() *string {
	return s.Origin
}

func (s *DescribeApisecRulesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeApisecRulesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeApisecRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecRulesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecRulesRequest) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeApisecRulesRequest) GetType() *string {
	return s.Type
}

func (s *DescribeApisecRulesRequest) SetInstanceId(v string) *DescribeApisecRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecRulesRequest) SetLang(v string) *DescribeApisecRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeApisecRulesRequest) SetLevel(v string) *DescribeApisecRulesRequest {
	s.Level = &v
	return s
}

func (s *DescribeApisecRulesRequest) SetName(v string) *DescribeApisecRulesRequest {
	s.Name = &v
	return s
}

func (s *DescribeApisecRulesRequest) SetOrigin(v string) *DescribeApisecRulesRequest {
	s.Origin = &v
	return s
}

func (s *DescribeApisecRulesRequest) SetPageNumber(v int64) *DescribeApisecRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisecRulesRequest) SetPageSize(v int64) *DescribeApisecRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisecRulesRequest) SetRegionId(v string) *DescribeApisecRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecRulesRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecRulesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecRulesRequest) SetStatus(v int64) *DescribeApisecRulesRequest {
	s.Status = &v
	return s
}

func (s *DescribeApisecRulesRequest) SetType(v string) *DescribeApisecRulesRequest {
	s.Type = &v
	return s
}

func (s *DescribeApisecRulesRequest) Validate() error {
	return dara.Validate(s)
}
