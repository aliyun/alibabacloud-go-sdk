// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvironmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteEnvironmentResponseBody
	GetCode() *int32
	SetData(v string) *DeleteEnvironmentResponseBody
	GetData() *string
	SetMessage(v string) *DeleteEnvironmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEnvironmentResponseBody
	GetRequestId() *string
}

type DeleteEnvironmentResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34ED024E-9E31-434A-9E4E-D9D15C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnvironmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteEnvironmentResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteEnvironmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEnvironmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnvironmentResponseBody) SetCode(v int32) *DeleteEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEnvironmentResponseBody) SetData(v string) *DeleteEnvironmentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteEnvironmentResponseBody) SetMessage(v string) *DeleteEnvironmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEnvironmentResponseBody) SetRequestId(v string) *DeleteEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnvironmentResponseBody) Validate() error {
	return dara.Validate(s)
}
