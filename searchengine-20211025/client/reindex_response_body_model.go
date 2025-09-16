// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReindexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReindexResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ReindexResponseBody
	GetResult() map[string]interface{}
}

type ReindexResponseBody struct {
	// requestId
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ReindexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReindexResponseBody) GoString() string {
	return s.String()
}

func (s *ReindexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReindexResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ReindexResponseBody) SetRequestId(v string) *ReindexResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReindexResponseBody) SetResult(v map[string]interface{}) *ReindexResponseBody {
	s.Result = v
	return s
}

func (s *ReindexResponseBody) Validate() error {
	return dara.Validate(s)
}
