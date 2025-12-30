// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDashscopeAsyncTaskFinishEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DashscopeAsyncTaskFinishEventResponseBody
	GetCode() *string
	SetMessage(v string) *DashscopeAsyncTaskFinishEventResponseBody
	GetMessage() *string
	SetRetryAble(v bool) *DashscopeAsyncTaskFinishEventResponseBody
	GetRetryAble() *bool
	SetSuccess(v bool) *DashscopeAsyncTaskFinishEventResponseBody
	GetSuccess() *bool
}

type DashscopeAsyncTaskFinishEventResponseBody struct {
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RetryAble *bool   `json:"retryAble,omitempty" xml:"retryAble,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DashscopeAsyncTaskFinishEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DashscopeAsyncTaskFinishEventResponseBody) GoString() string {
	return s.String()
}

func (s *DashscopeAsyncTaskFinishEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *DashscopeAsyncTaskFinishEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DashscopeAsyncTaskFinishEventResponseBody) GetRetryAble() *bool {
	return s.RetryAble
}

func (s *DashscopeAsyncTaskFinishEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DashscopeAsyncTaskFinishEventResponseBody) SetCode(v string) *DashscopeAsyncTaskFinishEventResponseBody {
	s.Code = &v
	return s
}

func (s *DashscopeAsyncTaskFinishEventResponseBody) SetMessage(v string) *DashscopeAsyncTaskFinishEventResponseBody {
	s.Message = &v
	return s
}

func (s *DashscopeAsyncTaskFinishEventResponseBody) SetRetryAble(v bool) *DashscopeAsyncTaskFinishEventResponseBody {
	s.RetryAble = &v
	return s
}

func (s *DashscopeAsyncTaskFinishEventResponseBody) SetSuccess(v bool) *DashscopeAsyncTaskFinishEventResponseBody {
	s.Success = &v
	return s
}

func (s *DashscopeAsyncTaskFinishEventResponseBody) Validate() error {
	return dara.Validate(s)
}
