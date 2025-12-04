// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVccGrantRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetVccGrantRuleResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetVccGrantRuleResponseBody
	GetCode() *int32
	SetContent(v *GetVccGrantRuleResponseBodyContent) *GetVccGrantRuleResponseBody
	GetContent() *GetVccGrantRuleResponseBodyContent
	SetMessage(v string) *GetVccGrantRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVccGrantRuleResponseBody
	GetRequestId() *string
}

type GetVccGrantRuleResponseBody struct {
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
	// The returned data.
	Content *GetVccGrantRuleResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVccGrantRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVccGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetVccGrantRuleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetVccGrantRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetVccGrantRuleResponseBody) GetContent() *GetVccGrantRuleResponseBodyContent {
	return s.Content
}

func (s *GetVccGrantRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVccGrantRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVccGrantRuleResponseBody) SetAccessDeniedDetail(v string) *GetVccGrantRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetVccGrantRuleResponseBody) SetCode(v int32) *GetVccGrantRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetVccGrantRuleResponseBody) SetContent(v *GetVccGrantRuleResponseBodyContent) *GetVccGrantRuleResponseBody {
	s.Content = v
	return s
}

func (s *GetVccGrantRuleResponseBody) SetMessage(v string) *GetVccGrantRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetVccGrantRuleResponseBody) SetRequestId(v string) *GetVccGrantRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVccGrantRuleResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVccGrantRuleResponseBodyContent struct {
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
	// er-aueyxxsy
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorized Resource ID
	//
	// example:
	//
	// grant-rule-jaj34d75h01
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
	// vcc-cn-jaj34d75h01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Network Instance Name
	//
	// example:
	//
	// vcc-1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Network Product Code:
	//
	// 	- **VPD**: Lingjun CIDR block
	//
	// 	- **VCC**: Lingjun Connection
	//
	// example:
	//
	// VCC
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
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166279
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Whether the current authorization information has been used; optional values:
	//
	// 	- **true**: Used
	//
	// 	- **false**: Not used
	//
	// example:
	//
	// false
	Used *bool `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s GetVccGrantRuleResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetVccGrantRuleResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetVccGrantRuleResponseBodyContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetVccGrantRuleResponseBodyContent) GetErId() *string {
	return s.ErId
}

func (s *GetVccGrantRuleResponseBodyContent) GetGrantRuleId() *string {
	return s.GrantRuleId
}

func (s *GetVccGrantRuleResponseBodyContent) GetGrantTenantId() *string {
	return s.GrantTenantId
}

func (s *GetVccGrantRuleResponseBodyContent) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetVccGrantRuleResponseBodyContent) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetVccGrantRuleResponseBodyContent) GetProduct() *string {
	return s.Product
}

func (s *GetVccGrantRuleResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVccGrantRuleResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetVccGrantRuleResponseBodyContent) GetTenantId() *string {
	return s.TenantId
}

func (s *GetVccGrantRuleResponseBodyContent) GetUsed() *bool {
	return s.Used
}

func (s *GetVccGrantRuleResponseBodyContent) SetCreateTime(v string) *GetVccGrantRuleResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetErId(v string) *GetVccGrantRuleResponseBodyContent {
	s.ErId = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetGrantRuleId(v string) *GetVccGrantRuleResponseBodyContent {
	s.GrantRuleId = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetGrantTenantId(v string) *GetVccGrantRuleResponseBodyContent {
	s.GrantTenantId = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetInstanceId(v string) *GetVccGrantRuleResponseBodyContent {
	s.InstanceId = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetInstanceName(v string) *GetVccGrantRuleResponseBodyContent {
	s.InstanceName = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetProduct(v string) *GetVccGrantRuleResponseBodyContent {
	s.Product = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetRegionId(v string) *GetVccGrantRuleResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetResourceGroupId(v string) *GetVccGrantRuleResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetTenantId(v string) *GetVccGrantRuleResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetUsed(v bool) *GetVccGrantRuleResponseBodyContent {
	s.Used = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
