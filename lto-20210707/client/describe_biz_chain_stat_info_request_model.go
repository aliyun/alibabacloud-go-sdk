// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBizChainStatInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeBizChainStatInfoRequest
	GetRegionId() *string
}

type DescribeBizChainStatInfoRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBizChainStatInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBizChainStatInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeBizChainStatInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBizChainStatInfoRequest) SetRegionId(v string) *DescribeBizChainStatInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBizChainStatInfoRequest) Validate() error {
	return dara.Validate(s)
}
