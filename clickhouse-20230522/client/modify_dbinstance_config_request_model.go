// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceConfigRequest
	GetDBInstanceId() *string
	SetParameters(v string) *ModifyDBInstanceConfigRequest
	GetParameters() *string
	SetRegionId(v string) *ModifyDBInstanceConfigRequest
	GetRegionId() *string
}

type ModifyDBInstanceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-uf6lkzf*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// {"max_concurrent_queries":"100"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s *ModifyDBInstanceConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceConfigRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyDBInstanceConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBInstanceConfigRequest) SetDBInstanceId(v string) *ModifyDBInstanceConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetParameters(v string) *ModifyDBInstanceConfigRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetRegionId(v string) *ModifyDBInstanceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
