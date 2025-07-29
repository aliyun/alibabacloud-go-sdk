// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpgradeClusterResponseBody
	GetClusterId() *string
	SetRequestId(v string) *UpgradeClusterResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpgradeClusterResponseBody
	GetTaskId() *string
}

type UpgradeClusterResponseBody struct {
	// Cluster ID.
	//
	// example:
	//
	// c82e6987e2961451182edacd74faf****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 0527ac9a-c899-4341-a21a-****
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Task ID.
	//
	// example:
	//
	// T-5faa48fb31b6b8078d00****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UpgradeClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpgradeClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeClusterResponseBody) SetClusterId(v string) *UpgradeClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *UpgradeClusterResponseBody) SetRequestId(v string) *UpgradeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeClusterResponseBody) SetTaskId(v string) *UpgradeClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpgradeClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
