// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployCustomContainerOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeployCustomContainerOutputData) *DeployCustomContainerOutput
	GetData() *DeployCustomContainerOutputData
	SetErrCode(v string) *DeployCustomContainerOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeployCustomContainerOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeployCustomContainerOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeployCustomContainerOutput
	GetSuccess() *bool
}

type DeployCustomContainerOutput struct {
	Data    *DeployCustomContainerOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                          `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                          `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployCustomContainerOutput) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerOutput) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerOutput) GetData() *DeployCustomContainerOutputData {
	return s.Data
}

func (s *DeployCustomContainerOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeployCustomContainerOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeployCustomContainerOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployCustomContainerOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeployCustomContainerOutput) SetData(v *DeployCustomContainerOutputData) *DeployCustomContainerOutput {
	s.Data = v
	return s
}

func (s *DeployCustomContainerOutput) SetErrCode(v string) *DeployCustomContainerOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployCustomContainerOutput) SetErrMsg(v string) *DeployCustomContainerOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployCustomContainerOutput) SetRequestId(v string) *DeployCustomContainerOutput {
	s.RequestId = &v
	return s
}

func (s *DeployCustomContainerOutput) SetSuccess(v bool) *DeployCustomContainerOutput {
	s.Success = &v
	return s
}

func (s *DeployCustomContainerOutput) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerOutputData struct {
	DeploymentTaskID *string `json:"deploymentTaskID,omitempty" xml:"deploymentTaskID,omitempty"`
	ErrorMessage     *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Finished         *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	NasConfigStr     *string `json:"nasConfigStr,omitempty" xml:"nasConfigStr,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TraceID          *string `json:"traceID,omitempty" xml:"traceID,omitempty"`
	UrlInternet      *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet      *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
	VpcConfigStr     *string `json:"vpcConfigStr,omitempty" xml:"vpcConfigStr,omitempty"`
}

func (s DeployCustomContainerOutputData) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerOutputData) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerOutputData) GetDeploymentTaskID() *string {
	return s.DeploymentTaskID
}

func (s *DeployCustomContainerOutputData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeployCustomContainerOutputData) GetFinished() *bool {
	return s.Finished
}

func (s *DeployCustomContainerOutputData) GetNasConfigStr() *string {
	return s.NasConfigStr
}

func (s *DeployCustomContainerOutputData) GetServiceName() *string {
	return s.ServiceName
}

func (s *DeployCustomContainerOutputData) GetTraceID() *string {
	return s.TraceID
}

func (s *DeployCustomContainerOutputData) GetUrlInternet() *string {
	return s.UrlInternet
}

func (s *DeployCustomContainerOutputData) GetUrlIntranet() *string {
	return s.UrlIntranet
}

func (s *DeployCustomContainerOutputData) GetVpcConfigStr() *string {
	return s.VpcConfigStr
}

func (s *DeployCustomContainerOutputData) SetDeploymentTaskID(v string) *DeployCustomContainerOutputData {
	s.DeploymentTaskID = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetErrorMessage(v string) *DeployCustomContainerOutputData {
	s.ErrorMessage = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetFinished(v bool) *DeployCustomContainerOutputData {
	s.Finished = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetNasConfigStr(v string) *DeployCustomContainerOutputData {
	s.NasConfigStr = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetServiceName(v string) *DeployCustomContainerOutputData {
	s.ServiceName = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetTraceID(v string) *DeployCustomContainerOutputData {
	s.TraceID = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetUrlInternet(v string) *DeployCustomContainerOutputData {
	s.UrlInternet = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetUrlIntranet(v string) *DeployCustomContainerOutputData {
	s.UrlIntranet = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetVpcConfigStr(v string) *DeployCustomContainerOutputData {
	s.VpcConfigStr = &v
	return s
}

func (s *DeployCustomContainerOutputData) Validate() error {
	return dara.Validate(s)
}
