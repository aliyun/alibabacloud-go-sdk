// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWhatsappCallShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *WhatsappCallShrinkRequest
	GetBusinessNumber() *string
	SetCallAction(v string) *WhatsappCallShrinkRequest
	GetCallAction() *string
	SetCallId(v string) *WhatsappCallShrinkRequest
	GetCallId() *string
	SetCustSpaceId(v string) *WhatsappCallShrinkRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *WhatsappCallShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *WhatsappCallShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *WhatsappCallShrinkRequest
	GetResourceOwnerId() *int64
	SetSessionShrink(v string) *WhatsappCallShrinkRequest
	GetSessionShrink() *string
	SetUserNumber(v string) *WhatsappCallShrinkRequest
	GetUserNumber() *string
}

type WhatsappCallShrinkRequest struct {
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
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SessionShrink        *string `json:"Session,omitempty" xml:"Session,omitempty"`
	// example:
	//
	// 示例值示例值
	UserNumber *string `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
}

func (s WhatsappCallShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s WhatsappCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *WhatsappCallShrinkRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *WhatsappCallShrinkRequest) GetCallAction() *string {
	return s.CallAction
}

func (s *WhatsappCallShrinkRequest) GetCallId() *string {
	return s.CallId
}

func (s *WhatsappCallShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *WhatsappCallShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *WhatsappCallShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *WhatsappCallShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *WhatsappCallShrinkRequest) GetSessionShrink() *string {
	return s.SessionShrink
}

func (s *WhatsappCallShrinkRequest) GetUserNumber() *string {
	return s.UserNumber
}

func (s *WhatsappCallShrinkRequest) SetBusinessNumber(v string) *WhatsappCallShrinkRequest {
	s.BusinessNumber = &v
	return s
}

func (s *WhatsappCallShrinkRequest) SetCallAction(v string) *WhatsappCallShrinkRequest {
	s.CallAction = &v
	return s
}

func (s *WhatsappCallShrinkRequest) SetCallId(v string) *WhatsappCallShrinkRequest {
	s.CallId = &v
	return s
}

func (s *WhatsappCallShrinkRequest) SetCustSpaceId(v string) *WhatsappCallShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *WhatsappCallShrinkRequest) SetOwnerId(v int64) *WhatsappCallShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *WhatsappCallShrinkRequest) SetResourceOwnerAccount(v string) *WhatsappCallShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *WhatsappCallShrinkRequest) SetResourceOwnerId(v int64) *WhatsappCallShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *WhatsappCallShrinkRequest) SetSessionShrink(v string) *WhatsappCallShrinkRequest {
	s.SessionShrink = &v
	return s
}

func (s *WhatsappCallShrinkRequest) SetUserNumber(v string) *WhatsappCallShrinkRequest {
	s.UserNumber = &v
	return s
}

func (s *WhatsappCallShrinkRequest) Validate() error {
	return dara.Validate(s)
}
