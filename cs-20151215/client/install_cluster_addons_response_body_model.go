// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallClusterAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *InstallClusterAddonsResponseBody
	GetClusterId() *string
	SetRequestId(v string) *InstallClusterAddonsResponseBody
	GetRequestId() *string
	SetTaskId(v string) *InstallClusterAddonsResponseBody
	GetTaskId() *string
}

type InstallClusterAddonsResponseBody struct {
	// 集群ID。
	//
	// example:
	//
	// c82e6987e2961451182edacd74faf****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// 687C5BAA-D103-4993-884B-C35E4314****
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 任务ID。
	//
	// example:
	//
	// T-5a54309c80282e39ea0****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s InstallClusterAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallClusterAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *InstallClusterAddonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallClusterAddonsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *InstallClusterAddonsResponseBody) SetClusterId(v string) *InstallClusterAddonsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *InstallClusterAddonsResponseBody) SetRequestId(v string) *InstallClusterAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallClusterAddonsResponseBody) SetTaskId(v string) *InstallClusterAddonsResponseBody {
	s.TaskId = &v
	return s
}

func (s *InstallClusterAddonsResponseBody) Validate() error {
	return dara.Validate(s)
}
