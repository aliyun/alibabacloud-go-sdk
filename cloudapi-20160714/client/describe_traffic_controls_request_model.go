// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrafficControlsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeTrafficControlsRequest
	GetApiId() *string
	SetGroupId(v string) *DescribeTrafficControlsRequest
	GetGroupId() *string
	SetPageNumber(v int32) *DescribeTrafficControlsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTrafficControlsRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeTrafficControlsRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeTrafficControlsRequest
	GetStageName() *string
	SetTrafficControlId(v string) *DescribeTrafficControlsRequest
	GetTrafficControlId() *string
	SetTrafficControlName(v string) *DescribeTrafficControlsRequest
	GetTrafficControlName() *string
}

type DescribeTrafficControlsRequest struct {
	// The specified API ID. This parameter must be specified together with GroupId and StageName.
	//
	// example:
	//
	// 3b81fd160f5645e097cc8855d75a1cf6
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The specified group ID. This parameter must be specified together with ApiId and StageName.
	//
	// example:
	//
	// 0009db9c828549768a200320714b8930
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// 436fa39b-b3b9-40c5-ae5d-ce3e000e38c5
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The environment name. This parameter must be specified together with GroupId and ApiId. Valid values:********
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The ID of the throttling policy.
	//
	// example:
	//
	// tf123456
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
	// The name of the throttling policy.
	//
	// example:
	//
	// ThrottlingTest
	TrafficControlName *string `json:"TrafficControlName,omitempty" xml:"TrafficControlName,omitempty"`
}

func (s DescribeTrafficControlsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeTrafficControlsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeTrafficControlsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTrafficControlsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTrafficControlsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeTrafficControlsRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeTrafficControlsRequest) GetTrafficControlId() *string {
	return s.TrafficControlId
}

func (s *DescribeTrafficControlsRequest) GetTrafficControlName() *string {
	return s.TrafficControlName
}

func (s *DescribeTrafficControlsRequest) SetApiId(v string) *DescribeTrafficControlsRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetGroupId(v string) *DescribeTrafficControlsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetPageNumber(v int32) *DescribeTrafficControlsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetPageSize(v int32) *DescribeTrafficControlsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetSecurityToken(v string) *DescribeTrafficControlsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetStageName(v string) *DescribeTrafficControlsRequest {
	s.StageName = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetTrafficControlId(v string) *DescribeTrafficControlsRequest {
	s.TrafficControlId = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetTrafficControlName(v string) *DescribeTrafficControlsRequest {
	s.TrafficControlName = &v
	return s
}

func (s *DescribeTrafficControlsRequest) Validate() error {
	return dara.Validate(s)
}
