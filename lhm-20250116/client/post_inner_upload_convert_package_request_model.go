// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostInnerUploadConvertPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileContentBase64(v string) *PostInnerUploadConvertPackageRequest
	GetFileContentBase64() *string
	SetFileName(v string) *PostInnerUploadConvertPackageRequest
	GetFileName() *string
	SetTaskId(v string) *PostInnerUploadConvertPackageRequest
	GetTaskId() *string
}

type PostInnerUploadConvertPackageRequest struct {
	// The file content, Base64-encoded.
	//
	// example:
	//
	// U0VMRUNUICogRlJPTSB0Ow==
	FileContentBase64 *string `json:"fileContentBase64,omitempty" xml:"fileContentBase64,omitempty"`
	// The file name.
	//
	// example:
	//
	// demo_file
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The task ID that uniquely identifies a task.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s PostInnerUploadConvertPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s PostInnerUploadConvertPackageRequest) GoString() string {
	return s.String()
}

func (s *PostInnerUploadConvertPackageRequest) GetFileContentBase64() *string {
	return s.FileContentBase64
}

func (s *PostInnerUploadConvertPackageRequest) GetFileName() *string {
	return s.FileName
}

func (s *PostInnerUploadConvertPackageRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *PostInnerUploadConvertPackageRequest) SetFileContentBase64(v string) *PostInnerUploadConvertPackageRequest {
	s.FileContentBase64 = &v
	return s
}

func (s *PostInnerUploadConvertPackageRequest) SetFileName(v string) *PostInnerUploadConvertPackageRequest {
	s.FileName = &v
	return s
}

func (s *PostInnerUploadConvertPackageRequest) SetTaskId(v string) *PostInnerUploadConvertPackageRequest {
	s.TaskId = &v
	return s
}

func (s *PostInnerUploadConvertPackageRequest) Validate() error {
	return dara.Validate(s)
}
