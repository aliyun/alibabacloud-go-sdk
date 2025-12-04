// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVccGrantRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListVccGrantRulesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListVccGrantRulesResponseBody
	GetCode() *int32
	SetContent(v *ListVccGrantRulesResponseBodyContent) *ListVccGrantRulesResponseBody
	GetContent() *ListVccGrantRulesResponseBodyContent
	SetMessage(v string) *ListVccGrantRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVccGrantRulesResponseBody
	GetRequestId() *string
}

type ListVccGrantRulesResponseBody struct {
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
	Content *ListVccGrantRulesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
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

func (s ListVccGrantRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVccGrantRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVccGrantRulesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListVccGrantRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListVccGrantRulesResponseBody) GetContent() *ListVccGrantRulesResponseBodyContent {
	return s.Content
}

func (s *ListVccGrantRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVccGrantRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVccGrantRulesResponseBody) SetAccessDeniedDetail(v string) *ListVccGrantRulesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListVccGrantRulesResponseBody) SetCode(v int32) *ListVccGrantRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVccGrantRulesResponseBody) SetContent(v *ListVccGrantRulesResponseBodyContent) *ListVccGrantRulesResponseBody {
	s.Content = v
	return s
}

func (s *ListVccGrantRulesResponseBody) SetMessage(v string) *ListVccGrantRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVccGrantRulesResponseBody) SetRequestId(v string) *ListVccGrantRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVccGrantRulesResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVccGrantRulesResponseBodyContent struct {
	// List of cross-account authorization information of Lingjun connection
	Data []*ListVccGrantRulesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVccGrantRulesResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListVccGrantRulesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVccGrantRulesResponseBodyContent) GetData() []*ListVccGrantRulesResponseBodyContentData {
	return s.Data
}

func (s *ListVccGrantRulesResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListVccGrantRulesResponseBodyContent) SetData(v []*ListVccGrantRulesResponseBodyContentData) *ListVccGrantRulesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVccGrantRulesResponseBodyContent) SetTotal(v int64) *ListVccGrantRulesResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContent) Validate() error {
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

type ListVccGrantRulesResponseBodyContentData struct {
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Lingjun HUB ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Cross-account authorization information Instance ID
	//
	// example:
	//
	// grant-rule-jpumgwvp
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1013666993027780
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Network Instance ID
	//
	// example:
	//
	// vcc-cn-jaj33d1kb05
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the ECU.
	//
	// example:
	//
	// vcc-1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the authorized product. Valid values:
	//
	// 	- **VPD**: indicates a VPD instance of the Lingjun network segment.
	//
	// 	- **VCC**: indicates that Lingjun connects to the VCC instance.
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
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Whether the current cross-account resource has been bound to the cross-account Lingjun HUB. Valid values:
	//
	// 	- **true**: Used
	//
	// 	- **false**: Not used
	//
	// example:
	//
	// true
	Used *bool `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListVccGrantRulesResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListVccGrantRulesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVccGrantRulesResponseBodyContentData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVccGrantRulesResponseBodyContentData) GetErId() *string {
	return s.ErId
}

func (s *ListVccGrantRulesResponseBodyContentData) GetGrantRuleId() *string {
	return s.GrantRuleId
}

func (s *ListVccGrantRulesResponseBodyContentData) GetGrantTenantId() *string {
	return s.GrantTenantId
}

func (s *ListVccGrantRulesResponseBodyContentData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVccGrantRulesResponseBodyContentData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListVccGrantRulesResponseBodyContentData) GetProduct() *string {
	return s.Product
}

func (s *ListVccGrantRulesResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVccGrantRulesResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVccGrantRulesResponseBodyContentData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVccGrantRulesResponseBodyContentData) GetUsed() *bool {
	return s.Used
}

func (s *ListVccGrantRulesResponseBodyContentData) SetCreateTime(v string) *ListVccGrantRulesResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetErId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.ErId = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetGrantRuleId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.GrantRuleId = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetGrantTenantId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.GrantTenantId = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetInstanceId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.InstanceId = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetInstanceName(v string) *ListVccGrantRulesResponseBodyContentData {
	s.InstanceName = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetProduct(v string) *ListVccGrantRulesResponseBodyContentData {
	s.Product = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetRegionId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetResourceGroupId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetTenantId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetUsed(v bool) *ListVccGrantRulesResponseBodyContentData {
	s.Used = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) Validate() error {
	return dara.Validate(s)
}
