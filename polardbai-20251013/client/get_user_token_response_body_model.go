// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *GetUserTokenResponseBody
	GetData() interface{}
	SetErrCode(v string) *GetUserTokenResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetUserTokenResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetUserTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserTokenResponseBody
	GetSuccess() *bool
}

type GetUserTokenResponseBody struct {
	// example:
	//
	// {"token": "xxxxxx", "expire": 3600}
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

func (s GetUserTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserTokenResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetUserTokenResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetUserTokenResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetUserTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserTokenResponseBody) SetData(v interface{}) *GetUserTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetUserTokenResponseBody) SetErrCode(v string) *GetUserTokenResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetUserTokenResponseBody) SetErrMessage(v string) *GetUserTokenResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetUserTokenResponseBody) SetRequestId(v string) *GetUserTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserTokenResponseBody) SetSuccess(v bool) *GetUserTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
