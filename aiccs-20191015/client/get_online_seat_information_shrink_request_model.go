// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnlineSeatInformationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdsShrink(v string) *GetOnlineSeatInformationShrinkRequest
	GetAgentIdsShrink() *string
	SetCurrentPage(v int32) *GetOnlineSeatInformationShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetOnlineSeatInformationShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetOnlineSeatInformationShrinkRequest
	GetEndDate() *int64
	SetInstanceId(v string) *GetOnlineSeatInformationShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetOnlineSeatInformationShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetOnlineSeatInformationShrinkRequest
	GetStartDate() *int64
}

type GetOnlineSeatInformationShrinkRequest struct {
	// List of agent IDs.
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// Current page number. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of department IDs.
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// End date UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1617761765000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// AICCS instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page size. The value must be greater than **0**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start date UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetOnlineSeatInformationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineSeatInformationShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetOnlineSeatInformationShrinkRequest) GetAgentIdsShrink() *string {
	return s.AgentIdsShrink
}

func (s *GetOnlineSeatInformationShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetOnlineSeatInformationShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetOnlineSeatInformationShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetOnlineSeatInformationShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetOnlineSeatInformationShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOnlineSeatInformationShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetOnlineSeatInformationShrinkRequest) SetAgentIdsShrink(v string) *GetOnlineSeatInformationShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetOnlineSeatInformationShrinkRequest) SetCurrentPage(v int32) *GetOnlineSeatInformationShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetOnlineSeatInformationShrinkRequest) SetDepIdsShrink(v string) *GetOnlineSeatInformationShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetOnlineSeatInformationShrinkRequest) SetEndDate(v int64) *GetOnlineSeatInformationShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetOnlineSeatInformationShrinkRequest) SetInstanceId(v string) *GetOnlineSeatInformationShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOnlineSeatInformationShrinkRequest) SetPageSize(v int32) *GetOnlineSeatInformationShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetOnlineSeatInformationShrinkRequest) SetStartDate(v int64) *GetOnlineSeatInformationShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetOnlineSeatInformationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
