// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CloseAlarmResponseBody
	GetCode() *int64
	SetMessage(v string) *CloseAlarmResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloseAlarmResponseBody
	GetRequestId() *string
	SetResult(v bool) *CloseAlarmResponseBody
	GetResult() *bool
	SetSuccess(v bool) *CloseAlarmResponseBody
	GetSuccess() *bool
}

type CloseAlarmResponseBody struct {
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
	// 46355DD8-FC56-40C5-BFC6-269DE4F9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CloseAlarmResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CloseAlarmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloseAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseAlarmResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CloseAlarmResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CloseAlarmResponseBody) SetCode(v int64) *CloseAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *CloseAlarmResponseBody) SetMessage(v string) *CloseAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *CloseAlarmResponseBody) SetRequestId(v string) *CloseAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseAlarmResponseBody) SetResult(v bool) *CloseAlarmResponseBody {
	s.Result = &v
	return s
}

func (s *CloseAlarmResponseBody) SetSuccess(v bool) *CloseAlarmResponseBody {
	s.Success = &v
	return s
}

func (s *CloseAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
