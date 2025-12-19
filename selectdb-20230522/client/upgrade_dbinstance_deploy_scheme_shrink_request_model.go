// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceDeploySchemeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeDBInstanceDeploySchemeShrinkRequest
	GetDBInstanceId() *string
	SetMultiZoneShrink(v string) *UpgradeDBInstanceDeploySchemeShrinkRequest
	GetMultiZoneShrink() *string
	SetRegionId(v string) *UpgradeDBInstanceDeploySchemeShrinkRequest
	GetRegionId() *string
	SetSecurityToken(v string) *UpgradeDBInstanceDeploySchemeShrinkRequest
	GetSecurityToken() *string
}

type UpgradeDBInstanceDeploySchemeShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	MultiZoneShrink *string `json:"MultiZone,omitempty" xml:"MultiZone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpgradeDBInstanceDeploySchemeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceDeploySchemeShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceDeploySchemeShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBInstanceDeploySchemeShrinkRequest) GetMultiZoneShrink() *string {
	return s.MultiZoneShrink
}

func (s *UpgradeDBInstanceDeploySchemeShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeDBInstanceDeploySchemeShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpgradeDBInstanceDeploySchemeShrinkRequest) SetDBInstanceId(v string) *UpgradeDBInstanceDeploySchemeShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeShrinkRequest) SetMultiZoneShrink(v string) *UpgradeDBInstanceDeploySchemeShrinkRequest {
	s.MultiZoneShrink = &v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeShrinkRequest) SetRegionId(v string) *UpgradeDBInstanceDeploySchemeShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeShrinkRequest) SetSecurityToken(v string) *UpgradeDBInstanceDeploySchemeShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
