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
	// The action type. Valid values:
	//
	// 	- **DOMAINAUDIT**: review a domain name review.
	//
	// 	- **AUDITCONTACT**: review a contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// AUDITCONTACT
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// Thee instance ID of the domain name. You can call the [QueryDomainList](https://help.aliyun.com/document_detail/67712.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// S2019270W570xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the error message to return if the request fails. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the client. Set the value to **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.1
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
