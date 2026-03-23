// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateMigrationTargetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ActivateMigrationTargetInstanceResponseBody
	GetDBInstanceName() *string
	SetRequestId(v string) *ActivateMigrationTargetInstanceResponseBody
	GetRequestId() *string
	SetSourceIpAddress(v string) *ActivateMigrationTargetInstanceResponseBody
	GetSourceIpAddress() *string
	SetSourcePort(v int64) *ActivateMigrationTargetInstanceResponseBody
	GetSourcePort() *int64
	SetTaskId(v int64) *ActivateMigrationTargetInstanceResponseBody
	GetTaskId() *int64
}

type ActivateMigrationTargetInstanceResponseBody struct {
	DBInstanceName  *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	SourcePort      *int64  `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	TaskId          *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ActivateMigrationTargetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateMigrationTargetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateMigrationTargetInstanceResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ActivateMigrationTargetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateMigrationTargetInstanceResponseBody) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *ActivateMigrationTargetInstanceResponseBody) GetSourcePort() *int64 {
	return s.SourcePort
}

func (s *ActivateMigrationTargetInstanceResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ActivateMigrationTargetInstanceResponseBody) SetDBInstanceName(v string) *ActivateMigrationTargetInstanceResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *ActivateMigrationTargetInstanceResponseBody) SetRequestId(v string) *ActivateMigrationTargetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateMigrationTargetInstanceResponseBody) SetSourceIpAddress(v string) *ActivateMigrationTargetInstanceResponseBody {
	s.SourceIpAddress = &v
	return s
}

func (s *ActivateMigrationTargetInstanceResponseBody) SetSourcePort(v int64) *ActivateMigrationTargetInstanceResponseBody {
	s.SourcePort = &v
	return s
}

func (s *ActivateMigrationTargetInstanceResponseBody) SetTaskId(v int64) *ActivateMigrationTargetInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *ActivateMigrationTargetInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
