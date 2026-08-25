// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *DeleteDirectoryRequest
	GetDirectoryId() *string
}

type DeleteDirectoryRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s DeleteDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDirectoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeleteDirectoryRequest) SetDirectoryId(v string) *DeleteDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
