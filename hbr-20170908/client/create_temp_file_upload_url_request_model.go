// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTempFileUploadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreateTempFileUploadUrlRequest
	GetFileName() *string
}

type CreateTempFileUploadUrlRequest struct {
	// The name of the file to be uploaded.
	//
	// This parameter is required.
	//
	// example:
	//
	// file-list.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s CreateTempFileUploadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTempFileUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateTempFileUploadUrlRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateTempFileUploadUrlRequest) SetFileName(v string) *CreateTempFileUploadUrlRequest {
	s.FileName = &v
	return s
}

func (s *CreateTempFileUploadUrlRequest) Validate() error {
	return dara.Validate(s)
}
