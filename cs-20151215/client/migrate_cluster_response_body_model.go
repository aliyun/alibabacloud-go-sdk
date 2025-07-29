// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *MigrateClusterResponseBody
	GetClusterId() *string
	SetRequestId(v string) *MigrateClusterResponseBody
	GetRequestId() *string
	SetTaskId(v string) *MigrateClusterResponseBody
	GetTaskId() *string
}

type MigrateClusterResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// c8155823d057948c69a****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The task ID.
	//
	// example:
	//
	// T-62ccd14aacb8db06ca00****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s MigrateClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *MigrateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *MigrateClusterResponseBody) SetClusterId(v string) *MigrateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *MigrateClusterResponseBody) SetRequestId(v string) *MigrateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateClusterResponseBody) SetTaskId(v string) *MigrateClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *MigrateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
