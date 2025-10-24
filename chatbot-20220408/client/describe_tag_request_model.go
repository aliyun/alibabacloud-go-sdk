// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DescribeTagRequest
	GetAgentKey() *string
	SetClientToken(v string) *DescribeTagRequest
	GetClientToken() *string
	SetGroupId(v int64) *DescribeTagRequest
	GetGroupId() *int64
	SetId(v int64) *DescribeTagRequest
	GetId() *int64
}

type DescribeTagRequest struct {
	// example:
	//
	// sssxx
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 45353
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 443434
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DescribeTagRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeTagRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeTagRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeTagRequest) SetAgentKey(v string) *DescribeTagRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeTagRequest) SetClientToken(v string) *DescribeTagRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeTagRequest) SetGroupId(v int64) *DescribeTagRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeTagRequest) SetId(v int64) *DescribeTagRequest {
	s.Id = &v
	return s
}

func (s *DescribeTagRequest) Validate() error {
	return dara.Validate(s)
}
