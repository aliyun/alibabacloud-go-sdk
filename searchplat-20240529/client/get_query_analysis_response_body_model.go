// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetQueryAnalysisResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetQueryAnalysisResponseBody
	GetRequestId() *string
	SetResult(v *GetQueryAnalysisResponseBodyResult) *GetQueryAnalysisResponseBody
	GetResult() *GetQueryAnalysisResponseBodyResult
	SetUsage(v *GetQueryAnalysisResponseBodyUsage) *GetQueryAnalysisResponseBody
	GetUsage() *GetQueryAnalysisResponseBodyUsage
}

type GetQueryAnalysisResponseBody struct {
	Latency   *int32                              `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                             `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetQueryAnalysisResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetQueryAnalysisResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetQueryAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueryAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryAnalysisResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetQueryAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueryAnalysisResponseBody) GetResult() *GetQueryAnalysisResponseBodyResult {
	return s.Result
}

func (s *GetQueryAnalysisResponseBody) GetUsage() *GetQueryAnalysisResponseBodyUsage {
	return s.Usage
}

func (s *GetQueryAnalysisResponseBody) SetLatency(v int32) *GetQueryAnalysisResponseBody {
	s.Latency = &v
	return s
}

func (s *GetQueryAnalysisResponseBody) SetRequestId(v string) *GetQueryAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryAnalysisResponseBody) SetResult(v *GetQueryAnalysisResponseBodyResult) *GetQueryAnalysisResponseBody {
	s.Result = v
	return s
}

func (s *GetQueryAnalysisResponseBody) SetUsage(v *GetQueryAnalysisResponseBodyUsage) *GetQueryAnalysisResponseBody {
	s.Usage = v
	return s
}

func (s *GetQueryAnalysisResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQueryAnalysisResponseBodyResult struct {
	Intent  *string                `json:"intent,omitempty" xml:"intent,omitempty"`
	Queries []*string              `json:"queries,omitempty" xml:"queries,omitempty" type:"Repeated"`
	Query   *string                `json:"query,omitempty" xml:"query,omitempty"`
	Sql     map[string]interface{} `json:"sql,omitempty" xml:"sql,omitempty"`
}

func (s GetQueryAnalysisResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetQueryAnalysisResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetQueryAnalysisResponseBodyResult) GetIntent() *string {
	return s.Intent
}

func (s *GetQueryAnalysisResponseBodyResult) GetQueries() []*string {
	return s.Queries
}

func (s *GetQueryAnalysisResponseBodyResult) GetQuery() *string {
	return s.Query
}

func (s *GetQueryAnalysisResponseBodyResult) GetSql() map[string]interface{} {
	return s.Sql
}

func (s *GetQueryAnalysisResponseBodyResult) SetIntent(v string) *GetQueryAnalysisResponseBodyResult {
	s.Intent = &v
	return s
}

func (s *GetQueryAnalysisResponseBodyResult) SetQueries(v []*string) *GetQueryAnalysisResponseBodyResult {
	s.Queries = v
	return s
}

func (s *GetQueryAnalysisResponseBodyResult) SetQuery(v string) *GetQueryAnalysisResponseBodyResult {
	s.Query = &v
	return s
}

func (s *GetQueryAnalysisResponseBodyResult) SetSql(v map[string]interface{}) *GetQueryAnalysisResponseBodyResult {
	s.Sql = v
	return s
}

func (s *GetQueryAnalysisResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetQueryAnalysisResponseBodyUsage struct {
	InputTokens  *int64 `json:"input_tokens,omitempty" xml:"input_tokens,omitempty"`
	OutputTokens *int64 `json:"output_tokens,omitempty" xml:"output_tokens,omitempty"`
	TotalTokens  *int64 `json:"total_tokens,omitempty" xml:"total_tokens,omitempty"`
}

func (s GetQueryAnalysisResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetQueryAnalysisResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetQueryAnalysisResponseBodyUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetQueryAnalysisResponseBodyUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetQueryAnalysisResponseBodyUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetQueryAnalysisResponseBodyUsage) SetInputTokens(v int64) *GetQueryAnalysisResponseBodyUsage {
	s.InputTokens = &v
	return s
}

func (s *GetQueryAnalysisResponseBodyUsage) SetOutputTokens(v int64) *GetQueryAnalysisResponseBodyUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetQueryAnalysisResponseBodyUsage) SetTotalTokens(v int64) *GetQueryAnalysisResponseBodyUsage {
	s.TotalTokens = &v
	return s
}

func (s *GetQueryAnalysisResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
