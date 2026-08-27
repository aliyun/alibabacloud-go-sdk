// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentNames(v []*string) *ListAuthorizedAgentsResponseBody
	GetAgentNames() []*string
	SetCode(v string) *ListAuthorizedAgentsResponseBody
	GetCode() *string
	SetMessage(v string) *ListAuthorizedAgentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAuthorizedAgentsResponseBody
	GetRequestId() *string
}

type ListAuthorizedAgentsResponseBody struct {
	// The agent names.
	//
	// example:
	//
	// string_value
	AgentNames []*string `json:"agentNames,omitempty" xml:"agentNames,omitempty" type:"Repeated"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAuthorizedAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAgentsResponseBody) GetAgentNames() []*string {
	return s.AgentNames
}

func (s *ListAuthorizedAgentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAuthorizedAgentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAuthorizedAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizedAgentsResponseBody) SetAgentNames(v []*string) *ListAuthorizedAgentsResponseBody {
	s.AgentNames = v
	return s
}

func (s *ListAuthorizedAgentsResponseBody) SetCode(v string) *ListAuthorizedAgentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAuthorizedAgentsResponseBody) SetMessage(v string) *ListAuthorizedAgentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAuthorizedAgentsResponseBody) SetRequestId(v string) *ListAuthorizedAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedAgentsResponseBody) Validate() error {
	return dara.Validate(s)
}
