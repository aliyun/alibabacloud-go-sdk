// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectAttachmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeProjectAttachmentsRequest
	GetInstanceId() *string
}

type DescribeProjectAttachmentsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4****89
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeProjectAttachmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectAttachmentsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeProjectAttachmentsRequest) SetInstanceId(v string) *DescribeProjectAttachmentsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeProjectAttachmentsRequest) Validate() error {
	return dara.Validate(s)
}
