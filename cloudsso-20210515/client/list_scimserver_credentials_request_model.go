// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSCIMServerCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListSCIMServerCredentialsRequest
	GetDirectoryId() *string
}

type ListSCIMServerCredentialsRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s ListSCIMServerCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSCIMServerCredentialsRequest) GoString() string {
	return s.String()
}

func (s *ListSCIMServerCredentialsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListSCIMServerCredentialsRequest) SetDirectoryId(v string) *ListSCIMServerCredentialsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListSCIMServerCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
