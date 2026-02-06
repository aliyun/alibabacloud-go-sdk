// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterTimedScheduleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstanceName(v string) *DescribeParameterTimedScheduleTaskRequest
	GetDbInstanceName() *string
}

type DescribeParameterTimedScheduleTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze2za3is7baay****
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
}

func (s DescribeParameterTimedScheduleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTimedScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterTimedScheduleTaskRequest) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DescribeParameterTimedScheduleTaskRequest) SetDbInstanceName(v string) *DescribeParameterTimedScheduleTaskRequest {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeParameterTimedScheduleTaskRequest) Validate() error {
	return dara.Validate(s)
}
