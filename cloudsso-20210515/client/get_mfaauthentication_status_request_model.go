// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMFAAuthenticationStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetMFAAuthenticationStatusRequest
	GetDirectoryId() *string
}

type GetMFAAuthenticationStatusRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetMFAAuthenticationStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMFAAuthenticationStatusRequest) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationStatusRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetMFAAuthenticationStatusRequest) SetDirectoryId(v string) *GetMFAAuthenticationStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetMFAAuthenticationStatusRequest) Validate() error {
	return dara.Validate(s)
}
