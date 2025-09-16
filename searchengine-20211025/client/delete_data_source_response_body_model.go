// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataSourceResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DeleteDataSourceResponseBody
	GetResult() map[string]interface{}
}

type DeleteDataSourceResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataSourceResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DeleteDataSourceResponseBody) SetRequestId(v string) *DeleteDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetResult(v map[string]interface{}) *DeleteDataSourceResponseBody {
	s.Result = v
	return s
}

func (s *DeleteDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
