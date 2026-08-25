// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearExternalSAMLIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ClearExternalSAMLIdentityProviderRequest
	GetDirectoryId() *string
}

type ClearExternalSAMLIdentityProviderRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s ClearExternalSAMLIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearExternalSAMLIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *ClearExternalSAMLIdentityProviderRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ClearExternalSAMLIdentityProviderRequest) SetDirectoryId(v string) *ClearExternalSAMLIdentityProviderRequest {
	s.DirectoryId = &v
	return s
}

func (s *ClearExternalSAMLIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
