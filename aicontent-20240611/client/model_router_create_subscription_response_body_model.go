// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubscriptionDTO) *ModelRouterCreateSubscriptionResponseBody
	GetData() *SubscriptionDTO
	SetErrCode(v string) *ModelRouterCreateSubscriptionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterCreateSubscriptionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterCreateSubscriptionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterCreateSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterCreateSubscriptionResponseBody
	GetSuccess() *bool
}

type ModelRouterCreateSubscriptionResponseBody struct {
	// example:
	//
	// []
	Data *SubscriptionDTO `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// null
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

func (s ModelRouterCreateSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateSubscriptionResponseBody) GetData() *SubscriptionDTO {
	return s.Data
}

func (s *ModelRouterCreateSubscriptionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterCreateSubscriptionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterCreateSubscriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterCreateSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterCreateSubscriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterCreateSubscriptionResponseBody) SetData(v *SubscriptionDTO) *ModelRouterCreateSubscriptionResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterCreateSubscriptionResponseBody) SetErrCode(v string) *ModelRouterCreateSubscriptionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterCreateSubscriptionResponseBody) SetErrMessage(v string) *ModelRouterCreateSubscriptionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterCreateSubscriptionResponseBody) SetHttpStatusCode(v int32) *ModelRouterCreateSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterCreateSubscriptionResponseBody) SetRequestId(v string) *ModelRouterCreateSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterCreateSubscriptionResponseBody) SetSuccess(v bool) *ModelRouterCreateSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterCreateSubscriptionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
