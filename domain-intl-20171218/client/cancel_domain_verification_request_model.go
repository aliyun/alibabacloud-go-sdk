// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDomainVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *CancelDomainVerificationRequest
	GetActionType() *string
	SetInstanceId(v string) *CancelDomainVerificationRequest
	GetInstanceId() *string
	SetLang(v string) *CancelDomainVerificationRequest
	GetLang() *string
	SetUserClientIp(v string) *CancelDomainVerificationRequest
	GetUserClientIp() *string
}

type CancelDomainVerificationRequest struct {
	// This parameter is required.
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CancelDomainVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelDomainVerificationRequest) GoString() string {
	return s.String()
}

func (s *CancelDomainVerificationRequest) GetActionType() *string {
	return s.ActionType
}

func (s *CancelDomainVerificationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelDomainVerificationRequest) GetLang() *string {
	return s.Lang
}

func (s *CancelDomainVerificationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *CancelDomainVerificationRequest) SetActionType(v string) *CancelDomainVerificationRequest {
	s.ActionType = &v
	return s
}

func (s *CancelDomainVerificationRequest) SetInstanceId(v string) *CancelDomainVerificationRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelDomainVerificationRequest) SetLang(v string) *CancelDomainVerificationRequest {
	s.Lang = &v
	return s
}

func (s *CancelDomainVerificationRequest) SetUserClientIp(v string) *CancelDomainVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *CancelDomainVerificationRequest) Validate() error {
	return dara.Validate(s)
}
