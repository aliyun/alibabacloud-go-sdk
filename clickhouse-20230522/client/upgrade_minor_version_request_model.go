// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMinorVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeMinorVersionRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *UpgradeMinorVersionRequest
	GetRegionId() *string
	SetSwitchTime(v string) *UpgradeMinorVersionRequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v string) *UpgradeMinorVersionRequest
	GetSwitchTimeMode() *string
	SetTargetMinorVersion(v string) *UpgradeMinorVersionRequest
	GetTargetMinorVersion() *string
}

type UpgradeMinorVersionRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1jyis8p15we****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The update time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  If you set SwitchTimeMode to SpecifyTime, you must configure this parameter to specify the update time.
	//
	// example:
	//
	// 2023-01-09T05:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// Specifies whether to update the minor engine version of the cluster immediately. Valid values:
	//
	// 	- **Immediate**: The system immediately performs the update.
	//
	// 	- **MaintainTime**: The system performs the update during the specified maintenance window.
	//
	// 	- **SpecifyTime**: The system performs the update at a specified time.
	//
	// example:
	//
	// Immediate
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	// The minor engine version to which you want to update.
	//
	// >  By default, TargetMinorVersion is not set and the minor engine version of the cluster is updated to the latest version.
	//
	// example:
	//
	// 23.8.1.41495_6
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
}

func (s UpgradeMinorVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeMinorVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeMinorVersionRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *UpgradeMinorVersionRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *UpgradeMinorVersionRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *UpgradeMinorVersionRequest) SetDBInstanceId(v string) *UpgradeMinorVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetRegionId(v string) *UpgradeMinorVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetSwitchTime(v string) *UpgradeMinorVersionRequest {
	s.SwitchTime = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetSwitchTimeMode(v string) *UpgradeMinorVersionRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetTargetMinorVersion(v string) *UpgradeMinorVersionRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *UpgradeMinorVersionRequest) Validate() error {
	return dara.Validate(s)
}
