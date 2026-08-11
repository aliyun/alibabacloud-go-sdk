// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppTraceDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysis(v string) *GetAiAppTraceDetailResponseBody
	GetAnalysis() *string
	SetAppId(v string) *GetAiAppTraceDetailResponseBody
	GetAppId() *string
	SetAppName(v string) *GetAiAppTraceDetailResponseBody
	GetAppName() *string
	SetChannel(v string) *GetAiAppTraceDetailResponseBody
	GetChannel() *string
	SetLabels(v []*GetAiAppTraceDetailResponseBodyLabels) *GetAiAppTraceDetailResponseBody
	GetLabels() []*GetAiAppTraceDetailResponseBodyLabels
	SetRequestId(v string) *GetAiAppTraceDetailResponseBody
	GetRequestId() *string
	SetTraceId(v string) *GetAiAppTraceDetailResponseBody
	GetTraceId() *string
	SetWarningTime(v string) *GetAiAppTraceDetailResponseBody
	GetWarningTime() *string
}

type GetAiAppTraceDetailResponseBody struct {
	// The AI analysis result.
	//
	// example:
	//
	// xxxx
	Analysis *string `json:"Analysis,omitempty" xml:"Analysis,omitempty"`
	// The application ID.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// appxxx
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application channel.
	//
	// example:
	//
	// bailian
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The list of labels.
	Labels []*GetAiAppTraceDetailResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The backend-assigned ID that uniquely identifies a request. You can use this ID for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The trace ID used to correlate and track alert events.
	//
	// example:
	//
	// 0abb7ee117615311812886711e0a15
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// The alert time.
	//
	// example:
	//
	// 2026-01-01 16:08:38
	WarningTime *string `json:"WarningTime,omitempty" xml:"WarningTime,omitempty"`
}

func (s GetAiAppTraceDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppTraceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiAppTraceDetailResponseBody) GetAnalysis() *string {
	return s.Analysis
}

func (s *GetAiAppTraceDetailResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GetAiAppTraceDetailResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAiAppTraceDetailResponseBody) GetChannel() *string {
	return s.Channel
}

func (s *GetAiAppTraceDetailResponseBody) GetLabels() []*GetAiAppTraceDetailResponseBodyLabels {
	return s.Labels
}

func (s *GetAiAppTraceDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiAppTraceDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GetAiAppTraceDetailResponseBody) GetWarningTime() *string {
	return s.WarningTime
}

func (s *GetAiAppTraceDetailResponseBody) SetAnalysis(v string) *GetAiAppTraceDetailResponseBody {
	s.Analysis = &v
	return s
}

func (s *GetAiAppTraceDetailResponseBody) SetAppId(v string) *GetAiAppTraceDetailResponseBody {
	s.AppId = &v
	return s
}

func (s *GetAiAppTraceDetailResponseBody) SetAppName(v string) *GetAiAppTraceDetailResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAiAppTraceDetailResponseBody) SetChannel(v string) *GetAiAppTraceDetailResponseBody {
	s.Channel = &v
	return s
}

func (s *GetAiAppTraceDetailResponseBody) SetLabels(v []*GetAiAppTraceDetailResponseBodyLabels) *GetAiAppTraceDetailResponseBody {
	s.Labels = v
	return s
}

func (s *GetAiAppTraceDetailResponseBody) SetRequestId(v string) *GetAiAppTraceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiAppTraceDetailResponseBody) SetTraceId(v string) *GetAiAppTraceDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *GetAiAppTraceDetailResponseBody) SetWarningTime(v string) *GetAiAppTraceDetailResponseBody {
	s.WarningTime = &v
	return s
}

func (s *GetAiAppTraceDetailResponseBody) Validate() error {
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

type GetAiAppTraceDetailResponseBodyLabels struct {
	// The count.
	//
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
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
	// xxx
	LabelDesc *string `json:"LabelDesc,omitempty" xml:"LabelDesc,omitempty"`
	// The type.
	//
	// example:
	//
	// sensitiveData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAiAppTraceDetailResponseBodyLabels) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppTraceDetailResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetAiAppTraceDetailResponseBodyLabels) GetCount() *int64 {
	return s.Count
}

func (s *GetAiAppTraceDetailResponseBodyLabels) GetLabel() *string {
	return s.Label
}

func (s *GetAiAppTraceDetailResponseBodyLabels) GetLabelDesc() *string {
	return s.LabelDesc
}

func (s *GetAiAppTraceDetailResponseBodyLabels) GetType() *string {
	return s.Type
}

func (s *GetAiAppTraceDetailResponseBodyLabels) SetCount(v int64) *GetAiAppTraceDetailResponseBodyLabels {
	s.Count = &v
	return s
}

func (s *GetAiAppTraceDetailResponseBodyLabels) SetLabel(v string) *GetAiAppTraceDetailResponseBodyLabels {
	s.Label = &v
	return s
}

func (s *GetAiAppTraceDetailResponseBodyLabels) SetLabelDesc(v string) *GetAiAppTraceDetailResponseBodyLabels {
	s.LabelDesc = &v
	return s
}

func (s *GetAiAppTraceDetailResponseBodyLabels) SetType(v string) *GetAiAppTraceDetailResponseBodyLabels {
	s.Type = &v
	return s
}

func (s *GetAiAppTraceDetailResponseBodyLabels) Validate() error {
	return dara.Validate(s)
}
