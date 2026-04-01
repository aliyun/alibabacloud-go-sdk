// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRdsCustomInitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CheckRdsCustomInitRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CheckRdsCustomInitRequest
	GetResourceGroupId() *string
	SetServiceLinkedRole(v string) *CheckRdsCustomInitRequest
	GetServiceLinkedRole() *string
}

type CheckRdsCustomInitRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// None
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AliyunServiceRoleForRds
	ServiceLinkedRole *string `json:"ServiceLinkedRole,omitempty" xml:"ServiceLinkedRole,omitempty"`
}

func (s CheckRdsCustomInitRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckRdsCustomInitRequest) GoString() string {
	return s.String()
}

func (s *CheckRdsCustomInitRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckRdsCustomInitRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CheckRdsCustomInitRequest) GetServiceLinkedRole() *string {
	return s.ServiceLinkedRole
}

func (s *CheckRdsCustomInitRequest) SetRegionId(v string) *CheckRdsCustomInitRequest {
	s.RegionId = &v
	return s
}

func (s *CheckRdsCustomInitRequest) SetResourceGroupId(v string) *CheckRdsCustomInitRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CheckRdsCustomInitRequest) SetServiceLinkedRole(v string) *CheckRdsCustomInitRequest {
	s.ServiceLinkedRole = &v
	return s
}

func (s *CheckRdsCustomInitRequest) Validate() error {
	return dara.Validate(s)
}
