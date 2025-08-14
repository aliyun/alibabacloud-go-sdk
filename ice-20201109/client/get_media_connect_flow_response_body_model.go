// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *GetMediaConnectFlowResponseBodyContent) *GetMediaConnectFlowResponseBody
	GetContent() *GetMediaConnectFlowResponseBodyContent
	SetDescription(v string) *GetMediaConnectFlowResponseBody
	GetDescription() *string
	SetRequestId(v string) *GetMediaConnectFlowResponseBody
	GetRequestId() *string
	SetRetcode(v int32) *GetMediaConnectFlowResponseBody
	GetRetcode() *int32
}

type GetMediaConnectFlowResponseBody struct {
	// The response body.
	Content *GetMediaConnectFlowResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The call description.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FB503AEF-118E-1516-89E2-7B227EA1AC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	Retcode *int32 `json:"Retcode,omitempty" xml:"Retcode,omitempty"`
}

func (s GetMediaConnectFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowResponseBody) GetContent() *GetMediaConnectFlowResponseBodyContent {
	return s.Content
}

func (s *GetMediaConnectFlowResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetMediaConnectFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaConnectFlowResponseBody) GetRetcode() *int32 {
	return s.Retcode
}

func (s *GetMediaConnectFlowResponseBody) SetContent(v *GetMediaConnectFlowResponseBodyContent) *GetMediaConnectFlowResponseBody {
	s.Content = v
	return s
}

func (s *GetMediaConnectFlowResponseBody) SetDescription(v string) *GetMediaConnectFlowResponseBody {
	s.Description = &v
	return s
}

func (s *GetMediaConnectFlowResponseBody) SetRequestId(v string) *GetMediaConnectFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaConnectFlowResponseBody) SetRetcode(v int32) *GetMediaConnectFlowResponseBody {
	s.Retcode = &v
	return s
}

func (s *GetMediaConnectFlowResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMediaConnectFlowResponseBodyContent struct {
	// The time when the flow was created.
	//
	// example:
	//
	// 2024-07-18T01:29:24Z
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FlowFailover *string `json:"FlowFailover,omitempty" xml:"FlowFailover,omitempty"`
	// The flow ID.
	//
	// example:
	//
	// 34900dc6-90ec-4968-af3c-fcd87f231a5f
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The flow name.
	//
	// example:
	//
	// AliTestFlow
	FlowName   *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	FlowRegion *string `json:"FlowRegion,omitempty" xml:"FlowRegion,omitempty"`
	// The state of the flow.
	//
	// example:
	//
	// online
	FlowStatus *string `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty"`
	// The time when the flow is started.
	//
	// example:
	//
	// 2024-07-18T01:39:24Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetMediaConnectFlowResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowResponseBodyContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMediaConnectFlowResponseBodyContent) GetFlowFailover() *string {
	return s.FlowFailover
}

func (s *GetMediaConnectFlowResponseBodyContent) GetFlowId() *string {
	return s.FlowId
}

func (s *GetMediaConnectFlowResponseBodyContent) GetFlowName() *string {
	return s.FlowName
}

func (s *GetMediaConnectFlowResponseBodyContent) GetFlowRegion() *string {
	return s.FlowRegion
}

func (s *GetMediaConnectFlowResponseBodyContent) GetFlowStatus() *string {
	return s.FlowStatus
}

func (s *GetMediaConnectFlowResponseBodyContent) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMediaConnectFlowResponseBodyContent) SetCreateTime(v string) *GetMediaConnectFlowResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetMediaConnectFlowResponseBodyContent) SetFlowFailover(v string) *GetMediaConnectFlowResponseBodyContent {
	s.FlowFailover = &v
	return s
}

func (s *GetMediaConnectFlowResponseBodyContent) SetFlowId(v string) *GetMediaConnectFlowResponseBodyContent {
	s.FlowId = &v
	return s
}

func (s *GetMediaConnectFlowResponseBodyContent) SetFlowName(v string) *GetMediaConnectFlowResponseBodyContent {
	s.FlowName = &v
	return s
}

func (s *GetMediaConnectFlowResponseBodyContent) SetFlowRegion(v string) *GetMediaConnectFlowResponseBodyContent {
	s.FlowRegion = &v
	return s
}

func (s *GetMediaConnectFlowResponseBodyContent) SetFlowStatus(v string) *GetMediaConnectFlowResponseBodyContent {
	s.FlowStatus = &v
	return s
}

func (s *GetMediaConnectFlowResponseBodyContent) SetStartTime(v string) *GetMediaConnectFlowResponseBodyContent {
	s.StartTime = &v
	return s
}

func (s *GetMediaConnectFlowResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
