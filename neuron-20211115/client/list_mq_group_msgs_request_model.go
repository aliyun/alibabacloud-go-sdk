// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMqGroupMsgsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListMqGroupMsgsRequest
	GetEndTime() *string
	SetMsgId(v string) *ListMqGroupMsgsRequest
	GetMsgId() *string
	SetOrderBy(v string) *ListMqGroupMsgsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMqGroupMsgsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListMqGroupMsgsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMqGroupMsgsRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListMqGroupMsgsRequest
	GetStartTime() *string
}

type ListMqGroupMsgsRequest struct {
	EndTime        *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MsgId          *string `json:"msgId,omitempty" xml:"msgId,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime      *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListMqGroupMsgsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMqGroupMsgsRequest) GoString() string {
	return s.String()
}

func (s *ListMqGroupMsgsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListMqGroupMsgsRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *ListMqGroupMsgsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMqGroupMsgsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMqGroupMsgsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMqGroupMsgsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMqGroupMsgsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListMqGroupMsgsRequest) SetEndTime(v string) *ListMqGroupMsgsRequest {
	s.EndTime = &v
	return s
}

func (s *ListMqGroupMsgsRequest) SetMsgId(v string) *ListMqGroupMsgsRequest {
	s.MsgId = &v
	return s
}

func (s *ListMqGroupMsgsRequest) SetOrderBy(v string) *ListMqGroupMsgsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMqGroupMsgsRequest) SetOrderDirection(v string) *ListMqGroupMsgsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMqGroupMsgsRequest) SetPageNumber(v int32) *ListMqGroupMsgsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMqGroupMsgsRequest) SetPageSize(v int32) *ListMqGroupMsgsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMqGroupMsgsRequest) SetStartTime(v string) *ListMqGroupMsgsRequest {
	s.StartTime = &v
	return s
}

func (s *ListMqGroupMsgsRequest) Validate() error {
	return dara.Validate(s)
}
