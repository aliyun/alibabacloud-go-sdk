// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageBatchModerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImageBatchModerationResponseBody
	GetCode() *int32
	SetData(v *ImageBatchModerationResponseBodyData) *ImageBatchModerationResponseBody
	GetData() *ImageBatchModerationResponseBodyData
	SetMsg(v string) *ImageBatchModerationResponseBody
	GetMsg() *string
	SetRequestId(v string) *ImageBatchModerationResponseBody
	GetRequestId() *string
}

type ImageBatchModerationResponseBody struct {
	// The return code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The results of the image content moderation.
	Data *ImageBatchModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message for the request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The unique ID of the request. Alibaba Cloud generates this ID for each request. Use this ID to troubleshoot issues.
	//
	// example:
	//
	// 6CF2815C-C8C7-4A01-B52E-FF6E24F53492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageBatchModerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBody) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImageBatchModerationResponseBody) GetData() *ImageBatchModerationResponseBodyData {
	return s.Data
}

func (s *ImageBatchModerationResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ImageBatchModerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageBatchModerationResponseBody) SetCode(v int32) *ImageBatchModerationResponseBody {
	s.Code = &v
	return s
}

func (s *ImageBatchModerationResponseBody) SetData(v *ImageBatchModerationResponseBodyData) *ImageBatchModerationResponseBody {
	s.Data = v
	return s
}

func (s *ImageBatchModerationResponseBody) SetMsg(v string) *ImageBatchModerationResponseBody {
	s.Msg = &v
	return s
}

func (s *ImageBatchModerationResponseBody) SetRequestId(v string) *ImageBatchModerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageBatchModerationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageBatchModerationResponseBodyData struct {
	// The data ID of the moderated object.
	//
	// example:
	//
	// 26769ada6e264e7ba9aa048241e12be9
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The ID of the manual review task.
	//
	// example:
	//
	// xxxxx-xxxxx
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// An array of results for the image moderation. The results contain parameters such as threat labels and confidence scores.
	Result []*ImageBatchModerationResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The detailed moderation results for each detection service. This is an array.
	Results []*ImageBatchModerationResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The risk level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ImageBatchModerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *ImageBatchModerationResponseBodyData) GetManualTaskId() *string {
	return s.ManualTaskId
}

func (s *ImageBatchModerationResponseBodyData) GetResult() []*ImageBatchModerationResponseBodyDataResult {
	return s.Result
}

func (s *ImageBatchModerationResponseBodyData) GetResults() []*ImageBatchModerationResponseBodyDataResults {
	return s.Results
}

func (s *ImageBatchModerationResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ImageBatchModerationResponseBodyData) SetDataId(v string) *ImageBatchModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ImageBatchModerationResponseBodyData) SetManualTaskId(v string) *ImageBatchModerationResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *ImageBatchModerationResponseBodyData) SetResult(v []*ImageBatchModerationResponseBodyDataResult) *ImageBatchModerationResponseBodyData {
	s.Result = v
	return s
}

func (s *ImageBatchModerationResponseBodyData) SetResults(v []*ImageBatchModerationResponseBodyDataResults) *ImageBatchModerationResponseBodyData {
	s.Results = v
	return s
}

func (s *ImageBatchModerationResponseBodyData) SetRiskLevel(v string) *ImageBatchModerationResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *ImageBatchModerationResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImageBatchModerationResponseBodyDataResult struct {
	// The confidence score. The value ranges from 0 to 100, with two decimal places. Some labels do not have a confidence score.
	//
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The description.
	//
	// example:
	//
	// 未检测出风险
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label returned after the image content moderation. An image may have multiple labels and scores.
	//
	// example:
	//
	// violent_explosion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageBatchModerationResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *ImageBatchModerationResponseBodyDataResult) GetLabel() *string {
	return s.Label
}

func (s *ImageBatchModerationResponseBodyDataResult) SetConfidence(v float32) *ImageBatchModerationResponseBodyDataResult {
	s.Confidence = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResult) SetDescription(v string) *ImageBatchModerationResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResult) SetLabel(v string) *ImageBatchModerationResponseBodyDataResult {
	s.Label = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type ImageBatchModerationResponseBodyDataResults struct {
	// Additional reference information for the image.
	Ext *ImageBatchModerationResponseBodyDataResultsExt `json:"Ext,omitempty" xml:"Ext,omitempty" type:"Struct"`
	// The results of the image detection, including threat labels and confidence scores. This is an array.
	Result []*ImageBatchModerationResponseBodyDataResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The risk level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The detection service supported by Image Moderation Pro.
	//
	// example:
	//
	// baselineCheck
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResults) GetExt() *ImageBatchModerationResponseBodyDataResultsExt {
	return s.Ext
}

func (s *ImageBatchModerationResponseBodyDataResults) GetResult() []*ImageBatchModerationResponseBodyDataResultsResult {
	return s.Result
}

func (s *ImageBatchModerationResponseBodyDataResults) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ImageBatchModerationResponseBodyDataResults) GetService() *string {
	return s.Service
}

func (s *ImageBatchModerationResponseBodyDataResults) SetExt(v *ImageBatchModerationResponseBodyDataResultsExt) *ImageBatchModerationResponseBodyDataResults {
	s.Ext = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResults) SetResult(v []*ImageBatchModerationResponseBodyDataResultsResult) *ImageBatchModerationResponseBodyDataResults {
	s.Result = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResults) SetRiskLevel(v string) *ImageBatchModerationResponseBodyDataResults {
	s.RiskLevel = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResults) SetService(v string) *ImageBatchModerationResponseBodyDataResults {
	s.Service = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResults) Validate() error {
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImageBatchModerationResponseBodyDataResultsExt struct {
	// A list of hits in custom image libraries.
	CustomImage []*ImageBatchModerationResponseBodyDataResultsExtCustomImage `json:"CustomImage,omitempty" xml:"CustomImage,omitempty" type:"Repeated"`
	// Logo information.
	LogoData *ImageBatchModerationResponseBodyDataResultsExtLogoData `json:"LogoData,omitempty" xml:"LogoData,omitempty" type:"Struct"`
	// A list of public figures.
	PublicFigure []*ImageBatchModerationResponseBodyDataResultsExtPublicFigure `json:"PublicFigure,omitempty" xml:"PublicFigure,omitempty" type:"Repeated"`
	// The text detected in the image.
	TextInImage *ImageBatchModerationResponseBodyDataResultsExtTextInImage `json:"TextInImage,omitempty" xml:"TextInImage,omitempty" type:"Struct"`
}

func (s ImageBatchModerationResponseBodyDataResultsExt) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExt) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExt) GetCustomImage() []*ImageBatchModerationResponseBodyDataResultsExtCustomImage {
	return s.CustomImage
}

func (s *ImageBatchModerationResponseBodyDataResultsExt) GetLogoData() *ImageBatchModerationResponseBodyDataResultsExtLogoData {
	return s.LogoData
}

func (s *ImageBatchModerationResponseBodyDataResultsExt) GetPublicFigure() []*ImageBatchModerationResponseBodyDataResultsExtPublicFigure {
	return s.PublicFigure
}

func (s *ImageBatchModerationResponseBodyDataResultsExt) GetTextInImage() *ImageBatchModerationResponseBodyDataResultsExtTextInImage {
	return s.TextInImage
}

func (s *ImageBatchModerationResponseBodyDataResultsExt) SetCustomImage(v []*ImageBatchModerationResponseBodyDataResultsExtCustomImage) *ImageBatchModerationResponseBodyDataResultsExt {
	s.CustomImage = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExt) SetLogoData(v *ImageBatchModerationResponseBodyDataResultsExtLogoData) *ImageBatchModerationResponseBodyDataResultsExt {
	s.LogoData = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExt) SetPublicFigure(v []*ImageBatchModerationResponseBodyDataResultsExtPublicFigure) *ImageBatchModerationResponseBodyDataResultsExt {
	s.PublicFigure = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExt) SetTextInImage(v *ImageBatchModerationResponseBodyDataResultsExtTextInImage) *ImageBatchModerationResponseBodyDataResultsExt {
	s.TextInImage = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExt) Validate() error {
	if s.CustomImage != nil {
		for _, item := range s.CustomImage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LogoData != nil {
		if err := s.LogoData.Validate(); err != nil {
			return err
		}
	}
	if s.PublicFigure != nil {
		for _, item := range s.PublicFigure {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TextInImage != nil {
		if err := s.TextInImage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageBatchModerationResponseBodyDataResultsExtCustomImage struct {
	// The ID of the hit custom image.
	//
	// example:
	//
	// 1965304870002
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the custom library.
	//
	// example:
	//
	// 1965304870002
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The name of the hit custom image library.
	//
	// example:
	//
	// 白名单
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtCustomImage) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtCustomImage) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtCustomImage) GetImageId() *string {
	return s.ImageId
}

func (s *ImageBatchModerationResponseBodyDataResultsExtCustomImage) GetLibId() *string {
	return s.LibId
}

func (s *ImageBatchModerationResponseBodyDataResultsExtCustomImage) GetLibName() *string {
	return s.LibName
}

func (s *ImageBatchModerationResponseBodyDataResultsExtCustomImage) SetImageId(v string) *ImageBatchModerationResponseBodyDataResultsExtCustomImage {
	s.ImageId = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtCustomImage) SetLibId(v string) *ImageBatchModerationResponseBodyDataResultsExtCustomImage {
	s.LibId = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtCustomImage) SetLibName(v string) *ImageBatchModerationResponseBodyDataResultsExtCustomImage {
	s.LibName = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtCustomImage) Validate() error {
	return dara.Validate(s)
}

type ImageBatchModerationResponseBodyDataResultsExtLogoData struct {
	// The location of the recognized object.
	Location *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// Identity information.
	Logo []*ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo `json:"Logo,omitempty" xml:"Logo,omitempty" type:"Repeated"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtLogoData) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtLogoData) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoData) GetLocation() *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation {
	return s.Location
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoData) GetLogo() []*ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo {
	return s.Logo
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoData) SetLocation(v *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) *ImageBatchModerationResponseBodyDataResultsExtLogoData {
	s.Location = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoData) SetLogo(v []*ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) *ImageBatchModerationResponseBodyDataResultsExtLogoData {
	s.Logo = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoData) Validate() error {
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	if s.Logo != nil {
		for _, item := range s.Logo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation struct {
	// The height of the logo area, in pixels.
	//
	// example:
	//
	// 440
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the logo area, in pixels.
	//
	// example:
	//
	// 330
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The x-coordinate of the upper-left corner of the area, in pixels. The origin (0,0) is the upper-left corner of the image.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate of the upper-left corner of the area, in pixels. The origin (0,0) is the upper-left corner of the image.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) GetH() *int32 {
	return s.H
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) GetW() *int32 {
	return s.W
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) GetX() *int32 {
	return s.X
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) GetY() *int32 {
	return s.Y
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) SetH(v int32) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation {
	s.H = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) SetW(v int32) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation {
	s.W = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) SetX(v int32) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation {
	s.X = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) SetY(v int32) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation {
	s.Y = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) Validate() error {
	return dara.Validate(s)
}

type ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo struct {
	// The confidence score. The value ranges from 0 to 100, with two decimal places.
	//
	// example:
	//
	// 99.1
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The category of the logo.
	//
	// example:
	//
	// logo_sns
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The name of the logo.
	//
	// example:
	//
	// 阿里云
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) GetLabel() *string {
	return s.Label
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) GetName() *string {
	return s.Name
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) SetConfidence(v float32) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo {
	s.Confidence = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) SetLabel(v string) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo {
	s.Label = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) SetName(v string) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo {
	s.Name = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) Validate() error {
	return dara.Validate(s)
}

type ImageBatchModerationResponseBodyDataResultsExtPublicFigure struct {
	// The ID of the recognized public figure.
	//
	// example:
	//
	// 12324222
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	// The name of the recognized public figure.
	//
	// example:
	//
	// xxxxx
	FigureName *string `json:"FigureName,omitempty" xml:"FigureName,omitempty"`
	// The location of the recognized object.
	Location []*ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtPublicFigure) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtPublicFigure) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigure) GetFigureId() *string {
	return s.FigureId
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigure) GetFigureName() *string {
	return s.FigureName
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigure) GetLocation() []*ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation {
	return s.Location
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigure) SetFigureId(v string) *ImageBatchModerationResponseBodyDataResultsExtPublicFigure {
	s.FigureId = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigure) SetFigureName(v string) *ImageBatchModerationResponseBodyDataResultsExtPublicFigure {
	s.FigureName = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigure) SetLocation(v []*ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) *ImageBatchModerationResponseBodyDataResultsExtPublicFigure {
	s.Location = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigure) Validate() error {
	if s.Location != nil {
		for _, item := range s.Location {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation struct {
	// The height of the area, in pixels.
	//
	// example:
	//
	// 440
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the area, in pixels.
	//
	// example:
	//
	// 330
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The x-coordinate of the upper-left corner of the area, in pixels. The origin (0,0) is the upper-left corner of the image.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate of the upper-left corner of the area, in pixels. The origin (0,0) is the upper-left corner of the image.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) GetH() *int32 {
	return s.H
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) GetW() *int32 {
	return s.W
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) GetX() *int32 {
	return s.X
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) GetY() *int32 {
	return s.Y
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) SetH(v int32) *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation {
	s.H = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) SetW(v int32) *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation {
	s.W = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) SetX(v int32) *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation {
	s.X = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) SetY(v int32) *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation {
	s.Y = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) Validate() error {
	return dara.Validate(s)
}

type ImageBatchModerationResponseBodyDataResultsExtTextInImage struct {
	// If a custom text library is hit, the ID and name of the library, and the hit keywords are returned.
	CustomText []*ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText `json:"CustomText,omitempty" xml:"CustomText,omitempty" type:"Repeated"`
	// The information for each line of text recognized in the image.
	OcrResult []*ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult `json:"OcrResult,omitempty" xml:"OcrResult,omitempty" type:"Repeated"`
	// The detected risk keywords.
	RiskWord []*string `json:"RiskWord,omitempty" xml:"RiskWord,omitempty" type:"Repeated"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImage) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImage) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImage) GetCustomText() []*ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText {
	return s.CustomText
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImage) GetOcrResult() []*ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult {
	return s.OcrResult
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImage) GetRiskWord() []*string {
	return s.RiskWord
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImage) SetCustomText(v []*ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) *ImageBatchModerationResponseBodyDataResultsExtTextInImage {
	s.CustomText = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImage) SetOcrResult(v []*ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult) *ImageBatchModerationResponseBodyDataResultsExtTextInImage {
	s.OcrResult = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImage) SetRiskWord(v []*string) *ImageBatchModerationResponseBodyDataResultsExtTextInImage {
	s.RiskWord = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImage) Validate() error {
	if s.CustomText != nil {
		for _, item := range s.CustomText {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OcrResult != nil {
		for _, item := range s.OcrResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText struct {
	// The custom keywords. Separate multiple keywords with a comma.
	//
	// example:
	//
	// 自定义词1,自定义词2
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// The ID of the custom library.
	//
	// example:
	//
	// 123456
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The name of the custom library.
	//
	// example:
	//
	// 自定义库1
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) GetKeyWords() *string {
	return s.KeyWords
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) GetLibId() *string {
	return s.LibId
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) GetLibName() *string {
	return s.LibName
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) SetKeyWords(v string) *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText {
	s.KeyWords = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) SetLibId(v string) *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText {
	s.LibId = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) SetLibName(v string) *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText {
	s.LibName = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) Validate() error {
	return dara.Validate(s)
}

type ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult struct {
	// The coordinates of the text line.
	Location *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// The text.
	//
	// example:
	//
	// abcd
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult) GetLocation() *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation {
	return s.Location
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult) GetText() *string {
	return s.Text
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult) SetLocation(v *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult {
	s.Location = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult) SetText(v string) *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult {
	s.Text = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult) Validate() error {
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation struct {
	// The height of the text area, in pixels.
	//
	// example:
	//
	// 33
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the text area, in pixels.
	//
	// example:
	//
	// 44
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The x-coordinate of the upper-left corner of the text area, in pixels. The origin (0,0) is the upper-left corner of the image.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate of the upper-left corner of the text area, in pixels. The origin (0,0) is the upper-left corner of the image.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) GetH() *int32 {
	return s.H
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) GetW() *int32 {
	return s.W
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) GetX() *int32 {
	return s.X
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) GetY() *int32 {
	return s.Y
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) SetH(v int32) *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation {
	s.H = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) SetW(v int32) *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation {
	s.W = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) SetX(v int32) *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation {
	s.X = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) SetY(v int32) *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation {
	s.Y = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) Validate() error {
	return dara.Validate(s)
}

type ImageBatchModerationResponseBodyDataResultsResult struct {
	// The confidence score. The value ranges from 0 to 100, with two decimal places. Some labels do not have a confidence score.
	//
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The description.
	//
	// example:
	//
	// 未检测出风险
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label returned after the image content moderation. An image may have multiple labels and scores.
	//
	// example:
	//
	// violent_explosion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsResult) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsResult) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageBatchModerationResponseBodyDataResultsResult) GetDescription() *string {
	return s.Description
}

func (s *ImageBatchModerationResponseBodyDataResultsResult) GetLabel() *string {
	return s.Label
}

func (s *ImageBatchModerationResponseBodyDataResultsResult) SetConfidence(v float32) *ImageBatchModerationResponseBodyDataResultsResult {
	s.Confidence = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsResult) SetDescription(v string) *ImageBatchModerationResponseBodyDataResultsResult {
	s.Description = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsResult) SetLabel(v string) *ImageBatchModerationResponseBodyDataResultsResult {
	s.Label = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsResult) Validate() error {
	return dara.Validate(s)
}
