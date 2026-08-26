// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoShowListTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeAutoShowListTasksRequest
	GetCasterId() *string
	SetOwnerId(v int64) *DescribeAutoShowListTasksRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAutoShowListTasksRequest
	GetRegionId() *string
}

type DescribeAutoShowListTasksRequest struct {
	// The ID of the production studio whose scheduled tasks you want to query.
	//
	// > If you create a scheduled task by calling the [InitializeAutoShowListTask](https://help.aliyun.com/document_detail/2848056.html) operation, use the CasterId value returned by that operation. If you leave this parameter empty, all scheduled tasks under your account are queried by default.
	//
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAutoShowListTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoShowListTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoShowListTasksRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeAutoShowListTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAutoShowListTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoShowListTasksRequest) SetCasterId(v string) *DescribeAutoShowListTasksRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeAutoShowListTasksRequest) SetOwnerId(v int64) *DescribeAutoShowListTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoShowListTasksRequest) SetRegionId(v string) *DescribeAutoShowListTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoShowListTasksRequest) Validate() error {
	return dara.Validate(s)
}
