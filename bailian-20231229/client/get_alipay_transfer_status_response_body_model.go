// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlipayTransferStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAlipayTransferStatusResponseBodyData) *GetAlipayTransferStatusResponseBody
	GetData() *GetAlipayTransferStatusResponseBodyData
	SetRequestId(v string) *GetAlipayTransferStatusResponseBody
	GetRequestId() *string
}

type GetAlipayTransferStatusResponseBody struct {
	Data *GetAlipayTransferStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-11402910xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAlipayTransferStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlipayTransferStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlipayTransferStatusResponseBody) GetData() *GetAlipayTransferStatusResponseBodyData {
	return s.Data
}

func (s *GetAlipayTransferStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlipayTransferStatusResponseBody) SetData(v *GetAlipayTransferStatusResponseBodyData) *GetAlipayTransferStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetAlipayTransferStatusResponseBody) SetRequestId(v string) *GetAlipayTransferStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAlipayTransferStatusResponseBodyData struct {
	// example:
	//
	// 1348393307144609
	AccountId         *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AlipayOrderDetail *string `json:"alipayOrderDetail,omitempty" xml:"alipayOrderDetail,omitempty"`
	AlipayOrderId     *string `json:"alipayOrderId,omitempty" xml:"alipayOrderId,omitempty"`
	// example:
	//
	// 0
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// 1007576424487905
	MainAccountId *string `json:"mainAccountId,omitempty" xml:"mainAccountId,omitempty"`
	Modifier      *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	QrURL         *string `json:"qrURL,omitempty" xml:"qrURL,omitempty"`
	Scope         *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// example:
	//
	// 1
	Status         *int64  `json:"status,omitempty" xml:"status,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
	TransAmount    *string `json:"transAmount,omitempty" xml:"transAmount,omitempty"`
	WalletItemCode *string `json:"walletItemCode,omitempty" xml:"walletItemCode,omitempty"`
}

func (s GetAlipayTransferStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAlipayTransferStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAlipayTransferStatusResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAlipayTransferStatusResponseBodyData) GetAlipayOrderDetail() *string {
	return s.AlipayOrderDetail
}

func (s *GetAlipayTransferStatusResponseBodyData) GetAlipayOrderId() *string {
	return s.AlipayOrderId
}

func (s *GetAlipayTransferStatusResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetAlipayTransferStatusResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *GetAlipayTransferStatusResponseBodyData) GetMainAccountId() *string {
	return s.MainAccountId
}

func (s *GetAlipayTransferStatusResponseBodyData) GetModifier() *string {
	return s.Modifier
}

func (s *GetAlipayTransferStatusResponseBodyData) GetQrURL() *string {
	return s.QrURL
}

func (s *GetAlipayTransferStatusResponseBodyData) GetScope() *string {
	return s.Scope
}

func (s *GetAlipayTransferStatusResponseBodyData) GetStatus() *int64 {
	return s.Status
}

func (s *GetAlipayTransferStatusResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetAlipayTransferStatusResponseBodyData) GetTransAmount() *string {
	return s.TransAmount
}

func (s *GetAlipayTransferStatusResponseBodyData) GetWalletItemCode() *string {
	return s.WalletItemCode
}

func (s *GetAlipayTransferStatusResponseBodyData) SetAccountId(v string) *GetAlipayTransferStatusResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBodyData) SetAlipayOrderDetail(v string) *GetAlipayTransferStatusResponseBodyData {
	s.AlipayOrderDetail = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBodyData) SetAlipayOrderId(v string) *GetAlipayTransferStatusResponseBodyData {
	s.AlipayOrderId = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBodyData) SetCode(v string) *GetAlipayTransferStatusResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBodyData) SetCreator(v string) *GetAlipayTransferStatusResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBodyData) SetMainAccountId(v string) *GetAlipayTransferStatusResponseBodyData {
	s.MainAccountId = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBodyData) SetModifier(v string) *GetAlipayTransferStatusResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBodyData) SetQrURL(v string) *GetAlipayTransferStatusResponseBodyData {
	s.QrURL = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBodyData) SetScope(v string) *GetAlipayTransferStatusResponseBodyData {
	s.Scope = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBodyData) SetStatus(v int64) *GetAlipayTransferStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBodyData) SetTitle(v string) *GetAlipayTransferStatusResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBodyData) SetTransAmount(v string) *GetAlipayTransferStatusResponseBodyData {
	s.TransAmount = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBodyData) SetWalletItemCode(v string) *GetAlipayTransferStatusResponseBodyData {
	s.WalletItemCode = &v
	return s
}

func (s *GetAlipayTransferStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
