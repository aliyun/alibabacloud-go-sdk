// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMilkPowderCanBottomCodeRephotographyDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MilkPowderCanBottomCodeRephotographyDetectionResponseBody
	GetCode() *string
	SetData(v *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData) *MilkPowderCanBottomCodeRephotographyDetectionResponseBody
	GetData() *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData
	SetMessage(v string) *MilkPowderCanBottomCodeRephotographyDetectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *MilkPowderCanBottomCodeRephotographyDetectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MilkPowderCanBottomCodeRephotographyDetectionResponseBody
	GetSuccess() *bool
}

type MilkPowderCanBottomCodeRephotographyDetectionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The recapture detection result.
	Data *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Specifies whether the call is successful. A valid business conclusion does not mean that the review is passed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MilkPowderCanBottomCodeRephotographyDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCanBottomCodeRephotographyDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBody) GetData() *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData {
	return s.Data
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBody) SetCode(v string) *MilkPowderCanBottomCodeRephotographyDetectionResponseBody {
	s.Code = &v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBody) SetData(v *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData) *MilkPowderCanBottomCodeRephotographyDetectionResponseBody {
	s.Data = v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBody) SetMessage(v string) *MilkPowderCanBottomCodeRephotographyDetectionResponseBody {
	s.Message = &v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBody) SetRequestId(v string) *MilkPowderCanBottomCodeRephotographyDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBody) SetSuccess(v bool) *MilkPowderCanBottomCodeRephotographyDetectionResponseBody {
	s.Success = &v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData struct {
	// The details of the recapture detection result.
	Result *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The usage information. If the call is successful, the value is `{"ProcessingCount":1}`. If the call fails, an empty object is returned.
	//
	// example:
	//
	// {"ProcessingCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData) GetResult() *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult {
	return s.Result
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData) SetResult(v *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult) *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData {
	s.Result = v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData) SetUsageMap(v map[string]*int64) *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult struct {
	// The localization information.
	Grounding *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGrounding `json:"Grounding,omitempty" xml:"Grounding,omitempty" type:"Struct"`
	// Specifies whether the image is recaptured.
	//
	// example:
	//
	// true
	IsFake *bool `json:"IsFake,omitempty" xml:"IsFake,omitempty"`
	// The detection type. The value is fixed to canBottom.
	//
	// example:
	//
	// canBottom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult) GetGrounding() *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGrounding {
	return s.Grounding
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult) GetIsFake() *bool {
	return s.IsFake
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult) GetType() *string {
	return s.Type
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult) SetGrounding(v *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGrounding) *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult {
	s.Grounding = v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult) SetIsFake(v bool) *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult {
	s.IsFake = &v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult) SetType(v string) *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult {
	s.Type = &v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResult) Validate() error {
	if s.Grounding != nil {
		if err := s.Grounding.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGrounding struct {
	// The coverage. Valid values:
	//
	// - complete: The relevant visible targets are complete.
	//
	// - partial: Only some targets are valid or recognizable.
	//
	// example:
	//
	// complete
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The list of regions.
	Regions []*MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGrounding) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGrounding) GoString() string {
	return s.String()
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGrounding) GetCoverage() *string {
	return s.Coverage
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGrounding) GetRegions() []*MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions {
	return s.Regions
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGrounding) SetCoverage(v string) *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGrounding {
	s.Coverage = &v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGrounding) SetRegions(v []*MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGrounding {
	s.Regions = v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGrounding) Validate() error {
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

type MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions struct {
	// The bounding box coordinates in the format of `[top-left x, top-left y, bottom-right x, bottom-right y]`. Valid values: 0 to 1000. The left value must be less than the right value, and the top value must be less than the bottom value.
	Bbox2d []*float32 `json:"Bbox2d,omitempty" xml:"Bbox2d,omitempty" type:"Repeated"`
	// The target category. For valid values, see the following table.
	//
	// example:
	//
	// milk_powder
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The confirmed short name or text of up to 24 characters. If the text is unclear, an empty string is returned.
	//
	// example:
	//
	// Milk powder
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) GoString() string {
	return s.String()
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) GetBbox2d() []*float32 {
	return s.Bbox2d
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) GetLabel() *string {
	return s.Label
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) GetText() *string {
	return s.Text
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) SetBbox2d(v []*float32) *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions {
	s.Bbox2d = v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) SetLabel(v string) *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions {
	s.Label = &v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) SetText(v string) *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions {
	s.Text = &v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) Validate() error {
	return dara.Validate(s)
}
