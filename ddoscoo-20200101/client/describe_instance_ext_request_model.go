// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceExtRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceExtRequest
	GetInstanceId() *string
	SetPageNumber(v string) *DescribeInstanceExtRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeInstanceExtRequest
	GetPageSize() *string
}

type DescribeInstanceExtRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// example:
	//
	// ddoscoo-cn-i7m25564****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page. For example, to query the returned results on the first page, set the value to **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceExtRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceExtRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceExtRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceExtRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeInstanceExtRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInstanceExtRequest) SetInstanceId(v string) *DescribeInstanceExtRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceExtRequest) SetPageNumber(v string) *DescribeInstanceExtRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceExtRequest) SetPageSize(v string) *DescribeInstanceExtRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceExtRequest) Validate() error {
	return dara.Validate(s)
}
