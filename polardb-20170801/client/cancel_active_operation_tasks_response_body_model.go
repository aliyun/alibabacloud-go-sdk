// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelActiveOperationTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelActiveOperationTasksResponseBody
	GetRequestId() *string
	SetTaskIds(v string) *CancelActiveOperationTasksResponseBody
	GetTaskIds() *string
}

type CancelActiveOperationTasksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 25C70FF3-D49B-594D-BECE-0DE2BA1D8BBB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the O&M events to cancel in batches. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 2355,2352
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s CancelActiveOperationTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelActiveOperationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *CancelActiveOperationTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelActiveOperationTasksResponseBody) GetTaskIds() *string {
	return s.TaskIds
}

func (s *CancelActiveOperationTasksResponseBody) SetRequestId(v string) *CancelActiveOperationTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelActiveOperationTasksResponseBody) SetTaskIds(v string) *CancelActiveOperationTasksResponseBody {
	s.TaskIds = &v
	return s
}

func (s *CancelActiveOperationTasksResponseBody) Validate() error {
	return dara.Validate(s)
}
