// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadTaskId(v string) *CreateDownloadUrlRequest
	GetDownloadTaskId() *string
	SetFileId(v string) *CreateDownloadUrlRequest
	GetFileId() *string
}

type CreateDownloadUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 137da36b41304bcd999a0a7895dc6881
	DownloadTaskId *string `json:"DownloadTaskId,omitempty" xml:"DownloadTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6f91885fa24b4c408d8f4eb392fd8ae6
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s CreateDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateDownloadUrlRequest) GetDownloadTaskId() *string {
	return s.DownloadTaskId
}

func (s *CreateDownloadUrlRequest) GetFileId() *string {
	return s.FileId
}

func (s *CreateDownloadUrlRequest) SetDownloadTaskId(v string) *CreateDownloadUrlRequest {
	s.DownloadTaskId = &v
	return s
}

func (s *CreateDownloadUrlRequest) SetFileId(v string) *CreateDownloadUrlRequest {
	s.FileId = &v
	return s
}

func (s *CreateDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
