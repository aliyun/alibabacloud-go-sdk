// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTableResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *CreateTableResponseBody
	GetResult() map[string]interface{}
}

type CreateTableResponseBody struct {
	// id of request
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTableResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *CreateTableResponseBody) SetRequestId(v string) *CreateTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTableResponseBody) SetResult(v map[string]interface{}) *CreateTableResponseBody {
	s.Result = v
	return s
}

func (s *CreateTableResponseBody) Validate() error {
	return dara.Validate(s)
}
