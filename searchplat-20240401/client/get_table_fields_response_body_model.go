// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTableFieldsResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *GetTableFieldsResponseBody
	GetResult() map[string]interface{}
}

type GetTableFieldsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 58113A95-1858-5674-87E5-192AEE6FD9DD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// {
	//
	//         "classic": "VARCHAR",
	//
	//         "address": "CHAR",
	//
	//         "string": "STRING",
	//
	//         "price": "DECIMAL",
	//
	//         "name": "STRING",
	//
	//         "id": "INT",
	//
	//         "tint": "TINYINT",
	//
	//         "home": "DECIMAL"
	//
	//     }
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetTableFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableFieldsResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *GetTableFieldsResponseBody) SetRequestId(v string) *GetTableFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableFieldsResponseBody) SetResult(v map[string]interface{}) *GetTableFieldsResponseBody {
	s.Result = v
	return s
}

func (s *GetTableFieldsResponseBody) Validate() error {
	return dara.Validate(s)
}
