// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParentPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeParentPlatformRequest
	GetId() *string
	SetOwnerId(v int64) *DescribeParentPlatformRequest
	GetOwnerId() *int64
}

type DescribeParentPlatformRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 359*****374-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeParentPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentPlatformRequest) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformRequest) GetId() *string {
	return s.Id
}

func (s *DescribeParentPlatformRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParentPlatformRequest) SetId(v string) *DescribeParentPlatformRequest {
	s.Id = &v
	return s
}

func (s *DescribeParentPlatformRequest) SetOwnerId(v int64) *DescribeParentPlatformRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParentPlatformRequest) Validate() error {
	return dara.Validate(s)
}
