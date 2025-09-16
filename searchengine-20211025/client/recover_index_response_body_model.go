// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RecoverIndexResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *RecoverIndexResponseBody
	GetResult() map[string]interface{}
}

type RecoverIndexResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned by data search.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RecoverIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecoverIndexResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecoverIndexResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *RecoverIndexResponseBody) SetRequestId(v string) *RecoverIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverIndexResponseBody) SetResult(v map[string]interface{}) *RecoverIndexResponseBody {
	s.Result = v
	return s
}

func (s *RecoverIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
