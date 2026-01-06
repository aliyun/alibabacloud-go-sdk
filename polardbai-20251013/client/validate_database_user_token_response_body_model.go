// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateDatabaseUserTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ValidateDatabaseUserTokenResponseBody
	GetData() interface{}
	SetErrCode(v string) *ValidateDatabaseUserTokenResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ValidateDatabaseUserTokenResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ValidateDatabaseUserTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ValidateDatabaseUserTokenResponseBody
	GetSuccess() *bool
}

type ValidateDatabaseUserTokenResponseBody struct {
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

func (s ValidateDatabaseUserTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateDatabaseUserTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateDatabaseUserTokenResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ValidateDatabaseUserTokenResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ValidateDatabaseUserTokenResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ValidateDatabaseUserTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateDatabaseUserTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ValidateDatabaseUserTokenResponseBody) SetData(v interface{}) *ValidateDatabaseUserTokenResponseBody {
	s.Data = v
	return s
}

func (s *ValidateDatabaseUserTokenResponseBody) SetErrCode(v string) *ValidateDatabaseUserTokenResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ValidateDatabaseUserTokenResponseBody) SetErrMessage(v string) *ValidateDatabaseUserTokenResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ValidateDatabaseUserTokenResponseBody) SetRequestId(v string) *ValidateDatabaseUserTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateDatabaseUserTokenResponseBody) SetSuccess(v bool) *ValidateDatabaseUserTokenResponseBody {
	s.Success = &v
	return s
}

func (s *ValidateDatabaseUserTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
