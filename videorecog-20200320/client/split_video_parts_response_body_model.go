// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSplitVideoPartsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SplitVideoPartsResponseBodyData) *SplitVideoPartsResponseBody
	GetData() *SplitVideoPartsResponseBodyData
	SetMessage(v string) *SplitVideoPartsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SplitVideoPartsResponseBody
	GetRequestId() *string
}

type SplitVideoPartsResponseBody struct {
	Data    *SplitVideoPartsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A00A3C17-61D5-1489-860D-B709F83A7C40
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SplitVideoPartsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SplitVideoPartsResponseBody) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsResponseBody) GetData() *SplitVideoPartsResponseBodyData {
	return s.Data
}

func (s *SplitVideoPartsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SplitVideoPartsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SplitVideoPartsResponseBody) SetData(v *SplitVideoPartsResponseBodyData) *SplitVideoPartsResponseBody {
	s.Data = v
	return s
}

func (s *SplitVideoPartsResponseBody) SetMessage(v string) *SplitVideoPartsResponseBody {
	s.Message = &v
	return s
}

func (s *SplitVideoPartsResponseBody) SetRequestId(v string) *SplitVideoPartsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SplitVideoPartsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SplitVideoPartsResponseBodyData struct {
	Elements              []*SplitVideoPartsResponseBodyDataElements              `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	SplitVideoPartResults []*SplitVideoPartsResponseBodyDataSplitVideoPartResults `json:"SplitVideoPartResults,omitempty" xml:"SplitVideoPartResults,omitempty" type:"Repeated"`
}

func (s SplitVideoPartsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SplitVideoPartsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsResponseBodyData) GetElements() []*SplitVideoPartsResponseBodyDataElements {
	return s.Elements
}

func (s *SplitVideoPartsResponseBodyData) GetSplitVideoPartResults() []*SplitVideoPartsResponseBodyDataSplitVideoPartResults {
	return s.SplitVideoPartResults
}

func (s *SplitVideoPartsResponseBodyData) SetElements(v []*SplitVideoPartsResponseBodyDataElements) *SplitVideoPartsResponseBodyData {
	s.Elements = v
	return s
}

func (s *SplitVideoPartsResponseBodyData) SetSplitVideoPartResults(v []*SplitVideoPartsResponseBodyDataSplitVideoPartResults) *SplitVideoPartsResponseBodyData {
	s.SplitVideoPartResults = v
	return s
}

func (s *SplitVideoPartsResponseBodyData) Validate() error {
	if s.Elements != nil {
		for _, item := range s.Elements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SplitVideoPartResults != nil {
		for _, item := range s.SplitVideoPartResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SplitVideoPartsResponseBodyDataElements struct {
	// example:
	//
	// 10.06
	BeginTime *float32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 17.3
	EndTime *float32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
}

func (s SplitVideoPartsResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s SplitVideoPartsResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsResponseBodyDataElements) GetBeginTime() *float32 {
	return s.BeginTime
}

func (s *SplitVideoPartsResponseBodyDataElements) GetEndTime() *float32 {
	return s.EndTime
}

func (s *SplitVideoPartsResponseBodyDataElements) GetIndex() *int64 {
	return s.Index
}

func (s *SplitVideoPartsResponseBodyDataElements) SetBeginTime(v float32) *SplitVideoPartsResponseBodyDataElements {
	s.BeginTime = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataElements) SetEndTime(v float32) *SplitVideoPartsResponseBodyDataElements {
	s.EndTime = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataElements) SetIndex(v int64) *SplitVideoPartsResponseBodyDataElements {
	s.Index = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}

type SplitVideoPartsResponseBodyDataSplitVideoPartResults struct {
	BeginTime *float32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	By        *string  `json:"By,omitempty" xml:"By,omitempty"`
	EndTime   *float32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Theme     *string  `json:"Theme,omitempty" xml:"Theme,omitempty"`
	Type      *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SplitVideoPartsResponseBodyDataSplitVideoPartResults) String() string {
	return dara.Prettify(s)
}

func (s SplitVideoPartsResponseBodyDataSplitVideoPartResults) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) GetBeginTime() *float32 {
	return s.BeginTime
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) GetBy() *string {
	return s.By
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) GetEndTime() *float32 {
	return s.EndTime
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) GetTheme() *string {
	return s.Theme
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) GetType() *string {
	return s.Type
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) SetBeginTime(v float32) *SplitVideoPartsResponseBodyDataSplitVideoPartResults {
	s.BeginTime = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) SetBy(v string) *SplitVideoPartsResponseBodyDataSplitVideoPartResults {
	s.By = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) SetEndTime(v float32) *SplitVideoPartsResponseBodyDataSplitVideoPartResults {
	s.EndTime = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) SetTheme(v string) *SplitVideoPartsResponseBodyDataSplitVideoPartResults {
	s.Theme = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) SetType(v string) *SplitVideoPartsResponseBodyDataSplitVideoPartResults {
	s.Type = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) Validate() error {
	return dara.Validate(s)
}
