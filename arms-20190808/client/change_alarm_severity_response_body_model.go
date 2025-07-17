// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAlarmSeverityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ChangeAlarmSeverityResponseBody
	GetCode() *int64
	SetMessage(v string) *ChangeAlarmSeverityResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeAlarmSeverityResponseBody
	GetRequestId() *string
	SetResult(v bool) *ChangeAlarmSeverityResponseBody
	GetResult() *bool
	SetSuccess(v bool) *ChangeAlarmSeverityResponseBody
	GetSuccess() *bool
}

type ChangeAlarmSeverityResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F7781D4A-2818-41E7-B7BB-79D809E9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the severity level was modified.
	//
	// - `true`: The severity level was modified.
	//
	// - `false`: The severity level failed to be modified.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeAlarmSeverityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeAlarmSeverityResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeAlarmSeverityResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ChangeAlarmSeverityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeAlarmSeverityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeAlarmSeverityResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ChangeAlarmSeverityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeAlarmSeverityResponseBody) SetCode(v int64) *ChangeAlarmSeverityResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeAlarmSeverityResponseBody) SetMessage(v string) *ChangeAlarmSeverityResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeAlarmSeverityResponseBody) SetRequestId(v string) *ChangeAlarmSeverityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeAlarmSeverityResponseBody) SetResult(v bool) *ChangeAlarmSeverityResponseBody {
	s.Result = &v
	return s
}

func (s *ChangeAlarmSeverityResponseBody) SetSuccess(v bool) *ChangeAlarmSeverityResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeAlarmSeverityResponseBody) Validate() error {
	return dara.Validate(s)
}
