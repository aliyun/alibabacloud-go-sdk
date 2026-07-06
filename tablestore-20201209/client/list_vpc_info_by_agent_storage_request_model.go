// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcInfoByAgentStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageName(v string) *ListVpcInfoByAgentStorageRequest
	GetAgentStorageName() *string
	SetPageNum(v int64) *ListVpcInfoByAgentStorageRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListVpcInfoByAgentStorageRequest
	GetPageSize() *int64
}

type ListVpcInfoByAgentStorageRequest struct {
	// The agent storage name.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-test
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of VPCs per page for the query.
	//
	// example:
	//
	// 8
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVpcInfoByAgentStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcInfoByAgentStorageRequest) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByAgentStorageRequest) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *ListVpcInfoByAgentStorageRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListVpcInfoByAgentStorageRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVpcInfoByAgentStorageRequest) SetAgentStorageName(v string) *ListVpcInfoByAgentStorageRequest {
	s.AgentStorageName = &v
	return s
}

func (s *ListVpcInfoByAgentStorageRequest) SetPageNum(v int64) *ListVpcInfoByAgentStorageRequest {
	s.PageNum = &v
	return s
}

func (s *ListVpcInfoByAgentStorageRequest) SetPageSize(v int64) *ListVpcInfoByAgentStorageRequest {
	s.PageSize = &v
	return s
}

func (s *ListVpcInfoByAgentStorageRequest) Validate() error {
	return dara.Validate(s)
}
