// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResult4QueryInstancePrice4ModifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetResult4QueryInstancePrice4ModifyResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetResult4QueryInstancePrice4ModifyResponseBody
	GetCode() *string
	SetData(v *GetResult4QueryInstancePrice4ModifyResponseBodyData) *GetResult4QueryInstancePrice4ModifyResponseBody
	GetData() *GetResult4QueryInstancePrice4ModifyResponseBodyData
	SetMessage(v string) *GetResult4QueryInstancePrice4ModifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetResult4QueryInstancePrice4ModifyResponseBody
	GetRequestId() *string
}

type GetResult4QueryInstancePrice4ModifyResponseBody struct {
	// example:
	//
	// {
	//
	//     "PolicyType": "",
	//
	//     "AuthPrincipalOwnerId": "",
	//
	//     "EncodedDiagnosticMessage": "",
	//
	//     "AuthPrincipalType": "",
	//
	//     "AuthPrincipalDisplayName": "",
	//
	//     "NoPermissionType": "",
	//
	//     "AuthAction": ""
	//
	//   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetResult4QueryInstancePrice4ModifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResult4QueryInstancePrice4ModifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResult4QueryInstancePrice4ModifyResponseBody) GoString() string {
	return s.String()
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) GetData() *GetResult4QueryInstancePrice4ModifyResponseBodyData {
	return s.Data
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) SetAccessDeniedDetail(v string) *GetResult4QueryInstancePrice4ModifyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) SetCode(v string) *GetResult4QueryInstancePrice4ModifyResponseBody {
	s.Code = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) SetData(v *GetResult4QueryInstancePrice4ModifyResponseBodyData) *GetResult4QueryInstancePrice4ModifyResponseBody {
	s.Data = v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) SetMessage(v string) *GetResult4QueryInstancePrice4ModifyResponseBody {
	s.Message = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) SetRequestId(v string) *GetResult4QueryInstancePrice4ModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetResult4QueryInstancePrice4ModifyResponseBodyData struct {
	PriceList []*GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList `json:"PriceList,omitempty" xml:"PriceList,omitempty" type:"Repeated"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// d9a3e99b-6954-4a16-ad51-954db4a528b7
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetResult4QueryInstancePrice4ModifyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetResult4QueryInstancePrice4ModifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyData) GetPriceList() []*GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	return s.PriceList
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyData) SetPriceList(v []*GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) *GetResult4QueryInstancePrice4ModifyResponseBodyData {
	s.PriceList = v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyData) SetStatus(v string) *GetResult4QueryInstancePrice4ModifyResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyData) SetTaskId(v string) *GetResult4QueryInstancePrice4ModifyResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList struct {
	// example:
	//
	// 2
	DiscountAmount *float64 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	Error          *string  `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// vpc
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// example:
	//
	// 3
	OriginalAmount *float64 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	PriceUnit      *string  `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	PromotionName  *string  `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// 1
	TradeAmount *float64 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) String() string {
	return dara.Prettify(s)
}

func (s GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) GoString() string {
	return s.String()
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) GetDiscountAmount() *float64 {
	return s.DiscountAmount
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) GetError() *string {
	return s.Error
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) GetNodeType() *string {
	return s.NodeType
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) GetOriginalAmount() *float64 {
	return s.OriginalAmount
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) GetPromotionName() *string {
	return s.PromotionName
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) GetTradeAmount() *float64 {
	return s.TradeAmount
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetDiscountAmount(v float64) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.DiscountAmount = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetError(v string) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.Error = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetNodeType(v string) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.NodeType = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetOriginalAmount(v float64) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.OriginalAmount = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetPriceUnit(v string) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.PriceUnit = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetPromotionName(v string) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.PromotionName = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetTradeAmount(v float64) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.TradeAmount = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) Validate() error {
	return dara.Validate(s)
}
