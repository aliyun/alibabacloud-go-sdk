// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeStreamRequest
	GetId() *string
	SetOwnerId(v int64) *DescribeStreamRequest
	GetOwnerId() *int64
}

type DescribeStreamRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 323*****997-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamRequest) GetId() *string {
	return s.Id
}

func (s *DescribeStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeStreamRequest) SetId(v string) *DescribeStreamRequest {
	s.Id = &v
	return s
}

func (s *DescribeStreamRequest) SetOwnerId(v int64) *DescribeStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStreamRequest) Validate() error {
	return dara.Validate(s)
}
