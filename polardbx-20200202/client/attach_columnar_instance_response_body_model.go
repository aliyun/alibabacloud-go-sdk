// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachColumnarInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *AttachColumnarInstanceResponseBody
	GetDBInstanceName() *string
	SetRequestId(v string) *AttachColumnarInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *AttachColumnarInstanceResponseBody
	GetTaskId() *string
}

type AttachColumnarInstanceResponseBody struct {
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 422922413
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AttachColumnarInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachColumnarInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachColumnarInstanceResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *AttachColumnarInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachColumnarInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *AttachColumnarInstanceResponseBody) SetDBInstanceName(v string) *AttachColumnarInstanceResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *AttachColumnarInstanceResponseBody) SetRequestId(v string) *AttachColumnarInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachColumnarInstanceResponseBody) SetTaskId(v string) *AttachColumnarInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *AttachColumnarInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
