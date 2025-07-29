// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyClusterResponseBody
	GetClusterId() *string
	SetRequestId(v string) *ModifyClusterResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyClusterResponseBody
	GetTaskId() *string
}

type ModifyClusterResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// cb95aa626a47740afbf6aa09****
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

func (s ModifyClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyClusterResponseBody) SetClusterId(v string) *ModifyClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterResponseBody) SetRequestId(v string) *ModifyClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterResponseBody) SetTaskId(v string) *ModifyClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
