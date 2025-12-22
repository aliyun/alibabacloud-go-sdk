// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageDetectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileInfo(v *CreateImageDetectionTaskRequestFileInfo) *CreateImageDetectionTaskRequest
	GetFileInfo() *CreateImageDetectionTaskRequestFileInfo
	SetFileUrl(v string) *CreateImageDetectionTaskRequest
	GetFileUrl() *string
	SetRequestId(v string) *CreateImageDetectionTaskRequest
	GetRequestId() *string
	SetUserId(v string) *CreateImageDetectionTaskRequest
	GetUserId() *string
}

type CreateImageDetectionTaskRequest struct {
	FileInfo *CreateImageDetectionTaskRequestFileInfo `json:"fileInfo,omitempty" xml:"fileInfo,omitempty" type:"Struct"`
	// example:
	//
	// fileUrl
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sfkhwjf
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateImageDetectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageDetectionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageDetectionTaskRequest) GetFileInfo() *CreateImageDetectionTaskRequestFileInfo {
	return s.FileInfo
}

func (s *CreateImageDetectionTaskRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateImageDetectionTaskRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageDetectionTaskRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateImageDetectionTaskRequest) SetFileInfo(v *CreateImageDetectionTaskRequestFileInfo) *CreateImageDetectionTaskRequest {
	s.FileInfo = v
	return s
}

func (s *CreateImageDetectionTaskRequest) SetFileUrl(v string) *CreateImageDetectionTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateImageDetectionTaskRequest) SetRequestId(v string) *CreateImageDetectionTaskRequest {
	s.RequestId = &v
	return s
}

func (s *CreateImageDetectionTaskRequest) SetUserId(v string) *CreateImageDetectionTaskRequest {
	s.UserId = &v
	return s
}

func (s *CreateImageDetectionTaskRequest) Validate() error {
	if s.FileInfo != nil {
		if err := s.FileInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateImageDetectionTaskRequestFileInfo struct {
	// example:
	//
	// wanx-demo-1.png
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// wanx-demo-1.png
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// fileTraceId
	FileTraceId *string `json:"fileTraceId,omitempty" xml:"fileTraceId,omitempty"`
	// example:
	//
	// wanx-demo-1.png
	OssKey *string `json:"ossKey,omitempty" xml:"ossKey,omitempty"`
}

func (s CreateImageDetectionTaskRequestFileInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateImageDetectionTaskRequestFileInfo) GoString() string {
	return s.String()
}

func (s *CreateImageDetectionTaskRequestFileInfo) GetFileId() *string {
	return s.FileId
}

func (s *CreateImageDetectionTaskRequestFileInfo) GetFileName() *string {
	return s.FileName
}

func (s *CreateImageDetectionTaskRequestFileInfo) GetFileTraceId() *string {
	return s.FileTraceId
}

func (s *CreateImageDetectionTaskRequestFileInfo) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateImageDetectionTaskRequestFileInfo) SetFileId(v string) *CreateImageDetectionTaskRequestFileInfo {
	s.FileId = &v
	return s
}

func (s *CreateImageDetectionTaskRequestFileInfo) SetFileName(v string) *CreateImageDetectionTaskRequestFileInfo {
	s.FileName = &v
	return s
}

func (s *CreateImageDetectionTaskRequestFileInfo) SetFileTraceId(v string) *CreateImageDetectionTaskRequestFileInfo {
	s.FileTraceId = &v
	return s
}

func (s *CreateImageDetectionTaskRequestFileInfo) SetOssKey(v string) *CreateImageDetectionTaskRequestFileInfo {
	s.OssKey = &v
	return s
}

func (s *CreateImageDetectionTaskRequestFileInfo) Validate() error {
	return dara.Validate(s)
}
