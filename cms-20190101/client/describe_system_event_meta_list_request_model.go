// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemEventMetaListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeSystemEventMetaListRequest
	GetRegionId() *string
}

type DescribeSystemEventMetaListRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSystemEventMetaListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventMetaListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventMetaListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSystemEventMetaListRequest) SetRegionId(v string) *DescribeSystemEventMetaListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSystemEventMetaListRequest) Validate() error {
	return dara.Validate(s)
}
