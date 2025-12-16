// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveScheduledTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveScheduledTaskResponseBody
	GetRequestId() *string
	SetResult(v []*int32) *RemoveScheduledTaskResponseBody
	GetResult() []*int32
}

type RemoveScheduledTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result []*int32 `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s RemoveScheduledTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveScheduledTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveScheduledTaskResponseBody) GetResult() []*int32 {
	return s.Result
}

func (s *RemoveScheduledTaskResponseBody) SetRequestId(v string) *RemoveScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveScheduledTaskResponseBody) SetResult(v []*int32) *RemoveScheduledTaskResponseBody {
	s.Result = v
	return s
}

func (s *RemoveScheduledTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
