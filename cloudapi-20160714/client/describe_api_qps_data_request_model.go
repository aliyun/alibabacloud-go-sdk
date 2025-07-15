// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiQpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeApiQpsDataRequest
	GetApiId() *string
	SetEndTime(v string) *DescribeApiQpsDataRequest
	GetEndTime() *string
	SetGroupId(v string) *DescribeApiQpsDataRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DescribeApiQpsDataRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeApiQpsDataRequest
	GetStageName() *string
	SetStartTime(v string) *DescribeApiQpsDataRequest
	GetStartTime() *string
}

type DescribeApiQpsDataRequest struct {
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

func (s DescribeApiQpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiQpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApiQpsDataRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiQpsDataRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApiQpsDataRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApiQpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApiQpsDataRequest) SetApiId(v string) *DescribeApiQpsDataRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiQpsDataRequest) SetEndTime(v string) *DescribeApiQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApiQpsDataRequest) SetGroupId(v string) *DescribeApiQpsDataRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiQpsDataRequest) SetSecurityToken(v string) *DescribeApiQpsDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiQpsDataRequest) SetStageName(v string) *DescribeApiQpsDataRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApiQpsDataRequest) SetStartTime(v string) *DescribeApiQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApiQpsDataRequest) Validate() error {
	return dara.Validate(s)
}
