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
	// The ID of the cross-region data backup file.
	//
	// example:
	//
	// 14377
	CrossBackupId *int32 `json:"CrossBackupId,omitempty" xml:"CrossBackupId,omitempty"`
	// The start time from which data can be restored. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-06-12T05:22:29Z
	RecoveryBeginTime *string `json:"RecoveryBeginTime,omitempty" xml:"RecoveryBeginTime,omitempty"`
	// The end time to which data can be restored. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-06-12T07:33:12Z
	RecoveryEndTime *string `json:"RecoveryEndTime,omitempty" xml:"RecoveryEndTime,omitempty"`
	// The region where the source instance resides.
	//
	// example:
	//
	// cn-hangzhou
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
