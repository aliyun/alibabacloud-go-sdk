// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMemoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SearchMemoriesRequestBody) *SearchMemoriesRequest
	GetBody() *SearchMemoriesRequestBody
}

type SearchMemoriesRequest struct {
	// The semantic retrieval request.
	//
	// This parameter is required.
	Body *SearchMemoriesRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s SearchMemoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesRequest) GoString() string {
	return s.String()
}

func (s *SearchMemoriesRequest) GetBody() *SearchMemoriesRequestBody {
	return s.Body
}

func (s *SearchMemoriesRequest) SetBody(v *SearchMemoriesRequestBody) *SearchMemoriesRequest {
	s.Body = v
	return s
}

func (s *SearchMemoriesRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchMemoriesRequestBody struct {
	// Specifies whether to enable reranking. Default value: true.
	EnableRerank *bool `json:"enableRerank,omitempty" xml:"enableRerank,omitempty"`
	// The string key-value metadata for exact match. You can specify up to 16 keys. Each key can be up to 64 characters in length, and each value can be up to 1,024 characters in length.
	Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The minimum relevance score. Valid values: 0 to 1. Default value: 0.2. If both minScore and minSimilarity are specified, the AND logic is applied for filtering.
	//
	// example:
	//
	// 0.2
	MinScore *float32 `json:"minScore,omitempty" xml:"minScore,omitempty"`
	// The minimum similarity. Valid values: 0 to 1. Default value: 0.3. If both minScore and minSimilarity are specified, the AND logic is applied for filtering.
	//
	// example:
	//
	// 0.3
	MinSimilarity *float32 `json:"minSimilarity,omitempty" xml:"minSimilarity,omitempty"`
	// The query text. The text can be up to 8,192 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// What sports does the user like?
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The retrieval scope. UserId, AgentId, and SessionId are all optional.
	Scope *SearchMemoriesRequestBodyScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The number of results to return.
	//
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s SearchMemoriesRequestBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesRequestBody) GoString() string {
	return s.String()
}

func (s *SearchMemoriesRequestBody) GetEnableRerank() *bool {
	return s.EnableRerank
}

func (s *SearchMemoriesRequestBody) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *SearchMemoriesRequestBody) GetMinScore() *float32 {
	return s.MinScore
}

func (s *SearchMemoriesRequestBody) GetMinSimilarity() *float32 {
	return s.MinSimilarity
}

func (s *SearchMemoriesRequestBody) GetQuery() *string {
	return s.Query
}

func (s *SearchMemoriesRequestBody) GetScope() *SearchMemoriesRequestBodyScope {
	return s.Scope
}

func (s *SearchMemoriesRequestBody) GetTopK() *int32 {
	return s.TopK
}

func (s *SearchMemoriesRequestBody) SetEnableRerank(v bool) *SearchMemoriesRequestBody {
	s.EnableRerank = &v
	return s
}

func (s *SearchMemoriesRequestBody) SetMetadata(v map[string]*string) *SearchMemoriesRequestBody {
	s.Metadata = v
	return s
}

func (s *SearchMemoriesRequestBody) SetMinScore(v float32) *SearchMemoriesRequestBody {
	s.MinScore = &v
	return s
}

func (s *SearchMemoriesRequestBody) SetMinSimilarity(v float32) *SearchMemoriesRequestBody {
	s.MinSimilarity = &v
	return s
}

func (s *SearchMemoriesRequestBody) SetQuery(v string) *SearchMemoriesRequestBody {
	s.Query = &v
	return s
}

func (s *SearchMemoriesRequestBody) SetScope(v *SearchMemoriesRequestBodyScope) *SearchMemoriesRequestBody {
	s.Scope = v
	return s
}

func (s *SearchMemoriesRequestBody) SetTopK(v int32) *SearchMemoriesRequestBody {
	s.TopK = &v
	return s
}

func (s *SearchMemoriesRequestBody) Validate() error {
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchMemoriesRequestBodyScope struct {
	// The agent ID. The ID must be 1 to 128 bytes in length and encoded in UTF-8. The ID cannot contain a forward slash (/) or an asterisk (*). If the memory is not associated with an agent ID, this field is not returned.
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The session ID. The ID must be 1 to 128 bytes in length and encoded in UTF-8. The ID cannot contain a forward slash (/) or an asterisk (*). If the memory is not associated with a session ID, this field is not returned.
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The user ID. The ID must be 1 to 128 bytes in length and encoded in UTF-8. The ID cannot contain a forward slash (/) or an asterisk (*). If the memory is not associated with a user ID, this field is not returned.
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchMemoriesRequestBodyScope) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesRequestBodyScope) GoString() string {
	return s.String()
}

func (s *SearchMemoriesRequestBodyScope) GetAgentId() *string {
	return s.AgentId
}

func (s *SearchMemoriesRequestBodyScope) GetSessionId() *string {
	return s.SessionId
}

func (s *SearchMemoriesRequestBodyScope) GetUserId() *string {
	return s.UserId
}

func (s *SearchMemoriesRequestBodyScope) SetAgentId(v string) *SearchMemoriesRequestBodyScope {
	s.AgentId = &v
	return s
}

func (s *SearchMemoriesRequestBodyScope) SetSessionId(v string) *SearchMemoriesRequestBodyScope {
	s.SessionId = &v
	return s
}

func (s *SearchMemoriesRequestBodyScope) SetUserId(v string) *SearchMemoriesRequestBodyScope {
	s.UserId = &v
	return s
}

func (s *SearchMemoriesRequestBodyScope) Validate() error {
	return dara.Validate(s)
}
