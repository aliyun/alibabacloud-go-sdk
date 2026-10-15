// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStorefrontSceneRecognitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StorefrontSceneRecognitionResponseBody
	GetCode() *string
	SetData(v *StorefrontSceneRecognitionResponseBodyData) *StorefrontSceneRecognitionResponseBody
	GetData() *StorefrontSceneRecognitionResponseBodyData
	SetMessage(v string) *StorefrontSceneRecognitionResponseBody
	GetMessage() *string
	SetRequestId(v string) *StorefrontSceneRecognitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StorefrontSceneRecognitionResponseBody
	GetSuccess() *bool
}

type StorefrontSceneRecognitionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The scene recognition result.
	Data *StorefrontSceneRecognitionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message or failure description.
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

func (s StorefrontSceneRecognitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StorefrontSceneRecognitionResponseBody) GoString() string {
	return s.String()
}

func (s *StorefrontSceneRecognitionResponseBody) GetCode() *string {
	return s.Code
}

func (s *StorefrontSceneRecognitionResponseBody) GetData() *StorefrontSceneRecognitionResponseBodyData {
	return s.Data
}

func (s *StorefrontSceneRecognitionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StorefrontSceneRecognitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StorefrontSceneRecognitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StorefrontSceneRecognitionResponseBody) SetCode(v string) *StorefrontSceneRecognitionResponseBody {
	s.Code = &v
	return s
}

func (s *StorefrontSceneRecognitionResponseBody) SetData(v *StorefrontSceneRecognitionResponseBodyData) *StorefrontSceneRecognitionResponseBody {
	s.Data = v
	return s
}

func (s *StorefrontSceneRecognitionResponseBody) SetMessage(v string) *StorefrontSceneRecognitionResponseBody {
	s.Message = &v
	return s
}

func (s *StorefrontSceneRecognitionResponseBody) SetRequestId(v string) *StorefrontSceneRecognitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StorefrontSceneRecognitionResponseBody) SetSuccess(v bool) *StorefrontSceneRecognitionResponseBody {
	s.Success = &v
	return s
}

func (s *StorefrontSceneRecognitionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StorefrontSceneRecognitionResponseBodyData struct {
	// The scene recognition result details.
	Result *StorefrontSceneRecognitionResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The usage information.
	//
	// example:
	//
	// {"ProcessingCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s StorefrontSceneRecognitionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StorefrontSceneRecognitionResponseBodyData) GoString() string {
	return s.String()
}

func (s *StorefrontSceneRecognitionResponseBodyData) GetResult() *StorefrontSceneRecognitionResponseBodyDataResult {
	return s.Result
}

func (s *StorefrontSceneRecognitionResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *StorefrontSceneRecognitionResponseBodyData) SetResult(v *StorefrontSceneRecognitionResponseBodyDataResult) *StorefrontSceneRecognitionResponseBodyData {
	s.Result = v
	return s
}

func (s *StorefrontSceneRecognitionResponseBodyData) SetUsageMap(v map[string]*int64) *StorefrontSceneRecognitionResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *StorefrontSceneRecognitionResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StorefrontSceneRecognitionResponseBodyDataResult struct {
	// The classification details. This can be an empty array.
	Details []*StorefrontSceneRecognitionResponseBodyDataResultDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The localization information.
	Grounding *StorefrontSceneRecognitionResponseBodyDataResultGrounding `json:"Grounding,omitempty" xml:"Grounding,omitempty" type:"Struct"`
	// The non-empty scene classification array.
	Scenes []*string `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
}

func (s StorefrontSceneRecognitionResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s StorefrontSceneRecognitionResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *StorefrontSceneRecognitionResponseBodyDataResult) GetDetails() []*StorefrontSceneRecognitionResponseBodyDataResultDetails {
	return s.Details
}

func (s *StorefrontSceneRecognitionResponseBodyDataResult) GetGrounding() *StorefrontSceneRecognitionResponseBodyDataResultGrounding {
	return s.Grounding
}

func (s *StorefrontSceneRecognitionResponseBodyDataResult) GetScenes() []*string {
	return s.Scenes
}

func (s *StorefrontSceneRecognitionResponseBodyDataResult) SetDetails(v []*StorefrontSceneRecognitionResponseBodyDataResultDetails) *StorefrontSceneRecognitionResponseBodyDataResult {
	s.Details = v
	return s
}

func (s *StorefrontSceneRecognitionResponseBodyDataResult) SetGrounding(v *StorefrontSceneRecognitionResponseBodyDataResultGrounding) *StorefrontSceneRecognitionResponseBodyDataResult {
	s.Grounding = v
	return s
}

func (s *StorefrontSceneRecognitionResponseBodyDataResult) SetScenes(v []*string) *StorefrontSceneRecognitionResponseBodyDataResult {
	s.Scenes = v
	return s
}

func (s *StorefrontSceneRecognitionResponseBodyDataResult) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Grounding != nil {
		if err := s.Grounding.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StorefrontSceneRecognitionResponseBodyDataResultDetails struct {
	// The non-empty category name.
	//
	// example:
	//
	// oppo vivo smartphone specialty store
	Clazz *string `json:"Clazz,omitempty" xml:"Clazz,omitempty"`
	// A finite numerical value. The source document does not specify a value range.
	//
	// example:
	//
	// 0.99
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s StorefrontSceneRecognitionResponseBodyDataResultDetails) String() string {
	return dara.Prettify(s)
}

func (s StorefrontSceneRecognitionResponseBodyDataResultDetails) GoString() string {
	return s.String()
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultDetails) GetClazz() *string {
	return s.Clazz
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultDetails) GetConfidence() *float32 {
	return s.Confidence
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultDetails) SetClazz(v string) *StorefrontSceneRecognitionResponseBodyDataResultDetails {
	s.Clazz = &v
	return s
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultDetails) SetConfidence(v float32) *StorefrontSceneRecognitionResponseBodyDataResultDetails {
	s.Confidence = &v
	return s
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultDetails) Validate() error {
	return dara.Validate(s)
}

type StorefrontSceneRecognitionResponseBodyDataResultGrounding struct {
	// The coverage.
	//
	// example:
	//
	// partial
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The region list.
	Regions []*StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s StorefrontSceneRecognitionResponseBodyDataResultGrounding) String() string {
	return dara.Prettify(s)
}

func (s StorefrontSceneRecognitionResponseBodyDataResultGrounding) GoString() string {
	return s.String()
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultGrounding) GetCoverage() *string {
	return s.Coverage
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultGrounding) GetRegions() []*StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions {
	return s.Regions
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultGrounding) SetCoverage(v string) *StorefrontSceneRecognitionResponseBodyDataResultGrounding {
	s.Coverage = &v
	return s
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultGrounding) SetRegions(v []*StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions) *StorefrontSceneRecognitionResponseBodyDataResultGrounding {
	s.Regions = v
	return s
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultGrounding) Validate() error {
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

type StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions struct {
	// The bounding box coordinates.
	Bbox2d []*float32 `json:"Bbox2d,omitempty" xml:"Bbox2d,omitempty" type:"Repeated"`
	// The label.
	//
	// example:
	//
	// store_sign
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The recognized text.
	//
	// example:
	//
	// oppo
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions) String() string {
	return dara.Prettify(s)
}

func (s StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions) GoString() string {
	return s.String()
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions) GetBbox2d() []*float32 {
	return s.Bbox2d
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions) GetLabel() *string {
	return s.Label
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions) GetText() *string {
	return s.Text
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions) SetBbox2d(v []*float32) *StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions {
	s.Bbox2d = v
	return s
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions) SetLabel(v string) *StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions {
	s.Label = &v
	return s
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions) SetText(v string) *StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions {
	s.Text = &v
	return s
}

func (s *StorefrontSceneRecognitionResponseBodyDataResultGroundingRegions) Validate() error {
	return dara.Validate(s)
}
