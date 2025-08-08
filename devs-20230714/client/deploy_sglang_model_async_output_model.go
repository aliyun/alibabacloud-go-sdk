// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploySGLangModelAsyncOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeploySGLangModelAsyncOutput
	GetData() *string
	SetErrCode(v string) *DeploySGLangModelAsyncOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeploySGLangModelAsyncOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeploySGLangModelAsyncOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeploySGLangModelAsyncOutput
	GetSuccess() *bool
}

type DeploySGLangModelAsyncOutput struct {
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeploySGLangModelAsyncOutput) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelAsyncOutput) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelAsyncOutput) GetData() *string {
	return s.Data
}

func (s *DeploySGLangModelAsyncOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeploySGLangModelAsyncOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeploySGLangModelAsyncOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeploySGLangModelAsyncOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeploySGLangModelAsyncOutput) SetData(v string) *DeploySGLangModelAsyncOutput {
	s.Data = &v
	return s
}

func (s *DeploySGLangModelAsyncOutput) SetErrCode(v string) *DeploySGLangModelAsyncOutput {
	s.ErrCode = &v
	return s
}

func (s *DeploySGLangModelAsyncOutput) SetErrMsg(v string) *DeploySGLangModelAsyncOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeploySGLangModelAsyncOutput) SetRequestId(v string) *DeploySGLangModelAsyncOutput {
	s.RequestId = &v
	return s
}

func (s *DeploySGLangModelAsyncOutput) SetSuccess(v bool) *DeploySGLangModelAsyncOutput {
	s.Success = &v
	return s
}

func (s *DeploySGLangModelAsyncOutput) Validate() error {
	return dara.Validate(s)
}
