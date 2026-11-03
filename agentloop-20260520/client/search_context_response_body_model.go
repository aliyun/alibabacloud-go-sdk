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
	// The call ID for the multi-strategy recall of the experience store.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	CallId *string `json:"callId,omitempty" xml:"callId,omitempty"`
	// The diagnostics information for the current experience recall, including duration, model call volume, and filtering. The actual diagnostics items are determined by the running mode.
	Diagnostics map[string]interface{} `json:"diagnostics,omitempty" xml:"diagnostics,omitempty"`
	// The actual experience recall mode used for the current request. Valid values: direct, agentic, and hybrid.
	//
	// example:
	//
	// agentic
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The ID of the recall event. You can use this ID to call the GetContextStoreRecall operation to view audit details.
	//
	// example:
	//
	// 0190f1c2-7d3e-7a1b-9c4d-2e5f6a7b8c9d
	RecallEventId *string `json:"recallEventId,omitempty" xml:"recallEventId,omitempty"`
	// The list of relevance judgment results for candidate experiences.
	Relevance []map[string]interface{} `json:"relevance,omitempty" xml:"relevance,omitempty" type:"Repeated"`
	// The ID of the request. You can use this ID to locate the request for troubleshooting.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of retrieval results, sorted by similarity in descending order.
	Results []map[string]interface{} `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// The list of rewritten queries generated based on the original retrieval query. The original query is not included.
	RewrittenQueries []*string `json:"rewrittenQueries,omitempty" xml:"rewrittenQueries,omitempty" type:"Repeated"`
	// The processing status of the current experience recall call. Valid values: success, degraded, failed, and error.
	//
	// example:
	//
	// success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The strategy implementation version used for the current experience recall. This is used for auditing and result tracing.
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
