// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterStopSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubscriptionDTO) *ModelRouterStopSubscriptionResponseBody
	GetData() *SubscriptionDTO
	SetErrCode(v string) *ModelRouterStopSubscriptionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterStopSubscriptionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterStopSubscriptionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterStopSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterStopSubscriptionResponseBody
	GetSuccess() *bool
}

type ModelRouterStopSubscriptionResponseBody struct {
	Data *SubscriptionDTO `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterStopSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterStopSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterStopSubscriptionResponseBody) GetData() *SubscriptionDTO {
	return s.Data
}

func (s *ModelRouterStopSubscriptionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterStopSubscriptionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterStopSubscriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterStopSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterStopSubscriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterStopSubscriptionResponseBody) SetData(v *SubscriptionDTO) *ModelRouterStopSubscriptionResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterStopSubscriptionResponseBody) SetErrCode(v string) *ModelRouterStopSubscriptionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterStopSubscriptionResponseBody) SetErrMessage(v string) *ModelRouterStopSubscriptionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterStopSubscriptionResponseBody) SetHttpStatusCode(v int32) *ModelRouterStopSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterStopSubscriptionResponseBody) SetRequestId(v string) *ModelRouterStopSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterStopSubscriptionResponseBody) SetSuccess(v bool) *ModelRouterStopSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterStopSubscriptionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
