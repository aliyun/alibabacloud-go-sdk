// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeTaxiInvoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeTaxiInvoiceResponseBodyData) *RecognizeTaxiInvoiceResponseBody
	GetData() *RecognizeTaxiInvoiceResponseBodyData
	SetRequestId(v string) *RecognizeTaxiInvoiceResponseBody
	GetRequestId() *string
}

type RecognizeTaxiInvoiceResponseBody struct {
	Data *RecognizeTaxiInvoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// B2BBBD26-1D3E-4CFA-A80B-6A9266B8D125
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBody) GetData() *RecognizeTaxiInvoiceResponseBodyData {
	return s.Data
}

func (s *RecognizeTaxiInvoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeTaxiInvoiceResponseBody) SetData(v *RecognizeTaxiInvoiceResponseBodyData) *RecognizeTaxiInvoiceResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBody) SetRequestId(v string) *RecognizeTaxiInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeTaxiInvoiceResponseBodyData struct {
	Invoices []*RecognizeTaxiInvoiceResponseBodyDataInvoices `json:"Invoices,omitempty" xml:"Invoices,omitempty" type:"Repeated"`
}

func (s RecognizeTaxiInvoiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyData) GetInvoices() []*RecognizeTaxiInvoiceResponseBodyDataInvoices {
	return s.Invoices
}

func (s *RecognizeTaxiInvoiceResponseBodyData) SetInvoices(v []*RecognizeTaxiInvoiceResponseBodyDataInvoices) *RecognizeTaxiInvoiceResponseBodyData {
	s.Invoices = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecognizeTaxiInvoiceResponseBodyDataInvoices struct {
	InvoiceRoi *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi `json:"InvoiceRoi,omitempty" xml:"InvoiceRoi,omitempty" type:"Struct"`
	Items      []*RecognizeTaxiInvoiceResponseBodyDataInvoicesItems    `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	RotateType *int32 `json:"RotateType,omitempty" xml:"RotateType,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoices) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoices) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoices) GetInvoiceRoi() *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi {
	return s.InvoiceRoi
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoices) GetItems() []*RecognizeTaxiInvoiceResponseBodyDataInvoicesItems {
	return s.Items
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoices) GetRotateType() *int32 {
	return s.RotateType
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoices) SetInvoiceRoi(v *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) *RecognizeTaxiInvoiceResponseBodyDataInvoices {
	s.InvoiceRoi = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoices) SetItems(v []*RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) *RecognizeTaxiInvoiceResponseBodyDataInvoices {
	s.Items = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoices) SetRotateType(v int32) *RecognizeTaxiInvoiceResponseBodyDataInvoices {
	s.RotateType = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoices) Validate() error {
	return dara.Validate(s)
}

type RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi struct {
	// example:
	//
	// 3625
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
	// example:
	//
	// 1773
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
	// example:
	//
	// 513
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 302
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) GetH() *float32 {
	return s.H
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) GetW() *float32 {
	return s.W
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) GetX() *float32 {
	return s.X
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) GetY() *float32 {
	return s.Y
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) SetH(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi {
	s.H = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) SetW(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi {
	s.W = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) SetX(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi {
	s.X = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) SetY(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi {
	s.Y = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) Validate() error {
	return dara.Validate(s)
}

type RecognizeTaxiInvoiceResponseBodyDataInvoicesItems struct {
	ItemRoi *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi `json:"ItemRoi,omitempty" xml:"ItemRoi,omitempty" type:"Struct"`
	// example:
	//
	// 86655664
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) GetItemRoi() *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi {
	return s.ItemRoi
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) GetText() *string {
	return s.Text
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) SetItemRoi(v *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItems {
	s.ItemRoi = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) SetText(v string) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItems {
	s.Text = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) Validate() error {
	return dara.Validate(s)
}

type RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi struct {
	// example:
	//
	// -90
	Angle  *float32                                                        `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Center *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter `json:"Center,omitempty" xml:"Center,omitempty" type:"Struct"`
	Size   *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize   `json:"Size,omitempty" xml:"Size,omitempty" type:"Struct"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) GetAngle() *float32 {
	return s.Angle
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) GetCenter() *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter {
	return s.Center
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) GetSize() *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize {
	return s.Size
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) SetAngle(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi {
	s.Angle = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) SetCenter(v *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi {
	s.Center = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) SetSize(v *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi {
	s.Size = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) Validate() error {
	return dara.Validate(s)
}

type RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter struct {
	// example:
	//
	// 1593
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 1360
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) GetX() *float32 {
	return s.X
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) GetY() *float32 {
	return s.Y
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) SetX(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter {
	s.X = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) SetY(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter {
	s.Y = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) Validate() error {
	return dara.Validate(s)
}

type RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize struct {
	// example:
	//
	// 81.999984741210938
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
	// example:
	//
	// 887.9998779296875
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) GetH() *float32 {
	return s.H
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) GetW() *float32 {
	return s.W
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) SetH(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize {
	s.H = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) SetW(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize {
	s.W = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) Validate() error {
	return dara.Validate(s)
}
