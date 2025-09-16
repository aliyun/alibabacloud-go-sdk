// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRestQueryResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRestQueryResultResponseBody
	GetRequestId() *string
	SetResult(v interface{}) *ListRestQueryResultResponseBody
	GetResult() interface{}
}

type ListRestQueryResultResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F43E8AB4-419C-5F4C-90D6-615590DFAA3C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListRestQueryResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRestQueryResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListRestQueryResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRestQueryResultResponseBody) GetResult() interface{} {
	return s.Result
}

func (s *ListRestQueryResultResponseBody) SetRequestId(v string) *ListRestQueryResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRestQueryResultResponseBody) SetResult(v interface{}) *ListRestQueryResultResponseBody {
	s.Result = v
	return s
}

func (s *ListRestQueryResultResponseBody) Validate() error {
	return dara.Validate(s)
}
