// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMemoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *SearchMemoriesRequest
	GetApplicationId() *string
	SetCreateTimeBegin(v string) *SearchMemoriesRequest
	GetCreateTimeBegin() *string
	SetCreateTimeEnd(v string) *SearchMemoriesRequest
	GetCreateTimeEnd() *string
	SetMemoryAgentId(v string) *SearchMemoriesRequest
	GetMemoryAgentId() *string
	SetMemoryUserId(v string) *SearchMemoriesRequest
	GetMemoryUserId() *string
	SetQuery(v string) *SearchMemoriesRequest
	GetQuery() *string
	SetTopK(v string) *SearchMemoriesRequest
	GetTopK() *string
}

type SearchMemoriesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// yyyy-MM-ddTHH:mm:ssZ
	CreateTimeBegin *string `json:"CreateTimeBegin,omitempty" xml:"CreateTimeBegin,omitempty"`
	// example:
	//
	// yyyy-MM-ddTHH:mm:ssZ
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// agent1
	MemoryAgentId *string `json:"MemoryAgentId,omitempty" xml:"MemoryAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user1
	MemoryUserId *string `json:"MemoryUserId,omitempty" xml:"MemoryUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// who are you
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 3
	TopK *string `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s SearchMemoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesRequest) GoString() string {
	return s.String()
}

func (s *SearchMemoriesRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *SearchMemoriesRequest) GetCreateTimeBegin() *string {
	return s.CreateTimeBegin
}

func (s *SearchMemoriesRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *SearchMemoriesRequest) GetMemoryAgentId() *string {
	return s.MemoryAgentId
}

func (s *SearchMemoriesRequest) GetMemoryUserId() *string {
	return s.MemoryUserId
}

func (s *SearchMemoriesRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchMemoriesRequest) GetTopK() *string {
	return s.TopK
}

func (s *SearchMemoriesRequest) SetApplicationId(v string) *SearchMemoriesRequest {
	s.ApplicationId = &v
	return s
}

func (s *SearchMemoriesRequest) SetCreateTimeBegin(v string) *SearchMemoriesRequest {
	s.CreateTimeBegin = &v
	return s
}

func (s *SearchMemoriesRequest) SetCreateTimeEnd(v string) *SearchMemoriesRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchMemoriesRequest) SetMemoryAgentId(v string) *SearchMemoriesRequest {
	s.MemoryAgentId = &v
	return s
}

func (s *SearchMemoriesRequest) SetMemoryUserId(v string) *SearchMemoriesRequest {
	s.MemoryUserId = &v
	return s
}

func (s *SearchMemoriesRequest) SetQuery(v string) *SearchMemoriesRequest {
	s.Query = &v
	return s
}

func (s *SearchMemoriesRequest) SetTopK(v string) *SearchMemoriesRequest {
	s.TopK = &v
	return s
}

func (s *SearchMemoriesRequest) Validate() error {
	return dara.Validate(s)
}
