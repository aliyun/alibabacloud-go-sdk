// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *DeleteFileRequest
	GetFileId() *string
}

type DeleteFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f-167131acd45omat771813f7141d28db2f7
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *DeleteFileRequest) SetFileId(v string) *DeleteFileRequest {
	s.FileId = &v
	return s
}

func (s *DeleteFileRequest) Validate() error {
	return dara.Validate(s)
}
