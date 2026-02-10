// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySendDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySendDetailsResponseBody
	GetCode() *string
	SetMessage(v string) *QuerySendDetailsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySendDetailsResponseBody
	GetRequestId() *string
	SetSmsSendDetailDTOs(v *QuerySendDetailsResponseBodySmsSendDetailDTOs) *QuerySendDetailsResponseBody
	GetSmsSendDetailDTOs() *QuerySendDetailsResponseBodySmsSendDetailDTOs
	SetTotalCount(v string) *QuerySendDetailsResponseBody
	GetTotalCount() *string
}

type QuerySendDetailsResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E477085AAF
	RequestId         *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SmsSendDetailDTOs *QuerySendDetailsResponseBodySmsSendDetailDTOs `json:"SmsSendDetailDTOs,omitempty" xml:"SmsSendDetailDTOs,omitempty" type:"Struct"`
	// The number of sent messages.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySendDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySendDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySendDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySendDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySendDetailsResponseBody) GetSmsSendDetailDTOs() *QuerySendDetailsResponseBodySmsSendDetailDTOs {
	return s.SmsSendDetailDTOs
}

func (s *QuerySendDetailsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *QuerySendDetailsResponseBody) SetCode(v string) *QuerySendDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetMessage(v string) *QuerySendDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetRequestId(v string) *QuerySendDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetSmsSendDetailDTOs(v *QuerySendDetailsResponseBodySmsSendDetailDTOs) *QuerySendDetailsResponseBody {
	s.SmsSendDetailDTOs = v
	return s
}

func (s *QuerySendDetailsResponseBody) SetTotalCount(v string) *QuerySendDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QuerySendDetailsResponseBody) Validate() error {
	if s.SmsSendDetailDTOs != nil {
		if err := s.SmsSendDetailDTOs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySendDetailsResponseBodySmsSendDetailDTOs struct {
	SmsSendDetailDTO []*QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO `json:"SmsSendDetailDTO,omitempty" xml:"SmsSendDetailDTO,omitempty" type:"Repeated"`
}

func (s QuerySendDetailsResponseBodySmsSendDetailDTOs) String() string {
	return dara.Prettify(s)
}

func (s QuerySendDetailsResponseBodySmsSendDetailDTOs) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOs) GetSmsSendDetailDTO() []*QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	return s.SmsSendDetailDTO
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOs) SetSmsSendDetailDTO(v []*QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) *QuerySendDetailsResponseBodySmsSendDetailDTOs {
	s.SmsSendDetailDTO = v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOs) Validate() error {
	if s.SmsSendDetailDTO != nil {
		for _, item := range s.SmsSendDetailDTO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO struct {
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ErrCode      *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	OutId        *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	PhoneNum     *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	ReceiveDate  *string `json:"ReceiveDate,omitempty" xml:"ReceiveDate,omitempty"`
	SendDate     *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	SendStatus   *int64  `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) String() string {
	return dara.Prettify(s)
}

func (s QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) GetContent() *string {
	return s.Content
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) GetErrCode() *string {
	return s.ErrCode
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) GetOutId() *string {
	return s.OutId
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) GetPhoneNum() *string {
	return s.PhoneNum
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) GetReceiveDate() *string {
	return s.ReceiveDate
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) GetSendDate() *string {
	return s.SendDate
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) GetSendStatus() *int64 {
	return s.SendStatus
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetContent(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.Content = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetErrCode(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.ErrCode = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetOutId(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.OutId = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetPhoneNum(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.PhoneNum = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetReceiveDate(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.ReceiveDate = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetSendDate(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.SendDate = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetSendStatus(v int64) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.SendStatus = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetTemplateCode(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.TemplateCode = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) Validate() error {
	return dara.Validate(s)
}
