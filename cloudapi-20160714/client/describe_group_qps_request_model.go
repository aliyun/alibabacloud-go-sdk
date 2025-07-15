// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupQpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeGroupQpsRequest
	GetEndTime() *string
	SetGroupId(v string) *DescribeGroupQpsRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DescribeGroupQpsRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeGroupQpsRequest
	GetStageName() *string
	SetStartTime(v string) *DescribeGroupQpsRequest
	GetStartTime() *string
}

type DescribeGroupQpsRequest struct {
	// The end time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-08-23T07:27:44Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The API group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 264c60db9f764345a13ac5c825b229b9
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The environment to which the API group is published. Valid values:
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
	// The start time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-08-12T06:09:52Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeGroupQpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupQpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupQpsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeGroupQpsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGroupQpsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeGroupQpsRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeGroupQpsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeGroupQpsRequest) SetEndTime(v string) *DescribeGroupQpsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeGroupQpsRequest) SetGroupId(v string) *DescribeGroupQpsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupQpsRequest) SetSecurityToken(v string) *DescribeGroupQpsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGroupQpsRequest) SetStageName(v string) *DescribeGroupQpsRequest {
	s.StageName = &v
	return s
}

func (s *DescribeGroupQpsRequest) SetStartTime(v string) *DescribeGroupQpsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeGroupQpsRequest) Validate() error {
	return dara.Validate(s)
}
