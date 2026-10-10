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
	SetRecallEventId(v string) *SearchContextResponseBody
	GetRecallEventId() *string
	SetRequestId(v string) *SearchContextResponseBody
	GetRequestId() *string
	SetResults(v []map[string]interface{}) *SearchContextResponseBody
	GetResults() []map[string]interface{}
}

type SearchContextResponseBody struct {
	// example:
	//
	// ok
	AuditStatus *string `json:"auditStatus,omitempty" xml:"auditStatus,omitempty"`
	// example:
	//
	// 0190f1c2-7d3e-7a1b-9c4d-2e5f6a7b8c9d
	RecallEventId *string `json:"recallEventId,omitempty" xml:"recallEventId,omitempty"`
	// The request ID. You can use this ID to locate and troubleshoot issues.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of retrieval results, sorted by similarity in descending order.
	Results []map[string]interface{} `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
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

func (s *SearchContextResponseBody) GetRecallEventId() *string {
	return s.RecallEventId
}

func (s *SearchContextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchContextResponseBody) GetResults() []map[string]interface{} {
	return s.Results
}

func (s *SearchContextResponseBody) SetAuditStatus(v string) *SearchContextResponseBody {
	s.AuditStatus = &v
	return s
}

func (s *SearchContextResponseBody) SetRecallEventId(v string) *SearchContextResponseBody {
	s.RecallEventId = &v
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

func (s *SearchContextResponseBody) Validate() error {
	return dara.Validate(s)
}
