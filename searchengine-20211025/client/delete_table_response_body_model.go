// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTableResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DeleteTableResponseBody
	GetResult() map[string]interface{}
}

type DeleteTableResponseBody struct {
	// requestId
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTableResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DeleteTableResponseBody) SetRequestId(v string) *DeleteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTableResponseBody) SetResult(v map[string]interface{}) *DeleteTableResponseBody {
	s.Result = v
	return s
}

func (s *DeleteTableResponseBody) Validate() error {
	return dara.Validate(s)
}
