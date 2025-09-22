// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVLExtractionResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetVLExtractionResultResponseBodyData) *GetVLExtractionResultResponseBody
	GetData() *GetVLExtractionResultResponseBodyData
	SetRequestId(v string) *GetVLExtractionResultResponseBody
	GetRequestId() *string
}

type GetVLExtractionResultResponseBody struct {
	// Returned Data
	Data *GetVLExtractionResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetVLExtractionResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVLExtractionResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultResponseBody) GetData() *GetVLExtractionResultResponseBodyData {
	return s.Data
}

func (s *GetVLExtractionResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVLExtractionResultResponseBody) SetData(v *GetVLExtractionResultResponseBodyData) *GetVLExtractionResultResponseBody {
	s.Data = v
	return s
}

func (s *GetVLExtractionResultResponseBody) SetRequestId(v string) *GetVLExtractionResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVLExtractionResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVLExtractionResultResponseBodyData struct {
	// Document Parsing Result
	KvListInfo []*GetVLExtractionResultResponseBodyDataKvListInfo `json:"kvListInfo,omitempty" xml:"kvListInfo,omitempty" type:"Repeated"`
}

func (s GetVLExtractionResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVLExtractionResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultResponseBodyData) GetKvListInfo() []*GetVLExtractionResultResponseBodyDataKvListInfo {
	return s.KvListInfo
}

func (s *GetVLExtractionResultResponseBodyData) SetKvListInfo(v []*GetVLExtractionResultResponseBodyDataKvListInfo) *GetVLExtractionResultResponseBodyData {
	s.KvListInfo = v
	return s
}

func (s *GetVLExtractionResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetVLExtractionResultResponseBodyDataKvListInfo struct {
	// Recall content
	Context *GetVLExtractionResultResponseBodyDataKvListInfoContext `json:"context,omitempty" xml:"context,omitempty" type:"Struct"`
	// Field Key name
	//
	// example:
	//
	// Tenant
	KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty"`
	// Field key value
	//
	// example:
	//
	// Alibaba Cloud XXX Co., Ltd.
	KeyValue *string `json:"keyValue,omitempty" xml:"keyValue,omitempty"`
}

func (s GetVLExtractionResultResponseBodyDataKvListInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVLExtractionResultResponseBodyDataKvListInfo) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfo) GetContext() *GetVLExtractionResultResponseBodyDataKvListInfoContext {
	return s.Context
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfo) GetKeyName() *string {
	return s.KeyName
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfo) GetKeyValue() *string {
	return s.KeyValue
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfo) SetContext(v *GetVLExtractionResultResponseBodyDataKvListInfoContext) *GetVLExtractionResultResponseBodyDataKvListInfo {
	s.Context = v
	return s
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfo) SetKeyName(v string) *GetVLExtractionResultResponseBodyDataKvListInfo {
	s.KeyName = &v
	return s
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfo) SetKeyValue(v string) *GetVLExtractionResultResponseBodyDataKvListInfo {
	s.KeyValue = &v
	return s
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfo) Validate() error {
	return dara.Validate(s)
}

type GetVLExtractionResultResponseBodyDataKvListInfoContext struct {
	// Confidence
	Confidence *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence `json:"confidence,omitempty" xml:"confidence,omitempty" type:"Struct"`
	// Key recall information details
	Key []*ContentItem `json:"key,omitempty" xml:"key,omitempty" type:"Repeated"`
	// Value Recall Information
	Value []*ContentItem `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s GetVLExtractionResultResponseBodyDataKvListInfoContext) String() string {
	return dara.Prettify(s)
}

func (s GetVLExtractionResultResponseBodyDataKvListInfoContext) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContext) GetConfidence() *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence {
	return s.Confidence
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContext) GetKey() []*ContentItem {
	return s.Key
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContext) GetValue() []*ContentItem {
	return s.Value
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContext) SetConfidence(v *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence) *GetVLExtractionResultResponseBodyDataKvListInfoContext {
	s.Confidence = v
	return s
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContext) SetKey(v []*ContentItem) *GetVLExtractionResultResponseBodyDataKvListInfoContext {
	s.Key = v
	return s
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContext) SetValue(v []*ContentItem) *GetVLExtractionResultResponseBodyDataKvListInfoContext {
	s.Value = v
	return s
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContext) Validate() error {
	return dara.Validate(s)
}

type GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence struct {
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

func (s GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence) String() string {
	return dara.Prettify(s)
}

func (s GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence) GetKeyConfidence() *float64 {
	return s.KeyConfidence
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence) GetValueConfidence() *float64 {
	return s.ValueConfidence
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence) SetKeyConfidence(v float64) *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence {
	s.KeyConfidence = &v
	return s
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence) SetValueConfidence(v float64) *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence {
	s.ValueConfidence = &v
	return s
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence) Validate() error {
	return dara.Validate(s)
}
