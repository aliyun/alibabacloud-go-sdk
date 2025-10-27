// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateAccountResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *CreateAccountResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *CreateAccountResponseBody
	GetTaskId() *int32
}

type CreateAccountResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// example:
	//
	// am-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1564657730
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccountResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *CreateAccountResponseBody) SetDBClusterId(v string) *CreateAccountResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountResponseBody) SetTaskId(v int32) *CreateAccountResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
