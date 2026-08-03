// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateAICoachTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchCreateAICoachTaskResponseBody
	GetRequestId() *string
	SetTaskIds(v []*string) *BatchCreateAICoachTaskResponseBody
	GetTaskIds() []*string
}

type BatchCreateAICoachTaskResponseBody struct {
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskIds   []*string `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
}

func (s BatchCreateAICoachTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateAICoachTaskResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateAICoachTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateAICoachTaskResponseBody) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *BatchCreateAICoachTaskResponseBody) SetRequestId(v string) *BatchCreateAICoachTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateAICoachTaskResponseBody) SetTaskIds(v []*string) *BatchCreateAICoachTaskResponseBody {
	s.TaskIds = v
	return s
}

func (s *BatchCreateAICoachTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
