// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsResourceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpiredEndTime(v string) *DescribeEnsResourceUsageRequest
	GetExpiredEndTime() *string
	SetExpiredStartTime(v string) *DescribeEnsResourceUsageRequest
	GetExpiredStartTime() *string
}

type DescribeEnsResourceUsageRequest struct {
	// The end of the time range to query. Format: yyyy-MM-dd or yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2018-06-15T09:07:23Z
	ExpiredEndTime *string `json:"ExpiredEndTime,omitempty" xml:"ExpiredEndTime,omitempty"`
	// The beginning of the time range to query. Format: yyyy-MM-dd or yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2018-06-15T09:07:23Z
	ExpiredStartTime *string `json:"ExpiredStartTime,omitempty" xml:"ExpiredStartTime,omitempty"`
}

func (s DescribeEnsResourceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsResourceUsageRequest) GetExpiredEndTime() *string {
	return s.ExpiredEndTime
}

func (s *DescribeEnsResourceUsageRequest) GetExpiredStartTime() *string {
	return s.ExpiredStartTime
}

func (s *DescribeEnsResourceUsageRequest) SetExpiredEndTime(v string) *DescribeEnsResourceUsageRequest {
	s.ExpiredEndTime = &v
	return s
}

func (s *DescribeEnsResourceUsageRequest) SetExpiredStartTime(v string) *DescribeEnsResourceUsageRequest {
	s.ExpiredStartTime = &v
	return s
}

func (s *DescribeEnsResourceUsageRequest) Validate() error {
	return dara.Validate(s)
}
