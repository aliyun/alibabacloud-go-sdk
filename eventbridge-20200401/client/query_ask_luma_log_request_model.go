// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAskLumaLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfter(v string) *QueryAskLumaLogRequest
	GetAfter() *string
	SetAgentName(v string) *QueryAskLumaLogRequest
	GetAgentName() *string
	SetLimit(v int32) *QueryAskLumaLogRequest
	GetLimit() *int32
}

type QueryAskLumaLogRequest struct {
	// The cursor. Set this parameter to the messageId of the last entry on the previous page.
	//
	// example:
	//
	// eyJ0cyI6MTcxN...
	After *string `json:"After,omitempty" xml:"After,omitempty"`
	// The agent name. If this parameter is not specified, logs of all agents are queried.
	//
	// example:
	//
	// demo-luma-agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The number of entries to return. Default value: 20. Maximum value: 50.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s QueryAskLumaLogRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAskLumaLogRequest) GoString() string {
	return s.String()
}

func (s *QueryAskLumaLogRequest) GetAfter() *string {
	return s.After
}

func (s *QueryAskLumaLogRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *QueryAskLumaLogRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *QueryAskLumaLogRequest) SetAfter(v string) *QueryAskLumaLogRequest {
	s.After = &v
	return s
}

func (s *QueryAskLumaLogRequest) SetAgentName(v string) *QueryAskLumaLogRequest {
	s.AgentName = &v
	return s
}

func (s *QueryAskLumaLogRequest) SetLimit(v int32) *QueryAskLumaLogRequest {
	s.Limit = &v
	return s
}

func (s *QueryAskLumaLogRequest) Validate() error {
	return dara.Validate(s)
}
