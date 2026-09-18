// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeDBVersionRequest
	GetDBInstanceId() *string
	SetEffectiveTime(v string) *UpgradeDBVersionRequest
	GetEffectiveTime() *string
	SetMajorVersion(v string) *UpgradeDBVersionRequest
	GetMajorVersion() *string
	SetMinorVersion(v string) *UpgradeDBVersionRequest
	GetMinorVersion() *string
	SetOwnerId(v int64) *UpgradeDBVersionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpgradeDBVersionRequest
	GetRegionId() *string
	SetSwitchTime(v string) *UpgradeDBVersionRequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v string) *UpgradeDBVersionRequest
	GetSwitchTimeMode() *string
}

type UpgradeDBVersionRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in the specified region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-wz9kmr708m155j***
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The effective period. Valid values:
	//
	// 	- **Immediate*	- (default): The upgrade takes effect immediately.
	//
	// 	- **MaintainTime**: The upgrade takes effect during the O&M window. For more information, see ModifyDBInstanceMaintainTime.
	//
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. You do not need to specify this parameter.
	//
	// example:
	//
	// null
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// The minor version.
	//
	// example:
	//
	// 6.3.6.1-202112012048
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. You do not need to specify this parameter.
	//
	// example:
	//
	// null
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. You do not need to specify this parameter.
	//
	// example:
	//
	// null
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s UpgradeDBVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBVersionRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *UpgradeDBVersionRequest) GetMajorVersion() *string {
	return s.MajorVersion
}

func (s *UpgradeDBVersionRequest) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *UpgradeDBVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeDBVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeDBVersionRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *UpgradeDBVersionRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *UpgradeDBVersionRequest) SetDBInstanceId(v string) *UpgradeDBVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetEffectiveTime(v string) *UpgradeDBVersionRequest {
	s.EffectiveTime = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetMajorVersion(v string) *UpgradeDBVersionRequest {
	s.MajorVersion = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetMinorVersion(v string) *UpgradeDBVersionRequest {
	s.MinorVersion = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetOwnerId(v int64) *UpgradeDBVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetRegionId(v string) *UpgradeDBVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetSwitchTime(v string) *UpgradeDBVersionRequest {
	s.SwitchTime = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetSwitchTimeMode(v string) *UpgradeDBVersionRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *UpgradeDBVersionRequest) Validate() error {
	return dara.Validate(s)
}
