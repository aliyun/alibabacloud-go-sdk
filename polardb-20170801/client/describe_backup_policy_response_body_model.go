// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedDataPolicies(v *DescribeBackupPolicyResponseBodyAdvancedDataPolicies) *DescribeBackupPolicyResponseBody
	GetAdvancedDataPolicies() *DescribeBackupPolicyResponseBodyAdvancedDataPolicies
	SetAdvancedPolicyOption(v string) *DescribeBackupPolicyResponseBody
	GetAdvancedPolicyOption() *string
	SetBackupFrequency(v string) *DescribeBackupPolicyResponseBody
	GetBackupFrequency() *string
	SetBackupPolicyLevel(v string) *DescribeBackupPolicyResponseBody
	GetBackupPolicyLevel() *string
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
	SetEnableCrossRegionImmutableBackup(v bool) *DescribeBackupPolicyResponseBody
	GetEnableCrossRegionImmutableBackup() *bool
	SetEnableImmutableBackup(v bool) *DescribeBackupPolicyResponseBody
	GetEnableImmutableBackup() *bool
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
	AdvancedDataPolicies *DescribeBackupPolicyResponseBodyAdvancedDataPolicies `json:"AdvancedDataPolicies,omitempty" xml:"AdvancedDataPolicies,omitempty" type:"Struct"`
	// The advanced backup policy option. Valid values:
	//
	// 	- **enable**: Advanced backup is enabled.
	//
	// 	- **disable**: Advanced backup is not enabled but can be enabled.
	//
	// 	- **notSupport**: Advanced backup is not supported.
	//
	// > 	- This parameter is not supported for PolarDB for PostgreSQL (Compatible with Oracle) or PolarDB for PostgreSQL.
	//
	// example:
	//
	// disable
	AdvancedPolicyOption *string `json:"AdvancedPolicyOption,omitempty" xml:"AdvancedPolicyOption,omitempty"`
	// The backup frequency. Valid values:
	//
	// - **Normal*	- (default): regular backup. A backup is performed once a day at a scheduled time.
	//
	// - **2/24H**: enhanced backup. A backup is performed every 2 hours.
	//
	// - **3/24H**: enhanced backup. A backup is performed every 3 hours.
	//
	// - **4/24H**: enhanced backup. A backup is performed every 4 hours.
	//
	// > 	- After enhanced backup is enabled, all backups completed within 24 hours are retained. For backups older than 24 hours, only the first backup completed after 00:00 each day is retained, and all others are deleted.
	//
	// >	- After enhanced backup is enabled, the backup cycle parameter **PreferredBackupPeriod*	- is set to all days of the week by default (Monday through Sunday).
	//
	// example:
	//
	// Normal
	BackupFrequency *string `json:"BackupFrequency,omitempty" xml:"BackupFrequency,omitempty"`
	// The current backup policy level. Valid values:
	//
	// 	- **Normal**: regular backup
	//
	// 	- **Advanced**: advanced backup
	//
	// > 	- This parameter is not supported for PolarDB for PostgreSQL (Compatible with Oracle) or PolarDB for PostgreSQL.
	//
	// example:
	//
	// Normal
	BackupPolicyLevel *string `json:"BackupPolicyLevel,omitempty" xml:"BackupPolicyLevel,omitempty"`
	// Specifies whether to retain backups when the cluster is deleted. Valid values:
	//
	// 	- **ALL**: All backups are retained with long-term retention (LTR).
	//
	// 	- **LATEST**: The last backup is retained with long-term retention (LTR).
	//
	// 	- **NONE*	- (default): No backups are retained.
	//
	// example:
	//
	// NONE
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// The backup frequency. Valid values:
	//
	// - **Normal*	- (default): regular backup. A backup is performed once a day at a scheduled time.
	//
	// - **2/24H**: high-frequency backup. A backup is performed every 2 hours.
	//
	// - **3/24H**: high-frequency backup. A backup is performed every 3 hours.
	//
	// - **4/24H**: high-frequency backup. A backup is performed every 4 hours.
	//
	// > 	- 	- This parameter is not supported for PolarDB for PostgreSQL (Compatible with Oracle) or PolarDB for PostgreSQL.
	//
	// >	- 	- If the region of your PolarDB for MySQL cluster does not support the cross-region backup feature, this parameter is not supported. For the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// >	- 	- After advanced backup is enabled, use the AdvancedDataPolicies parameter instead of this parameter.
	//
	// example:
	//
	// Normal
	DataLevel1BackupFrequency *string `json:"DataLevel1BackupFrequency,omitempty" xml:"DataLevel1BackupFrequency,omitempty"`
	// The level-1 backup cycle. Valid values:
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
	// > 	- 	- At least two days must be selected. Separate multiple values with commas (,).
	//
	// >	- 	- This parameter is not supported for PolarDB for PostgreSQL (Compatible with Oracle) or PolarDB for PostgreSQL.
	//
	// >	- 	- If the region of your PolarDB for MySQL cluster does not support the cross-region backup feature, this parameter is not supported. For the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// >	- 	- After advanced backup is enabled, use the AdvancedDataPolicies parameter instead of this parameter.
	//
	// example:
	//
	// Monday,Tuesday
	DataLevel1BackupPeriod *string `json:"DataLevel1BackupPeriod,omitempty" xml:"DataLevel1BackupPeriod,omitempty"`
	// The retention period of level-1 backups. Valid values: 3 to 14. Unit: days.
	//
	// > After advanced backup is enabled, use the AdvancedDataPolicies parameter instead of this parameter.
	//
	// example:
	//
	// 7
	DataLevel1BackupRetentionPeriod *string `json:"DataLevel1BackupRetentionPeriod,omitempty" xml:"DataLevel1BackupRetentionPeriod,omitempty"`
	// The time period during which automatic backups are performed. The value is in the `hh:mmZ-hh:mmZ` format (UTC). The start and end times must be on the hour and exactly 1 hour apart. Example: `14:00Z-15:00Z`.
	//
	// > 	- This parameter is not supported for PolarDB for PostgreSQL (Compatible with Oracle) or PolarDB for PostgreSQL.
	//
	// >	- If the region of your PolarDB for MySQL cluster does not support the cross-region backup feature, this parameter is not supported. For the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// 15:00Z-16:00Z
	DataLevel1BackupTime *string `json:"DataLevel1BackupTime,omitempty" xml:"DataLevel1BackupTime,omitempty"`
	// The cross-region backup region for level-2 backups. For the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// > After advanced backup is enabled, use the AdvancedDataPolicies parameter instead of this parameter.
	//
	// example:
	//
	// cn-hangzhou
	DataLevel2BackupAnotherRegionRegion *string `json:"DataLevel2BackupAnotherRegionRegion,omitempty" xml:"DataLevel2BackupAnotherRegionRegion,omitempty"`
	// The retention epoch of cross-region backups for level-2 backups. Valid values:
	//
	// - **0**: The level-2 backup feature is shutdown.
	//
	// - **30 to 7300**: The retention epoch of level-2 backups. Unit: days.
	//
	// - **-1**: Level-2 backups are retained with long-term retention (LTR).
	//
	//  >
	//
	// >- - When a cluster is created, the default value is **0**, which means the cross-region backup feature for level-2 backups is shutdown.
	//
	// >- - After advanced backup is enabled, use the AdvancedDataPolicies parameter instead of this parameter.
	//
	// example:
	//
	// 30
	DataLevel2BackupAnotherRegionRetentionPeriod *string `json:"DataLevel2BackupAnotherRegionRetentionPeriod,omitempty" xml:"DataLevel2BackupAnotherRegionRetentionPeriod,omitempty"`
	// The level-2 backup cycle. Valid values:
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
	// > 	- 	- At least two days must be selected. Separate multiple values with commas (,).
	//
	// >	- 	- This parameter is not supported for PolarDB for PostgreSQL (Compatible with Oracle) or PolarDB for PostgreSQL.
	//
	// >	- 	- If the region of your PolarDB for MySQL cluster does not support the cross-region backup feature, this parameter is not supported. For the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// >	- 	- After advanced backup is enabled, use the AdvancedDataPolicies parameter instead of this parameter.
	//
	// example:
	//
	// Monday,Tuesday
	DataLevel2BackupPeriod *string `json:"DataLevel2BackupPeriod,omitempty" xml:"DataLevel2BackupPeriod,omitempty"`
	// The retention epoch of level-2 backups. Valid values:
	//
	//  	- 0: The level-2 backup feature is shutdown.
	//
	//  	- 30 to 7300: The retention epoch of level-2 backups. Unit: days.
	//
	//  	- -1: Level-2 backups are retained with long-term retention (LTR).
	//
	//  >
	//
	// >- 	- When a cluster is created, the default value is **0**, which means the level-2 backup feature is shutdown.
	//
	// >- 	- After advanced backup is enabled, use the AdvancedDataPolicies parameter instead of this parameter.
	//
	// example:
	//
	// 0
	DataLevel2BackupRetentionPeriod *string `json:"DataLevel2BackupRetentionPeriod,omitempty" xml:"DataLevel2BackupRetentionPeriod,omitempty"`
	// Indicates whether immutable cross-region backup is enabled.
	EnableCrossRegionImmutableBackup *bool `json:"EnableCrossRegionImmutableBackup,omitempty" xml:"EnableCrossRegionImmutableBackup,omitempty"`
	// Indicates whether immutable backup is enabled.
	EnableImmutableBackup *bool `json:"EnableImmutableBackup,omitempty" xml:"EnableImmutableBackup,omitempty"`
	// The data backup cycle. Valid values:
	//
	// - Monday
	//
	// - Tuesday
	//
	// - Wednesday
	//
	// - Thursday
	//
	// - Friday
	//
	// - Saturday
	//
	// - Sunday
	//
	// > After advanced backup is enabled, use the AdvancedDataPolicies parameter instead of this parameter.
	//
	// example:
	//
	// Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time period during which automatic backups are performed. The value is in the `HH:mmZ-HH:mmZ` format (UTC).
	//
	// example:
	//
	// 07:00Z-08:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The time of the next backup. The value is in the `YYYY-MM-DDThh:mmZ` format (UTC).
	//
	// example:
	//
	// 2020-11-16T07:30Z
	PreferredNextBackupTime *string `json:"PreferredNextBackupTime,omitempty" xml:"PreferredNextBackupTime,omitempty"`
	// The request ID.
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

func (s *DescribeBackupPolicyResponseBody) GetAdvancedDataPolicies() *DescribeBackupPolicyResponseBodyAdvancedDataPolicies {
	return s.AdvancedDataPolicies
}

func (s *DescribeBackupPolicyResponseBody) GetAdvancedPolicyOption() *string {
	return s.AdvancedPolicyOption
}

func (s *DescribeBackupPolicyResponseBody) GetBackupFrequency() *string {
	return s.BackupFrequency
}

func (s *DescribeBackupPolicyResponseBody) GetBackupPolicyLevel() *string {
	return s.BackupPolicyLevel
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

func (s *DescribeBackupPolicyResponseBody) GetEnableCrossRegionImmutableBackup() *bool {
	return s.EnableCrossRegionImmutableBackup
}

func (s *DescribeBackupPolicyResponseBody) GetEnableImmutableBackup() *bool {
	return s.EnableImmutableBackup
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

func (s *DescribeBackupPolicyResponseBody) SetAdvancedDataPolicies(v *DescribeBackupPolicyResponseBodyAdvancedDataPolicies) *DescribeBackupPolicyResponseBody {
	s.AdvancedDataPolicies = v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetAdvancedPolicyOption(v string) *DescribeBackupPolicyResponseBody {
	s.AdvancedPolicyOption = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupFrequency(v string) *DescribeBackupPolicyResponseBody {
	s.BackupFrequency = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupPolicyLevel(v string) *DescribeBackupPolicyResponseBody {
	s.BackupPolicyLevel = &v
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

func (s *DescribeBackupPolicyResponseBody) SetEnableCrossRegionImmutableBackup(v bool) *DescribeBackupPolicyResponseBody {
	s.EnableCrossRegionImmutableBackup = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnableImmutableBackup(v bool) *DescribeBackupPolicyResponseBody {
	s.EnableImmutableBackup = &v
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
	if s.AdvancedDataPolicies != nil {
		if err := s.AdvancedDataPolicies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupPolicyResponseBodyAdvancedDataPolicies struct {
	AdvancedDataPolicy []*DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy `json:"AdvancedDataPolicy,omitempty" xml:"AdvancedDataPolicy,omitempty" type:"Repeated"`
}

func (s DescribeBackupPolicyResponseBodyAdvancedDataPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyAdvancedDataPolicies) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPolicies) GetAdvancedDataPolicy() []*DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	return s.AdvancedDataPolicy
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPolicies) SetAdvancedDataPolicy(v []*DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) *DescribeBackupPolicyResponseBodyAdvancedDataPolicies {
	s.AdvancedDataPolicy = v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPolicies) Validate() error {
	if s.AdvancedDataPolicy != nil {
		for _, item := range s.AdvancedDataPolicy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy struct {
	AutoCreated             *bool   `json:"AutoCreated,omitempty" xml:"AutoCreated,omitempty"`
	BakType                 *string `json:"BakType,omitempty" xml:"BakType,omitempty"`
	DestRegion              *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	DestType                *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	DumpAction              *string `json:"DumpAction,omitempty" xml:"DumpAction,omitempty"`
	FilterKey               *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterType              *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	FilterValue             *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	OnlyPreserveOneEachDay  *bool   `json:"OnlyPreserveOneEachDay,omitempty" xml:"OnlyPreserveOneEachDay,omitempty"`
	OnlyPreserveOneEachHour *bool   `json:"OnlyPreserveOneEachHour,omitempty" xml:"OnlyPreserveOneEachHour,omitempty"`
	PolicyId                *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RetentionType           *string `json:"RetentionType,omitempty" xml:"RetentionType,omitempty"`
	RetentionValue          *string `json:"RetentionValue,omitempty" xml:"RetentionValue,omitempty"`
	SrcRegion               *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	SrcType                 *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	StorageClass            *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetAutoCreated() *bool {
	return s.AutoCreated
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetBakType() *string {
	return s.BakType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetDestType() *string {
	return s.DestType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetDumpAction() *string {
	return s.DumpAction
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetFilterKey() *string {
	return s.FilterKey
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetFilterType() *string {
	return s.FilterType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetFilterValue() *string {
	return s.FilterValue
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetOnlyPreserveOneEachDay() *bool {
	return s.OnlyPreserveOneEachDay
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetOnlyPreserveOneEachHour() *bool {
	return s.OnlyPreserveOneEachHour
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetRetentionType() *string {
	return s.RetentionType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetRetentionValue() *string {
	return s.RetentionValue
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetSrcType() *string {
	return s.SrcType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetAutoCreated(v bool) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.AutoCreated = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetBakType(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.BakType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetDestRegion(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.DestRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetDestType(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.DestType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetDumpAction(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.DumpAction = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetFilterKey(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.FilterKey = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetFilterType(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.FilterType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetFilterValue(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.FilterValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetOnlyPreserveOneEachDay(v bool) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.OnlyPreserveOneEachDay = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetOnlyPreserveOneEachHour(v bool) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.OnlyPreserveOneEachHour = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetPolicyId(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.PolicyId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetRetentionType(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.RetentionType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetRetentionValue(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.RetentionValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetSrcRegion(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.SrcRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetSrcType(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.SrcType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetStorageClass(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.StorageClass = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) Validate() error {
	return dara.Validate(s)
}
