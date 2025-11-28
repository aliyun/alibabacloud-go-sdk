// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFreezeMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberId(v string) *FreezeMemberRequest
	GetMemberId() *string
	SetRegionId(v string) *FreezeMemberRequest
	GetRegionId() *string
}

type FreezeMemberRequest struct {
	// This parameter is required.
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s FreezeMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s FreezeMemberRequest) GoString() string {
	return s.String()
}

func (s *FreezeMemberRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *FreezeMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *FreezeMemberRequest) SetMemberId(v string) *FreezeMemberRequest {
	s.MemberId = &v
	return s
}

func (s *FreezeMemberRequest) SetRegionId(v string) *FreezeMemberRequest {
	s.RegionId = &v
	return s
}

func (s *FreezeMemberRequest) Validate() error {
	return dara.Validate(s)
}
