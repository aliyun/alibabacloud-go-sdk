// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployModelScopeModelAsyncOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeployModelScopeModelAsyncOutput
	GetData() *string
	SetErrCode(v string) *DeployModelScopeModelAsyncOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeployModelScopeModelAsyncOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeployModelScopeModelAsyncOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeployModelScopeModelAsyncOutput
	GetSuccess() *bool
}

type DeployModelScopeModelAsyncOutput struct {
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployModelScopeModelAsyncOutput) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelAsyncOutput) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelAsyncOutput) GetData() *string {
	return s.Data
}

func (s *DeployModelScopeModelAsyncOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeployModelScopeModelAsyncOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeployModelScopeModelAsyncOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployModelScopeModelAsyncOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeployModelScopeModelAsyncOutput) SetData(v string) *DeployModelScopeModelAsyncOutput {
	s.Data = &v
	return s
}

func (s *DeployModelScopeModelAsyncOutput) SetErrCode(v string) *DeployModelScopeModelAsyncOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployModelScopeModelAsyncOutput) SetErrMsg(v string) *DeployModelScopeModelAsyncOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployModelScopeModelAsyncOutput) SetRequestId(v string) *DeployModelScopeModelAsyncOutput {
	s.RequestId = &v
	return s
}

func (s *DeployModelScopeModelAsyncOutput) SetSuccess(v bool) *DeployModelScopeModelAsyncOutput {
	s.Success = &v
	return s
}

func (s *DeployModelScopeModelAsyncOutput) Validate() error {
	return dara.Validate(s)
}
