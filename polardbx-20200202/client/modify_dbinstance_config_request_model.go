// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigName(v string) *ModifyDBInstanceConfigRequest
	GetConfigName() *string
	SetConfigValue(v string) *ModifyDBInstanceConfigRequest
	GetConfigValue() *string
	SetDBInstanceName(v string) *ModifyDBInstanceConfigRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ModifyDBInstanceConfigRequest
	GetRegionId() *string
}

type ModifyDBInstanceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ENABLE_CONSISTENT_REPLICA_READ
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *ModifyDBInstanceConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ModifyDBInstanceConfigRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyDBInstanceConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBInstanceConfigRequest) SetConfigName(v string) *ModifyDBInstanceConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetConfigValue(v string) *ModifyDBInstanceConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetDBInstanceName(v string) *ModifyDBInstanceConfigRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetRegionId(v string) *ModifyDBInstanceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
