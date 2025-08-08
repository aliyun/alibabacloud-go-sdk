// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployModelScopeModelOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeployModelScopeModelOutputData) *DeployModelScopeModelOutput
	GetData() *DeployModelScopeModelOutputData
	SetErrCode(v string) *DeployModelScopeModelOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeployModelScopeModelOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeployModelScopeModelOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeployModelScopeModelOutput
	GetSuccess() *bool
}

type DeployModelScopeModelOutput struct {
	Data    *DeployModelScopeModelOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                          `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                          `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployModelScopeModelOutput) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelOutput) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelOutput) GetData() *DeployModelScopeModelOutputData {
	return s.Data
}

func (s *DeployModelScopeModelOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeployModelScopeModelOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeployModelScopeModelOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployModelScopeModelOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeployModelScopeModelOutput) SetData(v *DeployModelScopeModelOutputData) *DeployModelScopeModelOutput {
	s.Data = v
	return s
}

func (s *DeployModelScopeModelOutput) SetErrCode(v string) *DeployModelScopeModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployModelScopeModelOutput) SetErrMsg(v string) *DeployModelScopeModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployModelScopeModelOutput) SetRequestId(v string) *DeployModelScopeModelOutput {
	s.RequestId = &v
	return s
}

func (s *DeployModelScopeModelOutput) SetSuccess(v bool) *DeployModelScopeModelOutput {
	s.Success = &v
	return s
}

func (s *DeployModelScopeModelOutput) Validate() error {
	return dara.Validate(s)
}

type DeployModelScopeModelOutputData struct {
	DeploymentTaskID *string `json:"deploymentTaskID,omitempty" xml:"deploymentTaskID,omitempty"`
	ErrorMessage     *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Finished         *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TaskType         *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TraceID          *string `json:"traceID,omitempty" xml:"traceID,omitempty"`
	UrlInternet      *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet      *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
}

func (s DeployModelScopeModelOutputData) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelOutputData) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelOutputData) GetDeploymentTaskID() *string {
	return s.DeploymentTaskID
}

func (s *DeployModelScopeModelOutputData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeployModelScopeModelOutputData) GetFinished() *bool {
	return s.Finished
}

func (s *DeployModelScopeModelOutputData) GetServiceName() *string {
	return s.ServiceName
}

func (s *DeployModelScopeModelOutputData) GetTaskType() *string {
	return s.TaskType
}

func (s *DeployModelScopeModelOutputData) GetTraceID() *string {
	return s.TraceID
}

func (s *DeployModelScopeModelOutputData) GetUrlInternet() *string {
	return s.UrlInternet
}

func (s *DeployModelScopeModelOutputData) GetUrlIntranet() *string {
	return s.UrlIntranet
}

func (s *DeployModelScopeModelOutputData) SetDeploymentTaskID(v string) *DeployModelScopeModelOutputData {
	s.DeploymentTaskID = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetErrorMessage(v string) *DeployModelScopeModelOutputData {
	s.ErrorMessage = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetFinished(v bool) *DeployModelScopeModelOutputData {
	s.Finished = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetServiceName(v string) *DeployModelScopeModelOutputData {
	s.ServiceName = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetTaskType(v string) *DeployModelScopeModelOutputData {
	s.TaskType = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetTraceID(v string) *DeployModelScopeModelOutputData {
	s.TraceID = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetUrlInternet(v string) *DeployModelScopeModelOutputData {
	s.UrlInternet = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetUrlIntranet(v string) *DeployModelScopeModelOutputData {
	s.UrlIntranet = &v
	return s
}

func (s *DeployModelScopeModelOutputData) Validate() error {
	return dara.Validate(s)
}
