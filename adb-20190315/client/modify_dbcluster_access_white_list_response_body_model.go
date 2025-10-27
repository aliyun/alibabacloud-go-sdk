// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAccessWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterAccessWhiteListResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *ModifyDBClusterAccessWhiteListResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *ModifyDBClusterAccessWhiteListResponseBody
	GetTaskId() *int32
}

type ModifyDBClusterAccessWhiteListResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1564657730
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDBClusterAccessWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetDBClusterId(v string) *ModifyDBClusterAccessWhiteListResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetRequestId(v string) *ModifyDBClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetTaskId(v int32) *ModifyDBClusterAccessWhiteListResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
