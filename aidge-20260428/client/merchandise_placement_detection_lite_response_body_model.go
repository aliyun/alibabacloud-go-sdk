// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandisePlacementDetectionLiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MerchandisePlacementDetectionLiteResponseBody
	GetCode() *string
	SetData(v *MerchandisePlacementDetectionLiteResponseBodyData) *MerchandisePlacementDetectionLiteResponseBody
	GetData() *MerchandisePlacementDetectionLiteResponseBodyData
	SetMessage(v string) *MerchandisePlacementDetectionLiteResponseBody
	GetMessage() *string
	SetRequestId(v string) *MerchandisePlacementDetectionLiteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MerchandisePlacementDetectionLiteResponseBody
	GetSuccess() *bool
}

type MerchandisePlacementDetectionLiteResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *MerchandisePlacementDetectionLiteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message or failure description.
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

func (s MerchandisePlacementDetectionLiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionLiteResponseBody) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionLiteResponseBody) GetCode() *string {
	return s.Code
}

func (s *MerchandisePlacementDetectionLiteResponseBody) GetData() *MerchandisePlacementDetectionLiteResponseBodyData {
	return s.Data
}

func (s *MerchandisePlacementDetectionLiteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MerchandisePlacementDetectionLiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MerchandisePlacementDetectionLiteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MerchandisePlacementDetectionLiteResponseBody) SetCode(v string) *MerchandisePlacementDetectionLiteResponseBody {
	s.Code = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBody) SetData(v *MerchandisePlacementDetectionLiteResponseBodyData) *MerchandisePlacementDetectionLiteResponseBody {
	s.Data = v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBody) SetMessage(v string) *MerchandisePlacementDetectionLiteResponseBody {
	s.Message = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBody) SetRequestId(v string) *MerchandisePlacementDetectionLiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBody) SetSuccess(v bool) *MerchandisePlacementDetectionLiteResponseBody {
	s.Success = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MerchandisePlacementDetectionLiteResponseBodyData struct {
	// The result.
	Result *MerchandisePlacementDetectionLiteResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The usage map. The value is {"ProcessingCount":1} for a successful call, and an empty object for a failed call.
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s MerchandisePlacementDetectionLiteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionLiteResponseBodyData) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionLiteResponseBodyData) GetResult() *MerchandisePlacementDetectionLiteResponseBodyDataResult {
	return s.Result
}

func (s *MerchandisePlacementDetectionLiteResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *MerchandisePlacementDetectionLiteResponseBodyData) SetResult(v *MerchandisePlacementDetectionLiteResponseBodyDataResult) *MerchandisePlacementDetectionLiteResponseBodyData {
	s.Result = v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyData) SetUsageMap(v map[string]*int64) *MerchandisePlacementDetectionLiteResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MerchandisePlacementDetectionLiteResponseBodyDataResult struct {
	// The supplementary element grounding results.
	Grounding *MerchandisePlacementDetectionLiteResponseBodyDataResultGrounding `json:"Grounding,omitempty" xml:"Grounding,omitempty" type:"Struct"`
	// The product items.
	Products []*MerchandisePlacementDetectionLiteResponseBodyDataResultProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
	// The deduplicated list of SkuIds.
	SkuIds []*string `json:"SkuIds,omitempty" xml:"SkuIds,omitempty" type:"Repeated"`
	// The result type, which is fixed to sku.
	//
	// example:
	//
	// sku
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MerchandisePlacementDetectionLiteResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionLiteResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResult) GetGrounding() *MerchandisePlacementDetectionLiteResponseBodyDataResultGrounding {
	return s.Grounding
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResult) GetProducts() []*MerchandisePlacementDetectionLiteResponseBodyDataResultProducts {
	return s.Products
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResult) GetSkuIds() []*string {
	return s.SkuIds
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResult) GetType() *string {
	return s.Type
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResult) SetGrounding(v *MerchandisePlacementDetectionLiteResponseBodyDataResultGrounding) *MerchandisePlacementDetectionLiteResponseBodyDataResult {
	s.Grounding = v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResult) SetProducts(v []*MerchandisePlacementDetectionLiteResponseBodyDataResultProducts) *MerchandisePlacementDetectionLiteResponseBodyDataResult {
	s.Products = v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResult) SetSkuIds(v []*string) *MerchandisePlacementDetectionLiteResponseBodyDataResult {
	s.SkuIds = v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResult) SetType(v string) *MerchandisePlacementDetectionLiteResponseBodyDataResult {
	s.Type = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResult) Validate() error {
	if s.Grounding != nil {
		if err := s.Grounding.Validate(); err != nil {
			return err
		}
	}
	if s.Products != nil {
		for _, item := range s.Products {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MerchandisePlacementDetectionLiteResponseBodyDataResultGrounding struct {
	// The coverage status. Valid values:
	//
	// - complete: The relevant visible targets are complete.
	//
	// - partial: Only some targets are valid or recognizable.
	//
	// - Empty string: No valid supplementary grounding results are obtained.
	//
	// example:
	//
	// complete
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The target array, which contains up to 12 items. This array can be empty when no relevant targets are found.
	Regions []*MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s MerchandisePlacementDetectionLiteResponseBodyDataResultGrounding) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionLiteResponseBodyDataResultGrounding) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultGrounding) GetCoverage() *string {
	return s.Coverage
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultGrounding) GetRegions() []*MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions {
	return s.Regions
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultGrounding) SetCoverage(v string) *MerchandisePlacementDetectionLiteResponseBodyDataResultGrounding {
	s.Coverage = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultGrounding) SetRegions(v []*MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions) *MerchandisePlacementDetectionLiteResponseBodyDataResultGrounding {
	s.Regions = v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultGrounding) Validate() error {
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

type MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions struct {
	// The 2D bounding box in the format of [top-left x, top-left y, bottom-right x, bottom-right y]. The values range from 0 to 1000. The left coordinate is less than the right coordinate, and the top coordinate is less than the bottom coordinate.
	Bbox2d []*int64 `json:"Bbox_2d,omitempty" xml:"Bbox_2d,omitempty" type:"Repeated"`
	// The target category.
	//
	// example:
	//
	// price_tag
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The confirmable short name or text of up to 24 characters. This field is an empty string when the text is unclear.
	//
	// example:
	//
	// 5.50
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions) GetBbox2d() []*int64 {
	return s.Bbox2d
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions) GetLabel() *string {
	return s.Label
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions) GetText() *string {
	return s.Text
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions) SetBbox2d(v []*int64) *MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions {
	s.Bbox2d = v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions) SetLabel(v string) *MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions {
	s.Label = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions) SetText(v string) *MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions {
	s.Text = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultGroundingRegions) Validate() error {
	return dara.Validate(s)
}

type MerchandisePlacementDetectionLiteResponseBodyDataResultProducts struct {
	// The bounding boxes corresponding to the SKU.
	Boxes []*MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	// The recognition code, which is obtained by removing the trailing _ce from the recognition name.
	//
	// example:
	//
	// sku_test
	SkuId *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
}

func (s MerchandisePlacementDetectionLiteResponseBodyDataResultProducts) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionLiteResponseBodyDataResultProducts) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProducts) GetBoxes() []*MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes {
	return s.Boxes
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProducts) GetSkuId() *string {
	return s.SkuId
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProducts) SetBoxes(v []*MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes) *MerchandisePlacementDetectionLiteResponseBodyDataResultProducts {
	s.Boxes = v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProducts) SetSkuId(v string) *MerchandisePlacementDetectionLiteResponseBodyDataResultProducts {
	s.SkuId = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProducts) Validate() error {
	if s.Boxes != nil {
		for _, item := range s.Boxes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes struct {
	// The bottom coordinate.
	//
	// example:
	//
	// 200
	Bottom *int64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	// The left coordinate.
	//
	// example:
	//
	// 100
	Left *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	// The right coordinate.
	//
	// example:
	//
	// (\\"874\\")
	Right *int64 `json:"Right,omitempty" xml:"Right,omitempty"`
	// The bottom coordinate.
	//
	// example:
	//
	// 100
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes) GetBottom() *int64 {
	return s.Bottom
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes) GetLeft() *int64 {
	return s.Left
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes) GetRight() *int64 {
	return s.Right
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes) GetTop() *int64 {
	return s.Top
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes) SetBottom(v int64) *MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes {
	s.Bottom = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes) SetLeft(v int64) *MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes {
	s.Left = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes) SetRight(v int64) *MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes {
	s.Right = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes) SetTop(v int64) *MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes {
	s.Top = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponseBodyDataResultProductsBoxes) Validate() error {
	return dara.Validate(s)
}
