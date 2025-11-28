// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberTotalStatInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeMemberTotalStatInfoRequest
	GetAccountId() *string
	SetRegionId(v string) *DescribeMemberTotalStatInfoRequest
	GetRegionId() *string
}

type DescribeMemberTotalStatInfoRequest struct {
	// This parameter is required.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMemberTotalStatInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberTotalStatInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeMemberTotalStatInfoRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeMemberTotalStatInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMemberTotalStatInfoRequest) SetAccountId(v string) *DescribeMemberTotalStatInfoRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeMemberTotalStatInfoRequest) SetRegionId(v string) *DescribeMemberTotalStatInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMemberTotalStatInfoRequest) Validate() error {
	return dara.Validate(s)
}
