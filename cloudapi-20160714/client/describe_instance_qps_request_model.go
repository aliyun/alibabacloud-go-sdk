// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceQpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInstanceQpsRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeInstanceQpsRequest
	GetInstanceId() *string
	SetSecurityToken(v string) *DescribeInstanceQpsRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeInstanceQpsRequest
	GetStageName() *string
	SetStartTime(v string) *DescribeInstanceQpsRequest
	GetStartTime() *string
}

type DescribeInstanceQpsRequest struct {
	// The end time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-03-29T06:25:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-bj-6c219f1fd5d4
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
	// 2022-02-10T06:03:47Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceQpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceQpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceQpsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceQpsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceQpsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceQpsRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeInstanceQpsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceQpsRequest) SetEndTime(v string) *DescribeInstanceQpsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceQpsRequest) SetInstanceId(v string) *DescribeInstanceQpsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceQpsRequest) SetSecurityToken(v string) *DescribeInstanceQpsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceQpsRequest) SetStageName(v string) *DescribeInstanceQpsRequest {
	s.StageName = &v
	return s
}

func (s *DescribeInstanceQpsRequest) SetStartTime(v string) *DescribeInstanceQpsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceQpsRequest) Validate() error {
	return dara.Validate(s)
}
