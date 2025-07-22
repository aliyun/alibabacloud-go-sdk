// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceShutdownTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceShutdownTimerResponseBody
	GetCode() *string
	SetDueTime(v string) *GetInstanceShutdownTimerResponseBody
	GetDueTime() *string
	SetGmtCreateTime(v string) *GetInstanceShutdownTimerResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetInstanceShutdownTimerResponseBody
	GetGmtModifiedTime() *string
	SetHttpStatusCode(v int32) *GetInstanceShutdownTimerResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *GetInstanceShutdownTimerResponseBody
	GetInstanceId() *string
	SetMessage(v string) *GetInstanceShutdownTimerResponseBody
	GetMessage() *string
	SetRemainingTimeInMs(v int64) *GetInstanceShutdownTimerResponseBody
	GetRemainingTimeInMs() *int64
	SetRequestId(v string) *GetInstanceShutdownTimerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceShutdownTimerResponseBody
	GetSuccess() *bool
}

type GetInstanceShutdownTimerResponseBody struct {
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// null
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3600000
	RemainingTimeInMs *int64 `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceShutdownTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceShutdownTimerResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceShutdownTimerResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceShutdownTimerResponseBody) GetDueTime() *string {
	return s.DueTime
}

func (s *GetInstanceShutdownTimerResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetInstanceShutdownTimerResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetInstanceShutdownTimerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceShutdownTimerResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceShutdownTimerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceShutdownTimerResponseBody) GetRemainingTimeInMs() *int64 {
	return s.RemainingTimeInMs
}

func (s *GetInstanceShutdownTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceShutdownTimerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceShutdownTimerResponseBody) SetCode(v string) *GetInstanceShutdownTimerResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetDueTime(v string) *GetInstanceShutdownTimerResponseBody {
	s.DueTime = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetGmtCreateTime(v string) *GetInstanceShutdownTimerResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetGmtModifiedTime(v string) *GetInstanceShutdownTimerResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetHttpStatusCode(v int32) *GetInstanceShutdownTimerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetInstanceId(v string) *GetInstanceShutdownTimerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetMessage(v string) *GetInstanceShutdownTimerResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetRemainingTimeInMs(v int64) *GetInstanceShutdownTimerResponseBody {
	s.RemainingTimeInMs = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetRequestId(v string) *GetInstanceShutdownTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetSuccess(v bool) *GetInstanceShutdownTimerResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) Validate() error {
	return dara.Validate(s)
}
