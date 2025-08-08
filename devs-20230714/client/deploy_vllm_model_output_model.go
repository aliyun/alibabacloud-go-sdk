// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployVllmModelOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeployVllmModelOutputData) *DeployVllmModelOutput
	GetData() *DeployVllmModelOutputData
	SetErrCode(v string) *DeployVllmModelOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeployVllmModelOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeployVllmModelOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeployVllmModelOutput
	GetSuccess() *bool
}

type DeployVllmModelOutput struct {
	Data    *DeployVllmModelOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                    `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                    `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployVllmModelOutput) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelOutput) GoString() string {
	return s.String()
}

func (s *DeployVllmModelOutput) GetData() *DeployVllmModelOutputData {
	return s.Data
}

func (s *DeployVllmModelOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeployVllmModelOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeployVllmModelOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployVllmModelOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeployVllmModelOutput) SetData(v *DeployVllmModelOutputData) *DeployVllmModelOutput {
	s.Data = v
	return s
}

func (s *DeployVllmModelOutput) SetErrCode(v string) *DeployVllmModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployVllmModelOutput) SetErrMsg(v string) *DeployVllmModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployVllmModelOutput) SetRequestId(v string) *DeployVllmModelOutput {
	s.RequestId = &v
	return s
}

func (s *DeployVllmModelOutput) SetSuccess(v bool) *DeployVllmModelOutput {
	s.Success = &v
	return s
}

func (s *DeployVllmModelOutput) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelOutputData struct {
	DeploymentTaskID *string `json:"deploymentTaskID,omitempty" xml:"deploymentTaskID,omitempty"`
	ErrorMessage     *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Finished         *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	ModelName        *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TraceID          *string `json:"traceID,omitempty" xml:"traceID,omitempty"`
	UrlInternet      *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet      *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
}

func (s DeployVllmModelOutputData) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelOutputData) GoString() string {
	return s.String()
}

func (s *DeployVllmModelOutputData) GetDeploymentTaskID() *string {
	return s.DeploymentTaskID
}

func (s *DeployVllmModelOutputData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeployVllmModelOutputData) GetFinished() *bool {
	return s.Finished
}

func (s *DeployVllmModelOutputData) GetModelName() *string {
	return s.ModelName
}

func (s *DeployVllmModelOutputData) GetServiceName() *string {
	return s.ServiceName
}

func (s *DeployVllmModelOutputData) GetTraceID() *string {
	return s.TraceID
}

func (s *DeployVllmModelOutputData) GetUrlInternet() *string {
	return s.UrlInternet
}

func (s *DeployVllmModelOutputData) GetUrlIntranet() *string {
	return s.UrlIntranet
}

func (s *DeployVllmModelOutputData) SetDeploymentTaskID(v string) *DeployVllmModelOutputData {
	s.DeploymentTaskID = &v
	return s
}

func (s *DeployVllmModelOutputData) SetErrorMessage(v string) *DeployVllmModelOutputData {
	s.ErrorMessage = &v
	return s
}

func (s *DeployVllmModelOutputData) SetFinished(v bool) *DeployVllmModelOutputData {
	s.Finished = &v
	return s
}

func (s *DeployVllmModelOutputData) SetModelName(v string) *DeployVllmModelOutputData {
	s.ModelName = &v
	return s
}

func (s *DeployVllmModelOutputData) SetServiceName(v string) *DeployVllmModelOutputData {
	s.ServiceName = &v
	return s
}

func (s *DeployVllmModelOutputData) SetTraceID(v string) *DeployVllmModelOutputData {
	s.TraceID = &v
	return s
}

func (s *DeployVllmModelOutputData) SetUrlInternet(v string) *DeployVllmModelOutputData {
	s.UrlInternet = &v
	return s
}

func (s *DeployVllmModelOutputData) SetUrlIntranet(v string) *DeployVllmModelOutputData {
	s.UrlIntranet = &v
	return s
}

func (s *DeployVllmModelOutputData) Validate() error {
	return dara.Validate(s)
}
