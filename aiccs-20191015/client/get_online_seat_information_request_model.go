// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnlineSeatInformationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v []*int64) *GetOnlineSeatInformationRequest
	GetAgentIds() []*int64
	SetCurrentPage(v int32) *GetOnlineSeatInformationRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetOnlineSeatInformationRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetOnlineSeatInformationRequest
	GetEndDate() *int64
	SetInstanceId(v string) *GetOnlineSeatInformationRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetOnlineSeatInformationRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetOnlineSeatInformationRequest
	GetStartDate() *int64
}

type GetOnlineSeatInformationRequest struct {
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

func (s GetOnlineSeatInformationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineSeatInformationRequest) GoString() string {
	return s.String()
}

func (s *GetOnlineSeatInformationRequest) GetAgentIds() []*int64 {
	return s.AgentIds
}

func (s *GetOnlineSeatInformationRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetOnlineSeatInformationRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetOnlineSeatInformationRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetOnlineSeatInformationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetOnlineSeatInformationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOnlineSeatInformationRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetOnlineSeatInformationRequest) SetAgentIds(v []*int64) *GetOnlineSeatInformationRequest {
	s.AgentIds = v
	return s
}

func (s *GetOnlineSeatInformationRequest) SetCurrentPage(v int32) *GetOnlineSeatInformationRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetOnlineSeatInformationRequest) SetDepIds(v []*int64) *GetOnlineSeatInformationRequest {
	s.DepIds = v
	return s
}

func (s *GetOnlineSeatInformationRequest) SetEndDate(v int64) *GetOnlineSeatInformationRequest {
	s.EndDate = &v
	return s
}

func (s *GetOnlineSeatInformationRequest) SetInstanceId(v string) *GetOnlineSeatInformationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOnlineSeatInformationRequest) SetPageSize(v int32) *GetOnlineSeatInformationRequest {
	s.PageSize = &v
	return s
}

func (s *GetOnlineSeatInformationRequest) SetStartDate(v int64) *GetOnlineSeatInformationRequest {
	s.StartDate = &v
	return s
}

func (s *GetOnlineSeatInformationRequest) Validate() error {
	return dara.Validate(s)
}
