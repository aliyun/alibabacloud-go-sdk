// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateControlPlaneLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateControlPlaneLogResponseBody
	GetClusterId() *string
	SetRequestId(v string) *UpdateControlPlaneLogResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateControlPlaneLogResponseBody
	GetTaskId() *string
}

type UpdateControlPlaneLogResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// cb95aa626a47740afbf6aa099b650****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 687C5BAA-D103-4993-884B-C35E4314****
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The task ID.
	//
	// example:
	//
	// T-5a54309c80282e39ea00****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UpdateControlPlaneLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateControlPlaneLogResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateControlPlaneLogResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateControlPlaneLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateControlPlaneLogResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateControlPlaneLogResponseBody) SetClusterId(v string) *UpdateControlPlaneLogResponseBody {
	s.ClusterId = &v
	return s
}

func (s *UpdateControlPlaneLogResponseBody) SetRequestId(v string) *UpdateControlPlaneLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateControlPlaneLogResponseBody) SetTaskId(v string) *UpdateControlPlaneLogResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateControlPlaneLogResponseBody) Validate() error {
	return dara.Validate(s)
}
