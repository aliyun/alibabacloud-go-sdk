// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployOllamaModelOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeployOllamaModelOutputData) *DeployOllamaModelOutput
	GetData() *DeployOllamaModelOutputData
	SetErrCode(v string) *DeployOllamaModelOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeployOllamaModelOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeployOllamaModelOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeployOllamaModelOutput
	GetSuccess() *bool
}

type DeployOllamaModelOutput struct {
	Data    *DeployOllamaModelOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                      `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                      `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployOllamaModelOutput) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelOutput) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelOutput) GetData() *DeployOllamaModelOutputData {
	return s.Data
}

func (s *DeployOllamaModelOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeployOllamaModelOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeployOllamaModelOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployOllamaModelOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeployOllamaModelOutput) SetData(v *DeployOllamaModelOutputData) *DeployOllamaModelOutput {
	s.Data = v
	return s
}

func (s *DeployOllamaModelOutput) SetErrCode(v string) *DeployOllamaModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployOllamaModelOutput) SetErrMsg(v string) *DeployOllamaModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployOllamaModelOutput) SetRequestId(v string) *DeployOllamaModelOutput {
	s.RequestId = &v
	return s
}

func (s *DeployOllamaModelOutput) SetSuccess(v bool) *DeployOllamaModelOutput {
	s.Success = &v
	return s
}

func (s *DeployOllamaModelOutput) Validate() error {
	return dara.Validate(s)
}

type DeployOllamaModelOutputData struct {
	DeploymentTaskID *string `json:"deploymentTaskID,omitempty" xml:"deploymentTaskID,omitempty"`
	ErrorMessage     *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Finished         *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	ModelName        *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TraceID          *string `json:"traceID,omitempty" xml:"traceID,omitempty"`
	UrlInternet      *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet      *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
}

func (s DeployOllamaModelOutputData) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelOutputData) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelOutputData) GetDeploymentTaskID() *string {
	return s.DeploymentTaskID
}

func (s *DeployOllamaModelOutputData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeployOllamaModelOutputData) GetFinished() *bool {
	return s.Finished
}

func (s *DeployOllamaModelOutputData) GetModelName() *string {
	return s.ModelName
}

func (s *DeployOllamaModelOutputData) GetServiceName() *string {
	return s.ServiceName
}

func (s *DeployOllamaModelOutputData) GetTraceID() *string {
	return s.TraceID
}

func (s *DeployOllamaModelOutputData) GetUrlInternet() *string {
	return s.UrlInternet
}

func (s *DeployOllamaModelOutputData) GetUrlIntranet() *string {
	return s.UrlIntranet
}

func (s *DeployOllamaModelOutputData) SetDeploymentTaskID(v string) *DeployOllamaModelOutputData {
	s.DeploymentTaskID = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetErrorMessage(v string) *DeployOllamaModelOutputData {
	s.ErrorMessage = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetFinished(v bool) *DeployOllamaModelOutputData {
	s.Finished = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetModelName(v string) *DeployOllamaModelOutputData {
	s.ModelName = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetServiceName(v string) *DeployOllamaModelOutputData {
	s.ServiceName = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetTraceID(v string) *DeployOllamaModelOutputData {
	s.TraceID = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetUrlInternet(v string) *DeployOllamaModelOutputData {
	s.UrlInternet = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetUrlIntranet(v string) *DeployOllamaModelOutputData {
	s.UrlIntranet = &v
	return s
}

func (s *DeployOllamaModelOutputData) Validate() error {
	return dara.Validate(s)
}
