// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthAndRefreshLoginTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AuthAndRefreshLoginTicketResponseBody
	GetCode() *string
	SetData(v *AuthAndRefreshLoginTicketResponseBodyData) *AuthAndRefreshLoginTicketResponseBody
	GetData() *AuthAndRefreshLoginTicketResponseBodyData
	SetMsg(v string) *AuthAndRefreshLoginTicketResponseBody
	GetMsg() *string
	SetRequestId(v string) *AuthAndRefreshLoginTicketResponseBody
	GetRequestId() *string
}

type AuthAndRefreshLoginTicketResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AuthAndRefreshLoginTicketResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                                    `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthAndRefreshLoginTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthAndRefreshLoginTicketResponseBody) GoString() string {
	return s.String()
}

func (s *AuthAndRefreshLoginTicketResponseBody) GetCode() *string {
	return s.Code
}

func (s *AuthAndRefreshLoginTicketResponseBody) GetData() *AuthAndRefreshLoginTicketResponseBodyData {
	return s.Data
}

func (s *AuthAndRefreshLoginTicketResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *AuthAndRefreshLoginTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthAndRefreshLoginTicketResponseBody) SetCode(v string) *AuthAndRefreshLoginTicketResponseBody {
	s.Code = &v
	return s
}

func (s *AuthAndRefreshLoginTicketResponseBody) SetData(v *AuthAndRefreshLoginTicketResponseBodyData) *AuthAndRefreshLoginTicketResponseBody {
	s.Data = v
	return s
}

func (s *AuthAndRefreshLoginTicketResponseBody) SetMsg(v string) *AuthAndRefreshLoginTicketResponseBody {
	s.Msg = &v
	return s
}

func (s *AuthAndRefreshLoginTicketResponseBody) SetRequestId(v string) *AuthAndRefreshLoginTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthAndRefreshLoginTicketResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AuthAndRefreshLoginTicketResponseBodyData struct {
	NewLoginTicket *string `json:"NewLoginTicket,omitempty" xml:"NewLoginTicket,omitempty"`
}

func (s AuthAndRefreshLoginTicketResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AuthAndRefreshLoginTicketResponseBodyData) GoString() string {
	return s.String()
}

func (s *AuthAndRefreshLoginTicketResponseBodyData) GetNewLoginTicket() *string {
	return s.NewLoginTicket
}

func (s *AuthAndRefreshLoginTicketResponseBodyData) SetNewLoginTicket(v string) *AuthAndRefreshLoginTicketResponseBodyData {
	s.NewLoginTicket = &v
	return s
}

func (s *AuthAndRefreshLoginTicketResponseBodyData) Validate() error {
	return dara.Validate(s)
}
