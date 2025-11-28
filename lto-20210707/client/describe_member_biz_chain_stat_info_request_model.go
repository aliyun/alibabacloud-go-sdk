// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberBizChainStatInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeMemberBizChainStatInfoRequest
	GetAccountId() *string
	SetRegionId(v string) *DescribeMemberBizChainStatInfoRequest
	GetRegionId() *string
}

type DescribeMemberBizChainStatInfoRequest struct {
	// This parameter is required.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMemberBizChainStatInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberBizChainStatInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeMemberBizChainStatInfoRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeMemberBizChainStatInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMemberBizChainStatInfoRequest) SetAccountId(v string) *DescribeMemberBizChainStatInfoRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeMemberBizChainStatInfoRequest) SetRegionId(v string) *DescribeMemberBizChainStatInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMemberBizChainStatInfoRequest) Validate() error {
	return dara.Validate(s)
}
