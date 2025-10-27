// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimilarSecurityEventsQueryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateSimilarSecurityEventsQueryTaskResponse(v *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) *CreateSimilarSecurityEventsQueryTaskResponseBody
	GetCreateSimilarSecurityEventsQueryTaskResponse() *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse
	SetRequestId(v string) *CreateSimilarSecurityEventsQueryTaskResponseBody
	GetRequestId() *string
}

type CreateSimilarSecurityEventsQueryTaskResponseBody struct {
	// The information about the task that queries alert events of the same alert type.
	CreateSimilarSecurityEventsQueryTaskResponse *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse `json:"CreateSimilarSecurityEventsQueryTaskResponse,omitempty" xml:"CreateSimilarSecurityEventsQueryTaskResponse,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5861EE3E-F0B3-48B8-A5DC-A5080BFBE052
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSimilarSecurityEventsQueryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSimilarSecurityEventsQueryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBody) GetCreateSimilarSecurityEventsQueryTaskResponse() *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse {
	return s.CreateSimilarSecurityEventsQueryTaskResponse
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBody) SetCreateSimilarSecurityEventsQueryTaskResponse(v *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) *CreateSimilarSecurityEventsQueryTaskResponseBody {
	s.CreateSimilarSecurityEventsQueryTaskResponse = v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBody) SetRequestId(v string) *CreateSimilarSecurityEventsQueryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBody) Validate() error {
	if s.CreateSimilarSecurityEventsQueryTaskResponse != nil {
		if err := s.CreateSimilarSecurityEventsQueryTaskResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse struct {
	// The status of the task. Valid values:
	//
	// 	- **New**: The task is created.
	//
	// 	- **RetrievingData**: Data is being retrieved.
	//
	// 	- **DataRetrieved**: Data is retrieved.
	//
	// 	- **Processing**: The task is running.
	//
	// 	- **Success**: The task is successful.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **PartialFailed**: The task partially failed.
	//
	// example:
	//
	// New
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 2915
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) GetStatus() *string {
	return s.Status
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) SetStatus(v string) *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse {
	s.Status = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) SetTaskId(v int64) *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse {
	s.TaskId = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) Validate() error {
	return dara.Validate(s)
}
