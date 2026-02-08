// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GenerateUploadUrlRequest
	GetFileName() *string
	SetKey(v string) *GenerateUploadUrlRequest
	GetKey() *string
}

type GenerateUploadUrlRequest struct {
	// File name, including the extension.
	//
	// example:
	//
	// faaf8508-9542-4ac4-84a2-0ddcbb5f79a6 (2).json
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Upload URL key
	//
	// example:
	//
	// 3b9b5dc6d67ee9fa
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s GenerateUploadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadUrlRequest) GetFileName() *string {
	return s.FileName
}

func (s *GenerateUploadUrlRequest) GetKey() *string {
	return s.Key
}

func (s *GenerateUploadUrlRequest) SetFileName(v string) *GenerateUploadUrlRequest {
	s.FileName = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetKey(v string) *GenerateUploadUrlRequest {
	s.Key = &v
	return s
}

func (s *GenerateUploadUrlRequest) Validate() error {
	return dara.Validate(s)
}
