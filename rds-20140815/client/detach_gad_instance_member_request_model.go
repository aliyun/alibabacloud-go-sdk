// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGadInstanceMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGadInstanceName(v string) *DetachGadInstanceMemberRequest
	GetGadInstanceName() *string
	SetMemberInstanceName(v string) *DetachGadInstanceMemberRequest
	GetMemberInstanceName() *string
	SetRegionId(v string) *DetachGadInstanceMemberRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DetachGadInstanceMemberRequest
	GetResourceGroupId() *string
}

type DetachGadInstanceMemberRequest struct {
	// This parameter is required.
	GadInstanceName *string `json:"GadInstanceName,omitempty" xml:"GadInstanceName,omitempty"`
	// This parameter is required.
	MemberInstanceName *string `json:"MemberInstanceName,omitempty" xml:"MemberInstanceName,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId    *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DetachGadInstanceMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachGadInstanceMemberRequest) GoString() string {
	return s.String()
}

func (s *DetachGadInstanceMemberRequest) GetGadInstanceName() *string {
	return s.GadInstanceName
}

func (s *DetachGadInstanceMemberRequest) GetMemberInstanceName() *string {
	return s.MemberInstanceName
}

func (s *DetachGadInstanceMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachGadInstanceMemberRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DetachGadInstanceMemberRequest) SetGadInstanceName(v string) *DetachGadInstanceMemberRequest {
	s.GadInstanceName = &v
	return s
}

func (s *DetachGadInstanceMemberRequest) SetMemberInstanceName(v string) *DetachGadInstanceMemberRequest {
	s.MemberInstanceName = &v
	return s
}

func (s *DetachGadInstanceMemberRequest) SetRegionId(v string) *DetachGadInstanceMemberRequest {
	s.RegionId = &v
	return s
}

func (s *DetachGadInstanceMemberRequest) SetResourceGroupId(v string) *DetachGadInstanceMemberRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DetachGadInstanceMemberRequest) Validate() error {
	return dara.Validate(s)
}
