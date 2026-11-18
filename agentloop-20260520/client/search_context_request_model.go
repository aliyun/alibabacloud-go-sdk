// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchContextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v map[string]interface{}) *SearchContextRequest
	GetFilter() map[string]interface{}
	SetFormatted(v bool) *SearchContextRequest
	GetFormatted() *bool
	SetIncludeInactive(v bool) *SearchContextRequest
	GetIncludeInactive() *bool
	SetLimit(v int32) *SearchContextRequest
	GetLimit() *int32
	SetQuery(v string) *SearchContextRequest
	GetQuery() *string
	SetRetrievalOption(v string) *SearchContextRequest
	GetRetrievalOption() *string
	SetScope(v *SearchContextRequestScope) *SearchContextRequest
	GetScope() *SearchContextRequestScope
	SetStrategyIds(v []*string) *SearchContextRequest
	GetStrategyIds() []*string
	SetThreshold(v float64) *SearchContextRequest
	GetThreshold() *float64
}

type SearchContextRequest struct {
	// The structured filter conditions. The experience store supports scene. When __agentloopControl.mode is set to direct, hybrid, or agentic, multi-strategy recall with auditing is used, and you can pass source, callId, application, fields, channels, and budget simultaneously. The management plane can additionally pass contentStatus=all, active, candidate, gray, or deprecated to perform content retrieval, which includes disabled or unreviewed content and is not affected by store-level runtime switches. The API key runtime entry prohibits contentStatus and retains the original governance thresholds. The maximum value of timeoutMs is 60000. If mode is not passed, the existing semantics of fields, channels, minQualityScore, expansion, and debug are retained. When query is *, you can use includeCandidates, cursor, and offset to browse experiences (management plane only).
	//
	// example:
	//
	// {"scene":"incident-recovery","__agentloopControl":{"fields":["when_to_use","summary"],"channels":["vector","fulltext"],"minQualityScore":0.6,"expansion":false,"debug":true}}
	Filter map[string]interface{} `json:"filter,omitempty" xml:"filter,omitempty"`
	// Specifies whether to apply structured formatting to the returned results.
	//
	// example:
	//
	// true
	Formatted *bool `json:"formatted,omitempty" xml:"formatted,omitempty"`
	// Specifies whether to include memories with sourceState set to inactive. Default value: false.
	//
	// example:
	//
	// false
	IncludeInactive *bool `json:"includeInactive,omitempty" xml:"includeInactive,omitempty"`
	// The maximum number of returned results (Top-N similarity).
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The retrieval query text. Natural language is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// Recent user preference settings
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The retrieval options. The experience store supports normal and reranker (deterministic reranking). The memory store uses its own retrieval strategy. For real-time recall in the experience store, the limit ranges from 1 to 20.
	//
	// example:
	//
	// normal
	RetrievalOption *string `json:"retrievalOption,omitempty" xml:"retrievalOption,omitempty"`
	// The recall scope. If this field conflicts with a field of the same name in filter, a 400 ScopeFilterConflict error is returned.
	Scope *SearchContextRequestScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// Filters results by extraction strategy IDs. The conditions in the list are combined by using the OR operator. If this parameter is not passed, all strategies are recalled. An empty list returns a 400 error. If the format is valid but no matching IDs are found, an empty result is returned. Historical removed strategy IDs are supported.
	//
	// example:
	//
	// ["preferences","profile"]
	StrategyIds []*string `json:"strategyIds,omitempty" xml:"strategyIds,omitempty" type:"Repeated"`
	// The similarity threshold. Results with a score lower than this value are filtered out. Valid values: 0 to 1.
	//
	// example:
	//
	// 0.5
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s SearchContextRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchContextRequest) GoString() string {
	return s.String()
}

func (s *SearchContextRequest) GetFilter() map[string]interface{} {
	return s.Filter
}

func (s *SearchContextRequest) GetFormatted() *bool {
	return s.Formatted
}

func (s *SearchContextRequest) GetIncludeInactive() *bool {
	return s.IncludeInactive
}

func (s *SearchContextRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *SearchContextRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchContextRequest) GetRetrievalOption() *string {
	return s.RetrievalOption
}

func (s *SearchContextRequest) GetScope() *SearchContextRequestScope {
	return s.Scope
}

func (s *SearchContextRequest) GetStrategyIds() []*string {
	return s.StrategyIds
}

func (s *SearchContextRequest) GetThreshold() *float64 {
	return s.Threshold
}

func (s *SearchContextRequest) SetFilter(v map[string]interface{}) *SearchContextRequest {
	s.Filter = v
	return s
}

func (s *SearchContextRequest) SetFormatted(v bool) *SearchContextRequest {
	s.Formatted = &v
	return s
}

func (s *SearchContextRequest) SetIncludeInactive(v bool) *SearchContextRequest {
	s.IncludeInactive = &v
	return s
}

func (s *SearchContextRequest) SetLimit(v int32) *SearchContextRequest {
	s.Limit = &v
	return s
}

func (s *SearchContextRequest) SetQuery(v string) *SearchContextRequest {
	s.Query = &v
	return s
}

func (s *SearchContextRequest) SetRetrievalOption(v string) *SearchContextRequest {
	s.RetrievalOption = &v
	return s
}

func (s *SearchContextRequest) SetScope(v *SearchContextRequestScope) *SearchContextRequest {
	s.Scope = v
	return s
}

func (s *SearchContextRequest) SetStrategyIds(v []*string) *SearchContextRequest {
	s.StrategyIds = v
	return s
}

func (s *SearchContextRequest) SetThreshold(v float64) *SearchContextRequest {
	s.Threshold = &v
	return s
}

func (s *SearchContextRequest) Validate() error {
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchContextRequestScope struct {
	// The ID of the agent.
	//
	// example:
	//
	// sales-copilot
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// crm-service
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The ID of the run.
	//
	// example:
	//
	// run-001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// u-10001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchContextRequestScope) String() string {
	return dara.Prettify(s)
}

func (s SearchContextRequestScope) GoString() string {
	return s.String()
}

func (s *SearchContextRequestScope) GetAgentId() *string {
	return s.AgentId
}

func (s *SearchContextRequestScope) GetAppId() *string {
	return s.AppId
}

func (s *SearchContextRequestScope) GetRunId() *string {
	return s.RunId
}

func (s *SearchContextRequestScope) GetUserId() *string {
	return s.UserId
}

func (s *SearchContextRequestScope) SetAgentId(v string) *SearchContextRequestScope {
	s.AgentId = &v
	return s
}

func (s *SearchContextRequestScope) SetAppId(v string) *SearchContextRequestScope {
	s.AppId = &v
	return s
}

func (s *SearchContextRequestScope) SetRunId(v string) *SearchContextRequestScope {
	s.RunId = &v
	return s
}

func (s *SearchContextRequestScope) SetUserId(v string) *SearchContextRequestScope {
	s.UserId = &v
	return s
}

func (s *SearchContextRequestScope) Validate() error {
	return dara.Validate(s)
}
