// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilePutUserTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *FilePutUserTagsResponseBody
	GetFileId() *string
}

type FilePutUserTagsResponseBody struct {
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s FilePutUserTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FilePutUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *FilePutUserTagsResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *FilePutUserTagsResponseBody) SetFileId(v string) *FilePutUserTagsResponseBody {
	s.FileId = &v
	return s
}

func (s *FilePutUserTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
