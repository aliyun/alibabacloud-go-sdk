// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWhatsappCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *WhatsappCallRequest
	GetBusinessNumber() *string
	SetCallAction(v string) *WhatsappCallRequest
	GetCallAction() *string
	SetCallId(v string) *WhatsappCallRequest
	GetCallId() *string
	SetCustSpaceId(v string) *WhatsappCallRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *WhatsappCallRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *WhatsappCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *WhatsappCallRequest
	GetResourceOwnerId() *int64
	SetSession(v *WhatsappCallRequestSession) *WhatsappCallRequest
	GetSession() *WhatsappCallRequestSession
	SetUserNumber(v string) *WhatsappCallRequest
	GetUserNumber() *string
}

type WhatsappCallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	BusinessNumber *string `json:"BusinessNumber,omitempty" xml:"BusinessNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	CallAction *string `json:"CallAction,omitempty" xml:"CallAction,omitempty"`
	// example:
	//
	// 示例值示例值
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	CustSpaceId          *string                     `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64                      `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                     `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                      `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Session              *WhatsappCallRequestSession `json:"Session,omitempty" xml:"Session,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	UserNumber *string `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
}

func (s WhatsappCallRequest) String() string {
	return dara.Prettify(s)
}

func (s WhatsappCallRequest) GoString() string {
	return s.String()
}

func (s *WhatsappCallRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *WhatsappCallRequest) GetCallAction() *string {
	return s.CallAction
}

func (s *WhatsappCallRequest) GetCallId() *string {
	return s.CallId
}

func (s *WhatsappCallRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *WhatsappCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *WhatsappCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *WhatsappCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *WhatsappCallRequest) GetSession() *WhatsappCallRequestSession {
	return s.Session
}

func (s *WhatsappCallRequest) GetUserNumber() *string {
	return s.UserNumber
}

func (s *WhatsappCallRequest) SetBusinessNumber(v string) *WhatsappCallRequest {
	s.BusinessNumber = &v
	return s
}

func (s *WhatsappCallRequest) SetCallAction(v string) *WhatsappCallRequest {
	s.CallAction = &v
	return s
}

func (s *WhatsappCallRequest) SetCallId(v string) *WhatsappCallRequest {
	s.CallId = &v
	return s
}

func (s *WhatsappCallRequest) SetCustSpaceId(v string) *WhatsappCallRequest {
	s.CustSpaceId = &v
	return s
}

func (s *WhatsappCallRequest) SetOwnerId(v int64) *WhatsappCallRequest {
	s.OwnerId = &v
	return s
}

func (s *WhatsappCallRequest) SetResourceOwnerAccount(v string) *WhatsappCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *WhatsappCallRequest) SetResourceOwnerId(v int64) *WhatsappCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *WhatsappCallRequest) SetSession(v *WhatsappCallRequestSession) *WhatsappCallRequest {
	s.Session = v
	return s
}

func (s *WhatsappCallRequest) SetUserNumber(v string) *WhatsappCallRequest {
	s.UserNumber = &v
	return s
}

func (s *WhatsappCallRequest) Validate() error {
	if s.Session != nil {
		if err := s.Session.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WhatsappCallRequestSession struct {
	// example:
	//
	// 示例值
	Sdp *string `json:"Sdp,omitempty" xml:"Sdp,omitempty"`
	// example:
	//
	// 示例值
	SdpType *string `json:"SdpType,omitempty" xml:"SdpType,omitempty"`
}

func (s WhatsappCallRequestSession) String() string {
	return dara.Prettify(s)
}

func (s WhatsappCallRequestSession) GoString() string {
	return s.String()
}

func (s *WhatsappCallRequestSession) GetSdp() *string {
	return s.Sdp
}

func (s *WhatsappCallRequestSession) GetSdpType() *string {
	return s.SdpType
}

func (s *WhatsappCallRequestSession) SetSdp(v string) *WhatsappCallRequestSession {
	s.Sdp = &v
	return s
}

func (s *WhatsappCallRequestSession) SetSdpType(v string) *WhatsappCallRequestSession {
	s.SdpType = &v
	return s
}

func (s *WhatsappCallRequestSession) Validate() error {
	return dara.Validate(s)
}
