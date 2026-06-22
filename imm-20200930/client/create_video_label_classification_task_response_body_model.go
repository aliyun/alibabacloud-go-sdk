// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoLabelClassificationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateVideoLabelClassificationTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateVideoLabelClassificationTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateVideoLabelClassificationTaskResponseBody
	GetTaskId() *string
}

type CreateVideoLabelClassificationTaskResponseBody struct {
	// The event ID of the task. You can use this ID to query task information notifications in the [EventBridge](https://www.aliyun.com/product/aliware/eventbridge) service.
	//
	// example:
	//
	// 03F-1Qt1Yn5RZZ0Zh3ZdYlDblv7****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CA995EFD-083D-4F40-BE8A-BDF75FFFE0B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task. You can use the [GetTask](https://help.aliyun.com/document_detail/478241.html) operation to view task information or the [GetVideoLabelClassificationResult](https://help.aliyun.com/document_detail/478224.html) operation to obtain the results of the video label detection task.
	//
	// example:
	//
	// VideoLabelClassification-2f157087-91df-4fda-8c3e-232407ec*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateVideoLabelClassificationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoLabelClassificationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoLabelClassificationTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateVideoLabelClassificationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVideoLabelClassificationTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVideoLabelClassificationTaskResponseBody) SetEventId(v string) *CreateVideoLabelClassificationTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskResponseBody) SetRequestId(v string) *CreateVideoLabelClassificationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskResponseBody) SetTaskId(v string) *CreateVideoLabelClassificationTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
