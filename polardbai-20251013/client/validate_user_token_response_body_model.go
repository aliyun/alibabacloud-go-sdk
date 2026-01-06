// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateUserTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ValidateUserTokenResponseBody
	GetData() interface{}
	SetErrCode(v string) *ValidateUserTokenResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ValidateUserTokenResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ValidateUserTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ValidateUserTokenResponseBody
	GetSuccess() *bool
}

type ValidateUserTokenResponseBody struct {
	// example:
	//
	// {"valid": true, "token": "xxxxxx", "expire": 3600}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 983863E2-****-1D15-A4AE-3468CD75383D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ValidateUserTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateUserTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateUserTokenResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ValidateUserTokenResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ValidateUserTokenResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ValidateUserTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateUserTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ValidateUserTokenResponseBody) SetData(v interface{}) *ValidateUserTokenResponseBody {
	s.Data = v
	return s
}

func (s *ValidateUserTokenResponseBody) SetErrCode(v string) *ValidateUserTokenResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ValidateUserTokenResponseBody) SetErrMessage(v string) *ValidateUserTokenResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ValidateUserTokenResponseBody) SetRequestId(v string) *ValidateUserTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateUserTokenResponseBody) SetSuccess(v bool) *ValidateUserTokenResponseBody {
	s.Success = &v
	return s
}

func (s *ValidateUserTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
