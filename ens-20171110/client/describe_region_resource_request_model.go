// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeRegionResourceRequest
	GetEnsRegionId() *string
	SetIspType(v string) *DescribeRegionResourceRequest
	GetIspType() *string
}

type DescribeRegionResourceRequest struct {
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	IspType     *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
}

func (s DescribeRegionResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeRegionResourceRequest) GetIspType() *string {
	return s.IspType
}

func (s *DescribeRegionResourceRequest) SetEnsRegionId(v string) *DescribeRegionResourceRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeRegionResourceRequest) SetIspType(v string) *DescribeRegionResourceRequest {
	s.IspType = &v
	return s
}

func (s *DescribeRegionResourceRequest) Validate() error {
	return dara.Validate(s)
}
