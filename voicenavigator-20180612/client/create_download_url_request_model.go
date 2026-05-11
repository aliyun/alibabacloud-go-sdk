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
	// 073f092da0a847b9bf76eb88b5931c7a
	DownloadTaskId *string `json:"DownloadTaskId,omitempty" xml:"DownloadTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 22626c39603744f5a08d4d715315561a
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
