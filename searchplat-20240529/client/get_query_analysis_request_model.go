// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctions(v []*GetQueryAnalysisRequestFunctions) *GetQueryAnalysisRequest
	GetFunctions() []*GetQueryAnalysisRequestFunctions
	SetHistory(v []*GetQueryAnalysisRequestHistory) *GetQueryAnalysisRequest
	GetHistory() []*GetQueryAnalysisRequestHistory
	SetQuery(v string) *GetQueryAnalysisRequest
	GetQuery() *string
}

type GetQueryAnalysisRequest struct {
	Functions []*GetQueryAnalysisRequestFunctions `json:"functions,omitempty" xml:"functions,omitempty" type:"Repeated"`
	History   []*GetQueryAnalysisRequestHistory   `json:"history,omitempty" xml:"history,omitempty" type:"Repeated"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s GetQueryAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueryAnalysisRequest) GoString() string {
	return s.String()
}

func (s *GetQueryAnalysisRequest) GetFunctions() []*GetQueryAnalysisRequestFunctions {
	return s.Functions
}

func (s *GetQueryAnalysisRequest) GetHistory() []*GetQueryAnalysisRequestHistory {
	return s.History
}

func (s *GetQueryAnalysisRequest) GetQuery() *string {
	return s.Query
}

func (s *GetQueryAnalysisRequest) SetFunctions(v []*GetQueryAnalysisRequestFunctions) *GetQueryAnalysisRequest {
	s.Functions = v
	return s
}

func (s *GetQueryAnalysisRequest) SetHistory(v []*GetQueryAnalysisRequestHistory) *GetQueryAnalysisRequest {
	s.History = v
	return s
}

func (s *GetQueryAnalysisRequest) SetQuery(v string) *GetQueryAnalysisRequest {
	s.Query = &v
	return s
}

func (s *GetQueryAnalysisRequest) Validate() error {
	if s.Functions != nil {
		for _, item := range s.Functions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.History != nil {
		for _, item := range s.History {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQueryAnalysisRequestFunctions struct {
	Name       *string                `json:"name,omitempty" xml:"name,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s GetQueryAnalysisRequestFunctions) String() string {
	return dara.Prettify(s)
}

func (s GetQueryAnalysisRequestFunctions) GoString() string {
	return s.String()
}

func (s *GetQueryAnalysisRequestFunctions) GetName() *string {
	return s.Name
}

func (s *GetQueryAnalysisRequestFunctions) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *GetQueryAnalysisRequestFunctions) SetName(v string) *GetQueryAnalysisRequestFunctions {
	s.Name = &v
	return s
}

func (s *GetQueryAnalysisRequestFunctions) SetParameters(v map[string]interface{}) *GetQueryAnalysisRequestFunctions {
	s.Parameters = v
	return s
}

func (s *GetQueryAnalysisRequestFunctions) Validate() error {
	return dara.Validate(s)
}

type GetQueryAnalysisRequestHistory struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetQueryAnalysisRequestHistory) String() string {
	return dara.Prettify(s)
}

func (s GetQueryAnalysisRequestHistory) GoString() string {
	return s.String()
}

func (s *GetQueryAnalysisRequestHistory) GetContent() *string {
	return s.Content
}

func (s *GetQueryAnalysisRequestHistory) GetRole() *string {
	return s.Role
}

func (s *GetQueryAnalysisRequestHistory) SetContent(v string) *GetQueryAnalysisRequestHistory {
	s.Content = &v
	return s
}

func (s *GetQueryAnalysisRequestHistory) SetRole(v string) *GetQueryAnalysisRequestHistory {
	s.Role = &v
	return s
}

func (s *GetQueryAnalysisRequestHistory) Validate() error {
	return dara.Validate(s)
}
