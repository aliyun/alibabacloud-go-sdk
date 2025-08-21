// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceChannelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeDeviceChannelsRequest
	GetId() *string
	SetOwnerId(v int64) *DescribeDeviceChannelsRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *DescribeDeviceChannelsRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeDeviceChannelsRequest
	GetPageSize() *int64
}

type DescribeDeviceChannelsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDeviceChannelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceChannelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceChannelsRequest) GetId() *string {
	return s.Id
}

func (s *DescribeDeviceChannelsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDeviceChannelsRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeDeviceChannelsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDeviceChannelsRequest) SetId(v string) *DescribeDeviceChannelsRequest {
	s.Id = &v
	return s
}

func (s *DescribeDeviceChannelsRequest) SetOwnerId(v int64) *DescribeDeviceChannelsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDeviceChannelsRequest) SetPageNum(v int64) *DescribeDeviceChannelsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeDeviceChannelsRequest) SetPageSize(v int64) *DescribeDeviceChannelsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceChannelsRequest) Validate() error {
	return dara.Validate(s)
}
