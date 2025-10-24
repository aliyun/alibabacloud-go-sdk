// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DescribeTagGroupRequest
	GetAgentKey() *string
	SetClientToken(v string) *DescribeTagGroupRequest
	GetClientToken() *string
	SetId(v int64) *DescribeTagGroupRequest
	GetId() *int64
}

type DescribeTagGroupRequest struct {
	// example:
	//
	// xxx
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 443434
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeTagGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagGroupRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DescribeTagGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeTagGroupRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeTagGroupRequest) SetAgentKey(v string) *DescribeTagGroupRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeTagGroupRequest) SetClientToken(v string) *DescribeTagGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeTagGroupRequest) SetId(v int64) *DescribeTagGroupRequest {
	s.Id = &v
	return s
}

func (s *DescribeTagGroupRequest) Validate() error {
	return dara.Validate(s)
}
