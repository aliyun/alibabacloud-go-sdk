// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandisePriceTagDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MerchandisePriceTagDetectionResponseBody
	GetCode() *string
	SetData(v *MerchandisePriceTagDetectionResponseBodyData) *MerchandisePriceTagDetectionResponseBody
	GetData() *MerchandisePriceTagDetectionResponseBodyData
	SetMessage(v string) *MerchandisePriceTagDetectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *MerchandisePriceTagDetectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MerchandisePriceTagDetectionResponseBody
	GetSuccess() *bool
}

type MerchandisePriceTagDetectionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The scenario recognition result.
	Data *MerchandisePriceTagDetectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The return message or failure description.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 70CBEFDF-BB17-1EB3-8A21-569F3124738F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MerchandisePriceTagDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePriceTagDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *MerchandisePriceTagDetectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *MerchandisePriceTagDetectionResponseBody) GetData() *MerchandisePriceTagDetectionResponseBodyData {
	return s.Data
}

func (s *MerchandisePriceTagDetectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MerchandisePriceTagDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MerchandisePriceTagDetectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MerchandisePriceTagDetectionResponseBody) SetCode(v string) *MerchandisePriceTagDetectionResponseBody {
	s.Code = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBody) SetData(v *MerchandisePriceTagDetectionResponseBodyData) *MerchandisePriceTagDetectionResponseBody {
	s.Data = v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBody) SetMessage(v string) *MerchandisePriceTagDetectionResponseBody {
	s.Message = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBody) SetRequestId(v string) *MerchandisePriceTagDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBody) SetSuccess(v bool) *MerchandisePriceTagDetectionResponseBody {
	s.Success = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MerchandisePriceTagDetectionResponseBodyData struct {
	// The detection and recognition result details.
	Result *MerchandisePriceTagDetectionResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The usage information.
	//
	// example:
	//
	// {"ProcessingCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s MerchandisePriceTagDetectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePriceTagDetectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *MerchandisePriceTagDetectionResponseBodyData) GetResult() *MerchandisePriceTagDetectionResponseBodyDataResult {
	return s.Result
}

func (s *MerchandisePriceTagDetectionResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *MerchandisePriceTagDetectionResponseBodyData) SetResult(v *MerchandisePriceTagDetectionResponseBodyDataResult) *MerchandisePriceTagDetectionResponseBodyData {
	s.Result = v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyData) SetUsageMap(v map[string]*int64) *MerchandisePriceTagDetectionResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MerchandisePriceTagDetectionResponseBodyDataResult struct {
	// The localization information.
	Grounding *MerchandisePriceTagDetectionResponseBodyDataResultGrounding `json:"Grounding,omitempty" xml:"Grounding,omitempty" type:"Struct"`
	// The price tag list. This can be an empty array if no price tags are recognized under normal conditions.
	PriceTags []*MerchandisePriceTagDetectionResponseBodyDataResultPriceTags `json:"PriceTags,omitempty" xml:"PriceTags,omitempty" type:"Repeated"`
	// The value is fixed to priceTag.
	//
	// example:
	//
	// priceTag
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MerchandisePriceTagDetectionResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePriceTagDetectionResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResult) GetGrounding() *MerchandisePriceTagDetectionResponseBodyDataResultGrounding {
	return s.Grounding
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResult) GetPriceTags() []*MerchandisePriceTagDetectionResponseBodyDataResultPriceTags {
	return s.PriceTags
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResult) GetType() *string {
	return s.Type
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResult) SetGrounding(v *MerchandisePriceTagDetectionResponseBodyDataResultGrounding) *MerchandisePriceTagDetectionResponseBodyDataResult {
	s.Grounding = v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResult) SetPriceTags(v []*MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) *MerchandisePriceTagDetectionResponseBodyDataResult {
	s.PriceTags = v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResult) SetType(v string) *MerchandisePriceTagDetectionResponseBodyDataResult {
	s.Type = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResult) Validate() error {
	if s.Grounding != nil {
		if err := s.Grounding.Validate(); err != nil {
			return err
		}
	}
	if s.PriceTags != nil {
		for _, item := range s.PriceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MerchandisePriceTagDetectionResponseBodyDataResultGrounding struct {
	// The coverage.
	//
	// example:
	//
	// complete
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The region list.
	Regions []*MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s MerchandisePriceTagDetectionResponseBodyDataResultGrounding) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePriceTagDetectionResponseBodyDataResultGrounding) GoString() string {
	return s.String()
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultGrounding) GetCoverage() *string {
	return s.Coverage
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultGrounding) GetRegions() []*MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions {
	return s.Regions
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultGrounding) SetCoverage(v string) *MerchandisePriceTagDetectionResponseBodyDataResultGrounding {
	s.Coverage = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultGrounding) SetRegions(v []*MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions) *MerchandisePriceTagDetectionResponseBodyDataResultGrounding {
	s.Regions = v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultGrounding) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions struct {
	// The bounding box coordinates.
	Bbox2d []*float64 `json:"Bbox2d,omitempty" xml:"Bbox2d,omitempty" type:"Repeated"`
	// The label.
	//
	// example:
	//
	// price_tag
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The recognized text.
	//
	// example:
	//
	// 5.50
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions) GoString() string {
	return s.String()
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions) GetBbox2d() []*float64 {
	return s.Bbox2d
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions) GetLabel() *string {
	return s.Label
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions) GetText() *string {
	return s.Text
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions) SetBbox2d(v []*float64) *MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions {
	s.Bbox2d = v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions) SetLabel(v string) *MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions {
	s.Label = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions) SetText(v string) *MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions {
	s.Text = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultGroundingRegions) Validate() error {
	return dara.Validate(s)
}

type MerchandisePriceTagDetectionResponseBodyDataResultPriceTags struct {
	// The location.
	Location *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// The original price.
	//
	// example:
	//
	// 6
	OriginalPrice *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The selling price, retained as a string to avoid being parsed as a reserved word or reserved character.
	//
	// example:
	//
	// 5.50
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
	// The associated product code.
	//
	// example:
	//
	// 001
	SkuCode *string `json:"SkuCode,omitempty" xml:"SkuCode,omitempty"`
	// The associated product name.
	//
	// example:
	//
	// sku_test
	SkuName *string `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
}

func (s MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) GoString() string {
	return s.String()
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) GetLocation() *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation {
	return s.Location
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) GetOriginalPrice() *string {
	return s.OriginalPrice
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) GetPrice() *string {
	return s.Price
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) GetSkuCode() *string {
	return s.SkuCode
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) GetSkuName() *string {
	return s.SkuName
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) SetLocation(v *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation) *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags {
	s.Location = v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) SetOriginalPrice(v string) *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags {
	s.OriginalPrice = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) SetPrice(v string) *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags {
	s.Price = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) SetSkuCode(v string) *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags {
	s.SkuCode = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) SetSkuName(v string) *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags {
	s.SkuName = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTags) Validate() error {
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation struct {
	// The bottom boundary of the detection box.
	//
	// example:
	//
	// 369
	Bottom *string `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	// The left boundary of the detection box.
	//
	// example:
	//
	// 796
	Left *string `json:"Left,omitempty" xml:"Left,omitempty"`
	// The right boundary of the detection box.
	//
	// example:
	//
	// 820
	Right *string `json:"Right,omitempty" xml:"Right,omitempty"`
	// The top boundary of the detection box.
	//
	// example:
	//
	// 338
	Top *string `json:"top,omitempty" xml:"top,omitempty"`
}

func (s MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation) GoString() string {
	return s.String()
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation) GetBottom() *string {
	return s.Bottom
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation) GetLeft() *string {
	return s.Left
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation) GetRight() *string {
	return s.Right
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation) GetTop() *string {
	return s.Top
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation) SetBottom(v string) *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation {
	s.Bottom = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation) SetLeft(v string) *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation {
	s.Left = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation) SetRight(v string) *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation {
	s.Right = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation) SetTop(v string) *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation {
	s.Top = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponseBodyDataResultPriceTagsLocation) Validate() error {
	return dara.Validate(s)
}
