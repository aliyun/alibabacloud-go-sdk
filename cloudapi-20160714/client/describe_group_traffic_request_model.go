// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupTrafficRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeGroupTrafficRequest
	GetEndTime() *string
	SetGroupId(v string) *DescribeGroupTrafficRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DescribeGroupTrafficRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeGroupTrafficRequest
	GetStageName() *string
	SetStartTime(v string) *DescribeGroupTrafficRequest
	GetStartTime() *string
}

type DescribeGroupTrafficRequest struct {
	// The end time for the query. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-06-16T02:16:53Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16df9d11caa04900bcafe23b38a81600
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The environment to which the APIs in the API group are published. Valid values:
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
	// The start time for the query. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-29T01:27:43Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeGroupTrafficRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupTrafficRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeGroupTrafficRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGroupTrafficRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeGroupTrafficRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeGroupTrafficRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeGroupTrafficRequest) SetEndTime(v string) *DescribeGroupTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeGroupTrafficRequest) SetGroupId(v string) *DescribeGroupTrafficRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupTrafficRequest) SetSecurityToken(v string) *DescribeGroupTrafficRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGroupTrafficRequest) SetStageName(v string) *DescribeGroupTrafficRequest {
	s.StageName = &v
	return s
}

func (s *DescribeGroupTrafficRequest) SetStartTime(v string) *DescribeGroupTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeGroupTrafficRequest) Validate() error {
	return dara.Validate(s)
}
