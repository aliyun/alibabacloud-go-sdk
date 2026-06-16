// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCategoryMatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CategoryMatchResponseBody
	GetCode() *string
	SetData(v *CategoryMatchResponseBodyData) *CategoryMatchResponseBody
	GetData() *CategoryMatchResponseBodyData
	SetMessage(v string) *CategoryMatchResponseBody
	GetMessage() *string
	SetRequestId(v string) *CategoryMatchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CategoryMatchResponseBody
	GetSuccess() *bool
}

type CategoryMatchResponseBody struct {
	// The status code. The value "success" is returned for a successful call.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The product category matching result.
	Data *CategoryMatchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. The value "Success" is returned for a successful call.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which uniquely identifies the request.
	//
	// example:
	//
	// 2157065A-D6C8-1F3E-A4D0-B1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CategoryMatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CategoryMatchResponseBody) GoString() string {
	return s.String()
}

func (s *CategoryMatchResponseBody) GetCode() *string {
	return s.Code
}

func (s *CategoryMatchResponseBody) GetData() *CategoryMatchResponseBodyData {
	return s.Data
}

func (s *CategoryMatchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CategoryMatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CategoryMatchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CategoryMatchResponseBody) SetCode(v string) *CategoryMatchResponseBody {
	s.Code = &v
	return s
}

func (s *CategoryMatchResponseBody) SetData(v *CategoryMatchResponseBodyData) *CategoryMatchResponseBody {
	s.Data = v
	return s
}

func (s *CategoryMatchResponseBody) SetMessage(v string) *CategoryMatchResponseBody {
	s.Message = &v
	return s
}

func (s *CategoryMatchResponseBody) SetRequestId(v string) *CategoryMatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CategoryMatchResponseBody) SetSuccess(v bool) *CategoryMatchResponseBody {
	s.Success = &v
	return s
}

func (s *CategoryMatchResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CategoryMatchResponseBodyData struct {
	// The ID of the matched category.
	//
	// example:
	//
	// 1522
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The name of the matched category.
	//
	// example:
	//
	// 位置和活动跟踪器
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The full path of the category, separated by forward slashes (/).
	//
	// example:
	//
	// 宠物用品/猫用品/猫挂饰、项圈、牵引带/位置和活动跟踪器
	CategoryPath *string `json:"CategoryPath,omitempty" xml:"CategoryPath,omitempty"`
	// The matching confidence score, ranging from 0 to 100.
	//
	// example:
	//
	// 96
	Confidence *int32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Indicates whether the match is successful.
	//
	// example:
	//
	// true
	MatchSuccessful *bool `json:"MatchSuccessful,omitempty" xml:"MatchSuccessful,omitempty"`
	// The reason for the match.
	//
	// example:
	//
	// 商品核心为带 AirTag 定位功能的猫项圈，属\\"位置和活动跟踪器\\"类目，叶子节点语义精准匹配其追踪功能与猫用属性。
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The usage information.
	UsageMap map[string]*int32 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s CategoryMatchResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CategoryMatchResponseBodyData) GoString() string {
	return s.String()
}

func (s *CategoryMatchResponseBodyData) GetCategoryId() *string {
	return s.CategoryId
}

func (s *CategoryMatchResponseBodyData) GetCategoryName() *string {
	return s.CategoryName
}

func (s *CategoryMatchResponseBodyData) GetCategoryPath() *string {
	return s.CategoryPath
}

func (s *CategoryMatchResponseBodyData) GetConfidence() *int32 {
	return s.Confidence
}

func (s *CategoryMatchResponseBodyData) GetMatchSuccessful() *bool {
	return s.MatchSuccessful
}

func (s *CategoryMatchResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *CategoryMatchResponseBodyData) GetUsageMap() map[string]*int32 {
	return s.UsageMap
}

func (s *CategoryMatchResponseBodyData) SetCategoryId(v string) *CategoryMatchResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *CategoryMatchResponseBodyData) SetCategoryName(v string) *CategoryMatchResponseBodyData {
	s.CategoryName = &v
	return s
}

func (s *CategoryMatchResponseBodyData) SetCategoryPath(v string) *CategoryMatchResponseBodyData {
	s.CategoryPath = &v
	return s
}

func (s *CategoryMatchResponseBodyData) SetConfidence(v int32) *CategoryMatchResponseBodyData {
	s.Confidence = &v
	return s
}

func (s *CategoryMatchResponseBodyData) SetMatchSuccessful(v bool) *CategoryMatchResponseBodyData {
	s.MatchSuccessful = &v
	return s
}

func (s *CategoryMatchResponseBodyData) SetReason(v string) *CategoryMatchResponseBodyData {
	s.Reason = &v
	return s
}

func (s *CategoryMatchResponseBodyData) SetUsageMap(v map[string]*int32) *CategoryMatchResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *CategoryMatchResponseBodyData) Validate() error {
	return dara.Validate(s)
}
