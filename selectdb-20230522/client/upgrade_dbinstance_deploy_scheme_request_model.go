// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceDeploySchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeDBInstanceDeploySchemeRequest
	GetDBInstanceId() *string
	SetMultiZone(v []*UpgradeDBInstanceDeploySchemeRequestMultiZone) *UpgradeDBInstanceDeploySchemeRequest
	GetMultiZone() []*UpgradeDBInstanceDeploySchemeRequestMultiZone
	SetRegionId(v string) *UpgradeDBInstanceDeploySchemeRequest
	GetRegionId() *string
	SetSecurityToken(v string) *UpgradeDBInstanceDeploySchemeRequest
	GetSecurityToken() *string
}

type UpgradeDBInstanceDeploySchemeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	MultiZone []*UpgradeDBInstanceDeploySchemeRequestMultiZone `json:"MultiZone,omitempty" xml:"MultiZone,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpgradeDBInstanceDeploySchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceDeploySchemeRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceDeploySchemeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBInstanceDeploySchemeRequest) GetMultiZone() []*UpgradeDBInstanceDeploySchemeRequestMultiZone {
	return s.MultiZone
}

func (s *UpgradeDBInstanceDeploySchemeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeDBInstanceDeploySchemeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpgradeDBInstanceDeploySchemeRequest) SetDBInstanceId(v string) *UpgradeDBInstanceDeploySchemeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeRequest) SetMultiZone(v []*UpgradeDBInstanceDeploySchemeRequestMultiZone) *UpgradeDBInstanceDeploySchemeRequest {
	s.MultiZone = v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeRequest) SetRegionId(v string) *UpgradeDBInstanceDeploySchemeRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeRequest) SetSecurityToken(v string) *UpgradeDBInstanceDeploySchemeRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeRequest) Validate() error {
	if s.MultiZone != nil {
		for _, item := range s.MultiZone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpgradeDBInstanceDeploySchemeRequestMultiZone struct {
	// This parameter is required.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpgradeDBInstanceDeploySchemeRequestMultiZone) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceDeploySchemeRequestMultiZone) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceDeploySchemeRequestMultiZone) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *UpgradeDBInstanceDeploySchemeRequestMultiZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpgradeDBInstanceDeploySchemeRequestMultiZone) SetVSwitchIds(v []*string) *UpgradeDBInstanceDeploySchemeRequestMultiZone {
	s.VSwitchIds = v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeRequestMultiZone) SetZoneId(v string) *UpgradeDBInstanceDeploySchemeRequestMultiZone {
	s.ZoneId = &v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeRequestMultiZone) Validate() error {
	return dara.Validate(s)
}
