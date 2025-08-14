// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaConnectFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *CreateMediaConnectFlowResponseBodyContent) *CreateMediaConnectFlowResponseBody
	GetContent() *CreateMediaConnectFlowResponseBodyContent
	SetDescription(v string) *CreateMediaConnectFlowResponseBody
	GetDescription() *string
	SetRequestId(v string) *CreateMediaConnectFlowResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *CreateMediaConnectFlowResponseBody
	GetRetCode() *int32
}

type CreateMediaConnectFlowResponseBody struct {
	// The response body.
	Content *CreateMediaConnectFlowResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 86D92F9D-65E8-58A2-85D1-9DEEECC172E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s CreateMediaConnectFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConnectFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMediaConnectFlowResponseBody) GetContent() *CreateMediaConnectFlowResponseBodyContent {
	return s.Content
}

func (s *CreateMediaConnectFlowResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateMediaConnectFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMediaConnectFlowResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *CreateMediaConnectFlowResponseBody) SetContent(v *CreateMediaConnectFlowResponseBodyContent) *CreateMediaConnectFlowResponseBody {
	s.Content = v
	return s
}

func (s *CreateMediaConnectFlowResponseBody) SetDescription(v string) *CreateMediaConnectFlowResponseBody {
	s.Description = &v
	return s
}

func (s *CreateMediaConnectFlowResponseBody) SetRequestId(v string) *CreateMediaConnectFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMediaConnectFlowResponseBody) SetRetCode(v int32) *CreateMediaConnectFlowResponseBody {
	s.RetCode = &v
	return s
}

func (s *CreateMediaConnectFlowResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMediaConnectFlowResponseBodyContent struct {
	// The flow ID.
	//
	// example:
	//
	// 34900dc6-90ec-4968-af3c-fcd87f231a5f
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s CreateMediaConnectFlowResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConnectFlowResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateMediaConnectFlowResponseBodyContent) GetFlowId() *string {
	return s.FlowId
}

func (s *CreateMediaConnectFlowResponseBodyContent) SetFlowId(v string) *CreateMediaConnectFlowResponseBodyContent {
	s.FlowId = &v
	return s
}

func (s *CreateMediaConnectFlowResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
