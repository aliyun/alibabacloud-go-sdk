// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemberAuthorizedBizChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberId(v string) *ListMemberAuthorizedBizChainRequest
	GetMemberId() *string
	SetRegionId(v string) *ListMemberAuthorizedBizChainRequest
	GetRegionId() *string
}

type ListMemberAuthorizedBizChainRequest struct {
	// This parameter is required.
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListMemberAuthorizedBizChainRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemberAuthorizedBizChainRequest) GoString() string {
	return s.String()
}

func (s *ListMemberAuthorizedBizChainRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *ListMemberAuthorizedBizChainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListMemberAuthorizedBizChainRequest) SetMemberId(v string) *ListMemberAuthorizedBizChainRequest {
	s.MemberId = &v
	return s
}

func (s *ListMemberAuthorizedBizChainRequest) SetRegionId(v string) *ListMemberAuthorizedBizChainRequest {
	s.RegionId = &v
	return s
}

func (s *ListMemberAuthorizedBizChainRequest) Validate() error {
	return dara.Validate(s)
}
