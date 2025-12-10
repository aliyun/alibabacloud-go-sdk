// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListAgentRequest
	GetCount() *int32
	SetMarker(v string) *ListAgentRequest
	GetMarker() *string
}

type ListAgentRequest struct {
	// Specifies the number of agents to be returned.\\
	//
	// Valid values: 0 - 1000.\\
	//
	// Default value: 1000.
	//
	// example:
	//
	// 100
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The marker after which the agents are listed.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// test_agent
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRequest) GoString() string {
	return s.String()
}

func (s *ListAgentRequest) GetCount() *int32 {
	return s.Count
}

func (s *ListAgentRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListAgentRequest) SetCount(v int32) *ListAgentRequest {
	s.Count = &v
	return s
}

func (s *ListAgentRequest) SetMarker(v string) *ListAgentRequest {
	s.Marker = &v
	return s
}

func (s *ListAgentRequest) Validate() error {
	return dara.Validate(s)
}
