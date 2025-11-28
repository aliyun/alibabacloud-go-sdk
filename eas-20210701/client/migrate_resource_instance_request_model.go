// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateResourceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestResourceId(v string) *MigrateResourceInstanceRequest
	GetDestResourceId() *string
	SetInstanceIds(v []*string) *MigrateResourceInstanceRequest
	GetInstanceIds() []*string
	SetMigrateToHybrid(v bool) *MigrateResourceInstanceRequest
	GetMigrateToHybrid() *bool
}

type MigrateResourceInstanceRequest struct {
	// The ID of the destination resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// eas-r-asdasdasd****
	DestResourceId *string `json:"DestResourceId,omitempty" xml:"DestResourceId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	MigrateToHybrid *bool     `json:"MigrateToHybrid,omitempty" xml:"MigrateToHybrid,omitempty"`
}

func (s MigrateResourceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateResourceInstanceRequest) GoString() string {
	return s.String()
}

func (s *MigrateResourceInstanceRequest) GetDestResourceId() *string {
	return s.DestResourceId
}

func (s *MigrateResourceInstanceRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *MigrateResourceInstanceRequest) GetMigrateToHybrid() *bool {
	return s.MigrateToHybrid
}

func (s *MigrateResourceInstanceRequest) SetDestResourceId(v string) *MigrateResourceInstanceRequest {
	s.DestResourceId = &v
	return s
}

func (s *MigrateResourceInstanceRequest) SetInstanceIds(v []*string) *MigrateResourceInstanceRequest {
	s.InstanceIds = v
	return s
}

func (s *MigrateResourceInstanceRequest) SetMigrateToHybrid(v bool) *MigrateResourceInstanceRequest {
	s.MigrateToHybrid = &v
	return s
}

func (s *MigrateResourceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
