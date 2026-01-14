// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentRecoverCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *AgentRecoverCallRequest
	GetAgentId() *int64
	SetAgentTag(v string) *AgentRecoverCallRequest
	GetAgentTag() *string
	SetBeginImportTime(v string) *AgentRecoverCallRequest
	GetBeginImportTime() *string
	SetEndImportTime(v string) *AgentRecoverCallRequest
	GetEndImportTime() *string
	SetNumbers(v []*string) *AgentRecoverCallRequest
	GetNumbers() []*string
	SetOwnerId(v int64) *AgentRecoverCallRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AgentRecoverCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AgentRecoverCallRequest
	GetResourceOwnerId() *int64
	SetTags(v []*string) *AgentRecoverCallRequest
	GetTags() []*string
}

type AgentRecoverCallRequest struct {
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
	Numbers              []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户自定义标签列表
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s AgentRecoverCallRequest) String() string {
	return dara.Prettify(s)
}

func (s AgentRecoverCallRequest) GoString() string {
	return s.String()
}

func (s *AgentRecoverCallRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *AgentRecoverCallRequest) GetAgentTag() *string {
	return s.AgentTag
}

func (s *AgentRecoverCallRequest) GetBeginImportTime() *string {
	return s.BeginImportTime
}

func (s *AgentRecoverCallRequest) GetEndImportTime() *string {
	return s.EndImportTime
}

func (s *AgentRecoverCallRequest) GetNumbers() []*string {
	return s.Numbers
}

func (s *AgentRecoverCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AgentRecoverCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AgentRecoverCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AgentRecoverCallRequest) GetTags() []*string {
	return s.Tags
}

func (s *AgentRecoverCallRequest) SetAgentId(v int64) *AgentRecoverCallRequest {
	s.AgentId = &v
	return s
}

func (s *AgentRecoverCallRequest) SetAgentTag(v string) *AgentRecoverCallRequest {
	s.AgentTag = &v
	return s
}

func (s *AgentRecoverCallRequest) SetBeginImportTime(v string) *AgentRecoverCallRequest {
	s.BeginImportTime = &v
	return s
}

func (s *AgentRecoverCallRequest) SetEndImportTime(v string) *AgentRecoverCallRequest {
	s.EndImportTime = &v
	return s
}

func (s *AgentRecoverCallRequest) SetNumbers(v []*string) *AgentRecoverCallRequest {
	s.Numbers = v
	return s
}

func (s *AgentRecoverCallRequest) SetOwnerId(v int64) *AgentRecoverCallRequest {
	s.OwnerId = &v
	return s
}

func (s *AgentRecoverCallRequest) SetResourceOwnerAccount(v string) *AgentRecoverCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AgentRecoverCallRequest) SetResourceOwnerId(v int64) *AgentRecoverCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AgentRecoverCallRequest) SetTags(v []*string) *AgentRecoverCallRequest {
	s.Tags = v
	return s
}

func (s *AgentRecoverCallRequest) Validate() error {
	return dara.Validate(s)
}
