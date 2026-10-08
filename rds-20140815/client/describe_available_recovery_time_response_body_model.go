// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableRecoveryTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCrossBackupId(v int32) *DescribeAvailableRecoveryTimeResponseBody
	GetCrossBackupId() *int32
	SetRecoveryBeginTime(v string) *DescribeAvailableRecoveryTimeResponseBody
	GetRecoveryBeginTime() *string
	SetRecoveryEndTime(v string) *DescribeAvailableRecoveryTimeResponseBody
	GetRecoveryEndTime() *string
	SetRegionId(v string) *DescribeAvailableRecoveryTimeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeAvailableRecoveryTimeResponseBody
	GetRequestId() *string
}

type DescribeAvailableRecoveryTimeResponseBody struct {
	// The ID of the cross-region backup file.
	//
	// example:
	//
	// 1249****
	CrossBackupId *int32 `json:"CrossBackupId,omitempty" xml:"CrossBackupId,omitempty"`
	// The start time of the restorable time range for the cross-region backup file. The time follows the format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// example:
	//
	// 2024-03-04T21:00:47Z
	RecoveryBeginTime *string `json:"RecoveryBeginTime,omitempty" xml:"RecoveryBeginTime,omitempty"`
	// The end time of the restorable time range for the cross-region backup file. The time follows the format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// example:
	//
	// 2024-03-07T02:23:26Z
	RecoveryEndTime *string `json:"RecoveryEndTime,omitempty" xml:"RecoveryEndTime,omitempty"`
	// The region where the source instance resides.
	//
	// example:
	//
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8CCBF4BA-7CE1-47E1-B49F-E97EA200A40D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableRecoveryTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableRecoveryTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableRecoveryTimeResponseBody) GetCrossBackupId() *int32 {
	return s.CrossBackupId
}

func (s *DescribeAvailableRecoveryTimeResponseBody) GetRecoveryBeginTime() *string {
	return s.RecoveryBeginTime
}

func (s *DescribeAvailableRecoveryTimeResponseBody) GetRecoveryEndTime() *string {
	return s.RecoveryEndTime
}

func (s *DescribeAvailableRecoveryTimeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableRecoveryTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableRecoveryTimeResponseBody) SetCrossBackupId(v int32) *DescribeAvailableRecoveryTimeResponseBody {
	s.CrossBackupId = &v
	return s
}

func (s *DescribeAvailableRecoveryTimeResponseBody) SetRecoveryBeginTime(v string) *DescribeAvailableRecoveryTimeResponseBody {
	s.RecoveryBeginTime = &v
	return s
}

func (s *DescribeAvailableRecoveryTimeResponseBody) SetRecoveryEndTime(v string) *DescribeAvailableRecoveryTimeResponseBody {
	s.RecoveryEndTime = &v
	return s
}

func (s *DescribeAvailableRecoveryTimeResponseBody) SetRegionId(v string) *DescribeAvailableRecoveryTimeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableRecoveryTimeResponseBody) SetRequestId(v string) *DescribeAvailableRecoveryTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableRecoveryTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
