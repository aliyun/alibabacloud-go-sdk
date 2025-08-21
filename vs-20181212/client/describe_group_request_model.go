// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeGroupRequest
	GetId() *string
	SetIncludeStats(v bool) *DescribeGroupRequest
	GetIncludeStats() *bool
	SetOwnerId(v int64) *DescribeGroupRequest
	GetOwnerId() *int64
}

type DescribeGroupRequest struct {
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IncludeStats *bool  `json:"IncludeStats,omitempty" xml:"IncludeStats,omitempty"`
	OwnerId      *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupRequest) GetId() *string {
	return s.Id
}

func (s *DescribeGroupRequest) GetIncludeStats() *bool {
	return s.IncludeStats
}

func (s *DescribeGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGroupRequest) SetId(v string) *DescribeGroupRequest {
	s.Id = &v
	return s
}

func (s *DescribeGroupRequest) SetIncludeStats(v bool) *DescribeGroupRequest {
	s.IncludeStats = &v
	return s
}

func (s *DescribeGroupRequest) SetOwnerId(v int64) *DescribeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGroupRequest) Validate() error {
	return dara.Validate(s)
}
