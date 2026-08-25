// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExternalSAMLIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetExternalSAMLIdentityProviderRequest
	GetDirectoryId() *string
}

type GetExternalSAMLIdentityProviderRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetExternalSAMLIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExternalSAMLIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *GetExternalSAMLIdentityProviderRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetExternalSAMLIdentityProviderRequest) SetDirectoryId(v string) *GetExternalSAMLIdentityProviderRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
