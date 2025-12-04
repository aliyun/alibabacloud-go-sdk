// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpdGrantRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetVpdGrantRuleResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetVpdGrantRuleResponseBody
	GetCode() *int32
	SetContent(v *GetVpdGrantRuleResponseBodyContent) *GetVpdGrantRuleResponseBody
	GetContent() *GetVpdGrantRuleResponseBodyContent
	SetMessage(v string) *GetVpdGrantRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVpdGrantRuleResponseBody
	GetRequestId() *string
}

type GetVpdGrantRuleResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetVpdGrantRuleResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVpdGrantRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVpdGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpdGrantRuleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetVpdGrantRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetVpdGrantRuleResponseBody) GetContent() *GetVpdGrantRuleResponseBodyContent {
	return s.Content
}

func (s *GetVpdGrantRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVpdGrantRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVpdGrantRuleResponseBody) SetAccessDeniedDetail(v string) *GetVpdGrantRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetVpdGrantRuleResponseBody) SetCode(v int32) *GetVpdGrantRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetVpdGrantRuleResponseBody) SetContent(v *GetVpdGrantRuleResponseBodyContent) *GetVpdGrantRuleResponseBody {
	s.Content = v
	return s
}

func (s *GetVpdGrantRuleResponseBody) SetMessage(v string) *GetVpdGrantRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetVpdGrantRuleResponseBody) SetRequestId(v string) *GetVpdGrantRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVpdGrantRuleResponseBodyContent struct {
	// The time when the data address was created.
	//
	// example:
	//
	// 1648085472000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorized Resource ID
	//
	// example:
	//
	// grant-rule-xxxxxx
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1620939556166277
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Network Instance ID
	//
	// example:
	//
	// vpd-xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Network Instance Name
	//
	// example:
	//
	// vpd-lingjun
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Network Product Code:
	//
	// 	- **VPD**: Lingjun CIDR block
	//
	// 	- **VCC**: Lingjun Connection
	//
	// example:
	//
	// VPD
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aek2l4sq6l7u***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Whether the current authorization information has been used; default is false
	//
	// example:
	//
	// 0
	Used *bool `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s GetVpdGrantRuleResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetVpdGrantRuleResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetVpdGrantRuleResponseBodyContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetVpdGrantRuleResponseBodyContent) GetErId() *string {
	return s.ErId
}

func (s *GetVpdGrantRuleResponseBodyContent) GetGrantRuleId() *string {
	return s.GrantRuleId
}

func (s *GetVpdGrantRuleResponseBodyContent) GetGrantTenantId() *string {
	return s.GrantTenantId
}

func (s *GetVpdGrantRuleResponseBodyContent) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetVpdGrantRuleResponseBodyContent) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetVpdGrantRuleResponseBodyContent) GetProduct() *string {
	return s.Product
}

func (s *GetVpdGrantRuleResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpdGrantRuleResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetVpdGrantRuleResponseBodyContent) GetTenantId() *string {
	return s.TenantId
}

func (s *GetVpdGrantRuleResponseBodyContent) GetUsed() *bool {
	return s.Used
}

func (s *GetVpdGrantRuleResponseBodyContent) SetCreateTime(v string) *GetVpdGrantRuleResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetErId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.ErId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetGrantRuleId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.GrantRuleId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetGrantTenantId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.GrantTenantId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetInstanceId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.InstanceId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetInstanceName(v string) *GetVpdGrantRuleResponseBodyContent {
	s.InstanceName = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetProduct(v string) *GetVpdGrantRuleResponseBodyContent {
	s.Product = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetRegionId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetResourceGroupId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetTenantId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetUsed(v bool) *GetVpdGrantRuleResponseBodyContent {
	s.Used = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
