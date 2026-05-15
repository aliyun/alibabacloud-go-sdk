// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v int64) *GetAiOutboundTaskListRequest
	GetCreateTimeEnd() *int64
	SetCreateTimeStart(v int64) *GetAiOutboundTaskListRequest
	GetCreateTimeStart() *int64
	SetCurrentPage(v int32) *GetAiOutboundTaskListRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *GetAiOutboundTaskListRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetAiOutboundTaskListRequest
	GetPageSize() *int32
	SetSearchKey(v string) *GetAiOutboundTaskListRequest
	GetSearchKey() *string
	SetStatus(v int32) *GetAiOutboundTaskListRequest
	GetStatus() *int32
	SetType(v int32) *GetAiOutboundTaskListRequest
	GetType() *int32
}

type GetAiOutboundTaskListRequest struct {
	// example:
	//
	// 1617761765000
	CreateTimeEnd *int64 `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// 1615083365000
	CreateTimeStart *int64 `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent_****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1763****
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAiOutboundTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskListRequest) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskListRequest) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *GetAiOutboundTaskListRequest) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *GetAiOutboundTaskListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAiOutboundTaskListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAiOutboundTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAiOutboundTaskListRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *GetAiOutboundTaskListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetAiOutboundTaskListRequest) GetType() *int32 {
	return s.Type
}

func (s *GetAiOutboundTaskListRequest) SetCreateTimeEnd(v int64) *GetAiOutboundTaskListRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *GetAiOutboundTaskListRequest) SetCreateTimeStart(v int64) *GetAiOutboundTaskListRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *GetAiOutboundTaskListRequest) SetCurrentPage(v int32) *GetAiOutboundTaskListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAiOutboundTaskListRequest) SetInstanceId(v string) *GetAiOutboundTaskListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAiOutboundTaskListRequest) SetPageSize(v int32) *GetAiOutboundTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *GetAiOutboundTaskListRequest) SetSearchKey(v string) *GetAiOutboundTaskListRequest {
	s.SearchKey = &v
	return s
}

func (s *GetAiOutboundTaskListRequest) SetStatus(v int32) *GetAiOutboundTaskListRequest {
	s.Status = &v
	return s
}

func (s *GetAiOutboundTaskListRequest) SetType(v int32) *GetAiOutboundTaskListRequest {
	s.Type = &v
	return s
}

func (s *GetAiOutboundTaskListRequest) Validate() error {
	return dara.Validate(s)
}
