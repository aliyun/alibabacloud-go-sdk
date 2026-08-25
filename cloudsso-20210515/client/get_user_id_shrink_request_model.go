// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetUserIdShrinkRequest
	GetDirectoryId() *string
	SetExternalIdShrink(v string) *GetUserIdShrinkRequest
	GetExternalIdShrink() *string
}

type GetUserIdShrinkRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The identifier information about the user that is synchronized from an external identity provider (IdP).
	ExternalIdShrink *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
}

func (s GetUserIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetUserIdShrinkRequest) GetExternalIdShrink() *string {
	return s.ExternalIdShrink
}

func (s *GetUserIdShrinkRequest) SetDirectoryId(v string) *GetUserIdShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserIdShrinkRequest) SetExternalIdShrink(v string) *GetUserIdShrinkRequest {
	s.ExternalIdShrink = &v
	return s
}

func (s *GetUserIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
