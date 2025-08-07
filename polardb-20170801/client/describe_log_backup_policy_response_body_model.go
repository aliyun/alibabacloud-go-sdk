// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnableBackupLog(v int32) *DescribeLogBackupPolicyResponseBody
	GetEnableBackupLog() *int32
	SetLogBackupAnotherRegionRegion(v string) *DescribeLogBackupPolicyResponseBody
	GetLogBackupAnotherRegionRegion() *string
	SetLogBackupAnotherRegionRetentionPeriod(v string) *DescribeLogBackupPolicyResponseBody
	GetLogBackupAnotherRegionRetentionPeriod() *string
	SetLogBackupRetentionPeriod(v int32) *DescribeLogBackupPolicyResponseBody
	GetLogBackupRetentionPeriod() *int32
	SetRequestId(v string) *DescribeLogBackupPolicyResponseBody
	GetRequestId() *string
}

type DescribeLogBackupPolicyResponseBody struct {
	// Indicates whether the log backup feature is enabled. Valid values:
	//
	// 	- 0: The log backup feature is disabled.
	//
	// 	- 1: The log backup feature is enabled. By default, the log backup feature is enabled and cannot be disabled.
	//
	// example:
	//
	// 1
	EnableBackupLog *int32 `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// The region in which you want to store cross-region log backups. For more information about regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// cn-beijing
	LogBackupAnotherRegionRegion *string `json:"LogBackupAnotherRegionRegion,omitempty" xml:"LogBackupAnotherRegionRegion,omitempty"`
	// The retention period of cross-region log backups. Valid values:
	//
	// 	- **0**: The cross-region backup feature is disabled.
	//
	// 	- **30 to 7300**: Cross-region log backups are retained for 30 to 7,300 days.
	//
	// 	- **-1**: The log backups are permanently retained.
	//
	// >  When you create a cluster, the default value of this parameter is **0**.
	//
	// example:
	//
	// 0
	LogBackupAnotherRegionRetentionPeriod *string `json:"LogBackupAnotherRegionRetentionPeriod,omitempty" xml:"LogBackupAnotherRegionRetentionPeriod,omitempty"`
	// The retention period of the log backups. Valid values:
	//
	// 	- 3 to 7300: The log backups are retained for 3 to 7,300 days.
	//
	// 	- \\-1: The log backups are permanently retained.
	//
	// example:
	//
	// 7
	LogBackupRetentionPeriod *int32 `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 62EE0051-102B-488D-9C79-D607B8******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLogBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupPolicyResponseBody) GetEnableBackupLog() *int32 {
	return s.EnableBackupLog
}

func (s *DescribeLogBackupPolicyResponseBody) GetLogBackupAnotherRegionRegion() *string {
	return s.LogBackupAnotherRegionRegion
}

func (s *DescribeLogBackupPolicyResponseBody) GetLogBackupAnotherRegionRetentionPeriod() *string {
	return s.LogBackupAnotherRegionRetentionPeriod
}

func (s *DescribeLogBackupPolicyResponseBody) GetLogBackupRetentionPeriod() *int32 {
	return s.LogBackupRetentionPeriod
}

func (s *DescribeLogBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogBackupPolicyResponseBody) SetEnableBackupLog(v int32) *DescribeLogBackupPolicyResponseBody {
	s.EnableBackupLog = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBody) SetLogBackupAnotherRegionRegion(v string) *DescribeLogBackupPolicyResponseBody {
	s.LogBackupAnotherRegionRegion = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBody) SetLogBackupAnotherRegionRetentionPeriod(v string) *DescribeLogBackupPolicyResponseBody {
	s.LogBackupAnotherRegionRetentionPeriod = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBody) SetLogBackupRetentionPeriod(v int32) *DescribeLogBackupPolicyResponseBody {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBody) SetRequestId(v string) *DescribeLogBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
