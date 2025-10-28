// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateEcuRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *MigrateEcuRequest
	GetInstanceIds() *string
	SetLogicalRegionId(v string) *MigrateEcuRequest
	GetLogicalRegionId() *string
}

type MigrateEcuRequest struct {
	// The ID of the ECS instance. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2zej4i2jdf3ntwhj****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The ID of the custom namespace.
	//
	// 	- The ID of a custom namespace is in the `region ID:custom namespace ID` format. Example: cn-beijing:tdy218.
	//
	// 	- The ID of the default namespace is in the `region ID` format. Example: cn-beijing.
	//
	// example:
	//
	// cn-hangzhou:test_region
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
}

func (s MigrateEcuRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateEcuRequest) GoString() string {
	return s.String()
}

func (s *MigrateEcuRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *MigrateEcuRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *MigrateEcuRequest) SetInstanceIds(v string) *MigrateEcuRequest {
	s.InstanceIds = &v
	return s
}

func (s *MigrateEcuRequest) SetLogicalRegionId(v string) *MigrateEcuRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *MigrateEcuRequest) Validate() error {
	return dara.Validate(s)
}
