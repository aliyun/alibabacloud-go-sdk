// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTongyiChatDebugInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *TongyiChatDebugInfoResponseBody
	GetMessageId() *string
	SetPipeline(v []*TongyiChatDebugInfoResponseBodyPipeline) *TongyiChatDebugInfoResponseBody
	GetPipeline() []*TongyiChatDebugInfoResponseBodyPipeline
	SetRequestId(v string) *TongyiChatDebugInfoResponseBody
	GetRequestId() *string
}

type TongyiChatDebugInfoResponseBody struct {
	// example:
	//
	// 2828708A-2C7A-1BAE-B810-87DB9DA9C661
	MessageId *string                                    `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Pipeline  []*TongyiChatDebugInfoResponseBodyPipeline `json:"Pipeline,omitempty" xml:"Pipeline,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// E3E5C779-A630-45AC-B0F2-A4506A4212F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TongyiChatDebugInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TongyiChatDebugInfoResponseBody) GoString() string {
	return s.String()
}

func (s *TongyiChatDebugInfoResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *TongyiChatDebugInfoResponseBody) GetPipeline() []*TongyiChatDebugInfoResponseBodyPipeline {
	return s.Pipeline
}

func (s *TongyiChatDebugInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TongyiChatDebugInfoResponseBody) SetMessageId(v string) *TongyiChatDebugInfoResponseBody {
	s.MessageId = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBody) SetPipeline(v []*TongyiChatDebugInfoResponseBodyPipeline) *TongyiChatDebugInfoResponseBody {
	s.Pipeline = v
	return s
}

func (s *TongyiChatDebugInfoResponseBody) SetRequestId(v string) *TongyiChatDebugInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type TongyiChatDebugInfoResponseBodyPipeline struct {
	Input interface{} `json:"Input,omitempty" xml:"Input,omitempty"`
	Name  *string     `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// system_strategy
	NodeType *string     `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Output   interface{} `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s TongyiChatDebugInfoResponseBodyPipeline) String() string {
	return dara.Prettify(s)
}

func (s TongyiChatDebugInfoResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) GetInput() interface{} {
	return s.Input
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) GetName() *string {
	return s.Name
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) GetNodeType() *string {
	return s.NodeType
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) GetOutput() interface{} {
	return s.Output
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) SetInput(v interface{}) *TongyiChatDebugInfoResponseBodyPipeline {
	s.Input = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) SetName(v string) *TongyiChatDebugInfoResponseBodyPipeline {
	s.Name = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) SetNodeType(v string) *TongyiChatDebugInfoResponseBodyPipeline {
	s.NodeType = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) SetOutput(v interface{}) *TongyiChatDebugInfoResponseBodyPipeline {
	s.Output = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) Validate() error {
	return dara.Validate(s)
}
