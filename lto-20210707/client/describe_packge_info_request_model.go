// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackgeInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribePackgeInfoRequest
	GetAccountId() *string
	SetRegionId(v string) *DescribePackgeInfoRequest
	GetRegionId() *string
}

type DescribePackgeInfoRequest struct {
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePackgeInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePackgeInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribePackgeInfoRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribePackgeInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePackgeInfoRequest) SetAccountId(v string) *DescribePackgeInfoRequest {
	s.AccountId = &v
	return s
}

func (s *DescribePackgeInfoRequest) SetRegionId(v string) *DescribePackgeInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePackgeInfoRequest) Validate() error {
	return dara.Validate(s)
}
