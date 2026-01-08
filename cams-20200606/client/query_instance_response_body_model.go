// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryInstanceResponseBody
	GetCode() *string
	SetData(v *QueryInstanceResponseBodyData) *QueryInstanceResponseBody
	GetData() *QueryInstanceResponseBodyData
	SetMessage(v string) *QueryInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryInstanceResponseBody
	GetSuccess() *bool
}

type QueryInstanceResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// NULL
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 29kskkd******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryInstanceResponseBody) GetData() *QueryInstanceResponseBodyData {
	return s.Data
}

func (s *QueryInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryInstanceResponseBody) SetAccessDeniedDetail(v string) *QueryInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryInstanceResponseBody) SetCode(v string) *QueryInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInstanceResponseBody) SetData(v *QueryInstanceResponseBodyData) *QueryInstanceResponseBody {
	s.Data = v
	return s
}

func (s *QueryInstanceResponseBody) SetMessage(v string) *QueryInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstanceResponseBody) SetRequestId(v string) *QueryInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstanceResponseBody) SetSuccess(v bool) *QueryInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryInstanceResponseBodyData struct {
	// example:
	//
	// VIBER
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// 1@alibaba.com
	ContactMail *string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty"`
	// example:
	//
	// 1
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// example:
	//
	// 01
	CustType *string `json:"CustType,omitempty" xml:"CustType,omitempty"`
	// FacebookBmId
	//
	// example:
	//
	// 399298882
	FacebookBmId *string `json:"FacebookBmId,omitempty" xml:"FacebookBmId,omitempty"`
	// example:
	//
	// ins
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// example:
	//
	// 293939*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// viber_ins
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// https://alibaba.com/1.pdf
	IsvTerms *string `json:"IsvTerms,omitempty" xml:"IsvTerms,omitempty"`
	// example:
	//
	// 长沙麓谷
	OfficeAddress *string `json:"OfficeAddress,omitempty" xml:"OfficeAddress,omitempty"`
	// example:
	//
	// 111
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// 140092992
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s QueryInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInstanceResponseBodyData) GetChannelType() *string {
	return s.ChannelType
}

func (s *QueryInstanceResponseBodyData) GetContactMail() *string {
	return s.ContactMail
}

func (s *QueryInstanceResponseBodyData) GetCountryId() *string {
	return s.CountryId
}

func (s *QueryInstanceResponseBodyData) GetCustType() *string {
	return s.CustType
}

func (s *QueryInstanceResponseBodyData) GetFacebookBmId() *string {
	return s.FacebookBmId
}

func (s *QueryInstanceResponseBodyData) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *QueryInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryInstanceResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *QueryInstanceResponseBodyData) GetIsvTerms() *string {
	return s.IsvTerms
}

func (s *QueryInstanceResponseBodyData) GetOfficeAddress() *string {
	return s.OfficeAddress
}

func (s *QueryInstanceResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *QueryInstanceResponseBodyData) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *QueryInstanceResponseBodyData) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *QueryInstanceResponseBodyData) SetChannelType(v string) *QueryInstanceResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *QueryInstanceResponseBodyData) SetContactMail(v string) *QueryInstanceResponseBodyData {
	s.ContactMail = &v
	return s
}

func (s *QueryInstanceResponseBodyData) SetCountryId(v string) *QueryInstanceResponseBodyData {
	s.CountryId = &v
	return s
}

func (s *QueryInstanceResponseBodyData) SetCustType(v string) *QueryInstanceResponseBodyData {
	s.CustType = &v
	return s
}

func (s *QueryInstanceResponseBodyData) SetFacebookBmId(v string) *QueryInstanceResponseBodyData {
	s.FacebookBmId = &v
	return s
}

func (s *QueryInstanceResponseBodyData) SetInstanceDescription(v string) *QueryInstanceResponseBodyData {
	s.InstanceDescription = &v
	return s
}

func (s *QueryInstanceResponseBodyData) SetInstanceId(v string) *QueryInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *QueryInstanceResponseBodyData) SetInstanceName(v string) *QueryInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *QueryInstanceResponseBodyData) SetIsvTerms(v string) *QueryInstanceResponseBodyData {
	s.IsvTerms = &v
	return s
}

func (s *QueryInstanceResponseBodyData) SetOfficeAddress(v string) *QueryInstanceResponseBodyData {
	s.OfficeAddress = &v
	return s
}

func (s *QueryInstanceResponseBodyData) SetResourceGroupId(v string) *QueryInstanceResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *QueryInstanceResponseBodyData) SetResourceRegionId(v string) *QueryInstanceResponseBodyData {
	s.ResourceRegionId = &v
	return s
}

func (s *QueryInstanceResponseBodyData) SetSubmitTime(v string) *QueryInstanceResponseBodyData {
	s.SubmitTime = &v
	return s
}

func (s *QueryInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
