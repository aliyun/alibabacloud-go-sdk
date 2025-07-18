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
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-wz9kmr708m155j***
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is no longer used and does not need to be specified.
	//
	// example:
	//
	// null
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// The minor version of the instance.
	//
	// example:
	//
	// 6.3.6.1-202112012048
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is no longer used and does not need to be specified.
	//
	// example:
	//
	// null
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// This parameter is no longer used and does not need to be specified.
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
