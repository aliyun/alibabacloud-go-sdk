// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaItemDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightOtaItemDetailResponseBody
	GetCode() *string
	SetMessage(v string) *FlightOtaItemDetailResponseBody
	GetMessage() *string
	SetModule(v *FlightOtaItemDetailResponseBodyModule) *FlightOtaItemDetailResponseBody
	GetModule() *FlightOtaItemDetailResponseBodyModule
	SetRequestId(v string) *FlightOtaItemDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightOtaItemDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightOtaItemDetailResponseBody
	GetTraceId() *string
}

type FlightOtaItemDetailResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightOtaItemDetailResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 92359A00-85D8-16C4-AED0-249618DEEC17
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210bc60a16916374201706365d2a16
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightOtaItemDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBody) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightOtaItemDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightOtaItemDetailResponseBody) GetModule() *FlightOtaItemDetailResponseBodyModule {
	return s.Module
}

func (s *FlightOtaItemDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightOtaItemDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightOtaItemDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightOtaItemDetailResponseBody) SetCode(v string) *FlightOtaItemDetailResponseBody {
	s.Code = &v
	return s
}

func (s *FlightOtaItemDetailResponseBody) SetMessage(v string) *FlightOtaItemDetailResponseBody {
	s.Message = &v
	return s
}

func (s *FlightOtaItemDetailResponseBody) SetModule(v *FlightOtaItemDetailResponseBodyModule) *FlightOtaItemDetailResponseBody {
	s.Module = v
	return s
}

func (s *FlightOtaItemDetailResponseBody) SetRequestId(v string) *FlightOtaItemDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightOtaItemDetailResponseBody) SetSuccess(v bool) *FlightOtaItemDetailResponseBody {
	s.Success = &v
	return s
}

func (s *FlightOtaItemDetailResponseBody) SetTraceId(v string) *FlightOtaItemDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightOtaItemDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModule struct {
	BaggageRule []*FlightOtaItemDetailResponseBodyModuleBaggageRule `json:"baggage_rule,omitempty" xml:"baggage_rule,omitempty" type:"Repeated"`
	ChangeRule  []*FlightOtaItemDetailResponseBodyModuleChangeRule  `json:"change_rule,omitempty" xml:"change_rule,omitempty" type:"Repeated"`
	RefundRule  []*FlightOtaItemDetailResponseBodyModuleRefundRule  `json:"refund_rule,omitempty" xml:"refund_rule,omitempty" type:"Repeated"`
	// example:
	//
	// 1830
	SellPrice     *int32   `json:"sell_price,omitempty" xml:"sell_price,omitempty"`
	SellPriceList []*int32 `json:"sell_price_list,omitempty" xml:"sell_price_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModule) GetBaggageRule() []*FlightOtaItemDetailResponseBodyModuleBaggageRule {
	return s.BaggageRule
}

func (s *FlightOtaItemDetailResponseBodyModule) GetChangeRule() []*FlightOtaItemDetailResponseBodyModuleChangeRule {
	return s.ChangeRule
}

func (s *FlightOtaItemDetailResponseBodyModule) GetRefundRule() []*FlightOtaItemDetailResponseBodyModuleRefundRule {
	return s.RefundRule
}

func (s *FlightOtaItemDetailResponseBodyModule) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *FlightOtaItemDetailResponseBodyModule) GetSellPriceList() []*int32 {
	return s.SellPriceList
}

func (s *FlightOtaItemDetailResponseBodyModule) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightOtaItemDetailResponseBodyModule) SetBaggageRule(v []*FlightOtaItemDetailResponseBodyModuleBaggageRule) *FlightOtaItemDetailResponseBodyModule {
	s.BaggageRule = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModule) SetChangeRule(v []*FlightOtaItemDetailResponseBodyModuleChangeRule) *FlightOtaItemDetailResponseBodyModule {
	s.ChangeRule = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModule) SetRefundRule(v []*FlightOtaItemDetailResponseBodyModuleRefundRule) *FlightOtaItemDetailResponseBodyModule {
	s.RefundRule = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModule) SetSellPrice(v int32) *FlightOtaItemDetailResponseBodyModule {
	s.SellPrice = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModule) SetSellPriceList(v []*int32) *FlightOtaItemDetailResponseBodyModule {
	s.SellPriceList = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModule) SetTripType(v int32) *FlightOtaItemDetailResponseBodyModule {
	s.TripType = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleBaggageRule struct {
	BaggageSubItems []*FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems `json:"baggage_sub_items,omitempty" xml:"baggage_sub_items,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// tableHead
	TableHead *string                                               `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Tips      *FlightOtaItemDetailResponseBodyModuleBaggageRuleTips `json:"tips,omitempty" xml:"tips,omitempty" type:"Struct"`
	Title     *string                                               `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleBaggageRule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleBaggageRule) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRule) GetBaggageSubItems() []*FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems {
	return s.BaggageSubItems
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRule) GetIndex() *int32 {
	return s.Index
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRule) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRule) GetTips() *FlightOtaItemDetailResponseBodyModuleBaggageRuleTips {
	return s.Tips
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRule) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRule) GetType() *int32 {
	return s.Type
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRule) SetBaggageSubItems(v []*FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) *FlightOtaItemDetailResponseBodyModuleBaggageRule {
	s.BaggageSubItems = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRule) SetIndex(v int32) *FlightOtaItemDetailResponseBodyModuleBaggageRule {
	s.Index = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRule) SetTableHead(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRule {
	s.TableHead = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRule) SetTips(v *FlightOtaItemDetailResponseBodyModuleBaggageRuleTips) *FlightOtaItemDetailResponseBodyModuleBaggageRule {
	s.Tips = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRule) SetTitle(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRule {
	s.Title = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRule) SetType(v int32) *FlightOtaItemDetailResponseBodyModuleBaggageRule {
	s.Type = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems struct {
	BaggageSubContentVisualizes []*FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes `json:"baggage_sub_content_visualizes,omitempty" xml:"baggage_sub_content_visualizes,omitempty" type:"Repeated"`
	ExtraContentVisualizes      []interface{}                                                                                 `json:"extra_content_visualizes,omitempty" xml:"extra_content_visualizes,omitempty" type:"Repeated"`
	// example:
	//
	// false
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// example:
	//
	// ADT
	Ptc   *string `json:"ptc,omitempty" xml:"ptc,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) GetBaggageSubContentVisualizes() []*FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes {
	return s.BaggageSubContentVisualizes
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) GetExtraContentVisualizes() []interface{} {
	return s.ExtraContentVisualizes
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) SetBaggageSubContentVisualizes(v []*FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems {
	s.BaggageSubContentVisualizes = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) SetExtraContentVisualizes(v []interface{}) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems {
	s.ExtraContentVisualizes = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) SetIsStruct(v bool) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) SetPtc(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) SetTitle(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems {
	s.Title = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes struct {
	// baggage_desc
	BaggageDesc []*string `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	BaggageSubContentType *int32                                                                                                 `json:"baggage_sub_content_type,omitempty" xml:"baggage_sub_content_type,omitempty"`
	Description           *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription `json:"description,omitempty" xml:"description,omitempty" type:"Struct"`
	ImageDO               *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO     `json:"image_d_o,omitempty" xml:"image_d_o,omitempty" type:"Struct"`
	// example:
	//
	// false
	IsHighlight *bool   `json:"is_highlight,omitempty" xml:"is_highlight,omitempty"`
	SubTitle    *string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) GetBaggageDesc() []*string {
	return s.BaggageDesc
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) GetBaggageSubContentType() *int32 {
	return s.BaggageSubContentType
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) GetDescription() *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription {
	return s.Description
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) GetImageDO() *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO {
	return s.ImageDO
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) GetIsHighlight() *bool {
	return s.IsHighlight
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) GetSubTitle() *string {
	return s.SubTitle
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) SetBaggageDesc(v []*string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageDesc = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) SetBaggageSubContentType(v int32) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageSubContentType = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) SetDescription(v *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes {
	s.Description = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) SetImageDO(v *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes {
	s.ImageDO = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) SetIsHighlight(v bool) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes {
	s.IsHighlight = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) SetSubTitle(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes {
	s.SubTitle = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizes) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/imgextra/i4/O1CN01UynXG31pjsEtA3tMF_!!6000000005397-2-tps-36-36.png
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/imgextra/i1/O1CN01qe7wL21gJ0SmEXXL7_!!6000000004120-2-tps-1206-768.png
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription) GetDesc() *string {
	return s.Desc
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription) GetIcon() *string {
	return s.Icon
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription) GetImage() *string {
	return s.Image
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription) SetDesc(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Desc = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription) SetIcon(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Icon = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription) SetImage(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Image = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription) SetTitle(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Title = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesDescription) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO struct {
	// example:
	//
	// https://gw.alicdn.com/imgextra/i3/O1CN01kLt3m923XsUs6WVif_!!6000000007266-2-tps-280-300.png
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// example:
	//
	// 55
	Largest *string `json:"largest,omitempty" xml:"largest,omitempty"`
	// example:
	//
	// 40
	Middle *string `json:"middle,omitempty" xml:"middle,omitempty"`
	// example:
	//
	// 20
	Smallest *string `json:"smallest,omitempty" xml:"smallest,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO) GetImage() *string {
	return s.Image
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO) GetLargest() *string {
	return s.Largest
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO) GetMiddle() *string {
	return s.Middle
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO) GetSmallest() *string {
	return s.Smallest
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO) SetImage(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Image = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO) SetLargest(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Largest = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO) SetMiddle(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Middle = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO) SetSmallest(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Smallest = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleBaggageSubItemsBaggageSubContentVisualizesImageDO) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleBaggageRuleTips struct {
	// example:
	//
	// https://gw.alicdn.com/imgextra/i1/O1CN019zl3WZ22fNLxzx2cR_!!6000000007147-2-tps-125-45.png
	Logo     *string `json:"logo,omitempty" xml:"logo,omitempty"`
	TipsDesc *string `json:"tips_desc,omitempty" xml:"tips_desc,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/imgextra/i3/O1CN01rJxjw61f3bXNHAmlk_!!6000000003951-2-tps-1050-675.png
	TipsImage *string `json:"tips_image,omitempty" xml:"tips_image,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleBaggageRuleTips) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleBaggageRuleTips) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleTips) GetLogo() *string {
	return s.Logo
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleTips) GetTipsDesc() *string {
	return s.TipsDesc
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleTips) GetTipsImage() *string {
	return s.TipsImage
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleTips) SetLogo(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleTips {
	s.Logo = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleTips) SetTipsDesc(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleTips {
	s.TipsDesc = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleTips) SetTipsImage(v string) *FlightOtaItemDetailResponseBodyModuleBaggageRuleTips {
	s.TipsImage = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleBaggageRuleTips) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleChangeRule struct {
	ExtraContents []*FlightOtaItemDetailResponseBodyModuleChangeRuleExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// HO3925
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 0
	FreeBaggage *int32 `json:"free_baggage,omitempty" xml:"free_baggage,omitempty"`
	// example:
	//
	// 0
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// 0
	Level          *int32                                                           `json:"level,omitempty" xml:"level,omitempty"`
	RefundSubItems []*FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	// subTableHead
	SubTableHead []*string `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	// example:
	//
	// tableHead
	TableHead *string `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleChangeRule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleChangeRule) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) GetExtraContents() []*FlightOtaItemDetailResponseBodyModuleChangeRuleExtraContents {
	return s.ExtraContents
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) GetFreeBaggage() *int32 {
	return s.FreeBaggage
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) GetIndex() *int32 {
	return s.Index
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) GetLevel() *int32 {
	return s.Level
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) GetRefundSubItems() []*FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) GetType() *int32 {
	return s.Type
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) SetExtraContents(v []*FlightOtaItemDetailResponseBodyModuleChangeRuleExtraContents) *FlightOtaItemDetailResponseBodyModuleChangeRule {
	s.ExtraContents = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) SetFlightNo(v string) *FlightOtaItemDetailResponseBodyModuleChangeRule {
	s.FlightNo = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) SetFreeBaggage(v int32) *FlightOtaItemDetailResponseBodyModuleChangeRule {
	s.FreeBaggage = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) SetIndex(v int32) *FlightOtaItemDetailResponseBodyModuleChangeRule {
	s.Index = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) SetLevel(v int32) *FlightOtaItemDetailResponseBodyModuleChangeRule {
	s.Level = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) SetRefundSubItems(v []*FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems) *FlightOtaItemDetailResponseBodyModuleChangeRule {
	s.RefundSubItems = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) SetSubTableHead(v []*string) *FlightOtaItemDetailResponseBodyModuleChangeRule {
	s.SubTableHead = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) SetTableHead(v string) *FlightOtaItemDetailResponseBodyModuleChangeRule {
	s.TableHead = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) SetTitle(v string) *FlightOtaItemDetailResponseBodyModuleChangeRule {
	s.Title = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) SetType(v int32) *FlightOtaItemDetailResponseBodyModuleChangeRule {
	s.Type = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleChangeRuleExtraContents struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleChangeRuleExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleChangeRuleExtraContents) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleExtraContents) SetContent(v string) *FlightOtaItemDetailResponseBodyModuleChangeRuleExtraContents {
	s.Content = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleExtraContents) SetTitle(v string) *FlightOtaItemDetailResponseBodyModuleChangeRuleExtraContents {
	s.Title = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems struct {
	// example:
	//
	// false
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// example:
	//
	// ADT
	Ptc               *string                                                                           `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                           `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems) GetRefundSubContents() []*FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems) SetIsStruct(v bool) *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems) SetPtc(v string) *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems) SetRefundSubContents(v []*FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents) *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems) SetTitle(v string) *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleChangeRuleRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleRefundRule struct {
	ExtraContents []*FlightOtaItemDetailResponseBodyModuleRefundRuleExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// HO3925
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 0
	FreeBaggage *int32 `json:"free_baggage,omitempty" xml:"free_baggage,omitempty"`
	// example:
	//
	// 0
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// 0
	Level          *int32                                                           `json:"level,omitempty" xml:"level,omitempty"`
	RefundSubItems []*FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	// subTableHead
	SubTableHead []*string `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	// example:
	//
	// tableHead
	TableHead *string `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleRefundRule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleRefundRule) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) GetExtraContents() []*FlightOtaItemDetailResponseBodyModuleRefundRuleExtraContents {
	return s.ExtraContents
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) GetFreeBaggage() *int32 {
	return s.FreeBaggage
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) GetIndex() *int32 {
	return s.Index
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) GetLevel() *int32 {
	return s.Level
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) GetRefundSubItems() []*FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) GetType() *int32 {
	return s.Type
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) SetExtraContents(v []*FlightOtaItemDetailResponseBodyModuleRefundRuleExtraContents) *FlightOtaItemDetailResponseBodyModuleRefundRule {
	s.ExtraContents = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) SetFlightNo(v string) *FlightOtaItemDetailResponseBodyModuleRefundRule {
	s.FlightNo = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) SetFreeBaggage(v int32) *FlightOtaItemDetailResponseBodyModuleRefundRule {
	s.FreeBaggage = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) SetIndex(v int32) *FlightOtaItemDetailResponseBodyModuleRefundRule {
	s.Index = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) SetLevel(v int32) *FlightOtaItemDetailResponseBodyModuleRefundRule {
	s.Level = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) SetRefundSubItems(v []*FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems) *FlightOtaItemDetailResponseBodyModuleRefundRule {
	s.RefundSubItems = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) SetSubTableHead(v []*string) *FlightOtaItemDetailResponseBodyModuleRefundRule {
	s.SubTableHead = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) SetTableHead(v string) *FlightOtaItemDetailResponseBodyModuleRefundRule {
	s.TableHead = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) SetTitle(v string) *FlightOtaItemDetailResponseBodyModuleRefundRule {
	s.Title = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) SetType(v int32) *FlightOtaItemDetailResponseBodyModuleRefundRule {
	s.Type = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleRefundRuleExtraContents struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleRefundRuleExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleRefundRuleExtraContents) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleExtraContents) SetContent(v string) *FlightOtaItemDetailResponseBodyModuleRefundRuleExtraContents {
	s.Content = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleExtraContents) SetTitle(v string) *FlightOtaItemDetailResponseBodyModuleRefundRuleExtraContents {
	s.Title = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems struct {
	// example:
	//
	// false
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// example:
	//
	// ADT
	Ptc               *string                                                                           `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                           `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems) GetRefundSubContents() []*FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems) SetIsStruct(v bool) *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems) SetPtc(v string) *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems) SetRefundSubContents(v []*FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents) *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems) SetTitle(v string) *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightOtaItemDetailResponseBodyModuleRefundRuleRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}
