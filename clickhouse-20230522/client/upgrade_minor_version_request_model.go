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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The specified upgrade time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// > This parameter is required when SwitchTimeMode is set to SpecifyTime.
	//
	// example:
	//
	// 2023-01-09T05:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// Specifies when to upgrade. Valid values:
	//
	// - **Immediate**: upgrades immediately.
	//
	// - **MaintainTime**: upgrades during the O&M window.
	//
	// - **SpecifyTime**: upgrades at a specified time.
	//
	// example:
	//
	// Immediate
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	// The target minor engine version.
	//
	// >By default, leave this parameter empty to upgrade to the latest minor engine version.
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
