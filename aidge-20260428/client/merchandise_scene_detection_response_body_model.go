// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandiseSceneDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MerchandiseSceneDetectionResponseBody
	GetCode() *string
	SetData(v *MerchandiseSceneDetectionResponseBodyData) *MerchandiseSceneDetectionResponseBody
	GetData() *MerchandiseSceneDetectionResponseBodyData
	SetMessage(v string) *MerchandiseSceneDetectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *MerchandiseSceneDetectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MerchandiseSceneDetectionResponseBody
	GetSuccess() *bool
}

type MerchandiseSceneDetectionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The display scene detection result.
	Data *MerchandiseSceneDetectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s MerchandiseSceneDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MerchandiseSceneDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *MerchandiseSceneDetectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *MerchandiseSceneDetectionResponseBody) GetData() *MerchandiseSceneDetectionResponseBodyData {
	return s.Data
}

func (s *MerchandiseSceneDetectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MerchandiseSceneDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MerchandiseSceneDetectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MerchandiseSceneDetectionResponseBody) SetCode(v string) *MerchandiseSceneDetectionResponseBody {
	s.Code = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBody) SetData(v *MerchandiseSceneDetectionResponseBodyData) *MerchandiseSceneDetectionResponseBody {
	s.Data = v
	return s
}

func (s *MerchandiseSceneDetectionResponseBody) SetMessage(v string) *MerchandiseSceneDetectionResponseBody {
	s.Message = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBody) SetRequestId(v string) *MerchandiseSceneDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBody) SetSuccess(v bool) *MerchandiseSceneDetectionResponseBody {
	s.Success = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MerchandiseSceneDetectionResponseBodyData struct {
	// The details of the display scene detection result.
	Result *MerchandiseSceneDetectionResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The usage information.
	//
	// example:
	//
	// {"ProcessingCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s MerchandiseSceneDetectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MerchandiseSceneDetectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *MerchandiseSceneDetectionResponseBodyData) GetResult() *MerchandiseSceneDetectionResponseBodyDataResult {
	return s.Result
}

func (s *MerchandiseSceneDetectionResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *MerchandiseSceneDetectionResponseBodyData) SetResult(v *MerchandiseSceneDetectionResponseBodyDataResult) *MerchandiseSceneDetectionResponseBodyData {
	s.Result = v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyData) SetUsageMap(v map[string]*int64) *MerchandiseSceneDetectionResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MerchandiseSceneDetectionResponseBodyDataResult struct {
	// The positioning information.
	Grounding *MerchandiseSceneDetectionResponseBodyDataResultGrounding `json:"Grounding,omitempty" xml:"Grounding,omitempty" type:"Struct"`
	// The list of recognized scenes.
	Scenes []*MerchandiseSceneDetectionResponseBodyDataResultScenes `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
	// Fixed value: scene.
	//
	// example:
	//
	// scene
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MerchandiseSceneDetectionResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s MerchandiseSceneDetectionResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *MerchandiseSceneDetectionResponseBodyDataResult) GetGrounding() *MerchandiseSceneDetectionResponseBodyDataResultGrounding {
	return s.Grounding
}

func (s *MerchandiseSceneDetectionResponseBodyDataResult) GetScenes() []*MerchandiseSceneDetectionResponseBodyDataResultScenes {
	return s.Scenes
}

func (s *MerchandiseSceneDetectionResponseBodyDataResult) GetType() *string {
	return s.Type
}

func (s *MerchandiseSceneDetectionResponseBodyDataResult) SetGrounding(v *MerchandiseSceneDetectionResponseBodyDataResultGrounding) *MerchandiseSceneDetectionResponseBodyDataResult {
	s.Grounding = v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResult) SetScenes(v []*MerchandiseSceneDetectionResponseBodyDataResultScenes) *MerchandiseSceneDetectionResponseBodyDataResult {
	s.Scenes = v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResult) SetType(v string) *MerchandiseSceneDetectionResponseBodyDataResult {
	s.Type = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResult) Validate() error {
	if s.Grounding != nil {
		if err := s.Grounding.Validate(); err != nil {
			return err
		}
	}
	if s.Scenes != nil {
		for _, item := range s.Scenes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MerchandiseSceneDetectionResponseBodyDataResultGrounding struct {
	// The coverage.
	//
	// example:
	//
	// complete
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The list of regions.
	Regions []*MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s MerchandiseSceneDetectionResponseBodyDataResultGrounding) String() string {
	return dara.Prettify(s)
}

func (s MerchandiseSceneDetectionResponseBodyDataResultGrounding) GoString() string {
	return s.String()
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultGrounding) GetCoverage() *string {
	return s.Coverage
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultGrounding) GetRegions() []*MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions {
	return s.Regions
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultGrounding) SetCoverage(v string) *MerchandiseSceneDetectionResponseBodyDataResultGrounding {
	s.Coverage = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultGrounding) SetRegions(v []*MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions) *MerchandiseSceneDetectionResponseBodyDataResultGrounding {
	s.Regions = v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultGrounding) Validate() error {
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

type MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions struct {
	// The 2D bounding box in the format of [top-left x, top-left y, bottom-right x, bottom-right y]. Valid values: 0 to 1000. The left coordinate must be less than the right coordinate, and the top coordinate must be less than the bottom coordinate.
	Bbox2d []*float64 `json:"Bbox_2d,omitempty" xml:"Bbox_2d,omitempty" type:"Repeated"`
	// The label.
	//
	// example:
	//
	// shelf
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The recognized text.
	//
	// example:
	//
	// Multi-tier product display shelf
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions) String() string {
	return dara.Prettify(s)
}

func (s MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions) GoString() string {
	return s.String()
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions) GetBbox2d() []*float64 {
	return s.Bbox2d
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions) GetLabel() *string {
	return s.Label
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions) GetText() *string {
	return s.Text
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions) SetBbox2d(v []*float64) *MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions {
	s.Bbox2d = v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions) SetLabel(v string) *MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions {
	s.Label = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions) SetText(v string) *MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions {
	s.Text = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultGroundingRegions) Validate() error {
	return dara.Validate(s)
}

type MerchandiseSceneDetectionResponseBodyDataResultScenes struct {
	// Optional. The area proportion.
	//
	// example:
	//
	// 0.985
	AreaRate *float64 `json:"AreaRate,omitempty" xml:"AreaRate,omitempty"`
	// Optional. The purity metric for proprietary products. The specific business criteria are subject to the recognition service.
	//
	// example:
	//
	// 0
	OwnPurity *float64 `json:"OwnPurity,omitempty" xml:"OwnPurity,omitempty"`
	// Optional. The display plumpness metric.
	//
	// example:
	//
	// 0
	Plumpness *float64 `json:"Plumpness,omitempty" xml:"Plumpness,omitempty"`
	// The display name of the scene. If omitted, the value is an empty string.
	//
	// example:
	//
	// Air-cooled display cabinet
	SceneActualName *string `json:"SceneActualName,omitempty" xml:"SceneActualName,omitempty"`
	// The center point of the scene, which retains the original normalized coordinates. If omitted, the value is an empty array.
	SceneCenter []*float64 `json:"SceneCenter,omitempty" xml:"SceneCenter,omitempty" type:"Repeated"`
	// The list of region points. Each item is in the {x,y} format and retains the original normalized coordinates. If omitted, the value is an empty array.
	SceneLocation []*MerchandiseSceneDetectionResponseBodyDataResultScenesSceneLocation `json:"SceneLocation,omitempty" xml:"SceneLocation,omitempty" type:"Repeated"`
	// The scene identifier. This parameter is required and cannot be empty.
	//
	// example:
	//
	// yl_cj_lfg
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// Optional. The shelf distribution proportion metric. The specific business criteria are subject to the recognition service.
	//
	// example:
	//
	// 0
	ShelfDistributionRate *float64 `json:"ShelfDistributionRate,omitempty" xml:"ShelfDistributionRate,omitempty"`
}

func (s MerchandiseSceneDetectionResponseBodyDataResultScenes) String() string {
	return dara.Prettify(s)
}

func (s MerchandiseSceneDetectionResponseBodyDataResultScenes) GoString() string {
	return s.String()
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) GetAreaRate() *float64 {
	return s.AreaRate
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) GetOwnPurity() *float64 {
	return s.OwnPurity
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) GetPlumpness() *float64 {
	return s.Plumpness
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) GetSceneActualName() *string {
	return s.SceneActualName
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) GetSceneCenter() []*float64 {
	return s.SceneCenter
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) GetSceneLocation() []*MerchandiseSceneDetectionResponseBodyDataResultScenesSceneLocation {
	return s.SceneLocation
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) GetSceneName() *string {
	return s.SceneName
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) GetShelfDistributionRate() *float64 {
	return s.ShelfDistributionRate
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) SetAreaRate(v float64) *MerchandiseSceneDetectionResponseBodyDataResultScenes {
	s.AreaRate = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) SetOwnPurity(v float64) *MerchandiseSceneDetectionResponseBodyDataResultScenes {
	s.OwnPurity = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) SetPlumpness(v float64) *MerchandiseSceneDetectionResponseBodyDataResultScenes {
	s.Plumpness = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) SetSceneActualName(v string) *MerchandiseSceneDetectionResponseBodyDataResultScenes {
	s.SceneActualName = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) SetSceneCenter(v []*float64) *MerchandiseSceneDetectionResponseBodyDataResultScenes {
	s.SceneCenter = v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) SetSceneLocation(v []*MerchandiseSceneDetectionResponseBodyDataResultScenesSceneLocation) *MerchandiseSceneDetectionResponseBodyDataResultScenes {
	s.SceneLocation = v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) SetSceneName(v string) *MerchandiseSceneDetectionResponseBodyDataResultScenes {
	s.SceneName = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) SetShelfDistributionRate(v float64) *MerchandiseSceneDetectionResponseBodyDataResultScenes {
	s.ShelfDistributionRate = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenes) Validate() error {
	if s.SceneLocation != nil {
		for _, item := range s.SceneLocation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MerchandiseSceneDetectionResponseBodyDataResultScenesSceneLocation struct {
	// The x-coordinate of the region point.
	//
	// example:
	//
	// 0.431
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate of the region point.
	//
	// example:
	//
	// 0.973
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s MerchandiseSceneDetectionResponseBodyDataResultScenesSceneLocation) String() string {
	return dara.Prettify(s)
}

func (s MerchandiseSceneDetectionResponseBodyDataResultScenesSceneLocation) GoString() string {
	return s.String()
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenesSceneLocation) GetX() *float64 {
	return s.X
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenesSceneLocation) GetY() *float64 {
	return s.Y
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenesSceneLocation) SetX(v float64) *MerchandiseSceneDetectionResponseBodyDataResultScenesSceneLocation {
	s.X = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenesSceneLocation) SetY(v float64) *MerchandiseSceneDetectionResponseBodyDataResultScenesSceneLocation {
	s.Y = &v
	return s
}

func (s *MerchandiseSceneDetectionResponseBodyDataResultScenesSceneLocation) Validate() error {
	return dara.Validate(s)
}
