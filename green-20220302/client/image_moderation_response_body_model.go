// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageModerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImageModerationResponseBody
	GetCode() *int32
	SetData(v *ImageModerationResponseBodyData) *ImageModerationResponseBody
	GetData() *ImageModerationResponseBodyData
	SetMsg(v string) *ImageModerationResponseBody
	GetMsg() *string
	SetRequestId(v string) *ImageModerationResponseBody
	GetRequestId() *string
}

type ImageModerationResponseBody struct {
	// The return code. A value of 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The results of the image content moderation.
	Data *ImageModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned for the request.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID of the request. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 6CF2815C-C8C7-4A01-B52E-FF6E24F53492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageModerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBody) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImageModerationResponseBody) GetData() *ImageModerationResponseBodyData {
	return s.Data
}

func (s *ImageModerationResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ImageModerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageModerationResponseBody) SetCode(v int32) *ImageModerationResponseBody {
	s.Code = &v
	return s
}

func (s *ImageModerationResponseBody) SetData(v *ImageModerationResponseBodyData) *ImageModerationResponseBody {
	s.Data = v
	return s
}

func (s *ImageModerationResponseBody) SetMsg(v string) *ImageModerationResponseBody {
	s.Msg = &v
	return s
}

func (s *ImageModerationResponseBody) SetRequestId(v string) *ImageModerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageModerationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageModerationResponseBodyData struct {
	// The AccountId specified in the request.
	//
	// example:
	//
	// testaccountid123
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The data ID of the detected object.
	//
	// > If you specify the dataId parameter in the request, the corresponding dataId is returned.
	//
	// example:
	//
	// fb5ffab1-993b-449f-b8d6-b97d5e3331f2
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Auxiliary reference information for the image.
	Ext *ImageModerationResponseBodyDataExt `json:"Ext,omitempty" xml:"Ext,omitempty" type:"Struct"`
	// The ID of the manual review task.
	//
	// example:
	//
	// xxxxx-xxxxx
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// The results of the image moderation, including the threat labels and confidence levels. The value is an array.
	Result []*ImageModerationResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The threat level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ImageModerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *ImageModerationResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *ImageModerationResponseBodyData) GetExt() *ImageModerationResponseBodyDataExt {
	return s.Ext
}

func (s *ImageModerationResponseBodyData) GetManualTaskId() *string {
	return s.ManualTaskId
}

func (s *ImageModerationResponseBodyData) GetResult() []*ImageModerationResponseBodyDataResult {
	return s.Result
}

func (s *ImageModerationResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ImageModerationResponseBodyData) SetAccountId(v string) *ImageModerationResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ImageModerationResponseBodyData) SetDataId(v string) *ImageModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ImageModerationResponseBodyData) SetExt(v *ImageModerationResponseBodyDataExt) *ImageModerationResponseBodyData {
	s.Ext = v
	return s
}

func (s *ImageModerationResponseBodyData) SetManualTaskId(v string) *ImageModerationResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *ImageModerationResponseBodyData) SetResult(v []*ImageModerationResponseBodyDataResult) *ImageModerationResponseBodyData {
	s.Result = v
	return s
}

func (s *ImageModerationResponseBodyData) SetRiskLevel(v string) *ImageModerationResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *ImageModerationResponseBodyData) Validate() error {
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

type ImageModerationResponseBodyDataExt struct {
	// The detection information for the implicit AIGC identity in the image.
	AigcData *ImageModerationResponseBodyDataExtAigcData `json:"AigcData,omitempty" xml:"AigcData,omitempty" type:"Struct"`
	// A list of hits from the custom image library.
	CustomImage []*ImageModerationResponseBodyDataExtCustomImage `json:"CustomImage,omitempty" xml:"CustomImage,omitempty" type:"Repeated"`
	// The facial attribute detection results.
	FaceData []*ImageModerationResponseBodyDataExtFaceData `json:"FaceData,omitempty" xml:"FaceData,omitempty" type:"Repeated"`
	// The identity information.
	LogoData []*ImageModerationResponseBodyDataExtLogoData `json:"LogoData,omitempty" xml:"LogoData,omitempty" type:"Repeated"`
	// The results of optical character recognition (OCR).
	OcrResult []*ImageModerationResponseBodyDataExtOcrResult `json:"OcrResult,omitempty" xml:"OcrResult,omitempty" type:"Repeated"`
	// A list of public figures.
	PublicFigure []*ImageModerationResponseBodyDataExtPublicFigure `json:"PublicFigure,omitempty" xml:"PublicFigure,omitempty" type:"Repeated"`
	// The results of image object recognition.
	Recognition []*ImageModerationResponseBodyDataExtRecognition `json:"Recognition,omitempty" xml:"Recognition,omitempty" type:"Repeated"`
	// The text information that is hit in the image.
	TextInImage *ImageModerationResponseBodyDataExtTextInImage `json:"TextInImage,omitempty" xml:"TextInImage,omitempty" type:"Struct"`
	// The output content.
	VlContent *ImageModerationResponseBodyDataExtVlContent `json:"VlContent,omitempty" xml:"VlContent,omitempty" type:"Struct"`
}

func (s ImageModerationResponseBodyDataExt) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExt) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExt) GetAigcData() *ImageModerationResponseBodyDataExtAigcData {
	return s.AigcData
}

func (s *ImageModerationResponseBodyDataExt) GetCustomImage() []*ImageModerationResponseBodyDataExtCustomImage {
	return s.CustomImage
}

func (s *ImageModerationResponseBodyDataExt) GetFaceData() []*ImageModerationResponseBodyDataExtFaceData {
	return s.FaceData
}

func (s *ImageModerationResponseBodyDataExt) GetLogoData() []*ImageModerationResponseBodyDataExtLogoData {
	return s.LogoData
}

func (s *ImageModerationResponseBodyDataExt) GetOcrResult() []*ImageModerationResponseBodyDataExtOcrResult {
	return s.OcrResult
}

func (s *ImageModerationResponseBodyDataExt) GetPublicFigure() []*ImageModerationResponseBodyDataExtPublicFigure {
	return s.PublicFigure
}

func (s *ImageModerationResponseBodyDataExt) GetRecognition() []*ImageModerationResponseBodyDataExtRecognition {
	return s.Recognition
}

func (s *ImageModerationResponseBodyDataExt) GetTextInImage() *ImageModerationResponseBodyDataExtTextInImage {
	return s.TextInImage
}

func (s *ImageModerationResponseBodyDataExt) GetVlContent() *ImageModerationResponseBodyDataExtVlContent {
	return s.VlContent
}

func (s *ImageModerationResponseBodyDataExt) SetAigcData(v *ImageModerationResponseBodyDataExtAigcData) *ImageModerationResponseBodyDataExt {
	s.AigcData = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetCustomImage(v []*ImageModerationResponseBodyDataExtCustomImage) *ImageModerationResponseBodyDataExt {
	s.CustomImage = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetFaceData(v []*ImageModerationResponseBodyDataExtFaceData) *ImageModerationResponseBodyDataExt {
	s.FaceData = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetLogoData(v []*ImageModerationResponseBodyDataExtLogoData) *ImageModerationResponseBodyDataExt {
	s.LogoData = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetOcrResult(v []*ImageModerationResponseBodyDataExtOcrResult) *ImageModerationResponseBodyDataExt {
	s.OcrResult = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetPublicFigure(v []*ImageModerationResponseBodyDataExtPublicFigure) *ImageModerationResponseBodyDataExt {
	s.PublicFigure = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetRecognition(v []*ImageModerationResponseBodyDataExtRecognition) *ImageModerationResponseBodyDataExt {
	s.Recognition = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetTextInImage(v *ImageModerationResponseBodyDataExtTextInImage) *ImageModerationResponseBodyDataExt {
	s.TextInImage = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetVlContent(v *ImageModerationResponseBodyDataExtVlContent) *ImageModerationResponseBodyDataExt {
	s.VlContent = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) Validate() error {
	if s.AigcData != nil {
		if err := s.AigcData.Validate(); err != nil {
			return err
		}
	}
	if s.CustomImage != nil {
		for _, item := range s.CustomImage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FaceData != nil {
		for _, item := range s.FaceData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LogoData != nil {
		for _, item := range s.LogoData {
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
	if s.PublicFigure != nil {
		for _, item := range s.PublicFigure {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Recognition != nil {
		for _, item := range s.Recognition {
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
	if s.VlContent != nil {
		if err := s.VlContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageModerationResponseBodyDataExtAigcData struct {
	// The detection information for the implicit AIGC identity.
	AIGC *ImageModerationResponseBodyDataExtAigcDataAIGC `json:"AIGC,omitempty" xml:"AIGC,omitempty" type:"Struct"`
}

func (s ImageModerationResponseBodyDataExtAigcData) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtAigcData) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtAigcData) GetAIGC() *ImageModerationResponseBodyDataExtAigcDataAIGC {
	return s.AIGC
}

func (s *ImageModerationResponseBodyDataExtAigcData) SetAIGC(v *ImageModerationResponseBodyDataExtAigcDataAIGC) *ImageModerationResponseBodyDataExtAigcData {
	s.AIGC = v
	return s
}

func (s *ImageModerationResponseBodyDataExtAigcData) Validate() error {
	if s.AIGC != nil {
		if err := s.AIGC.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageModerationResponseBodyDataExtAigcDataAIGC struct {
	// The code or name of the service provider, which identifies the content producer.
	//
	// example:
	//
	// 001191******M000100Y43
	ContentProducer *string `json:"ContentProducer,omitempty" xml:"ContentProducer,omitempty"`
	// The name, ID, or code of the propagation platform. For services that provide AI-generated content, this can be the same as the value of ContentProducer.
	//
	// example:
	//
	// 001191******M000100Y43
	ContentPropagator *string `json:"ContentPropagator,omitempty" xml:"ContentPropagator,omitempty"`
	// Indicates whether the content is generated by artificial intelligence (AI). Valid values:
	//
	// - 1: The content is generated by AI.
	//
	// - 2: (For distribution platforms only) The content may be generated by AI.
	//
	// - 3: (For distribution platforms only) The content is suspected to be generated by AI.
	//
	// example:
	//
	// 1
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The content production ID. This is a unique ID used on the production platform to trace the source of synthesized content.
	//
	// example:
	//
	// 123******456
	ProduceID *string `json:"ProduceID,omitempty" xml:"ProduceID,omitempty"`
	// The content propagation ID. This is a unique ID that the propagation platform assigns to the distributed synthetic content.
	//
	// example:
	//
	// 123******456
	PropagateID *string `json:"PropagateID,omitempty" xml:"PropagateID,omitempty"`
	// A reserved field.
	//
	// This field can store information that the generative service provider uses for security protection to ensure the integrity of content and identities. A hashing mechanism based on ContentProducer and ProduceID can be used to securely store and verify key information.
	//
	// example:
	//
	// d41d**********427e
	ReservedCode1 *string `json:"ReservedCode1,omitempty" xml:"ReservedCode1,omitempty"`
	// A reserved field.
	//
	// This field can be used by content distribution service providers for security protection to ensure the integrity of content and identities. A hashing mechanism based on ContentProducer and ProduceID can be used to securely store and verify key information.
	//
	// example:
	//
	// d41d**********427e
	ReservedCode2 *string `json:"ReservedCode2,omitempty" xml:"ReservedCode2,omitempty"`
}

func (s ImageModerationResponseBodyDataExtAigcDataAIGC) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtAigcDataAIGC) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) GetContentProducer() *string {
	return s.ContentProducer
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) GetContentPropagator() *string {
	return s.ContentPropagator
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) GetLabel() *string {
	return s.Label
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) GetProduceID() *string {
	return s.ProduceID
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) GetPropagateID() *string {
	return s.PropagateID
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) GetReservedCode1() *string {
	return s.ReservedCode1
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) GetReservedCode2() *string {
	return s.ReservedCode2
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) SetContentProducer(v string) *ImageModerationResponseBodyDataExtAigcDataAIGC {
	s.ContentProducer = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) SetContentPropagator(v string) *ImageModerationResponseBodyDataExtAigcDataAIGC {
	s.ContentPropagator = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) SetLabel(v string) *ImageModerationResponseBodyDataExtAigcDataAIGC {
	s.Label = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) SetProduceID(v string) *ImageModerationResponseBodyDataExtAigcDataAIGC {
	s.ProduceID = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) SetPropagateID(v string) *ImageModerationResponseBodyDataExtAigcDataAIGC {
	s.PropagateID = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) SetReservedCode1(v string) *ImageModerationResponseBodyDataExtAigcDataAIGC {
	s.ReservedCode1 = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) SetReservedCode2(v string) *ImageModerationResponseBodyDataExtAigcDataAIGC {
	s.ReservedCode2 = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtAigcDataAIGC) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtCustomImage struct {
	// The ID of the hit custom image.
	//
	// example:
	//
	// 123456
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the hit custom image library.
	//
	// example:
	//
	// 图库123
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The name of the hit custom image library.
	//
	// example:
	//
	// 图库123
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s ImageModerationResponseBodyDataExtCustomImage) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtCustomImage) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtCustomImage) GetImageId() *string {
	return s.ImageId
}

func (s *ImageModerationResponseBodyDataExtCustomImage) GetLibId() *string {
	return s.LibId
}

func (s *ImageModerationResponseBodyDataExtCustomImage) GetLibName() *string {
	return s.LibName
}

func (s *ImageModerationResponseBodyDataExtCustomImage) SetImageId(v string) *ImageModerationResponseBodyDataExtCustomImage {
	s.ImageId = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtCustomImage) SetLibId(v string) *ImageModerationResponseBodyDataExtCustomImage {
	s.LibId = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtCustomImage) SetLibName(v string) *ImageModerationResponseBodyDataExtCustomImage {
	s.LibName = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtCustomImage) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtFaceData struct {
	// The detected age.
	//
	// example:
	//
	// 18
	Age *int32 `json:"Age,omitempty" xml:"Age,omitempty"`
	// The detection result for bangs.
	Bang *ImageModerationResponseBodyDataExtFaceDataBang `json:"Bang,omitempty" xml:"Bang,omitempty" type:"Struct"`
	// The gender detection result.
	Gender *ImageModerationResponseBodyDataExtFaceDataGender `json:"Gender,omitempty" xml:"Gender,omitempty" type:"Struct"`
	// Indicates whether the person is wearing glasses. Valid values:
	//
	// - None: The person is not wearing glasses.
	//
	// - Common: The person is wearing regular glasses.
	//
	// - Sunglass: The person is wearing sunglasses.
	//
	// example:
	//
	// Common
	Glasses *string `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	// The hairstyle detection result.
	Hairstyle *ImageModerationResponseBodyDataExtFaceDataHairstyle `json:"Hairstyle,omitempty" xml:"Hairstyle,omitempty" type:"Struct"`
	// The result of hat detection.
	Hat *ImageModerationResponseBodyDataExtFaceDataHat `json:"Hat,omitempty" xml:"Hat,omitempty" type:"Struct"`
	// The location of the face.
	Location *ImageModerationResponseBodyDataExtFaceDataLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// The result of mask detection.
	Mask *ImageModerationResponseBodyDataExtFaceDataMask `json:"Mask,omitempty" xml:"Mask,omitempty" type:"Struct"`
	// The result of mustache detection.
	Mustache *ImageModerationResponseBodyDataExtFaceDataMustache `json:"Mustache,omitempty" xml:"Mustache,omitempty" type:"Struct"`
	// The quality of the face image.
	Quality *ImageModerationResponseBodyDataExtFaceDataQuality `json:"Quality,omitempty" xml:"Quality,omitempty" type:"Struct"`
	// The degree of the smile. The value ranges from 0 to 100. A higher score indicates a wider smile.
	//
	// example:
	//
	// 85.88
	Smile *float32 `json:"Smile,omitempty" xml:"Smile,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceData) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceData) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceData) GetAge() *int32 {
	return s.Age
}

func (s *ImageModerationResponseBodyDataExtFaceData) GetBang() *ImageModerationResponseBodyDataExtFaceDataBang {
	return s.Bang
}

func (s *ImageModerationResponseBodyDataExtFaceData) GetGender() *ImageModerationResponseBodyDataExtFaceDataGender {
	return s.Gender
}

func (s *ImageModerationResponseBodyDataExtFaceData) GetGlasses() *string {
	return s.Glasses
}

func (s *ImageModerationResponseBodyDataExtFaceData) GetHairstyle() *ImageModerationResponseBodyDataExtFaceDataHairstyle {
	return s.Hairstyle
}

func (s *ImageModerationResponseBodyDataExtFaceData) GetHat() *ImageModerationResponseBodyDataExtFaceDataHat {
	return s.Hat
}

func (s *ImageModerationResponseBodyDataExtFaceData) GetLocation() *ImageModerationResponseBodyDataExtFaceDataLocation {
	return s.Location
}

func (s *ImageModerationResponseBodyDataExtFaceData) GetMask() *ImageModerationResponseBodyDataExtFaceDataMask {
	return s.Mask
}

func (s *ImageModerationResponseBodyDataExtFaceData) GetMustache() *ImageModerationResponseBodyDataExtFaceDataMustache {
	return s.Mustache
}

func (s *ImageModerationResponseBodyDataExtFaceData) GetQuality() *ImageModerationResponseBodyDataExtFaceDataQuality {
	return s.Quality
}

func (s *ImageModerationResponseBodyDataExtFaceData) GetSmile() *float32 {
	return s.Smile
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetAge(v int32) *ImageModerationResponseBodyDataExtFaceData {
	s.Age = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetBang(v *ImageModerationResponseBodyDataExtFaceDataBang) *ImageModerationResponseBodyDataExtFaceData {
	s.Bang = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetGender(v *ImageModerationResponseBodyDataExtFaceDataGender) *ImageModerationResponseBodyDataExtFaceData {
	s.Gender = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetGlasses(v string) *ImageModerationResponseBodyDataExtFaceData {
	s.Glasses = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetHairstyle(v *ImageModerationResponseBodyDataExtFaceDataHairstyle) *ImageModerationResponseBodyDataExtFaceData {
	s.Hairstyle = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetHat(v *ImageModerationResponseBodyDataExtFaceDataHat) *ImageModerationResponseBodyDataExtFaceData {
	s.Hat = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetLocation(v *ImageModerationResponseBodyDataExtFaceDataLocation) *ImageModerationResponseBodyDataExtFaceData {
	s.Location = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetMask(v *ImageModerationResponseBodyDataExtFaceDataMask) *ImageModerationResponseBodyDataExtFaceData {
	s.Mask = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetMustache(v *ImageModerationResponseBodyDataExtFaceDataMustache) *ImageModerationResponseBodyDataExtFaceData {
	s.Mustache = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetQuality(v *ImageModerationResponseBodyDataExtFaceDataQuality) *ImageModerationResponseBodyDataExtFaceData {
	s.Quality = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetSmile(v float32) *ImageModerationResponseBodyDataExtFaceData {
	s.Smile = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) Validate() error {
	if s.Bang != nil {
		if err := s.Bang.Validate(); err != nil {
			return err
		}
	}
	if s.Gender != nil {
		if err := s.Gender.Validate(); err != nil {
			return err
		}
	}
	if s.Hairstyle != nil {
		if err := s.Hairstyle.Validate(); err != nil {
			return err
		}
	}
	if s.Hat != nil {
		if err := s.Hat.Validate(); err != nil {
			return err
		}
	}
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	if s.Mask != nil {
		if err := s.Mask.Validate(); err != nil {
			return err
		}
	}
	if s.Mustache != nil {
		if err := s.Mustache.Validate(); err != nil {
			return err
		}
	}
	if s.Quality != nil {
		if err := s.Quality.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageModerationResponseBodyDataExtFaceDataBang struct {
	// The confidence level of the bangs detection. The value ranges from 0 to 100. A higher value indicates a more reliable result.
	//
	// example:
	//
	// 81.88
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The detection result for bangs. Valid values:
	//
	// - Has: The person has bangs.
	//
	// - None: The person does not have bangs.
	//
	// example:
	//
	// Has
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataBang) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataBang) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataBang) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageModerationResponseBodyDataExtFaceDataBang) GetValue() *string {
	return s.Value
}

func (s *ImageModerationResponseBodyDataExtFaceDataBang) SetConfidence(v float32) *ImageModerationResponseBodyDataExtFaceDataBang {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataBang) SetValue(v string) *ImageModerationResponseBodyDataExtFaceDataBang {
	s.Value = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataBang) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtFaceDataGender struct {
	// The confidence level of the gender detection. The value ranges from 0 to 100. A higher value indicates a more reliable result.
	//
	// example:
	//
	// 81.88
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The detected gender. Valid values:
	//
	// - Male: male
	//
	// - FeMale: female
	//
	// example:
	//
	// FeMale
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataGender) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataGender) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataGender) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageModerationResponseBodyDataExtFaceDataGender) GetValue() *string {
	return s.Value
}

func (s *ImageModerationResponseBodyDataExtFaceDataGender) SetConfidence(v float32) *ImageModerationResponseBodyDataExtFaceDataGender {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataGender) SetValue(v string) *ImageModerationResponseBodyDataExtFaceDataGender {
	s.Value = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataGender) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtFaceDataHairstyle struct {
	// The confidence level of the hairstyle detection. The value ranges from 0 to 100. A higher value indicates a more reliable result.
	//
	// example:
	//
	// 81.88
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The detected hairstyle. Valid values:
	//
	// - Bald: bald
	//
	// - Long: long hair
	//
	// - Short: short hair
	//
	// example:
	//
	// Short
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataHairstyle) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataHairstyle) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataHairstyle) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageModerationResponseBodyDataExtFaceDataHairstyle) GetValue() *string {
	return s.Value
}

func (s *ImageModerationResponseBodyDataExtFaceDataHairstyle) SetConfidence(v float32) *ImageModerationResponseBodyDataExtFaceDataHairstyle {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataHairstyle) SetValue(v string) *ImageModerationResponseBodyDataExtFaceDataHairstyle {
	s.Value = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataHairstyle) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtFaceDataHat struct {
	// The confidence level of the hat detection. The value ranges from 0 to 100. A higher value indicates a more reliable result.
	//
	// example:
	//
	// 88.88
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Indicates whether a hat is detected. Valid values:
	//
	// - Wear: A hat is worn.
	//
	// - None: No hat is worn.
	//
	// example:
	//
	// Wear
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataHat) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataHat) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataHat) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageModerationResponseBodyDataExtFaceDataHat) GetValue() *string {
	return s.Value
}

func (s *ImageModerationResponseBodyDataExtFaceDataHat) SetConfidence(v float32) *ImageModerationResponseBodyDataExtFaceDataHat {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataHat) SetValue(v string) *ImageModerationResponseBodyDataExtFaceDataHat {
	s.Value = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataHat) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtFaceDataLocation struct {
	// The height of the face area. Unit: pixel.
	//
	// example:
	//
	// 26
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the face area. Unit: pixel.
	//
	// example:
	//
	// 83
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance from the upper-left corner of the face area to the y-axis. The origin is the upper-left corner of the image. Unit: pixel.
	//
	// example:
	//
	// 41
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance from the upper-left corner of the face area to the x-axis. The origin is the upper-left corner of the image. Unit: pixel.
	//
	// example:
	//
	// 84
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataLocation) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataLocation) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataLocation) GetH() *int32 {
	return s.H
}

func (s *ImageModerationResponseBodyDataExtFaceDataLocation) GetW() *int32 {
	return s.W
}

func (s *ImageModerationResponseBodyDataExtFaceDataLocation) GetX() *int32 {
	return s.X
}

func (s *ImageModerationResponseBodyDataExtFaceDataLocation) GetY() *int32 {
	return s.Y
}

func (s *ImageModerationResponseBodyDataExtFaceDataLocation) SetH(v int32) *ImageModerationResponseBodyDataExtFaceDataLocation {
	s.H = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataLocation) SetW(v int32) *ImageModerationResponseBodyDataExtFaceDataLocation {
	s.W = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataLocation) SetX(v int32) *ImageModerationResponseBodyDataExtFaceDataLocation {
	s.X = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataLocation) SetY(v int32) *ImageModerationResponseBodyDataExtFaceDataLocation {
	s.Y = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataLocation) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtFaceDataMask struct {
	// The confidence level of the mask detection. The value ranges from 0 to 100. A higher value indicates a more reliable result.
	//
	// example:
	//
	// 99.99
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Indicates whether a mask is worn. Valid values:
	//
	// - Wear: A mask is worn.
	//
	// - None: No mask is worn.
	//
	// example:
	//
	// Wear
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataMask) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataMask) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataMask) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageModerationResponseBodyDataExtFaceDataMask) GetValue() *string {
	return s.Value
}

func (s *ImageModerationResponseBodyDataExtFaceDataMask) SetConfidence(v float32) *ImageModerationResponseBodyDataExtFaceDataMask {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataMask) SetValue(v string) *ImageModerationResponseBodyDataExtFaceDataMask {
	s.Value = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataMask) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtFaceDataMustache struct {
	// The confidence level of the mustache detection. The value ranges from 0 to 100. A higher value indicates a more reliable result.
	//
	// example:
	//
	// 99.99
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Indicates whether a mustache is present. Valid values:
	//
	// - Has: A mustache is present.
	//
	// - None: No mustache is present.
	//
	// example:
	//
	// Has
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataMustache) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataMustache) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataMustache) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageModerationResponseBodyDataExtFaceDataMustache) GetValue() *string {
	return s.Value
}

func (s *ImageModerationResponseBodyDataExtFaceDataMustache) SetConfidence(v float32) *ImageModerationResponseBodyDataExtFaceDataMustache {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataMustache) SetValue(v string) *ImageModerationResponseBodyDataExtFaceDataMustache {
	s.Value = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataMustache) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtFaceDataQuality struct {
	// The blurriness of the face image. The value ranges from 0 to 100. A higher score indicates a blurrier image.
	//
	// A value from 0 to 25 is recommended.
	//
	// example:
	//
	// 5.88
	Blur *float32 `json:"Blur,omitempty" xml:"Blur,omitempty"`
	// The integrity of the face. The value ranges from 0 to 100. A higher score indicates a more complete face.
	//
	// A value from 80 to 100 is recommended.
	//
	// example:
	//
	// 100.0
	Integrity *float32 `json:"Integrity,omitempty" xml:"Integrity,omitempty"`
	// The pitch angle of the face.
	//
	// A value from -30 to 30 is recommended.
	//
	// example:
	//
	// 5.88
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	// The roll angle of the face.
	//
	// A value from -30 to 30 is recommended.
	//
	// example:
	//
	// 5.18
	Roll *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	// The yaw angle of the face.
	//
	// A value from -30 to 30 is recommended.
	//
	// example:
	//
	// 5.18
	Yaw *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataQuality) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataQuality) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) GetBlur() *float32 {
	return s.Blur
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) GetIntegrity() *float32 {
	return s.Integrity
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) GetPitch() *float32 {
	return s.Pitch
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) GetRoll() *float32 {
	return s.Roll
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) GetYaw() *float32 {
	return s.Yaw
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) SetBlur(v float32) *ImageModerationResponseBodyDataExtFaceDataQuality {
	s.Blur = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) SetIntegrity(v float32) *ImageModerationResponseBodyDataExtFaceDataQuality {
	s.Integrity = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) SetPitch(v float32) *ImageModerationResponseBodyDataExtFaceDataQuality {
	s.Pitch = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) SetRoll(v float32) *ImageModerationResponseBodyDataExtFaceDataQuality {
	s.Roll = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) SetYaw(v float32) *ImageModerationResponseBodyDataExtFaceDataQuality {
	s.Yaw = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtLogoData struct {
	// The location of the logo.
	Location *ImageModerationResponseBodyDataExtLogoDataLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// The identity information.
	Logo []*ImageModerationResponseBodyDataExtLogoDataLogo `json:"Logo,omitempty" xml:"Logo,omitempty" type:"Repeated"`
}

func (s ImageModerationResponseBodyDataExtLogoData) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtLogoData) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtLogoData) GetLocation() *ImageModerationResponseBodyDataExtLogoDataLocation {
	return s.Location
}

func (s *ImageModerationResponseBodyDataExtLogoData) GetLogo() []*ImageModerationResponseBodyDataExtLogoDataLogo {
	return s.Logo
}

func (s *ImageModerationResponseBodyDataExtLogoData) SetLocation(v *ImageModerationResponseBodyDataExtLogoDataLocation) *ImageModerationResponseBodyDataExtLogoData {
	s.Location = v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoData) SetLogo(v []*ImageModerationResponseBodyDataExtLogoDataLogo) *ImageModerationResponseBodyDataExtLogoData {
	s.Logo = v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoData) Validate() error {
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

type ImageModerationResponseBodyDataExtLogoDataLocation struct {
	// The height of the detected area. Unit: pixel.
	//
	// example:
	//
	// 440
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the detected area. Unit: pixel.
	//
	// example:
	//
	// 330
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance from the upper-left corner of the detected area to the y-axis. The origin is the upper-left corner of the image. Unit: pixel.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance from the upper-left corner of the detected area to the x-axis. The origin is the upper-left corner of the image. Unit: pixel.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageModerationResponseBodyDataExtLogoDataLocation) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtLogoDataLocation) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtLogoDataLocation) GetH() *int32 {
	return s.H
}

func (s *ImageModerationResponseBodyDataExtLogoDataLocation) GetW() *int32 {
	return s.W
}

func (s *ImageModerationResponseBodyDataExtLogoDataLocation) GetX() *int32 {
	return s.X
}

func (s *ImageModerationResponseBodyDataExtLogoDataLocation) GetY() *int32 {
	return s.Y
}

func (s *ImageModerationResponseBodyDataExtLogoDataLocation) SetH(v int32) *ImageModerationResponseBodyDataExtLogoDataLocation {
	s.H = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoDataLocation) SetW(v int32) *ImageModerationResponseBodyDataExtLogoDataLocation {
	s.W = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoDataLocation) SetX(v int32) *ImageModerationResponseBodyDataExtLogoDataLocation {
	s.X = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoDataLocation) SetY(v int32) *ImageModerationResponseBodyDataExtLogoDataLocation {
	s.Y = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoDataLocation) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtLogoDataLogo struct {
	// The confidence score. The value ranges from 0 to 100, with two decimal places retained.
	//
	// example:
	//
	// 99.1
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The identity category.
	//
	// example:
	//
	// logo_sns
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The identity name.
	//
	// example:
	//
	// 钉钉
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ImageModerationResponseBodyDataExtLogoDataLogo) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtLogoDataLogo) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtLogoDataLogo) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageModerationResponseBodyDataExtLogoDataLogo) GetLabel() *string {
	return s.Label
}

func (s *ImageModerationResponseBodyDataExtLogoDataLogo) GetName() *string {
	return s.Name
}

func (s *ImageModerationResponseBodyDataExtLogoDataLogo) SetConfidence(v float32) *ImageModerationResponseBodyDataExtLogoDataLogo {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoDataLogo) SetLabel(v string) *ImageModerationResponseBodyDataExtLogoDataLogo {
	s.Label = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoDataLogo) SetName(v string) *ImageModerationResponseBodyDataExtLogoDataLogo {
	s.Name = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoDataLogo) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtOcrResult struct {
	// The coordinates of the text line.
	Location *ImageModerationResponseBodyDataExtOcrResultLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// A single line of recognized text.
	//
	// example:
	//
	// abcd
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ImageModerationResponseBodyDataExtOcrResult) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtOcrResult) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtOcrResult) GetLocation() *ImageModerationResponseBodyDataExtOcrResultLocation {
	return s.Location
}

func (s *ImageModerationResponseBodyDataExtOcrResult) GetText() *string {
	return s.Text
}

func (s *ImageModerationResponseBodyDataExtOcrResult) SetLocation(v *ImageModerationResponseBodyDataExtOcrResultLocation) *ImageModerationResponseBodyDataExtOcrResult {
	s.Location = v
	return s
}

func (s *ImageModerationResponseBodyDataExtOcrResult) SetText(v string) *ImageModerationResponseBodyDataExtOcrResult {
	s.Text = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtOcrResult) Validate() error {
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageModerationResponseBodyDataExtOcrResultLocation struct {
	// The height of the text area. Unit: pixel.
	//
	// example:
	//
	// 44
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the text area. Unit: pixel.
	//
	// example:
	//
	// 33
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance from the upper-left corner of the text area to the y-axis. The origin is the upper-left corner of the image. Unit: pixel.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance from the upper-left corner of the text area to the x-axis. The origin is the upper-left corner of the image. Unit: pixel.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageModerationResponseBodyDataExtOcrResultLocation) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtOcrResultLocation) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtOcrResultLocation) GetH() *int32 {
	return s.H
}

func (s *ImageModerationResponseBodyDataExtOcrResultLocation) GetW() *int32 {
	return s.W
}

func (s *ImageModerationResponseBodyDataExtOcrResultLocation) GetX() *int32 {
	return s.X
}

func (s *ImageModerationResponseBodyDataExtOcrResultLocation) GetY() *int32 {
	return s.Y
}

func (s *ImageModerationResponseBodyDataExtOcrResultLocation) SetH(v int32) *ImageModerationResponseBodyDataExtOcrResultLocation {
	s.H = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtOcrResultLocation) SetW(v int32) *ImageModerationResponseBodyDataExtOcrResultLocation {
	s.W = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtOcrResultLocation) SetX(v int32) *ImageModerationResponseBodyDataExtOcrResultLocation {
	s.X = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtOcrResultLocation) SetY(v int32) *ImageModerationResponseBodyDataExtOcrResultLocation {
	s.Y = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtOcrResultLocation) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtPublicFigure struct {
	// The ID of the detected public figure.
	//
	// example:
	//
	// xxx001
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	// The name of the detected public figure.
	//
	// example:
	//
	// yzazhzou
	FigureName *string `json:"FigureName,omitempty" xml:"FigureName,omitempty"`
	// The location of the identity.
	Location []*ImageModerationResponseBodyDataExtPublicFigureLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
}

func (s ImageModerationResponseBodyDataExtPublicFigure) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtPublicFigure) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtPublicFigure) GetFigureId() *string {
	return s.FigureId
}

func (s *ImageModerationResponseBodyDataExtPublicFigure) GetFigureName() *string {
	return s.FigureName
}

func (s *ImageModerationResponseBodyDataExtPublicFigure) GetLocation() []*ImageModerationResponseBodyDataExtPublicFigureLocation {
	return s.Location
}

func (s *ImageModerationResponseBodyDataExtPublicFigure) SetFigureId(v string) *ImageModerationResponseBodyDataExtPublicFigure {
	s.FigureId = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtPublicFigure) SetFigureName(v string) *ImageModerationResponseBodyDataExtPublicFigure {
	s.FigureName = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtPublicFigure) SetLocation(v []*ImageModerationResponseBodyDataExtPublicFigureLocation) *ImageModerationResponseBodyDataExtPublicFigure {
	s.Location = v
	return s
}

func (s *ImageModerationResponseBodyDataExtPublicFigure) Validate() error {
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

type ImageModerationResponseBodyDataExtPublicFigureLocation struct {
	// The height of the detected area. Unit: pixel.
	//
	// example:
	//
	// 440
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the detected area. Unit: pixel.
	//
	// example:
	//
	// 330
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance from the upper-left corner of the detected area to the y-axis. The origin is the upper-left corner of the image. Unit: pixel.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance from the upper-left corner of the detected area to the x-axis. The origin is the upper-left corner of the image. Unit: pixel.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageModerationResponseBodyDataExtPublicFigureLocation) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtPublicFigureLocation) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtPublicFigureLocation) GetH() *int32 {
	return s.H
}

func (s *ImageModerationResponseBodyDataExtPublicFigureLocation) GetW() *int32 {
	return s.W
}

func (s *ImageModerationResponseBodyDataExtPublicFigureLocation) GetX() *int32 {
	return s.X
}

func (s *ImageModerationResponseBodyDataExtPublicFigureLocation) GetY() *int32 {
	return s.Y
}

func (s *ImageModerationResponseBodyDataExtPublicFigureLocation) SetH(v int32) *ImageModerationResponseBodyDataExtPublicFigureLocation {
	s.H = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtPublicFigureLocation) SetW(v int32) *ImageModerationResponseBodyDataExtPublicFigureLocation {
	s.W = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtPublicFigureLocation) SetX(v int32) *ImageModerationResponseBodyDataExtPublicFigureLocation {
	s.X = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtPublicFigureLocation) SetY(v int32) *ImageModerationResponseBodyDataExtPublicFigureLocation {
	s.Y = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtPublicFigureLocation) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtRecognition struct {
	// The category of the recognized object in the image.
	//
	// example:
	//
	// 办公大楼
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	// The confidence level. The value ranges from 0 to 100, with two decimal places retained. No confidence level is returned when the value is nonLabel.
	//
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s ImageModerationResponseBodyDataExtRecognition) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtRecognition) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtRecognition) GetClassification() *string {
	return s.Classification
}

func (s *ImageModerationResponseBodyDataExtRecognition) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageModerationResponseBodyDataExtRecognition) SetClassification(v string) *ImageModerationResponseBodyDataExtRecognition {
	s.Classification = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtRecognition) SetConfidence(v float32) *ImageModerationResponseBodyDataExtRecognition {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtRecognition) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtTextInImage struct {
	// If a custom text library is hit, the custom library ID, custom library name, and custom word are returned.
	CustomText []*ImageModerationResponseBodyDataExtTextInImageCustomText `json:"CustomText,omitempty" xml:"CustomText,omitempty" type:"Repeated"`
	// Each line of text recognized in the image.
	OcrResult []*ImageModerationResponseBodyDataExtTextInImageOcrResult `json:"OcrResult,omitempty" xml:"OcrResult,omitempty" type:"Repeated"`
	// The hit threat keywords.
	RiskWord []*string `json:"RiskWord,omitempty" xml:"RiskWord,omitempty" type:"Repeated"`
}

func (s ImageModerationResponseBodyDataExtTextInImage) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtTextInImage) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtTextInImage) GetCustomText() []*ImageModerationResponseBodyDataExtTextInImageCustomText {
	return s.CustomText
}

func (s *ImageModerationResponseBodyDataExtTextInImage) GetOcrResult() []*ImageModerationResponseBodyDataExtTextInImageOcrResult {
	return s.OcrResult
}

func (s *ImageModerationResponseBodyDataExtTextInImage) GetRiskWord() []*string {
	return s.RiskWord
}

func (s *ImageModerationResponseBodyDataExtTextInImage) SetCustomText(v []*ImageModerationResponseBodyDataExtTextInImageCustomText) *ImageModerationResponseBodyDataExtTextInImage {
	s.CustomText = v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImage) SetOcrResult(v []*ImageModerationResponseBodyDataExtTextInImageOcrResult) *ImageModerationResponseBodyDataExtTextInImage {
	s.OcrResult = v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImage) SetRiskWord(v []*string) *ImageModerationResponseBodyDataExtTextInImage {
	s.RiskWord = v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImage) Validate() error {
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

type ImageModerationResponseBodyDataExtTextInImageCustomText struct {
	// The custom words. Separate multiple words with commas.
	//
	// example:
	//
	// 自定义词1,自定义词2
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// The custom library ID.
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

func (s ImageModerationResponseBodyDataExtTextInImageCustomText) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtTextInImageCustomText) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtTextInImageCustomText) GetKeyWords() *string {
	return s.KeyWords
}

func (s *ImageModerationResponseBodyDataExtTextInImageCustomText) GetLibId() *string {
	return s.LibId
}

func (s *ImageModerationResponseBodyDataExtTextInImageCustomText) GetLibName() *string {
	return s.LibName
}

func (s *ImageModerationResponseBodyDataExtTextInImageCustomText) SetKeyWords(v string) *ImageModerationResponseBodyDataExtTextInImageCustomText {
	s.KeyWords = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageCustomText) SetLibId(v string) *ImageModerationResponseBodyDataExtTextInImageCustomText {
	s.LibId = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageCustomText) SetLibName(v string) *ImageModerationResponseBodyDataExtTextInImageCustomText {
	s.LibName = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageCustomText) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtTextInImageOcrResult struct {
	// The coordinates of the text line.
	Location *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// The text.
	//
	// example:
	//
	// abcd
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ImageModerationResponseBodyDataExtTextInImageOcrResult) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtTextInImageOcrResult) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResult) GetLocation() *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation {
	return s.Location
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResult) GetText() *string {
	return s.Text
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResult) SetLocation(v *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) *ImageModerationResponseBodyDataExtTextInImageOcrResult {
	s.Location = v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResult) SetText(v string) *ImageModerationResponseBodyDataExtTextInImageOcrResult {
	s.Text = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResult) Validate() error {
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageModerationResponseBodyDataExtTextInImageOcrResultLocation struct {
	// The height of the text area. Unit: pixel.
	//
	// example:
	//
	// 33
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the text area. Unit: pixel.
	//
	// example:
	//
	// 44
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance from the upper-left corner of the text area to the y-axis. The origin is the upper-left corner of the image. Unit: pixel.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance from the upper-left corner of the text area to the x-axis. The origin is the upper-left corner of the image. Unit: pixel.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) GetH() *int32 {
	return s.H
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) GetW() *int32 {
	return s.W
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) GetX() *int32 {
	return s.X
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) GetY() *int32 {
	return s.Y
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) SetH(v int32) *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation {
	s.H = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) SetW(v int32) *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation {
	s.W = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) SetX(v int32) *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation {
	s.X = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) SetY(v int32) *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation {
	s.Y = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataExtVlContent struct {
	// The output content.
	//
	// example:
	//
	// 这是一段描述
	OutputText *string `json:"OutputText,omitempty" xml:"OutputText,omitempty"`
}

func (s ImageModerationResponseBodyDataExtVlContent) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtVlContent) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtVlContent) GetOutputText() *string {
	return s.OutputText
}

func (s *ImageModerationResponseBodyDataExtVlContent) SetOutputText(v string) *ImageModerationResponseBodyDataExtVlContent {
	s.OutputText = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtVlContent) Validate() error {
	return dara.Validate(s)
}

type ImageModerationResponseBodyDataResult struct {
	// The confidence level. The value ranges from 0 to 100, with two decimal places retained. Some labels do not have a confidence level.
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
	// The label returned after the image content is moderated. Multiple labels and scores may be returned for a single image.
	//
	// example:
	//
	// violent_explosion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The threat level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ImageModerationResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageModerationResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *ImageModerationResponseBodyDataResult) GetLabel() *string {
	return s.Label
}

func (s *ImageModerationResponseBodyDataResult) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ImageModerationResponseBodyDataResult) SetConfidence(v float32) *ImageModerationResponseBodyDataResult {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataResult) SetDescription(v string) *ImageModerationResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *ImageModerationResponseBodyDataResult) SetLabel(v string) *ImageModerationResponseBodyDataResult {
	s.Label = &v
	return s
}

func (s *ImageModerationResponseBodyDataResult) SetRiskLevel(v string) *ImageModerationResponseBodyDataResult {
	s.RiskLevel = &v
	return s
}

func (s *ImageModerationResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
