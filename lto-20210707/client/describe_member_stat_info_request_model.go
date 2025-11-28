// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberStatInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizChainId(v string) *DescribeMemberStatInfoRequest
	GetBizChainId() *string
	SetRegionId(v string) *DescribeMemberStatInfoRequest
	GetRegionId() *string
}

type DescribeMemberStatInfoRequest struct {
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMemberStatInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberStatInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeMemberStatInfoRequest) GetBizChainId() *string {
	return s.BizChainId
}

func (s *DescribeMemberStatInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMemberStatInfoRequest) SetBizChainId(v string) *DescribeMemberStatInfoRequest {
	s.BizChainId = &v
	return s
}

func (s *DescribeMemberStatInfoRequest) SetRegionId(v string) *DescribeMemberStatInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMemberStatInfoRequest) Validate() error {
	return dara.Validate(s)
}
