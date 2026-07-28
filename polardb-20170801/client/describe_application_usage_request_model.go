// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationUsageRequest
	GetApplicationId() *string
	SetDays(v int32) *DescribeApplicationUsageRequest
	GetDays() *int32
}

type DescribeApplicationUsageRequest struct {
	// The Hermes application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-123456
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The number of recent days to query. Valid values: 1 to 365. Default value: 30.
	//
	// example:
	//
	// 30
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
}

func (s DescribeApplicationUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationUsageRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationUsageRequest) GetDays() *int32 {
	return s.Days
}

func (s *DescribeApplicationUsageRequest) SetApplicationId(v string) *DescribeApplicationUsageRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationUsageRequest) SetDays(v int32) *DescribeApplicationUsageRequest {
	s.Days = &v
	return s
}

func (s *DescribeApplicationUsageRequest) Validate() error {
	return dara.Validate(s)
}
