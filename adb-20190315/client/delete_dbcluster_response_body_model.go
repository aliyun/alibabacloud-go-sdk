// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDBClusterResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *DeleteDBClusterResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *DeleteDBClusterResponseBody
	GetTaskId() *int32
}

type DeleteDBClusterResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// example:
	//
	// am-bp1u8c0mgfg58****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 421693038
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBClusterResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteDBClusterResponseBody) SetDBClusterId(v string) *DeleteDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) SetRequestId(v string) *DeleteDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) SetTaskId(v int32) *DeleteDBClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
