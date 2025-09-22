// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentAnalyzeResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDocumentAnalyzeResultResponseBodyData) *GetDocumentAnalyzeResultResponseBody
	GetData() *GetDocumentAnalyzeResultResponseBodyData
	SetRequestId(v string) *GetDocumentAnalyzeResultResponseBody
	GetRequestId() *string
}

type GetDocumentAnalyzeResultResponseBody struct {
	// Returned Data
	Data *GetDocumentAnalyzeResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDocumentAnalyzeResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalyzeResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultResponseBody) GetData() *GetDocumentAnalyzeResultResponseBodyData {
	return s.Data
}

func (s *GetDocumentAnalyzeResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentAnalyzeResultResponseBody) SetData(v *GetDocumentAnalyzeResultResponseBodyData) *GetDocumentAnalyzeResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBody) SetRequestId(v string) *GetDocumentAnalyzeResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDocumentAnalyzeResultResponseBodyData struct {
	// Document Parsing Result
	KvListInfo []*GetDocumentAnalyzeResultResponseBodyDataKvListInfo `json:"kvListInfo,omitempty" xml:"kvListInfo,omitempty" type:"Repeated"`
}

func (s GetDocumentAnalyzeResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalyzeResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultResponseBodyData) GetKvListInfo() []*GetDocumentAnalyzeResultResponseBodyDataKvListInfo {
	return s.KvListInfo
}

func (s *GetDocumentAnalyzeResultResponseBodyData) SetKvListInfo(v []*GetDocumentAnalyzeResultResponseBodyDataKvListInfo) *GetDocumentAnalyzeResultResponseBodyData {
	s.KvListInfo = v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDocumentAnalyzeResultResponseBodyDataKvListInfo struct {
	// Recalled Content
	Context *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext `json:"context,omitempty" xml:"context,omitempty" type:"Struct"`
	// Field Key Name
	//
	// example:
	//
	// Tenant
	KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty"`
	// Field Key Value
	//
	// example:
	//
	// Aliyun XXX Co., Ltd.
	KeyValue *string `json:"keyValue,omitempty" xml:"keyValue,omitempty"`
}

func (s GetDocumentAnalyzeResultResponseBodyDataKvListInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalyzeResultResponseBodyDataKvListInfo) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfo) GetContext() *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext {
	return s.Context
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfo) GetKeyName() *string {
	return s.KeyName
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfo) GetKeyValue() *string {
	return s.KeyValue
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfo) SetContext(v *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) *GetDocumentAnalyzeResultResponseBodyDataKvListInfo {
	s.Context = v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfo) SetKeyName(v string) *GetDocumentAnalyzeResultResponseBodyDataKvListInfo {
	s.KeyName = &v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfo) SetKeyValue(v string) *GetDocumentAnalyzeResultResponseBodyDataKvListInfo {
	s.KeyValue = &v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfo) Validate() error {
	return dara.Validate(s)
}

type GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext struct {
	// Confidence
	Confidence *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence `json:"confidence,omitempty" xml:"confidence,omitempty" type:"Struct"`
	// Key Recall Information
	Key []*ContentItem `json:"key,omitempty" xml:"key,omitempty" type:"Repeated"`
	// Value Recall Information
	Value []*ContentItem `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) GetConfidence() *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence {
	return s.Confidence
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) GetKey() []*ContentItem {
	return s.Key
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) GetValue() []*ContentItem {
	return s.Value
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) SetConfidence(v *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence) *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext {
	s.Confidence = v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) SetKey(v []*ContentItem) *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext {
	s.Key = v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) SetValue(v []*ContentItem) *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext {
	s.Value = v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) Validate() error {
	return dara.Validate(s)
}

type GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence struct {
	// Confidence of Key
	//
	// example:
	//
	// 0.9994202852249146
	KeyConfidence *float64 `json:"keyConfidence,omitempty" xml:"keyConfidence,omitempty"`
	// Confidence of Value
	//
	// example:
	//
	// 0.9794202852249146
	ValueConfidence *float64 `json:"valueConfidence,omitempty" xml:"valueConfidence,omitempty"`
}

func (s GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence) GetKeyConfidence() *float64 {
	return s.KeyConfidence
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence) GetValueConfidence() *float64 {
	return s.ValueConfidence
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence) SetKeyConfidence(v float64) *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence {
	s.KeyConfidence = &v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence) SetValueConfidence(v float64) *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence {
	s.ValueConfidence = &v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence) Validate() error {
	return dara.Validate(s)
}
