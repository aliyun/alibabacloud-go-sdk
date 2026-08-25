// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSCIMServerCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialId(v string) *DeleteSCIMServerCredentialRequest
	GetCredentialId() *string
	SetDirectoryId(v string) *DeleteSCIMServerCredentialRequest
	GetDirectoryId() *string
}

type DeleteSCIMServerCredentialRequest struct {
	// The ID of the SCIM credential.
	//
	// example:
	//
	// scimcred-004whl0kvfwcypbi****
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s DeleteSCIMServerCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSCIMServerCredentialRequest) GoString() string {
	return s.String()
}

func (s *DeleteSCIMServerCredentialRequest) GetCredentialId() *string {
	return s.CredentialId
}

func (s *DeleteSCIMServerCredentialRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeleteSCIMServerCredentialRequest) SetCredentialId(v string) *DeleteSCIMServerCredentialRequest {
	s.CredentialId = &v
	return s
}

func (s *DeleteSCIMServerCredentialRequest) SetDirectoryId(v string) *DeleteSCIMServerCredentialRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteSCIMServerCredentialRequest) Validate() error {
	return dara.Validate(s)
}
