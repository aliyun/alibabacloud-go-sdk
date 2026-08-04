// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateMemberSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubscriptionDTO) *ModelRouterCreateMemberSubscriptionResponseBody
	GetData() *SubscriptionDTO
	SetErrCode(v string) *ModelRouterCreateMemberSubscriptionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterCreateMemberSubscriptionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterCreateMemberSubscriptionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterCreateMemberSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterCreateMemberSubscriptionResponseBody
	GetSuccess() *bool
}

type ModelRouterCreateMemberSubscriptionResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {}
	Data *SubscriptionDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The fault code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterCreateMemberSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateMemberSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateMemberSubscriptionResponseBody) GetData() *SubscriptionDTO {
	return s.Data
}

func (s *ModelRouterCreateMemberSubscriptionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterCreateMemberSubscriptionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterCreateMemberSubscriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterCreateMemberSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterCreateMemberSubscriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterCreateMemberSubscriptionResponseBody) SetData(v *SubscriptionDTO) *ModelRouterCreateMemberSubscriptionResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterCreateMemberSubscriptionResponseBody) SetErrCode(v string) *ModelRouterCreateMemberSubscriptionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterCreateMemberSubscriptionResponseBody) SetErrMessage(v string) *ModelRouterCreateMemberSubscriptionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterCreateMemberSubscriptionResponseBody) SetHttpStatusCode(v int32) *ModelRouterCreateMemberSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterCreateMemberSubscriptionResponseBody) SetRequestId(v string) *ModelRouterCreateMemberSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterCreateMemberSubscriptionResponseBody) SetSuccess(v bool) *ModelRouterCreateMemberSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterCreateMemberSubscriptionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
