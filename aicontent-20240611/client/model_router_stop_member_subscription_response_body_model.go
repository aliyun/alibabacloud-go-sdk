// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterStopMemberSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterStopMemberSubscriptionResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterStopMemberSubscriptionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterStopMemberSubscriptionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterStopMemberSubscriptionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterStopMemberSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterStopMemberSubscriptionResponseBody
	GetSuccess() *bool
}

type ModelRouterStopMemberSubscriptionResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterStopMemberSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterStopMemberSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterStopMemberSubscriptionResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterStopMemberSubscriptionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterStopMemberSubscriptionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterStopMemberSubscriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterStopMemberSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterStopMemberSubscriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterStopMemberSubscriptionResponseBody) SetData(v bool) *ModelRouterStopMemberSubscriptionResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterStopMemberSubscriptionResponseBody) SetErrCode(v string) *ModelRouterStopMemberSubscriptionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterStopMemberSubscriptionResponseBody) SetErrMessage(v string) *ModelRouterStopMemberSubscriptionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterStopMemberSubscriptionResponseBody) SetHttpStatusCode(v int32) *ModelRouterStopMemberSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterStopMemberSubscriptionResponseBody) SetRequestId(v string) *ModelRouterStopMemberSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterStopMemberSubscriptionResponseBody) SetSuccess(v bool) *ModelRouterStopMemberSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterStopMemberSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
