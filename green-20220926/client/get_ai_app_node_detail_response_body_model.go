// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppNodeDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAiAppNodeDetailResponseBody
	GetAppId() *string
	SetChannel(v string) *GetAiAppNodeDetailResponseBody
	GetChannel() *string
	SetEventData(v []*GetAiAppNodeDetailResponseBodyEventData) *GetAiAppNodeDetailResponseBody
	GetEventData() []*GetAiAppNodeDetailResponseBodyEventData
	SetNodeId(v string) *GetAiAppNodeDetailResponseBody
	GetNodeId() *string
	SetNodeName(v string) *GetAiAppNodeDetailResponseBody
	GetNodeName() *string
	SetNodeType(v string) *GetAiAppNodeDetailResponseBody
	GetNodeType() *string
	SetRequestId(v string) *GetAiAppNodeDetailResponseBody
	GetRequestId() *string
	SetRiskLevel(v string) *GetAiAppNodeDetailResponseBody
	GetRiskLevel() *string
}

type GetAiAppNodeDetailResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The channel information.
	//
	// example:
	//
	// bailian
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The list of event data.
	EventData []*GetAiAppNodeDetailResponseBodyEventData `json:"EventData,omitempty" xml:"EventData,omitempty" type:"Repeated"`
	// The node ID.
	//
	// example:
	//
	// node-xxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node name.
	//
	// example:
	//
	// namexxx
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The node type.
	//
	// example:
	//
	// TOOL
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The risk level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s GetAiAppNodeDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppNodeDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiAppNodeDetailResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GetAiAppNodeDetailResponseBody) GetChannel() *string {
	return s.Channel
}

func (s *GetAiAppNodeDetailResponseBody) GetEventData() []*GetAiAppNodeDetailResponseBodyEventData {
	return s.EventData
}

func (s *GetAiAppNodeDetailResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *GetAiAppNodeDetailResponseBody) GetNodeName() *string {
	return s.NodeName
}

func (s *GetAiAppNodeDetailResponseBody) GetNodeType() *string {
	return s.NodeType
}

func (s *GetAiAppNodeDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiAppNodeDetailResponseBody) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetAiAppNodeDetailResponseBody) SetAppId(v string) *GetAiAppNodeDetailResponseBody {
	s.AppId = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBody) SetChannel(v string) *GetAiAppNodeDetailResponseBody {
	s.Channel = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBody) SetEventData(v []*GetAiAppNodeDetailResponseBodyEventData) *GetAiAppNodeDetailResponseBody {
	s.EventData = v
	return s
}

func (s *GetAiAppNodeDetailResponseBody) SetNodeId(v string) *GetAiAppNodeDetailResponseBody {
	s.NodeId = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBody) SetNodeName(v string) *GetAiAppNodeDetailResponseBody {
	s.NodeName = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBody) SetNodeType(v string) *GetAiAppNodeDetailResponseBody {
	s.NodeType = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBody) SetRequestId(v string) *GetAiAppNodeDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBody) SetRiskLevel(v string) *GetAiAppNodeDetailResponseBody {
	s.RiskLevel = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBody) Validate() error {
	if s.EventData != nil {
		for _, item := range s.EventData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAiAppNodeDetailResponseBodyEventData struct {
	// The channel.
	//
	// example:
	//
	// bailian
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The list of labels.
	Labels []*GetAiAppNodeDetailResponseBodyEventDataLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name.
	//
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The risk level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The time when the event occurred.
	//
	// example:
	//
	// 2026-01-01 16:08:38
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The trace ID, which is used to query the exact call information.
	//
	// example:
	//
	// xxxxx
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// The event type.
	//
	// example:
	//
	// hit_sensitive_data
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAiAppNodeDetailResponseBodyEventData) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppNodeDetailResponseBodyEventData) GoString() string {
	return s.String()
}

func (s *GetAiAppNodeDetailResponseBodyEventData) GetChannel() *string {
	return s.Channel
}

func (s *GetAiAppNodeDetailResponseBodyEventData) GetLabels() []*GetAiAppNodeDetailResponseBodyEventDataLabels {
	return s.Labels
}

func (s *GetAiAppNodeDetailResponseBodyEventData) GetName() *string {
	return s.Name
}

func (s *GetAiAppNodeDetailResponseBodyEventData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetAiAppNodeDetailResponseBodyEventData) GetTime() *string {
	return s.Time
}

func (s *GetAiAppNodeDetailResponseBodyEventData) GetTraceId() *string {
	return s.TraceId
}

func (s *GetAiAppNodeDetailResponseBodyEventData) GetType() *string {
	return s.Type
}

func (s *GetAiAppNodeDetailResponseBodyEventData) SetChannel(v string) *GetAiAppNodeDetailResponseBodyEventData {
	s.Channel = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBodyEventData) SetLabels(v []*GetAiAppNodeDetailResponseBodyEventDataLabels) *GetAiAppNodeDetailResponseBodyEventData {
	s.Labels = v
	return s
}

func (s *GetAiAppNodeDetailResponseBodyEventData) SetName(v string) *GetAiAppNodeDetailResponseBodyEventData {
	s.Name = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBodyEventData) SetRiskLevel(v string) *GetAiAppNodeDetailResponseBodyEventData {
	s.RiskLevel = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBodyEventData) SetTime(v string) *GetAiAppNodeDetailResponseBodyEventData {
	s.Time = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBodyEventData) SetTraceId(v string) *GetAiAppNodeDetailResponseBodyEventData {
	s.TraceId = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBodyEventData) SetType(v string) *GetAiAppNodeDetailResponseBodyEventData {
	s.Type = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBodyEventData) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAiAppNodeDetailResponseBodyEventDataLabels struct {
	// The label name.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The label description.
	//
	// example:
	//
	// porn desc
	LabelDesc *string `json:"LabelDesc,omitempty" xml:"LabelDesc,omitempty"`
	// The risk level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s GetAiAppNodeDetailResponseBodyEventDataLabels) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppNodeDetailResponseBodyEventDataLabels) GoString() string {
	return s.String()
}

func (s *GetAiAppNodeDetailResponseBodyEventDataLabels) GetLabel() *string {
	return s.Label
}

func (s *GetAiAppNodeDetailResponseBodyEventDataLabels) GetLabelDesc() *string {
	return s.LabelDesc
}

func (s *GetAiAppNodeDetailResponseBodyEventDataLabels) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetAiAppNodeDetailResponseBodyEventDataLabels) SetLabel(v string) *GetAiAppNodeDetailResponseBodyEventDataLabels {
	s.Label = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBodyEventDataLabels) SetLabelDesc(v string) *GetAiAppNodeDetailResponseBodyEventDataLabels {
	s.LabelDesc = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBodyEventDataLabels) SetRiskLevel(v string) *GetAiAppNodeDetailResponseBodyEventDataLabels {
	s.RiskLevel = &v
	return s
}

func (s *GetAiAppNodeDetailResponseBodyEventDataLabels) Validate() error {
	return dara.Validate(s)
}
