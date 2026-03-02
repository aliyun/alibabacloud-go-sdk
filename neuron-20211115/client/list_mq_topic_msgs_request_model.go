// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMqTopicMsgsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListMqTopicMsgsRequest
	GetEndTime() *string
	SetMsgId(v string) *ListMqTopicMsgsRequest
	GetMsgId() *string
	SetOrderBy(v string) *ListMqTopicMsgsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMqTopicMsgsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListMqTopicMsgsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMqTopicMsgsRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListMqTopicMsgsRequest
	GetStartTime() *string
}

type ListMqTopicMsgsRequest struct {
	// This parameter is required.
	EndTime        *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MsgId          *string `json:"msgId,omitempty" xml:"msgId,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListMqTopicMsgsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMqTopicMsgsRequest) GoString() string {
	return s.String()
}

func (s *ListMqTopicMsgsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListMqTopicMsgsRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *ListMqTopicMsgsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMqTopicMsgsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMqTopicMsgsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMqTopicMsgsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMqTopicMsgsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListMqTopicMsgsRequest) SetEndTime(v string) *ListMqTopicMsgsRequest {
	s.EndTime = &v
	return s
}

func (s *ListMqTopicMsgsRequest) SetMsgId(v string) *ListMqTopicMsgsRequest {
	s.MsgId = &v
	return s
}

func (s *ListMqTopicMsgsRequest) SetOrderBy(v string) *ListMqTopicMsgsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMqTopicMsgsRequest) SetOrderDirection(v string) *ListMqTopicMsgsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMqTopicMsgsRequest) SetPageNumber(v int32) *ListMqTopicMsgsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMqTopicMsgsRequest) SetPageSize(v int32) *ListMqTopicMsgsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMqTopicMsgsRequest) SetStartTime(v string) *ListMqTopicMsgsRequest {
	s.StartTime = &v
	return s
}

func (s *ListMqTopicMsgsRequest) Validate() error {
	return dara.Validate(s)
}
