// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpdGrantRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListVpdGrantRulesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListVpdGrantRulesResponseBody
	GetCode() *int32
	SetContent(v *ListVpdGrantRulesResponseBodyContent) *ListVpdGrantRulesResponseBody
	GetContent() *ListVpdGrantRulesResponseBodyContent
	SetMessage(v string) *ListVpdGrantRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVpdGrantRulesResponseBody
	GetRequestId() *string
}

type ListVpdGrantRulesResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListVpdGrantRulesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// A56F7D3C-8850-5AF4-A342-2D71C9A9D1CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVpdGrantRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpdGrantRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpdGrantRulesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListVpdGrantRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListVpdGrantRulesResponseBody) GetContent() *ListVpdGrantRulesResponseBodyContent {
	return s.Content
}

func (s *ListVpdGrantRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVpdGrantRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpdGrantRulesResponseBody) SetAccessDeniedDetail(v string) *ListVpdGrantRulesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListVpdGrantRulesResponseBody) SetCode(v int32) *ListVpdGrantRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVpdGrantRulesResponseBody) SetContent(v *ListVpdGrantRulesResponseBodyContent) *ListVpdGrantRulesResponseBody {
	s.Content = v
	return s
}

func (s *ListVpdGrantRulesResponseBody) SetMessage(v string) *ListVpdGrantRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVpdGrantRulesResponseBody) SetRequestId(v string) *ListVpdGrantRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVpdGrantRulesResponseBodyContent struct {
	// Lingjun CIDR block authorization information list
	Data []*ListVpdGrantRulesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVpdGrantRulesResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListVpdGrantRulesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVpdGrantRulesResponseBodyContent) GetData() []*ListVpdGrantRulesResponseBodyContentData {
	return s.Data
}

func (s *ListVpdGrantRulesResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListVpdGrantRulesResponseBodyContent) SetData(v []*ListVpdGrantRulesResponseBodyContentData) *ListVpdGrantRulesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContent) SetTotal(v int64) *ListVpdGrantRulesResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContent) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpdGrantRulesResponseBodyContentData struct {
	// The time when the data address was created.
	//
	// example:
	//
	// 1643013506000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The current network sample is authorized to the specified Lingjun HUB sample ID.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorization Entry ID
	//
	// example:
	//
	// grant-rule-8rgvqazb
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// The ID of the tenant to which the current instance is authorized.
	//
	// example:
	//
	// 1672372231790
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-8rgvqazb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the ECU.
	//
	// example:
	//
	// vpd-1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the authorized product. Valid values:
	//
	// 	- **VPD**: indicates a VPD instance of the Lingjun network segment.
	//
	// 	- **VCC**: indicates that Lingjun connects to the VCC instance.
	//
	// The caller does not need to specify.
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
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Whether the current authorized instance has been bound
	//
	// example:
	//
	// true
	Used *bool `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListVpdGrantRulesResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListVpdGrantRulesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVpdGrantRulesResponseBodyContentData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVpdGrantRulesResponseBodyContentData) GetErId() *string {
	return s.ErId
}

func (s *ListVpdGrantRulesResponseBodyContentData) GetGrantRuleId() *string {
	return s.GrantRuleId
}

func (s *ListVpdGrantRulesResponseBodyContentData) GetGrantTenantId() *string {
	return s.GrantTenantId
}

func (s *ListVpdGrantRulesResponseBodyContentData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVpdGrantRulesResponseBodyContentData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListVpdGrantRulesResponseBodyContentData) GetProduct() *string {
	return s.Product
}

func (s *ListVpdGrantRulesResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpdGrantRulesResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpdGrantRulesResponseBodyContentData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVpdGrantRulesResponseBodyContentData) GetUsed() *bool {
	return s.Used
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetCreateTime(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetErId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.ErId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetGrantRuleId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.GrantRuleId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetGrantTenantId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.GrantTenantId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetInstanceId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.InstanceId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetInstanceName(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.InstanceName = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetProduct(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.Product = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetRegionId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetResourceGroupId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetTenantId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetUsed(v bool) *ListVpdGrantRulesResponseBodyContentData {
	s.Used = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) Validate() error {
	return dara.Validate(s)
}
