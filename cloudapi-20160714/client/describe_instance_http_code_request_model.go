// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceHttpCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInstanceHttpCodeRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeInstanceHttpCodeRequest
	GetInstanceId() *string
	SetSecurityToken(v string) *DescribeInstanceHttpCodeRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeInstanceHttpCodeRequest
	GetStageName() *string
	SetStartTime(v string) *DescribeInstanceHttpCodeRequest
	GetStartTime() *string
}

type DescribeInstanceHttpCodeRequest struct {
	// The end time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-07-21T06:05:52Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-cn-m7r227yy2004
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
	// 2022-04-14T02:12:10Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceHttpCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHttpCodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHttpCodeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceHttpCodeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceHttpCodeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceHttpCodeRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeInstanceHttpCodeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceHttpCodeRequest) SetEndTime(v string) *DescribeInstanceHttpCodeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceHttpCodeRequest) SetInstanceId(v string) *DescribeInstanceHttpCodeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceHttpCodeRequest) SetSecurityToken(v string) *DescribeInstanceHttpCodeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceHttpCodeRequest) SetStageName(v string) *DescribeInstanceHttpCodeRequest {
	s.StageName = &v
	return s
}

func (s *DescribeInstanceHttpCodeRequest) SetStartTime(v string) *DescribeInstanceHttpCodeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceHttpCodeRequest) Validate() error {
	return dara.Validate(s)
}
