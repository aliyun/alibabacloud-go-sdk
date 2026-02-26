// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialId(v string) *DeleteCredentialRequest
	GetCredentialId() *string
	SetInstanceId(v string) *DeleteCredentialRequest
	GetInstanceId() *string
}

type DeleteCredentialRequest struct {
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

func (s DeleteCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialRequest) GoString() string {
	return s.String()
}

func (s *DeleteCredentialRequest) GetCredentialId() *string {
	return s.CredentialId
}

func (s *DeleteCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCredentialRequest) SetCredentialId(v string) *DeleteCredentialRequest {
	s.CredentialId = &v
	return s
}

func (s *DeleteCredentialRequest) SetInstanceId(v string) *DeleteCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCredentialRequest) Validate() error {
	return dara.Validate(s)
}
