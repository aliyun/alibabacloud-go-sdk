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
	// The description of the file.
	//
	// example:
	//
	// 测试使用
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// A custom file name. The name must be unique and serves as a unique identifier for the file. The name must meet the following requirements:
	//
	// 1. It must be 8 to 255 characters in length.
	//
	// 2. It can contain lowercase letters, digits, underscores (_), hyphens (-), and periods (.).
	//
	// 3. The first and last characters must be a letter or a digit.
	//
	// This parameter is required.
	//
	// example:
	//
	// mytest
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The MD5 hash of the file. This is used to verify the integrity of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 150b6083f50dd08159d45a0d5e4b56f9
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The download URL of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxx.xxx.xxx.tar
	OriginUrl *string `json:"OriginUrl,omitempty" xml:"OriginUrl,omitempty"`
	// The destination path on the service instance. This must be an absolute path to a file. You cannot specify only a folder. The parent folder of the destination path is restricted to the following locations:
	//
	// 1. /data/local
	//
	// 2. /data/user
	//
	// 3. /data/data
	//
	// 4. /data/cache
	//
	// 5. /data/tmp
	//
	// 6. /data/storage
	//
	// 7. /data/media/0
	//
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
