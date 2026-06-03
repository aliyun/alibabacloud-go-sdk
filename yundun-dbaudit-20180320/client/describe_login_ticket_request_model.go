// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoginTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeLoginTicketRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeLoginTicketRequest
	GetRegionId() *string
}

type DescribeLoginTicketRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLoginTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoginTicketRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoginTicketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLoginTicketRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoginTicketRequest) SetInstanceId(v string) *DescribeLoginTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeLoginTicketRequest) SetRegionId(v string) *DescribeLoginTicketRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoginTicketRequest) Validate() error {
	return dara.Validate(s)
}
