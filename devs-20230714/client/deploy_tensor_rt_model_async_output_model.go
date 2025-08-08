// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployTensorRtModelAsyncOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeployTensorRtModelAsyncOutput
	GetData() *string
	SetErrCode(v string) *DeployTensorRtModelAsyncOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeployTensorRtModelAsyncOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeployTensorRtModelAsyncOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeployTensorRtModelAsyncOutput
	GetSuccess() *bool
}

type DeployTensorRtModelAsyncOutput struct {
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployTensorRtModelAsyncOutput) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelAsyncOutput) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelAsyncOutput) GetData() *string {
	return s.Data
}

func (s *DeployTensorRtModelAsyncOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeployTensorRtModelAsyncOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeployTensorRtModelAsyncOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployTensorRtModelAsyncOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeployTensorRtModelAsyncOutput) SetData(v string) *DeployTensorRtModelAsyncOutput {
	s.Data = &v
	return s
}

func (s *DeployTensorRtModelAsyncOutput) SetErrCode(v string) *DeployTensorRtModelAsyncOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployTensorRtModelAsyncOutput) SetErrMsg(v string) *DeployTensorRtModelAsyncOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployTensorRtModelAsyncOutput) SetRequestId(v string) *DeployTensorRtModelAsyncOutput {
	s.RequestId = &v
	return s
}

func (s *DeployTensorRtModelAsyncOutput) SetSuccess(v bool) *DeployTensorRtModelAsyncOutput {
	s.Success = &v
	return s
}

func (s *DeployTensorRtModelAsyncOutput) Validate() error {
	return dara.Validate(s)
}
