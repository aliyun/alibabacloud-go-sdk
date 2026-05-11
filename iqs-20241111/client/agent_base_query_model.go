// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentBaseQuery interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *AgentBaseQuery
	GetQuery() *string
}

type AgentBaseQuery struct {
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s AgentBaseQuery) String() string {
	return dara.Prettify(s)
}

func (s AgentBaseQuery) GoString() string {
	return s.String()
}

func (s *AgentBaseQuery) GetQuery() *string {
	return s.Query
}

func (s *AgentBaseQuery) SetQuery(v string) *AgentBaseQuery {
	s.Query = &v
	return s
}

func (s *AgentBaseQuery) Validate() error {
	return dara.Validate(s)
}
