// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest
	GetDBInstanceDescription() *string
	SetDBInstanceName(v string) *ModifyDBInstanceDescriptionRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ModifyDBInstanceDescriptionRequest
	GetRegionId() *string
}

type ModifyDBInstanceDescriptionRequest struct {
	// This parameter is required.
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
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

func (s ModifyDBInstanceDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *ModifyDBInstanceDescriptionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyDBInstanceDescriptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceName(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetRegionId(v string) *ModifyDBInstanceDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
