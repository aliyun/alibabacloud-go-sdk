// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployHuggingFaceModelOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeployHuggingFaceModelOutputData) *DeployHuggingFaceModelOutput
	GetData() *DeployHuggingFaceModelOutputData
	SetErrCode(v string) *DeployHuggingFaceModelOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeployHuggingFaceModelOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeployHuggingFaceModelOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeployHuggingFaceModelOutput
	GetSuccess() *bool
}

type DeployHuggingFaceModelOutput struct {
	Data    *DeployHuggingFaceModelOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                           `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                           `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployHuggingFaceModelOutput) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelOutput) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelOutput) GetData() *DeployHuggingFaceModelOutputData {
	return s.Data
}

func (s *DeployHuggingFaceModelOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeployHuggingFaceModelOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeployHuggingFaceModelOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployHuggingFaceModelOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeployHuggingFaceModelOutput) SetData(v *DeployHuggingFaceModelOutputData) *DeployHuggingFaceModelOutput {
	s.Data = v
	return s
}

func (s *DeployHuggingFaceModelOutput) SetErrCode(v string) *DeployHuggingFaceModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployHuggingFaceModelOutput) SetErrMsg(v string) *DeployHuggingFaceModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployHuggingFaceModelOutput) SetRequestId(v string) *DeployHuggingFaceModelOutput {
	s.RequestId = &v
	return s
}

func (s *DeployHuggingFaceModelOutput) SetSuccess(v bool) *DeployHuggingFaceModelOutput {
	s.Success = &v
	return s
}

func (s *DeployHuggingFaceModelOutput) Validate() error {
	return dara.Validate(s)
}

type DeployHuggingFaceModelOutputData struct {
	DeploymentTaskID *string `json:"deploymentTaskID,omitempty" xml:"deploymentTaskID,omitempty"`
	ErrorMessage     *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Finished         *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TaskType         *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TraceID          *string `json:"traceID,omitempty" xml:"traceID,omitempty"`
	UrlInternet      *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet      *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
}

func (s DeployHuggingFaceModelOutputData) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelOutputData) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelOutputData) GetDeploymentTaskID() *string {
	return s.DeploymentTaskID
}

func (s *DeployHuggingFaceModelOutputData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeployHuggingFaceModelOutputData) GetFinished() *bool {
	return s.Finished
}

func (s *DeployHuggingFaceModelOutputData) GetServiceName() *string {
	return s.ServiceName
}

func (s *DeployHuggingFaceModelOutputData) GetTaskType() *string {
	return s.TaskType
}

func (s *DeployHuggingFaceModelOutputData) GetTraceID() *string {
	return s.TraceID
}

func (s *DeployHuggingFaceModelOutputData) GetUrlInternet() *string {
	return s.UrlInternet
}

func (s *DeployHuggingFaceModelOutputData) GetUrlIntranet() *string {
	return s.UrlIntranet
}

func (s *DeployHuggingFaceModelOutputData) SetDeploymentTaskID(v string) *DeployHuggingFaceModelOutputData {
	s.DeploymentTaskID = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetErrorMessage(v string) *DeployHuggingFaceModelOutputData {
	s.ErrorMessage = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetFinished(v bool) *DeployHuggingFaceModelOutputData {
	s.Finished = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetServiceName(v string) *DeployHuggingFaceModelOutputData {
	s.ServiceName = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetTaskType(v string) *DeployHuggingFaceModelOutputData {
	s.TaskType = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetTraceID(v string) *DeployHuggingFaceModelOutputData {
	s.TraceID = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetUrlInternet(v string) *DeployHuggingFaceModelOutputData {
	s.UrlInternet = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetUrlIntranet(v string) *DeployHuggingFaceModelOutputData {
	s.UrlIntranet = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) Validate() error {
	return dara.Validate(s)
}
