// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeIntentionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *RecognizeIntentionResponseBody
	GetCost() *int64
	SetData(v *RecognizeIntentionResponseBodyData) *RecognizeIntentionResponseBody
	GetData() *RecognizeIntentionResponseBodyData
	SetDataType(v string) *RecognizeIntentionResponseBody
	GetDataType() *string
	SetErrCode(v string) *RecognizeIntentionResponseBody
	GetErrCode() *string
	SetMessage(v string) *RecognizeIntentionResponseBody
	GetMessage() *string
	SetRequestId(v string) *RecognizeIntentionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RecognizeIntentionResponseBody
	GetSuccess() *bool
	SetTime(v string) *RecognizeIntentionResponseBody
	GetTime() *string
}

type RecognizeIntentionResponseBody struct {
	// example:
	//
	// null
	Cost *int64                              `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *RecognizeIntentionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 003D019A-1BB3-53EC-A0D2-CE76DA5D73B1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RecognizeIntentionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIntentionResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *RecognizeIntentionResponseBody) GetData() *RecognizeIntentionResponseBodyData {
	return s.Data
}

func (s *RecognizeIntentionResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *RecognizeIntentionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RecognizeIntentionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RecognizeIntentionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeIntentionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RecognizeIntentionResponseBody) GetTime() *string {
	return s.Time
}

func (s *RecognizeIntentionResponseBody) SetCost(v int64) *RecognizeIntentionResponseBody {
	s.Cost = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetData(v *RecognizeIntentionResponseBodyData) *RecognizeIntentionResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeIntentionResponseBody) SetDataType(v string) *RecognizeIntentionResponseBody {
	s.DataType = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetErrCode(v string) *RecognizeIntentionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetMessage(v string) *RecognizeIntentionResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetRequestId(v string) *RecognizeIntentionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetSuccess(v bool) *RecognizeIntentionResponseBody {
	s.Success = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetTime(v string) *RecognizeIntentionResponseBody {
	s.Time = &v
	return s
}

func (s *RecognizeIntentionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RecognizeIntentionResponseBodyData struct {
	AnalysisProcess *string `json:"analysisProcess,omitempty" xml:"analysisProcess,omitempty"`
	// example:
	//
	// 1
	IntentionCode      *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	IntentionName      *string `json:"intentionName,omitempty" xml:"intentionName,omitempty"`
	IntentionScript    *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
	RecommendIntention *string `json:"recommendIntention,omitempty" xml:"recommendIntention,omitempty"`
	RecommendScript    *string `json:"recommendScript,omitempty" xml:"recommendScript,omitempty"`
}

func (s RecognizeIntentionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIntentionResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionResponseBodyData) GetAnalysisProcess() *string {
	return s.AnalysisProcess
}

func (s *RecognizeIntentionResponseBodyData) GetIntentionCode() *string {
	return s.IntentionCode
}

func (s *RecognizeIntentionResponseBodyData) GetIntentionName() *string {
	return s.IntentionName
}

func (s *RecognizeIntentionResponseBodyData) GetIntentionScript() *string {
	return s.IntentionScript
}

func (s *RecognizeIntentionResponseBodyData) GetRecommendIntention() *string {
	return s.RecommendIntention
}

func (s *RecognizeIntentionResponseBodyData) GetRecommendScript() *string {
	return s.RecommendScript
}

func (s *RecognizeIntentionResponseBodyData) SetAnalysisProcess(v string) *RecognizeIntentionResponseBodyData {
	s.AnalysisProcess = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetIntentionCode(v string) *RecognizeIntentionResponseBodyData {
	s.IntentionCode = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetIntentionName(v string) *RecognizeIntentionResponseBodyData {
	s.IntentionName = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetIntentionScript(v string) *RecognizeIntentionResponseBodyData {
	s.IntentionScript = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetRecommendIntention(v string) *RecognizeIntentionResponseBodyData {
	s.RecommendIntention = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetRecommendScript(v string) *RecognizeIntentionResponseBodyData {
	s.RecommendScript = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
