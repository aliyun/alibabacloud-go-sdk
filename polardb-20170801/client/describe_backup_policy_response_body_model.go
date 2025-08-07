// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupFrequency(v string) *DescribeBackupPolicyResponseBody
	GetBackupFrequency() *string
	SetBackupRetentionPolicyOnClusterDeletion(v string) *DescribeBackupPolicyResponseBody
	GetBackupRetentionPolicyOnClusterDeletion() *string
	SetDataLevel1BackupFrequency(v string) *DescribeBackupPolicyResponseBody
	GetDataLevel1BackupFrequency() *string
	SetDataLevel1BackupPeriod(v string) *DescribeBackupPolicyResponseBody
	GetDataLevel1BackupPeriod() *string
	SetDataLevel1BackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody
	GetDataLevel1BackupRetentionPeriod() *string
	SetDataLevel1BackupTime(v string) *DescribeBackupPolicyResponseBody
	GetDataLevel1BackupTime() *string
	SetDataLevel2BackupAnotherRegionRegion(v string) *DescribeBackupPolicyResponseBody
	GetDataLevel2BackupAnotherRegionRegion() *string
	SetDataLevel2BackupAnotherRegionRetentionPeriod(v string) *DescribeBackupPolicyResponseBody
	GetDataLevel2BackupAnotherRegionRetentionPeriod() *string
	SetDataLevel2BackupPeriod(v string) *DescribeBackupPolicyResponseBody
	GetDataLevel2BackupPeriod() *string
	SetDataLevel2BackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody
	GetDataLevel2BackupRetentionPeriod() *string
	SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupTime() *string
	SetPreferredNextBackupTime(v string) *DescribeBackupPolicyResponseBody
	GetPreferredNextBackupTime() *string
	SetRequestId(v string) *DescribeBackupPolicyResponseBody
	GetRequestId() *string
}

type DescribeBackupPolicyResponseBody struct {
	// The backup frequency. Default value: Normal. Valid values:
	//
	// 	- **Normal**: standard backup. The system backs up data once a day.
	//
	// 	- **2/24H**: enhanced backup. The system backs up data every 2 hours.
	//
	// 	- **3/24H**: enhanced backup. The system backs up data every 3 hours.
	//
	// 	- **4/24H**: enhanced backup. The system backs up data every 4 hours.
	//
	// > - If enhanced backup is enabled, all backups are retained for 24 hours. Backups are automatically deleted when the retention period ends. However, the system permanently retains the first backup that is created after 00:00 every day.
	//
	// >-  If enhanced backup is enabled, **PreferredBackupPeriod*	- is automatically set to all days in a week (from Monday to Sunday).
	//
	// example:
	//
	// Normal
	BackupFrequency *string `json:"BackupFrequency,omitempty" xml:"BackupFrequency,omitempty"`
	// Indicates whether backups are retained when you delete a cluster. Valid values:
	//
	// 	- **ALL**: permanently retains all backups.
	//
	// 	- **LATEST**: permanently retains the most recent backup.
	//
	// 	- **NONE**: does not retain backups.
	//
	// example:
	//
	// NONE
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// The backup frequency of level-1 backups. Default value: Normal. Valid values:
	//
	// 	- **Normal**: standard backup. The system backs up data once a day.
	//
	// 	- **2/24H**: frequent backup. The system backs up data every 2 hours.
	//
	// 	- **3/24H**: frequent backup. The system backs up data every 3 hours.
	//
	// 	- **4/24H**: frequent backup. The system backs up data every 4 hours.
	//
	// >- This parameter is not supported for PolarDB for PostgreSQL (Compatible with Oracle) clusters or PolarDB for PostgreSQL clusters.
	//
	// >- This parameter is unavailable if the region where your PolarDB for MySQL cluster is deployed does not support the cross-region backup feature. For information about regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// Normal
	DataLevel1BackupFrequency *string `json:"DataLevel1BackupFrequency,omitempty" xml:"DataLevel1BackupFrequency,omitempty"`
	// The backup cycle of level-1 backups. Valid values:
	//
	// 	- **Monday**
	//
	// 	- **Tuesday**
	//
	// 	- **Wednesday**
	//
	// 	- **Thursday**
	//
	// 	- **Friday**
	//
	// 	- **Saturday**
	//
	// 	- **Sunday**
	//
	// >- You need to specify at least two values. Separate multiple values with commas (,).
	//
	// >- This parameter is not supported for PolarDB for PostgreSQL (Compatible with Oracle) clusters or PolarDB for PostgreSQL clusters.
	//
	// >- This parameter is unavailable if the region where your PolarDB for MySQL cluster is deployed does not support the cross-region backup feature. For information about regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// Monday,Tuesday
	DataLevel1BackupPeriod *string `json:"DataLevel1BackupPeriod,omitempty" xml:"DataLevel1BackupPeriod,omitempty"`
	// The retention period of level-1 backups. Valid values: 3 to 14. Unit: day.
	//
	// example:
	//
	// 7
	DataLevel1BackupRetentionPeriod *string `json:"DataLevel1BackupRetentionPeriod,omitempty" xml:"DataLevel1BackupRetentionPeriod,omitempty"`
	// The period of time during which automatic backup is performed. The value must be in the `hh:mmZ-hh:mmZ` format. The time must be in UTC. The start time and the end time must be on the hour and must have an interval of 1 hour. Example: `14:00Z-15:00Z`.
	//
	// >- This parameter is not supported for PolarDB for PostgreSQL (Compatible with Oracle) clusters or PolarDB for PostgreSQL clusters.
	//
	// >- This parameter is unavailable if the region where your PolarDB for MySQL cluster is deployed does not support the cross-region backup feature. For information about regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// 15:00Z-16:00Z
	DataLevel1BackupTime *string `json:"DataLevel1BackupTime,omitempty" xml:"DataLevel1BackupTime,omitempty"`
	// The region where the cross-region level-2 backup is stored. For information about regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// cn-hangzhou
	DataLevel2BackupAnotherRegionRegion *string `json:"DataLevel2BackupAnotherRegionRegion,omitempty" xml:"DataLevel2BackupAnotherRegionRegion,omitempty"`
	// The retention period of cross-region level-2 backups. Valid values:
	//
	// 	- **0**: The cross-region level-2 backup feature is disabled.
	//
	// 	- **30 to 7300**: Cross-region level-2 backups are retained for 30 to 7,300 days.
	//
	// 	- **1**: Cross-region level-2 backups are retained for a long period of time.
	//
	// >  The default value of the parameter is **0**.
	//
	// example:
	//
	// 30
	DataLevel2BackupAnotherRegionRetentionPeriod *string `json:"DataLevel2BackupAnotherRegionRetentionPeriod,omitempty" xml:"DataLevel2BackupAnotherRegionRetentionPeriod,omitempty"`
	// The backup cycle of level-2 backups. Valid values:
	//
	// 	- **Monday**
	//
	// 	- **Tuesday**
	//
	// 	- **Wednesday**
	//
	// 	- **Thursday**
	//
	// 	- **Friday**
	//
	// 	- **Saturday**
	//
	// 	- **Sunday**
	//
	// >- You need to specify at least two values. Separate multiple values with commas (,).
	//
	// >- This parameter is not supported for PolarDB for PostgreSQL (Compatible with Oracle) clusters or PolarDB for PostgreSQL clusters.
	//
	// >- This parameter is unavailable if the region where your PolarDB for MySQL cluster is deployed does not support the cross-region backup feature. For information about regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// Monday,Tuesday
	DataLevel2BackupPeriod *string `json:"DataLevel2BackupPeriod,omitempty" xml:"DataLevel2BackupPeriod,omitempty"`
	// The retention period of level-2 backups. Valid values:
	//
	// 	- 0: The level-2 backup feature is disabled.
	//
	// 	- 30 to 7300: Level-2 backups are retained for 30 to 7,300 days.
	//
	// 	- \\-1: Level-2 backups are retained for a long period of time.
	//
	// >  The default value of this parameter is **0**.
	//
	// example:
	//
	// 0
	DataLevel2BackupRetentionPeriod *string `json:"DataLevel2BackupRetentionPeriod,omitempty" xml:"DataLevel2BackupRetentionPeriod,omitempty"`
	// The backup cycle. Valid values:
	//
	// 	- Monday
	//
	// 	- Tuesday
	//
	// 	- Wednesday
	//
	// 	- Thursday
	//
	// 	- Friday
	//
	// 	- Saturday
	//
	// 	- Sunday
	//
	// example:
	//
	// Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time period when automatic backups are performed. The format is `HH:mmZ-HH:mmZ`. The time is displayed in UTC.
	//
	// example:
	//
	// 07:00Z-08:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The time when the next automatic backup will be performed. The format is `YYYY-MM-DDThh:mmZ`. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-11-16T07:30Z
	PreferredNextBackupTime *string `json:"PreferredNextBackupTime,omitempty" xml:"PreferredNextBackupTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EADFCE0F-9FB5-4685-B395-1440B******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) GetBackupFrequency() *string {
	return s.BackupFrequency
}

func (s *DescribeBackupPolicyResponseBody) GetBackupRetentionPolicyOnClusterDeletion() *string {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *DescribeBackupPolicyResponseBody) GetDataLevel1BackupFrequency() *string {
	return s.DataLevel1BackupFrequency
}

func (s *DescribeBackupPolicyResponseBody) GetDataLevel1BackupPeriod() *string {
	return s.DataLevel1BackupPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetDataLevel1BackupRetentionPeriod() *string {
	return s.DataLevel1BackupRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetDataLevel1BackupTime() *string {
	return s.DataLevel1BackupTime
}

func (s *DescribeBackupPolicyResponseBody) GetDataLevel2BackupAnotherRegionRegion() *string {
	return s.DataLevel2BackupAnotherRegionRegion
}

func (s *DescribeBackupPolicyResponseBody) GetDataLevel2BackupAnotherRegionRetentionPeriod() *string {
	return s.DataLevel2BackupAnotherRegionRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetDataLevel2BackupPeriod() *string {
	return s.DataLevel2BackupPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetDataLevel2BackupRetentionPeriod() *string {
	return s.DataLevel2BackupRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredNextBackupTime() *string {
	return s.PreferredNextBackupTime
}

func (s *DescribeBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPolicyResponseBody) SetBackupFrequency(v string) *DescribeBackupPolicyResponseBody {
	s.BackupFrequency = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPolicyOnClusterDeletion(v string) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetDataLevel1BackupFrequency(v string) *DescribeBackupPolicyResponseBody {
	s.DataLevel1BackupFrequency = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetDataLevel1BackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.DataLevel1BackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetDataLevel1BackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.DataLevel1BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetDataLevel1BackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.DataLevel1BackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetDataLevel2BackupAnotherRegionRegion(v string) *DescribeBackupPolicyResponseBody {
	s.DataLevel2BackupAnotherRegionRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetDataLevel2BackupAnotherRegionRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.DataLevel2BackupAnotherRegionRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetDataLevel2BackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.DataLevel2BackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetDataLevel2BackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.DataLevel2BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredNextBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredNextBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
