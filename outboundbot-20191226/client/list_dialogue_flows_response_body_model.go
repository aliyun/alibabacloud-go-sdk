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
	// example:
	//
	// OK
	Code          *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	DialogueFlows []*ListDialogueFlowsResponseBodyDialogueFlows `json:"DialogueFlows,omitempty" xml:"DialogueFlows,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DialogueFlowDefinition *string `json:"DialogueFlowDefinition,omitempty" xml:"DialogueFlowDefinition,omitempty"`
	// example:
	//
	// dae01529-3c3e-458e-b07a-97643d09ebb9
	DialogueFlowId   *string `json:"DialogueFlowId,omitempty" xml:"DialogueFlowId,omitempty"`
	DialogueFlowName *string `json:"DialogueFlowName,omitempty" xml:"DialogueFlowName,omitempty"`
	// example:
	//
	// MainFlow
	DialogueFlowType *string `json:"DialogueFlowType,omitempty" xml:"DialogueFlowType,omitempty"`
	// example:
	//
	// 2d5aa451-661f-4f08-b0c4-28eec78decc4
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
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
