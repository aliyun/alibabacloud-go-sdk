// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentBasisStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v []*int64) *GetAgentBasisStatusRequest
	GetAgentIds() []*int64
	SetCurrentPage(v int32) *GetAgentBasisStatusRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetAgentBasisStatusRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetAgentBasisStatusRequest
	GetEndDate() *int64
	SetInstanceId(v string) *GetAgentBasisStatusRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetAgentBasisStatusRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetAgentBasisStatusRequest
	GetStartDate() *int64
}

type GetAgentBasisStatusRequest struct {
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CurrentPage *int32   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DepIds      []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1617761765000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetAgentBasisStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentBasisStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAgentBasisStatusRequest) GetAgentIds() []*int64 {
	return s.AgentIds
}

func (s *GetAgentBasisStatusRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAgentBasisStatusRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetAgentBasisStatusRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetAgentBasisStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentBasisStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAgentBasisStatusRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetAgentBasisStatusRequest) SetAgentIds(v []*int64) *GetAgentBasisStatusRequest {
	s.AgentIds = v
	return s
}

func (s *GetAgentBasisStatusRequest) SetCurrentPage(v int32) *GetAgentBasisStatusRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentBasisStatusRequest) SetDepIds(v []*int64) *GetAgentBasisStatusRequest {
	s.DepIds = v
	return s
}

func (s *GetAgentBasisStatusRequest) SetEndDate(v int64) *GetAgentBasisStatusRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentBasisStatusRequest) SetInstanceId(v string) *GetAgentBasisStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentBasisStatusRequest) SetPageSize(v int32) *GetAgentBasisStatusRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentBasisStatusRequest) SetStartDate(v int64) *GetAgentBasisStatusRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentBasisStatusRequest) Validate() error {
	return dara.Validate(s)
}
