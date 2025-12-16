// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTableFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDataSourceTableFieldsResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ListDataSourceTableFieldsResponseBody
	GetResult() map[string]interface{}
}

type ListDataSourceTableFieldsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListDataSourceTableFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTableFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceTableFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourceTableFieldsResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ListDataSourceTableFieldsResponseBody) SetRequestId(v string) *ListDataSourceTableFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceTableFieldsResponseBody) SetResult(v map[string]interface{}) *ListDataSourceTableFieldsResponseBody {
	s.Result = v
	return s
}

func (s *ListDataSourceTableFieldsResponseBody) Validate() error {
	return dara.Validate(s)
}
