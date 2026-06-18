// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentBasisStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdsShrink(v string) *GetAgentBasisStatusShrinkRequest
	GetAgentIdsShrink() *string
	SetCurrentPage(v int32) *GetAgentBasisStatusShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetAgentBasisStatusShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetAgentBasisStatusShrinkRequest
	GetEndDate() *int64
	SetInstanceId(v string) *GetAgentBasisStatusShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetAgentBasisStatusShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetAgentBasisStatusShrinkRequest
	GetStartDate() *int64
}

type GetAgentBasisStatusShrinkRequest struct {
	// A list of agent IDs.
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// The current page number. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// A list of department IDs.
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// End Datetime UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1617761765000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it in the **Instance Management*	- section of the left-side navigation pane in the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page. The value must be greater than **0**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UNIX timestamp of the start date. Unit: milliseconds.
	//
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetAgentBasisStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentBasisStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAgentBasisStatusShrinkRequest) GetAgentIdsShrink() *string {
	return s.AgentIdsShrink
}

func (s *GetAgentBasisStatusShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAgentBasisStatusShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetAgentBasisStatusShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetAgentBasisStatusShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentBasisStatusShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAgentBasisStatusShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetAgentBasisStatusShrinkRequest) SetAgentIdsShrink(v string) *GetAgentBasisStatusShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetAgentBasisStatusShrinkRequest) SetCurrentPage(v int32) *GetAgentBasisStatusShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentBasisStatusShrinkRequest) SetDepIdsShrink(v string) *GetAgentBasisStatusShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetAgentBasisStatusShrinkRequest) SetEndDate(v int64) *GetAgentBasisStatusShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentBasisStatusShrinkRequest) SetInstanceId(v string) *GetAgentBasisStatusShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentBasisStatusShrinkRequest) SetPageSize(v int32) *GetAgentBasisStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentBasisStatusShrinkRequest) SetStartDate(v int64) *GetAgentBasisStatusShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentBasisStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
