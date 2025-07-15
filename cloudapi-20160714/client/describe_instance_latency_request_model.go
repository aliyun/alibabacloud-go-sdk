// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceLatencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInstanceLatencyRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeInstanceLatencyRequest
	GetInstanceId() *string
	SetSecurityToken(v string) *DescribeInstanceLatencyRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeInstanceLatencyRequest
	GetStageName() *string
	SetStartTime(v string) *DescribeInstanceLatencyRequest
	GetStartTime() *string
}

type DescribeInstanceLatencyRequest struct {
	// The end time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-06T02:05:13Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-cn-v641jf5tt01v
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The environment in which the API is requested. Valid values:
	//
	// 	- **RELEASE**: the production environment
	//
	// 	- **PRE**: the pre-release environment
	//
	// 	- **TEST**: the test environment
	//
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The start time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-15T11:07:05Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceLatencyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLatencyRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLatencyRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceLatencyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceLatencyRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceLatencyRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeInstanceLatencyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceLatencyRequest) SetEndTime(v string) *DescribeInstanceLatencyRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceLatencyRequest) SetInstanceId(v string) *DescribeInstanceLatencyRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceLatencyRequest) SetSecurityToken(v string) *DescribeInstanceLatencyRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceLatencyRequest) SetStageName(v string) *DescribeInstanceLatencyRequest {
	s.StageName = &v
	return s
}

func (s *DescribeInstanceLatencyRequest) SetStartTime(v string) *DescribeInstanceLatencyRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceLatencyRequest) Validate() error {
	return dara.Validate(s)
}
