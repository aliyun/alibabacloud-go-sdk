// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialId(v string) *ObtainCredentialRequest
	GetCredentialId() *string
	SetInstanceId(v string) *ObtainCredentialRequest
	GetInstanceId() *string
}

type ObtainCredentialRequest struct {
	// 凭据ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// cred_mkv7rgt4d7i4u7zqtzev2mxxxx
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ObtainCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s ObtainCredentialRequest) GoString() string {
	return s.String()
}

func (s *ObtainCredentialRequest) GetCredentialId() *string {
	return s.CredentialId
}

func (s *ObtainCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ObtainCredentialRequest) SetCredentialId(v string) *ObtainCredentialRequest {
	s.CredentialId = &v
	return s
}

func (s *ObtainCredentialRequest) SetInstanceId(v string) *ObtainCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *ObtainCredentialRequest) Validate() error {
	return dara.Validate(s)
}
