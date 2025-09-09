// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstanceRdMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemoveInstanceRdMemberRequest
	GetInstanceId() *string
	SetMemberId(v string) *RemoveInstanceRdMemberRequest
	GetMemberId() *string
	SetRegionId(v string) *RemoveInstanceRdMemberRequest
	GetRegionId() *string
}

type RemoveInstanceRdMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1597141696147832
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveInstanceRdMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceRdMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveInstanceRdMemberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveInstanceRdMemberRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *RemoveInstanceRdMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveInstanceRdMemberRequest) SetInstanceId(v string) *RemoveInstanceRdMemberRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveInstanceRdMemberRequest) SetMemberId(v string) *RemoveInstanceRdMemberRequest {
	s.MemberId = &v
	return s
}

func (s *RemoveInstanceRdMemberRequest) SetRegionId(v string) *RemoveInstanceRdMemberRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveInstanceRdMemberRequest) Validate() error {
	return dara.Validate(s)
}
