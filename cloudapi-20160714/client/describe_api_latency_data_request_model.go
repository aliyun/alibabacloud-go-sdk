// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiLatencyDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeApiLatencyDataRequest
	GetApiId() *string
	SetEndTime(v string) *DescribeApiLatencyDataRequest
	GetEndTime() *string
	SetGroupId(v string) *DescribeApiLatencyDataRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DescribeApiLatencyDataRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeApiLatencyDataRequest
	GetStageName() *string
	SetStartTime(v string) *DescribeApiLatencyDataRequest
	GetStartTime() *string
}

type DescribeApiLatencyDataRequest struct {
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
	// 	- **PRE**: the pre-release environment
	//
	// 	- **TEST**
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

func (s DescribeApiLatencyDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiLatencyDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyDataRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiLatencyDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApiLatencyDataRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiLatencyDataRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApiLatencyDataRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApiLatencyDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApiLatencyDataRequest) SetApiId(v string) *DescribeApiLatencyDataRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiLatencyDataRequest) SetEndTime(v string) *DescribeApiLatencyDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApiLatencyDataRequest) SetGroupId(v string) *DescribeApiLatencyDataRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiLatencyDataRequest) SetSecurityToken(v string) *DescribeApiLatencyDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiLatencyDataRequest) SetStageName(v string) *DescribeApiLatencyDataRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApiLatencyDataRequest) SetStartTime(v string) *DescribeApiLatencyDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApiLatencyDataRequest) Validate() error {
	return dara.Validate(s)
}
