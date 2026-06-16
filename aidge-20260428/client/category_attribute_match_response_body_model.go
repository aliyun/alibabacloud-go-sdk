// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCategoryAttributeMatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CategoryAttributeMatchResponseBody
	GetCode() *string
	SetData(v *CategoryAttributeMatchResponseBodyData) *CategoryAttributeMatchResponseBody
	GetData() *CategoryAttributeMatchResponseBodyData
	SetMessage(v string) *CategoryAttributeMatchResponseBody
	GetMessage() *string
	SetRequestId(v string) *CategoryAttributeMatchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CategoryAttributeMatchResponseBody
	GetSuccess() *bool
}

type CategoryAttributeMatchResponseBody struct {
	// The error code. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *CategoryAttributeMatchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9927B72F-3C8F-54C0-AAE3-2B1AC94BB6EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CategoryAttributeMatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CategoryAttributeMatchResponseBody) GoString() string {
	return s.String()
}

func (s *CategoryAttributeMatchResponseBody) GetCode() *string {
	return s.Code
}

func (s *CategoryAttributeMatchResponseBody) GetData() *CategoryAttributeMatchResponseBodyData {
	return s.Data
}

func (s *CategoryAttributeMatchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CategoryAttributeMatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CategoryAttributeMatchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CategoryAttributeMatchResponseBody) SetCode(v string) *CategoryAttributeMatchResponseBody {
	s.Code = &v
	return s
}

func (s *CategoryAttributeMatchResponseBody) SetData(v *CategoryAttributeMatchResponseBodyData) *CategoryAttributeMatchResponseBody {
	s.Data = v
	return s
}

func (s *CategoryAttributeMatchResponseBody) SetMessage(v string) *CategoryAttributeMatchResponseBody {
	s.Message = &v
	return s
}

func (s *CategoryAttributeMatchResponseBody) SetRequestId(v string) *CategoryAttributeMatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CategoryAttributeMatchResponseBody) SetSuccess(v bool) *CategoryAttributeMatchResponseBody {
	s.Success = &v
	return s
}

func (s *CategoryAttributeMatchResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CategoryAttributeMatchResponseBodyData struct {
	// The list of attribute filling results.
	Attributes []*CategoryAttributeMatchResponseBodyDataAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The category ID.
	//
	// example:
	//
	// FC-F3A8A2802D10606D
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The category name.
	//
	// example:
	//
	// 猫项圈
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The full path of the category, separated by "/".
	//
	// example:
	//
	// 宠物用品/猫用品/猫挂饰、项圈、牵引带/猫项圈
	CategoryPath *string `json:"CategoryPath,omitempty" xml:"CategoryPath,omitempty"`
	// The number of attributes that are successfully filled.
	//
	// example:
	//
	// 7
	FilledCount *int32 `json:"FilledCount,omitempty" xml:"FilledCount,omitempty"`
	// Indicates whether the attribute is successfully matched. Valid values: true and false.
	//
	// example:
	//
	// True
	Matched *bool `json:"Matched,omitempty" xml:"Matched,omitempty"`
	// The total number of attributes under the category.
	//
	// example:
	//
	// 10
	TotalAttributes *int32 `json:"TotalAttributes,omitempty" xml:"TotalAttributes,omitempty"`
	// The usage fields.
	UsageMap *CategoryAttributeMatchResponseBodyDataUsageMap `json:"UsageMap,omitempty" xml:"UsageMap,omitempty" type:"Struct"`
}

func (s CategoryAttributeMatchResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CategoryAttributeMatchResponseBodyData) GoString() string {
	return s.String()
}

func (s *CategoryAttributeMatchResponseBodyData) GetAttributes() []*CategoryAttributeMatchResponseBodyDataAttributes {
	return s.Attributes
}

func (s *CategoryAttributeMatchResponseBodyData) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *CategoryAttributeMatchResponseBodyData) GetCategoryName() *string {
	return s.CategoryName
}

func (s *CategoryAttributeMatchResponseBodyData) GetCategoryPath() *string {
	return s.CategoryPath
}

func (s *CategoryAttributeMatchResponseBodyData) GetFilledCount() *int32 {
	return s.FilledCount
}

func (s *CategoryAttributeMatchResponseBodyData) GetMatched() *bool {
	return s.Matched
}

func (s *CategoryAttributeMatchResponseBodyData) GetTotalAttributes() *int32 {
	return s.TotalAttributes
}

func (s *CategoryAttributeMatchResponseBodyData) GetUsageMap() *CategoryAttributeMatchResponseBodyDataUsageMap {
	return s.UsageMap
}

func (s *CategoryAttributeMatchResponseBodyData) SetAttributes(v []*CategoryAttributeMatchResponseBodyDataAttributes) *CategoryAttributeMatchResponseBodyData {
	s.Attributes = v
	return s
}

func (s *CategoryAttributeMatchResponseBodyData) SetCategoryId(v int32) *CategoryAttributeMatchResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyData) SetCategoryName(v string) *CategoryAttributeMatchResponseBodyData {
	s.CategoryName = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyData) SetCategoryPath(v string) *CategoryAttributeMatchResponseBodyData {
	s.CategoryPath = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyData) SetFilledCount(v int32) *CategoryAttributeMatchResponseBodyData {
	s.FilledCount = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyData) SetMatched(v bool) *CategoryAttributeMatchResponseBodyData {
	s.Matched = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyData) SetTotalAttributes(v int32) *CategoryAttributeMatchResponseBodyData {
	s.TotalAttributes = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyData) SetUsageMap(v *CategoryAttributeMatchResponseBodyDataUsageMap) *CategoryAttributeMatchResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *CategoryAttributeMatchResponseBodyData) Validate() error {
	if s.Attributes != nil {
		for _, item := range s.Attributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UsageMap != nil {
		if err := s.UsageMap.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CategoryAttributeMatchResponseBodyDataAttributes struct {
	// The attribute ID.
	//
	// example:
	//
	// 682439
	AttrId *int32 `json:"AttrId,omitempty" xml:"AttrId,omitempty"`
	// The matching confidence score. Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	Confidence *int32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The input type of the attribute.
	//
	// example:
	//
	// 单选下拉
	InputType *string `json:"InputType,omitempty" xml:"InputType,omitempty"`
	// Indicates whether the attribute is successfully matched. Valid values: true and false.
	//
	// example:
	//
	// True
	Matched *bool `json:"Matched,omitempty" xml:"Matched,omitempty"`
	// The Chinese name of the attribute.
	//
	// example:
	//
	// netpila-backup-vpc-j5ekvvg5i5iquaeqbyf6b-cn-shenzhen-finance-1d
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The English name of the attribute.
	//
	// example:
	//
	// ABC Private POP
	NameEn *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	// The reason for the matching result.
	//
	// example:
	//
	// a_reason
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The list of selected attribute value texts, such as ["iOS","Android"\\].
	SelectedValues []*string `json:"SelectedValues,omitempty" xml:"SelectedValues,omitempty" type:"Repeated"`
	// The list of selected attribute value IDs, such as [30127,30128\\].
	SelectedVids []*int32 `json:"SelectedVids,omitempty" xml:"SelectedVids,omitempty" type:"Repeated"`
}

func (s CategoryAttributeMatchResponseBodyDataAttributes) String() string {
	return dara.Prettify(s)
}

func (s CategoryAttributeMatchResponseBodyDataAttributes) GoString() string {
	return s.String()
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) GetAttrId() *int32 {
	return s.AttrId
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) GetConfidence() *int32 {
	return s.Confidence
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) GetInputType() *string {
	return s.InputType
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) GetMatched() *bool {
	return s.Matched
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) GetName() *string {
	return s.Name
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) GetNameEn() *string {
	return s.NameEn
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) GetReason() *string {
	return s.Reason
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) GetSelectedValues() []*string {
	return s.SelectedValues
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) GetSelectedVids() []*int32 {
	return s.SelectedVids
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) SetAttrId(v int32) *CategoryAttributeMatchResponseBodyDataAttributes {
	s.AttrId = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) SetConfidence(v int32) *CategoryAttributeMatchResponseBodyDataAttributes {
	s.Confidence = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) SetInputType(v string) *CategoryAttributeMatchResponseBodyDataAttributes {
	s.InputType = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) SetMatched(v bool) *CategoryAttributeMatchResponseBodyDataAttributes {
	s.Matched = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) SetName(v string) *CategoryAttributeMatchResponseBodyDataAttributes {
	s.Name = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) SetNameEn(v string) *CategoryAttributeMatchResponseBodyDataAttributes {
	s.NameEn = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) SetReason(v string) *CategoryAttributeMatchResponseBodyDataAttributes {
	s.Reason = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) SetSelectedValues(v []*string) *CategoryAttributeMatchResponseBodyDataAttributes {
	s.SelectedValues = v
	return s
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) SetSelectedVids(v []*int32) *CategoryAttributeMatchResponseBodyDataAttributes {
	s.SelectedVids = v
	return s
}

func (s *CategoryAttributeMatchResponseBodyDataAttributes) Validate() error {
	return dara.Validate(s)
}

type CategoryAttributeMatchResponseBodyDataUsageMap struct {
	// The number of processing times.
	//
	// example:
	//
	// 1
	ProcessingCount *int32 `json:"ProcessingCount,omitempty" xml:"ProcessingCount,omitempty"`
}

func (s CategoryAttributeMatchResponseBodyDataUsageMap) String() string {
	return dara.Prettify(s)
}

func (s CategoryAttributeMatchResponseBodyDataUsageMap) GoString() string {
	return s.String()
}

func (s *CategoryAttributeMatchResponseBodyDataUsageMap) GetProcessingCount() *int32 {
	return s.ProcessingCount
}

func (s *CategoryAttributeMatchResponseBodyDataUsageMap) SetProcessingCount(v int32) *CategoryAttributeMatchResponseBodyDataUsageMap {
	s.ProcessingCount = &v
	return s
}

func (s *CategoryAttributeMatchResponseBodyDataUsageMap) Validate() error {
	return dara.Validate(s)
}
