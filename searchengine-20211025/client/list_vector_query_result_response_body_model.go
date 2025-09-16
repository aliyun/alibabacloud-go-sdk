// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVectorQueryResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVectorQueryResultResponseBody
	GetRequestId() *string
	SetResult(v interface{}) *ListVectorQueryResultResponseBody
	GetResult() interface{}
}

type ListVectorQueryResultResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListVectorQueryResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVectorQueryResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListVectorQueryResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVectorQueryResultResponseBody) GetResult() interface{} {
	return s.Result
}

func (s *ListVectorQueryResultResponseBody) SetRequestId(v string) *ListVectorQueryResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVectorQueryResultResponseBody) SetResult(v interface{}) *ListVectorQueryResultResponseBody {
	s.Result = v
	return s
}

func (s *ListVectorQueryResultResponseBody) Validate() error {
	return dara.Validate(s)
}
