// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoCreationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreationInstruction(v *CreateVideoCreationTaskRequestCreationInstruction) *CreateVideoCreationTaskRequest
	GetCreationInstruction() *CreateVideoCreationTaskRequestCreationInstruction
	SetFileInfo(v *CreateVideoCreationTaskRequestFileInfo) *CreateVideoCreationTaskRequest
	GetFileInfo() *CreateVideoCreationTaskRequestFileInfo
	SetImageDetectionTaskId(v string) *CreateVideoCreationTaskRequest
	GetImageDetectionTaskId() *string
	SetRequestId(v string) *CreateVideoCreationTaskRequest
	GetRequestId() *string
	SetUserId(v string) *CreateVideoCreationTaskRequest
	GetUserId() *string
}

type CreateVideoCreationTaskRequest struct {
	CreationInstruction *CreateVideoCreationTaskRequestCreationInstruction `json:"creationInstruction,omitempty" xml:"creationInstruction,omitempty" type:"Struct"`
	FileInfo            *CreateVideoCreationTaskRequestFileInfo            `json:"fileInfo,omitempty" xml:"fileInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 111
	ImageDetectionTaskId *string `json:"imageDetectionTaskId,omitempty" xml:"imageDetectionTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateVideoCreationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoCreationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoCreationTaskRequest) GetCreationInstruction() *CreateVideoCreationTaskRequestCreationInstruction {
	return s.CreationInstruction
}

func (s *CreateVideoCreationTaskRequest) GetFileInfo() *CreateVideoCreationTaskRequestFileInfo {
	return s.FileInfo
}

func (s *CreateVideoCreationTaskRequest) GetImageDetectionTaskId() *string {
	return s.ImageDetectionTaskId
}

func (s *CreateVideoCreationTaskRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVideoCreationTaskRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateVideoCreationTaskRequest) SetCreationInstruction(v *CreateVideoCreationTaskRequestCreationInstruction) *CreateVideoCreationTaskRequest {
	s.CreationInstruction = v
	return s
}

func (s *CreateVideoCreationTaskRequest) SetFileInfo(v *CreateVideoCreationTaskRequestFileInfo) *CreateVideoCreationTaskRequest {
	s.FileInfo = v
	return s
}

func (s *CreateVideoCreationTaskRequest) SetImageDetectionTaskId(v string) *CreateVideoCreationTaskRequest {
	s.ImageDetectionTaskId = &v
	return s
}

func (s *CreateVideoCreationTaskRequest) SetRequestId(v string) *CreateVideoCreationTaskRequest {
	s.RequestId = &v
	return s
}

func (s *CreateVideoCreationTaskRequest) SetUserId(v string) *CreateVideoCreationTaskRequest {
	s.UserId = &v
	return s
}

func (s *CreateVideoCreationTaskRequest) Validate() error {
	if s.CreationInstruction != nil {
		if err := s.CreationInstruction.Validate(); err != nil {
			return err
		}
	}
	if s.FileInfo != nil {
		if err := s.FileInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVideoCreationTaskRequestCreationInstruction struct {
	// example:
	//
	// xxx
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// example:
	//
	// xxx
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s CreateVideoCreationTaskRequestCreationInstruction) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoCreationTaskRequestCreationInstruction) GoString() string {
	return s.String()
}

func (s *CreateVideoCreationTaskRequestCreationInstruction) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateVideoCreationTaskRequestCreationInstruction) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateVideoCreationTaskRequestCreationInstruction) SetInstruction(v string) *CreateVideoCreationTaskRequestCreationInstruction {
	s.Instruction = &v
	return s
}

func (s *CreateVideoCreationTaskRequestCreationInstruction) SetTemplateId(v string) *CreateVideoCreationTaskRequestCreationInstruction {
	s.TemplateId = &v
	return s
}

func (s *CreateVideoCreationTaskRequestCreationInstruction) Validate() error {
	return dara.Validate(s)
}

type CreateVideoCreationTaskRequestFileInfo struct {
	// example:
	//
	// xxx
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// xxx
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// xxx
	FileTraceId *string `json:"fileTraceId,omitempty" xml:"fileTraceId,omitempty"`
	// example:
	//
	// xxx
	OssKey *string `json:"ossKey,omitempty" xml:"ossKey,omitempty"`
}

func (s CreateVideoCreationTaskRequestFileInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoCreationTaskRequestFileInfo) GoString() string {
	return s.String()
}

func (s *CreateVideoCreationTaskRequestFileInfo) GetFileId() *string {
	return s.FileId
}

func (s *CreateVideoCreationTaskRequestFileInfo) GetFileName() *string {
	return s.FileName
}

func (s *CreateVideoCreationTaskRequestFileInfo) GetFileTraceId() *string {
	return s.FileTraceId
}

func (s *CreateVideoCreationTaskRequestFileInfo) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateVideoCreationTaskRequestFileInfo) SetFileId(v string) *CreateVideoCreationTaskRequestFileInfo {
	s.FileId = &v
	return s
}

func (s *CreateVideoCreationTaskRequestFileInfo) SetFileName(v string) *CreateVideoCreationTaskRequestFileInfo {
	s.FileName = &v
	return s
}

func (s *CreateVideoCreationTaskRequestFileInfo) SetFileTraceId(v string) *CreateVideoCreationTaskRequestFileInfo {
	s.FileTraceId = &v
	return s
}

func (s *CreateVideoCreationTaskRequestFileInfo) SetOssKey(v string) *CreateVideoCreationTaskRequestFileInfo {
	s.OssKey = &v
	return s
}

func (s *CreateVideoCreationTaskRequestFileInfo) Validate() error {
	return dara.Validate(s)
}
