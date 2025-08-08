// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployOllamaModelAsyncOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeployOllamaModelAsyncOutput
	GetData() *string
	SetErrCode(v string) *DeployOllamaModelAsyncOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeployOllamaModelAsyncOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeployOllamaModelAsyncOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeployOllamaModelAsyncOutput
	GetSuccess() *bool
}

type DeployOllamaModelAsyncOutput struct {
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployOllamaModelAsyncOutput) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelAsyncOutput) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelAsyncOutput) GetData() *string {
	return s.Data
}

func (s *DeployOllamaModelAsyncOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeployOllamaModelAsyncOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeployOllamaModelAsyncOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployOllamaModelAsyncOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeployOllamaModelAsyncOutput) SetData(v string) *DeployOllamaModelAsyncOutput {
	s.Data = &v
	return s
}

func (s *DeployOllamaModelAsyncOutput) SetErrCode(v string) *DeployOllamaModelAsyncOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployOllamaModelAsyncOutput) SetErrMsg(v string) *DeployOllamaModelAsyncOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployOllamaModelAsyncOutput) SetRequestId(v string) *DeployOllamaModelAsyncOutput {
	s.RequestId = &v
	return s
}

func (s *DeployOllamaModelAsyncOutput) SetSuccess(v bool) *DeployOllamaModelAsyncOutput {
	s.Success = &v
	return s
}

func (s *DeployOllamaModelAsyncOutput) Validate() error {
	return dara.Validate(s)
}
