// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYaoChiAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *GetYaoChiAgentResponseBody
	GetContent() *string
	SetFunctionCall(v []*GetYaoChiAgentResponseBodyFunctionCall) *GetYaoChiAgentResponseBody
	GetFunctionCall() []*GetYaoChiAgentResponseBodyFunctionCall
	SetParentId(v string) *GetYaoChiAgentResponseBody
	GetParentId() *string
	SetProduct(v string) *GetYaoChiAgentResponseBody
	GetProduct() *string
	SetQueryId(v string) *GetYaoChiAgentResponseBody
	GetQueryId() *string
	SetReasoningContent(v string) *GetYaoChiAgentResponseBody
	GetReasoningContent() *string
	SetRequestId(v string) *GetYaoChiAgentResponseBody
	GetRequestId() *string
	SetSessionId(v string) *GetYaoChiAgentResponseBody
	GetSessionId() *string
	SetSubAgentCall(v []*GetYaoChiAgentResponseBodySubAgentCall) *GetYaoChiAgentResponseBody
	GetSubAgentCall() []*GetYaoChiAgentResponseBodySubAgentCall
	SetUiFunctionCall(v []*GetYaoChiAgentResponseBodyUiFunctionCall) *GetYaoChiAgentResponseBody
	GetUiFunctionCall() []*GetYaoChiAgentResponseBodyUiFunctionCall
}

type GetYaoChiAgentResponseBody struct {
	// example:
	//
	// xxx
	Content      *string                                   `json:"Content,omitempty" xml:"Content,omitempty"`
	FunctionCall []*GetYaoChiAgentResponseBodyFunctionCall `json:"FunctionCall,omitempty" xml:"FunctionCall,omitempty" type:"Repeated"`
	ParentId     *string                                   `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// polardb
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-xxxxxxxxxxxx
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// example:
	//
	// xxx
	ReasoningContent *string `json:"ReasoningContent,omitempty" xml:"ReasoningContent,omitempty"`
	// example:
	//
	// 7172BECE-588A-5961-8126-C216E16B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-xxxxxxxxxxxx
	SessionId      *string                                     `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SubAgentCall   []*GetYaoChiAgentResponseBodySubAgentCall   `json:"SubAgentCall,omitempty" xml:"SubAgentCall,omitempty" type:"Repeated"`
	UiFunctionCall []*GetYaoChiAgentResponseBodyUiFunctionCall `json:"UiFunctionCall,omitempty" xml:"UiFunctionCall,omitempty" type:"Repeated"`
}

func (s GetYaoChiAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYaoChiAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetYaoChiAgentResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetYaoChiAgentResponseBody) GetFunctionCall() []*GetYaoChiAgentResponseBodyFunctionCall {
	return s.FunctionCall
}

func (s *GetYaoChiAgentResponseBody) GetParentId() *string {
	return s.ParentId
}

func (s *GetYaoChiAgentResponseBody) GetProduct() *string {
	return s.Product
}

func (s *GetYaoChiAgentResponseBody) GetQueryId() *string {
	return s.QueryId
}

func (s *GetYaoChiAgentResponseBody) GetReasoningContent() *string {
	return s.ReasoningContent
}

func (s *GetYaoChiAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYaoChiAgentResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *GetYaoChiAgentResponseBody) GetSubAgentCall() []*GetYaoChiAgentResponseBodySubAgentCall {
	return s.SubAgentCall
}

func (s *GetYaoChiAgentResponseBody) GetUiFunctionCall() []*GetYaoChiAgentResponseBodyUiFunctionCall {
	return s.UiFunctionCall
}

func (s *GetYaoChiAgentResponseBody) SetContent(v string) *GetYaoChiAgentResponseBody {
	s.Content = &v
	return s
}

func (s *GetYaoChiAgentResponseBody) SetFunctionCall(v []*GetYaoChiAgentResponseBodyFunctionCall) *GetYaoChiAgentResponseBody {
	s.FunctionCall = v
	return s
}

func (s *GetYaoChiAgentResponseBody) SetParentId(v string) *GetYaoChiAgentResponseBody {
	s.ParentId = &v
	return s
}

func (s *GetYaoChiAgentResponseBody) SetProduct(v string) *GetYaoChiAgentResponseBody {
	s.Product = &v
	return s
}

func (s *GetYaoChiAgentResponseBody) SetQueryId(v string) *GetYaoChiAgentResponseBody {
	s.QueryId = &v
	return s
}

func (s *GetYaoChiAgentResponseBody) SetReasoningContent(v string) *GetYaoChiAgentResponseBody {
	s.ReasoningContent = &v
	return s
}

func (s *GetYaoChiAgentResponseBody) SetRequestId(v string) *GetYaoChiAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYaoChiAgentResponseBody) SetSessionId(v string) *GetYaoChiAgentResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetYaoChiAgentResponseBody) SetSubAgentCall(v []*GetYaoChiAgentResponseBodySubAgentCall) *GetYaoChiAgentResponseBody {
	s.SubAgentCall = v
	return s
}

func (s *GetYaoChiAgentResponseBody) SetUiFunctionCall(v []*GetYaoChiAgentResponseBodyUiFunctionCall) *GetYaoChiAgentResponseBody {
	s.UiFunctionCall = v
	return s
}

func (s *GetYaoChiAgentResponseBody) Validate() error {
	if s.FunctionCall != nil {
		for _, item := range s.FunctionCall {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubAgentCall != nil {
		for _, item := range s.SubAgentCall {
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

type GetYaoChiAgentResponseBodyFunctionCall struct {
	// example:
	//
	// {"arg": "xxx"}
	Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	// example:
	//
	// 123447
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// sqlReview
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetYaoChiAgentResponseBodyFunctionCall) String() string {
	return dara.Prettify(s)
}

func (s GetYaoChiAgentResponseBodyFunctionCall) GoString() string {
	return s.String()
}

func (s *GetYaoChiAgentResponseBodyFunctionCall) GetArguments() *string {
	return s.Arguments
}

func (s *GetYaoChiAgentResponseBodyFunctionCall) GetId() *string {
	return s.Id
}

func (s *GetYaoChiAgentResponseBodyFunctionCall) GetName() *string {
	return s.Name
}

func (s *GetYaoChiAgentResponseBodyFunctionCall) GetStatus() *string {
	return s.Status
}

func (s *GetYaoChiAgentResponseBodyFunctionCall) SetArguments(v string) *GetYaoChiAgentResponseBodyFunctionCall {
	s.Arguments = &v
	return s
}

func (s *GetYaoChiAgentResponseBodyFunctionCall) SetId(v string) *GetYaoChiAgentResponseBodyFunctionCall {
	s.Id = &v
	return s
}

func (s *GetYaoChiAgentResponseBodyFunctionCall) SetName(v string) *GetYaoChiAgentResponseBodyFunctionCall {
	s.Name = &v
	return s
}

func (s *GetYaoChiAgentResponseBodyFunctionCall) SetStatus(v string) *GetYaoChiAgentResponseBodyFunctionCall {
	s.Status = &v
	return s
}

func (s *GetYaoChiAgentResponseBodyFunctionCall) Validate() error {
	return dara.Validate(s)
}

type GetYaoChiAgentResponseBodySubAgentCall struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubAgentId   *string `json:"SubAgentId,omitempty" xml:"SubAgentId,omitempty"`
	SubAgentName *string `json:"SubAgentName,omitempty" xml:"SubAgentName,omitempty"`
}

func (s GetYaoChiAgentResponseBodySubAgentCall) String() string {
	return dara.Prettify(s)
}

func (s GetYaoChiAgentResponseBodySubAgentCall) GoString() string {
	return s.String()
}

func (s *GetYaoChiAgentResponseBodySubAgentCall) GetStatus() *string {
	return s.Status
}

func (s *GetYaoChiAgentResponseBodySubAgentCall) GetSubAgentId() *string {
	return s.SubAgentId
}

func (s *GetYaoChiAgentResponseBodySubAgentCall) GetSubAgentName() *string {
	return s.SubAgentName
}

func (s *GetYaoChiAgentResponseBodySubAgentCall) SetStatus(v string) *GetYaoChiAgentResponseBodySubAgentCall {
	s.Status = &v
	return s
}

func (s *GetYaoChiAgentResponseBodySubAgentCall) SetSubAgentId(v string) *GetYaoChiAgentResponseBodySubAgentCall {
	s.SubAgentId = &v
	return s
}

func (s *GetYaoChiAgentResponseBodySubAgentCall) SetSubAgentName(v string) *GetYaoChiAgentResponseBodySubAgentCall {
	s.SubAgentName = &v
	return s
}

func (s *GetYaoChiAgentResponseBodySubAgentCall) Validate() error {
	return dara.Validate(s)
}

type GetYaoChiAgentResponseBodyUiFunctionCall struct {
	// example:
	//
	// {"arg": "xxx"}
	ArgsText *string `json:"ArgsText,omitempty" xml:"ArgsText,omitempty"`
	// example:
	//
	// card
	ToolName *string `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
}

func (s GetYaoChiAgentResponseBodyUiFunctionCall) String() string {
	return dara.Prettify(s)
}

func (s GetYaoChiAgentResponseBodyUiFunctionCall) GoString() string {
	return s.String()
}

func (s *GetYaoChiAgentResponseBodyUiFunctionCall) GetArgsText() *string {
	return s.ArgsText
}

func (s *GetYaoChiAgentResponseBodyUiFunctionCall) GetToolName() *string {
	return s.ToolName
}

func (s *GetYaoChiAgentResponseBodyUiFunctionCall) SetArgsText(v string) *GetYaoChiAgentResponseBodyUiFunctionCall {
	s.ArgsText = &v
	return s
}

func (s *GetYaoChiAgentResponseBodyUiFunctionCall) SetToolName(v string) *GetYaoChiAgentResponseBodyUiFunctionCall {
	s.ToolName = &v
	return s
}

func (s *GetYaoChiAgentResponseBodyUiFunctionCall) Validate() error {
	return dara.Validate(s)
}
