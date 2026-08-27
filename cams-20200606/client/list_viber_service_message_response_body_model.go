// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListViberServiceMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListViberServiceMessageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListViberServiceMessageResponseBody
	GetCode() *string
	SetData(v []*ListViberServiceMessageResponseBodyData) *ListViberServiceMessageResponseBody
	GetData() []*ListViberServiceMessageResponseBodyData
	SetMessage(v string) *ListViberServiceMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListViberServiceMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListViberServiceMessageResponseBody
	GetSuccess() *bool
}

type ListViberServiceMessageResponseBody struct {
	// The access denial details.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data object.
	Data []*ListViberServiceMessageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. This is used for troubleshooting when an error occurs.
	//
	// example:
	//
	// ddhjdn-dnjdnkdjknd**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListViberServiceMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListViberServiceMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ListViberServiceMessageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListViberServiceMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListViberServiceMessageResponseBody) GetData() []*ListViberServiceMessageResponseBodyData {
	return s.Data
}

func (s *ListViberServiceMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListViberServiceMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListViberServiceMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListViberServiceMessageResponseBody) SetAccessDeniedDetail(v string) *ListViberServiceMessageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListViberServiceMessageResponseBody) SetCode(v string) *ListViberServiceMessageResponseBody {
	s.Code = &v
	return s
}

func (s *ListViberServiceMessageResponseBody) SetData(v []*ListViberServiceMessageResponseBodyData) *ListViberServiceMessageResponseBody {
	s.Data = v
	return s
}

func (s *ListViberServiceMessageResponseBody) SetMessage(v string) *ListViberServiceMessageResponseBody {
	s.Message = &v
	return s
}

func (s *ListViberServiceMessageResponseBody) SetRequestId(v string) *ListViberServiceMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListViberServiceMessageResponseBody) SetSuccess(v bool) *ListViberServiceMessageResponseBody {
	s.Success = &v
	return s
}

func (s *ListViberServiceMessageResponseBody) Validate() error {
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

type ListViberServiceMessageResponseBodyData struct {
	// The business account name.
	//
	// example:
	//
	// Test.
	BusinessAccountName *string `json:"BusinessAccountName,omitempty" xml:"BusinessAccountName,omitempty"`
	// The destination country or region ID.
	DestinationCountryId []*string `json:"DestinationCountryId,omitempty" xml:"DestinationCountryId,omitempty" type:"Repeated"`
	// The destination country or region ID.
	DestinationInternationalCountryId []*string `json:"DestinationInternationalCountryId,omitempty" xml:"DestinationInternationalCountryId,omitempty" type:"Repeated"`
	// The industry involved.
	//
	// example:
	//
	// Healthcare.
	IndustryInvolved *string `json:"IndustryInvolved,omitempty" xml:"IndustryInvolved,omitempty"`
	// The collection of destination countries or regions.
	MessageDestinationCountry []*string `json:"MessageDestinationCountry,omitempty" xml:"MessageDestinationCountry,omitempty" type:"Repeated"`
	// The list of international destination countries or regions.
	MessageDestinationInternationalCountry []*string `json:"MessageDestinationInternationalCountry,omitempty" xml:"MessageDestinationInternationalCountry,omitempty" type:"Repeated"`
	// The service ID.
	//
	// example:
	//
	// 25644
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The resource status.
	//
	// example:
	//
	// stop
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListViberServiceMessageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListViberServiceMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListViberServiceMessageResponseBodyData) GetBusinessAccountName() *string {
	return s.BusinessAccountName
}

func (s *ListViberServiceMessageResponseBodyData) GetDestinationCountryId() []*string {
	return s.DestinationCountryId
}

func (s *ListViberServiceMessageResponseBodyData) GetDestinationInternationalCountryId() []*string {
	return s.DestinationInternationalCountryId
}

func (s *ListViberServiceMessageResponseBodyData) GetIndustryInvolved() *string {
	return s.IndustryInvolved
}

func (s *ListViberServiceMessageResponseBodyData) GetMessageDestinationCountry() []*string {
	return s.MessageDestinationCountry
}

func (s *ListViberServiceMessageResponseBodyData) GetMessageDestinationInternationalCountry() []*string {
	return s.MessageDestinationInternationalCountry
}

func (s *ListViberServiceMessageResponseBodyData) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListViberServiceMessageResponseBodyData) GetState() *string {
	return s.State
}

func (s *ListViberServiceMessageResponseBodyData) SetBusinessAccountName(v string) *ListViberServiceMessageResponseBodyData {
	s.BusinessAccountName = &v
	return s
}

func (s *ListViberServiceMessageResponseBodyData) SetDestinationCountryId(v []*string) *ListViberServiceMessageResponseBodyData {
	s.DestinationCountryId = v
	return s
}

func (s *ListViberServiceMessageResponseBodyData) SetDestinationInternationalCountryId(v []*string) *ListViberServiceMessageResponseBodyData {
	s.DestinationInternationalCountryId = v
	return s
}

func (s *ListViberServiceMessageResponseBodyData) SetIndustryInvolved(v string) *ListViberServiceMessageResponseBodyData {
	s.IndustryInvolved = &v
	return s
}

func (s *ListViberServiceMessageResponseBodyData) SetMessageDestinationCountry(v []*string) *ListViberServiceMessageResponseBodyData {
	s.MessageDestinationCountry = v
	return s
}

func (s *ListViberServiceMessageResponseBodyData) SetMessageDestinationInternationalCountry(v []*string) *ListViberServiceMessageResponseBodyData {
	s.MessageDestinationInternationalCountry = v
	return s
}

func (s *ListViberServiceMessageResponseBodyData) SetServiceId(v string) *ListViberServiceMessageResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *ListViberServiceMessageResponseBodyData) SetState(v string) *ListViberServiceMessageResponseBodyData {
	s.State = &v
	return s
}

func (s *ListViberServiceMessageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
