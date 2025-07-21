// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhatsappHealthStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetWhatsappHealthStatusResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetWhatsappHealthStatusResponseBody
	GetCode() *string
	SetData(v *GetWhatsappHealthStatusResponseBodyData) *GetWhatsappHealthStatusResponseBody
	GetData() *GetWhatsappHealthStatusResponseBodyData
	SetMessage(v string) *GetWhatsappHealthStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWhatsappHealthStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWhatsappHealthStatusResponseBody
	GetSuccess() *bool
}

type GetWhatsappHealthStatusResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetWhatsappHealthStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DAC72B08-3327-33EF-BEDC-8EC3E83A6575
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWhatsappHealthStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetWhatsappHealthStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetWhatsappHealthStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWhatsappHealthStatusResponseBody) GetData() *GetWhatsappHealthStatusResponseBodyData {
	return s.Data
}

func (s *GetWhatsappHealthStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWhatsappHealthStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWhatsappHealthStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWhatsappHealthStatusResponseBody) SetAccessDeniedDetail(v string) *GetWhatsappHealthStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBody) SetCode(v string) *GetWhatsappHealthStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBody) SetData(v *GetWhatsappHealthStatusResponseBodyData) *GetWhatsappHealthStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetWhatsappHealthStatusResponseBody) SetMessage(v string) *GetWhatsappHealthStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBody) SetRequestId(v string) *GetWhatsappHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBody) SetSuccess(v bool) *GetWhatsappHealthStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetWhatsappHealthStatusResponseBodyData struct {
	// Indicates whether the messages can be sent.
	//
	// example:
	//
	// AVAILABLE
	CanSendMessage *string `json:"CanSendMessage,omitempty" xml:"CanSendMessage,omitempty"`
	// The queried entities.
	Entities []*GetWhatsappHealthStatusResponseBodyDataEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
}

func (s GetWhatsappHealthStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappHealthStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWhatsappHealthStatusResponseBodyData) GetCanSendMessage() *string {
	return s.CanSendMessage
}

func (s *GetWhatsappHealthStatusResponseBodyData) GetEntities() []*GetWhatsappHealthStatusResponseBodyDataEntities {
	return s.Entities
}

func (s *GetWhatsappHealthStatusResponseBodyData) SetCanSendMessage(v string) *GetWhatsappHealthStatusResponseBodyData {
	s.CanSendMessage = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBodyData) SetEntities(v []*GetWhatsappHealthStatusResponseBodyDataEntities) *GetWhatsappHealthStatusResponseBodyData {
	s.Entities = v
	return s
}

func (s *GetWhatsappHealthStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetWhatsappHealthStatusResponseBodyDataEntities struct {
	// The Business Manager ID.
	//
	// example:
	//
	// 3992****
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// Indicates whether the messages can be sent.
	//
	// example:
	//
	// AVAILABLE
	CanSendMessage *string `json:"CanSendMessage,omitempty" xml:"CanSendMessage,omitempty"`
	// The entity type.
	//
	// example:
	//
	// PHONE_NUMBER
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The reasons why the messages failed to be sent.
	Errors []*GetWhatsappHealthStatusResponseBodyDataEntitiesErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// The template language.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The phone number to which the messages are sent.
	//
	// example:
	//
	// 86138****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The template code. This parameter is returned when the NodeType parameter is set to **template**.
	//
	// example:
	//
	// 939928****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The WABA ID. You can view the WABA ID in the Chat App Message Service console after you create the WABA.
	//
	// example:
	//
	// 39939***
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s GetWhatsappHealthStatusResponseBodyDataEntities) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappHealthStatusResponseBodyDataEntities) GoString() string {
	return s.String()
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) GetBusinessId() *string {
	return s.BusinessId
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) GetCanSendMessage() *string {
	return s.CanSendMessage
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) GetErrors() []*GetWhatsappHealthStatusResponseBodyDataEntitiesErrors {
	return s.Errors
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) GetLanguage() *string {
	return s.Language
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) GetWabaId() *string {
	return s.WabaId
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) SetBusinessId(v string) *GetWhatsappHealthStatusResponseBodyDataEntities {
	s.BusinessId = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) SetCanSendMessage(v string) *GetWhatsappHealthStatusResponseBodyDataEntities {
	s.CanSendMessage = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) SetEntityType(v string) *GetWhatsappHealthStatusResponseBodyDataEntities {
	s.EntityType = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) SetErrors(v []*GetWhatsappHealthStatusResponseBodyDataEntitiesErrors) *GetWhatsappHealthStatusResponseBodyDataEntities {
	s.Errors = v
	return s
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) SetLanguage(v string) *GetWhatsappHealthStatusResponseBodyDataEntities {
	s.Language = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) SetPhoneNumber(v string) *GetWhatsappHealthStatusResponseBodyDataEntities {
	s.PhoneNumber = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) SetTemplateCode(v string) *GetWhatsappHealthStatusResponseBodyDataEntities {
	s.TemplateCode = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) SetWabaId(v string) *GetWhatsappHealthStatusResponseBodyDataEntities {
	s.WabaId = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntities) Validate() error {
	return dara.Validate(s)
}

type GetWhatsappHealthStatusResponseBodyDataEntitiesErrors struct {
	// The error code.
	//
	// example:
	//
	// 141006
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The description of the error.
	//
	// example:
	//
	// There is an error with the payment method.
	ErrorDescription *string `json:"ErrorDescription,omitempty" xml:"ErrorDescription,omitempty"`
	// The possible solution to the error.
	//
	// example:
	//
	// There was an error with your payment method. Please add a new payment method to the account.
	PossibleSolution *string `json:"PossibleSolution,omitempty" xml:"PossibleSolution,omitempty"`
}

func (s GetWhatsappHealthStatusResponseBodyDataEntitiesErrors) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappHealthStatusResponseBodyDataEntitiesErrors) GoString() string {
	return s.String()
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntitiesErrors) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntitiesErrors) GetErrorDescription() *string {
	return s.ErrorDescription
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntitiesErrors) GetPossibleSolution() *string {
	return s.PossibleSolution
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntitiesErrors) SetErrorCode(v string) *GetWhatsappHealthStatusResponseBodyDataEntitiesErrors {
	s.ErrorCode = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntitiesErrors) SetErrorDescription(v string) *GetWhatsappHealthStatusResponseBodyDataEntitiesErrors {
	s.ErrorDescription = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntitiesErrors) SetPossibleSolution(v string) *GetWhatsappHealthStatusResponseBodyDataEntitiesErrors {
	s.PossibleSolution = &v
	return s
}

func (s *GetWhatsappHealthStatusResponseBodyDataEntitiesErrors) Validate() error {
	return dara.Validate(s)
}
