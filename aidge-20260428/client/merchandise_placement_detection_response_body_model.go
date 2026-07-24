// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandisePlacementDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MerchandisePlacementDetectionResponseBody
	GetCode() *string
	SetData(v *MerchandisePlacementDetectionResponseBodyData) *MerchandisePlacementDetectionResponseBody
	GetData() *MerchandisePlacementDetectionResponseBodyData
	SetMessage(v string) *MerchandisePlacementDetectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *MerchandisePlacementDetectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MerchandisePlacementDetectionResponseBody
	GetSuccess() *bool
}

type MerchandisePlacementDetectionResponseBody struct {
	// The error code. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The display detection result.
	Data *MerchandisePlacementDetectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MerchandisePlacementDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *MerchandisePlacementDetectionResponseBody) GetData() *MerchandisePlacementDetectionResponseBodyData {
	return s.Data
}

func (s *MerchandisePlacementDetectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MerchandisePlacementDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MerchandisePlacementDetectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MerchandisePlacementDetectionResponseBody) SetCode(v string) *MerchandisePlacementDetectionResponseBody {
	s.Code = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBody) SetData(v *MerchandisePlacementDetectionResponseBodyData) *MerchandisePlacementDetectionResponseBody {
	s.Data = v
	return s
}

func (s *MerchandisePlacementDetectionResponseBody) SetMessage(v string) *MerchandisePlacementDetectionResponseBody {
	s.Message = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBody) SetRequestId(v string) *MerchandisePlacementDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBody) SetSuccess(v bool) *MerchandisePlacementDetectionResponseBody {
	s.Success = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MerchandisePlacementDetectionResponseBodyData struct {
	// The number of valid detection boxes.
	//
	// example:
	//
	// 100
	BoxCount *int32 `json:"BoxCount,omitempty" xml:"BoxCount,omitempty"`
	// The list of retrieval details for each detection box.
	Data []*MerchandisePlacementDetectionResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The usage information. The key is the usage metric name, and the value is the count.
	//
	// example:
	//
	// {"ProcessingCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s MerchandisePlacementDetectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionResponseBodyData) GetBoxCount() *int32 {
	return s.BoxCount
}

func (s *MerchandisePlacementDetectionResponseBodyData) GetData() []*MerchandisePlacementDetectionResponseBodyDataData {
	return s.Data
}

func (s *MerchandisePlacementDetectionResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *MerchandisePlacementDetectionResponseBodyData) SetBoxCount(v int32) *MerchandisePlacementDetectionResponseBodyData {
	s.BoxCount = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyData) SetData(v []*MerchandisePlacementDetectionResponseBodyDataData) *MerchandisePlacementDetectionResponseBodyData {
	s.Data = v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyData) SetUsageMap(v map[string]*int64) *MerchandisePlacementDetectionResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyData) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MerchandisePlacementDetectionResponseBodyDataData struct {
	// The position coordinates of the detection box in the format [x1,y1,x2,y2].
	Bbox []*float32 `json:"Bbox,omitempty" xml:"Bbox,omitempty" type:"Repeated"`
	// The failure reason for the detection box. The value is null if the detection is successful.
	//
	// example:
	//
	// "embedding failed"
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The index of the detection box.
	//
	// example:
	//
	// 1
	Idx *int32 `json:"Idx,omitempty" xml:"Idx,omitempty"`
	// The top-1 recalled product for the detection box.
	Top1 *MerchandisePlacementDetectionResponseBodyDataDataTop1 `json:"Top1,omitempty" xml:"Top1,omitempty" type:"Struct"`
	// The list of top-K recalled products for the detection box.
	Topk []*MerchandisePlacementDetectionResponseBodyDataDataTopk `json:"Topk,omitempty" xml:"Topk,omitempty" type:"Repeated"`
}

func (s MerchandisePlacementDetectionResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionResponseBodyDataData) GetBbox() []*float32 {
	return s.Bbox
}

func (s *MerchandisePlacementDetectionResponseBodyDataData) GetError() *string {
	return s.Error
}

func (s *MerchandisePlacementDetectionResponseBodyDataData) GetIdx() *int32 {
	return s.Idx
}

func (s *MerchandisePlacementDetectionResponseBodyDataData) GetTop1() *MerchandisePlacementDetectionResponseBodyDataDataTop1 {
	return s.Top1
}

func (s *MerchandisePlacementDetectionResponseBodyDataData) GetTopk() []*MerchandisePlacementDetectionResponseBodyDataDataTopk {
	return s.Topk
}

func (s *MerchandisePlacementDetectionResponseBodyDataData) SetBbox(v []*float32) *MerchandisePlacementDetectionResponseBodyDataData {
	s.Bbox = v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyDataData) SetError(v string) *MerchandisePlacementDetectionResponseBodyDataData {
	s.Error = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyDataData) SetIdx(v int32) *MerchandisePlacementDetectionResponseBodyDataData {
	s.Idx = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyDataData) SetTop1(v *MerchandisePlacementDetectionResponseBodyDataDataTop1) *MerchandisePlacementDetectionResponseBodyDataData {
	s.Top1 = v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyDataData) SetTopk(v []*MerchandisePlacementDetectionResponseBodyDataDataTopk) *MerchandisePlacementDetectionResponseBodyDataData {
	s.Topk = v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyDataData) Validate() error {
	if s.Top1 != nil {
		if err := s.Top1.Validate(); err != nil {
			return err
		}
	}
	if s.Topk != nil {
		for _, item := range s.Topk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MerchandisePlacementDetectionResponseBodyDataDataTop1 struct {
	// The similarity score, ranging from 0 to 1.
	//
	// example:
	//
	// 0.53
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The ID of the recalled product.
	//
	// example:
	//
	// 123
	SkuId *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	// The name of the recalled product.
	//
	// example:
	//
	// Bright Milk 500ml.
	SkuName *string `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
}

func (s MerchandisePlacementDetectionResponseBodyDataDataTop1) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionResponseBodyDataDataTop1) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTop1) GetScore() *float32 {
	return s.Score
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTop1) GetSkuId() *string {
	return s.SkuId
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTop1) GetSkuName() *string {
	return s.SkuName
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTop1) SetScore(v float32) *MerchandisePlacementDetectionResponseBodyDataDataTop1 {
	s.Score = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTop1) SetSkuId(v string) *MerchandisePlacementDetectionResponseBodyDataDataTop1 {
	s.SkuId = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTop1) SetSkuName(v string) *MerchandisePlacementDetectionResponseBodyDataDataTop1 {
	s.SkuName = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTop1) Validate() error {
	return dara.Validate(s)
}

type MerchandisePlacementDetectionResponseBodyDataDataTopk struct {
	// The recall rank.
	//
	// example:
	//
	// 1
	Rank *int32 `json:"Rank,omitempty" xml:"Rank,omitempty"`
	// The similarity score, ranging from 0 to 1.
	//
	// example:
	//
	// 0.82
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The ID of the recalled product.
	//
	// example:
	//
	// 123
	SkuId *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	// The name of the recalled product.
	//
	// example:
	//
	// Bright Milk 500ml.
	SkuName *string `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
}

func (s MerchandisePlacementDetectionResponseBodyDataDataTopk) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionResponseBodyDataDataTopk) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTopk) GetRank() *int32 {
	return s.Rank
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTopk) GetScore() *float32 {
	return s.Score
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTopk) GetSkuId() *string {
	return s.SkuId
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTopk) GetSkuName() *string {
	return s.SkuName
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTopk) SetRank(v int32) *MerchandisePlacementDetectionResponseBodyDataDataTopk {
	s.Rank = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTopk) SetScore(v float32) *MerchandisePlacementDetectionResponseBodyDataDataTopk {
	s.Score = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTopk) SetSkuId(v string) *MerchandisePlacementDetectionResponseBodyDataDataTopk {
	s.SkuId = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTopk) SetSkuName(v string) *MerchandisePlacementDetectionResponseBodyDataDataTopk {
	s.SkuName = &v
	return s
}

func (s *MerchandisePlacementDetectionResponseBodyDataDataTopk) Validate() error {
	return dara.Validate(s)
}
