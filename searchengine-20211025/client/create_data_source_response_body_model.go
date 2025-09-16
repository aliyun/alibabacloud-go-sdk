// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDataSourceResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *CreateDataSourceResponseBody
	GetResult() map[string]interface{}
}

type CreateDataSourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataSourceResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *CreateDataSourceResponseBody) SetRequestId(v string) *CreateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetResult(v map[string]interface{}) *CreateDataSourceResponseBody {
	s.Result = v
	return s
}

func (s *CreateDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
