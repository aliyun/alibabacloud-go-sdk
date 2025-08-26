// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeQuotaInvoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeQuotaInvoiceResponseBodyData) *RecognizeQuotaInvoiceResponseBody
	GetData() *RecognizeQuotaInvoiceResponseBodyData
	SetRequestId(v string) *RecognizeQuotaInvoiceResponseBody
	GetRequestId() *string
}

type RecognizeQuotaInvoiceResponseBody struct {
	Data *RecognizeQuotaInvoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// BC4C12D0-7FD3-419A-B997-A91212DF6D82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeQuotaInvoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponseBody) GetData() *RecognizeQuotaInvoiceResponseBodyData {
	return s.Data
}

func (s *RecognizeQuotaInvoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeQuotaInvoiceResponseBody) SetData(v *RecognizeQuotaInvoiceResponseBodyData) *RecognizeQuotaInvoiceResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBody) SetRequestId(v string) *RecognizeQuotaInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeQuotaInvoiceResponseBodyData struct {
	// example:
	//
	// 1
	Angle   *int64                                        `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Content *RecognizeQuotaInvoiceResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// example:
	//
	// 624
	Height        *int64                                                `json:"Height,omitempty" xml:"Height,omitempty"`
	KeyValueInfos []*RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos `json:"KeyValueInfos,omitempty" xml:"KeyValueInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 610
	OrgHeight *int64 `json:"OrgHeight,omitempty" xml:"OrgHeight,omitempty"`
	// example:
	//
	// 855
	OrgWidth *int64 `json:"OrgWidth,omitempty" xml:"OrgWidth,omitempty"`
	// example:
	//
	// 865
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeQuotaInvoiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponseBodyData) GetAngle() *int64 {
	return s.Angle
}

func (s *RecognizeQuotaInvoiceResponseBodyData) GetContent() *RecognizeQuotaInvoiceResponseBodyDataContent {
	return s.Content
}

func (s *RecognizeQuotaInvoiceResponseBodyData) GetHeight() *int64 {
	return s.Height
}

func (s *RecognizeQuotaInvoiceResponseBodyData) GetKeyValueInfos() []*RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos {
	return s.KeyValueInfos
}

func (s *RecognizeQuotaInvoiceResponseBodyData) GetOrgHeight() *int64 {
	return s.OrgHeight
}

func (s *RecognizeQuotaInvoiceResponseBodyData) GetOrgWidth() *int64 {
	return s.OrgWidth
}

func (s *RecognizeQuotaInvoiceResponseBodyData) GetWidth() *int64 {
	return s.Width
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetAngle(v int64) *RecognizeQuotaInvoiceResponseBodyData {
	s.Angle = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetContent(v *RecognizeQuotaInvoiceResponseBodyDataContent) *RecognizeQuotaInvoiceResponseBodyData {
	s.Content = v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetHeight(v int64) *RecognizeQuotaInvoiceResponseBodyData {
	s.Height = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetKeyValueInfos(v []*RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) *RecognizeQuotaInvoiceResponseBodyData {
	s.KeyValueInfos = v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetOrgHeight(v int64) *RecognizeQuotaInvoiceResponseBodyData {
	s.OrgHeight = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetOrgWidth(v int64) *RecognizeQuotaInvoiceResponseBodyData {
	s.OrgWidth = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetWidth(v int64) *RecognizeQuotaInvoiceResponseBodyData {
	s.Width = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecognizeQuotaInvoiceResponseBodyDataContent struct {
	// example:
	//
	// 10
	InvoiceAmount *string `json:"InvoiceAmount,omitempty" xml:"InvoiceAmount,omitempty"`
	// example:
	//
	// 144031800103
	InvoiceCode    *string `json:"InvoiceCode,omitempty" xml:"InvoiceCode,omitempty"`
	InvoiceDetails *string `json:"InvoiceDetails,omitempty" xml:"InvoiceDetails,omitempty"`
	// example:
	//
	// 40637706
	InvoiceNo *string `json:"InvoiceNo,omitempty" xml:"InvoiceNo,omitempty"`
	SumAmount *string `json:"SumAmount,omitempty" xml:"SumAmount,omitempty"`
}

func (s RecognizeQuotaInvoiceResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) GetInvoiceAmount() *string {
	return s.InvoiceAmount
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) GetInvoiceCode() *string {
	return s.InvoiceCode
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) GetInvoiceDetails() *string {
	return s.InvoiceDetails
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) GetInvoiceNo() *string {
	return s.InvoiceNo
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) GetSumAmount() *string {
	return s.SumAmount
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) SetInvoiceAmount(v string) *RecognizeQuotaInvoiceResponseBodyDataContent {
	s.InvoiceAmount = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) SetInvoiceCode(v string) *RecognizeQuotaInvoiceResponseBodyDataContent {
	s.InvoiceCode = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) SetInvoiceDetails(v string) *RecognizeQuotaInvoiceResponseBodyDataContent {
	s.InvoiceDetails = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) SetInvoiceNo(v string) *RecognizeQuotaInvoiceResponseBodyDataContent {
	s.InvoiceNo = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) SetSumAmount(v string) *RecognizeQuotaInvoiceResponseBodyDataContent {
	s.SumAmount = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}

type RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos struct {
	Key            *string                                                             `json:"Key,omitempty" xml:"Key,omitempty"`
	Value          *string                                                             `json:"Value,omitempty" xml:"Value,omitempty"`
	ValuePositions []*RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions `json:"ValuePositions,omitempty" xml:"ValuePositions,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	ValueScore *float32 `json:"ValueScore,omitempty" xml:"ValueScore,omitempty"`
}

func (s RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) GetKey() *string {
	return s.Key
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) GetValue() *string {
	return s.Value
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) GetValuePositions() []*RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions {
	return s.ValuePositions
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) GetValueScore() *float32 {
	return s.ValueScore
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) SetKey(v string) *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos {
	s.Key = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) SetValue(v string) *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos {
	s.Value = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) SetValuePositions(v []*RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions) *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos {
	s.ValuePositions = v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) SetValueScore(v float32) *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos {
	s.ValueScore = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) Validate() error {
	return dara.Validate(s)
}

type RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions struct {
	// example:
	//
	// 544
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 387
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions) GetX() *int64 {
	return s.X
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions) GetY() *int64 {
	return s.Y
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions) SetX(v int64) *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions {
	s.X = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions) SetY(v int64) *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions {
	s.Y = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions) Validate() error {
	return dara.Validate(s)
}
