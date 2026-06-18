// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAiOutboundTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAiOutboundTaskResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteAiOutboundTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAiOutboundTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAiOutboundTaskResponseBody
	GetSuccess() *bool
}

type DeleteAiOutboundTaskResponseBody struct {
	// The request status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAiOutboundTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAiOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAiOutboundTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAiOutboundTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAiOutboundTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAiOutboundTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAiOutboundTaskResponseBody) SetCode(v string) *DeleteAiOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAiOutboundTaskResponseBody) SetMessage(v string) *DeleteAiOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAiOutboundTaskResponseBody) SetRequestId(v string) *DeleteAiOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAiOutboundTaskResponseBody) SetSuccess(v bool) *DeleteAiOutboundTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAiOutboundTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
