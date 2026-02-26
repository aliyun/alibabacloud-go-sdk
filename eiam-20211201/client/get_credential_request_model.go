// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialId(v string) *GetCredentialRequest
	GetCredentialId() *string
	SetInstanceId(v string) *GetCredentialRequest
	GetInstanceId() *string
}

type GetCredentialRequest struct {
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

func (s GetCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetCredentialRequest) GetCredentialId() *string {
	return s.CredentialId
}

func (s *GetCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCredentialRequest) SetCredentialId(v string) *GetCredentialRequest {
	s.CredentialId = &v
	return s
}

func (s *GetCredentialRequest) SetInstanceId(v string) *GetCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCredentialRequest) Validate() error {
	return dara.Validate(s)
}
