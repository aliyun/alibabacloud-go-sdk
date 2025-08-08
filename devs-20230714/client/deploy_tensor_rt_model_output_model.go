// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployTensorRtModelOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeployTensorRtModelOutputData) *DeployTensorRtModelOutput
	GetData() *DeployTensorRtModelOutputData
	SetErrCode(v string) *DeployTensorRtModelOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeployTensorRtModelOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeployTensorRtModelOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeployTensorRtModelOutput
	GetSuccess() *bool
}

type DeployTensorRtModelOutput struct {
	Data    *DeployTensorRtModelOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                        `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                        `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployTensorRtModelOutput) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelOutput) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelOutput) GetData() *DeployTensorRtModelOutputData {
	return s.Data
}

func (s *DeployTensorRtModelOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeployTensorRtModelOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeployTensorRtModelOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployTensorRtModelOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeployTensorRtModelOutput) SetData(v *DeployTensorRtModelOutputData) *DeployTensorRtModelOutput {
	s.Data = v
	return s
}

func (s *DeployTensorRtModelOutput) SetErrCode(v string) *DeployTensorRtModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployTensorRtModelOutput) SetErrMsg(v string) *DeployTensorRtModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployTensorRtModelOutput) SetRequestId(v string) *DeployTensorRtModelOutput {
	s.RequestId = &v
	return s
}

func (s *DeployTensorRtModelOutput) SetSuccess(v bool) *DeployTensorRtModelOutput {
	s.Success = &v
	return s
}

func (s *DeployTensorRtModelOutput) Validate() error {
	return dara.Validate(s)
}

type DeployTensorRtModelOutputData struct {
	DeploymentTaskID *string `json:"deploymentTaskID,omitempty" xml:"deploymentTaskID,omitempty"`
	ErrorMessage     *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Finished         *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TraceID          *string `json:"traceID,omitempty" xml:"traceID,omitempty"`
	UrlInternet      *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet      *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
}

func (s DeployTensorRtModelOutputData) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelOutputData) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelOutputData) GetDeploymentTaskID() *string {
	return s.DeploymentTaskID
}

func (s *DeployTensorRtModelOutputData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeployTensorRtModelOutputData) GetFinished() *bool {
	return s.Finished
}

func (s *DeployTensorRtModelOutputData) GetServiceName() *string {
	return s.ServiceName
}

func (s *DeployTensorRtModelOutputData) GetTraceID() *string {
	return s.TraceID
}

func (s *DeployTensorRtModelOutputData) GetUrlInternet() *string {
	return s.UrlInternet
}

func (s *DeployTensorRtModelOutputData) GetUrlIntranet() *string {
	return s.UrlIntranet
}

func (s *DeployTensorRtModelOutputData) SetDeploymentTaskID(v string) *DeployTensorRtModelOutputData {
	s.DeploymentTaskID = &v
	return s
}

func (s *DeployTensorRtModelOutputData) SetErrorMessage(v string) *DeployTensorRtModelOutputData {
	s.ErrorMessage = &v
	return s
}

func (s *DeployTensorRtModelOutputData) SetFinished(v bool) *DeployTensorRtModelOutputData {
	s.Finished = &v
	return s
}

func (s *DeployTensorRtModelOutputData) SetServiceName(v string) *DeployTensorRtModelOutputData {
	s.ServiceName = &v
	return s
}

func (s *DeployTensorRtModelOutputData) SetTraceID(v string) *DeployTensorRtModelOutputData {
	s.TraceID = &v
	return s
}

func (s *DeployTensorRtModelOutputData) SetUrlInternet(v string) *DeployTensorRtModelOutputData {
	s.UrlInternet = &v
	return s
}

func (s *DeployTensorRtModelOutputData) SetUrlIntranet(v string) *DeployTensorRtModelOutputData {
	s.UrlIntranet = &v
	return s
}

func (s *DeployTensorRtModelOutputData) Validate() error {
	return dara.Validate(s)
}
