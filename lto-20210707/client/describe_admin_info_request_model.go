// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdminInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeAdminInfoRequest
	GetRegionId() *string
}

type DescribeAdminInfoRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAdminInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdminInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdminInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAdminInfoRequest) SetRegionId(v string) *DescribeAdminInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAdminInfoRequest) Validate() error {
	return dara.Validate(s)
}
