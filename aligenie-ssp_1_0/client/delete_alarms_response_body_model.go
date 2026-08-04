// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlarmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteAlarmsResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteAlarmsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAlarmsResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteAlarmsResponseBody
	GetResult() *bool
}

type DeleteAlarmsResponseBody struct {
	// Status code returned by the alarm service
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// error message
	//
	// example:
	//
	// 闹钟id为空
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// request ID
	//
	// example:
	//
	// 43***28C-A810-5***-8747-EC226A086881
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the alarm deletion was executed successfully
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteAlarmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteAlarmsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAlarmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlarmsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteAlarmsResponseBody) SetCode(v int32) *DeleteAlarmsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAlarmsResponseBody) SetMessage(v string) *DeleteAlarmsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAlarmsResponseBody) SetRequestId(v string) *DeleteAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlarmsResponseBody) SetResult(v bool) *DeleteAlarmsResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteAlarmsResponseBody) Validate() error {
	return dara.Validate(s)
}
