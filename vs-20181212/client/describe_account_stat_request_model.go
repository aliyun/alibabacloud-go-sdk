// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeAccountStatRequest
	GetId() *string
	SetOwnerId(v int64) *DescribeAccountStatRequest
	GetOwnerId() *int64
}

type DescribeAccountStatRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeAccountStatRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountStatRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountStatRequest) GetId() *string {
	return s.Id
}

func (s *DescribeAccountStatRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAccountStatRequest) SetId(v string) *DescribeAccountStatRequest {
	s.Id = &v
	return s
}

func (s *DescribeAccountStatRequest) SetOwnerId(v int64) *DescribeAccountStatRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccountStatRequest) Validate() error {
	return dara.Validate(s)
}
