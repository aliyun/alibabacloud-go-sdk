// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpgradeClusterAddonsResponseBody
	GetClusterId() *string
	SetRequestId(v string) *UpgradeClusterAddonsResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpgradeClusterAddonsResponseBody
	GetTaskId() *string
}

type UpgradeClusterAddonsResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// cf4299b79b3e34226abfdc80a4bda****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// bfd12953-31cb-42f1-8a36-7b80ec345094
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The task ID.
	//
	// example:
	//
	// T-62a944794ee141074400****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UpgradeClusterAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpgradeClusterAddonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeClusterAddonsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeClusterAddonsResponseBody) SetClusterId(v string) *UpgradeClusterAddonsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *UpgradeClusterAddonsResponseBody) SetRequestId(v string) *UpgradeClusterAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeClusterAddonsResponseBody) SetTaskId(v string) *UpgradeClusterAddonsResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpgradeClusterAddonsResponseBody) Validate() error {
	return dara.Validate(s)
}
