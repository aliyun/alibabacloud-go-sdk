// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageQueryAgentListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *PageQueryAgentListRequest
	GetAgentId() *string
	SetAgentName(v string) *PageQueryAgentListRequest
	GetAgentName() *string
	SetOwnerId(v int64) *PageQueryAgentListRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *PageQueryAgentListRequest
	GetPageNo() *int64
	SetPageSize(v int64) *PageQueryAgentListRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *PageQueryAgentListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PageQueryAgentListRequest
	GetResourceOwnerId() *int64
	SetStatus(v int64) *PageQueryAgentListRequest
	GetStatus() *int64
}

type PageQueryAgentListRequest struct {
	// example:
	//
	// 1231*******
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 示例值示例值
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 7
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PageQueryAgentListRequest) String() string {
	return dara.Prettify(s)
}

func (s PageQueryAgentListRequest) GoString() string {
	return s.String()
}

func (s *PageQueryAgentListRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *PageQueryAgentListRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *PageQueryAgentListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PageQueryAgentListRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *PageQueryAgentListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *PageQueryAgentListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PageQueryAgentListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PageQueryAgentListRequest) GetStatus() *int64 {
	return s.Status
}

func (s *PageQueryAgentListRequest) SetAgentId(v string) *PageQueryAgentListRequest {
	s.AgentId = &v
	return s
}

func (s *PageQueryAgentListRequest) SetAgentName(v string) *PageQueryAgentListRequest {
	s.AgentName = &v
	return s
}

func (s *PageQueryAgentListRequest) SetOwnerId(v int64) *PageQueryAgentListRequest {
	s.OwnerId = &v
	return s
}

func (s *PageQueryAgentListRequest) SetPageNo(v int64) *PageQueryAgentListRequest {
	s.PageNo = &v
	return s
}

func (s *PageQueryAgentListRequest) SetPageSize(v int64) *PageQueryAgentListRequest {
	s.PageSize = &v
	return s
}

func (s *PageQueryAgentListRequest) SetResourceOwnerAccount(v string) *PageQueryAgentListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PageQueryAgentListRequest) SetResourceOwnerId(v int64) *PageQueryAgentListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PageQueryAgentListRequest) SetStatus(v int64) *PageQueryAgentListRequest {
	s.Status = &v
	return s
}

func (s *PageQueryAgentListRequest) Validate() error {
	return dara.Validate(s)
}
