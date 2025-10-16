// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMemoryResponseBody
	GetCode() *string
	SetRequestId(v string) *CreateMemoryResponseBody
	GetRequestId() *string
}

type CreateMemoryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0FB1162C-D50B-5DA7-AD04-3417ABBF133A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemoryResponseBody) SetCode(v string) *CreateMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMemoryResponseBody) SetRequestId(v string) *CreateMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
