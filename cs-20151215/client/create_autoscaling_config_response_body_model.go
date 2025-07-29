// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoscalingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateAutoscalingConfigResponseBody
	GetClusterId() *string
	SetRequestId(v string) *CreateAutoscalingConfigResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateAutoscalingConfigResponseBody
	GetTaskId() *string
}

type CreateAutoscalingConfigResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// cc212d04dfe184547bffaa596********
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AF8BE105-C32B-1269-9774-5510********
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The task ID.
	//
	// example:
	//
	// T-5fd211e924e1d007********
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateAutoscalingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoscalingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoscalingConfigResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateAutoscalingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAutoscalingConfigResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateAutoscalingConfigResponseBody) SetClusterId(v string) *CreateAutoscalingConfigResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateAutoscalingConfigResponseBody) SetRequestId(v string) *CreateAutoscalingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAutoscalingConfigResponseBody) SetTaskId(v string) *CreateAutoscalingConfigResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateAutoscalingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
