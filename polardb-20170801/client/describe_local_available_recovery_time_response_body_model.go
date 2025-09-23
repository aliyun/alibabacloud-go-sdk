// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLocalAvailableRecoveryTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeLocalAvailableRecoveryTimeResponseBody
	GetDBClusterId() *string
	SetRecoveryBeginTime(v string) *DescribeLocalAvailableRecoveryTimeResponseBody
	GetRecoveryBeginTime() *string
	SetRecoveryEndTime(v string) *DescribeLocalAvailableRecoveryTimeResponseBody
	GetRecoveryEndTime() *string
	SetRequestId(v string) *DescribeLocalAvailableRecoveryTimeResponseBody
	GetRequestId() *string
}

type DescribeLocalAvailableRecoveryTimeResponseBody struct {
	// example:
	//
	// pc-2ze3ngi149b313***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 2025-09-10T14:19:48Z
	RecoveryBeginTime *string `json:"RecoveryBeginTime,omitempty" xml:"RecoveryBeginTime,omitempty"`
	// example:
	//
	// 2025-09-17T08:56:45Z
	RecoveryEndTime *string `json:"RecoveryEndTime,omitempty" xml:"RecoveryEndTime,omitempty"`
	// example:
	//
	// 4EA0E6F8-BDB2-17B2-9567-591F6B3D7***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLocalAvailableRecoveryTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocalAvailableRecoveryTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLocalAvailableRecoveryTimeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
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

func (s *DescribeLocalAvailableRecoveryTimeResponseBody) SetDBClusterId(v string) *DescribeLocalAvailableRecoveryTimeResponseBody {
	s.DBClusterId = &v
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
