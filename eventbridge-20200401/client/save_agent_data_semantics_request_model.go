// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAgentDataSemanticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *SaveAgentDataSemanticsRequest
	GetAgentName() *string
	SetExamples(v []*AgentDataSemanticsExample) *SaveAgentDataSemanticsRequest
	GetExamples() []*AgentDataSemanticsExample
	SetJoins(v []*AgentDataSemanticsJoin) *SaveAgentDataSemanticsRequest
	GetJoins() []*AgentDataSemanticsJoin
	SetMetrics(v []*AgentDataSemanticsMetric) *SaveAgentDataSemanticsRequest
	GetMetrics() []*AgentDataSemanticsMetric
	SetText(v *AgentDataSemanticsText) *SaveAgentDataSemanticsRequest
	GetText() *AgentDataSemanticsText
}

type SaveAgentDataSemanticsRequest struct {
	// The name of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// bakehouse_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The SQL example knowledge. If this parameter is specified, the current content is saved. If this parameter is not specified, the existing content is cleared. A maximum of 50 entries are supported, and the maximum size of each knowledge category is 16 KB.
	Examples []*AgentDataSemanticsExample `json:"Examples,omitempty" xml:"Examples,omitempty" type:"Repeated"`
	// The data association knowledge. If this parameter is specified, the current content is saved. If this parameter is not specified, the existing content is cleared. A maximum of 100 entries are supported, and the maximum size of each knowledge category is 16 KB.
	Joins []*AgentDataSemanticsJoin `json:"Joins,omitempty" xml:"Joins,omitempty" type:"Repeated"`
	// The SQL expression knowledge. If this parameter is specified, the current content is saved. If this parameter is not specified, the existing content is cleared. A maximum of 100 entries are supported, and the maximum size of each knowledge category is 16 KB.
	Metrics []*AgentDataSemanticsMetric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The text knowledge in Markdown format. If this parameter is specified, the current content is saved. If this parameter is not specified, the existing content is cleared. The maximum size of each knowledge category is 16 KB.
	Text *AgentDataSemanticsText `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SaveAgentDataSemanticsRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveAgentDataSemanticsRequest) GoString() string {
	return s.String()
}

func (s *SaveAgentDataSemanticsRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *SaveAgentDataSemanticsRequest) GetExamples() []*AgentDataSemanticsExample {
	return s.Examples
}

func (s *SaveAgentDataSemanticsRequest) GetJoins() []*AgentDataSemanticsJoin {
	return s.Joins
}

func (s *SaveAgentDataSemanticsRequest) GetMetrics() []*AgentDataSemanticsMetric {
	return s.Metrics
}

func (s *SaveAgentDataSemanticsRequest) GetText() *AgentDataSemanticsText {
	return s.Text
}

func (s *SaveAgentDataSemanticsRequest) SetAgentName(v string) *SaveAgentDataSemanticsRequest {
	s.AgentName = &v
	return s
}

func (s *SaveAgentDataSemanticsRequest) SetExamples(v []*AgentDataSemanticsExample) *SaveAgentDataSemanticsRequest {
	s.Examples = v
	return s
}

func (s *SaveAgentDataSemanticsRequest) SetJoins(v []*AgentDataSemanticsJoin) *SaveAgentDataSemanticsRequest {
	s.Joins = v
	return s
}

func (s *SaveAgentDataSemanticsRequest) SetMetrics(v []*AgentDataSemanticsMetric) *SaveAgentDataSemanticsRequest {
	s.Metrics = v
	return s
}

func (s *SaveAgentDataSemanticsRequest) SetText(v *AgentDataSemanticsText) *SaveAgentDataSemanticsRequest {
	s.Text = v
	return s
}

func (s *SaveAgentDataSemanticsRequest) Validate() error {
	if s.Examples != nil {
		for _, item := range s.Examples {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Joins != nil {
		for _, item := range s.Joins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Text != nil {
		if err := s.Text.Validate(); err != nil {
			return err
		}
	}
	return nil
}
