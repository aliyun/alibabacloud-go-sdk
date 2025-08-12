// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupMonitoringAgentProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribeGroupMonitoringAgentProcessRequest
	GetGroupId() *string
	SetPageNumber(v int32) *DescribeGroupMonitoringAgentProcessRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGroupMonitoringAgentProcessRequest
	GetPageSize() *int32
	SetProcessName(v string) *DescribeGroupMonitoringAgentProcessRequest
	GetProcessName() *string
	SetRegionId(v string) *DescribeGroupMonitoringAgentProcessRequest
	GetRegionId() *string
}

type DescribeGroupMonitoringAgentProcessRequest struct {
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The page number. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The process name.
	//
	// example:
	//
	// sshd
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGroupMonitoringAgentProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGroupMonitoringAgentProcessRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGroupMonitoringAgentProcessRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGroupMonitoringAgentProcessRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *DescribeGroupMonitoringAgentProcessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGroupMonitoringAgentProcessRequest) SetGroupId(v string) *DescribeGroupMonitoringAgentProcessRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessRequest) SetPageNumber(v int32) *DescribeGroupMonitoringAgentProcessRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessRequest) SetPageSize(v int32) *DescribeGroupMonitoringAgentProcessRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessRequest) SetProcessName(v string) *DescribeGroupMonitoringAgentProcessRequest {
	s.ProcessName = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessRequest) SetRegionId(v string) *DescribeGroupMonitoringAgentProcessRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessRequest) Validate() error {
	return dara.Validate(s)
}
