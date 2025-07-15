// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTrafficRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInstanceTrafficRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeInstanceTrafficRequest
	GetInstanceId() *string
	SetSecurityToken(v string) *DescribeInstanceTrafficRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeInstanceTrafficRequest
	GetStageName() *string
	SetStartTime(v string) *DescribeInstanceTrafficRequest
	GetStartTime() *string
}

type DescribeInstanceTrafficRequest struct {
	// The end time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-04-01T06:34:03Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-cn-2r426lavr001
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The environment in which the API runs. Valid values:
	//
	// 	- **RELEASE**: the production environment
	//
	// 	- **PRE**: the pre-release environment
	//
	// 	- **TEST: the test environment**
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
	// 2022-03-31T03:42:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceTrafficRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTrafficRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceTrafficRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceTrafficRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceTrafficRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeInstanceTrafficRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceTrafficRequest) SetEndTime(v string) *DescribeInstanceTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceTrafficRequest) SetInstanceId(v string) *DescribeInstanceTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceTrafficRequest) SetSecurityToken(v string) *DescribeInstanceTrafficRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceTrafficRequest) SetStageName(v string) *DescribeInstanceTrafficRequest {
	s.StageName = &v
	return s
}

func (s *DescribeInstanceTrafficRequest) SetStartTime(v string) *DescribeInstanceTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceTrafficRequest) Validate() error {
	return dara.Validate(s)
}
