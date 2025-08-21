// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDiskTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeCloudDiskTypesRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v []*string) *DescribeCloudDiskTypesRequest
	GetEnsRegionIds() []*string
}

type DescribeCloudDiskTypesRequest struct {
	// The ID of the edge node.
	//
	// example:
	//
	// cn-chongqing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The edge nodes.
	EnsRegionIds []*string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Repeated"`
}

func (s DescribeCloudDiskTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskTypesRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeCloudDiskTypesRequest) GetEnsRegionIds() []*string {
	return s.EnsRegionIds
}

func (s *DescribeCloudDiskTypesRequest) SetEnsRegionId(v string) *DescribeCloudDiskTypesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeCloudDiskTypesRequest) SetEnsRegionIds(v []*string) *DescribeCloudDiskTypesRequest {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeCloudDiskTypesRequest) Validate() error {
	return dara.Validate(s)
}
