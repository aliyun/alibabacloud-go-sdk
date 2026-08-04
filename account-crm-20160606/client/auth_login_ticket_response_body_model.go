// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AuthLoginTicketResponseBody
	GetCode() *string
	SetLoginTicketDto(v *AuthLoginTicketResponseBodyLoginTicketDto) *AuthLoginTicketResponseBody
	GetLoginTicketDto() *AuthLoginTicketResponseBodyLoginTicketDto
	SetMessage(v string) *AuthLoginTicketResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthLoginTicketResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AuthLoginTicketResponseBody
	GetSuccess() *bool
}

type AuthLoginTicketResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	LoginTicketDto *AuthLoginTicketResponseBodyLoginTicketDto `json:"LoginTicketDto,omitempty" xml:"LoginTicketDto,omitempty" type:"Struct"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthLoginTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginTicketResponseBody) GoString() string {
	return s.String()
}

func (s *AuthLoginTicketResponseBody) GetCode() *string {
	return s.Code
}

func (s *AuthLoginTicketResponseBody) GetLoginTicketDto() *AuthLoginTicketResponseBodyLoginTicketDto {
	return s.LoginTicketDto
}

func (s *AuthLoginTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthLoginTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthLoginTicketResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthLoginTicketResponseBody) SetCode(v string) *AuthLoginTicketResponseBody {
	s.Code = &v
	return s
}

func (s *AuthLoginTicketResponseBody) SetLoginTicketDto(v *AuthLoginTicketResponseBodyLoginTicketDto) *AuthLoginTicketResponseBody {
	s.LoginTicketDto = v
	return s
}

func (s *AuthLoginTicketResponseBody) SetMessage(v string) *AuthLoginTicketResponseBody {
	s.Message = &v
	return s
}

func (s *AuthLoginTicketResponseBody) SetRequestId(v string) *AuthLoginTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthLoginTicketResponseBody) SetSuccess(v bool) *AuthLoginTicketResponseBody {
	s.Success = &v
	return s
}

func (s *AuthLoginTicketResponseBody) Validate() error {
	if s.LoginTicketDto != nil {
		if err := s.LoginTicketDto.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AuthLoginTicketResponseBodyLoginTicketDto struct {
	LoginTicket *string `json:"LoginTicket,omitempty" xml:"LoginTicket,omitempty"`
}

func (s AuthLoginTicketResponseBodyLoginTicketDto) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginTicketResponseBodyLoginTicketDto) GoString() string {
	return s.String()
}

func (s *AuthLoginTicketResponseBodyLoginTicketDto) GetLoginTicket() *string {
	return s.LoginTicket
}

func (s *AuthLoginTicketResponseBodyLoginTicketDto) SetLoginTicket(v string) *AuthLoginTicketResponseBodyLoginTicketDto {
	s.LoginTicket = &v
	return s
}

func (s *AuthLoginTicketResponseBodyLoginTicketDto) Validate() error {
	return dara.Validate(s)
}
