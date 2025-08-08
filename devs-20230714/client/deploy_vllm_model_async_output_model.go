// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployVllmModelAsyncOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeployVllmModelAsyncOutput
	GetData() *string
	SetErrCode(v string) *DeployVllmModelAsyncOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeployVllmModelAsyncOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeployVllmModelAsyncOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeployVllmModelAsyncOutput
	GetSuccess() *bool
}

type DeployVllmModelAsyncOutput struct {
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployVllmModelAsyncOutput) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelAsyncOutput) GoString() string {
	return s.String()
}

func (s *DeployVllmModelAsyncOutput) GetData() *string {
	return s.Data
}

func (s *DeployVllmModelAsyncOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeployVllmModelAsyncOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeployVllmModelAsyncOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployVllmModelAsyncOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeployVllmModelAsyncOutput) SetData(v string) *DeployVllmModelAsyncOutput {
	s.Data = &v
	return s
}

func (s *DeployVllmModelAsyncOutput) SetErrCode(v string) *DeployVllmModelAsyncOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployVllmModelAsyncOutput) SetErrMsg(v string) *DeployVllmModelAsyncOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployVllmModelAsyncOutput) SetRequestId(v string) *DeployVllmModelAsyncOutput {
	s.RequestId = &v
	return s
}

func (s *DeployVllmModelAsyncOutput) SetSuccess(v bool) *DeployVllmModelAsyncOutput {
	s.Success = &v
	return s
}

func (s *DeployVllmModelAsyncOutput) Validate() error {
	return dara.Validate(s)
}
