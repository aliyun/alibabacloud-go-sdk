// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DeleteConsumerGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteConsumerGroupResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DeleteConsumerGroupResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteConsumerGroupResponseBody
	GetSuccess() *string
}

type DeleteConsumerGroupResponseBody struct {
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
	// The ID of the request.
	//
	// example:
	//
	// 4D0ADAD5-DD97-41B6-B78F-D1961AB1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteConsumerGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConsumerGroupResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteConsumerGroupResponseBody) SetErrCode(v string) *DeleteConsumerGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetErrMessage(v string) *DeleteConsumerGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetRequestId(v string) *DeleteConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetSuccess(v string) *DeleteConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
