// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeTagsRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeTagsRequest
	GetRegionId() *string
	SetTagKey(v string) *DescribeTagsRequest
	GetTagKey() *string
}

type DescribeTagsRequest struct {
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagsRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagsRequest) SetDBInstanceName(v string) *DescribeTagsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeTagsRequest) SetRegionId(v string) *DescribeTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagsRequest) SetTagKey(v string) *DescribeTagsRequest {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsRequest) Validate() error {
	return dara.Validate(s)
}
