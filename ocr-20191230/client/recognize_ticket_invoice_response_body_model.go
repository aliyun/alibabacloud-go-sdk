// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeTicketInvoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeTicketInvoiceResponseBodyData) *RecognizeTicketInvoiceResponseBody
	GetData() *RecognizeTicketInvoiceResponseBodyData
	SetRequestId(v string) *RecognizeTicketInvoiceResponseBody
	GetRequestId() *string
}

type RecognizeTicketInvoiceResponseBody struct {
	Data *RecognizeTicketInvoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 063C0178-7EA3-4754-96FB-C0C9AE6B9AAE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBody) GetData() *RecognizeTicketInvoiceResponseBodyData {
	return s.Data
}

func (s *RecognizeTicketInvoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeTicketInvoiceResponseBody) SetData(v *RecognizeTicketInvoiceResponseBodyData) *RecognizeTicketInvoiceResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeTicketInvoiceResponseBody) SetRequestId(v string) *RecognizeTicketInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeTicketInvoiceResponseBodyData struct {
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 594
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 1417
	OrgHeight *int64 `json:"OrgHeight,omitempty" xml:"OrgHeight,omitempty"`
	// example:
	//
	// 1417
	OrgWidth *int64                                           `json:"OrgWidth,omitempty" xml:"OrgWidth,omitempty"`
	Results  []*RecognizeTicketInvoiceResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// 594
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *RecognizeTicketInvoiceResponseBodyData) GetHeight() *int64 {
	return s.Height
}

func (s *RecognizeTicketInvoiceResponseBodyData) GetOrgHeight() *int64 {
	return s.OrgHeight
}

func (s *RecognizeTicketInvoiceResponseBodyData) GetOrgWidth() *int64 {
	return s.OrgWidth
}

func (s *RecognizeTicketInvoiceResponseBodyData) GetResults() []*RecognizeTicketInvoiceResponseBodyDataResults {
	return s.Results
}

func (s *RecognizeTicketInvoiceResponseBodyData) GetWidth() *int64 {
	return s.Width
}

func (s *RecognizeTicketInvoiceResponseBodyData) SetCount(v int64) *RecognizeTicketInvoiceResponseBodyData {
	s.Count = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyData) SetHeight(v int64) *RecognizeTicketInvoiceResponseBodyData {
	s.Height = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyData) SetOrgHeight(v int64) *RecognizeTicketInvoiceResponseBodyData {
	s.OrgHeight = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyData) SetOrgWidth(v int64) *RecognizeTicketInvoiceResponseBodyData {
	s.OrgWidth = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyData) SetResults(v []*RecognizeTicketInvoiceResponseBodyDataResults) *RecognizeTicketInvoiceResponseBodyData {
	s.Results = v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyData) SetWidth(v int64) *RecognizeTicketInvoiceResponseBodyData {
	s.Width = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecognizeTicketInvoiceResponseBodyDataResults struct {
	Content *RecognizeTicketInvoiceResponseBodyDataResultsContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Index          *int64                                                         `json:"Index,omitempty" xml:"Index,omitempty"`
	KeyValueInfos  []*RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos  `json:"KeyValueInfos,omitempty" xml:"KeyValueInfos,omitempty" type:"Repeated"`
	SliceRectangle []*RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle `json:"SliceRectangle,omitempty" xml:"SliceRectangle,omitempty" type:"Repeated"`
	Type           *string                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) GetContent() *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	return s.Content
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) GetIndex() *int64 {
	return s.Index
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) GetKeyValueInfos() []*RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos {
	return s.KeyValueInfos
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) GetSliceRectangle() []*RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle {
	return s.SliceRectangle
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) GetType() *string {
	return s.Type
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) SetContent(v *RecognizeTicketInvoiceResponseBodyDataResultsContent) *RecognizeTicketInvoiceResponseBodyDataResults {
	s.Content = v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) SetIndex(v int64) *RecognizeTicketInvoiceResponseBodyDataResults {
	s.Index = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) SetKeyValueInfos(v []*RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) *RecognizeTicketInvoiceResponseBodyDataResults {
	s.KeyValueInfos = v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) SetSliceRectangle(v []*RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle) *RecognizeTicketInvoiceResponseBodyDataResults {
	s.SliceRectangle = v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) SetType(v string) *RecognizeTicketInvoiceResponseBodyDataResults {
	s.Type = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}

type RecognizeTicketInvoiceResponseBodyDataResultsContent struct {
	// example:
	//
	// 81931914902643039780
	AntiFakeCode *string `json:"AntiFakeCode,omitempty" xml:"AntiFakeCode,omitempty"`
	// example:
	//
	// 044031860107
	InvoiceCode *string `json:"InvoiceCode,omitempty" xml:"InvoiceCode,omitempty"`
	// example:
	//
	// 2018-09-20
	InvoiceDate *string `json:"InvoiceDate,omitempty" xml:"InvoiceDate,omitempty"`
	// example:
	//
	// 09267581
	InvoiceNumber *string `json:"InvoiceNumber,omitempty" xml:"InvoiceNumber,omitempty"`
	PayeeName     *string `json:"PayeeName,omitempty" xml:"PayeeName,omitempty"`
	// example:
	//
	// 914403002794492693
	PayeeRegisterNo *string `json:"PayeeRegisterNo,omitempty" xml:"PayeeRegisterNo,omitempty"`
	PayerName       *string `json:"PayerName,omitempty" xml:"PayerName,omitempty"`
	// example:
	//
	// 91440300MA5EXWHW6F
	PayerRegisterNo *string `json:"PayerRegisterNo,omitempty" xml:"PayerRegisterNo,omitempty"`
	// example:
	//
	// ￥220.00
	SumAmount *string `json:"SumAmount,omitempty" xml:"SumAmount,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsContent) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsContent) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) GetAntiFakeCode() *string {
	return s.AntiFakeCode
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) GetInvoiceCode() *string {
	return s.InvoiceCode
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) GetInvoiceDate() *string {
	return s.InvoiceDate
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) GetInvoiceNumber() *string {
	return s.InvoiceNumber
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) GetPayeeName() *string {
	return s.PayeeName
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) GetPayeeRegisterNo() *string {
	return s.PayeeRegisterNo
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) GetPayerName() *string {
	return s.PayerName
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) GetPayerRegisterNo() *string {
	return s.PayerRegisterNo
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) GetSumAmount() *string {
	return s.SumAmount
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetAntiFakeCode(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.AntiFakeCode = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetInvoiceCode(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.InvoiceCode = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetInvoiceDate(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.InvoiceDate = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetInvoiceNumber(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.InvoiceNumber = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetPayeeName(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.PayeeName = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetPayeeRegisterNo(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.PayeeRegisterNo = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetPayerName(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.PayerName = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetPayerRegisterNo(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.PayerRegisterNo = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetSumAmount(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.SumAmount = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) Validate() error {
	return dara.Validate(s)
}

type RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos struct {
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 044031860107
	Value          *string                                                                     `json:"Value,omitempty" xml:"Value,omitempty"`
	ValuePositions []*RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions `json:"ValuePositions,omitempty" xml:"ValuePositions,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	ValueScore *float32 `json:"ValueScore,omitempty" xml:"ValueScore,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) GetKey() *string {
	return s.Key
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) GetValue() *string {
	return s.Value
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) GetValuePositions() []*RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions {
	return s.ValuePositions
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) GetValueScore() *float32 {
	return s.ValueScore
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) SetKey(v string) *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos {
	s.Key = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) SetValue(v string) *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos {
	s.Value = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) SetValuePositions(v []*RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions) *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos {
	s.ValuePositions = v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) SetValueScore(v float32) *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos {
	s.ValueScore = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) Validate() error {
	return dara.Validate(s)
}

type RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions struct {
	// example:
	//
	// 586
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 16
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions) GetX() *int64 {
	return s.X
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions) GetY() *int64 {
	return s.Y
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions) SetX(v int64) *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions {
	s.X = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions) SetY(v int64) *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions {
	s.Y = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions) Validate() error {
	return dara.Validate(s)
}

type RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle struct {
	// example:
	//
	// 586
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 16
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle) GetX() *int64 {
	return s.X
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle) GetY() *int64 {
	return s.Y
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle) SetX(v int64) *RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle {
	s.X = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle) SetY(v int64) *RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle {
	s.Y = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle) Validate() error {
	return dara.Validate(s)
}
