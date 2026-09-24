// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeEsAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvokeEsAgentResponseBody
	GetCode() *string
	SetData(v map[string]*string) *InvokeEsAgentResponseBody
	GetData() map[string]*string
	SetMessage(v string) *InvokeEsAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvokeEsAgentResponseBody
	GetRequestId() *string
}

type InvokeEsAgentResponseBody struct {
	// The status code. A value of 200 indicates a successful call. For non-200 values, the message field contains the error description.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The JSON-RPC 2.0 response body. data.result contains the actual return content of the called ACP method. data.id is the id passed in the request. data.jsonrpc is fixed to 2.0. data.timestamp is the UNIX timestamp in milliseconds when the response was generated.
	Data map[string]*string `json:"data,omitempty" xml:"data,omitempty"`
	// The error description. The value is null when the call is successful. A specific error message is returned when the call fails.
	//
	// example:
	//
	// Agent subscription is inactive or expired
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1df6d614-96db-41bd-b8b6-ab061a1724ca
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s InvokeEsAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeEsAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeEsAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvokeEsAgentResponseBody) GetData() map[string]*string {
	return s.Data
}

func (s *InvokeEsAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvokeEsAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeEsAgentResponseBody) SetCode(v string) *InvokeEsAgentResponseBody {
	s.Code = &v
	return s
}

func (s *InvokeEsAgentResponseBody) SetData(v map[string]*string) *InvokeEsAgentResponseBody {
	s.Data = v
	return s
}

func (s *InvokeEsAgentResponseBody) SetMessage(v string) *InvokeEsAgentResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeEsAgentResponseBody) SetRequestId(v string) *InvokeEsAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeEsAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
