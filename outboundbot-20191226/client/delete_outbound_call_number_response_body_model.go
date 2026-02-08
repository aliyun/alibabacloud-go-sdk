// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOutboundCallNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteOutboundCallNumberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteOutboundCallNumberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteOutboundCallNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteOutboundCallNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteOutboundCallNumberResponseBody
	GetSuccess() *bool
}

type DeleteOutboundCallNumberResponseBody struct {
	// The status code returned by the API.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the query succeeded:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteOutboundCallNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOutboundCallNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOutboundCallNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteOutboundCallNumberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteOutboundCallNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteOutboundCallNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOutboundCallNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteOutboundCallNumberResponseBody) SetCode(v string) *DeleteOutboundCallNumberResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteOutboundCallNumberResponseBody) SetHttpStatusCode(v int32) *DeleteOutboundCallNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteOutboundCallNumberResponseBody) SetMessage(v string) *DeleteOutboundCallNumberResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteOutboundCallNumberResponseBody) SetRequestId(v string) *DeleteOutboundCallNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOutboundCallNumberResponseBody) SetSuccess(v bool) *DeleteOutboundCallNumberResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteOutboundCallNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
