// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateScheduledTaskResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *CreateScheduledTaskResponseBody
	GetResult() map[string]interface{}
}

type CreateScheduledTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// For more information about a scheduled task, see [ScheduledTask](https://help.aliyun.com/document_detail/173610.html).
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateScheduledTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScheduledTaskResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *CreateScheduledTaskResponseBody) SetRequestId(v string) *CreateScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetResult(v map[string]interface{}) *CreateScheduledTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateScheduledTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
