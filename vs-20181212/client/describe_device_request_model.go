// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeDeviceRequest
	GetId() *string
	SetIncludeDirectory(v bool) *DescribeDeviceRequest
	GetIncludeDirectory() *bool
	SetIncludeStats(v bool) *DescribeDeviceRequest
	GetIncludeStats() *bool
	SetOwnerId(v int64) *DescribeDeviceRequest
	GetOwnerId() *int64
}

type DescribeDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3238848****092996-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IncludeDirectory *bool `json:"IncludeDirectory,omitempty" xml:"IncludeDirectory,omitempty"`
	// example:
	//
	// false
	IncludeStats *bool  `json:"IncludeStats,omitempty" xml:"IncludeStats,omitempty"`
	OwnerId      *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceRequest) GetId() *string {
	return s.Id
}

func (s *DescribeDeviceRequest) GetIncludeDirectory() *bool {
	return s.IncludeDirectory
}

func (s *DescribeDeviceRequest) GetIncludeStats() *bool {
	return s.IncludeStats
}

func (s *DescribeDeviceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDeviceRequest) SetId(v string) *DescribeDeviceRequest {
	s.Id = &v
	return s
}

func (s *DescribeDeviceRequest) SetIncludeDirectory(v bool) *DescribeDeviceRequest {
	s.IncludeDirectory = &v
	return s
}

func (s *DescribeDeviceRequest) SetIncludeStats(v bool) *DescribeDeviceRequest {
	s.IncludeStats = &v
	return s
}

func (s *DescribeDeviceRequest) SetOwnerId(v int64) *DescribeDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDeviceRequest) Validate() error {
	return dara.Validate(s)
}
