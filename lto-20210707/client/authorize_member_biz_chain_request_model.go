// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeMemberBizChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizChainInfo(v string) *AuthorizeMemberBizChainRequest
	GetBizChainInfo() *string
	SetMemberId(v string) *AuthorizeMemberBizChainRequest
	GetMemberId() *string
	SetRegionId(v string) *AuthorizeMemberBizChainRequest
	GetRegionId() *string
}

type AuthorizeMemberBizChainRequest struct {
	// This parameter is required.
	BizChainInfo *string `json:"BizChainInfo,omitempty" xml:"BizChainInfo,omitempty"`
	// This parameter is required.
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AuthorizeMemberBizChainRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeMemberBizChainRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeMemberBizChainRequest) GetBizChainInfo() *string {
	return s.BizChainInfo
}

func (s *AuthorizeMemberBizChainRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *AuthorizeMemberBizChainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AuthorizeMemberBizChainRequest) SetBizChainInfo(v string) *AuthorizeMemberBizChainRequest {
	s.BizChainInfo = &v
	return s
}

func (s *AuthorizeMemberBizChainRequest) SetMemberId(v string) *AuthorizeMemberBizChainRequest {
	s.MemberId = &v
	return s
}

func (s *AuthorizeMemberBizChainRequest) SetRegionId(v string) *AuthorizeMemberBizChainRequest {
	s.RegionId = &v
	return s
}

func (s *AuthorizeMemberBizChainRequest) Validate() error {
	return dara.Validate(s)
}
