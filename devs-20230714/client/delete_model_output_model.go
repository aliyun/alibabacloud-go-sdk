// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelOutput interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DeleteModelOutput
	GetErrCode() *string
	SetErrMsg(v string) *DeleteModelOutput
	GetErrMsg() *string
	SetRequestId(v string) *DeleteModelOutput
	GetRequestId() *string
	SetSuccess(v bool) *DeleteModelOutput
	GetSuccess() *bool
}

type DeleteModelOutput struct {
	ErrCode   *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg    *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteModelOutput) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelOutput) GoString() string {
	return s.String()
}

func (s *DeleteModelOutput) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteModelOutput) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DeleteModelOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelOutput) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteModelOutput) SetErrCode(v string) *DeleteModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DeleteModelOutput) SetErrMsg(v string) *DeleteModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeleteModelOutput) SetRequestId(v string) *DeleteModelOutput {
	s.RequestId = &v
	return s
}

func (s *DeleteModelOutput) SetSuccess(v bool) *DeleteModelOutput {
	s.Success = &v
	return s
}

func (s *DeleteModelOutput) Validate() error {
	return dara.Validate(s)
}
