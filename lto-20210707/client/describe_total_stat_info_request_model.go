// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTotalStatInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeTotalStatInfoRequest
	GetRegionId() *string
}

type DescribeTotalStatInfoRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTotalStatInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTotalStatInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeTotalStatInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTotalStatInfoRequest) SetRegionId(v string) *DescribeTotalStatInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTotalStatInfoRequest) Validate() error {
	return dara.Validate(s)
}
