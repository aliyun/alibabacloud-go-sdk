// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardForBase64ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *MultiModalGuardForBase64ResponseBody
	GetCode() *int32
	SetData(v *MultiModalGuardForBase64ResponseBodyData) *MultiModalGuardForBase64ResponseBody
	GetData() *MultiModalGuardForBase64ResponseBodyData
	SetMessage(v string) *MultiModalGuardForBase64ResponseBody
	GetMessage() *string
	SetRequestId(v string) *MultiModalGuardForBase64ResponseBody
	GetRequestId() *string
}

type MultiModalGuardForBase64ResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *MultiModalGuardForBase64ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MultiModalGuardForBase64ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardForBase64ResponseBody) GoString() string {
	return s.String()
}

func (s *MultiModalGuardForBase64ResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *MultiModalGuardForBase64ResponseBody) GetData() *MultiModalGuardForBase64ResponseBodyData {
	return s.Data
}

func (s *MultiModalGuardForBase64ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MultiModalGuardForBase64ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MultiModalGuardForBase64ResponseBody) SetCode(v int32) *MultiModalGuardForBase64ResponseBody {
	s.Code = &v
	return s
}

func (s *MultiModalGuardForBase64ResponseBody) SetData(v *MultiModalGuardForBase64ResponseBodyData) *MultiModalGuardForBase64ResponseBody {
	s.Data = v
	return s
}

func (s *MultiModalGuardForBase64ResponseBody) SetMessage(v string) *MultiModalGuardForBase64ResponseBody {
	s.Message = &v
	return s
}

func (s *MultiModalGuardForBase64ResponseBody) SetRequestId(v string) *MultiModalGuardForBase64ResponseBody {
	s.RequestId = &v
	return s
}

func (s *MultiModalGuardForBase64ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MultiModalGuardForBase64ResponseBodyData struct {
	// example:
	//
	// xxx
	DataId *string                                           `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Detail []*MultiModalGuardForBase64ResponseBodyDataDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s MultiModalGuardForBase64ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardForBase64ResponseBodyData) GoString() string {
	return s.String()
}

func (s *MultiModalGuardForBase64ResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *MultiModalGuardForBase64ResponseBodyData) GetDetail() []*MultiModalGuardForBase64ResponseBodyDataDetail {
	return s.Detail
}

func (s *MultiModalGuardForBase64ResponseBodyData) GetSuggestion() *string {
	return s.Suggestion
}

func (s *MultiModalGuardForBase64ResponseBodyData) SetDataId(v string) *MultiModalGuardForBase64ResponseBodyData {
	s.DataId = &v
	return s
}

func (s *MultiModalGuardForBase64ResponseBodyData) SetDetail(v []*MultiModalGuardForBase64ResponseBodyDataDetail) *MultiModalGuardForBase64ResponseBodyData {
	s.Detail = v
	return s
}

func (s *MultiModalGuardForBase64ResponseBodyData) SetSuggestion(v string) *MultiModalGuardForBase64ResponseBodyData {
	s.Suggestion = &v
	return s
}

func (s *MultiModalGuardForBase64ResponseBodyData) Validate() error {
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

type MultiModalGuardForBase64ResponseBodyDataDetail struct {
	// example:
	//
	// low
	Level  *string                                                 `json:"Level,omitempty" xml:"Level,omitempty"`
	Result []*MultiModalGuardForBase64ResponseBodyDataDetailResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// example:
	//
	// contentModeration
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MultiModalGuardForBase64ResponseBodyDataDetail) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardForBase64ResponseBodyDataDetail) GoString() string {
	return s.String()
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetail) GetLevel() *string {
	return s.Level
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetail) GetResult() []*MultiModalGuardForBase64ResponseBodyDataDetailResult {
	return s.Result
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetail) GetSuggestion() *string {
	return s.Suggestion
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetail) GetType() *string {
	return s.Type
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetail) SetLevel(v string) *MultiModalGuardForBase64ResponseBodyDataDetail {
	s.Level = &v
	return s
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetail) SetResult(v []*MultiModalGuardForBase64ResponseBodyDataDetailResult) *MultiModalGuardForBase64ResponseBodyDataDetail {
	s.Result = v
	return s
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetail) SetSuggestion(v string) *MultiModalGuardForBase64ResponseBodyDataDetail {
	s.Suggestion = &v
	return s
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetail) SetType(v string) *MultiModalGuardForBase64ResponseBodyDataDetail {
	s.Type = &v
	return s
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetail) Validate() error {
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

type MultiModalGuardForBase64ResponseBodyDataDetailResult struct {
	// example:
	//
	// 100
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// 未检测出风险
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// json格式数据
	Ext interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// nonLable
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// low
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s MultiModalGuardForBase64ResponseBodyDataDetailResult) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardForBase64ResponseBodyDataDetailResult) GoString() string {
	return s.String()
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetailResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetailResult) GetDescription() *string {
	return s.Description
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetailResult) GetExt() interface{} {
	return s.Ext
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetailResult) GetLabel() *string {
	return s.Label
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetailResult) GetLevel() *string {
	return s.Level
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetailResult) SetConfidence(v float32) *MultiModalGuardForBase64ResponseBodyDataDetailResult {
	s.Confidence = &v
	return s
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetailResult) SetDescription(v string) *MultiModalGuardForBase64ResponseBodyDataDetailResult {
	s.Description = &v
	return s
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetailResult) SetExt(v interface{}) *MultiModalGuardForBase64ResponseBodyDataDetailResult {
	s.Ext = v
	return s
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetailResult) SetLabel(v string) *MultiModalGuardForBase64ResponseBodyDataDetailResult {
	s.Label = &v
	return s
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetailResult) SetLevel(v string) *MultiModalGuardForBase64ResponseBodyDataDetailResult {
	s.Level = &v
	return s
}

func (s *MultiModalGuardForBase64ResponseBodyDataDetailResult) Validate() error {
	return dara.Validate(s)
}
