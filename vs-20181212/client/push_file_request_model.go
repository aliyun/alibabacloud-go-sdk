// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *PushFileRequest
	GetFileId() *string
	SetRenderingInstanceId(v string) *PushFileRequest
	GetRenderingInstanceId() *string
}

type PushFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f-1671accd4dafdag3er256cvgewt13f7141db2f7
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s PushFileRequest) String() string {
	return dara.Prettify(s)
}

func (s PushFileRequest) GoString() string {
	return s.String()
}

func (s *PushFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *PushFileRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *PushFileRequest) SetFileId(v string) *PushFileRequest {
	s.FileId = &v
	return s
}

func (s *PushFileRequest) SetRenderingInstanceId(v string) *PushFileRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *PushFileRequest) Validate() error {
	return dara.Validate(s)
}
