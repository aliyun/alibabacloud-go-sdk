// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAgentDataSemanticsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *SaveAgentDataSemanticsShrinkRequest
	GetAgentName() *string
	SetExamplesShrink(v string) *SaveAgentDataSemanticsShrinkRequest
	GetExamplesShrink() *string
	SetJoinsShrink(v string) *SaveAgentDataSemanticsShrinkRequest
	GetJoinsShrink() *string
	SetMetricsShrink(v string) *SaveAgentDataSemanticsShrinkRequest
	GetMetricsShrink() *string
	SetTextShrink(v string) *SaveAgentDataSemanticsShrinkRequest
	GetTextShrink() *string
}

type SaveAgentDataSemanticsShrinkRequest struct {
	// The name of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// bakehouse_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The SQL example knowledge. If this parameter is specified, the current content is saved. If this parameter is not specified, the existing content is cleared. A maximum of 50 entries are supported, and the maximum size of each knowledge category is 16 KB.
	ExamplesShrink *string `json:"Examples,omitempty" xml:"Examples,omitempty"`
	// The data association knowledge. If this parameter is specified, the current content is saved. If this parameter is not specified, the existing content is cleared. A maximum of 100 entries are supported, and the maximum size of each knowledge category is 16 KB.
	JoinsShrink *string `json:"Joins,omitempty" xml:"Joins,omitempty"`
	// The SQL expression knowledge. If this parameter is specified, the current content is saved. If this parameter is not specified, the existing content is cleared. A maximum of 100 entries are supported, and the maximum size of each knowledge category is 16 KB.
	MetricsShrink *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The text knowledge in Markdown format. If this parameter is specified, the current content is saved. If this parameter is not specified, the existing content is cleared. The maximum size of each knowledge category is 16 KB.
	TextShrink *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SaveAgentDataSemanticsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveAgentDataSemanticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveAgentDataSemanticsShrinkRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *SaveAgentDataSemanticsShrinkRequest) GetExamplesShrink() *string {
	return s.ExamplesShrink
}

func (s *SaveAgentDataSemanticsShrinkRequest) GetJoinsShrink() *string {
	return s.JoinsShrink
}

func (s *SaveAgentDataSemanticsShrinkRequest) GetMetricsShrink() *string {
	return s.MetricsShrink
}

func (s *SaveAgentDataSemanticsShrinkRequest) GetTextShrink() *string {
	return s.TextShrink
}

func (s *SaveAgentDataSemanticsShrinkRequest) SetAgentName(v string) *SaveAgentDataSemanticsShrinkRequest {
	s.AgentName = &v
	return s
}

func (s *SaveAgentDataSemanticsShrinkRequest) SetExamplesShrink(v string) *SaveAgentDataSemanticsShrinkRequest {
	s.ExamplesShrink = &v
	return s
}

func (s *SaveAgentDataSemanticsShrinkRequest) SetJoinsShrink(v string) *SaveAgentDataSemanticsShrinkRequest {
	s.JoinsShrink = &v
	return s
}

func (s *SaveAgentDataSemanticsShrinkRequest) SetMetricsShrink(v string) *SaveAgentDataSemanticsShrinkRequest {
	s.MetricsShrink = &v
	return s
}

func (s *SaveAgentDataSemanticsShrinkRequest) SetTextShrink(v string) *SaveAgentDataSemanticsShrinkRequest {
	s.TextShrink = &v
	return s
}

func (s *SaveAgentDataSemanticsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
