// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRdsCustomInitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *RdsCustomInitRequest
	GetRegionId() *string
	SetServiceLinkedRole(v string) *RdsCustomInitRequest
	GetServiceLinkedRole() *string
}

type RdsCustomInitRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AliyunServiceRoleForRds
	ServiceLinkedRole *string `json:"ServiceLinkedRole,omitempty" xml:"ServiceLinkedRole,omitempty"`
}

func (s RdsCustomInitRequest) String() string {
	return dara.Prettify(s)
}

func (s RdsCustomInitRequest) GoString() string {
	return s.String()
}

func (s *RdsCustomInitRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RdsCustomInitRequest) GetServiceLinkedRole() *string {
	return s.ServiceLinkedRole
}

func (s *RdsCustomInitRequest) SetRegionId(v string) *RdsCustomInitRequest {
	s.RegionId = &v
	return s
}

func (s *RdsCustomInitRequest) SetServiceLinkedRole(v string) *RdsCustomInitRequest {
	s.ServiceLinkedRole = &v
	return s
}

func (s *RdsCustomInitRequest) Validate() error {
	return dara.Validate(s)
}
