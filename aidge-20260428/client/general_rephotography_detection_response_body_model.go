// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGeneralRephotographyDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GeneralRephotographyDetectionResponseBody
	GetCode() *string
	SetData(v *GeneralRephotographyDetectionResponseBodyData) *GeneralRephotographyDetectionResponseBody
	GetData() *GeneralRephotographyDetectionResponseBodyData
	SetMessage(v string) *GeneralRephotographyDetectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GeneralRephotographyDetectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GeneralRephotographyDetectionResponseBody
	GetSuccess() *bool
}

type GeneralRephotographyDetectionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The recapture detection result.
	Data *GeneralRephotographyDetectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GeneralRephotographyDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GeneralRephotographyDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *GeneralRephotographyDetectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GeneralRephotographyDetectionResponseBody) GetData() *GeneralRephotographyDetectionResponseBodyData {
	return s.Data
}

func (s *GeneralRephotographyDetectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GeneralRephotographyDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GeneralRephotographyDetectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GeneralRephotographyDetectionResponseBody) SetCode(v string) *GeneralRephotographyDetectionResponseBody {
	s.Code = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBody) SetData(v *GeneralRephotographyDetectionResponseBodyData) *GeneralRephotographyDetectionResponseBody {
	s.Data = v
	return s
}

func (s *GeneralRephotographyDetectionResponseBody) SetMessage(v string) *GeneralRephotographyDetectionResponseBody {
	s.Message = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBody) SetRequestId(v string) *GeneralRephotographyDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBody) SetSuccess(v bool) *GeneralRephotographyDetectionResponseBody {
	s.Success = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GeneralRephotographyDetectionResponseBodyData struct {
	// The business result. This value is an empty object if the request fails.
	Result *GeneralRephotographyDetectionResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The usage information. The value is `{"ProcessingCount":1}` on success, or an empty object on failure.
	//
	// example:
	//
	// {"ProcessingCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s GeneralRephotographyDetectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GeneralRephotographyDetectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GeneralRephotographyDetectionResponseBodyData) GetResult() *GeneralRephotographyDetectionResponseBodyDataResult {
	return s.Result
}

func (s *GeneralRephotographyDetectionResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *GeneralRephotographyDetectionResponseBodyData) SetResult(v *GeneralRephotographyDetectionResponseBodyDataResult) *GeneralRephotographyDetectionResponseBodyData {
	s.Result = v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyData) SetUsageMap(v map[string]*int64) *GeneralRephotographyDetectionResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GeneralRephotographyDetectionResponseBodyDataResult struct {
	// The supplementary element localization result.
	Grounding *GeneralRephotographyDetectionResponseBodyDataResultGrounding `json:"Grounding,omitempty" xml:"Grounding,omitempty" type:"Struct"`
	// Indicates whether the image is a recaptured photo.
	//
	// example:
	//
	// true
	IsFake *bool `json:"IsFake,omitempty" xml:"IsFake,omitempty"`
	// The detection type. The value is fixed as general.
	//
	// example:
	//
	// general
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GeneralRephotographyDetectionResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GeneralRephotographyDetectionResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GeneralRephotographyDetectionResponseBodyDataResult) GetGrounding() *GeneralRephotographyDetectionResponseBodyDataResultGrounding {
	return s.Grounding
}

func (s *GeneralRephotographyDetectionResponseBodyDataResult) GetIsFake() *bool {
	return s.IsFake
}

func (s *GeneralRephotographyDetectionResponseBodyDataResult) GetType() *string {
	return s.Type
}

func (s *GeneralRephotographyDetectionResponseBodyDataResult) SetGrounding(v *GeneralRephotographyDetectionResponseBodyDataResultGrounding) *GeneralRephotographyDetectionResponseBodyDataResult {
	s.Grounding = v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyDataResult) SetIsFake(v bool) *GeneralRephotographyDetectionResponseBodyDataResult {
	s.IsFake = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyDataResult) SetType(v string) *GeneralRephotographyDetectionResponseBodyDataResult {
	s.Type = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyDataResult) Validate() error {
	if s.Grounding != nil {
		if err := s.Grounding.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GeneralRephotographyDetectionResponseBodyDataResultGrounding struct {
	// The coverage of the localization. Valid values:
	//
	// - complete: All relevant visible targets are fully covered.
	//
	// - partial: Only some targets are valid or recognizable.
	//
	// example:
	//
	// complete
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The array of targets. A maximum of 12 items are returned. This value can be empty if no relevant targets exist.
	Regions []*GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s GeneralRephotographyDetectionResponseBodyDataResultGrounding) String() string {
	return dara.Prettify(s)
}

func (s GeneralRephotographyDetectionResponseBodyDataResultGrounding) GoString() string {
	return s.String()
}

func (s *GeneralRephotographyDetectionResponseBodyDataResultGrounding) GetCoverage() *string {
	return s.Coverage
}

func (s *GeneralRephotographyDetectionResponseBodyDataResultGrounding) GetRegions() []*GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions {
	return s.Regions
}

func (s *GeneralRephotographyDetectionResponseBodyDataResultGrounding) SetCoverage(v string) *GeneralRephotographyDetectionResponseBodyDataResultGrounding {
	s.Coverage = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyDataResultGrounding) SetRegions(v []*GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions) *GeneralRephotographyDetectionResponseBodyDataResultGrounding {
	s.Regions = v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyDataResultGrounding) Validate() error {
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

type GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions struct {
	// The bounding box coordinates.
	Bbox2d []*float32 `json:"Bbox2d,omitempty" xml:"Bbox2d,omitempty" type:"Repeated"`
	// The target category. For valid values, see the table below.
	//
	// example:
	//
	// product
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The confirmed short name or text of up to 24 characters. This value is an empty string if the text is unreadable.
	//
	// example:
	//
	// Product
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions) String() string {
	return dara.Prettify(s)
}

func (s GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions) GoString() string {
	return s.String()
}

func (s *GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions) GetBbox2d() []*float32 {
	return s.Bbox2d
}

func (s *GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions) GetLabel() *string {
	return s.Label
}

func (s *GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions) GetText() *string {
	return s.Text
}

func (s *GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions) SetBbox2d(v []*float32) *GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions {
	s.Bbox2d = v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions) SetLabel(v string) *GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions {
	s.Label = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions) SetText(v string) *GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions {
	s.Text = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyDataResultGroundingRegions) Validate() error {
	return dara.Validate(s)
}
