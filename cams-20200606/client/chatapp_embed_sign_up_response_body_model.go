// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappEmbedSignUpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ChatappEmbedSignUpResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ChatappEmbedSignUpResponseBody
	GetCode() *string
	SetMessage(v string) *ChatappEmbedSignUpResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChatappEmbedSignUpResponseBody
	GetRequestId() *string
	SetWabas(v []*ChatappEmbedSignUpResponseBodyWabas) *ChatappEmbedSignUpResponseBody
	GetWabas() []*ChatappEmbedSignUpResponseBodyWabas
}

type ChatappEmbedSignUpResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of the WhatsApp Business accounts.
	Wabas []*ChatappEmbedSignUpResponseBodyWabas `json:"Wabas,omitempty" xml:"Wabas,omitempty" type:"Repeated"`
}

func (s ChatappEmbedSignUpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatappEmbedSignUpResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappEmbedSignUpResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ChatappEmbedSignUpResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChatappEmbedSignUpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChatappEmbedSignUpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatappEmbedSignUpResponseBody) GetWabas() []*ChatappEmbedSignUpResponseBodyWabas {
	return s.Wabas
}

func (s *ChatappEmbedSignUpResponseBody) SetAccessDeniedDetail(v string) *ChatappEmbedSignUpResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBody) SetCode(v string) *ChatappEmbedSignUpResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBody) SetMessage(v string) *ChatappEmbedSignUpResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBody) SetRequestId(v string) *ChatappEmbedSignUpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBody) SetWabas(v []*ChatappEmbedSignUpResponseBodyWabas) *ChatappEmbedSignUpResponseBody {
	s.Wabas = v
	return s
}

func (s *ChatappEmbedSignUpResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChatappEmbedSignUpResponseBodyWabas struct {
	// The review state of the WABA.
	//
	// example:
	//
	// VERIFIED
	AccountReviewStatus *string `json:"AccountReviewStatus,omitempty" xml:"AccountReviewStatus,omitempty"`
	// The currency.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The ID of the WABA.
	//
	// example:
	//
	// 2939933992*****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The namespace of the message template.
	//
	// example:
	//
	// alals-lsslls-slslsos-slsl
	MessageTemplateNamespace *string `json:"MessageTemplateNamespace,omitempty" xml:"MessageTemplateNamespace,omitempty"`
	// The name of the WABA.
	//
	// example:
	//
	// Alibaba
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ChatappEmbedSignUpResponseBodyWabas) String() string {
	return dara.Prettify(s)
}

func (s ChatappEmbedSignUpResponseBodyWabas) GoString() string {
	return s.String()
}

func (s *ChatappEmbedSignUpResponseBodyWabas) GetAccountReviewStatus() *string {
	return s.AccountReviewStatus
}

func (s *ChatappEmbedSignUpResponseBodyWabas) GetCurrency() *string {
	return s.Currency
}

func (s *ChatappEmbedSignUpResponseBodyWabas) GetId() *string {
	return s.Id
}

func (s *ChatappEmbedSignUpResponseBodyWabas) GetMessageTemplateNamespace() *string {
	return s.MessageTemplateNamespace
}

func (s *ChatappEmbedSignUpResponseBodyWabas) GetName() *string {
	return s.Name
}

func (s *ChatappEmbedSignUpResponseBodyWabas) SetAccountReviewStatus(v string) *ChatappEmbedSignUpResponseBodyWabas {
	s.AccountReviewStatus = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBodyWabas) SetCurrency(v string) *ChatappEmbedSignUpResponseBodyWabas {
	s.Currency = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBodyWabas) SetId(v string) *ChatappEmbedSignUpResponseBodyWabas {
	s.Id = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBodyWabas) SetMessageTemplateNamespace(v string) *ChatappEmbedSignUpResponseBodyWabas {
	s.MessageTemplateNamespace = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBodyWabas) SetName(v string) *ChatappEmbedSignUpResponseBodyWabas {
	s.Name = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBodyWabas) Validate() error {
	return dara.Validate(s)
}
