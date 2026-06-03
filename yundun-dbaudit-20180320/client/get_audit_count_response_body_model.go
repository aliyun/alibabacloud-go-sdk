// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAuditCountResponseBody
	GetRequestId() *string
	SetRiskCount(v int64) *GetAuditCountResponseBody
	GetRiskCount() *int64
	SetSessionCount(v int64) *GetAuditCountResponseBody
	GetSessionCount() *int64
	SetSqlCount(v int64) *GetAuditCountResponseBody
	GetSqlCount() *int64
}

type GetAuditCountResponseBody struct {
	// example:
	//
	// E6A08A8A-F962-4FAD-AF0C-86B393E1F9C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1000
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// example:
	//
	// 2000
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// example:
	//
	// 100000
	SqlCount *int64 `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
}

func (s GetAuditCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuditCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuditCountResponseBody) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *GetAuditCountResponseBody) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *GetAuditCountResponseBody) GetSqlCount() *int64 {
	return s.SqlCount
}

func (s *GetAuditCountResponseBody) SetRequestId(v string) *GetAuditCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditCountResponseBody) SetRiskCount(v int64) *GetAuditCountResponseBody {
	s.RiskCount = &v
	return s
}

func (s *GetAuditCountResponseBody) SetSessionCount(v int64) *GetAuditCountResponseBody {
	s.SessionCount = &v
	return s
}

func (s *GetAuditCountResponseBody) SetSqlCount(v int64) *GetAuditCountResponseBody {
	s.SqlCount = &v
	return s
}

func (s *GetAuditCountResponseBody) Validate() error {
	return dara.Validate(s)
}
