// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMemoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRelations(v []*SearchMemoriesResponseBodyRelations) *SearchMemoriesResponseBody
	GetRelations() []*SearchMemoriesResponseBodyRelations
	SetRequestId(v string) *SearchMemoriesResponseBody
	GetRequestId() *string
	SetResults(v []*SearchMemoriesResponseBodyResults) *SearchMemoriesResponseBody
	GetResults() []*SearchMemoriesResponseBodyResults
}

type SearchMemoriesResponseBody struct {
	Relations []*SearchMemoriesResponseBodyRelations `json:"relations,omitempty" xml:"relations,omitempty" type:"Repeated"`
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Results   []*SearchMemoriesResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s SearchMemoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMemoriesResponseBody) GetRelations() []*SearchMemoriesResponseBodyRelations {
	return s.Relations
}

func (s *SearchMemoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMemoriesResponseBody) GetResults() []*SearchMemoriesResponseBodyResults {
	return s.Results
}

func (s *SearchMemoriesResponseBody) SetRelations(v []*SearchMemoriesResponseBodyRelations) *SearchMemoriesResponseBody {
	s.Relations = v
	return s
}

func (s *SearchMemoriesResponseBody) SetRequestId(v string) *SearchMemoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMemoriesResponseBody) SetResults(v []*SearchMemoriesResponseBodyResults) *SearchMemoriesResponseBody {
	s.Results = v
	return s
}

func (s *SearchMemoriesResponseBody) Validate() error {
	if s.Relations != nil {
		for _, item := range s.Relations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMemoriesResponseBodyRelations struct {
	// example:
	//
	// test
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty"`
	// example:
	//
	// test
	Relationship *string `json:"relationship,omitempty" xml:"relationship,omitempty"`
	// example:
	//
	// todo_open_dingoj06pvqfeayy3lkr
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s SearchMemoriesResponseBodyRelations) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesResponseBodyRelations) GoString() string {
	return s.String()
}

func (s *SearchMemoriesResponseBodyRelations) GetDestination() *string {
	return s.Destination
}

func (s *SearchMemoriesResponseBodyRelations) GetRelationship() *string {
	return s.Relationship
}

func (s *SearchMemoriesResponseBodyRelations) GetSource() *string {
	return s.Source
}

func (s *SearchMemoriesResponseBodyRelations) SetDestination(v string) *SearchMemoriesResponseBodyRelations {
	s.Destination = &v
	return s
}

func (s *SearchMemoriesResponseBodyRelations) SetRelationship(v string) *SearchMemoriesResponseBodyRelations {
	s.Relationship = &v
	return s
}

func (s *SearchMemoriesResponseBodyRelations) SetSource(v string) *SearchMemoriesResponseBodyRelations {
	s.Source = &v
	return s
}

func (s *SearchMemoriesResponseBodyRelations) Validate() error {
	return dara.Validate(s)
}

type SearchMemoriesResponseBodyResults struct {
	// example:
	//
	// test_session_001
	ActorId *string `json:"actorId,omitempty" xml:"actorId,omitempty"`
	// example:
	//
	// 972772996913709056
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// 1762773128968
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// d791bf27236c5532056a13279baad3517042bb8d5b1bdb02e7871fa632debffe
	Hash *string `json:"hash,omitempty" xml:"hash,omitempty"`
	// example:
	//
	// 019ca1e5-7307-7d50-b943-5e628326a8ed
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// My name is Zhang San and I live in Hangzhou.
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
	// example:
	//
	// {"sessionId":"test_session_001"}
	Metadata *string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// test_session_001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// 13.21
	Score *float64 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// 1764902557784
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// test_session_001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchMemoriesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *SearchMemoriesResponseBodyResults) GetActorId() *string {
	return s.ActorId
}

func (s *SearchMemoriesResponseBodyResults) GetAgentId() *string {
	return s.AgentId
}

func (s *SearchMemoriesResponseBodyResults) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *SearchMemoriesResponseBodyResults) GetHash() *string {
	return s.Hash
}

func (s *SearchMemoriesResponseBodyResults) GetId() *string {
	return s.Id
}

func (s *SearchMemoriesResponseBodyResults) GetMemory() *string {
	return s.Memory
}

func (s *SearchMemoriesResponseBodyResults) GetMetadata() *string {
	return s.Metadata
}

func (s *SearchMemoriesResponseBodyResults) GetRole() *string {
	return s.Role
}

func (s *SearchMemoriesResponseBodyResults) GetRunId() *string {
	return s.RunId
}

func (s *SearchMemoriesResponseBodyResults) GetScore() *float64 {
	return s.Score
}

func (s *SearchMemoriesResponseBodyResults) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *SearchMemoriesResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *SearchMemoriesResponseBodyResults) SetActorId(v string) *SearchMemoriesResponseBodyResults {
	s.ActorId = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetAgentId(v string) *SearchMemoriesResponseBodyResults {
	s.AgentId = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetCreatedAt(v string) *SearchMemoriesResponseBodyResults {
	s.CreatedAt = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetHash(v string) *SearchMemoriesResponseBodyResults {
	s.Hash = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetId(v string) *SearchMemoriesResponseBodyResults {
	s.Id = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetMemory(v string) *SearchMemoriesResponseBodyResults {
	s.Memory = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetMetadata(v string) *SearchMemoriesResponseBodyResults {
	s.Metadata = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetRole(v string) *SearchMemoriesResponseBodyResults {
	s.Role = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetRunId(v string) *SearchMemoriesResponseBodyResults {
	s.RunId = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetScore(v float64) *SearchMemoriesResponseBodyResults {
	s.Score = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetUpdatedAt(v string) *SearchMemoriesResponseBodyResults {
	s.UpdatedAt = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetUserId(v string) *SearchMemoriesResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
