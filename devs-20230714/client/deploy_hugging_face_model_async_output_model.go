// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployHuggingFaceModelAsyncOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeployHuggingFaceModelAsyncOutput
	GetData() *string
	SetErrCode(v string) *DeployHuggingFaceModelAsyncOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeployHuggingFaceModelAsyncOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeployHuggingFaceModelAsyncOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeployHuggingFaceModelAsyncOutput
	GetSuccess() *bool
}

type DeployHuggingFaceModelAsyncOutput struct {
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployHuggingFaceModelAsyncOutput) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelAsyncOutput) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelAsyncOutput) GetData() *string {
	return s.Data
}

func (s *DeployHuggingFaceModelAsyncOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeployHuggingFaceModelAsyncOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeployHuggingFaceModelAsyncOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployHuggingFaceModelAsyncOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeployHuggingFaceModelAsyncOutput) SetData(v string) *DeployHuggingFaceModelAsyncOutput {
	s.Data = &v
	return s
}

func (s *DeployHuggingFaceModelAsyncOutput) SetErrCode(v string) *DeployHuggingFaceModelAsyncOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployHuggingFaceModelAsyncOutput) SetErrMsg(v string) *DeployHuggingFaceModelAsyncOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployHuggingFaceModelAsyncOutput) SetRequestId(v string) *DeployHuggingFaceModelAsyncOutput {
	s.RequestId = &v
	return s
}

func (s *DeployHuggingFaceModelAsyncOutput) SetSuccess(v bool) *DeployHuggingFaceModelAsyncOutput {
	s.Success = &v
	return s
}

func (s *DeployHuggingFaceModelAsyncOutput) Validate() error {
	return dara.Validate(s)
}
