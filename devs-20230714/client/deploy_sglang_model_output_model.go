// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploySGLangModelOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeploySGLangModelOutputData) *DeploySGLangModelOutput
	GetData() *DeploySGLangModelOutputData
	SetErrCode(v string) *DeploySGLangModelOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeploySGLangModelOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeploySGLangModelOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeploySGLangModelOutput
	GetSuccess() *bool
}

type DeploySGLangModelOutput struct {
	Data    *DeploySGLangModelOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                      `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                      `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeploySGLangModelOutput) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelOutput) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelOutput) GetData() *DeploySGLangModelOutputData {
	return s.Data
}

func (s *DeploySGLangModelOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeploySGLangModelOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeploySGLangModelOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeploySGLangModelOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeploySGLangModelOutput) SetData(v *DeploySGLangModelOutputData) *DeploySGLangModelOutput {
	s.Data = v
	return s
}

func (s *DeploySGLangModelOutput) SetErrCode(v string) *DeploySGLangModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DeploySGLangModelOutput) SetErrMsg(v string) *DeploySGLangModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeploySGLangModelOutput) SetRequestId(v string) *DeploySGLangModelOutput {
	s.RequestId = &v
	return s
}

func (s *DeploySGLangModelOutput) SetSuccess(v bool) *DeploySGLangModelOutput {
	s.Success = &v
	return s
}

func (s *DeploySGLangModelOutput) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelOutputData struct {
	DeploymentTaskID *string `json:"deploymentTaskID,omitempty" xml:"deploymentTaskID,omitempty"`
	ErrorMessage     *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Finished         *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	ModelName        *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TraceID          *string `json:"traceID,omitempty" xml:"traceID,omitempty"`
	UrlInternet      *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet      *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
}

func (s DeploySGLangModelOutputData) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelOutputData) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelOutputData) GetDeploymentTaskID() *string {
	return s.DeploymentTaskID
}

func (s *DeploySGLangModelOutputData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeploySGLangModelOutputData) GetFinished() *bool {
	return s.Finished
}

func (s *DeploySGLangModelOutputData) GetModelName() *string {
	return s.ModelName
}

func (s *DeploySGLangModelOutputData) GetServiceName() *string {
	return s.ServiceName
}

func (s *DeploySGLangModelOutputData) GetTraceID() *string {
	return s.TraceID
}

func (s *DeploySGLangModelOutputData) GetUrlInternet() *string {
	return s.UrlInternet
}

func (s *DeploySGLangModelOutputData) GetUrlIntranet() *string {
	return s.UrlIntranet
}

func (s *DeploySGLangModelOutputData) SetDeploymentTaskID(v string) *DeploySGLangModelOutputData {
	s.DeploymentTaskID = &v
	return s
}

func (s *DeploySGLangModelOutputData) SetErrorMessage(v string) *DeploySGLangModelOutputData {
	s.ErrorMessage = &v
	return s
}

func (s *DeploySGLangModelOutputData) SetFinished(v bool) *DeploySGLangModelOutputData {
	s.Finished = &v
	return s
}

func (s *DeploySGLangModelOutputData) SetModelName(v string) *DeploySGLangModelOutputData {
	s.ModelName = &v
	return s
}

func (s *DeploySGLangModelOutputData) SetServiceName(v string) *DeploySGLangModelOutputData {
	s.ServiceName = &v
	return s
}

func (s *DeploySGLangModelOutputData) SetTraceID(v string) *DeploySGLangModelOutputData {
	s.TraceID = &v
	return s
}

func (s *DeploySGLangModelOutputData) SetUrlInternet(v string) *DeploySGLangModelOutputData {
	s.UrlInternet = &v
	return s
}

func (s *DeploySGLangModelOutputData) SetUrlIntranet(v string) *DeploySGLangModelOutputData {
	s.UrlIntranet = &v
	return s
}

func (s *DeploySGLangModelOutputData) Validate() error {
	return dara.Validate(s)
}
