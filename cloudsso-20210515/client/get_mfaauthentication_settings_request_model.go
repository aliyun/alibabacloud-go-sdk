// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMFAAuthenticationSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetMFAAuthenticationSettingsRequest
	GetDirectoryId() *string
}

type GetMFAAuthenticationSettingsRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetMFAAuthenticationSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMFAAuthenticationSettingsRequest) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetMFAAuthenticationSettingsRequest) SetDirectoryId(v string) *GetMFAAuthenticationSettingsRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetMFAAuthenticationSettingsRequest) Validate() error {
	return dara.Validate(s)
}
