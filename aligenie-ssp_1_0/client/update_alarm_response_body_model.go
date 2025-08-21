// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateAlarmResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateAlarmResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAlarmResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateAlarmResponseBody
	GetResult() *bool
}

type UpdateAlarmResponseBody struct {
	// example:
	//
	// 200
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43***28C-A810-5***-8747-EC226A086881
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlarmResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateAlarmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlarmResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateAlarmResponseBody) SetCode(v int32) *UpdateAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAlarmResponseBody) SetMessage(v string) *UpdateAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAlarmResponseBody) SetRequestId(v string) *UpdateAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlarmResponseBody) SetResult(v bool) *UpdateAlarmResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
