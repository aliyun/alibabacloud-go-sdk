// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRelations(v []*GetMemoriesResponseBodyRelations) *GetMemoriesResponseBody
	GetRelations() []*GetMemoriesResponseBodyRelations
	SetRequestId(v string) *GetMemoriesResponseBody
	GetRequestId() *string
	SetResults(v []*GetMemoriesResponseBodyResults) *GetMemoriesResponseBody
	GetResults() []*GetMemoriesResponseBodyResults
}

type GetMemoriesResponseBody struct {
	Relations []*GetMemoriesResponseBodyRelations `json:"relations,omitempty" xml:"relations,omitempty" type:"Repeated"`
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Results   []*GetMemoriesResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s GetMemoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoriesResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoriesResponseBody) GetRelations() []*GetMemoriesResponseBodyRelations {
	return s.Relations
}

func (s *GetMemoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoriesResponseBody) GetResults() []*GetMemoriesResponseBodyResults {
	return s.Results
}

func (s *GetMemoriesResponseBody) SetRelations(v []*GetMemoriesResponseBodyRelations) *GetMemoriesResponseBody {
	s.Relations = v
	return s
}

func (s *GetMemoriesResponseBody) SetRequestId(v string) *GetMemoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoriesResponseBody) SetResults(v []*GetMemoriesResponseBodyResults) *GetMemoriesResponseBody {
	s.Results = v
	return s
}

func (s *GetMemoriesResponseBody) Validate() error {
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

type GetMemoriesResponseBodyRelations struct {
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
	// test_namespace/RegistryModule-test2153b9f3-0be6-455e-9efd-05fd31f62885/alicloud
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s GetMemoriesResponseBodyRelations) String() string {
	return dara.Prettify(s)
}

func (s GetMemoriesResponseBodyRelations) GoString() string {
	return s.String()
}

func (s *GetMemoriesResponseBodyRelations) GetDestination() *string {
	return s.Destination
}

func (s *GetMemoriesResponseBodyRelations) GetRelationship() *string {
	return s.Relationship
}

func (s *GetMemoriesResponseBodyRelations) GetSource() *string {
	return s.Source
}

func (s *GetMemoriesResponseBodyRelations) SetDestination(v string) *GetMemoriesResponseBodyRelations {
	s.Destination = &v
	return s
}

func (s *GetMemoriesResponseBodyRelations) SetRelationship(v string) *GetMemoriesResponseBodyRelations {
	s.Relationship = &v
	return s
}

func (s *GetMemoriesResponseBodyRelations) SetSource(v string) *GetMemoriesResponseBodyRelations {
	s.Source = &v
	return s
}

func (s *GetMemoriesResponseBodyRelations) Validate() error {
	return dara.Validate(s)
}

type GetMemoriesResponseBodyResults struct {
	// example:
	//
	// test_session_001
	ActorId *string `json:"actorId,omitempty" xml:"actorId,omitempty"`
	// example:
	//
	// 980565235819266048
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// test_user_001
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// 1747623093939
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 73ad89f2d56c2f8615e5dd0cef7b4c41c074277c91fa0e31fc5b41802c0481f2
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
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// jr-965a0b00cb42a43b
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// 30.12
	Score *float64 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// 1752825865045
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// test_session_001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetMemoriesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s GetMemoriesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetMemoriesResponseBodyResults) GetActorId() *string {
	return s.ActorId
}

func (s *GetMemoriesResponseBodyResults) GetAgentId() *string {
	return s.AgentId
}

func (s *GetMemoriesResponseBodyResults) GetAppId() *string {
	return s.AppId
}

func (s *GetMemoriesResponseBodyResults) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetMemoriesResponseBodyResults) GetHash() *string {
	return s.Hash
}

func (s *GetMemoriesResponseBodyResults) GetId() *string {
	return s.Id
}

func (s *GetMemoriesResponseBodyResults) GetMemory() *string {
	return s.Memory
}

func (s *GetMemoriesResponseBodyResults) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *GetMemoriesResponseBodyResults) GetRole() *string {
	return s.Role
}

func (s *GetMemoriesResponseBodyResults) GetRunId() *string {
	return s.RunId
}

func (s *GetMemoriesResponseBodyResults) GetScore() *float64 {
	return s.Score
}

func (s *GetMemoriesResponseBodyResults) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetMemoriesResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *GetMemoriesResponseBodyResults) SetActorId(v string) *GetMemoriesResponseBodyResults {
	s.ActorId = &v
	return s
}

func (s *GetMemoriesResponseBodyResults) SetAgentId(v string) *GetMemoriesResponseBodyResults {
	s.AgentId = &v
	return s
}

func (s *GetMemoriesResponseBodyResults) SetAppId(v string) *GetMemoriesResponseBodyResults {
	s.AppId = &v
	return s
}

func (s *GetMemoriesResponseBodyResults) SetCreatedAt(v string) *GetMemoriesResponseBodyResults {
	s.CreatedAt = &v
	return s
}

func (s *GetMemoriesResponseBodyResults) SetHash(v string) *GetMemoriesResponseBodyResults {
	s.Hash = &v
	return s
}

func (s *GetMemoriesResponseBodyResults) SetId(v string) *GetMemoriesResponseBodyResults {
	s.Id = &v
	return s
}

func (s *GetMemoriesResponseBodyResults) SetMemory(v string) *GetMemoriesResponseBodyResults {
	s.Memory = &v
	return s
}

func (s *GetMemoriesResponseBodyResults) SetMetadata(v map[string]interface{}) *GetMemoriesResponseBodyResults {
	s.Metadata = v
	return s
}

func (s *GetMemoriesResponseBodyResults) SetRole(v string) *GetMemoriesResponseBodyResults {
	s.Role = &v
	return s
}

func (s *GetMemoriesResponseBodyResults) SetRunId(v string) *GetMemoriesResponseBodyResults {
	s.RunId = &v
	return s
}

func (s *GetMemoriesResponseBodyResults) SetScore(v float64) *GetMemoriesResponseBodyResults {
	s.Score = &v
	return s
}

func (s *GetMemoriesResponseBodyResults) SetUpdatedAt(v string) *GetMemoriesResponseBodyResults {
	s.UpdatedAt = &v
	return s
}

func (s *GetMemoriesResponseBodyResults) SetUserId(v string) *GetMemoriesResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *GetMemoriesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
