// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateAlarmResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateAlarmResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAlarmResponseBody
	GetRequestId() *string
	SetResult(v int64) *CreateAlarmResponseBody
	GetResult() *int64
}

type CreateAlarmResponseBody struct {
	// Request result status code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// error message
	//
	// example:
	//
	// xx为空
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// request ID
	//
	// example:
	//
	// 43***28C-A810-5***-8747-EC226A086881
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Alarm ID
	//
	// example:
	//
	// 1234567
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlarmResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateAlarmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlarmResponseBody) GetResult() *int64 {
	return s.Result
}

func (s *CreateAlarmResponseBody) SetCode(v int32) *CreateAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAlarmResponseBody) SetMessage(v string) *CreateAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAlarmResponseBody) SetRequestId(v string) *CreateAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlarmResponseBody) SetResult(v int64) *CreateAlarmResponseBody {
	s.Result = &v
	return s
}

func (s *CreateAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
