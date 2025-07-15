// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupLatencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeGroupLatencyRequest
	GetEndTime() *string
	SetGroupId(v string) *DescribeGroupLatencyRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DescribeGroupLatencyRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeGroupLatencyRequest
	GetStageName() *string
	SetStartTime(v string) *DescribeGroupLatencyRequest
	GetStartTime() *string
}

type DescribeGroupLatencyRequest struct {
	// The end time of the time range to query. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-14T06:26:14Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// d825094fdd114a869f5adb443d9b7ead
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The environment in which you want to perform the query. Valid values:
	//
	// 	- **RELEASE**: the production environment
	//
	// 	- **PRE**: the staging environment
	//
	// 	- **TEST**: the test environment
	//
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The start time of the time range to query. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-03-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeGroupLatencyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupLatencyRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupLatencyRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeGroupLatencyRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGroupLatencyRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeGroupLatencyRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeGroupLatencyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeGroupLatencyRequest) SetEndTime(v string) *DescribeGroupLatencyRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeGroupLatencyRequest) SetGroupId(v string) *DescribeGroupLatencyRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupLatencyRequest) SetSecurityToken(v string) *DescribeGroupLatencyRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGroupLatencyRequest) SetStageName(v string) *DescribeGroupLatencyRequest {
	s.StageName = &v
	return s
}

func (s *DescribeGroupLatencyRequest) SetStartTime(v string) *DescribeGroupLatencyRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeGroupLatencyRequest) Validate() error {
	return dara.Validate(s)
}
