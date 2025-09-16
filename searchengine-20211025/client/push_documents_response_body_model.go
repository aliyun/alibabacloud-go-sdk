// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PushDocumentsResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *PushDocumentsResponseBody
	GetResult() map[string]interface{}
}

type PushDocumentsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PushDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *PushDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushDocumentsResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *PushDocumentsResponseBody) SetRequestId(v string) *PushDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushDocumentsResponseBody) SetResult(v map[string]interface{}) *PushDocumentsResponseBody {
	s.Result = v
	return s
}

func (s *PushDocumentsResponseBody) Validate() error {
	return dara.Validate(s)
}
