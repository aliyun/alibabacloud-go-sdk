// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrafficControlsByApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeTrafficControlsByApiRequest
	GetApiId() *string
	SetGroupId(v string) *DescribeTrafficControlsByApiRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DescribeTrafficControlsByApiRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeTrafficControlsByApiRequest
	GetStageName() *string
}

type DescribeTrafficControlsByApiRequest struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3b81fd160f5645e097cc8855d75a1cf6
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the group to which the API to be queried belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0009db9c828549768a200320714b8930
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// 7c51b234-48d3-44e1-9b36-e2ddccc738e3
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The runtime environment of the API. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**
	//
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeTrafficControlsByApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsByApiRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsByApiRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeTrafficControlsByApiRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeTrafficControlsByApiRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeTrafficControlsByApiRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeTrafficControlsByApiRequest) SetApiId(v string) *DescribeTrafficControlsByApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeTrafficControlsByApiRequest) SetGroupId(v string) *DescribeTrafficControlsByApiRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeTrafficControlsByApiRequest) SetSecurityToken(v string) *DescribeTrafficControlsByApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTrafficControlsByApiRequest) SetStageName(v string) *DescribeTrafficControlsByApiRequest {
	s.StageName = &v
	return s
}

func (s *DescribeTrafficControlsByApiRequest) Validate() error {
	return dara.Validate(s)
}
