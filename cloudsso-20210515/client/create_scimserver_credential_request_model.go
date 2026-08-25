// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSCIMServerCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *CreateSCIMServerCredentialRequest
	GetDirectoryId() *string
}

type CreateSCIMServerCredentialRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s CreateSCIMServerCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSCIMServerCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateSCIMServerCredentialRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateSCIMServerCredentialRequest) SetDirectoryId(v string) *CreateSCIMServerCredentialRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateSCIMServerCredentialRequest) Validate() error {
	return dara.Validate(s)
}
