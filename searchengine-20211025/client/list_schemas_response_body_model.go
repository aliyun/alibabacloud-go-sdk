// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSchemasResponseBody
	GetRequestId() *string
	SetResult(v interface{}) *ListSchemasResponseBody
	GetResult() interface{}
}

type ListSchemasResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FE03180A-0E29-5474-8A86-33F0683294A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSchemasResponseBody) GetResult() interface{} {
	return s.Result
}

func (s *ListSchemasResponseBody) SetRequestId(v string) *ListSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSchemasResponseBody) SetResult(v interface{}) *ListSchemasResponseBody {
	s.Result = v
	return s
}

func (s *ListSchemasResponseBody) Validate() error {
	return dara.Validate(s)
}
