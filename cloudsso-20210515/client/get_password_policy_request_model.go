// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetPasswordPolicyRequest
	GetDirectoryId() *string
}

type GetPasswordPolicyRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetPasswordPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetPasswordPolicyRequest) SetDirectoryId(v string) *GetPasswordPolicyRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetPasswordPolicyRequest) Validate() error {
	return dara.Validate(s)
}
