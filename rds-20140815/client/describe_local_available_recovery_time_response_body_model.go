// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLocalAvailableRecoveryTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeLocalAvailableRecoveryTimeResponseBody
	GetDBInstanceId() *string
	SetRecoveryBeginTime(v string) *DescribeLocalAvailableRecoveryTimeResponseBody
	GetRecoveryBeginTime() *string
	SetRecoveryEndTime(v string) *DescribeLocalAvailableRecoveryTimeResponseBody
	GetRecoveryEndTime() *string
	SetRequestId(v string) *DescribeLocalAvailableRecoveryTimeResponseBody
	GetRequestId() *string
}

type DescribeLocalAvailableRecoveryTimeResponseBody struct {
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RecoveryBeginTime *string `json:"RecoveryBeginTime,omitempty" xml:"RecoveryBeginTime,omitempty"`
	RecoveryEndTime   *string `json:"RecoveryEndTime,omitempty" xml:"RecoveryEndTime,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLocalAvailableRecoveryTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocalAvailableRecoveryTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLocalAvailableRecoveryTimeResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeLocalAvailableRecoveryTimeResponseBody) GetRecoveryBeginTime() *string {
	return s.RecoveryBeginTime
}

func (s *DescribeLocalAvailableRecoveryTimeResponseBody) GetRecoveryEndTime() *string {
	return s.RecoveryEndTime
}

func (s *DescribeLocalAvailableRecoveryTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLocalAvailableRecoveryTimeResponseBody) SetDBInstanceId(v string) *DescribeLocalAvailableRecoveryTimeResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeResponseBody) SetRecoveryBeginTime(v string) *DescribeLocalAvailableRecoveryTimeResponseBody {
	s.RecoveryBeginTime = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeResponseBody) SetRecoveryEndTime(v string) *DescribeLocalAvailableRecoveryTimeResponseBody {
	s.RecoveryEndTime = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeResponseBody) SetRequestId(v string) *DescribeLocalAvailableRecoveryTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
