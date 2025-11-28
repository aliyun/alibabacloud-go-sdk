// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnFreezeMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberId(v string) *UnFreezeMemberRequest
	GetMemberId() *string
	SetRegionId(v string) *UnFreezeMemberRequest
	GetRegionId() *string
}

type UnFreezeMemberRequest struct {
	// This parameter is required.
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UnFreezeMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UnFreezeMemberRequest) GoString() string {
	return s.String()
}

func (s *UnFreezeMemberRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *UnFreezeMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnFreezeMemberRequest) SetMemberId(v string) *UnFreezeMemberRequest {
	s.MemberId = &v
	return s
}

func (s *UnFreezeMemberRequest) SetRegionId(v string) *UnFreezeMemberRequest {
	s.RegionId = &v
	return s
}

func (s *UnFreezeMemberRequest) Validate() error {
	return dara.Validate(s)
}
