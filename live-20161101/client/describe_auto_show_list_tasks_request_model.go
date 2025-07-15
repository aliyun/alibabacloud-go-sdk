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
	// The ID of the production studio for which you want to query scheduled tasks.
	//
	// >  If the scheduled tasks were created by calling the [InitializeAutoShowListTask](https://help.aliyun.com/document_detail/2848056.html) operation, check the value of the response parameter CasterId to obtain the ID. If you do not specify this parameter, the system queries all the scheduled tasks that belong to the user specified by the OwnerId parameter.
	//
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
