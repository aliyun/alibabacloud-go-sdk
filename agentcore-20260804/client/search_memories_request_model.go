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
	// The string key-value metadata for exact matching. A maximum of 16 keys are supported. Each key can be up to 64 characters in length, and each value can be up to 1,024 characters in length.
	//
	// example:
	//
	// {"topic":"hiking"}
	Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The minimum relevance score. Default value: 0.2. Valid values: 0 to 1. If both minScore and minSimilarity are specified, the results are filtered based on the AND logic.
	//
	// example:
	//
	// 0.2
	MinScore *float32 `json:"minScore,omitempty" xml:"minScore,omitempty"`
	// The minimum similarity score. Default value: 0.3. Valid values: 0 to 1. If both minSimilarity and minScore are specified, the results are filtered based on the AND logic.
	//
	// example:
	//
	// 0.3
	MinSimilarity *float32 `json:"minSimilarity,omitempty" xml:"minSimilarity,omitempty"`
	// The query text for retrieval, up to 8,192 characters.
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
	// The list of memory content types. Supported types include semantic, skill, user_profile, and session_summary. If omitted or left as an empty array, no type restriction is applied.
	Types []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
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

func (s *SearchMemoriesRequestBody) GetTypes() []*string {
	return s.Types
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

func (s *SearchMemoriesRequestBody) SetTypes(v []*string) *SearchMemoriesRequestBody {
	s.Types = v
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
	// The agent ID. The length is limited to 1 to 128 UTF-8 bytes. The value cannot contain a forward slash (/) or an asterisk (*). This field is not returned if the memory is not associated with an agent ID.
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The session ID. The length is limited to 1 to 128 UTF-8 bytes. The value cannot contain a forward slash (/) or an asterisk (*). This field is not returned if the memory is not associated with a session ID.
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The business team ID defined by the caller. If provided, it cannot be empty and can be up to 128 UTF-8 bytes in length.
	//
	// example:
	//
	// team-001
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// The user ID. The length is limited to 1 to 128 UTF-8 bytes. The value cannot contain a forward slash (/) or an asterisk (*). This field is not returned if the memory is not associated with a user ID.
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

func (s *SearchMemoriesRequestBodyScope) GetTeamId() *string {
	return s.TeamId
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

func (s *SearchMemoriesRequestBodyScope) SetTeamId(v string) *SearchMemoriesRequestBodyScope {
	s.TeamId = &v
	return s
}

func (s *SearchMemoriesRequestBodyScope) SetUserId(v string) *SearchMemoriesRequestBodyScope {
	s.UserId = &v
	return s
}

func (s *SearchMemoriesRequestBodyScope) Validate() error {
	return dara.Validate(s)
}
