// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeniedMemberAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberAccountId(v string) *DeniedMemberAccessRequest
	GetMemberAccountId() *string
	SetRegionId(v string) *DeniedMemberAccessRequest
	GetRegionId() *string
}

type DeniedMemberAccessRequest struct {
	// This parameter is required.
	MemberAccountId *string `json:"MemberAccountId,omitempty" xml:"MemberAccountId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeniedMemberAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s DeniedMemberAccessRequest) GoString() string {
	return s.String()
}

func (s *DeniedMemberAccessRequest) GetMemberAccountId() *string {
	return s.MemberAccountId
}

func (s *DeniedMemberAccessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeniedMemberAccessRequest) SetMemberAccountId(v string) *DeniedMemberAccessRequest {
	s.MemberAccountId = &v
	return s
}

func (s *DeniedMemberAccessRequest) SetRegionId(v string) *DeniedMemberAccessRequest {
	s.RegionId = &v
	return s
}

func (s *DeniedMemberAccessRequest) Validate() error {
	return dara.Validate(s)
}
