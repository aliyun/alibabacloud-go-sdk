// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationBackupPolicyResponseBody
	GetApplicationId() *string
	SetApplicationType(v string) *DescribeApplicationBackupPolicyResponseBody
	GetApplicationType() *string
	SetPreferredBackupPeriod(v string) *DescribeApplicationBackupPolicyResponseBody
	GetPreferredBackupPeriod() *string
	SetPreferredBackupRetentionPeriod(v int32) *DescribeApplicationBackupPolicyResponseBody
	GetPreferredBackupRetentionPeriod() *int32
	SetPreferredBackupTime(v string) *DescribeApplicationBackupPolicyResponseBody
	GetPreferredBackupTime() *string
	SetRequestId(v string) *DescribeApplicationBackupPolicyResponseBody
	GetRequestId() *string
}

type DescribeApplicationBackupPolicyResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type, such as polarclaw.
	//
	// example:
	//
	// polarclaw
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The days of the week when automatic backups are performed. Separate multiple values with commas. Valid values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, and Sunday.
	//
	// example:
	//
	// Monday,Tuesday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The retention period of level-1 backups. Unit: days.
	//
	// example:
	//
	// 7
	PreferredBackupRetentionPeriod *int32 `json:"PreferredBackupRetentionPeriod,omitempty" xml:"PreferredBackupRetentionPeriod,omitempty"`
	// The time window when automatic backups are performed. The format is HH:mmZ-HH:mmZ in UTC.
	//
	// example:
	//
	// 07:00Z-08:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApplicationBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationBackupPolicyResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationBackupPolicyResponseBody) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DescribeApplicationBackupPolicyResponseBody) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *DescribeApplicationBackupPolicyResponseBody) GetPreferredBackupRetentionPeriod() *int32 {
	return s.PreferredBackupRetentionPeriod
}

func (s *DescribeApplicationBackupPolicyResponseBody) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *DescribeApplicationBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationBackupPolicyResponseBody) SetApplicationId(v string) *DescribeApplicationBackupPolicyResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationBackupPolicyResponseBody) SetApplicationType(v string) *DescribeApplicationBackupPolicyResponseBody {
	s.ApplicationType = &v
	return s
}

func (s *DescribeApplicationBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeApplicationBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeApplicationBackupPolicyResponseBody) SetPreferredBackupRetentionPeriod(v int32) *DescribeApplicationBackupPolicyResponseBody {
	s.PreferredBackupRetentionPeriod = &v
	return s
}

func (s *DescribeApplicationBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeApplicationBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeApplicationBackupPolicyResponseBody) SetRequestId(v string) *DescribeApplicationBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
