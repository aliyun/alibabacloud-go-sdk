// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceForRebuildResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateDBInstanceForRebuildResponseBody
	GetDBInstanceId() *string
	SetMessage(v string) *CreateDBInstanceForRebuildResponseBody
	GetMessage() *string
	SetOrderId(v string) *CreateDBInstanceForRebuildResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateDBInstanceForRebuildResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateDBInstanceForRebuildResponseBody
	GetTaskId() *string
}

type CreateDBInstanceForRebuildResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDBInstanceForRebuildResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceForRebuildResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceForRebuildResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstanceForRebuildResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDBInstanceForRebuildResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDBInstanceForRebuildResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBInstanceForRebuildResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateDBInstanceForRebuildResponseBody) SetDBInstanceId(v string) *CreateDBInstanceForRebuildResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceForRebuildResponseBody) SetMessage(v string) *CreateDBInstanceForRebuildResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDBInstanceForRebuildResponseBody) SetOrderId(v string) *CreateDBInstanceForRebuildResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBInstanceForRebuildResponseBody) SetRequestId(v string) *CreateDBInstanceForRebuildResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceForRebuildResponseBody) SetTaskId(v string) *CreateDBInstanceForRebuildResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateDBInstanceForRebuildResponseBody) Validate() error {
	return dara.Validate(s)
}
