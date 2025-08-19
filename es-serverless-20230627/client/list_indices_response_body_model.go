// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListIndicesResponseBody
	GetRequestId() *string
	SetResult(v []interface{}) *ListIndicesResponseBody
	GetResult() []interface{}
}

type ListIndicesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8C85CCB3-C0C9-521C-B599-F903E14A8793
	RequestId *string       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListIndicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIndicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIndicesResponseBody) GetResult() []interface{} {
	return s.Result
}

func (s *ListIndicesResponseBody) SetRequestId(v string) *ListIndicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndicesResponseBody) SetResult(v []interface{}) *ListIndicesResponseBody {
	s.Result = v
	return s
}

func (s *ListIndicesResponseBody) Validate() error {
	return dara.Validate(s)
}
