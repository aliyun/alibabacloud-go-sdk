// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeApiTrafficDataRequest
	GetApiId() *string
	SetEndTime(v string) *DescribeApiTrafficDataRequest
	GetEndTime() *string
	SetGroupId(v string) *DescribeApiTrafficDataRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DescribeApiTrafficDataRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeApiTrafficDataRequest
	GetStageName() *string
	SetStartTime(v string) *DescribeApiTrafficDataRequest
	GetStartTime() *string
}

type DescribeApiTrafficDataRequest struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// d6f679aeb3be4b91b3688e887ca1fe16
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The end time in UTC. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-07-23T09:28:48Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 63be9002440b4778a61122f14c2b2bbb
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**: the test environment
	//
	// 	- PRE: the pre-release environment
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The start time in UTC. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-07-23T08:28:48Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApiTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApiTrafficDataRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiTrafficDataRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApiTrafficDataRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApiTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApiTrafficDataRequest) SetApiId(v string) *DescribeApiTrafficDataRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiTrafficDataRequest) SetEndTime(v string) *DescribeApiTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApiTrafficDataRequest) SetGroupId(v string) *DescribeApiTrafficDataRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiTrafficDataRequest) SetSecurityToken(v string) *DescribeApiTrafficDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiTrafficDataRequest) SetStageName(v string) *DescribeApiTrafficDataRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApiTrafficDataRequest) SetStartTime(v string) *DescribeApiTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApiTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
