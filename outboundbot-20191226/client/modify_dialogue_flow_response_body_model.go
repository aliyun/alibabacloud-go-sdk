// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDialogueFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyDialogueFlowResponseBody
	GetCode() *string
	SetDialogueFlowDefinition(v string) *ModifyDialogueFlowResponseBody
	GetDialogueFlowDefinition() *string
	SetDialogueFlowId(v string) *ModifyDialogueFlowResponseBody
	GetDialogueFlowId() *string
	SetHttpStatusCode(v int32) *ModifyDialogueFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyDialogueFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyDialogueFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDialogueFlowResponseBody
	GetSuccess() *bool
}

type ModifyDialogueFlowResponseBody struct {
	// Status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Conversation flow Data
	//
	// example:
	//
	// {"transitions":[{"id":"a91c4023","index":1,"source":"cc31e02b","sourceAnchor":0,"target":"946045be","targetAnchor":0}],"nodes":[{"collectedNumberEnabled":false,"content":{"branches":[{"branchId":"f5450420-09ab-11ea-b107-e9059c6a79d8","branchName":"发起对话"}]},"coordinates":{"x":180,"y":134},"id":"cc31e02b","index":0,"interruptible":false,"nodeIndex":0,"shape":"start-html","size":"170*34","type":"start","x":180,"y":134},{"collectedNumberEnabled":false,"content":{"actionParams":"","action":"Hangup"},"coordinates":{"x":487.65625,"y":155},"id":"946045be","index":2,"interruptible":false,"labels":[],"name":"功能节点","nodeIndex":1,"questions":["好的同学，您的情况已了解了，稍后我们会安排资深顾问老师给您做更详细的留学评估以及升学指导，请留意电话接听"],"script":"好的同学，您的情况已了解了，稍后我们会安排资深顾问老师给您做更详细的留学评估以及升学指导，请留意电话接听","shape":"function-html","size":"170*34","type":"transfer","x":500,"y":182}]}
	DialogueFlowDefinition *string `json:"DialogueFlowDefinition,omitempty" xml:"DialogueFlowDefinition,omitempty"`
	// Conversation flow ID
	//
	// example:
	//
	// 390515b5-6115-4ccf-83e2-52d5bfaf2ddf
	DialogueFlowId *string `json:"DialogueFlowId,omitempty" xml:"DialogueFlowId,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Succcess
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request Succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDialogueFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDialogueFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDialogueFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyDialogueFlowResponseBody) GetDialogueFlowDefinition() *string {
	return s.DialogueFlowDefinition
}

func (s *ModifyDialogueFlowResponseBody) GetDialogueFlowId() *string {
	return s.DialogueFlowId
}

func (s *ModifyDialogueFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyDialogueFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyDialogueFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDialogueFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDialogueFlowResponseBody) SetCode(v string) *ModifyDialogueFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) SetDialogueFlowDefinition(v string) *ModifyDialogueFlowResponseBody {
	s.DialogueFlowDefinition = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) SetDialogueFlowId(v string) *ModifyDialogueFlowResponseBody {
	s.DialogueFlowId = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) SetHttpStatusCode(v int32) *ModifyDialogueFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) SetMessage(v string) *ModifyDialogueFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) SetRequestId(v string) *ModifyDialogueFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) SetSuccess(v bool) *ModifyDialogueFlowResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
