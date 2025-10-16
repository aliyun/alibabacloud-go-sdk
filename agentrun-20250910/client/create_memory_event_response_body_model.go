// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMemoryEventResponseBody
	GetCode() *string
	SetRequestId(v string) *CreateMemoryEventResponseBody
	GetRequestId() *string
}

type CreateMemoryEventResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A8B33FA2-43F2-5E56-9032-51283F08018E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMemoryEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryEventResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMemoryEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemoryEventResponseBody) SetCode(v string) *CreateMemoryEventResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMemoryEventResponseBody) SetRequestId(v string) *CreateMemoryEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemoryEventResponseBody) Validate() error {
	return dara.Validate(s)
}
