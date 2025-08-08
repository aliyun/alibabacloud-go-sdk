// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelStatusOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetModelStatusOutputData) *GetModelStatusOutput
	GetData() *GetModelStatusOutputData
	SetErrCode(v string) *GetModelStatusOutput
	GetErrCode() *string
	SetErrMsg(v string) *GetModelStatusOutput
	GetErrMsg() *string
	SetRequestId(v string) *GetModelStatusOutput
	GetRequestId() *string
	SetSuccess(v bool) *GetModelStatusOutput
	GetSuccess() *bool
}

type GetModelStatusOutput struct {
	Data      *GetModelStatusOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode   *string                   `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg    *string                   `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	RequestId *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetModelStatusOutput) String() string {
	return dara.Prettify(s)
}

func (s GetModelStatusOutput) GoString() string {
	return s.String()
}

func (s *GetModelStatusOutput) GetData() *GetModelStatusOutputData {
	return s.Data
}

func (s *GetModelStatusOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetModelStatusOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *GetModelStatusOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelStatusOutput) GetSuccess() *bool {
	return s.Success
}

func (s *GetModelStatusOutput) SetData(v *GetModelStatusOutputData) *GetModelStatusOutput {
	s.Data = v
	return s
}

func (s *GetModelStatusOutput) SetErrCode(v string) *GetModelStatusOutput {
	s.ErrCode = &v
	return s
}

func (s *GetModelStatusOutput) SetErrMsg(v string) *GetModelStatusOutput {
	s.ErrMsg = &v
	return s
}

func (s *GetModelStatusOutput) SetRequestId(v string) *GetModelStatusOutput {
	s.RequestId = &v
	return s
}

func (s *GetModelStatusOutput) SetSuccess(v bool) *GetModelStatusOutput {
	s.Success = &v
	return s
}

func (s *GetModelStatusOutput) Validate() error {
	return dara.Validate(s)
}

type GetModelStatusOutputData struct {
	CurrentBytes *int64  `json:"currentBytes,omitempty" xml:"currentBytes,omitempty"`
	ErrMessage   *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	FileSize     *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Finished     *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	FinishedTime *int64  `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	Speed        *int64  `json:"speed,omitempty" xml:"speed,omitempty"`
	StartTime    *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Total        *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetModelStatusOutputData) String() string {
	return dara.Prettify(s)
}

func (s GetModelStatusOutputData) GoString() string {
	return s.String()
}

func (s *GetModelStatusOutputData) GetCurrentBytes() *int64 {
	return s.CurrentBytes
}

func (s *GetModelStatusOutputData) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetModelStatusOutputData) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetModelStatusOutputData) GetFinished() *bool {
	return s.Finished
}

func (s *GetModelStatusOutputData) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *GetModelStatusOutputData) GetSpeed() *int64 {
	return s.Speed
}

func (s *GetModelStatusOutputData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetModelStatusOutputData) GetTotal() *int64 {
	return s.Total
}

func (s *GetModelStatusOutputData) SetCurrentBytes(v int64) *GetModelStatusOutputData {
	s.CurrentBytes = &v
	return s
}

func (s *GetModelStatusOutputData) SetErrMessage(v string) *GetModelStatusOutputData {
	s.ErrMessage = &v
	return s
}

func (s *GetModelStatusOutputData) SetFileSize(v int64) *GetModelStatusOutputData {
	s.FileSize = &v
	return s
}

func (s *GetModelStatusOutputData) SetFinished(v bool) *GetModelStatusOutputData {
	s.Finished = &v
	return s
}

func (s *GetModelStatusOutputData) SetFinishedTime(v int64) *GetModelStatusOutputData {
	s.FinishedTime = &v
	return s
}

func (s *GetModelStatusOutputData) SetSpeed(v int64) *GetModelStatusOutputData {
	s.Speed = &v
	return s
}

func (s *GetModelStatusOutputData) SetStartTime(v int64) *GetModelStatusOutputData {
	s.StartTime = &v
	return s
}

func (s *GetModelStatusOutputData) SetTotal(v int64) *GetModelStatusOutputData {
	s.Total = &v
	return s
}

func (s *GetModelStatusOutputData) Validate() error {
	return dara.Validate(s)
}
