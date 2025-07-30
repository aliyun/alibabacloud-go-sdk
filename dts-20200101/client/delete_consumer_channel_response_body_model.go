// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DeleteConsumerChannelResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteConsumerChannelResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *DeleteConsumerChannelResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *DeleteConsumerChannelResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteConsumerChannelResponseBody
	GetSuccess() *string
}

type DeleteConsumerChannelResponseBody struct {
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
	// F9E00ABE-2AD9-40A9-8C3C-D817E648****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteConsumerChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerChannelResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteConsumerChannelResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteConsumerChannelResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DeleteConsumerChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConsumerChannelResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteConsumerChannelResponseBody) SetErrCode(v string) *DeleteConsumerChannelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteConsumerChannelResponseBody) SetErrMessage(v string) *DeleteConsumerChannelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteConsumerChannelResponseBody) SetHttpStatusCode(v string) *DeleteConsumerChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteConsumerChannelResponseBody) SetRequestId(v string) *DeleteConsumerChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerChannelResponseBody) SetSuccess(v string) *DeleteConsumerChannelResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteConsumerChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
