// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *ModifyAppInstanceShrinkRequest
	GetBranchName() *string
	SetClientToken(v string) *ModifyAppInstanceShrinkRequest
	GetClientToken() *string
	SetComponentsShrink(v string) *ModifyAppInstanceShrinkRequest
	GetComponentsShrink() *string
	SetDBInstanceName(v string) *ModifyAppInstanceShrinkRequest
	GetDBInstanceName() *string
	SetInstanceClass(v string) *ModifyAppInstanceShrinkRequest
	GetInstanceClass() *string
	SetInstanceName(v string) *ModifyAppInstanceShrinkRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyAppInstanceShrinkRequest
	GetRegionId() *string
}

type ModifyAppInstanceShrinkRequest struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The idempotency parameter.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of modules.
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	DBInstanceName   *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	InstanceClass    *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The instance ID of the AI application.
	//
	// example:
	//
	// ra-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAppInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceShrinkRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *ModifyAppInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAppInstanceShrinkRequest) GetComponentsShrink() *string {
	return s.ComponentsShrink
}

func (s *ModifyAppInstanceShrinkRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyAppInstanceShrinkRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *ModifyAppInstanceShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyAppInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAppInstanceShrinkRequest) SetBranchName(v string) *ModifyAppInstanceShrinkRequest {
	s.BranchName = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetClientToken(v string) *ModifyAppInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetComponentsShrink(v string) *ModifyAppInstanceShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetDBInstanceName(v string) *ModifyAppInstanceShrinkRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetInstanceClass(v string) *ModifyAppInstanceShrinkRequest {
	s.InstanceClass = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetInstanceName(v string) *ModifyAppInstanceShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetRegionId(v string) *ModifyAppInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
