// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeProjectMessagesRequest
	GetInstanceId() *string
	SetPageIndex(v int32) *DescribeProjectMessagesRequest
	GetPageIndex() *int32
}

type DescribeProjectMessagesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4****89
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
}

func (s DescribeProjectMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectMessagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectMessagesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeProjectMessagesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeProjectMessagesRequest) SetInstanceId(v string) *DescribeProjectMessagesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeProjectMessagesRequest) SetPageIndex(v int32) *DescribeProjectMessagesRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeProjectMessagesRequest) Validate() error {
	return dara.Validate(s)
}
