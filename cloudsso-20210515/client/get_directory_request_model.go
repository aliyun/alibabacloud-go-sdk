// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetDirectoryRequest
	GetDirectoryId() *string
}

type GetDirectoryRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryRequest) GoString() string {
	return s.String()
}

func (s *GetDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetDirectoryRequest) SetDirectoryId(v string) *GetDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
