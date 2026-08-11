// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsByFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScriptsByFlowResponseBody
	GetCode() *string
	SetData(v []*ListScriptsByFlowResponseBodyData) *ListScriptsByFlowResponseBody
	GetData() []*ListScriptsByFlowResponseBodyData
	SetHttpStatusCode(v int32) *ListScriptsByFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListScriptsByFlowResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListScriptsByFlowResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListScriptsByFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListScriptsByFlowResponseBody
	GetSuccess() *bool
}

type ListScriptsByFlowResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data []*ListScriptsByFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance does not exist. Instance=ob-1234567890
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 019FDAC7-13C5-1B64-A853-999DF105B9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListScriptsByFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsByFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListScriptsByFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScriptsByFlowResponseBody) GetData() []*ListScriptsByFlowResponseBodyData {
	return s.Data
}

func (s *ListScriptsByFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListScriptsByFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScriptsByFlowResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListScriptsByFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScriptsByFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScriptsByFlowResponseBody) SetCode(v string) *ListScriptsByFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ListScriptsByFlowResponseBody) SetData(v []*ListScriptsByFlowResponseBodyData) *ListScriptsByFlowResponseBody {
	s.Data = v
	return s
}

func (s *ListScriptsByFlowResponseBody) SetHttpStatusCode(v int32) *ListScriptsByFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListScriptsByFlowResponseBody) SetMessage(v string) *ListScriptsByFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ListScriptsByFlowResponseBody) SetParams(v []*string) *ListScriptsByFlowResponseBody {
	s.Params = v
	return s
}

func (s *ListScriptsByFlowResponseBody) SetRequestId(v string) *ListScriptsByFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScriptsByFlowResponseBody) SetSuccess(v bool) *ListScriptsByFlowResponseBody {
	s.Success = &v
	return s
}

func (s *ListScriptsByFlowResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScriptsByFlowResponseBodyData struct {
	// The chatbot ID.
	//
	// example:
	//
	// chatbot-cn-MQuyjjb666
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// The name.
	//
	// example:
	//
	// Satisfaction Survey
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b54
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ListScriptsByFlowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsByFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListScriptsByFlowResponseBodyData) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *ListScriptsByFlowResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListScriptsByFlowResponseBodyData) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListScriptsByFlowResponseBodyData) SetChatbotId(v string) *ListScriptsByFlowResponseBodyData {
	s.ChatbotId = &v
	return s
}

func (s *ListScriptsByFlowResponseBodyData) SetName(v string) *ListScriptsByFlowResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListScriptsByFlowResponseBodyData) SetScriptId(v string) *ListScriptsByFlowResponseBodyData {
	s.ScriptId = &v
	return s
}

func (s *ListScriptsByFlowResponseBodyData) Validate() error {
	return dara.Validate(s)
}
