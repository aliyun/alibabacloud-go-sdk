// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKmsKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeKmsKeysRequest
	GetRegionId() *string
}

type DescribeKmsKeysRequest struct {
	// The ID of the region. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeKmsKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeKmsKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKmsKeysRequest) SetRegionId(v string) *DescribeKmsKeysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKmsKeysRequest) Validate() error {
	return dara.Validate(s)
}
