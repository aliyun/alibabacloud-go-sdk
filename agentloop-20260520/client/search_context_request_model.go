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
	SetThreshold(v float64) *SearchContextRequest
	GetThreshold() *float64
}

type SearchContextRequest struct {
	// The structured filter conditions. The key is the field name, and the value is the expected matching value.
	//
	// example:
	//
	// {"userId":"alice"}
	Filter map[string]interface{} `json:"filter,omitempty" xml:"filter,omitempty"`
	// Specifies whether to apply structured formatting to the returned results.
	//
	// example:
	//
	// true
	Formatted *bool `json:"formatted,omitempty" xml:"formatted,omitempty"`
	// example:
	//
	// false
	IncludeInactive *bool `json:"includeInactive,omitempty" xml:"includeInactive,omitempty"`
	// The maximum number of returned results (similarity Top-N).
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
	// 用户最近的偏好设置
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The retrieval options that control the retrieval strategy.
	//
	// example:
	//
	// semantic
	RetrievalOption *string                    `json:"retrievalOption,omitempty" xml:"retrievalOption,omitempty"`
	Scope           *SearchContextRequestScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The similarity threshold. Results with a similarity score lower than this value are filtered out. Valid values: 0 to 1.
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
	// example:
	//
	// sales-copilot
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// crm-service
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// run-001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
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
