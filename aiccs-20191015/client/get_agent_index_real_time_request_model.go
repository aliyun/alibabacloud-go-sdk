// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentIndexRealTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetAgentIndexRealTimeRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetAgentIndexRealTimeRequest
	GetDepIds() []*int64
	SetGroupIds(v []*int64) *GetAgentIndexRealTimeRequest
	GetGroupIds() []*int64
	SetInstanceId(v string) *GetAgentIndexRealTimeRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetAgentIndexRealTimeRequest
	GetPageSize() *int32
}

type GetAgentIndexRealTimeRequest struct {
	// Current page number. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of department IDs.
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// List of skill group IDs.
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
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
}

func (s GetAgentIndexRealTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentIndexRealTimeRequest) GoString() string {
	return s.String()
}

func (s *GetAgentIndexRealTimeRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAgentIndexRealTimeRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetAgentIndexRealTimeRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *GetAgentIndexRealTimeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentIndexRealTimeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAgentIndexRealTimeRequest) SetCurrentPage(v int32) *GetAgentIndexRealTimeRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentIndexRealTimeRequest) SetDepIds(v []*int64) *GetAgentIndexRealTimeRequest {
	s.DepIds = v
	return s
}

func (s *GetAgentIndexRealTimeRequest) SetGroupIds(v []*int64) *GetAgentIndexRealTimeRequest {
	s.GroupIds = v
	return s
}

func (s *GetAgentIndexRealTimeRequest) SetInstanceId(v string) *GetAgentIndexRealTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentIndexRealTimeRequest) SetPageSize(v int32) *GetAgentIndexRealTimeRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentIndexRealTimeRequest) Validate() error {
	return dara.Validate(s)
}
