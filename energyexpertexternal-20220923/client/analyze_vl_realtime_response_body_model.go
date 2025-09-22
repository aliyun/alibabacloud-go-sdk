// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeVlRealtimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AnalyzeVlRealtimeResponseBodyData) *AnalyzeVlRealtimeResponseBody
	GetData() *AnalyzeVlRealtimeResponseBodyData
	SetRequestId(v string) *AnalyzeVlRealtimeResponseBody
	GetRequestId() *string
}

type AnalyzeVlRealtimeResponseBody struct {
	// Return result.
	Data *AnalyzeVlRealtimeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AnalyzeVlRealtimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeVlRealtimeResponseBody) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeResponseBody) GetData() *AnalyzeVlRealtimeResponseBodyData {
	return s.Data
}

func (s *AnalyzeVlRealtimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AnalyzeVlRealtimeResponseBody) SetData(v *AnalyzeVlRealtimeResponseBodyData) *AnalyzeVlRealtimeResponseBody {
	s.Data = v
	return s
}

func (s *AnalyzeVlRealtimeResponseBody) SetRequestId(v string) *AnalyzeVlRealtimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnalyzeVlRealtimeResponseBody) Validate() error {
	return dara.Validate(s)
}

type AnalyzeVlRealtimeResponseBodyData struct {
	// Document parsing result details
	KvListInfo []*AnalyzeVlRealtimeResponseBodyDataKvListInfo `json:"kvListInfo,omitempty" xml:"kvListInfo,omitempty" type:"Repeated"`
}

func (s AnalyzeVlRealtimeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeVlRealtimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeResponseBodyData) GetKvListInfo() []*AnalyzeVlRealtimeResponseBodyDataKvListInfo {
	return s.KvListInfo
}

func (s *AnalyzeVlRealtimeResponseBodyData) SetKvListInfo(v []*AnalyzeVlRealtimeResponseBodyDataKvListInfo) *AnalyzeVlRealtimeResponseBodyData {
	s.KvListInfo = v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type AnalyzeVlRealtimeResponseBodyDataKvListInfo struct {
	// Recall content
	Context *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext `json:"context,omitempty" xml:"context,omitempty" type:"Struct"`
	// Field Key name
	//
	// example:
	//
	// username
	KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty"`
	// Field key value
	//
	// example:
	//
	// bob
	KeyValue *string `json:"keyValue,omitempty" xml:"keyValue,omitempty"`
}

func (s AnalyzeVlRealtimeResponseBodyDataKvListInfo) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeVlRealtimeResponseBodyDataKvListInfo) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfo) GetContext() *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext {
	return s.Context
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfo) GetKeyName() *string {
	return s.KeyName
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfo) GetKeyValue() *string {
	return s.KeyValue
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfo) SetContext(v *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) *AnalyzeVlRealtimeResponseBodyDataKvListInfo {
	s.Context = v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfo) SetKeyName(v string) *AnalyzeVlRealtimeResponseBodyDataKvListInfo {
	s.KeyName = &v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfo) SetKeyValue(v string) *AnalyzeVlRealtimeResponseBodyDataKvListInfo {
	s.KeyValue = &v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfo) Validate() error {
	return dara.Validate(s)
}

type AnalyzeVlRealtimeResponseBodyDataKvListInfoContext struct {
	// Confidence
	Confidence *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence `json:"confidence,omitempty" xml:"confidence,omitempty" type:"Struct"`
	// Key recall information details
	Key []*ContentItem `json:"key,omitempty" xml:"key,omitempty" type:"Repeated"`
	// Value recall information details
	Value []*ContentItem `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) GetConfidence() *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence {
	return s.Confidence
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) GetKey() []*ContentItem {
	return s.Key
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) GetValue() []*ContentItem {
	return s.Value
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) SetConfidence(v *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence) *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext {
	s.Confidence = v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) SetKey(v []*ContentItem) *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext {
	s.Key = v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) SetValue(v []*ContentItem) *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext {
	s.Value = v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) Validate() error {
	return dara.Validate(s)
}

type AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence struct {
	// Key confidence
	//
	// example:
	//
	// 0.9994202852249146
	KeyConfidence *float64 `json:"keyConfidence,omitempty" xml:"keyConfidence,omitempty"`
	// Value confidence
	//
	// example:
	//
	// 0.9794202852249146
	ValueConfidence *float64 `json:"valueConfidence,omitempty" xml:"valueConfidence,omitempty"`
}

func (s AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence) GetKeyConfidence() *float64 {
	return s.KeyConfidence
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence) GetValueConfidence() *float64 {
	return s.ValueConfidence
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence) SetKeyConfidence(v float64) *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence {
	s.KeyConfidence = &v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence) SetValueConfidence(v float64) *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence {
	s.ValueConfidence = &v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence) Validate() error {
	return dara.Validate(s)
}
