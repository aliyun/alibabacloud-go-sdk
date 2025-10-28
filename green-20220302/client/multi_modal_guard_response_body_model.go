// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *MultiModalGuardResponseBody
	GetCode() *int32
	SetData(v *MultiModalGuardResponseBodyData) *MultiModalGuardResponseBody
	GetData() *MultiModalGuardResponseBodyData
	SetMessage(v string) *MultiModalGuardResponseBody
	GetMessage() *string
	SetRequestId(v string) *MultiModalGuardResponseBody
	GetRequestId() *string
}

type MultiModalGuardResponseBody struct {
	// example:
	//
	// 200
	Code *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *MultiModalGuardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MultiModalGuardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardResponseBody) GoString() string {
	return s.String()
}

func (s *MultiModalGuardResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *MultiModalGuardResponseBody) GetData() *MultiModalGuardResponseBodyData {
	return s.Data
}

func (s *MultiModalGuardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MultiModalGuardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MultiModalGuardResponseBody) SetCode(v int32) *MultiModalGuardResponseBody {
	s.Code = &v
	return s
}

func (s *MultiModalGuardResponseBody) SetData(v *MultiModalGuardResponseBodyData) *MultiModalGuardResponseBody {
	s.Data = v
	return s
}

func (s *MultiModalGuardResponseBody) SetMessage(v string) *MultiModalGuardResponseBody {
	s.Message = &v
	return s
}

func (s *MultiModalGuardResponseBody) SetRequestId(v string) *MultiModalGuardResponseBody {
	s.RequestId = &v
	return s
}

func (s *MultiModalGuardResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MultiModalGuardResponseBodyData struct {
	// example:
	//
	// data1234
	DataId *string                                  `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Detail []*MultiModalGuardResponseBodyDataDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s MultiModalGuardResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardResponseBodyData) GoString() string {
	return s.String()
}

func (s *MultiModalGuardResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *MultiModalGuardResponseBodyData) GetDetail() []*MultiModalGuardResponseBodyDataDetail {
	return s.Detail
}

func (s *MultiModalGuardResponseBodyData) GetSuggestion() *string {
	return s.Suggestion
}

func (s *MultiModalGuardResponseBodyData) SetDataId(v string) *MultiModalGuardResponseBodyData {
	s.DataId = &v
	return s
}

func (s *MultiModalGuardResponseBodyData) SetDetail(v []*MultiModalGuardResponseBodyDataDetail) *MultiModalGuardResponseBodyData {
	s.Detail = v
	return s
}

func (s *MultiModalGuardResponseBodyData) SetSuggestion(v string) *MultiModalGuardResponseBodyData {
	s.Suggestion = &v
	return s
}

func (s *MultiModalGuardResponseBodyData) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MultiModalGuardResponseBodyDataDetail struct {
	// example:
	//
	// none
	Level  *string                                        `json:"Level,omitempty" xml:"Level,omitempty"`
	Result []*MultiModalGuardResponseBodyDataDetailResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// example:
	//
	// contentModeration
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MultiModalGuardResponseBodyDataDetail) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardResponseBodyDataDetail) GoString() string {
	return s.String()
}

func (s *MultiModalGuardResponseBodyDataDetail) GetLevel() *string {
	return s.Level
}

func (s *MultiModalGuardResponseBodyDataDetail) GetResult() []*MultiModalGuardResponseBodyDataDetailResult {
	return s.Result
}

func (s *MultiModalGuardResponseBodyDataDetail) GetSuggestion() *string {
	return s.Suggestion
}

func (s *MultiModalGuardResponseBodyDataDetail) GetType() *string {
	return s.Type
}

func (s *MultiModalGuardResponseBodyDataDetail) SetLevel(v string) *MultiModalGuardResponseBodyDataDetail {
	s.Level = &v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetail) SetResult(v []*MultiModalGuardResponseBodyDataDetailResult) *MultiModalGuardResponseBodyDataDetail {
	s.Result = v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetail) SetSuggestion(v string) *MultiModalGuardResponseBodyDataDetail {
	s.Suggestion = &v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetail) SetType(v string) *MultiModalGuardResponseBodyDataDetail {
	s.Type = &v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetail) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MultiModalGuardResponseBodyDataDetailResult struct {
	// example:
	//
	// 0
	Confidence  *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {}
	Ext interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// contraband_act
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// none
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s MultiModalGuardResponseBodyDataDetailResult) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardResponseBodyDataDetailResult) GoString() string {
	return s.String()
}

func (s *MultiModalGuardResponseBodyDataDetailResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *MultiModalGuardResponseBodyDataDetailResult) GetDescription() *string {
	return s.Description
}

func (s *MultiModalGuardResponseBodyDataDetailResult) GetExt() interface{} {
	return s.Ext
}

func (s *MultiModalGuardResponseBodyDataDetailResult) GetLabel() *string {
	return s.Label
}

func (s *MultiModalGuardResponseBodyDataDetailResult) GetLevel() *string {
	return s.Level
}

func (s *MultiModalGuardResponseBodyDataDetailResult) SetConfidence(v float32) *MultiModalGuardResponseBodyDataDetailResult {
	s.Confidence = &v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetailResult) SetDescription(v string) *MultiModalGuardResponseBodyDataDetailResult {
	s.Description = &v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetailResult) SetExt(v interface{}) *MultiModalGuardResponseBodyDataDetailResult {
	s.Ext = v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetailResult) SetLabel(v string) *MultiModalGuardResponseBodyDataDetailResult {
	s.Label = &v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetailResult) SetLevel(v string) *MultiModalGuardResponseBodyDataDetailResult {
	s.Level = &v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetailResult) Validate() error {
	return dara.Validate(s)
}
