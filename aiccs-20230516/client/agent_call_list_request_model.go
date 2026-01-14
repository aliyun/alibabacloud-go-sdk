// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentCallListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *AgentCallListRequest
	GetAgentId() *int64
	SetAgentTag(v string) *AgentCallListRequest
	GetAgentTag() *string
	SetCallDate(v string) *AgentCallListRequest
	GetCallDate() *string
	SetEndCallDate(v string) *AgentCallListRequest
	GetEndCallDate() *string
	SetNumberMD5s(v []*string) *AgentCallListRequest
	GetNumberMD5s() []*string
	SetNumbers(v []*string) *AgentCallListRequest
	GetNumbers() []*string
	SetOwnerId(v int64) *AgentCallListRequest
	GetOwnerId() *int64
	SetPage(v int64) *AgentCallListRequest
	GetPage() *int64
	SetPageSize(v int64) *AgentCallListRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *AgentCallListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AgentCallListRequest
	GetResourceOwnerId() *int64
	SetTags(v []*string) *AgentCallListRequest
	GetTags() []*string
}

type AgentCallListRequest struct {
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
	NumberMD5s []*string `json:"NumberMD5s,omitempty" xml:"NumberMD5s,omitempty" type:"Repeated"`
	// 号码列表
	Numbers []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s AgentCallListRequest) String() string {
	return dara.Prettify(s)
}

func (s AgentCallListRequest) GoString() string {
	return s.String()
}

func (s *AgentCallListRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *AgentCallListRequest) GetAgentTag() *string {
	return s.AgentTag
}

func (s *AgentCallListRequest) GetCallDate() *string {
	return s.CallDate
}

func (s *AgentCallListRequest) GetEndCallDate() *string {
	return s.EndCallDate
}

func (s *AgentCallListRequest) GetNumberMD5s() []*string {
	return s.NumberMD5s
}

func (s *AgentCallListRequest) GetNumbers() []*string {
	return s.Numbers
}

func (s *AgentCallListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AgentCallListRequest) GetPage() *int64 {
	return s.Page
}

func (s *AgentCallListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *AgentCallListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AgentCallListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AgentCallListRequest) GetTags() []*string {
	return s.Tags
}

func (s *AgentCallListRequest) SetAgentId(v int64) *AgentCallListRequest {
	s.AgentId = &v
	return s
}

func (s *AgentCallListRequest) SetAgentTag(v string) *AgentCallListRequest {
	s.AgentTag = &v
	return s
}

func (s *AgentCallListRequest) SetCallDate(v string) *AgentCallListRequest {
	s.CallDate = &v
	return s
}

func (s *AgentCallListRequest) SetEndCallDate(v string) *AgentCallListRequest {
	s.EndCallDate = &v
	return s
}

func (s *AgentCallListRequest) SetNumberMD5s(v []*string) *AgentCallListRequest {
	s.NumberMD5s = v
	return s
}

func (s *AgentCallListRequest) SetNumbers(v []*string) *AgentCallListRequest {
	s.Numbers = v
	return s
}

func (s *AgentCallListRequest) SetOwnerId(v int64) *AgentCallListRequest {
	s.OwnerId = &v
	return s
}

func (s *AgentCallListRequest) SetPage(v int64) *AgentCallListRequest {
	s.Page = &v
	return s
}

func (s *AgentCallListRequest) SetPageSize(v int64) *AgentCallListRequest {
	s.PageSize = &v
	return s
}

func (s *AgentCallListRequest) SetResourceOwnerAccount(v string) *AgentCallListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AgentCallListRequest) SetResourceOwnerId(v int64) *AgentCallListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AgentCallListRequest) SetTags(v []*string) *AgentCallListRequest {
	s.Tags = v
	return s
}

func (s *AgentCallListRequest) Validate() error {
	return dara.Validate(s)
}
