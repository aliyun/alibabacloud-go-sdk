// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentGroupPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentGroupId(v int64) *AgentGroupPageRequest
	GetAgentGroupId() *int64
	SetOwnerId(v int64) *AgentGroupPageRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *AgentGroupPageRequest
	GetPageNum() *int64
	SetPageSize(v int64) *AgentGroupPageRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *AgentGroupPageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AgentGroupPageRequest
	GetResourceOwnerId() *int64
}

type AgentGroupPageRequest struct {
	// 坐席组ID
	//
	// example:
	//
	// 1
	AgentGroupId *int64 `json:"AgentGroupId,omitempty" xml:"AgentGroupId,omitempty"`
	OwnerId      *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 当前页
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 分页数量
	//
	// example:
	//
	// 1
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AgentGroupPageRequest) String() string {
	return dara.Prettify(s)
}

func (s AgentGroupPageRequest) GoString() string {
	return s.String()
}

func (s *AgentGroupPageRequest) GetAgentGroupId() *int64 {
	return s.AgentGroupId
}

func (s *AgentGroupPageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AgentGroupPageRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *AgentGroupPageRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *AgentGroupPageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AgentGroupPageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AgentGroupPageRequest) SetAgentGroupId(v int64) *AgentGroupPageRequest {
	s.AgentGroupId = &v
	return s
}

func (s *AgentGroupPageRequest) SetOwnerId(v int64) *AgentGroupPageRequest {
	s.OwnerId = &v
	return s
}

func (s *AgentGroupPageRequest) SetPageNum(v int64) *AgentGroupPageRequest {
	s.PageNum = &v
	return s
}

func (s *AgentGroupPageRequest) SetPageSize(v int64) *AgentGroupPageRequest {
	s.PageSize = &v
	return s
}

func (s *AgentGroupPageRequest) SetResourceOwnerAccount(v string) *AgentGroupPageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AgentGroupPageRequest) SetResourceOwnerId(v int64) *AgentGroupPageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AgentGroupPageRequest) Validate() error {
	return dara.Validate(s)
}
