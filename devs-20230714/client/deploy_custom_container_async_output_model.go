// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployCustomContainerAsyncOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeployCustomContainerAsyncOutput
	GetData() *string
	SetErrCode(v string) *DeployCustomContainerAsyncOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeployCustomContainerAsyncOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeployCustomContainerAsyncOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeployCustomContainerAsyncOutput
	GetSuccess() *bool
}

type DeployCustomContainerAsyncOutput struct {
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployCustomContainerAsyncOutput) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerAsyncOutput) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerAsyncOutput) GetData() *string {
	return s.Data
}

func (s *DeployCustomContainerAsyncOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeployCustomContainerAsyncOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeployCustomContainerAsyncOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployCustomContainerAsyncOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeployCustomContainerAsyncOutput) SetData(v string) *DeployCustomContainerAsyncOutput {
	s.Data = &v
	return s
}

func (s *DeployCustomContainerAsyncOutput) SetErrCode(v string) *DeployCustomContainerAsyncOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployCustomContainerAsyncOutput) SetErrMsg(v string) *DeployCustomContainerAsyncOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployCustomContainerAsyncOutput) SetRequestId(v string) *DeployCustomContainerAsyncOutput {
	s.RequestId = &v
	return s
}

func (s *DeployCustomContainerAsyncOutput) SetSuccess(v bool) *DeployCustomContainerAsyncOutput {
	s.Success = &v
	return s
}

func (s *DeployCustomContainerAsyncOutput) Validate() error {
	return dara.Validate(s)
}
