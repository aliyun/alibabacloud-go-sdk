// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ResetAccountPasswordResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *ResetAccountPasswordResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *ResetAccountPasswordResponseBody
	GetTaskId() *int32
}

type ResetAccountPasswordResponseBody struct {
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
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1564657730
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ResetAccountPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAccountPasswordResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ResetAccountPasswordResponseBody) SetDBClusterId(v string) *ResetAccountPasswordResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetTaskId(v int32) *ResetAccountPasswordResponseBody {
	s.TaskId = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
