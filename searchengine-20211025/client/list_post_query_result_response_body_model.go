// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPostQueryResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPostQueryResultResponseBody
	GetRequestId() *string
	SetResult(v interface{}) *ListPostQueryResultResponseBody
	GetResult() interface{}
}

type ListPostQueryResultResponseBody struct {
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

func (s ListPostQueryResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPostQueryResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListPostQueryResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPostQueryResultResponseBody) GetResult() interface{} {
	return s.Result
}

func (s *ListPostQueryResultResponseBody) SetRequestId(v string) *ListPostQueryResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPostQueryResultResponseBody) SetResult(v interface{}) *ListPostQueryResultResponseBody {
	s.Result = v
	return s
}

func (s *ListPostQueryResultResponseBody) Validate() error {
	return dara.Validate(s)
}
