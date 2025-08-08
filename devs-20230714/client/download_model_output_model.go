// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadModelOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DownloadModelOutputData) *DownloadModelOutput
	GetData() *DownloadModelOutputData
	SetErrCode(v string) *DownloadModelOutput
	GetErrCode() *string
	SetErrMsg(v string) *DownloadModelOutput
	GetErrMsg() *string
	SetRequestId(v string) *DownloadModelOutput
	GetRequestId() *string
	SetSuccess(v bool) *DownloadModelOutput
	GetSuccess() *bool
}

type DownloadModelOutput struct {
	Data    *DownloadModelOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                  `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                  `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DownloadModelOutput) String() string {
	return dara.Prettify(s)
}

func (s DownloadModelOutput) GoString() string {
	return s.String()
}

func (s *DownloadModelOutput) GetData() *DownloadModelOutputData {
	return s.Data
}

func (s *DownloadModelOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DownloadModelOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DownloadModelOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadModelOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DownloadModelOutput) SetData(v *DownloadModelOutputData) *DownloadModelOutput {
	s.Data = v
	return s
}

func (s *DownloadModelOutput) SetErrCode(v string) *DownloadModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DownloadModelOutput) SetErrMsg(v string) *DownloadModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DownloadModelOutput) SetRequestId(v string) *DownloadModelOutput {
	s.RequestId = &v
	return s
}

func (s *DownloadModelOutput) SetSuccess(v bool) *DownloadModelOutput {
	s.Success = &v
	return s
}

func (s *DownloadModelOutput) Validate() error {
	return dara.Validate(s)
}

type DownloadModelOutputData struct {
	ModelPath *string `json:"modelPath,omitempty" xml:"modelPath,omitempty"`
	TaskType  *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s DownloadModelOutputData) String() string {
	return dara.Prettify(s)
}

func (s DownloadModelOutputData) GoString() string {
	return s.String()
}

func (s *DownloadModelOutputData) GetModelPath() *string {
	return s.ModelPath
}

func (s *DownloadModelOutputData) GetTaskType() *string {
	return s.TaskType
}

func (s *DownloadModelOutputData) SetModelPath(v string) *DownloadModelOutputData {
	s.ModelPath = &v
	return s
}

func (s *DownloadModelOutputData) SetTaskType(v string) *DownloadModelOutputData {
	s.TaskType = &v
	return s
}

func (s *DownloadModelOutputData) Validate() error {
	return dara.Validate(s)
}
