// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompletedAt(v int64) *GetQueryResponseBody
	GetCompletedAt() *int64
	SetCreatedAt(v int64) *GetQueryResponseBody
	GetCreatedAt() *int64
	SetQueryId(v string) *GetQueryResponseBody
	GetQueryId() *string
	SetResults(v []*StatementResult) *GetQueryResponseBody
	GetResults() []*StatementResult
	SetSql(v string) *GetQueryResponseBody
	GetSql() *string
	SetStatus(v string) *GetQueryResponseBody
	GetStatus() *string
	SetTier(v string) *GetQueryResponseBody
	GetTier() *string
}

type GetQueryResponseBody struct {
	CompletedAt *int64             `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	CreatedAt   *int64             `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	QueryId     *string            `json:"queryId,omitempty" xml:"queryId,omitempty"`
	Results     []*StatementResult `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	Sql         *string            `json:"sql,omitempty" xml:"sql,omitempty"`
	Status      *string            `json:"status,omitempty" xml:"status,omitempty"`
	Tier        *string            `json:"tier,omitempty" xml:"tier,omitempty"`
}

func (s GetQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueryResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryResponseBody) GetCompletedAt() *int64 {
	return s.CompletedAt
}

func (s *GetQueryResponseBody) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetQueryResponseBody) GetQueryId() *string {
	return s.QueryId
}

func (s *GetQueryResponseBody) GetResults() []*StatementResult {
	return s.Results
}

func (s *GetQueryResponseBody) GetSql() *string {
	return s.Sql
}

func (s *GetQueryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetQueryResponseBody) GetTier() *string {
	return s.Tier
}

func (s *GetQueryResponseBody) SetCompletedAt(v int64) *GetQueryResponseBody {
	s.CompletedAt = &v
	return s
}

func (s *GetQueryResponseBody) SetCreatedAt(v int64) *GetQueryResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetQueryResponseBody) SetQueryId(v string) *GetQueryResponseBody {
	s.QueryId = &v
	return s
}

func (s *GetQueryResponseBody) SetResults(v []*StatementResult) *GetQueryResponseBody {
	s.Results = v
	return s
}

func (s *GetQueryResponseBody) SetSql(v string) *GetQueryResponseBody {
	s.Sql = &v
	return s
}

func (s *GetQueryResponseBody) SetStatus(v string) *GetQueryResponseBody {
	s.Status = &v
	return s
}

func (s *GetQueryResponseBody) SetTier(v string) *GetQueryResponseBody {
	s.Tier = &v
	return s
}

func (s *GetQueryResponseBody) Validate() error {
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
