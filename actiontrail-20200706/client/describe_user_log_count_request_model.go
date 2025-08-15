// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserLogCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *DescribeUserLogCountRequest
	GetEndDate() *string
	SetStartDate(v string) *DescribeUserLogCountRequest
	GetStartDate() *string
}

type DescribeUserLogCountRequest struct {
	// example:
	//
	// 2025-06-10
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 2025-05-12
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeUserLogCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLogCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserLogCountRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeUserLogCountRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeUserLogCountRequest) SetEndDate(v string) *DescribeUserLogCountRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUserLogCountRequest) SetStartDate(v string) *DescribeUserLogCountRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeUserLogCountRequest) Validate() error {
	return dara.Validate(s)
}
