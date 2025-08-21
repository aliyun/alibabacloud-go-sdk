// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UploadFileRequest
	GetDescription() *string
	SetFileName(v string) *UploadFileRequest
	GetFileName() *string
	SetMd5(v string) *UploadFileRequest
	GetMd5() *string
	SetOriginUrl(v string) *UploadFileRequest
	GetOriginUrl() *string
	SetTargetPath(v string) *UploadFileRequest
	GetTargetPath() *string
}

type UploadFileRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mytest
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 150b6083f50dd08159d45a0d5e4b56f9
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxx.xxx.xxx.tar
	OriginUrl *string `json:"OriginUrl,omitempty" xml:"OriginUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /data/tmp/test/xxx.tar
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
}

func (s UploadFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadFileRequest) GoString() string {
	return s.String()
}

func (s *UploadFileRequest) GetDescription() *string {
	return s.Description
}

func (s *UploadFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadFileRequest) GetMd5() *string {
	return s.Md5
}

func (s *UploadFileRequest) GetOriginUrl() *string {
	return s.OriginUrl
}

func (s *UploadFileRequest) GetTargetPath() *string {
	return s.TargetPath
}

func (s *UploadFileRequest) SetDescription(v string) *UploadFileRequest {
	s.Description = &v
	return s
}

func (s *UploadFileRequest) SetFileName(v string) *UploadFileRequest {
	s.FileName = &v
	return s
}

func (s *UploadFileRequest) SetMd5(v string) *UploadFileRequest {
	s.Md5 = &v
	return s
}

func (s *UploadFileRequest) SetOriginUrl(v string) *UploadFileRequest {
	s.OriginUrl = &v
	return s
}

func (s *UploadFileRequest) SetTargetPath(v string) *UploadFileRequest {
	s.TargetPath = &v
	return s
}

func (s *UploadFileRequest) Validate() error {
	return dara.Validate(s)
}
