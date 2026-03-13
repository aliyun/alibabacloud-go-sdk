// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWebAuthnAuthenticatorRegistrationUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *GenerateWebAuthnAuthenticatorRegistrationUrlRequest
	GetDomainId() *string
	SetInstanceId(v string) *GenerateWebAuthnAuthenticatorRegistrationUrlRequest
	GetInstanceId() *string
	SetUserId(v string) *GenerateWebAuthnAuthenticatorRegistrationUrlRequest
	GetUserId() *string
}

type GenerateWebAuthnAuthenticatorRegistrationUrlRequest struct {
	// example:
	//
	// dm_nfplcagj5cguo2267bkjxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// UserID
	//
	// This parameter is required.
	//
	// example:
	//
	// user_e2jvisn35ea5lmthk267xx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GenerateWebAuthnAuthenticatorRegistrationUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateWebAuthnAuthenticatorRegistrationUrlRequest) GoString() string {
	return s.String()
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlRequest) GetUserId() *string {
	return s.UserId
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlRequest) SetDomainId(v string) *GenerateWebAuthnAuthenticatorRegistrationUrlRequest {
	s.DomainId = &v
	return s
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlRequest) SetInstanceId(v string) *GenerateWebAuthnAuthenticatorRegistrationUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlRequest) SetUserId(v string) *GenerateWebAuthnAuthenticatorRegistrationUrlRequest {
	s.UserId = &v
	return s
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlRequest) Validate() error {
	return dara.Validate(s)
}
