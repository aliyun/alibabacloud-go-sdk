// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDialogueFlowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDialogueFlowsResponseBody
	GetCode() *string
	SetDialogueFlows(v []*ListDialogueFlowsResponseBodyDialogueFlows) *ListDialogueFlowsResponseBody
	GetDialogueFlows() []*ListDialogueFlowsResponseBodyDialogueFlows
	SetHttpStatusCode(v int32) *ListDialogueFlowsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDialogueFlowsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDialogueFlowsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDialogueFlowsResponseBody
	GetSuccess() *bool
}

type ListDialogueFlowsResponseBody struct {
	// Status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of conversation flow data
	DialogueFlows []*ListDialogueFlowsResponseBodyDialogueFlows `json:"DialogueFlows,omitempty" xml:"DialogueFlows,omitempty" type:"Repeated"`
	// HTTP return code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API prompt message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the invocation succeeded.
	//
	// - **true**: The invocation succeeded.
	//
	// - **false**: Failed to invoke.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDialogueFlowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDialogueFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDialogueFlowsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDialogueFlowsResponseBody) GetDialogueFlows() []*ListDialogueFlowsResponseBodyDialogueFlows {
	return s.DialogueFlows
}

func (s *ListDialogueFlowsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDialogueFlowsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDialogueFlowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDialogueFlowsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDialogueFlowsResponseBody) SetCode(v string) *ListDialogueFlowsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDialogueFlowsResponseBody) SetDialogueFlows(v []*ListDialogueFlowsResponseBodyDialogueFlows) *ListDialogueFlowsResponseBody {
	s.DialogueFlows = v
	return s
}

func (s *ListDialogueFlowsResponseBody) SetHttpStatusCode(v int32) *ListDialogueFlowsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDialogueFlowsResponseBody) SetMessage(v string) *ListDialogueFlowsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDialogueFlowsResponseBody) SetRequestId(v string) *ListDialogueFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDialogueFlowsResponseBody) SetSuccess(v bool) *ListDialogueFlowsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDialogueFlowsResponseBody) Validate() error {
	if s.DialogueFlows != nil {
		for _, item := range s.DialogueFlows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDialogueFlowsResponseBodyDialogueFlows struct {
	// Conversation flow data
	//
	// example:
	//
	// {"transitions":[{"source":"cc31e02b","sourceAnchor":0,"target":"345700de","targetAnchor":0,"id":"16d37c6e","index":1}],"nodes":[{"shape":"start-html","type":"start","size":"170*34","x":180,"y":134,"id":"cc31e02b","index":0,"nodeIndex":0,"content":{"branches":[{"branchName":"发起对话","branchId":"3c50a880-a7bc-11e9-80fc-5917e4f31864"}]},"coordinates":{"x":180,"y":134}},{"shape":"function-html","type":"transfer","size":"170*34","x":606,"y":134,"id":"345700de","nodeIndex":1,"name":"功能节点","script":"你是 @name ma","content":{"action":"Hangup","actionParams":"","label":["test1","test2","test3","test4","test5"]},"coordinates":{"x":606,"y":134},"index":2,"questions":["你是 @联系人姓名 ma"]}]}
	DialogueFlowDefinition *string `json:"DialogueFlowDefinition,omitempty" xml:"DialogueFlowDefinition,omitempty"`
	// Conversation flow ID
	//
	// example:
	//
	// dae01529-3c3e-458e-b07a-97643d09ebb9
	DialogueFlowId *string `json:"DialogueFlowId,omitempty" xml:"DialogueFlowId,omitempty"`
	// Conversation flow name
	//
	// example:
	//
	// 主流程
	DialogueFlowName *string `json:"DialogueFlowName,omitempty" xml:"DialogueFlowName,omitempty"`
	// Flow type.
	//
	// - SubFlow: child flow type
	//
	// - MainFlow: main flow type
	//
	// example:
	//
	// MainFlow
	DialogueFlowType *string `json:"DialogueFlowType,omitempty" xml:"DialogueFlowType,omitempty"`
	// Outbound scenario ID
	//
	// example:
	//
	// 2d5aa451-661f-4f08-b0c4-28eec78decc4
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Scenario version number
	//
	// example:
	//
	// 1569220563549
	ScriptVersion *string `json:"ScriptVersion,omitempty" xml:"ScriptVersion,omitempty"`
}

func (s ListDialogueFlowsResponseBodyDialogueFlows) String() string {
	return dara.Prettify(s)
}

func (s ListDialogueFlowsResponseBodyDialogueFlows) GoString() string {
	return s.String()
}

func (s *ListDialogueFlowsResponseBodyDialogueFlows) GetDialogueFlowDefinition() *string {
	return s.DialogueFlowDefinition
}

func (s *ListDialogueFlowsResponseBodyDialogueFlows) GetDialogueFlowId() *string {
	return s.DialogueFlowId
}

func (s *ListDialogueFlowsResponseBodyDialogueFlows) GetDialogueFlowName() *string {
	return s.DialogueFlowName
}

func (s *ListDialogueFlowsResponseBodyDialogueFlows) GetDialogueFlowType() *string {
	return s.DialogueFlowType
}

func (s *ListDialogueFlowsResponseBodyDialogueFlows) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListDialogueFlowsResponseBodyDialogueFlows) GetScriptVersion() *string {
	return s.ScriptVersion
}

func (s *ListDialogueFlowsResponseBodyDialogueFlows) SetDialogueFlowDefinition(v string) *ListDialogueFlowsResponseBodyDialogueFlows {
	s.DialogueFlowDefinition = &v
	return s
}

func (s *ListDialogueFlowsResponseBodyDialogueFlows) SetDialogueFlowId(v string) *ListDialogueFlowsResponseBodyDialogueFlows {
	s.DialogueFlowId = &v
	return s
}

func (s *ListDialogueFlowsResponseBodyDialogueFlows) SetDialogueFlowName(v string) *ListDialogueFlowsResponseBodyDialogueFlows {
	s.DialogueFlowName = &v
	return s
}

func (s *ListDialogueFlowsResponseBodyDialogueFlows) SetDialogueFlowType(v string) *ListDialogueFlowsResponseBodyDialogueFlows {
	s.DialogueFlowType = &v
	return s
}

func (s *ListDialogueFlowsResponseBodyDialogueFlows) SetScriptId(v string) *ListDialogueFlowsResponseBodyDialogueFlows {
	s.ScriptId = &v
	return s
}

func (s *ListDialogueFlowsResponseBodyDialogueFlows) SetScriptVersion(v string) *ListDialogueFlowsResponseBodyDialogueFlows {
	s.ScriptVersion = &v
	return s
}

func (s *ListDialogueFlowsResponseBodyDialogueFlows) Validate() error {
	return dara.Validate(s)
}
