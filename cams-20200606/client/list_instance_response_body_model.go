// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListInstanceResponseBody
	GetCode() *string
	SetData(v []*ListInstanceResponseBodyData) *ListInstanceResponseBody
	GetData() []*ListInstanceResponseBodyData
	SetMessage(v string) *ListInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstanceResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListInstanceResponseBody
	GetTotalCount() *int64
}

type ListInstanceResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - `OK` indicates that the request was successful.
	//
	// - For other error codes, see the [Error Code List](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// An array of objects, each representing an instance.
	Data []*ListInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// NULL
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2993*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call succeeded.
	//
	// - **true**: The call succeeded.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries that match the specified criteria.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceResponseBody) GetData() []*ListInstanceResponseBodyData {
	return s.Data
}

func (s *ListInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstanceResponseBody) SetAccessDeniedDetail(v string) *ListInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListInstanceResponseBody) SetCode(v string) *ListInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceResponseBody) SetData(v []*ListInstanceResponseBodyData) *ListInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceResponseBody) SetMessage(v string) *ListInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceResponseBody) SetRequestId(v string) *ListInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResponseBody) SetSuccess(v bool) *ListInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceResponseBody) SetTotalCount(v int64) *ListInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceResponseBody) Validate() error {
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

type ListInstanceResponseBodyData struct {
	// The channel type.
	//
	// example:
	//
	// VIBER
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The contact email address.
	//
	// example:
	//
	// ma**@gmail.com
	ContactMail *string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty"`
	// The country ID.
	//
	// example:
	//
	// 1
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// The customer space ID.
	//
	// example:
	//
	// dad-gf**
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The ID of the associated Facebook Business Manager account.
	//
	// example:
	//
	// 1
	FacebookBmId *string `json:"FacebookBmId,omitempty" xml:"FacebookBmId,omitempty"`
	// The instance description.
	//
	// example:
	//
	// ins
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 29339****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// viber_ins
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ISV terms.
	//
	// example:
	//
	// aa
	IsvTerms *string `json:"IsvTerms,omitempty" xml:"IsvTerms,omitempty"`
	// The office address.
	//
	// example:
	//
	// example
	OfficeAddress *string `json:"OfficeAddress,omitempty" xml:"OfficeAddress,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// 12
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the region where the resource is located.
	//
	// example:
	//
	// 11
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The state of the instance.
	//
	// example:
	//
	// published
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the instance was submitted.
	//
	// example:
	//
	// 2023-12-12 00:00:00
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s ListInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyData) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListInstanceResponseBodyData) GetContactMail() *string {
	return s.ContactMail
}

func (s *ListInstanceResponseBodyData) GetCountryId() *string {
	return s.CountryId
}

func (s *ListInstanceResponseBodyData) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListInstanceResponseBodyData) GetFacebookBmId() *string {
	return s.FacebookBmId
}

func (s *ListInstanceResponseBodyData) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *ListInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstanceResponseBodyData) GetIsvTerms() *string {
	return s.IsvTerms
}

func (s *ListInstanceResponseBodyData) GetOfficeAddress() *string {
	return s.OfficeAddress
}

func (s *ListInstanceResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstanceResponseBodyData) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ListInstanceResponseBodyData) GetState() *string {
	return s.State
}

func (s *ListInstanceResponseBodyData) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListInstanceResponseBodyData) SetChannelType(v string) *ListInstanceResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetContactMail(v string) *ListInstanceResponseBodyData {
	s.ContactMail = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetCountryId(v string) *ListInstanceResponseBodyData {
	s.CountryId = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetCustSpaceId(v string) *ListInstanceResponseBodyData {
	s.CustSpaceId = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetFacebookBmId(v string) *ListInstanceResponseBodyData {
	s.FacebookBmId = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetInstanceDescription(v string) *ListInstanceResponseBodyData {
	s.InstanceDescription = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetInstanceId(v string) *ListInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetInstanceName(v string) *ListInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetIsvTerms(v string) *ListInstanceResponseBodyData {
	s.IsvTerms = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetOfficeAddress(v string) *ListInstanceResponseBodyData {
	s.OfficeAddress = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetResourceGroupId(v string) *ListInstanceResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetResourceRegionId(v string) *ListInstanceResponseBodyData {
	s.ResourceRegionId = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetState(v string) *ListInstanceResponseBodyData {
	s.State = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetSubmitTime(v string) *ListInstanceResponseBodyData {
	s.SubmitTime = &v
	return s
}

func (s *ListInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
