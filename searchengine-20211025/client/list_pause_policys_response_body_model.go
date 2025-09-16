// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPausePolicysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPausePolicysResponseBody
	GetRequestId() *string
	SetResult(v map[string]*ResultValue) *ListPausePolicysResponseBody
	GetResult() map[string]*ResultValue
}

type ListPausePolicysResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	Result map[string]*ResultValue `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListPausePolicysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPausePolicysResponseBody) GoString() string {
	return s.String()
}

func (s *ListPausePolicysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPausePolicysResponseBody) GetResult() map[string]*ResultValue {
	return s.Result
}

func (s *ListPausePolicysResponseBody) SetRequestId(v string) *ListPausePolicysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPausePolicysResponseBody) SetResult(v map[string]*ResultValue) *ListPausePolicysResponseBody {
	s.Result = v
	return s
}

func (s *ListPausePolicysResponseBody) Validate() error {
	return dara.Validate(s)
}
