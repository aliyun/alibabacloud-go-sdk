// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupplierRegistrationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSupplierRegistrationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSupplierRegistrationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSupplierRegistrationsResponseBody
	GetRequestId() *string
	SetSupplierRegistrations(v []*ListSupplierRegistrationsResponseBodySupplierRegistrations) *ListSupplierRegistrationsResponseBody
	GetSupplierRegistrations() []*ListSupplierRegistrationsResponseBodySupplierRegistrations
	SetTotalCount(v int32) *ListSupplierRegistrationsResponseBody
	GetTotalCount() *int32
}

type ListSupplierRegistrationsResponseBody struct {
	// The number of entries returned per page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// AAAAAdx9kBO7qKpr9My/+XQo0oY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C6CC568D-xxxx-xxxx-xxxx-08EB8E9F9F20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of service provider registration requests.
	SupplierRegistrations []*ListSupplierRegistrationsResponseBodySupplierRegistrations `json:"SupplierRegistrations,omitempty" xml:"SupplierRegistrations,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSupplierRegistrationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSupplierRegistrationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupplierRegistrationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSupplierRegistrationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupplierRegistrationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSupplierRegistrationsResponseBody) GetSupplierRegistrations() []*ListSupplierRegistrationsResponseBodySupplierRegistrations {
	return s.SupplierRegistrations
}

func (s *ListSupplierRegistrationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSupplierRegistrationsResponseBody) SetMaxResults(v int32) *ListSupplierRegistrationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBody) SetNextToken(v string) *ListSupplierRegistrationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBody) SetRequestId(v string) *ListSupplierRegistrationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBody) SetSupplierRegistrations(v []*ListSupplierRegistrationsResponseBodySupplierRegistrations) *ListSupplierRegistrationsResponseBody {
	s.SupplierRegistrations = v
	return s
}

func (s *ListSupplierRegistrationsResponseBody) SetTotalCount(v int32) *ListSupplierRegistrationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBody) Validate() error {
	if s.SupplierRegistrations != nil {
		for _, item := range s.SupplierRegistrations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSupplierRegistrationsResponseBodySupplierRegistrations struct {
	// The review comments.
	//
	// example:
	//
	// 无
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The email address of the contact person.
	//
	// example:
	//
	// test@163.com
	ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	// The phone number of the contact person.
	//
	// example:
	//
	// 135xxxxxxxx
	ContactNumber *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	// The contact person.
	//
	// example:
	//
	// John Doe
	ContactPerson *string `json:"ContactPerson,omitempty" xml:"ContactPerson,omitempty"`
	// The title of the contact person.
	//
	// example:
	//
	// CTO
	ContactPersonTitle *string `json:"ContactPersonTitle,omitempty" xml:"ContactPersonTitle,omitempty"`
	// Indicates whether the reseller mode is supported.
	//
	// example:
	//
	// false
	EnableResellerMode *bool `json:"EnableResellerMode,omitempty" xml:"EnableResellerMode,omitempty"`
	// The annual revenue of the service provider\\"s product.
	//
	// example:
	//
	// 10000000
	ProductAnnualRevenue *string `json:"ProductAnnualRevenue,omitempty" xml:"ProductAnnualRevenue,omitempty"`
	// The industry of the service provider\\"s product.
	//
	// example:
	//
	// AI
	ProductBusiness *string `json:"ProductBusiness,omitempty" xml:"ProductBusiness,omitempty"`
	// The delivery method of the service provider\\"s product. Valid values:
	//
	// - SaaS
	//
	// - License
	//
	// - API
	//
	// - Desktop software
	//
	// - Other
	//
	// example:
	//
	// SaaS
	ProductDeliveryTypes *string `json:"ProductDeliveryTypes,omitempty" xml:"ProductDeliveryTypes,omitempty"`
	// The product launch date.
	//
	// example:
	//
	// 2024.10.24
	ProductPublishTime *string `json:"ProductPublishTime,omitempty" xml:"ProductPublishTime,omitempty"`
	// The sales model of the service provider\\"s product. Valid values:
	//
	// - Direct
	//
	// - Channel
	//
	// example:
	//
	// 直销
	ProductSellTypes *string `json:"ProductSellTypes,omitempty" xml:"ProductSellTypes,omitempty"`
	// The ID of the review request.
	//
	// example:
	//
	// sr-xxx
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The description of the reseller business.
	//
	// example:
	//
	// 无
	ResellBusinessDesc *string `json:"ResellBusinessDesc,omitempty" xml:"ResellBusinessDesc,omitempty"`
	// The review status. Valid values:
	//
	// - Submitted: The request is submitted.
	//
	// - Approved: The request is approved.
	//
	// - Rejected: The request is rejected.
	//
	// example:
	//
	// Submitted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the request was submitted.
	//
	// example:
	//
	// 2025-01-22 09:47:58
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The description of the service provider.
	//
	// example:
	//
	// 服务商测试申请
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
	// The icon of the service provider.
	//
	// example:
	//
	// https://service-info-public.oss-cn-hangzhou.aliyuncs.com/xxx/service-xxx/xxx.png
	SupplierLogo *string `json:"SupplierLogo,omitempty" xml:"SupplierLogo,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// ComputeNest community service
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The English name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud ComputeNest
	SupplierNameEn *string `json:"SupplierNameEn,omitempty" xml:"SupplierNameEn,omitempty"`
	// The UID of the service provider.
	//
	// example:
	//
	// 1256xxx23434
	SupplierUid *string `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	// The URL of the service provider\\"s official website.
	//
	// example:
	//
	// https://www.guangbao-uni.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
}

func (s ListSupplierRegistrationsResponseBodySupplierRegistrations) String() string {
	return dara.Prettify(s)
}

func (s ListSupplierRegistrationsResponseBodySupplierRegistrations) GoString() string {
	return s.String()
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetComment() *string {
	return s.Comment
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetContactNumber() *string {
	return s.ContactNumber
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetContactPerson() *string {
	return s.ContactPerson
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetContactPersonTitle() *string {
	return s.ContactPersonTitle
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetEnableResellerMode() *bool {
	return s.EnableResellerMode
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetProductAnnualRevenue() *string {
	return s.ProductAnnualRevenue
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetProductBusiness() *string {
	return s.ProductBusiness
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetProductDeliveryTypes() *string {
	return s.ProductDeliveryTypes
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetProductPublishTime() *string {
	return s.ProductPublishTime
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetProductSellTypes() *string {
	return s.ProductSellTypes
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetRegistrationId() *string {
	return s.RegistrationId
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetResellBusinessDesc() *string {
	return s.ResellBusinessDesc
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetStatus() *string {
	return s.Status
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetSupplierDesc() *string {
	return s.SupplierDesc
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetSupplierLogo() *string {
	return s.SupplierLogo
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetSupplierName() *string {
	return s.SupplierName
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetSupplierNameEn() *string {
	return s.SupplierNameEn
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetSupplierUid() *string {
	return s.SupplierUid
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) GetSupplierUrl() *string {
	return s.SupplierUrl
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetComment(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.Comment = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetContactEmail(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ContactEmail = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetContactNumber(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ContactNumber = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetContactPerson(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ContactPerson = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetContactPersonTitle(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ContactPersonTitle = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetEnableResellerMode(v bool) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.EnableResellerMode = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetProductAnnualRevenue(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ProductAnnualRevenue = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetProductBusiness(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ProductBusiness = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetProductDeliveryTypes(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ProductDeliveryTypes = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetProductPublishTime(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ProductPublishTime = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetProductSellTypes(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ProductSellTypes = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetRegistrationId(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.RegistrationId = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetResellBusinessDesc(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ResellBusinessDesc = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetStatus(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.Status = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSubmitTime(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SubmitTime = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSupplierDesc(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SupplierDesc = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSupplierLogo(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SupplierLogo = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSupplierName(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SupplierName = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSupplierNameEn(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SupplierNameEn = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSupplierUid(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SupplierUid = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSupplierUrl(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SupplierUrl = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) Validate() error {
	return dara.Validate(s)
}
