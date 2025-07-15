// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBusinessAccessPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListBusinessAccessPointsRequest
	GetRegionId() *string
}

type ListBusinessAccessPointsRequest struct {
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBusinessAccessPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessAccessPointsRequest) GoString() string {
	return s.String()
}

func (s *ListBusinessAccessPointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBusinessAccessPointsRequest) SetRegionId(v string) *ListBusinessAccessPointsRequest {
	s.RegionId = &v
	return s
}

func (s *ListBusinessAccessPointsRequest) Validate() error {
	return dara.Validate(s)
}
