// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentRecoverCallShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *AgentRecoverCallShrinkRequest
	GetAgentId() *int64
	SetAgentTag(v string) *AgentRecoverCallShrinkRequest
	GetAgentTag() *string
	SetBeginImportTime(v string) *AgentRecoverCallShrinkRequest
	GetBeginImportTime() *string
	SetEndImportTime(v string) *AgentRecoverCallShrinkRequest
	GetEndImportTime() *string
	SetNumbersShrink(v string) *AgentRecoverCallShrinkRequest
	GetNumbersShrink() *string
	SetOwnerId(v int64) *AgentRecoverCallShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AgentRecoverCallShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AgentRecoverCallShrinkRequest
	GetResourceOwnerId() *int64
	SetTagsShrink(v string) *AgentRecoverCallShrinkRequest
	GetTagsShrink() *string
}

type AgentRecoverCallShrinkRequest struct {
	// 坐席ID
	//
	// example:
	//
	// 5
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// abc
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 查询开始导入时间
	//
	// example:
	//
	// 2020-03-06 10:10:10
	BeginImportTime *string `json:"BeginImportTime,omitempty" xml:"BeginImportTime,omitempty"`
	// 查询结束导入时间
	//
	// example:
	//
	// 2021-03-06 10:10:10
	EndImportTime *string `json:"EndImportTime,omitempty" xml:"EndImportTime,omitempty"`
	// 号码列表
	NumbersShrink        *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户自定义标签列表
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AgentRecoverCallShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AgentRecoverCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *AgentRecoverCallShrinkRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *AgentRecoverCallShrinkRequest) GetAgentTag() *string {
	return s.AgentTag
}

func (s *AgentRecoverCallShrinkRequest) GetBeginImportTime() *string {
	return s.BeginImportTime
}

func (s *AgentRecoverCallShrinkRequest) GetEndImportTime() *string {
	return s.EndImportTime
}

func (s *AgentRecoverCallShrinkRequest) GetNumbersShrink() *string {
	return s.NumbersShrink
}

func (s *AgentRecoverCallShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AgentRecoverCallShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AgentRecoverCallShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AgentRecoverCallShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *AgentRecoverCallShrinkRequest) SetAgentId(v int64) *AgentRecoverCallShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetAgentTag(v string) *AgentRecoverCallShrinkRequest {
	s.AgentTag = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetBeginImportTime(v string) *AgentRecoverCallShrinkRequest {
	s.BeginImportTime = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetEndImportTime(v string) *AgentRecoverCallShrinkRequest {
	s.EndImportTime = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetNumbersShrink(v string) *AgentRecoverCallShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetOwnerId(v int64) *AgentRecoverCallShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetResourceOwnerAccount(v string) *AgentRecoverCallShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetResourceOwnerId(v int64) *AgentRecoverCallShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetTagsShrink(v string) *AgentRecoverCallShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) Validate() error {
	return dara.Validate(s)
}
