// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgreeMemberAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberAccountId(v string) *AgreeMemberAccessRequest
	GetMemberAccountId() *string
	SetRegionId(v string) *AgreeMemberAccessRequest
	GetRegionId() *string
}

type AgreeMemberAccessRequest struct {
	// This parameter is required.
	MemberAccountId *string `json:"MemberAccountId,omitempty" xml:"MemberAccountId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AgreeMemberAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s AgreeMemberAccessRequest) GoString() string {
	return s.String()
}

func (s *AgreeMemberAccessRequest) GetMemberAccountId() *string {
	return s.MemberAccountId
}

func (s *AgreeMemberAccessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AgreeMemberAccessRequest) SetMemberAccountId(v string) *AgreeMemberAccessRequest {
	s.MemberAccountId = &v
	return s
}

func (s *AgreeMemberAccessRequest) SetRegionId(v string) *AgreeMemberAccessRequest {
	s.RegionId = &v
	return s
}

func (s *AgreeMemberAccessRequest) Validate() error {
	return dara.Validate(s)
}
