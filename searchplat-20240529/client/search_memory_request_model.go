// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *SearchMemoryRequest
	GetAgentId() *string
	SetEnhancements(v map[string]interface{}) *SearchMemoryRequest
	GetEnhancements() map[string]interface{}
	SetQuery(v string) *SearchMemoryRequest
	GetQuery() *string
	SetRunId(v string) *SearchMemoryRequest
	GetRunId() *string
	SetSize(v int32) *SearchMemoryRequest
	GetSize() *int32
	SetUserId(v string) *SearchMemoryRequest
	GetUserId() *string
}

type SearchMemoryRequest struct {
	AgentId      *string                `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	Enhancements map[string]interface{} `json:"enhancements,omitempty" xml:"enhancements,omitempty"`
	// This parameter is required.
	Query  *string `json:"query,omitempty" xml:"query,omitempty"`
	RunId  *string `json:"run_id,omitempty" xml:"run_id,omitempty"`
	Size   *int32  `json:"size,omitempty" xml:"size,omitempty"`
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s SearchMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoryRequest) GoString() string {
	return s.String()
}

func (s *SearchMemoryRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *SearchMemoryRequest) GetEnhancements() map[string]interface{} {
	return s.Enhancements
}

func (s *SearchMemoryRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchMemoryRequest) GetRunId() *string {
	return s.RunId
}

func (s *SearchMemoryRequest) GetSize() *int32 {
	return s.Size
}

func (s *SearchMemoryRequest) GetUserId() *string {
	return s.UserId
}

func (s *SearchMemoryRequest) SetAgentId(v string) *SearchMemoryRequest {
	s.AgentId = &v
	return s
}

func (s *SearchMemoryRequest) SetEnhancements(v map[string]interface{}) *SearchMemoryRequest {
	s.Enhancements = v
	return s
}

func (s *SearchMemoryRequest) SetQuery(v string) *SearchMemoryRequest {
	s.Query = &v
	return s
}

func (s *SearchMemoryRequest) SetRunId(v string) *SearchMemoryRequest {
	s.RunId = &v
	return s
}

func (s *SearchMemoryRequest) SetSize(v int32) *SearchMemoryRequest {
	s.Size = &v
	return s
}

func (s *SearchMemoryRequest) SetUserId(v string) *SearchMemoryRequest {
	s.UserId = &v
	return s
}

func (s *SearchMemoryRequest) Validate() error {
	return dara.Validate(s)
}
