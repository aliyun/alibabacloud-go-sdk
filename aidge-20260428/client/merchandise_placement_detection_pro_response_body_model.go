// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandisePlacementDetectionProResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MerchandisePlacementDetectionProResponseBody
	GetCode() *string
	SetData(v *MerchandisePlacementDetectionProResponseBodyData) *MerchandisePlacementDetectionProResponseBody
	GetData() *MerchandisePlacementDetectionProResponseBodyData
	SetMessage(v string) *MerchandisePlacementDetectionProResponseBody
	GetMessage() *string
	SetRequestId(v string) *MerchandisePlacementDetectionProResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MerchandisePlacementDetectionProResponseBody
	GetSuccess() *bool
}

type MerchandisePlacementDetectionProResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detection result of product display detection Pro.
	Data *MerchandisePlacementDetectionProResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s MerchandisePlacementDetectionProResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionProResponseBody) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionProResponseBody) GetCode() *string {
	return s.Code
}

func (s *MerchandisePlacementDetectionProResponseBody) GetData() *MerchandisePlacementDetectionProResponseBodyData {
	return s.Data
}

func (s *MerchandisePlacementDetectionProResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MerchandisePlacementDetectionProResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MerchandisePlacementDetectionProResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MerchandisePlacementDetectionProResponseBody) SetCode(v string) *MerchandisePlacementDetectionProResponseBody {
	s.Code = &v
	return s
}

func (s *MerchandisePlacementDetectionProResponseBody) SetData(v *MerchandisePlacementDetectionProResponseBodyData) *MerchandisePlacementDetectionProResponseBody {
	s.Data = v
	return s
}

func (s *MerchandisePlacementDetectionProResponseBody) SetMessage(v string) *MerchandisePlacementDetectionProResponseBody {
	s.Message = &v
	return s
}

func (s *MerchandisePlacementDetectionProResponseBody) SetRequestId(v string) *MerchandisePlacementDetectionProResponseBody {
	s.RequestId = &v
	return s
}

func (s *MerchandisePlacementDetectionProResponseBody) SetSuccess(v bool) *MerchandisePlacementDetectionProResponseBody {
	s.Success = &v
	return s
}

func (s *MerchandisePlacementDetectionProResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MerchandisePlacementDetectionProResponseBodyData struct {
	// The number of valid bounding boxes.
	//
	// example:
	//
	// 3
	BoxCount *int32 `json:"BoxCount,omitempty" xml:"BoxCount,omitempty"`
	// The list of per-box detection details.
	Data []*MerchandisePlacementDetectionProResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The usage information.
	//
	// example:
	//
	// {"ProcessingCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s MerchandisePlacementDetectionProResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionProResponseBodyData) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionProResponseBodyData) GetBoxCount() *int32 {
	return s.BoxCount
}

func (s *MerchandisePlacementDetectionProResponseBodyData) GetData() []*MerchandisePlacementDetectionProResponseBodyDataData {
	return s.Data
}

func (s *MerchandisePlacementDetectionProResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *MerchandisePlacementDetectionProResponseBodyData) SetBoxCount(v int32) *MerchandisePlacementDetectionProResponseBodyData {
	s.BoxCount = &v
	return s
}

func (s *MerchandisePlacementDetectionProResponseBodyData) SetData(v []*MerchandisePlacementDetectionProResponseBodyDataData) *MerchandisePlacementDetectionProResponseBodyData {
	s.Data = v
	return s
}

func (s *MerchandisePlacementDetectionProResponseBodyData) SetUsageMap(v map[string]*int64) *MerchandisePlacementDetectionProResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *MerchandisePlacementDetectionProResponseBodyData) Validate() error {
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

type MerchandisePlacementDetectionProResponseBodyDataData struct {
	// The normalized bounding box coordinates [x1,y1,x2,y2], with values in the range 0–1000.
	Bbox2d []*int32 `json:"Bbox2d,omitempty" xml:"Bbox2d,omitempty" type:"Repeated"`
	// The detected product name. The value is unknown if the name cannot be determined.
	//
	// example:
	//
	// unknown
	DetectedSkuName *string `json:"DetectedSkuName,omitempty" xml:"DetectedSkuName,omitempty"`
	// The bounding box index, starting from 1.
	//
	// example:
	//
	// 1
	Idx *int32 `json:"Idx,omitempty" xml:"Idx,omitempty"`
}

func (s MerchandisePlacementDetectionProResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionProResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionProResponseBodyDataData) GetBbox2d() []*int32 {
	return s.Bbox2d
}

func (s *MerchandisePlacementDetectionProResponseBodyDataData) GetDetectedSkuName() *string {
	return s.DetectedSkuName
}

func (s *MerchandisePlacementDetectionProResponseBodyDataData) GetIdx() *int32 {
	return s.Idx
}

func (s *MerchandisePlacementDetectionProResponseBodyDataData) SetBbox2d(v []*int32) *MerchandisePlacementDetectionProResponseBodyDataData {
	s.Bbox2d = v
	return s
}

func (s *MerchandisePlacementDetectionProResponseBodyDataData) SetDetectedSkuName(v string) *MerchandisePlacementDetectionProResponseBodyDataData {
	s.DetectedSkuName = &v
	return s
}

func (s *MerchandisePlacementDetectionProResponseBodyDataData) SetIdx(v int32) *MerchandisePlacementDetectionProResponseBodyDataData {
	s.Idx = &v
	return s
}

func (s *MerchandisePlacementDetectionProResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}
