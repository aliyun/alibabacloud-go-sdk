// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduledTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyScheduledTaskResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyScheduledTaskResponseBody
	GetResult() map[string]interface{}
}

type ModifyScheduledTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the scheduled task.
	//
	// example:
	//
	// Array
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyScheduledTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyScheduledTaskResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyScheduledTaskResponseBody) SetRequestId(v string) *ModifyScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScheduledTaskResponseBody) SetResult(v map[string]interface{}) *ModifyScheduledTaskResponseBody {
	s.Result = v
	return s
}

func (s *ModifyScheduledTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
