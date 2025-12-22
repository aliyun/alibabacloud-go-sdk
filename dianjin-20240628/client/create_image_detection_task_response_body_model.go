// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageDetectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateImageDetectionTaskResponseBody
	GetCode() *string
	SetData(v *CreateImageDetectionTaskResponseBodyData) *CreateImageDetectionTaskResponseBody
	GetData() *CreateImageDetectionTaskResponseBodyData
	SetMessage(v string) *CreateImageDetectionTaskResponseBody
	GetMessage() *string
	SetRetryAble(v bool) *CreateImageDetectionTaskResponseBody
	GetRetryAble() *bool
	SetSuccess(v bool) *CreateImageDetectionTaskResponseBody
	GetSuccess() *bool
}

type CreateImageDetectionTaskResponseBody struct {
	// example:
	//
	// 0
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateImageDetectionTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 成功
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RetryAble *bool   `json:"retryAble,omitempty" xml:"retryAble,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateImageDetectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageDetectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageDetectionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateImageDetectionTaskResponseBody) GetData() *CreateImageDetectionTaskResponseBodyData {
	return s.Data
}

func (s *CreateImageDetectionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateImageDetectionTaskResponseBody) GetRetryAble() *bool {
	return s.RetryAble
}

func (s *CreateImageDetectionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateImageDetectionTaskResponseBody) SetCode(v string) *CreateImageDetectionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBody) SetData(v *CreateImageDetectionTaskResponseBodyData) *CreateImageDetectionTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateImageDetectionTaskResponseBody) SetMessage(v string) *CreateImageDetectionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBody) SetRetryAble(v bool) *CreateImageDetectionTaskResponseBody {
	s.RetryAble = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBody) SetSuccess(v bool) *CreateImageDetectionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateImageDetectionTaskResponseBodyData struct {
	FileInfo *CreateImageDetectionTaskResponseBodyDataFileInfo `json:"fileInfo,omitempty" xml:"fileInfo,omitempty" type:"Struct"`
	// example:
	//
	// a1112229
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// w4paqoezm2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateImageDetectionTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateImageDetectionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateImageDetectionTaskResponseBodyData) GetFileInfo() *CreateImageDetectionTaskResponseBodyDataFileInfo {
	return s.FileInfo
}

func (s *CreateImageDetectionTaskResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageDetectionTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateImageDetectionTaskResponseBodyData) SetFileInfo(v *CreateImageDetectionTaskResponseBodyDataFileInfo) *CreateImageDetectionTaskResponseBodyData {
	s.FileInfo = v
	return s
}

func (s *CreateImageDetectionTaskResponseBodyData) SetRequestId(v string) *CreateImageDetectionTaskResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBodyData) SetTaskId(v string) *CreateImageDetectionTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBodyData) Validate() error {
	if s.FileInfo != nil {
		if err := s.FileInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateImageDetectionTaskResponseBodyDataFileInfo struct {
	// example:
	//
	// 1
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 1
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 1
	FileTraceId *string `json:"fileTraceId,omitempty" xml:"fileTraceId,omitempty"`
	// example:
	//
	// 1
	OssKey *string `json:"ossKey,omitempty" xml:"ossKey,omitempty"`
}

func (s CreateImageDetectionTaskResponseBodyDataFileInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateImageDetectionTaskResponseBodyDataFileInfo) GoString() string {
	return s.String()
}

func (s *CreateImageDetectionTaskResponseBodyDataFileInfo) GetFileId() *string {
	return s.FileId
}

func (s *CreateImageDetectionTaskResponseBodyDataFileInfo) GetFileName() *string {
	return s.FileName
}

func (s *CreateImageDetectionTaskResponseBodyDataFileInfo) GetFileTraceId() *string {
	return s.FileTraceId
}

func (s *CreateImageDetectionTaskResponseBodyDataFileInfo) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateImageDetectionTaskResponseBodyDataFileInfo) SetFileId(v string) *CreateImageDetectionTaskResponseBodyDataFileInfo {
	s.FileId = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBodyDataFileInfo) SetFileName(v string) *CreateImageDetectionTaskResponseBodyDataFileInfo {
	s.FileName = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBodyDataFileInfo) SetFileTraceId(v string) *CreateImageDetectionTaskResponseBodyDataFileInfo {
	s.FileTraceId = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBodyDataFileInfo) SetOssKey(v string) *CreateImageDetectionTaskResponseBodyDataFileInfo {
	s.OssKey = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBodyDataFileInfo) Validate() error {
	return dara.Validate(s)
}
