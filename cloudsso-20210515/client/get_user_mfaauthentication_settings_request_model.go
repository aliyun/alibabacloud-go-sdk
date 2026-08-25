// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserMFAAuthenticationSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetUserMFAAuthenticationSettingsRequest
	GetDirectoryId() *string
	SetUserId(v string) *GetUserMFAAuthenticationSettingsRequest
	GetUserId() *string
}

type GetUserMFAAuthenticationSettingsRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserMFAAuthenticationSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserMFAAuthenticationSettingsRequest) GoString() string {
	return s.String()
}

func (s *GetUserMFAAuthenticationSettingsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetUserMFAAuthenticationSettingsRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetUserMFAAuthenticationSettingsRequest) SetDirectoryId(v string) *GetUserMFAAuthenticationSettingsRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserMFAAuthenticationSettingsRequest) SetUserId(v string) *GetUserMFAAuthenticationSettingsRequest {
	s.UserId = &v
	return s
}

func (s *GetUserMFAAuthenticationSettingsRequest) Validate() error {
	return dara.Validate(s)
}
