// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableCredentialRequest
	GetClientToken() *string
	SetCredentialId(v string) *DisableCredentialRequest
	GetCredentialId() *string
	SetInstanceId(v string) *DisableCredentialRequest
	GetInstanceId() *string
}

type DisableCredentialRequest struct {
	// Ensure idempotence. Generate a unique value on your client for each request. ClientToken supports only ASCII characters and must be no longer than 64 characters. For more information, see [How to Ensure Idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the credential.
	//
	// This parameter is required.
	//
	// example:
	//
	// cred_mkv7rgt4d7i4u7zqtzev2mxxxx
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableCredentialRequest) GoString() string {
	return s.String()
}

func (s *DisableCredentialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableCredentialRequest) GetCredentialId() *string {
	return s.CredentialId
}

func (s *DisableCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableCredentialRequest) SetClientToken(v string) *DisableCredentialRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableCredentialRequest) SetCredentialId(v string) *DisableCredentialRequest {
	s.CredentialId = &v
	return s
}

func (s *DisableCredentialRequest) SetInstanceId(v string) *DisableCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableCredentialRequest) Validate() error {
	return dara.Validate(s)
}
