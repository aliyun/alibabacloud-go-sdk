// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMemoryResponseBody
	GetCode() *string
	SetRequestId(v string) *UpdateMemoryResponseBody
	GetRequestId() *string
}

type UpdateMemoryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C0595DB0-D1EE-55C3-8DDD-790872C7EC2F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemoryResponseBody) SetCode(v string) *UpdateMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMemoryResponseBody) SetRequestId(v string) *UpdateMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
