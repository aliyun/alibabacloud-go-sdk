// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterAccessWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyClusterAccessWhiteListResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *ModifyClusterAccessWhiteListResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *ModifyClusterAccessWhiteListResponseBody
	GetTaskId() *int32
}

type ModifyClusterAccessWhiteListResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 370D09FD-442A-5225-AAD3-7362CAE39177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1564657730
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyClusterAccessWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterAccessWhiteListResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyClusterAccessWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterAccessWhiteListResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifyClusterAccessWhiteListResponseBody) SetDBClusterId(v string) *ModifyClusterAccessWhiteListResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyClusterAccessWhiteListResponseBody) SetRequestId(v string) *ModifyClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterAccessWhiteListResponseBody) SetTaskId(v int32) *ModifyClusterAccessWhiteListResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyClusterAccessWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
