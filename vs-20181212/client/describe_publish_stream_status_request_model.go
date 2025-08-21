// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePublishStreamStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribePublishStreamStatusRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *DescribePublishStreamStatusRequest
	GetOwnerId() *int64
}

type DescribePublishStreamStatusRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribePublishStreamStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePublishStreamStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribePublishStreamStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePublishStreamStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePublishStreamStatusRequest) SetInstanceId(v string) *DescribePublishStreamStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePublishStreamStatusRequest) SetOwnerId(v int64) *DescribePublishStreamStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePublishStreamStatusRequest) Validate() error {
	return dara.Validate(s)
}
