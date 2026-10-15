// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMilkPowderCartonCodeRephotographyDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MilkPowderCartonCodeRephotographyDetectionResponseBody
	GetCode() *string
	SetData(v *MilkPowderCartonCodeRephotographyDetectionResponseBodyData) *MilkPowderCartonCodeRephotographyDetectionResponseBody
	GetData() *MilkPowderCartonCodeRephotographyDetectionResponseBodyData
	SetMessage(v string) *MilkPowderCartonCodeRephotographyDetectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *MilkPowderCartonCodeRephotographyDetectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MilkPowderCartonCodeRephotographyDetectionResponseBody
	GetSuccess() *bool
}

type MilkPowderCartonCodeRephotographyDetectionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The recapture detection result.
	Data *MilkPowderCartonCodeRephotographyDetectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s MilkPowderCartonCodeRephotographyDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCartonCodeRephotographyDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBody) GetData() *MilkPowderCartonCodeRephotographyDetectionResponseBodyData {
	return s.Data
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBody) SetCode(v string) *MilkPowderCartonCodeRephotographyDetectionResponseBody {
	s.Code = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBody) SetData(v *MilkPowderCartonCodeRephotographyDetectionResponseBodyData) *MilkPowderCartonCodeRephotographyDetectionResponseBody {
	s.Data = v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBody) SetMessage(v string) *MilkPowderCartonCodeRephotographyDetectionResponseBody {
	s.Message = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBody) SetRequestId(v string) *MilkPowderCartonCodeRephotographyDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBody) SetSuccess(v bool) *MilkPowderCartonCodeRephotographyDetectionResponseBody {
	s.Success = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MilkPowderCartonCodeRephotographyDetectionResponseBodyData struct {
	// The details of the recapture detection result.
	Result *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The usage information. The value is `{"ProcessingCount":1}` if the call is successful, or an empty object if the call fails.
	//
	// example:
	//
	// {"ProcessingCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s MilkPowderCartonCodeRephotographyDetectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCartonCodeRephotographyDetectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyData) GetResult() *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult {
	return s.Result
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyData) SetResult(v *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) *MilkPowderCartonCodeRephotographyDetectionResponseBodyData {
	s.Result = v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyData) SetUsageMap(v map[string]*int64) *MilkPowderCartonCodeRephotographyDetectionResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult struct {
	// The supplementary element grounding results.
	Grounding *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGrounding `json:"Grounding,omitempty" xml:"Grounding,omitempty" type:"Struct"`
	// Indicates whether the image is recaptured.
	//
	// example:
	//
	// true
	IsFake *bool `json:"IsFake,omitempty" xml:"IsFake,omitempty"`
	// Indicates whether the image is suspected to be recaptured.
	//
	// example:
	//
	// true
	IsSuspect *bool `json:"IsSuspect,omitempty" xml:"IsSuspect,omitempty"`
	// The recapture status. Valid values:
	//
	// - `0`: Non-recapture.
	//
	// - `1`: Recapture.
	//
	// - `2`: Suspected recapture.
	//
	// example:
	//
	// 1
	PhotoStatus *int32 `json:"PhotoStatus,omitempty" xml:"PhotoStatus,omitempty"`
	// The confidence score for recapture detection.
	//
	// example:
	//
	// 0.89
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The detection type. The value is fixed to `boxCode`.
	//
	// example:
	//
	// boxCode
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) GetGrounding() *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGrounding {
	return s.Grounding
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) GetIsFake() *bool {
	return s.IsFake
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) GetIsSuspect() *bool {
	return s.IsSuspect
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) GetPhotoStatus() *int32 {
	return s.PhotoStatus
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) GetScore() *float32 {
	return s.Score
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) GetType() *string {
	return s.Type
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) SetGrounding(v *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGrounding) *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult {
	s.Grounding = v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) SetIsFake(v bool) *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult {
	s.IsFake = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) SetIsSuspect(v bool) *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult {
	s.IsSuspect = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) SetPhotoStatus(v int32) *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult {
	s.PhotoStatus = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) SetScore(v float32) *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult {
	s.Score = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) SetType(v string) *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult {
	s.Type = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResult) Validate() error {
	if s.Grounding != nil {
		if err := s.Grounding.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGrounding struct {
	// The coverage of the grounding. Valid values:
	//
	// - `complete`: The relevant visible targets are complete.
	//
	// - `partial`: Only some targets are valid or recognizable.
	//
	// example:
	//
	// complete
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The array of targets, which contains up to 12 items. This array can be empty if no relevant targets are found.
	Regions []*MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGrounding) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGrounding) GoString() string {
	return s.String()
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGrounding) GetCoverage() *string {
	return s.Coverage
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGrounding) GetRegions() []*MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions {
	return s.Regions
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGrounding) SetCoverage(v string) *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGrounding {
	s.Coverage = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGrounding) SetRegions(v []*MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGrounding {
	s.Regions = v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGrounding) Validate() error {
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

type MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions struct {
	// The coordinates of the bounding box in the `[top-left x, top-left y, bottom-right x, bottom-right y]` format. The value range is 0 to 1000. The left value must be less than the right value, and the top value must be less than the bottom value.
	Bbox2d []*float32 `json:"Bbox2d,omitempty" xml:"Bbox2d,omitempty" type:"Repeated"`
	// The category of the target. For enumerated values, see the following table.
	//
	// example:
	//
	// milk_powder
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The confirmed short name or text, which contains up to 24 characters. This field is an empty string if the text is unclear.
	//
	// example:
	//
	// Milk powder
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) GoString() string {
	return s.String()
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) GetBbox2d() []*float32 {
	return s.Bbox2d
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) GetLabel() *string {
	return s.Label
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) GetText() *string {
	return s.Text
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) SetBbox2d(v []*float32) *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions {
	s.Bbox2d = v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) SetLabel(v string) *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions {
	s.Label = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) SetText(v string) *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions {
	s.Text = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponseBodyDataResultGroundingRegions) Validate() error {
	return dara.Validate(s)
}
