// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocExtractionResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDocExtractionResultResponseBodyData) *GetDocExtractionResultResponseBody
	GetData() *GetDocExtractionResultResponseBodyData
	SetRequestId(v string) *GetDocExtractionResultResponseBody
	GetRequestId() *string
}

type GetDocExtractionResultResponseBody struct {
	// Returned data structure.
	Data *GetDocExtractionResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// ID of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDocExtractionResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocExtractionResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultResponseBody) GetData() *GetDocExtractionResultResponseBodyData {
	return s.Data
}

func (s *GetDocExtractionResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocExtractionResultResponseBody) SetData(v *GetDocExtractionResultResponseBodyData) *GetDocExtractionResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocExtractionResultResponseBody) SetRequestId(v string) *GetDocExtractionResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocExtractionResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDocExtractionResultResponseBodyData struct {
	// Details of document extraction results
	KvListInfo []*GetDocExtractionResultResponseBodyDataKvListInfo `json:"kvListInfo,omitempty" xml:"kvListInfo,omitempty" type:"Repeated"`
}

func (s GetDocExtractionResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocExtractionResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultResponseBodyData) GetKvListInfo() []*GetDocExtractionResultResponseBodyDataKvListInfo {
	return s.KvListInfo
}

func (s *GetDocExtractionResultResponseBodyData) SetKvListInfo(v []*GetDocExtractionResultResponseBodyDataKvListInfo) *GetDocExtractionResultResponseBodyData {
	s.KvListInfo = v
	return s
}

func (s *GetDocExtractionResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDocExtractionResultResponseBodyDataKvListInfo struct {
	// Recalled content
	Context *GetDocExtractionResultResponseBodyDataKvListInfoContext `json:"context,omitempty" xml:"context,omitempty" type:"Struct"`
	// Field key name
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

func (s GetDocExtractionResultResponseBodyDataKvListInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDocExtractionResultResponseBodyDataKvListInfo) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfo) GetContext() *GetDocExtractionResultResponseBodyDataKvListInfoContext {
	return s.Context
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfo) GetKeyName() *string {
	return s.KeyName
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfo) GetKeyValue() *string {
	return s.KeyValue
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfo) SetContext(v *GetDocExtractionResultResponseBodyDataKvListInfoContext) *GetDocExtractionResultResponseBodyDataKvListInfo {
	s.Context = v
	return s
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfo) SetKeyName(v string) *GetDocExtractionResultResponseBodyDataKvListInfo {
	s.KeyName = &v
	return s
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfo) SetKeyValue(v string) *GetDocExtractionResultResponseBodyDataKvListInfo {
	s.KeyValue = &v
	return s
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfo) Validate() error {
	return dara.Validate(s)
}

type GetDocExtractionResultResponseBodyDataKvListInfoContext struct {
	// Confidence level
	Confidence *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence `json:"confidence,omitempty" xml:"confidence,omitempty" type:"Struct"`
	// Details of key recall information
	Key []*ContentItem `json:"key,omitempty" xml:"key,omitempty" type:"Repeated"`
	// Details of value recall information
	Value []*ContentItem `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s GetDocExtractionResultResponseBodyDataKvListInfoContext) String() string {
	return dara.Prettify(s)
}

func (s GetDocExtractionResultResponseBodyDataKvListInfoContext) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContext) GetConfidence() *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence {
	return s.Confidence
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContext) GetKey() []*ContentItem {
	return s.Key
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContext) GetValue() []*ContentItem {
	return s.Value
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContext) SetConfidence(v *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence) *GetDocExtractionResultResponseBodyDataKvListInfoContext {
	s.Confidence = v
	return s
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContext) SetKey(v []*ContentItem) *GetDocExtractionResultResponseBodyDataKvListInfoContext {
	s.Key = v
	return s
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContext) SetValue(v []*ContentItem) *GetDocExtractionResultResponseBodyDataKvListInfoContext {
	s.Value = v
	return s
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContext) Validate() error {
	return dara.Validate(s)
}

type GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence struct {
	// Key confidence level
	//
	// example:
	//
	// 0.9994202852249146
	KeyConfidence *float64 `json:"keyConfidence,omitempty" xml:"keyConfidence,omitempty"`
	// value confidence level
	//
	// example:
	//
	// 0.9794202852249146
	ValueConfidence *float64 `json:"valueConfidence,omitempty" xml:"valueConfidence,omitempty"`
}

func (s GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence) String() string {
	return dara.Prettify(s)
}

func (s GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence) GetKeyConfidence() *float64 {
	return s.KeyConfidence
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence) GetValueConfidence() *float64 {
	return s.ValueConfidence
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence) SetKeyConfidence(v float64) *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence {
	s.KeyConfidence = &v
	return s
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence) SetValueConfidence(v float64) *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence {
	s.ValueConfidence = &v
	return s
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence) Validate() error {
	return dara.Validate(s)
}
