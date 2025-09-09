// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstanceRdMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddInstanceRdMemberRequest
	GetInstanceId() *string
	SetMemberId(v string) *AddInstanceRdMemberRequest
	GetMemberId() *string
	SetRegionId(v string) *AddInstanceRdMemberRequest
	GetRegionId() *string
}

type AddInstanceRdMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-7mz2ve7h00a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1857311509574932
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddInstanceRdMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceRdMemberRequest) GoString() string {
	return s.String()
}

func (s *AddInstanceRdMemberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddInstanceRdMemberRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *AddInstanceRdMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddInstanceRdMemberRequest) SetInstanceId(v string) *AddInstanceRdMemberRequest {
	s.InstanceId = &v
	return s
}

func (s *AddInstanceRdMemberRequest) SetMemberId(v string) *AddInstanceRdMemberRequest {
	s.MemberId = &v
	return s
}

func (s *AddInstanceRdMemberRequest) SetRegionId(v string) *AddInstanceRdMemberRequest {
	s.RegionId = &v
	return s
}

func (s *AddInstanceRdMemberRequest) Validate() error {
	return dara.Validate(s)
}
