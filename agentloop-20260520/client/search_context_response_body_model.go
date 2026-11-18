// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchContextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuditStatus(v string) *SearchContextResponseBody
	GetAuditStatus() *string
	SetCallId(v string) *SearchContextResponseBody
	GetCallId() *string
	SetDiagnostics(v map[string]interface{}) *SearchContextResponseBody
	GetDiagnostics() map[string]interface{}
	SetMode(v string) *SearchContextResponseBody
	GetMode() *string
	SetRecallEventId(v string) *SearchContextResponseBody
	GetRecallEventId() *string
	SetRelevance(v []map[string]interface{}) *SearchContextResponseBody
	GetRelevance() []map[string]interface{}
	SetRequestId(v string) *SearchContextResponseBody
	GetRequestId() *string
	SetResults(v []map[string]interface{}) *SearchContextResponseBody
	GetResults() []map[string]interface{}
	SetRewrittenQueries(v []*string) *SearchContextResponseBody
	GetRewrittenQueries() []*string
	SetStatus(v string) *SearchContextResponseBody
	GetStatus() *string
	SetStrategyVersion(v string) *SearchContextResponseBody
	GetStrategyVersion() *string
}

type SearchContextResponseBody struct {
	// The audit write status. Valid values: ok and degraded.
	//
	// example:
	//
	// ok
	AuditStatus *string `json:"auditStatus,omitempty" xml:"auditStatus,omitempty"`
	// The call ID for multi-strategy recall in the experience store.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	CallId *string `json:"callId,omitempty" xml:"callId,omitempty"`
	// The diagnostics information for the current experience recall, including duration, model call volume, and filtering. The actual diagnostics items are determined by the running mode.
	//
	// example:
	//
	// {
	//
	// 	"durationMs": 386.5,
	//
	// 	"modelCalls": 2,
	//
	// 	"relevanceValidation": "semantic",
	//
	// 	"candidateCount": 1,
	//
	// 	"filtering": {
	//
	// 		"model_not_relevant": 0
	//
	// 	}
	//
	// }
	Diagnostics map[string]interface{} `json:"diagnostics,omitempty" xml:"diagnostics,omitempty"`
	// The experience recall mode actually used for the current request. Valid values: direct, agentic, and hybrid.
	//
	// example:
	//
	// agentic
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The ID of the recall event. You can use this ID to call GetContextStoreRecall to view audit details.
	//
	// example:
	//
	// 0190f1c2-7d3e-7a1b-9c4d-2e5f6a7b8c9d
	RecallEventId *string `json:"recallEventId,omitempty" xml:"recallEventId,omitempty"`
	// The list of relevance judgment results for candidate experiences.
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "contextId": "a:1",
	//
	//     "judgment": "relevant",
	//
	//     "reason": "This experience explains how to skip tasks with unchanged inputs, which is relevant to the search query."
	//
	//   }
	//
	// ]
	Relevance []map[string]interface{} `json:"relevance,omitempty" xml:"relevance,omitempty" type:"Repeated"`
	// The ID of the request. You can use this ID to locate the request when troubleshooting issues.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of retrieval results, sorted by similarity in descending order. Each item contains contextId, context, metadata, score, rank, revision, strategyVersion, and strategyId.
	//
	// example:
	//
	// [{"contextId":"01a1cf08-c368-7c83******","strategyId":"default","context":"The assistant recommends steamed seafood because it is fresh and not spicy, which aligns with the user\\"s preference for seafood and constraint against spicy food.","metadata":{"contextType":"memory","memoryType":"long","userId":"test_user_002","agentId":"95273000000000","appId":"mm_480d961******","runId":"jr-80ded1d6******","strategyId":"default","category":"recommendation","kind":"state","sourceKind":"api","sourceId":"addctx:75395808-cb81-54cd******","revision":1,"reviewStatus":"unreviewed","sourceState":"active","factKey":"assistant:recommendation:dish******","sessionId":"test_session_001"},"score":0.6617179,"rank":1,"revision":1,"strategyVersion":1}]
	Results []map[string]interface{} `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// The list of rewritten queries generated based on the original retrieval query. The original query is not included.
	//
	// example:
	//
	// ["How to avoid repeatedly executing unchanged inputs"]
	RewrittenQueries []*string `json:"rewrittenQueries,omitempty" xml:"rewrittenQueries,omitempty" type:"Repeated"`
	// The processing status of the current experience recall call. Valid values: success, degraded, failed, and error.
	//
	// example:
	//
	// success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The implementation version of the strategy used for the current experience recall. This is used for auditing and result tracing.
	//
	// example:
	//
	// experience-recall-20261031-v5
	StrategyVersion *string `json:"strategyVersion,omitempty" xml:"strategyVersion,omitempty"`
}

func (s SearchContextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchContextResponseBody) GoString() string {
	return s.String()
}

func (s *SearchContextResponseBody) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *SearchContextResponseBody) GetCallId() *string {
	return s.CallId
}

func (s *SearchContextResponseBody) GetDiagnostics() map[string]interface{} {
	return s.Diagnostics
}

func (s *SearchContextResponseBody) GetMode() *string {
	return s.Mode
}

func (s *SearchContextResponseBody) GetRecallEventId() *string {
	return s.RecallEventId
}

func (s *SearchContextResponseBody) GetRelevance() []map[string]interface{} {
	return s.Relevance
}

func (s *SearchContextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchContextResponseBody) GetResults() []map[string]interface{} {
	return s.Results
}

func (s *SearchContextResponseBody) GetRewrittenQueries() []*string {
	return s.RewrittenQueries
}

func (s *SearchContextResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SearchContextResponseBody) GetStrategyVersion() *string {
	return s.StrategyVersion
}

func (s *SearchContextResponseBody) SetAuditStatus(v string) *SearchContextResponseBody {
	s.AuditStatus = &v
	return s
}

func (s *SearchContextResponseBody) SetCallId(v string) *SearchContextResponseBody {
	s.CallId = &v
	return s
}

func (s *SearchContextResponseBody) SetDiagnostics(v map[string]interface{}) *SearchContextResponseBody {
	s.Diagnostics = v
	return s
}

func (s *SearchContextResponseBody) SetMode(v string) *SearchContextResponseBody {
	s.Mode = &v
	return s
}

func (s *SearchContextResponseBody) SetRecallEventId(v string) *SearchContextResponseBody {
	s.RecallEventId = &v
	return s
}

func (s *SearchContextResponseBody) SetRelevance(v []map[string]interface{}) *SearchContextResponseBody {
	s.Relevance = v
	return s
}

func (s *SearchContextResponseBody) SetRequestId(v string) *SearchContextResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchContextResponseBody) SetResults(v []map[string]interface{}) *SearchContextResponseBody {
	s.Results = v
	return s
}

func (s *SearchContextResponseBody) SetRewrittenQueries(v []*string) *SearchContextResponseBody {
	s.RewrittenQueries = v
	return s
}

func (s *SearchContextResponseBody) SetStatus(v string) *SearchContextResponseBody {
	s.Status = &v
	return s
}

func (s *SearchContextResponseBody) SetStrategyVersion(v string) *SearchContextResponseBody {
	s.StrategyVersion = &v
	return s
}

func (s *SearchContextResponseBody) Validate() error {
	return dara.Validate(s)
}
