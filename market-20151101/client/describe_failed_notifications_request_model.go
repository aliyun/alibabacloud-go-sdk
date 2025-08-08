// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFailedNotificationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeFailedNotificationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFailedNotificationsRequest
	GetPageSize() *int32
}

type DescribeFailedNotificationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeFailedNotificationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFailedNotificationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFailedNotificationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFailedNotificationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFailedNotificationsRequest) SetPageNumber(v int32) *DescribeFailedNotificationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFailedNotificationsRequest) SetPageSize(v int32) *DescribeFailedNotificationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFailedNotificationsRequest) Validate() error {
	return dara.Validate(s)
}
