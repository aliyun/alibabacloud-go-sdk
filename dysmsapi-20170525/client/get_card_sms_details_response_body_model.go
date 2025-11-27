// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCardSmsDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetCardSmsDetailsResponseBody
	GetAccessDeniedDetail() *string
	SetCardSendDetailDTO(v *GetCardSmsDetailsResponseBodyCardSendDetailDTO) *GetCardSmsDetailsResponseBody
	GetCardSendDetailDTO() *GetCardSmsDetailsResponseBodyCardSendDetailDTO
	SetCode(v string) *GetCardSmsDetailsResponseBody
	GetCode() *string
	SetMessage(v string) *GetCardSmsDetailsResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetCardSmsDetailsResponseBody
	GetSuccess() *bool
}

type GetCardSmsDetailsResponseBody struct {
	// Access denied detail; this field is returned only if the RAM check fails.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Card SMS sending result
	CardSendDetailDTO *GetCardSmsDetailsResponseBodyCardSendDetailDTO `json:"CardSendDetailDTO,omitempty" xml:"CardSendDetailDTO,omitempty" type:"Struct"`
	// Request status code.
	//
	// 	- OK indicates a successful request.
	//
	// 	- For other error codes, see [API Error Codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the API call was successful. Values:
	//
	// - **true*	- - **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCardSmsDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCardSmsDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardSmsDetailsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetCardSmsDetailsResponseBody) GetCardSendDetailDTO() *GetCardSmsDetailsResponseBodyCardSendDetailDTO {
	return s.CardSendDetailDTO
}

func (s *GetCardSmsDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCardSmsDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCardSmsDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCardSmsDetailsResponseBody) SetAccessDeniedDetail(v string) *GetCardSmsDetailsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetCardSmsDetailsResponseBody) SetCardSendDetailDTO(v *GetCardSmsDetailsResponseBodyCardSendDetailDTO) *GetCardSmsDetailsResponseBody {
	s.CardSendDetailDTO = v
	return s
}

func (s *GetCardSmsDetailsResponseBody) SetCode(v string) *GetCardSmsDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *GetCardSmsDetailsResponseBody) SetMessage(v string) *GetCardSmsDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *GetCardSmsDetailsResponseBody) SetSuccess(v bool) *GetCardSmsDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *GetCardSmsDetailsResponseBody) Validate() error {
	if s.CardSendDetailDTO != nil {
		if err := s.CardSendDetailDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCardSmsDetailsResponseBodyCardSendDetailDTO struct {
	// Current page number
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of card SMS sending records
	Records []*GetCardSmsDetailsResponseBodyCardSendDetailDTORecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// Total count
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCardSmsDetailsResponseBodyCardSendDetailDTO) String() string {
	return dara.Prettify(s)
}

func (s GetCardSmsDetailsResponseBodyCardSendDetailDTO) GoString() string {
	return s.String()
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTO) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTO) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTO) GetRecords() []*GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	return s.Records
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTO) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTO) SetCurrentPage(v int64) *GetCardSmsDetailsResponseBodyCardSendDetailDTO {
	s.CurrentPage = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTO) SetPageSize(v int64) *GetCardSmsDetailsResponseBodyCardSendDetailDTO {
	s.PageSize = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTO) SetRecords(v []*GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) *GetCardSmsDetailsResponseBodyCardSendDetailDTO {
	s.Records = v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTO) SetTotalCount(v int64) *GetCardSmsDetailsResponseBodyCardSendDetailDTO {
	s.TotalCount = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTO) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCardSmsDetailsResponseBodyCardSendDetailDTORecords struct {
	// Error code for sending
	//
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// Customer-transmitted outId
	//
	// example:
	//
	// 12345678
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// Phone number that received the SMS
	//
	// example:
	//
	// 156****9080
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Receive date
	//
	// example:
	//
	// 2024-09-27 11:26:35
	ReceiveDate *string `json:"ReceiveDate,omitempty" xml:"ReceiveDate,omitempty"`
	// Receive SMS type
	//
	// example:
	//
	// CARD_SMS
	ReceiveType *string `json:"ReceiveType,omitempty" xml:"ReceiveType,omitempty"`
	// Render date
	//
	// example:
	//
	// 2024-09-27 12:13:39
	RenderDate *string `json:"RenderDate,omitempty" xml:"RenderDate,omitempty"`
	// Render status. 0: Not rendered; 1: Rendered successfully; 3: Not rendered
	//
	// example:
	//
	// 1
	RenderStatus *int64 `json:"RenderStatus,omitempty" xml:"RenderStatus,omitempty"`
	// Time when the SMS was sent
	//
	// example:
	//
	// 2024-09-27 11:26:32
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	// Sending status. 1: Sending; 2: Send failed; 3: Sent successfully; 4: Addressing failed
	//
	// example:
	//
	// 3
	SendStatus *int64 `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	// SMS content. Only applicable for text messages.
	//
	// example:
	//
	// 您收到一条短信消息
	SmsContent *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	// Template code
	//
	// example:
	//
	// CARD_SMS_6***
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) String() string {
	return dara.Prettify(s)
}

func (s GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) GoString() string {
	return s.String()
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) GetOutId() *string {
	return s.OutId
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) GetReceiveDate() *string {
	return s.ReceiveDate
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) GetReceiveType() *string {
	return s.ReceiveType
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) GetRenderDate() *string {
	return s.RenderDate
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) GetRenderStatus() *int64 {
	return s.RenderStatus
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) GetSendDate() *string {
	return s.SendDate
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) GetSendStatus() *int64 {
	return s.SendStatus
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) GetSmsContent() *string {
	return s.SmsContent
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetErrCode(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.ErrCode = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetOutId(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.OutId = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetPhoneNumber(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.PhoneNumber = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetReceiveDate(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.ReceiveDate = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetReceiveType(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.ReceiveType = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetRenderDate(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.RenderDate = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetRenderStatus(v int64) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.RenderStatus = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetSendDate(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.SendDate = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetSendStatus(v int64) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.SendStatus = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetSmsContent(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.SmsContent = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetTemplateCode(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.TemplateCode = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) Validate() error {
	return dara.Validate(s)
}
