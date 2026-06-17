// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolarAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *GetPolarAgentResponseBody
	GetContent() *string
	SetFunctionCall(v []*GetPolarAgentResponseBodyFunctionCall) *GetPolarAgentResponseBody
	GetFunctionCall() []*GetPolarAgentResponseBodyFunctionCall
	SetProduct(v string) *GetPolarAgentResponseBody
	GetProduct() *string
	SetQueryId(v string) *GetPolarAgentResponseBody
	GetQueryId() *string
	SetReasoningContent(v string) *GetPolarAgentResponseBody
	GetReasoningContent() *string
	SetRequestId(v string) *GetPolarAgentResponseBody
	GetRequestId() *string
	SetSessionId(v string) *GetPolarAgentResponseBody
	GetSessionId() *string
	SetUiFunctionCall(v []*GetPolarAgentResponseBodyUiFunctionCall) *GetPolarAgentResponseBody
	GetUiFunctionCall() []*GetPolarAgentResponseBodyUiFunctionCall
}

type GetPolarAgentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// QZ-13661677-80
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Indicates whether FunctionCall is supported.
	FunctionCall []*GetPolarAgentResponseBodyFunctionCall `json:"FunctionCall,omitempty" xml:"FunctionCall,omitempty" type:"Repeated"`
	// The cloud product type.
	//
	// example:
	//
	// polardb
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The ID of the user query pipeline task.
	//
	// example:
	//
	// 2548026401648157601743560466154
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The model\\"s reasoning content.
	//
	// example:
	//
	// xxxx
	ReasoningContent *string `json:"ReasoningContent,omitempty" xml:"ReasoningContent,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34458CD3-33E0-4624-BFEF-840C15******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The session ID for code execution.
	//
	// example:
	//
	// 40315d708f0806903b08813bf4c9db2e
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Indicates whether UiFunctionCall is supported.
	UiFunctionCall []*GetPolarAgentResponseBodyUiFunctionCall `json:"UiFunctionCall,omitempty" xml:"UiFunctionCall,omitempty" type:"Repeated"`
}

func (s GetPolarAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolarAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolarAgentResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetPolarAgentResponseBody) GetFunctionCall() []*GetPolarAgentResponseBodyFunctionCall {
	return s.FunctionCall
}

func (s *GetPolarAgentResponseBody) GetProduct() *string {
	return s.Product
}

func (s *GetPolarAgentResponseBody) GetQueryId() *string {
	return s.QueryId
}

func (s *GetPolarAgentResponseBody) GetReasoningContent() *string {
	return s.ReasoningContent
}

func (s *GetPolarAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolarAgentResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *GetPolarAgentResponseBody) GetUiFunctionCall() []*GetPolarAgentResponseBodyUiFunctionCall {
	return s.UiFunctionCall
}

func (s *GetPolarAgentResponseBody) SetContent(v string) *GetPolarAgentResponseBody {
	s.Content = &v
	return s
}

func (s *GetPolarAgentResponseBody) SetFunctionCall(v []*GetPolarAgentResponseBodyFunctionCall) *GetPolarAgentResponseBody {
	s.FunctionCall = v
	return s
}

func (s *GetPolarAgentResponseBody) SetProduct(v string) *GetPolarAgentResponseBody {
	s.Product = &v
	return s
}

func (s *GetPolarAgentResponseBody) SetQueryId(v string) *GetPolarAgentResponseBody {
	s.QueryId = &v
	return s
}

func (s *GetPolarAgentResponseBody) SetReasoningContent(v string) *GetPolarAgentResponseBody {
	s.ReasoningContent = &v
	return s
}

func (s *GetPolarAgentResponseBody) SetRequestId(v string) *GetPolarAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolarAgentResponseBody) SetSessionId(v string) *GetPolarAgentResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetPolarAgentResponseBody) SetUiFunctionCall(v []*GetPolarAgentResponseBodyUiFunctionCall) *GetPolarAgentResponseBody {
	s.UiFunctionCall = v
	return s
}

func (s *GetPolarAgentResponseBody) Validate() error {
	if s.FunctionCall != nil {
		for _, item := range s.FunctionCall {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UiFunctionCall != nil {
		for _, item := range s.UiFunctionCall {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPolarAgentResponseBodyFunctionCall struct {
	// The operation-related output result.
	//
	// example:
	//
	// ---narguments:n  parameters: []n
	Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// 393
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule04
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status.
	//
	// example:
	//
	// xxx
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPolarAgentResponseBodyFunctionCall) String() string {
	return dara.Prettify(s)
}

func (s GetPolarAgentResponseBodyFunctionCall) GoString() string {
	return s.String()
}

func (s *GetPolarAgentResponseBodyFunctionCall) GetArguments() *string {
	return s.Arguments
}

func (s *GetPolarAgentResponseBodyFunctionCall) GetId() *string {
	return s.Id
}

func (s *GetPolarAgentResponseBodyFunctionCall) GetName() *string {
	return s.Name
}

func (s *GetPolarAgentResponseBodyFunctionCall) GetStatus() *string {
	return s.Status
}

func (s *GetPolarAgentResponseBodyFunctionCall) SetArguments(v string) *GetPolarAgentResponseBodyFunctionCall {
	s.Arguments = &v
	return s
}

func (s *GetPolarAgentResponseBodyFunctionCall) SetId(v string) *GetPolarAgentResponseBodyFunctionCall {
	s.Id = &v
	return s
}

func (s *GetPolarAgentResponseBodyFunctionCall) SetName(v string) *GetPolarAgentResponseBodyFunctionCall {
	s.Name = &v
	return s
}

func (s *GetPolarAgentResponseBodyFunctionCall) SetStatus(v string) *GetPolarAgentResponseBodyFunctionCall {
	s.Status = &v
	return s
}

func (s *GetPolarAgentResponseBodyFunctionCall) Validate() error {
	return dara.Validate(s)
}

type GetPolarAgentResponseBodyUiFunctionCall struct {
	// xxx
	//
	// example:
	//
	// xxx
	ArgsText *string `json:"ArgsText,omitempty" xml:"ArgsText,omitempty"`
	// The tool command that is invoked.
	//
	// example:
	//
	// xxx
	ToolName *string `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
}

func (s GetPolarAgentResponseBodyUiFunctionCall) String() string {
	return dara.Prettify(s)
}

func (s GetPolarAgentResponseBodyUiFunctionCall) GoString() string {
	return s.String()
}

func (s *GetPolarAgentResponseBodyUiFunctionCall) GetArgsText() *string {
	return s.ArgsText
}

func (s *GetPolarAgentResponseBodyUiFunctionCall) GetToolName() *string {
	return s.ToolName
}

func (s *GetPolarAgentResponseBodyUiFunctionCall) SetArgsText(v string) *GetPolarAgentResponseBodyUiFunctionCall {
	s.ArgsText = &v
	return s
}

func (s *GetPolarAgentResponseBodyUiFunctionCall) SetToolName(v string) *GetPolarAgentResponseBodyUiFunctionCall {
	s.ToolName = &v
	return s
}

func (s *GetPolarAgentResponseBodyUiFunctionCall) Validate() error {
	return dara.Validate(s)
}
