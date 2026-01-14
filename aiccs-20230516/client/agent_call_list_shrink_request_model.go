// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentCallListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *AgentCallListShrinkRequest
	GetAgentId() *int64
	SetAgentTag(v string) *AgentCallListShrinkRequest
	GetAgentTag() *string
	SetCallDate(v string) *AgentCallListShrinkRequest
	GetCallDate() *string
	SetEndCallDate(v string) *AgentCallListShrinkRequest
	GetEndCallDate() *string
	SetNumberMD5sShrink(v string) *AgentCallListShrinkRequest
	GetNumberMD5sShrink() *string
	SetNumbersShrink(v string) *AgentCallListShrinkRequest
	GetNumbersShrink() *string
	SetOwnerId(v int64) *AgentCallListShrinkRequest
	GetOwnerId() *int64
	SetPage(v int64) *AgentCallListShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *AgentCallListShrinkRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *AgentCallListShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AgentCallListShrinkRequest
	GetResourceOwnerId() *int64
	SetTagsShrink(v string) *AgentCallListShrinkRequest
	GetTagsShrink() *string
}

type AgentCallListShrinkRequest struct {
	// 坐席ID
	//
	// example:
	//
	// 51
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// abcd
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 开始外呼时间
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-03-06 10:10:10
	CallDate *string `json:"CallDate,omitempty" xml:"CallDate,omitempty"`
	// 结束外呼时间
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-03-07 10:10:10
	EndCallDate *string `json:"EndCallDate,omitempty" xml:"EndCallDate,omitempty"`
	// 号码MD5列表
	NumberMD5sShrink *string `json:"NumberMD5s,omitempty" xml:"NumberMD5s,omitempty"`
	// 号码列表
	NumbersShrink *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 页数
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// 每页外呼记录数
	//
	// example:
	//
	// 14
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户自定义标签列表
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AgentCallListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AgentCallListShrinkRequest) GoString() string {
	return s.String()
}

func (s *AgentCallListShrinkRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *AgentCallListShrinkRequest) GetAgentTag() *string {
	return s.AgentTag
}

func (s *AgentCallListShrinkRequest) GetCallDate() *string {
	return s.CallDate
}

func (s *AgentCallListShrinkRequest) GetEndCallDate() *string {
	return s.EndCallDate
}

func (s *AgentCallListShrinkRequest) GetNumberMD5sShrink() *string {
	return s.NumberMD5sShrink
}

func (s *AgentCallListShrinkRequest) GetNumbersShrink() *string {
	return s.NumbersShrink
}

func (s *AgentCallListShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AgentCallListShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *AgentCallListShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *AgentCallListShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AgentCallListShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AgentCallListShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *AgentCallListShrinkRequest) SetAgentId(v int64) *AgentCallListShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *AgentCallListShrinkRequest) SetAgentTag(v string) *AgentCallListShrinkRequest {
	s.AgentTag = &v
	return s
}

func (s *AgentCallListShrinkRequest) SetCallDate(v string) *AgentCallListShrinkRequest {
	s.CallDate = &v
	return s
}

func (s *AgentCallListShrinkRequest) SetEndCallDate(v string) *AgentCallListShrinkRequest {
	s.EndCallDate = &v
	return s
}

func (s *AgentCallListShrinkRequest) SetNumberMD5sShrink(v string) *AgentCallListShrinkRequest {
	s.NumberMD5sShrink = &v
	return s
}

func (s *AgentCallListShrinkRequest) SetNumbersShrink(v string) *AgentCallListShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *AgentCallListShrinkRequest) SetOwnerId(v int64) *AgentCallListShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AgentCallListShrinkRequest) SetPage(v int64) *AgentCallListShrinkRequest {
	s.Page = &v
	return s
}

func (s *AgentCallListShrinkRequest) SetPageSize(v int64) *AgentCallListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *AgentCallListShrinkRequest) SetResourceOwnerAccount(v string) *AgentCallListShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AgentCallListShrinkRequest) SetResourceOwnerId(v int64) *AgentCallListShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AgentCallListShrinkRequest) SetTagsShrink(v string) *AgentCallListShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *AgentCallListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
