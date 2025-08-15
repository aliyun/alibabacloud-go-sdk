// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAlertCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *DescribeUserAlertCountRequest
	GetEndDate() *string
	SetStartDate(v string) *DescribeUserAlertCountRequest
	GetStartDate() *string
}

type DescribeUserAlertCountRequest struct {
	// example:
	//
	// 2025-06-10
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 2025-05-12
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeUserAlertCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAlertCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserAlertCountRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeUserAlertCountRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeUserAlertCountRequest) SetEndDate(v string) *DescribeUserAlertCountRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUserAlertCountRequest) SetStartDate(v string) *DescribeUserAlertCountRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeUserAlertCountRequest) Validate() error {
	return dara.Validate(s)
}
