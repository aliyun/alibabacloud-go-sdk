// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterNodepoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeClusterNodepoolResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpgradeClusterNodepoolResponseBody
	GetTaskId() *string
}

type UpgradeClusterNodepoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF518944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// T-5fd211e924e1d0078700xxxx
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UpgradeClusterNodepoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterNodepoolResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeClusterNodepoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeClusterNodepoolResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeClusterNodepoolResponseBody) SetRequestId(v string) *UpgradeClusterNodepoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeClusterNodepoolResponseBody) SetTaskId(v string) *UpgradeClusterNodepoolResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpgradeClusterNodepoolResponseBody) Validate() error {
	return dara.Validate(s)
}
