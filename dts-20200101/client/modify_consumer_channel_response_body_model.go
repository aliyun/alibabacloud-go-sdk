// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumerChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ModifyConsumerChannelResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyConsumerChannelResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *ModifyConsumerChannelResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *ModifyConsumerChannelResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyConsumerChannelResponseBody
	GetSuccess() *string
}

type ModifyConsumerChannelResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 055CAF9B-C15D-4799-BB9E-E62D417****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyConsumerChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumerChannelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyConsumerChannelResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyConsumerChannelResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyConsumerChannelResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ModifyConsumerChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyConsumerChannelResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyConsumerChannelResponseBody) SetErrCode(v string) *ModifyConsumerChannelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyConsumerChannelResponseBody) SetErrMessage(v string) *ModifyConsumerChannelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyConsumerChannelResponseBody) SetHttpStatusCode(v string) *ModifyConsumerChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyConsumerChannelResponseBody) SetRequestId(v string) *ModifyConsumerChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyConsumerChannelResponseBody) SetSuccess(v string) *ModifyConsumerChannelResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyConsumerChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
