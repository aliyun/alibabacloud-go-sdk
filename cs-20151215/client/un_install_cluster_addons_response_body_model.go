// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnInstallClusterAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UnInstallClusterAddonsResponseBody
	GetClusterId() *string
	SetRequestId(v string) *UnInstallClusterAddonsResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UnInstallClusterAddonsResponseBody
	GetTaskId() *string
}

type UnInstallClusterAddonsResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// c5b5e80b0b64a4bf6939d2d8fbbc5****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 74D1512F-67DA-54E8-99EA-4D50EB4898F4
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The task ID.
	//
	// example:
	//
	// T-66e39b39c0fdd001320005c0
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UnInstallClusterAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnInstallClusterAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *UnInstallClusterAddonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnInstallClusterAddonsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UnInstallClusterAddonsResponseBody) SetClusterId(v string) *UnInstallClusterAddonsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *UnInstallClusterAddonsResponseBody) SetRequestId(v string) *UnInstallClusterAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnInstallClusterAddonsResponseBody) SetTaskId(v string) *UnInstallClusterAddonsResponseBody {
	s.TaskId = &v
	return s
}

func (s *UnInstallClusterAddonsResponseBody) Validate() error {
	return dara.Validate(s)
}
