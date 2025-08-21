// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeDirectoryRequest
	GetId() *string
	SetOwnerId(v int64) *DescribeDirectoryRequest
	GetOwnerId() *int64
}

type DescribeDirectoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 399*****488-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoryRequest) GetId() *string {
	return s.Id
}

func (s *DescribeDirectoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDirectoryRequest) SetId(v string) *DescribeDirectoryRequest {
	s.Id = &v
	return s
}

func (s *DescribeDirectoryRequest) SetOwnerId(v int64) *DescribeDirectoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
