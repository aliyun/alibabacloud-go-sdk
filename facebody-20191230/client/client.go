// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type ExtractPedestrianFeatureAttrRequest struct {
	ImageURL       *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Mode           *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s ExtractPedestrianFeatureAttrRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttrRequest) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttrRequest) SetImageURL(v string) *ExtractPedestrianFeatureAttrRequest {
	s.ImageURL = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrRequest) SetMode(v string) *ExtractPedestrianFeatureAttrRequest {
	s.Mode = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrRequest) SetServiceVersion(v string) *ExtractPedestrianFeatureAttrRequest {
	s.ServiceVersion = &v
	return s
}

type ExtractPedestrianFeatureAttrAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	Mode           *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ServiceVersion *string   `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s ExtractPedestrianFeatureAttrAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttrAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttrAdvanceRequest) SetImageURLObject(v io.Reader) *ExtractPedestrianFeatureAttrAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *ExtractPedestrianFeatureAttrAdvanceRequest) SetMode(v string) *ExtractPedestrianFeatureAttrAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrAdvanceRequest) SetServiceVersion(v string) *ExtractPedestrianFeatureAttrAdvanceRequest {
	s.ServiceVersion = &v
	return s
}

type ExtractPedestrianFeatureAttrResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ExtractPedestrianFeatureAttrResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ExtractPedestrianFeatureAttrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttrResponseBody) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttrResponseBody) SetRequestId(v string) *ExtractPedestrianFeatureAttrResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBody) SetData(v *ExtractPedestrianFeatureAttrResponseBodyData) *ExtractPedestrianFeatureAttrResponseBody {
	s.Data = v
	return s
}

type ExtractPedestrianFeatureAttrResponseBodyData struct {
	QualityScore    *float32 `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
	ObjType         *string  `json:"ObjType,omitempty" xml:"ObjType,omitempty"`
	Feature         *string  `json:"Feature,omitempty" xml:"Feature,omitempty"`
	Gender          *string  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	LowerColorScore *float32 `json:"LowerColorScore,omitempty" xml:"LowerColorScore,omitempty"`
	ObjTypeScore    *float32 `json:"ObjTypeScore,omitempty" xml:"ObjTypeScore,omitempty"`
	Age             *string  `json:"Age,omitempty" xml:"Age,omitempty"`
	AgeScore        *float32 `json:"AgeScore,omitempty" xml:"AgeScore,omitempty"`
	UpperTypeScore  *float32 `json:"UpperTypeScore,omitempty" xml:"UpperTypeScore,omitempty"`
	LowerTypeScore  *float32 `json:"LowerTypeScore,omitempty" xml:"LowerTypeScore,omitempty"`
	LowerColor      *string  `json:"LowerColor,omitempty" xml:"LowerColor,omitempty"`
	Hair            *string  `json:"Hair,omitempty" xml:"Hair,omitempty"`
	UpperColor      *string  `json:"UpperColor,omitempty" xml:"UpperColor,omitempty"`
	GenderScore     *float32 `json:"GenderScore,omitempty" xml:"GenderScore,omitempty"`
	UpperType       *string  `json:"UpperType,omitempty" xml:"UpperType,omitempty"`
	HairScore       *float32 `json:"HairScore,omitempty" xml:"HairScore,omitempty"`
	LowerType       *string  `json:"LowerType,omitempty" xml:"LowerType,omitempty"`
	UpperColorScore *float32 `json:"UpperColorScore,omitempty" xml:"UpperColorScore,omitempty"`
}

func (s ExtractPedestrianFeatureAttrResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttrResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetQualityScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.QualityScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetObjType(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.ObjType = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetFeature(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.Feature = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetGender(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.Gender = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetLowerColorScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.LowerColorScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetObjTypeScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.ObjTypeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetAge(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.Age = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetAgeScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.AgeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetUpperTypeScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.UpperTypeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetLowerTypeScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.LowerTypeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetLowerColor(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.LowerColor = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetHair(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.Hair = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetUpperColor(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.UpperColor = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetGenderScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.GenderScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetUpperType(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.UpperType = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetHairScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.HairScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetLowerType(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.LowerType = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetUpperColorScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.UpperColorScore = &v
	return s
}

type ExtractPedestrianFeatureAttrResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExtractPedestrianFeatureAttrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExtractPedestrianFeatureAttrResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttrResponse) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttrResponse) SetHeaders(v map[string]*string) *ExtractPedestrianFeatureAttrResponse {
	s.Headers = v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponse) SetBody(v *ExtractPedestrianFeatureAttrResponseBody) *ExtractPedestrianFeatureAttrResponse {
	s.Body = v
	return s
}

type DetectBodyCountRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectBodyCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectBodyCountRequest) GoString() string {
	return s.String()
}

func (s *DetectBodyCountRequest) SetImageURL(v string) *DetectBodyCountRequest {
	s.ImageURL = &v
	return s
}

type DetectBodyCountAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s DetectBodyCountAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectBodyCountAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectBodyCountAdvanceRequest) SetImageURLObject(v io.Reader) *DetectBodyCountAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectBodyCountResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectBodyCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectBodyCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectBodyCountResponseBody) GoString() string {
	return s.String()
}

func (s *DetectBodyCountResponseBody) SetRequestId(v string) *DetectBodyCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectBodyCountResponseBody) SetData(v *DetectBodyCountResponseBodyData) *DetectBodyCountResponseBody {
	s.Data = v
	return s
}

type DetectBodyCountResponseBodyData struct {
	PersonNumber *int32 `json:"PersonNumber,omitempty" xml:"PersonNumber,omitempty"`
}

func (s DetectBodyCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectBodyCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectBodyCountResponseBodyData) SetPersonNumber(v int32) *DetectBodyCountResponseBodyData {
	s.PersonNumber = &v
	return s
}

type DetectBodyCountResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectBodyCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectBodyCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectBodyCountResponse) GoString() string {
	return s.String()
}

func (s *DetectBodyCountResponse) SetHeaders(v map[string]*string) *DetectBodyCountResponse {
	s.Headers = v
	return s
}

func (s *DetectBodyCountResponse) SetBody(v *DetectBodyCountResponseBody) *DetectBodyCountResponse {
	s.Body = v
	return s
}

type DetectVideoLivingFaceRequest struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s DetectVideoLivingFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoLivingFaceRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceRequest) SetVideoUrl(v string) *DetectVideoLivingFaceRequest {
	s.VideoUrl = &v
	return s
}

type DetectVideoLivingFaceAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
}

func (s DetectVideoLivingFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoLivingFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceAdvanceRequest) SetVideoUrlObject(v io.Reader) *DetectVideoLivingFaceAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type DetectVideoLivingFaceResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectVideoLivingFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectVideoLivingFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoLivingFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceResponseBody) SetRequestId(v string) *DetectVideoLivingFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectVideoLivingFaceResponseBody) SetData(v *DetectVideoLivingFaceResponseBodyData) *DetectVideoLivingFaceResponseBody {
	s.Data = v
	return s
}

type DetectVideoLivingFaceResponseBodyData struct {
	Elements []*DetectVideoLivingFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectVideoLivingFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoLivingFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceResponseBodyData) SetElements(v []*DetectVideoLivingFaceResponseBodyDataElements) *DetectVideoLivingFaceResponseBodyData {
	s.Elements = v
	return s
}

type DetectVideoLivingFaceResponseBodyDataElements struct {
	Rect           []*int32 `json:"Rect,omitempty" xml:"Rect,omitempty" type:"Repeated"`
	LiveConfidence *float32 `json:"LiveConfidence,omitempty" xml:"LiveConfidence,omitempty"`
	FaceConfidence *float32 `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
}

func (s DetectVideoLivingFaceResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoLivingFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceResponseBodyDataElements) SetRect(v []*int32) *DetectVideoLivingFaceResponseBodyDataElements {
	s.Rect = v
	return s
}

func (s *DetectVideoLivingFaceResponseBodyDataElements) SetLiveConfidence(v float32) *DetectVideoLivingFaceResponseBodyDataElements {
	s.LiveConfidence = &v
	return s
}

func (s *DetectVideoLivingFaceResponseBodyDataElements) SetFaceConfidence(v float32) *DetectVideoLivingFaceResponseBodyDataElements {
	s.FaceConfidence = &v
	return s
}

type DetectVideoLivingFaceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectVideoLivingFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectVideoLivingFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoLivingFaceResponse) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceResponse) SetHeaders(v map[string]*string) *DetectVideoLivingFaceResponse {
	s.Headers = v
	return s
}

func (s *DetectVideoLivingFaceResponse) SetBody(v *DetectVideoLivingFaceResponseBody) *DetectVideoLivingFaceResponse {
	s.Body = v
	return s
}

type RecognizeFaceRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFaceRequest) SetImageType(v int32) *RecognizeFaceRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeFaceRequest) SetImageURL(v string) *RecognizeFaceRequest {
	s.ImageURL = &v
	return s
}

type RecognizeFaceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s RecognizeFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFaceAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetImageType(v int32) *RecognizeFaceAdvanceRequest {
	s.ImageType = &v
	return s
}

type RecognizeFaceResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponseBody) SetRequestId(v string) *RecognizeFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeFaceResponseBody) SetData(v *RecognizeFaceResponseBodyData) *RecognizeFaceResponseBody {
	s.Data = v
	return s
}

type RecognizeFaceResponseBodyData struct {
	Pupils              []*float32                              `json:"Pupils,omitempty" xml:"Pupils,omitempty" type:"Repeated"`
	GenderList          []*int32                                `json:"GenderList,omitempty" xml:"GenderList,omitempty" type:"Repeated"`
	Expressions         []*int32                                `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	DenseFeatures       []*string                               `json:"DenseFeatures,omitempty" xml:"DenseFeatures,omitempty" type:"Repeated"`
	FaceCount           *int32                                  `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	Landmarks           []*float32                              `json:"Landmarks,omitempty" xml:"Landmarks,omitempty" type:"Repeated"`
	LandmarkCount       *int32                                  `json:"LandmarkCount,omitempty" xml:"LandmarkCount,omitempty"`
	Qualities           *RecognizeFaceResponseBodyDataQualities `json:"Qualities,omitempty" xml:"Qualities,omitempty" type:"Struct"`
	BeuatyList          []*float32                              `json:"BeuatyList,omitempty" xml:"BeuatyList,omitempty" type:"Repeated"`
	HatList             []*int32                                `json:"HatList,omitempty" xml:"HatList,omitempty" type:"Repeated"`
	FaceProbabilityList []*float32                              `json:"FaceProbabilityList,omitempty" xml:"FaceProbabilityList,omitempty" type:"Repeated"`
	Glasses             []*int32                                `json:"Glasses,omitempty" xml:"Glasses,omitempty" type:"Repeated"`
	FaceRectangles      []*int32                                `json:"FaceRectangles,omitempty" xml:"FaceRectangles,omitempty" type:"Repeated"`
	PoseList            []*float32                              `json:"PoseList,omitempty" xml:"PoseList,omitempty" type:"Repeated"`
	AgeList             []*int32                                `json:"AgeList,omitempty" xml:"AgeList,omitempty" type:"Repeated"`
	DenseFeatureLength  *int32                                  `json:"DenseFeatureLength,omitempty" xml:"DenseFeatureLength,omitempty"`
}

func (s RecognizeFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponseBodyData) SetPupils(v []*float32) *RecognizeFaceResponseBodyData {
	s.Pupils = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetGenderList(v []*int32) *RecognizeFaceResponseBodyData {
	s.GenderList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetExpressions(v []*int32) *RecognizeFaceResponseBodyData {
	s.Expressions = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetDenseFeatures(v []*string) *RecognizeFaceResponseBodyData {
	s.DenseFeatures = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetFaceCount(v int32) *RecognizeFaceResponseBodyData {
	s.FaceCount = &v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetLandmarks(v []*float32) *RecognizeFaceResponseBodyData {
	s.Landmarks = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetLandmarkCount(v int32) *RecognizeFaceResponseBodyData {
	s.LandmarkCount = &v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetQualities(v *RecognizeFaceResponseBodyDataQualities) *RecognizeFaceResponseBodyData {
	s.Qualities = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetBeuatyList(v []*float32) *RecognizeFaceResponseBodyData {
	s.BeuatyList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetHatList(v []*int32) *RecognizeFaceResponseBodyData {
	s.HatList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetFaceProbabilityList(v []*float32) *RecognizeFaceResponseBodyData {
	s.FaceProbabilityList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetGlasses(v []*int32) *RecognizeFaceResponseBodyData {
	s.Glasses = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetFaceRectangles(v []*int32) *RecognizeFaceResponseBodyData {
	s.FaceRectangles = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetPoseList(v []*float32) *RecognizeFaceResponseBodyData {
	s.PoseList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetAgeList(v []*int32) *RecognizeFaceResponseBodyData {
	s.AgeList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetDenseFeatureLength(v int32) *RecognizeFaceResponseBodyData {
	s.DenseFeatureLength = &v
	return s
}

type RecognizeFaceResponseBodyDataQualities struct {
	ScoreList []*float32 `json:"ScoreList,omitempty" xml:"ScoreList,omitempty" type:"Repeated"`
	BlurList  []*float32 `json:"BlurList,omitempty" xml:"BlurList,omitempty" type:"Repeated"`
	FnfList   []*float32 `json:"FnfList,omitempty" xml:"FnfList,omitempty" type:"Repeated"`
	GlassList []*float32 `json:"GlassList,omitempty" xml:"GlassList,omitempty" type:"Repeated"`
	IlluList  []*float32 `json:"IlluList,omitempty" xml:"IlluList,omitempty" type:"Repeated"`
	MaskList  []*float32 `json:"MaskList,omitempty" xml:"MaskList,omitempty" type:"Repeated"`
	NoiseList []*float32 `json:"NoiseList,omitempty" xml:"NoiseList,omitempty" type:"Repeated"`
	PoseList  []*float32 `json:"PoseList,omitempty" xml:"PoseList,omitempty" type:"Repeated"`
}

func (s RecognizeFaceResponseBodyDataQualities) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceResponseBodyDataQualities) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponseBodyDataQualities) SetScoreList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.ScoreList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetBlurList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.BlurList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetFnfList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.FnfList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetGlassList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.GlassList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetIlluList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.IlluList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetMaskList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.MaskList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetNoiseList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.NoiseList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetPoseList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.PoseList = v
	return s
}

type RecognizeFaceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponse) SetHeaders(v map[string]*string) *RecognizeFaceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeFaceResponse) SetBody(v *RecognizeFaceResponseBody) *RecognizeFaceResponse {
	s.Body = v
	return s
}

type VerifyFaceMaskRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RefUrl   *string `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
}

func (s VerifyFaceMaskRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyFaceMaskRequest) GoString() string {
	return s.String()
}

func (s *VerifyFaceMaskRequest) SetImageURL(v string) *VerifyFaceMaskRequest {
	s.ImageURL = &v
	return s
}

func (s *VerifyFaceMaskRequest) SetRefUrl(v string) *VerifyFaceMaskRequest {
	s.RefUrl = &v
	return s
}

type VerifyFaceMaskAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	RefUrl         *string   `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
}

func (s VerifyFaceMaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyFaceMaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *VerifyFaceMaskAdvanceRequest) SetImageURLObject(v io.Reader) *VerifyFaceMaskAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *VerifyFaceMaskAdvanceRequest) SetRefUrl(v string) *VerifyFaceMaskAdvanceRequest {
	s.RefUrl = &v
	return s
}

type VerifyFaceMaskResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *VerifyFaceMaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s VerifyFaceMaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyFaceMaskResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyFaceMaskResponseBody) SetRequestId(v string) *VerifyFaceMaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyFaceMaskResponseBody) SetData(v *VerifyFaceMaskResponseBodyData) *VerifyFaceMaskResponseBody {
	s.Data = v
	return s
}

type VerifyFaceMaskResponseBodyData struct {
	Thresholds   []*float32 `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Repeated"`
	Mask         *int32     `json:"Mask,omitempty" xml:"Mask,omitempty"`
	Confidence   *float32   `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Rectangle    []*int32   `json:"Rectangle,omitempty" xml:"Rectangle,omitempty" type:"Repeated"`
	RectangleRef []*int32   `json:"RectangleRef,omitempty" xml:"RectangleRef,omitempty" type:"Repeated"`
	MaskRef      *int32     `json:"MaskRef,omitempty" xml:"MaskRef,omitempty"`
}

func (s VerifyFaceMaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VerifyFaceMaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyFaceMaskResponseBodyData) SetThresholds(v []*float32) *VerifyFaceMaskResponseBodyData {
	s.Thresholds = v
	return s
}

func (s *VerifyFaceMaskResponseBodyData) SetMask(v int32) *VerifyFaceMaskResponseBodyData {
	s.Mask = &v
	return s
}

func (s *VerifyFaceMaskResponseBodyData) SetConfidence(v float32) *VerifyFaceMaskResponseBodyData {
	s.Confidence = &v
	return s
}

func (s *VerifyFaceMaskResponseBodyData) SetRectangle(v []*int32) *VerifyFaceMaskResponseBodyData {
	s.Rectangle = v
	return s
}

func (s *VerifyFaceMaskResponseBodyData) SetRectangleRef(v []*int32) *VerifyFaceMaskResponseBodyData {
	s.RectangleRef = v
	return s
}

func (s *VerifyFaceMaskResponseBodyData) SetMaskRef(v int32) *VerifyFaceMaskResponseBodyData {
	s.MaskRef = &v
	return s
}

type VerifyFaceMaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyFaceMaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyFaceMaskResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyFaceMaskResponse) GoString() string {
	return s.String()
}

func (s *VerifyFaceMaskResponse) SetHeaders(v map[string]*string) *VerifyFaceMaskResponse {
	s.Headers = v
	return s
}

func (s *VerifyFaceMaskResponse) SetBody(v *VerifyFaceMaskResponseBody) *VerifyFaceMaskResponse {
	s.Body = v
	return s
}

type DetectIPCPedestrianRequest struct {
	ContinueOnError *bool                                `json:"ContinueOnError,omitempty" xml:"ContinueOnError,omitempty"`
	ImageData       *string                              `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	Width           *int32                               `json:"Width,omitempty" xml:"Width,omitempty"`
	Height          *int32                               `json:"Height,omitempty" xml:"Height,omitempty"`
	URLList         []*DetectIPCPedestrianRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s DetectIPCPedestrianRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianRequest) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianRequest) SetContinueOnError(v bool) *DetectIPCPedestrianRequest {
	s.ContinueOnError = &v
	return s
}

func (s *DetectIPCPedestrianRequest) SetImageData(v string) *DetectIPCPedestrianRequest {
	s.ImageData = &v
	return s
}

func (s *DetectIPCPedestrianRequest) SetWidth(v int32) *DetectIPCPedestrianRequest {
	s.Width = &v
	return s
}

func (s *DetectIPCPedestrianRequest) SetHeight(v int32) *DetectIPCPedestrianRequest {
	s.Height = &v
	return s
}

func (s *DetectIPCPedestrianRequest) SetURLList(v []*DetectIPCPedestrianRequestURLList) *DetectIPCPedestrianRequest {
	s.URLList = v
	return s
}

type DetectIPCPedestrianRequestURLList struct {
}

func (s DetectIPCPedestrianRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianRequestURLList) GoString() string {
	return s.String()
}

type DetectIPCPedestrianResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectIPCPedestrianResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectIPCPedestrianResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianResponseBody) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianResponseBody) SetRequestId(v string) *DetectIPCPedestrianResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectIPCPedestrianResponseBody) SetData(v *DetectIPCPedestrianResponseBodyData) *DetectIPCPedestrianResponseBody {
	s.Data = v
	return s
}

type DetectIPCPedestrianResponseBodyData struct {
	ImageInfoList []*DetectIPCPedestrianResponseBodyDataImageInfoList `json:"ImageInfoList,omitempty" xml:"ImageInfoList,omitempty" type:"Repeated"`
}

func (s DetectIPCPedestrianResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianResponseBodyData) SetImageInfoList(v []*DetectIPCPedestrianResponseBodyDataImageInfoList) *DetectIPCPedestrianResponseBodyData {
	s.ImageInfoList = v
	return s
}

type DetectIPCPedestrianResponseBodyDataImageInfoList struct {
	Elements []*DetectIPCPedestrianResponseBodyDataImageInfoListElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectIPCPedestrianResponseBodyDataImageInfoList) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianResponseBodyDataImageInfoList) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianResponseBodyDataImageInfoList) SetElements(v []*DetectIPCPedestrianResponseBodyDataImageInfoListElements) *DetectIPCPedestrianResponseBodyDataImageInfoList {
	s.Elements = v
	return s
}

type DetectIPCPedestrianResponseBodyDataImageInfoListElements struct {
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s DetectIPCPedestrianResponseBodyDataImageInfoListElements) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianResponseBodyDataImageInfoListElements) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianResponseBodyDataImageInfoListElements) SetBoxes(v []*int32) *DetectIPCPedestrianResponseBodyDataImageInfoListElements {
	s.Boxes = v
	return s
}

func (s *DetectIPCPedestrianResponseBodyDataImageInfoListElements) SetScore(v float32) *DetectIPCPedestrianResponseBodyDataImageInfoListElements {
	s.Score = &v
	return s
}

type DetectIPCPedestrianResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectIPCPedestrianResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectIPCPedestrianResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianResponse) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianResponse) SetHeaders(v map[string]*string) *DetectIPCPedestrianResponse {
	s.Headers = v
	return s
}

func (s *DetectIPCPedestrianResponse) SetBody(v *DetectIPCPedestrianResponseBody) *DetectIPCPedestrianResponse {
	s.Body = v
	return s
}

type GetFaceEntityRequest struct {
	DbName   *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s GetFaceEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFaceEntityRequest) GoString() string {
	return s.String()
}

func (s *GetFaceEntityRequest) SetDbName(v string) *GetFaceEntityRequest {
	s.DbName = &v
	return s
}

func (s *GetFaceEntityRequest) SetEntityId(v string) *GetFaceEntityRequest {
	s.EntityId = &v
	return s
}

type GetFaceEntityResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetFaceEntityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetFaceEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFaceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetFaceEntityResponseBody) SetRequestId(v string) *GetFaceEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFaceEntityResponseBody) SetData(v *GetFaceEntityResponseBodyData) *GetFaceEntityResponseBody {
	s.Data = v
	return s
}

type GetFaceEntityResponseBodyData struct {
	DbName   *string                               `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId *string                               `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Labels   *string                               `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Faces    []*GetFaceEntityResponseBodyDataFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
}

func (s GetFaceEntityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFaceEntityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFaceEntityResponseBodyData) SetDbName(v string) *GetFaceEntityResponseBodyData {
	s.DbName = &v
	return s
}

func (s *GetFaceEntityResponseBodyData) SetEntityId(v string) *GetFaceEntityResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *GetFaceEntityResponseBodyData) SetLabels(v string) *GetFaceEntityResponseBodyData {
	s.Labels = &v
	return s
}

func (s *GetFaceEntityResponseBodyData) SetFaces(v []*GetFaceEntityResponseBodyDataFaces) *GetFaceEntityResponseBodyData {
	s.Faces = v
	return s
}

type GetFaceEntityResponseBodyDataFaces struct {
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
}

func (s GetFaceEntityResponseBodyDataFaces) String() string {
	return tea.Prettify(s)
}

func (s GetFaceEntityResponseBodyDataFaces) GoString() string {
	return s.String()
}

func (s *GetFaceEntityResponseBodyDataFaces) SetFaceId(v string) *GetFaceEntityResponseBodyDataFaces {
	s.FaceId = &v
	return s
}

type GetFaceEntityResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFaceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFaceEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFaceEntityResponse) GoString() string {
	return s.String()
}

func (s *GetFaceEntityResponse) SetHeaders(v map[string]*string) *GetFaceEntityResponse {
	s.Headers = v
	return s
}

func (s *GetFaceEntityResponse) SetBody(v *GetFaceEntityResponseBody) *GetFaceEntityResponse {
	s.Body = v
	return s
}

type CompareFaceRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURLA *string `json:"ImageURLA,omitempty" xml:"ImageURLA,omitempty"`
	ImageURLB *string `json:"ImageURLB,omitempty" xml:"ImageURLB,omitempty"`
}

func (s CompareFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceRequest) GoString() string {
	return s.String()
}

func (s *CompareFaceRequest) SetImageType(v int32) *CompareFaceRequest {
	s.ImageType = &v
	return s
}

func (s *CompareFaceRequest) SetImageURLA(v string) *CompareFaceRequest {
	s.ImageURLA = &v
	return s
}

func (s *CompareFaceRequest) SetImageURLB(v string) *CompareFaceRequest {
	s.ImageURLB = &v
	return s
}

type CompareFaceResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CompareFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CompareFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceResponseBody) GoString() string {
	return s.String()
}

func (s *CompareFaceResponseBody) SetRequestId(v string) *CompareFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareFaceResponseBody) SetData(v *CompareFaceResponseBodyData) *CompareFaceResponseBody {
	s.Data = v
	return s
}

type CompareFaceResponseBodyData struct {
	Thresholds []*float32 `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Repeated"`
	RectBList  []*int32   `json:"RectBList,omitempty" xml:"RectBList,omitempty" type:"Repeated"`
	Confidence *float32   `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	RectAList  []*int32   `json:"RectAList,omitempty" xml:"RectAList,omitempty" type:"Repeated"`
}

func (s CompareFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompareFaceResponseBodyData) SetThresholds(v []*float32) *CompareFaceResponseBodyData {
	s.Thresholds = v
	return s
}

func (s *CompareFaceResponseBodyData) SetRectBList(v []*int32) *CompareFaceResponseBodyData {
	s.RectBList = v
	return s
}

func (s *CompareFaceResponseBodyData) SetConfidence(v float32) *CompareFaceResponseBodyData {
	s.Confidence = &v
	return s
}

func (s *CompareFaceResponseBodyData) SetRectAList(v []*int32) *CompareFaceResponseBodyData {
	s.RectAList = v
	return s
}

type CompareFaceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CompareFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompareFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceResponse) GoString() string {
	return s.String()
}

func (s *CompareFaceResponse) SetHeaders(v map[string]*string) *CompareFaceResponse {
	s.Headers = v
	return s
}

func (s *CompareFaceResponse) SetBody(v *CompareFaceResponseBody) *CompareFaceResponse {
	s.Body = v
	return s
}

type PedestrianDetectAttributeRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s PedestrianDetectAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeRequest) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeRequest) SetImageURL(v string) *PedestrianDetectAttributeRequest {
	s.ImageURL = &v
	return s
}

type PedestrianDetectAttributeAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s PedestrianDetectAttributeAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeAdvanceRequest) SetImageURLObject(v io.Reader) *PedestrianDetectAttributeAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type PedestrianDetectAttributeResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *PedestrianDetectAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s PedestrianDetectAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBody) SetRequestId(v string) *PedestrianDetectAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBody) SetData(v *PedestrianDetectAttributeResponseBodyData) *PedestrianDetectAttributeResponseBody {
	s.Data = v
	return s
}

type PedestrianDetectAttributeResponseBodyData struct {
	Attributes   []*PedestrianDetectAttributeResponseBodyDataAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Boxes        []*PedestrianDetectAttributeResponseBodyDataBoxes      `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	PersonNumber *int32                                                 `json:"PersonNumber,omitempty" xml:"PersonNumber,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyData) SetAttributes(v []*PedestrianDetectAttributeResponseBodyDataAttributes) *PedestrianDetectAttributeResponseBodyData {
	s.Attributes = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyData) SetBoxes(v []*PedestrianDetectAttributeResponseBodyDataBoxes) *PedestrianDetectAttributeResponseBodyData {
	s.Boxes = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyData) SetPersonNumber(v int32) *PedestrianDetectAttributeResponseBodyData {
	s.PersonNumber = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributes struct {
	Gender      *PedestrianDetectAttributeResponseBodyDataAttributesGender      `json:"Gender,omitempty" xml:"Gender,omitempty" type:"Struct"`
	Orient      *PedestrianDetectAttributeResponseBodyDataAttributesOrient      `json:"Orient,omitempty" xml:"Orient,omitempty" type:"Struct"`
	Age         *PedestrianDetectAttributeResponseBodyDataAttributesAge         `json:"Age,omitempty" xml:"Age,omitempty" type:"Struct"`
	UpperWear   *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear   `json:"UpperWear,omitempty" xml:"UpperWear,omitempty" type:"Struct"`
	Glasses     *PedestrianDetectAttributeResponseBodyDataAttributesGlasses     `json:"Glasses,omitempty" xml:"Glasses,omitempty" type:"Struct"`
	LowerWear   *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear   `json:"LowerWear,omitempty" xml:"LowerWear,omitempty" type:"Struct"`
	LowerColor  *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor  `json:"LowerColor,omitempty" xml:"LowerColor,omitempty" type:"Struct"`
	Hat         *PedestrianDetectAttributeResponseBodyDataAttributesHat         `json:"Hat,omitempty" xml:"Hat,omitempty" type:"Struct"`
	Handbag     *PedestrianDetectAttributeResponseBodyDataAttributesHandbag     `json:"Handbag,omitempty" xml:"Handbag,omitempty" type:"Struct"`
	Backpack    *PedestrianDetectAttributeResponseBodyDataAttributesBackpack    `json:"Backpack,omitempty" xml:"Backpack,omitempty" type:"Struct"`
	UpperColor  *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor  `json:"UpperColor,omitempty" xml:"UpperColor,omitempty" type:"Struct"`
	ShoulderBag *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag `json:"ShoulderBag,omitempty" xml:"ShoulderBag,omitempty" type:"Struct"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributes) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributes) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetGender(v *PedestrianDetectAttributeResponseBodyDataAttributesGender) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Gender = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetOrient(v *PedestrianDetectAttributeResponseBodyDataAttributesOrient) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Orient = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetAge(v *PedestrianDetectAttributeResponseBodyDataAttributesAge) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Age = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetUpperWear(v *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.UpperWear = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetGlasses(v *PedestrianDetectAttributeResponseBodyDataAttributesGlasses) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Glasses = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetLowerWear(v *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.LowerWear = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetLowerColor(v *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.LowerColor = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetHat(v *PedestrianDetectAttributeResponseBodyDataAttributesHat) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Hat = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetHandbag(v *PedestrianDetectAttributeResponseBodyDataAttributesHandbag) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Handbag = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetBackpack(v *PedestrianDetectAttributeResponseBodyDataAttributesBackpack) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Backpack = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetUpperColor(v *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.UpperColor = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetShoulderBag(v *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.ShoulderBag = v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesGender struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesGender) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesGender) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGender) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesGender {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGender) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesGender {
	s.Name = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesOrient struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesOrient) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesOrient) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesOrient) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesOrient {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesOrient) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesOrient {
	s.Name = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesAge struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesAge) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesAge) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesAge) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesAge {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesAge) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesAge {
	s.Name = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesUpperWear struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear {
	s.Name = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesGlasses struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesGlasses) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesGlasses) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGlasses) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesGlasses {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGlasses) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesGlasses {
	s.Name = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesLowerWear struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear {
	s.Name = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesLowerColor struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor {
	s.Name = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesHat struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesHat) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesHat) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHat) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesHat {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHat) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesHat {
	s.Name = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesHandbag struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesHandbag) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesHandbag) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHandbag) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesHandbag {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHandbag) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesHandbag {
	s.Name = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesBackpack struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesBackpack) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesBackpack) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesBackpack) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesBackpack {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesBackpack) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesBackpack {
	s.Name = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesUpperColor struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor {
	s.Name = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag {
	s.Name = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataBoxes struct {
	BottomRightX *float32 `json:"BottomRightX,omitempty" xml:"BottomRightX,omitempty"`
	TopLeftY     *float32 `json:"TopLeftY,omitempty" xml:"TopLeftY,omitempty"`
	Score        *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	TopLeftX     *float32 `json:"TopLeftX,omitempty" xml:"TopLeftX,omitempty"`
	BottomRightY *float32 `json:"BottomRightY,omitempty" xml:"BottomRightY,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataBoxes) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataBoxes) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetBottomRightX(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.BottomRightX = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetTopLeftY(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.TopLeftY = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetTopLeftX(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.TopLeftX = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetBottomRightY(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.BottomRightY = &v
	return s
}

type PedestrianDetectAttributeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PedestrianDetectAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PedestrianDetectAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponse) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponse) SetHeaders(v map[string]*string) *PedestrianDetectAttributeResponse {
	s.Headers = v
	return s
}

func (s *PedestrianDetectAttributeResponse) SetBody(v *PedestrianDetectAttributeResponseBody) *PedestrianDetectAttributeResponse {
	s.Body = v
	return s
}

type FaceFilterRequest struct {
	ImageURL     *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ResourceType *string  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Strength     *float32 `json:"Strength,omitempty" xml:"Strength,omitempty"`
}

func (s FaceFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceFilterRequest) GoString() string {
	return s.String()
}

func (s *FaceFilterRequest) SetImageURL(v string) *FaceFilterRequest {
	s.ImageURL = &v
	return s
}

func (s *FaceFilterRequest) SetResourceType(v string) *FaceFilterRequest {
	s.ResourceType = &v
	return s
}

func (s *FaceFilterRequest) SetStrength(v float32) *FaceFilterRequest {
	s.Strength = &v
	return s
}

type FaceFilterAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ResourceType   *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Strength       *float32  `json:"Strength,omitempty" xml:"Strength,omitempty"`
}

func (s FaceFilterAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceFilterAdvanceRequest) GoString() string {
	return s.String()
}

func (s *FaceFilterAdvanceRequest) SetImageURLObject(v io.Reader) *FaceFilterAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *FaceFilterAdvanceRequest) SetResourceType(v string) *FaceFilterAdvanceRequest {
	s.ResourceType = &v
	return s
}

func (s *FaceFilterAdvanceRequest) SetStrength(v float32) *FaceFilterAdvanceRequest {
	s.Strength = &v
	return s
}

type FaceFilterResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *FaceFilterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s FaceFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceFilterResponseBody) GoString() string {
	return s.String()
}

func (s *FaceFilterResponseBody) SetRequestId(v string) *FaceFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceFilterResponseBody) SetData(v *FaceFilterResponseBodyData) *FaceFilterResponseBody {
	s.Data = v
	return s
}

type FaceFilterResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s FaceFilterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FaceFilterResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceFilterResponseBodyData) SetImageURL(v string) *FaceFilterResponseBodyData {
	s.ImageURL = &v
	return s
}

type FaceFilterResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FaceFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FaceFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceFilterResponse) GoString() string {
	return s.String()
}

func (s *FaceFilterResponse) SetHeaders(v map[string]*string) *FaceFilterResponse {
	s.Headers = v
	return s
}

func (s *FaceFilterResponse) SetBody(v *FaceFilterResponseBody) *FaceFilterResponse {
	s.Body = v
	return s
}

type FaceBeautyRequest struct {
	ImageURL *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Sharp    *float32 `json:"Sharp,omitempty" xml:"Sharp,omitempty"`
	Smooth   *float32 `json:"Smooth,omitempty" xml:"Smooth,omitempty"`
	White    *float32 `json:"White,omitempty" xml:"White,omitempty"`
}

func (s FaceBeautyRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceBeautyRequest) GoString() string {
	return s.String()
}

func (s *FaceBeautyRequest) SetImageURL(v string) *FaceBeautyRequest {
	s.ImageURL = &v
	return s
}

func (s *FaceBeautyRequest) SetSharp(v float32) *FaceBeautyRequest {
	s.Sharp = &v
	return s
}

func (s *FaceBeautyRequest) SetSmooth(v float32) *FaceBeautyRequest {
	s.Smooth = &v
	return s
}

func (s *FaceBeautyRequest) SetWhite(v float32) *FaceBeautyRequest {
	s.White = &v
	return s
}

type FaceBeautyAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	Sharp          *float32  `json:"Sharp,omitempty" xml:"Sharp,omitempty"`
	Smooth         *float32  `json:"Smooth,omitempty" xml:"Smooth,omitempty"`
	White          *float32  `json:"White,omitempty" xml:"White,omitempty"`
}

func (s FaceBeautyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceBeautyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *FaceBeautyAdvanceRequest) SetImageURLObject(v io.Reader) *FaceBeautyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *FaceBeautyAdvanceRequest) SetSharp(v float32) *FaceBeautyAdvanceRequest {
	s.Sharp = &v
	return s
}

func (s *FaceBeautyAdvanceRequest) SetSmooth(v float32) *FaceBeautyAdvanceRequest {
	s.Smooth = &v
	return s
}

func (s *FaceBeautyAdvanceRequest) SetWhite(v float32) *FaceBeautyAdvanceRequest {
	s.White = &v
	return s
}

type FaceBeautyResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *FaceBeautyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s FaceBeautyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceBeautyResponseBody) GoString() string {
	return s.String()
}

func (s *FaceBeautyResponseBody) SetRequestId(v string) *FaceBeautyResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceBeautyResponseBody) SetData(v *FaceBeautyResponseBodyData) *FaceBeautyResponseBody {
	s.Data = v
	return s
}

type FaceBeautyResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s FaceBeautyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FaceBeautyResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceBeautyResponseBodyData) SetImageURL(v string) *FaceBeautyResponseBodyData {
	s.ImageURL = &v
	return s
}

type FaceBeautyResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FaceBeautyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FaceBeautyResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceBeautyResponse) GoString() string {
	return s.String()
}

func (s *FaceBeautyResponse) SetHeaders(v map[string]*string) *FaceBeautyResponse {
	s.Headers = v
	return s
}

func (s *FaceBeautyResponse) SetBody(v *FaceBeautyResponseBody) *FaceBeautyResponse {
	s.Body = v
	return s
}

type GenerateHumanAnimeStyleRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	AlgoType *string `json:"AlgoType,omitempty" xml:"AlgoType,omitempty"`
}

func (s GenerateHumanAnimeStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleRequest) SetImageURL(v string) *GenerateHumanAnimeStyleRequest {
	s.ImageURL = &v
	return s
}

func (s *GenerateHumanAnimeStyleRequest) SetAlgoType(v string) *GenerateHumanAnimeStyleRequest {
	s.AlgoType = &v
	return s
}

type GenerateHumanAnimeStyleAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	AlgoType       *string   `json:"AlgoType,omitempty" xml:"AlgoType,omitempty"`
}

func (s GenerateHumanAnimeStyleAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleAdvanceRequest) SetImageURLObject(v io.Reader) *GenerateHumanAnimeStyleAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *GenerateHumanAnimeStyleAdvanceRequest) SetAlgoType(v string) *GenerateHumanAnimeStyleAdvanceRequest {
	s.AlgoType = &v
	return s
}

type GenerateHumanAnimeStyleResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GenerateHumanAnimeStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GenerateHumanAnimeStyleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleResponseBody) SetRequestId(v string) *GenerateHumanAnimeStyleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateHumanAnimeStyleResponseBody) SetData(v *GenerateHumanAnimeStyleResponseBodyData) *GenerateHumanAnimeStyleResponseBody {
	s.Data = v
	return s
}

type GenerateHumanAnimeStyleResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s GenerateHumanAnimeStyleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleResponseBodyData) SetImageURL(v string) *GenerateHumanAnimeStyleResponseBodyData {
	s.ImageURL = &v
	return s
}

type GenerateHumanAnimeStyleResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateHumanAnimeStyleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateHumanAnimeStyleResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleResponse) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleResponse) SetHeaders(v map[string]*string) *GenerateHumanAnimeStyleResponse {
	s.Headers = v
	return s
}

func (s *GenerateHumanAnimeStyleResponse) SetBody(v *GenerateHumanAnimeStyleResponseBody) *GenerateHumanAnimeStyleResponse {
	s.Body = v
	return s
}

type QueryFaceImageTemplateRequest struct {
	// A short description of struct
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryFaceImageTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceImageTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateRequest) SetUserId(v string) *QueryFaceImageTemplateRequest {
	s.UserId = &v
	return s
}

func (s *QueryFaceImageTemplateRequest) SetTemplateId(v string) *QueryFaceImageTemplateRequest {
	s.TemplateId = &v
	return s
}

type QueryFaceImageTemplateResponseBody struct {
	// Id of the request
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *QueryFaceImageTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryFaceImageTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceImageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponseBody) SetRequestId(v string) *QueryFaceImageTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBody) SetData(v *QueryFaceImageTemplateResponseBodyData) *QueryFaceImageTemplateResponseBody {
	s.Data = v
	return s
}

type QueryFaceImageTemplateResponseBodyData struct {
	Elements []*QueryFaceImageTemplateResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s QueryFaceImageTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceImageTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponseBodyData) SetElements(v []*QueryFaceImageTemplateResponseBodyDataElements) *QueryFaceImageTemplateResponseBodyData {
	s.Elements = v
	return s
}

type QueryFaceImageTemplateResponseBodyDataElements struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
}

func (s QueryFaceImageTemplateResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceImageTemplateResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetCreateTime(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.CreateTime = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetUpdateTime(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.UpdateTime = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetUserId(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.UserId = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetTemplateId(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.TemplateId = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetTemplateURL(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.TemplateURL = &v
	return s
}

type QueryFaceImageTemplateResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFaceImageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFaceImageTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceImageTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponse) SetHeaders(v map[string]*string) *QueryFaceImageTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceImageTemplateResponse) SetBody(v *QueryFaceImageTemplateResponseBody) *QueryFaceImageTemplateResponse {
	s.Body = v
	return s
}

type DetectFaceRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceRequest) SetImageType(v int32) *DetectFaceRequest {
	s.ImageType = &v
	return s
}

func (s *DetectFaceRequest) SetImageURL(v string) *DetectFaceRequest {
	s.ImageURL = &v
	return s
}

type DetectFaceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s DetectFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceAdvanceRequest) SetImageURLObject(v io.Reader) *DetectFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectFaceAdvanceRequest) SetImageType(v int32) *DetectFaceAdvanceRequest {
	s.ImageType = &v
	return s
}

type DetectFaceResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetectFaceResponseBody) SetRequestId(v string) *DetectFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectFaceResponseBody) SetData(v *DetectFaceResponseBodyData) *DetectFaceResponseBody {
	s.Data = v
	return s
}

type DetectFaceResponseBodyData struct {
	FaceProbabilityList []*float32                           `json:"FaceProbabilityList,omitempty" xml:"FaceProbabilityList,omitempty" type:"Repeated"`
	Pupils              []*float32                           `json:"Pupils,omitempty" xml:"Pupils,omitempty" type:"Repeated"`
	FaceRectangles      []*int32                             `json:"FaceRectangles,omitempty" xml:"FaceRectangles,omitempty" type:"Repeated"`
	FaceCount           *int32                               `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	PoseList            []*float32                           `json:"PoseList,omitempty" xml:"PoseList,omitempty" type:"Repeated"`
	Landmarks           []*float32                           `json:"Landmarks,omitempty" xml:"Landmarks,omitempty" type:"Repeated"`
	LandmarkCount       *int32                               `json:"LandmarkCount,omitempty" xml:"LandmarkCount,omitempty"`
	Qualities           *DetectFaceResponseBodyDataQualities `json:"Qualities,omitempty" xml:"Qualities,omitempty" type:"Struct"`
}

func (s DetectFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectFaceResponseBodyData) SetFaceProbabilityList(v []*float32) *DetectFaceResponseBodyData {
	s.FaceProbabilityList = v
	return s
}

func (s *DetectFaceResponseBodyData) SetPupils(v []*float32) *DetectFaceResponseBodyData {
	s.Pupils = v
	return s
}

func (s *DetectFaceResponseBodyData) SetFaceRectangles(v []*int32) *DetectFaceResponseBodyData {
	s.FaceRectangles = v
	return s
}

func (s *DetectFaceResponseBodyData) SetFaceCount(v int32) *DetectFaceResponseBodyData {
	s.FaceCount = &v
	return s
}

func (s *DetectFaceResponseBodyData) SetPoseList(v []*float32) *DetectFaceResponseBodyData {
	s.PoseList = v
	return s
}

func (s *DetectFaceResponseBodyData) SetLandmarks(v []*float32) *DetectFaceResponseBodyData {
	s.Landmarks = v
	return s
}

func (s *DetectFaceResponseBodyData) SetLandmarkCount(v int32) *DetectFaceResponseBodyData {
	s.LandmarkCount = &v
	return s
}

func (s *DetectFaceResponseBodyData) SetQualities(v *DetectFaceResponseBodyDataQualities) *DetectFaceResponseBodyData {
	s.Qualities = v
	return s
}

type DetectFaceResponseBodyDataQualities struct {
	ScoreList []*float32 `json:"ScoreList,omitempty" xml:"ScoreList,omitempty" type:"Repeated"`
	BlurList  []*float32 `json:"BlurList,omitempty" xml:"BlurList,omitempty" type:"Repeated"`
	FnfList   []*float32 `json:"FnfList,omitempty" xml:"FnfList,omitempty" type:"Repeated"`
	GlassList []*float32 `json:"GlassList,omitempty" xml:"GlassList,omitempty" type:"Repeated"`
	IlluList  []*float32 `json:"IlluList,omitempty" xml:"IlluList,omitempty" type:"Repeated"`
	MaskList  []*float32 `json:"MaskList,omitempty" xml:"MaskList,omitempty" type:"Repeated"`
	NoiseList []*float32 `json:"NoiseList,omitempty" xml:"NoiseList,omitempty" type:"Repeated"`
	PoseList  []*float32 `json:"PoseList,omitempty" xml:"PoseList,omitempty" type:"Repeated"`
}

func (s DetectFaceResponseBodyDataQualities) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceResponseBodyDataQualities) GoString() string {
	return s.String()
}

func (s *DetectFaceResponseBodyDataQualities) SetScoreList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.ScoreList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetBlurList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.BlurList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetFnfList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.FnfList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetGlassList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.GlassList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetIlluList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.IlluList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetMaskList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.MaskList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetNoiseList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.NoiseList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetPoseList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.PoseList = v
	return s
}

type DetectFaceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceResponse) GoString() string {
	return s.String()
}

func (s *DetectFaceResponse) SetHeaders(v map[string]*string) *DetectFaceResponse {
	s.Headers = v
	return s
}

func (s *DetectFaceResponse) SetBody(v *DetectFaceResponseBody) *DetectFaceResponse {
	s.Body = v
	return s
}

type DetectMaskRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectMaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectMaskRequest) GoString() string {
	return s.String()
}

func (s *DetectMaskRequest) SetImageURL(v string) *DetectMaskRequest {
	s.ImageURL = &v
	return s
}

type DetectMaskAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s DetectMaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectMaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectMaskAdvanceRequest) SetImageURLObject(v io.Reader) *DetectMaskAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectMaskResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectMaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectMaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectMaskResponseBody) GoString() string {
	return s.String()
}

func (s *DetectMaskResponseBody) SetRequestId(v string) *DetectMaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectMaskResponseBody) SetData(v *DetectMaskResponseBodyData) *DetectMaskResponseBody {
	s.Data = v
	return s
}

type DetectMaskResponseBodyData struct {
	Mask            *int32   `json:"Mask,omitempty" xml:"Mask,omitempty"`
	FaceProbability *float32 `json:"FaceProbability,omitempty" xml:"FaceProbability,omitempty"`
}

func (s DetectMaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectMaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectMaskResponseBodyData) SetMask(v int32) *DetectMaskResponseBodyData {
	s.Mask = &v
	return s
}

func (s *DetectMaskResponseBodyData) SetFaceProbability(v float32) *DetectMaskResponseBodyData {
	s.FaceProbability = &v
	return s
}

type DetectMaskResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectMaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectMaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectMaskResponse) GoString() string {
	return s.String()
}

func (s *DetectMaskResponse) SetHeaders(v map[string]*string) *DetectMaskResponse {
	s.Headers = v
	return s
}

func (s *DetectMaskResponse) SetBody(v *DetectMaskResponseBody) *DetectMaskResponse {
	s.Body = v
	return s
}

type GenRealPersonVerificationTokenRequest struct {
	CertificateName   *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	CertificateNumber *string `json:"CertificateNumber,omitempty" xml:"CertificateNumber,omitempty"`
	MetaInfo          *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
}

func (s GenRealPersonVerificationTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GenRealPersonVerificationTokenRequest) GoString() string {
	return s.String()
}

func (s *GenRealPersonVerificationTokenRequest) SetCertificateName(v string) *GenRealPersonVerificationTokenRequest {
	s.CertificateName = &v
	return s
}

func (s *GenRealPersonVerificationTokenRequest) SetCertificateNumber(v string) *GenRealPersonVerificationTokenRequest {
	s.CertificateNumber = &v
	return s
}

func (s *GenRealPersonVerificationTokenRequest) SetMetaInfo(v string) *GenRealPersonVerificationTokenRequest {
	s.MetaInfo = &v
	return s
}

type GenRealPersonVerificationTokenResponseBody struct {
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *GenRealPersonVerificationTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenRealPersonVerificationTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenRealPersonVerificationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenRealPersonVerificationTokenResponseBody) SetRequestId(v string) *GenRealPersonVerificationTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenRealPersonVerificationTokenResponseBody) SetData(v *GenRealPersonVerificationTokenResponseBodyData) *GenRealPersonVerificationTokenResponseBody {
	s.Data = v
	return s
}

func (s *GenRealPersonVerificationTokenResponseBody) SetErrorMessage(v string) *GenRealPersonVerificationTokenResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GenRealPersonVerificationTokenResponseBody) SetCode(v string) *GenRealPersonVerificationTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GenRealPersonVerificationTokenResponseBody) SetSuccess(v bool) *GenRealPersonVerificationTokenResponseBody {
	s.Success = &v
	return s
}

type GenRealPersonVerificationTokenResponseBodyData struct {
	VerificationToken *string `json:"VerificationToken,omitempty" xml:"VerificationToken,omitempty"`
}

func (s GenRealPersonVerificationTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenRealPersonVerificationTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenRealPersonVerificationTokenResponseBodyData) SetVerificationToken(v string) *GenRealPersonVerificationTokenResponseBodyData {
	s.VerificationToken = &v
	return s
}

type GenRealPersonVerificationTokenResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenRealPersonVerificationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenRealPersonVerificationTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GenRealPersonVerificationTokenResponse) GoString() string {
	return s.String()
}

func (s *GenRealPersonVerificationTokenResponse) SetHeaders(v map[string]*string) *GenRealPersonVerificationTokenResponse {
	s.Headers = v
	return s
}

func (s *GenRealPersonVerificationTokenResponse) SetBody(v *GenRealPersonVerificationTokenResponseBody) *GenRealPersonVerificationTokenResponse {
	s.Body = v
	return s
}

type ListFaceDbsResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListFaceDbsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListFaceDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFaceDbsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFaceDbsResponseBody) SetRequestId(v string) *ListFaceDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFaceDbsResponseBody) SetData(v *ListFaceDbsResponseBodyData) *ListFaceDbsResponseBody {
	s.Data = v
	return s
}

type ListFaceDbsResponseBodyData struct {
	DbList []*ListFaceDbsResponseBodyDataDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
}

func (s ListFaceDbsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFaceDbsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFaceDbsResponseBodyData) SetDbList(v []*ListFaceDbsResponseBodyDataDbList) *ListFaceDbsResponseBodyData {
	s.DbList = v
	return s
}

type ListFaceDbsResponseBodyDataDbList struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListFaceDbsResponseBodyDataDbList) String() string {
	return tea.Prettify(s)
}

func (s ListFaceDbsResponseBodyDataDbList) GoString() string {
	return s.String()
}

func (s *ListFaceDbsResponseBodyDataDbList) SetName(v string) *ListFaceDbsResponseBodyDataDbList {
	s.Name = &v
	return s
}

type ListFaceDbsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFaceDbsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFaceDbsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFaceDbsResponse) GoString() string {
	return s.String()
}

func (s *ListFaceDbsResponse) SetHeaders(v map[string]*string) *ListFaceDbsResponse {
	s.Headers = v
	return s
}

func (s *ListFaceDbsResponse) SetBody(v *ListFaceDbsResponseBody) *ListFaceDbsResponse {
	s.Body = v
	return s
}

type RecognizeActionRequest struct {
	Type     *int32                           `json:"Type,omitempty" xml:"Type,omitempty"`
	VideoUrl *string                          `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	URLList  []*RecognizeActionRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s RecognizeActionRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionRequest) GoString() string {
	return s.String()
}

func (s *RecognizeActionRequest) SetType(v int32) *RecognizeActionRequest {
	s.Type = &v
	return s
}

func (s *RecognizeActionRequest) SetVideoUrl(v string) *RecognizeActionRequest {
	s.VideoUrl = &v
	return s
}

func (s *RecognizeActionRequest) SetURLList(v []*RecognizeActionRequestURLList) *RecognizeActionRequest {
	s.URLList = v
	return s
}

type RecognizeActionRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s RecognizeActionRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionRequestURLList) GoString() string {
	return s.String()
}

func (s *RecognizeActionRequestURLList) SetURL(v string) *RecognizeActionRequestURLList {
	s.URL = &v
	return s
}

type RecognizeActionResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeActionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponseBody) SetRequestId(v string) *RecognizeActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeActionResponseBody) SetData(v *RecognizeActionResponseBodyData) *RecognizeActionResponseBody {
	s.Data = v
	return s
}

type RecognizeActionResponseBodyData struct {
	Elements []*RecognizeActionResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizeActionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponseBodyData) SetElements(v []*RecognizeActionResponseBodyDataElements) *RecognizeActionResponseBodyData {
	s.Elements = v
	return s
}

type RecognizeActionResponseBodyDataElements struct {
	Scores    []*float32                                      `json:"Scores,omitempty" xml:"Scores,omitempty" type:"Repeated"`
	Labels    []*string                                       `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Boxes     []*RecognizeActionResponseBodyDataElementsBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Timestamp *int32                                          `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s RecognizeActionResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponseBodyDataElements) SetScores(v []*float32) *RecognizeActionResponseBodyDataElements {
	s.Scores = v
	return s
}

func (s *RecognizeActionResponseBodyDataElements) SetLabels(v []*string) *RecognizeActionResponseBodyDataElements {
	s.Labels = v
	return s
}

func (s *RecognizeActionResponseBodyDataElements) SetBoxes(v []*RecognizeActionResponseBodyDataElementsBoxes) *RecognizeActionResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *RecognizeActionResponseBodyDataElements) SetTimestamp(v int32) *RecognizeActionResponseBodyDataElements {
	s.Timestamp = &v
	return s
}

type RecognizeActionResponseBodyDataElementsBoxes struct {
	Box []*int32 `json:"Box,omitempty" xml:"Box,omitempty" type:"Repeated"`
}

func (s RecognizeActionResponseBodyDataElementsBoxes) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionResponseBodyDataElementsBoxes) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponseBodyDataElementsBoxes) SetBox(v []*int32) *RecognizeActionResponseBodyDataElementsBoxes {
	s.Box = v
	return s
}

type RecognizeActionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeActionResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionResponse) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponse) SetHeaders(v map[string]*string) *RecognizeActionResponse {
	s.Headers = v
	return s
}

func (s *RecognizeActionResponse) SetBody(v *RecognizeActionResponseBody) *RecognizeActionResponse {
	s.Body = v
	return s
}

type DetectChefCapRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectChefCapRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectChefCapRequest) GoString() string {
	return s.String()
}

func (s *DetectChefCapRequest) SetImageURL(v string) *DetectChefCapRequest {
	s.ImageURL = &v
	return s
}

type DetectChefCapAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s DetectChefCapAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectChefCapAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectChefCapAdvanceRequest) SetImageURLObject(v io.Reader) *DetectChefCapAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectChefCapResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectChefCapResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectChefCapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectChefCapResponseBody) GoString() string {
	return s.String()
}

func (s *DetectChefCapResponseBody) SetRequestId(v string) *DetectChefCapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectChefCapResponseBody) SetData(v *DetectChefCapResponseBodyData) *DetectChefCapResponseBody {
	s.Data = v
	return s
}

type DetectChefCapResponseBodyData struct {
	Elements []*DetectChefCapResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectChefCapResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectChefCapResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectChefCapResponseBodyData) SetElements(v []*DetectChefCapResponseBodyDataElements) *DetectChefCapResponseBodyData {
	s.Elements = v
	return s
}

type DetectChefCapResponseBodyDataElements struct {
	Confidence *float32   `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Category   *string    `json:"Category,omitempty" xml:"Category,omitempty"`
	Box        []*float32 `json:"Box,omitempty" xml:"Box,omitempty" type:"Repeated"`
}

func (s DetectChefCapResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectChefCapResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectChefCapResponseBodyDataElements) SetConfidence(v float32) *DetectChefCapResponseBodyDataElements {
	s.Confidence = &v
	return s
}

func (s *DetectChefCapResponseBodyDataElements) SetCategory(v string) *DetectChefCapResponseBodyDataElements {
	s.Category = &v
	return s
}

func (s *DetectChefCapResponseBodyDataElements) SetBox(v []*float32) *DetectChefCapResponseBodyDataElements {
	s.Box = v
	return s
}

type DetectChefCapResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectChefCapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectChefCapResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectChefCapResponse) GoString() string {
	return s.String()
}

func (s *DetectChefCapResponse) SetHeaders(v map[string]*string) *DetectChefCapResponse {
	s.Headers = v
	return s
}

func (s *DetectChefCapResponse) SetBody(v *DetectChefCapResponseBody) *DetectChefCapResponse {
	s.Body = v
	return s
}

type DetectLivingFaceRequest struct {
	Tasks []*DetectLivingFaceRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DetectLivingFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceRequest) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceRequest) SetTasks(v []*DetectLivingFaceRequestTasks) *DetectLivingFaceRequest {
	s.Tasks = v
	return s
}

type DetectLivingFaceRequestTasks struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectLivingFaceRequestTasks) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceRequestTasks) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceRequestTasks) SetImageURL(v string) *DetectLivingFaceRequestTasks {
	s.ImageURL = &v
	return s
}

type DetectLivingFaceResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectLivingFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectLivingFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBody) SetRequestId(v string) *DetectLivingFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectLivingFaceResponseBody) SetData(v *DetectLivingFaceResponseBodyData) *DetectLivingFaceResponseBody {
	s.Data = v
	return s
}

type DetectLivingFaceResponseBodyData struct {
	Elements []*DetectLivingFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectLivingFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyData) SetElements(v []*DetectLivingFaceResponseBodyDataElements) *DetectLivingFaceResponseBodyData {
	s.Elements = v
	return s
}

type DetectLivingFaceResponseBodyDataElements struct {
	ImageURL *string                                            `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	TaskId   *string                                            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Results  []*DetectLivingFaceResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetectLivingFaceResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyDataElements) SetImageURL(v string) *DetectLivingFaceResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElements) SetTaskId(v string) *DetectLivingFaceResponseBodyDataElements {
	s.TaskId = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElements) SetResults(v []*DetectLivingFaceResponseBodyDataElementsResults) *DetectLivingFaceResponseBodyDataElements {
	s.Results = v
	return s
}

type DetectLivingFaceResponseBodyDataElementsResults struct {
	Suggestion *string                                                  `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	Frames     []*DetectLivingFaceResponseBodyDataElementsResultsFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	Label      *string                                                  `json:"Label,omitempty" xml:"Label,omitempty"`
	Rate       *float32                                                 `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s DetectLivingFaceResponseBodyDataElementsResults) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetSuggestion(v string) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Suggestion = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetFrames(v []*DetectLivingFaceResponseBodyDataElementsResultsFrames) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Frames = v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetLabel(v string) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Label = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetRate(v float32) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Rate = &v
	return s
}

type DetectLivingFaceResponseBodyDataElementsResultsFrames struct {
	Url  *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s DetectLivingFaceResponseBodyDataElementsResultsFrames) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceResponseBodyDataElementsResultsFrames) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsFrames) SetUrl(v string) *DetectLivingFaceResponseBodyDataElementsResultsFrames {
	s.Url = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsFrames) SetRate(v float32) *DetectLivingFaceResponseBodyDataElementsResultsFrames {
	s.Rate = &v
	return s
}

type DetectLivingFaceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectLivingFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectLivingFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceResponse) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponse) SetHeaders(v map[string]*string) *DetectLivingFaceResponse {
	s.Headers = v
	return s
}

func (s *DetectLivingFaceResponse) SetBody(v *DetectLivingFaceResponseBody) *DetectLivingFaceResponse {
	s.Body = v
	return s
}

type DetectCelebrityRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectCelebrityRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectCelebrityRequest) GoString() string {
	return s.String()
}

func (s *DetectCelebrityRequest) SetImageURL(v string) *DetectCelebrityRequest {
	s.ImageURL = &v
	return s
}

type DetectCelebrityAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s DetectCelebrityAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectCelebrityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectCelebrityAdvanceRequest) SetImageURLObject(v io.Reader) *DetectCelebrityAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectCelebrityResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectCelebrityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectCelebrityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectCelebrityResponseBody) GoString() string {
	return s.String()
}

func (s *DetectCelebrityResponseBody) SetRequestId(v string) *DetectCelebrityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectCelebrityResponseBody) SetData(v *DetectCelebrityResponseBodyData) *DetectCelebrityResponseBody {
	s.Data = v
	return s
}

type DetectCelebrityResponseBodyData struct {
	FaceRecognizeResults []*DetectCelebrityResponseBodyDataFaceRecognizeResults `json:"FaceRecognizeResults,omitempty" xml:"FaceRecognizeResults,omitempty" type:"Repeated"`
	Width                *int32                                                 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height               *int32                                                 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s DetectCelebrityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectCelebrityResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectCelebrityResponseBodyData) SetFaceRecognizeResults(v []*DetectCelebrityResponseBodyDataFaceRecognizeResults) *DetectCelebrityResponseBodyData {
	s.FaceRecognizeResults = v
	return s
}

func (s *DetectCelebrityResponseBodyData) SetWidth(v int32) *DetectCelebrityResponseBodyData {
	s.Width = &v
	return s
}

func (s *DetectCelebrityResponseBodyData) SetHeight(v int32) *DetectCelebrityResponseBodyData {
	s.Height = &v
	return s
}

type DetectCelebrityResponseBodyDataFaceRecognizeResults struct {
	FaceBoxes []*float32 `json:"FaceBoxes,omitempty" xml:"FaceBoxes,omitempty" type:"Repeated"`
	Name      *string    `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DetectCelebrityResponseBodyDataFaceRecognizeResults) String() string {
	return tea.Prettify(s)
}

func (s DetectCelebrityResponseBodyDataFaceRecognizeResults) GoString() string {
	return s.String()
}

func (s *DetectCelebrityResponseBodyDataFaceRecognizeResults) SetFaceBoxes(v []*float32) *DetectCelebrityResponseBodyDataFaceRecognizeResults {
	s.FaceBoxes = v
	return s
}

func (s *DetectCelebrityResponseBodyDataFaceRecognizeResults) SetName(v string) *DetectCelebrityResponseBodyDataFaceRecognizeResults {
	s.Name = &v
	return s
}

type DetectCelebrityResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectCelebrityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectCelebrityResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectCelebrityResponse) GoString() string {
	return s.String()
}

func (s *DetectCelebrityResponse) SetHeaders(v map[string]*string) *DetectCelebrityResponse {
	s.Headers = v
	return s
}

func (s *DetectCelebrityResponse) SetBody(v *DetectCelebrityResponseBody) *DetectCelebrityResponse {
	s.Body = v
	return s
}

type GetRealPersonVerificationResultRequest struct {
	VerificationToken *string `json:"VerificationToken,omitempty" xml:"VerificationToken,omitempty"`
	MaterialHash      *string `json:"MaterialHash,omitempty" xml:"MaterialHash,omitempty"`
}

func (s GetRealPersonVerificationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRealPersonVerificationResultRequest) GoString() string {
	return s.String()
}

func (s *GetRealPersonVerificationResultRequest) SetVerificationToken(v string) *GetRealPersonVerificationResultRequest {
	s.VerificationToken = &v
	return s
}

func (s *GetRealPersonVerificationResultRequest) SetMaterialHash(v string) *GetRealPersonVerificationResultRequest {
	s.MaterialHash = &v
	return s
}

type GetRealPersonVerificationResultResponseBody struct {
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *GetRealPersonVerificationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRealPersonVerificationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRealPersonVerificationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealPersonVerificationResultResponseBody) SetRequestId(v string) *GetRealPersonVerificationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealPersonVerificationResultResponseBody) SetData(v *GetRealPersonVerificationResultResponseBodyData) *GetRealPersonVerificationResultResponseBody {
	s.Data = v
	return s
}

func (s *GetRealPersonVerificationResultResponseBody) SetErrorMessage(v string) *GetRealPersonVerificationResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRealPersonVerificationResultResponseBody) SetCode(v string) *GetRealPersonVerificationResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetRealPersonVerificationResultResponseBody) SetSuccess(v bool) *GetRealPersonVerificationResultResponseBody {
	s.Success = &v
	return s
}

type GetRealPersonVerificationResultResponseBodyData struct {
	Pass          *bool   `json:"Pass,omitempty" xml:"Pass,omitempty"`
	IdentityInfo  *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty"`
	MaterialMatch *string `json:"MaterialMatch,omitempty" xml:"MaterialMatch,omitempty"`
}

func (s GetRealPersonVerificationResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRealPersonVerificationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRealPersonVerificationResultResponseBodyData) SetPass(v bool) *GetRealPersonVerificationResultResponseBodyData {
	s.Pass = &v
	return s
}

func (s *GetRealPersonVerificationResultResponseBodyData) SetIdentityInfo(v string) *GetRealPersonVerificationResultResponseBodyData {
	s.IdentityInfo = &v
	return s
}

func (s *GetRealPersonVerificationResultResponseBodyData) SetMaterialMatch(v string) *GetRealPersonVerificationResultResponseBodyData {
	s.MaterialMatch = &v
	return s
}

type GetRealPersonVerificationResultResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRealPersonVerificationResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRealPersonVerificationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRealPersonVerificationResultResponse) GoString() string {
	return s.String()
}

func (s *GetRealPersonVerificationResultResponse) SetHeaders(v map[string]*string) *GetRealPersonVerificationResultResponse {
	s.Headers = v
	return s
}

func (s *GetRealPersonVerificationResultResponse) SetBody(v *GetRealPersonVerificationResultResponseBody) *GetRealPersonVerificationResultResponse {
	s.Body = v
	return s
}

type DeleteFaceRequest struct {
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
}

func (s DeleteFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceRequest) SetDbName(v string) *DeleteFaceRequest {
	s.DbName = &v
	return s
}

func (s *DeleteFaceRequest) SetFaceId(v string) *DeleteFaceRequest {
	s.FaceId = &v
	return s
}

type DeleteFaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceResponseBody) SetRequestId(v string) *DeleteFaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFaceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceResponse) SetHeaders(v map[string]*string) *DeleteFaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceResponse) SetBody(v *DeleteFaceResponseBody) *DeleteFaceResponse {
	s.Body = v
	return s
}

type ExtractPedestrianFeatureAttributeRequest struct {
	Mode     *string                                            `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ImageURL *string                                            `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	UrlList  []*ExtractPedestrianFeatureAttributeRequestUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s ExtractPedestrianFeatureAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttributeRequest) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttributeRequest) SetMode(v string) *ExtractPedestrianFeatureAttributeRequest {
	s.Mode = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeRequest) SetImageURL(v string) *ExtractPedestrianFeatureAttributeRequest {
	s.ImageURL = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeRequest) SetUrlList(v []*ExtractPedestrianFeatureAttributeRequestUrlList) *ExtractPedestrianFeatureAttributeRequest {
	s.UrlList = v
	return s
}

type ExtractPedestrianFeatureAttributeRequestUrlList struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ExtractPedestrianFeatureAttributeRequestUrlList) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttributeRequestUrlList) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttributeRequestUrlList) SetUrl(v string) *ExtractPedestrianFeatureAttributeRequestUrlList {
	s.Url = &v
	return s
}

type ExtractPedestrianFeatureAttributeResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ExtractPedestrianFeatureAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ExtractPedestrianFeatureAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttributeResponseBody) SetRequestId(v string) *ExtractPedestrianFeatureAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBody) SetData(v *ExtractPedestrianFeatureAttributeResponseBodyData) *ExtractPedestrianFeatureAttributeResponseBody {
	s.Data = v
	return s
}

type ExtractPedestrianFeatureAttributeResponseBodyData struct {
	QualityScore    *float32                                                     `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
	ObjType         *string                                                      `json:"ObjType,omitempty" xml:"ObjType,omitempty"`
	Elements        []*ExtractPedestrianFeatureAttributeResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Feature         *string                                                      `json:"Feature,omitempty" xml:"Feature,omitempty"`
	Gender          *string                                                      `json:"Gender,omitempty" xml:"Gender,omitempty"`
	LowerColorScore *float32                                                     `json:"LowerColorScore,omitempty" xml:"LowerColorScore,omitempty"`
	ObjTypeScore    *float32                                                     `json:"ObjTypeScore,omitempty" xml:"ObjTypeScore,omitempty"`
	Age             *string                                                      `json:"Age,omitempty" xml:"Age,omitempty"`
	AgeScore        *float32                                                     `json:"AgeScore,omitempty" xml:"AgeScore,omitempty"`
	UpperTypeScore  *float32                                                     `json:"UpperTypeScore,omitempty" xml:"UpperTypeScore,omitempty"`
	LowerTypeScore  *float32                                                     `json:"LowerTypeScore,omitempty" xml:"LowerTypeScore,omitempty"`
	LowerColor      *string                                                      `json:"LowerColor,omitempty" xml:"LowerColor,omitempty"`
	Hair            *string                                                      `json:"Hair,omitempty" xml:"Hair,omitempty"`
	UpperColor      *string                                                      `json:"UpperColor,omitempty" xml:"UpperColor,omitempty"`
	GenderScore     *float32                                                     `json:"GenderScore,omitempty" xml:"GenderScore,omitempty"`
	UpperType       *string                                                      `json:"UpperType,omitempty" xml:"UpperType,omitempty"`
	HairScore       *float32                                                     `json:"HairScore,omitempty" xml:"HairScore,omitempty"`
	LowerType       *string                                                      `json:"LowerType,omitempty" xml:"LowerType,omitempty"`
	UpperColorScore *float32                                                     `json:"UpperColorScore,omitempty" xml:"UpperColorScore,omitempty"`
}

func (s ExtractPedestrianFeatureAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetQualityScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.QualityScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetObjType(v string) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.ObjType = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetElements(v []*ExtractPedestrianFeatureAttributeResponseBodyDataElements) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.Elements = v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetFeature(v string) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.Feature = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetGender(v string) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.Gender = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetLowerColorScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.LowerColorScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetObjTypeScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.ObjTypeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetAge(v string) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.Age = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetAgeScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.AgeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetUpperTypeScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.UpperTypeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetLowerTypeScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.LowerTypeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetLowerColor(v string) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.LowerColor = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetHair(v string) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.Hair = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetUpperColor(v string) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.UpperColor = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetGenderScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.GenderScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetUpperType(v string) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.UpperType = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetHairScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.HairScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetLowerType(v string) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.LowerType = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyData) SetUpperColorScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyData {
	s.UpperColorScore = &v
	return s
}

type ExtractPedestrianFeatureAttributeResponseBodyDataElements struct {
	QualityScore    *float32 `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
	ObjType         *string  `json:"ObjType,omitempty" xml:"ObjType,omitempty"`
	Feature         *string  `json:"Feature,omitempty" xml:"Feature,omitempty"`
	Gender          *string  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	LowerColorScore *float32 `json:"LowerColorScore,omitempty" xml:"LowerColorScore,omitempty"`
	ObjTypeScore    *float32 `json:"ObjTypeScore,omitempty" xml:"ObjTypeScore,omitempty"`
	Age             *string  `json:"Age,omitempty" xml:"Age,omitempty"`
	AgeScore        *float32 `json:"AgeScore,omitempty" xml:"AgeScore,omitempty"`
	UpperTypeScore  *float32 `json:"UpperTypeScore,omitempty" xml:"UpperTypeScore,omitempty"`
	LowerTypeScore  *float32 `json:"LowerTypeScore,omitempty" xml:"LowerTypeScore,omitempty"`
	LowerColor      *string  `json:"LowerColor,omitempty" xml:"LowerColor,omitempty"`
	Hair            *string  `json:"Hair,omitempty" xml:"Hair,omitempty"`
	UpperColor      *string  `json:"UpperColor,omitempty" xml:"UpperColor,omitempty"`
	GenderScore     *float32 `json:"GenderScore,omitempty" xml:"GenderScore,omitempty"`
	UpperType       *string  `json:"UpperType,omitempty" xml:"UpperType,omitempty"`
	HairScore       *float32 `json:"HairScore,omitempty" xml:"HairScore,omitempty"`
	LowerType       *string  `json:"LowerType,omitempty" xml:"LowerType,omitempty"`
	UpperColorScore *float32 `json:"UpperColorScore,omitempty" xml:"UpperColorScore,omitempty"`
}

func (s ExtractPedestrianFeatureAttributeResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttributeResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetQualityScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.QualityScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetObjType(v string) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.ObjType = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetFeature(v string) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.Feature = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetGender(v string) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.Gender = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetLowerColorScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.LowerColorScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetObjTypeScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.ObjTypeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetAge(v string) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.Age = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetAgeScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.AgeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetUpperTypeScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.UpperTypeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetLowerTypeScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.LowerTypeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetLowerColor(v string) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.LowerColor = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetHair(v string) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.Hair = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetUpperColor(v string) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.UpperColor = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetGenderScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.GenderScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetUpperType(v string) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.UpperType = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetHairScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.HairScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetLowerType(v string) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.LowerType = &v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponseBodyDataElements) SetUpperColorScore(v float32) *ExtractPedestrianFeatureAttributeResponseBodyDataElements {
	s.UpperColorScore = &v
	return s
}

type ExtractPedestrianFeatureAttributeResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExtractPedestrianFeatureAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExtractPedestrianFeatureAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttributeResponse) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttributeResponse) SetHeaders(v map[string]*string) *ExtractPedestrianFeatureAttributeResponse {
	s.Headers = v
	return s
}

func (s *ExtractPedestrianFeatureAttributeResponse) SetBody(v *ExtractPedestrianFeatureAttributeResponseBody) *ExtractPedestrianFeatureAttributeResponse {
	s.Body = v
	return s
}

type RecognizeExpressionRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeExpressionRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionRequest) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionRequest) SetImageURL(v string) *RecognizeExpressionRequest {
	s.ImageURL = &v
	return s
}

type RecognizeExpressionAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s RecognizeExpressionAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeExpressionAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeExpressionResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeExpressionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeExpressionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponseBody) SetRequestId(v string) *RecognizeExpressionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeExpressionResponseBody) SetData(v *RecognizeExpressionResponseBodyData) *RecognizeExpressionResponseBody {
	s.Data = v
	return s
}

type RecognizeExpressionResponseBodyData struct {
	Elements []*RecognizeExpressionResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizeExpressionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponseBodyData) SetElements(v []*RecognizeExpressionResponseBodyDataElements) *RecognizeExpressionResponseBodyData {
	s.Elements = v
	return s
}

type RecognizeExpressionResponseBodyDataElements struct {
	FaceRectangle   *RecognizeExpressionResponseBodyDataElementsFaceRectangle `json:"FaceRectangle,omitempty" xml:"FaceRectangle,omitempty" type:"Struct"`
	Expression      *string                                                   `json:"Expression,omitempty" xml:"Expression,omitempty"`
	FaceProbability *float32                                                  `json:"FaceProbability,omitempty" xml:"FaceProbability,omitempty"`
}

func (s RecognizeExpressionResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponseBodyDataElements) SetFaceRectangle(v *RecognizeExpressionResponseBodyDataElementsFaceRectangle) *RecognizeExpressionResponseBodyDataElements {
	s.FaceRectangle = v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElements) SetExpression(v string) *RecognizeExpressionResponseBodyDataElements {
	s.Expression = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElements) SetFaceProbability(v float32) *RecognizeExpressionResponseBodyDataElements {
	s.FaceProbability = &v
	return s
}

type RecognizeExpressionResponseBodyDataElementsFaceRectangle struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s RecognizeExpressionResponseBodyDataElementsFaceRectangle) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionResponseBodyDataElementsFaceRectangle) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) SetTop(v int32) *RecognizeExpressionResponseBodyDataElementsFaceRectangle {
	s.Top = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) SetWidth(v int32) *RecognizeExpressionResponseBodyDataElementsFaceRectangle {
	s.Width = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) SetHeight(v int32) *RecognizeExpressionResponseBodyDataElementsFaceRectangle {
	s.Height = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) SetLeft(v int32) *RecognizeExpressionResponseBodyDataElementsFaceRectangle {
	s.Left = &v
	return s
}

type RecognizeExpressionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeExpressionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeExpressionResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionResponse) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponse) SetHeaders(v map[string]*string) *RecognizeExpressionResponse {
	s.Headers = v
	return s
}

func (s *RecognizeExpressionResponse) SetBody(v *RecognizeExpressionResponseBody) *RecognizeExpressionResponse {
	s.Body = v
	return s
}

type MergeImageFaceRequest struct {
	// A short description of struct
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s MergeImageFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeImageFaceRequest) GoString() string {
	return s.String()
}

func (s *MergeImageFaceRequest) SetUserId(v string) *MergeImageFaceRequest {
	s.UserId = &v
	return s
}

func (s *MergeImageFaceRequest) SetTemplateId(v string) *MergeImageFaceRequest {
	s.TemplateId = &v
	return s
}

func (s *MergeImageFaceRequest) SetImageURL(v string) *MergeImageFaceRequest {
	s.ImageURL = &v
	return s
}

type MergeImageFaceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	// A short description of struct
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s MergeImageFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeImageFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *MergeImageFaceAdvanceRequest) SetImageURLObject(v io.Reader) *MergeImageFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *MergeImageFaceAdvanceRequest) SetUserId(v string) *MergeImageFaceAdvanceRequest {
	s.UserId = &v
	return s
}

func (s *MergeImageFaceAdvanceRequest) SetTemplateId(v string) *MergeImageFaceAdvanceRequest {
	s.TemplateId = &v
	return s
}

type MergeImageFaceResponseBody struct {
	// Id of the request
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *MergeImageFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s MergeImageFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MergeImageFaceResponseBody) GoString() string {
	return s.String()
}

func (s *MergeImageFaceResponseBody) SetRequestId(v string) *MergeImageFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MergeImageFaceResponseBody) SetData(v *MergeImageFaceResponseBodyData) *MergeImageFaceResponseBody {
	s.Data = v
	return s
}

type MergeImageFaceResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s MergeImageFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MergeImageFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *MergeImageFaceResponseBodyData) SetImageURL(v string) *MergeImageFaceResponseBodyData {
	s.ImageURL = &v
	return s
}

type MergeImageFaceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MergeImageFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MergeImageFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s MergeImageFaceResponse) GoString() string {
	return s.String()
}

func (s *MergeImageFaceResponse) SetHeaders(v map[string]*string) *MergeImageFaceResponse {
	s.Headers = v
	return s
}

func (s *MergeImageFaceResponse) SetBody(v *MergeImageFaceResponseBody) *MergeImageFaceResponse {
	s.Body = v
	return s
}

type DeleteBodyPersonRequest struct {
	// 数据库ID
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// 人员ID
	PersonId *int64 `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s DeleteBodyPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBodyPersonRequest) GoString() string {
	return s.String()
}

func (s *DeleteBodyPersonRequest) SetDbId(v int64) *DeleteBodyPersonRequest {
	s.DbId = &v
	return s
}

func (s *DeleteBodyPersonRequest) SetPersonId(v int64) *DeleteBodyPersonRequest {
	s.PersonId = &v
	return s
}

type DeleteBodyPersonResponseBody struct {
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBodyPersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBodyPersonResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBodyPersonResponseBody) SetRequestId(v string) *DeleteBodyPersonResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBodyPersonResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBodyPersonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBodyPersonResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBodyPersonResponse) GoString() string {
	return s.String()
}

func (s *DeleteBodyPersonResponse) SetHeaders(v map[string]*string) *DeleteBodyPersonResponse {
	s.Headers = v
	return s
}

func (s *DeleteBodyPersonResponse) SetBody(v *DeleteBodyPersonResponseBody) *DeleteBodyPersonResponse {
	s.Body = v
	return s
}

type DetectPedestrianRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectPedestrianRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianRequest) GoString() string {
	return s.String()
}

func (s *DetectPedestrianRequest) SetImageURL(v string) *DetectPedestrianRequest {
	s.ImageURL = &v
	return s
}

type DetectPedestrianAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s DetectPedestrianAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectPedestrianAdvanceRequest) SetImageURLObject(v io.Reader) *DetectPedestrianAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectPedestrianResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectPedestrianResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectPedestrianResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianResponseBody) GoString() string {
	return s.String()
}

func (s *DetectPedestrianResponseBody) SetRequestId(v string) *DetectPedestrianResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectPedestrianResponseBody) SetData(v *DetectPedestrianResponseBodyData) *DetectPedestrianResponseBody {
	s.Data = v
	return s
}

type DetectPedestrianResponseBodyData struct {
	Elements []*DetectPedestrianResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Width    *int32                                      `json:"Width,omitempty" xml:"Width,omitempty"`
	Height   *int32                                      `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s DetectPedestrianResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectPedestrianResponseBodyData) SetElements(v []*DetectPedestrianResponseBodyDataElements) *DetectPedestrianResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectPedestrianResponseBodyData) SetWidth(v int32) *DetectPedestrianResponseBodyData {
	s.Width = &v
	return s
}

func (s *DetectPedestrianResponseBodyData) SetHeight(v int32) *DetectPedestrianResponseBodyData {
	s.Height = &v
	return s
}

type DetectPedestrianResponseBodyDataElements struct {
	Type  *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s DetectPedestrianResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectPedestrianResponseBodyDataElements) SetType(v string) *DetectPedestrianResponseBodyDataElements {
	s.Type = &v
	return s
}

func (s *DetectPedestrianResponseBodyDataElements) SetBoxes(v []*int32) *DetectPedestrianResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *DetectPedestrianResponseBodyDataElements) SetScore(v float32) *DetectPedestrianResponseBodyDataElements {
	s.Score = &v
	return s
}

type DetectPedestrianResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectPedestrianResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectPedestrianResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianResponse) GoString() string {
	return s.String()
}

func (s *DetectPedestrianResponse) SetHeaders(v map[string]*string) *DetectPedestrianResponse {
	s.Headers = v
	return s
}

func (s *DetectPedestrianResponse) SetBody(v *DetectPedestrianResponseBody) *DetectPedestrianResponse {
	s.Body = v
	return s
}

type SwapFacialFeaturesRequest struct {
	SourceImageURL *string `json:"SourceImageURL,omitempty" xml:"SourceImageURL,omitempty"`
	EditPart       *string `json:"EditPart,omitempty" xml:"EditPart,omitempty"`
	TargetImageURL *string `json:"TargetImageURL,omitempty" xml:"TargetImageURL,omitempty"`
}

func (s SwapFacialFeaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s SwapFacialFeaturesRequest) GoString() string {
	return s.String()
}

func (s *SwapFacialFeaturesRequest) SetSourceImageURL(v string) *SwapFacialFeaturesRequest {
	s.SourceImageURL = &v
	return s
}

func (s *SwapFacialFeaturesRequest) SetEditPart(v string) *SwapFacialFeaturesRequest {
	s.EditPart = &v
	return s
}

func (s *SwapFacialFeaturesRequest) SetTargetImageURL(v string) *SwapFacialFeaturesRequest {
	s.TargetImageURL = &v
	return s
}

type SwapFacialFeaturesAdvanceRequest struct {
	SourceImageURLObject io.Reader `json:"SourceImageURLObject,omitempty" xml:"SourceImageURLObject,omitempty" require:"true"`
	EditPart             *string   `json:"EditPart,omitempty" xml:"EditPart,omitempty"`
	TargetImageURL       *string   `json:"TargetImageURL,omitempty" xml:"TargetImageURL,omitempty"`
}

func (s SwapFacialFeaturesAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SwapFacialFeaturesAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SwapFacialFeaturesAdvanceRequest) SetSourceImageURLObject(v io.Reader) *SwapFacialFeaturesAdvanceRequest {
	s.SourceImageURLObject = v
	return s
}

func (s *SwapFacialFeaturesAdvanceRequest) SetEditPart(v string) *SwapFacialFeaturesAdvanceRequest {
	s.EditPart = &v
	return s
}

func (s *SwapFacialFeaturesAdvanceRequest) SetTargetImageURL(v string) *SwapFacialFeaturesAdvanceRequest {
	s.TargetImageURL = &v
	return s
}

type SwapFacialFeaturesResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SwapFacialFeaturesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s SwapFacialFeaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwapFacialFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *SwapFacialFeaturesResponseBody) SetRequestId(v string) *SwapFacialFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwapFacialFeaturesResponseBody) SetData(v *SwapFacialFeaturesResponseBodyData) *SwapFacialFeaturesResponseBody {
	s.Data = v
	return s
}

type SwapFacialFeaturesResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SwapFacialFeaturesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SwapFacialFeaturesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SwapFacialFeaturesResponseBodyData) SetImageURL(v string) *SwapFacialFeaturesResponseBodyData {
	s.ImageURL = &v
	return s
}

type SwapFacialFeaturesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SwapFacialFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwapFacialFeaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s SwapFacialFeaturesResponse) GoString() string {
	return s.String()
}

func (s *SwapFacialFeaturesResponse) SetHeaders(v map[string]*string) *SwapFacialFeaturesResponse {
	s.Headers = v
	return s
}

func (s *SwapFacialFeaturesResponse) SetBody(v *SwapFacialFeaturesResponseBody) *SwapFacialFeaturesResponse {
	s.Body = v
	return s
}

type SearchFaceRequest struct {
	DbName   *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Limit    *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	DbNames  *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
}

func (s SearchFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceRequest) GoString() string {
	return s.String()
}

func (s *SearchFaceRequest) SetDbName(v string) *SearchFaceRequest {
	s.DbName = &v
	return s
}

func (s *SearchFaceRequest) SetImageUrl(v string) *SearchFaceRequest {
	s.ImageUrl = &v
	return s
}

func (s *SearchFaceRequest) SetLimit(v int32) *SearchFaceRequest {
	s.Limit = &v
	return s
}

func (s *SearchFaceRequest) SetDbNames(v string) *SearchFaceRequest {
	s.DbNames = &v
	return s
}

type SearchFaceAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	DbName         *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Limit          *int32    `json:"Limit,omitempty" xml:"Limit,omitempty"`
	DbNames        *string   `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
}

func (s SearchFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SearchFaceAdvanceRequest) SetImageUrlObject(v io.Reader) *SearchFaceAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *SearchFaceAdvanceRequest) SetDbName(v string) *SearchFaceAdvanceRequest {
	s.DbName = &v
	return s
}

func (s *SearchFaceAdvanceRequest) SetLimit(v int32) *SearchFaceAdvanceRequest {
	s.Limit = &v
	return s
}

func (s *SearchFaceAdvanceRequest) SetDbNames(v string) *SearchFaceAdvanceRequest {
	s.DbNames = &v
	return s
}

type SearchFaceResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SearchFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s SearchFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBody) SetRequestId(v string) *SearchFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchFaceResponseBody) SetData(v *SearchFaceResponseBodyData) *SearchFaceResponseBody {
	s.Data = v
	return s
}

type SearchFaceResponseBodyData struct {
	MatchList []*SearchFaceResponseBodyDataMatchList `json:"MatchList,omitempty" xml:"MatchList,omitempty" type:"Repeated"`
}

func (s SearchFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyData) SetMatchList(v []*SearchFaceResponseBodyDataMatchList) *SearchFaceResponseBodyData {
	s.MatchList = v
	return s
}

type SearchFaceResponseBodyDataMatchList struct {
	FaceItems []*SearchFaceResponseBodyDataMatchListFaceItems `json:"FaceItems,omitempty" xml:"FaceItems,omitempty" type:"Repeated"`
	Location  *SearchFaceResponseBodyDataMatchListLocation    `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
}

func (s SearchFaceResponseBodyDataMatchList) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBodyDataMatchList) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyDataMatchList) SetFaceItems(v []*SearchFaceResponseBodyDataMatchListFaceItems) *SearchFaceResponseBodyDataMatchList {
	s.FaceItems = v
	return s
}

func (s *SearchFaceResponseBodyDataMatchList) SetLocation(v *SearchFaceResponseBodyDataMatchListLocation) *SearchFaceResponseBodyDataMatchList {
	s.Location = v
	return s
}

type SearchFaceResponseBodyDataMatchListFaceItems struct {
	EntityId  *string  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	FaceId    *string  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	Score     *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	ExtraData *string  `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	DbName    *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s SearchFaceResponseBodyDataMatchListFaceItems) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBodyDataMatchListFaceItems) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetEntityId(v string) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.EntityId = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetFaceId(v string) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.FaceId = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetScore(v float32) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.Score = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetExtraData(v string) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.ExtraData = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetDbName(v string) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.DbName = &v
	return s
}

type SearchFaceResponseBodyDataMatchListLocation struct {
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y      *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X      *int32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s SearchFaceResponseBodyDataMatchListLocation) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBodyDataMatchListLocation) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyDataMatchListLocation) SetWidth(v int32) *SearchFaceResponseBodyDataMatchListLocation {
	s.Width = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListLocation) SetHeight(v int32) *SearchFaceResponseBodyDataMatchListLocation {
	s.Height = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListLocation) SetY(v int32) *SearchFaceResponseBodyDataMatchListLocation {
	s.Y = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListLocation) SetX(v int32) *SearchFaceResponseBodyDataMatchListLocation {
	s.X = &v
	return s
}

type SearchFaceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponse) GoString() string {
	return s.String()
}

func (s *SearchFaceResponse) SetHeaders(v map[string]*string) *SearchFaceResponse {
	s.Headers = v
	return s
}

func (s *SearchFaceResponse) SetBody(v *SearchFaceResponseBody) *SearchFaceResponse {
	s.Body = v
	return s
}

type UpdateFaceEntityRequest struct {
	DbName   *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Labels   *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s UpdateFaceEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceEntityRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaceEntityRequest) SetDbName(v string) *UpdateFaceEntityRequest {
	s.DbName = &v
	return s
}

func (s *UpdateFaceEntityRequest) SetEntityId(v string) *UpdateFaceEntityRequest {
	s.EntityId = &v
	return s
}

func (s *UpdateFaceEntityRequest) SetLabels(v string) *UpdateFaceEntityRequest {
	s.Labels = &v
	return s
}

type UpdateFaceEntityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFaceEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFaceEntityResponseBody) SetRequestId(v string) *UpdateFaceEntityResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFaceEntityResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFaceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFaceEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceEntityResponse) GoString() string {
	return s.String()
}

func (s *UpdateFaceEntityResponse) SetHeaders(v map[string]*string) *UpdateFaceEntityResponse {
	s.Headers = v
	return s
}

func (s *UpdateFaceEntityResponse) SetBody(v *UpdateFaceEntityResponseBody) *UpdateFaceEntityResponse {
	s.Body = v
	return s
}

type BlurFaceRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BlurFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s BlurFaceRequest) GoString() string {
	return s.String()
}

func (s *BlurFaceRequest) SetImageURL(v string) *BlurFaceRequest {
	s.ImageURL = &v
	return s
}

type BlurFaceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s BlurFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s BlurFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BlurFaceAdvanceRequest) SetImageURLObject(v io.Reader) *BlurFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type BlurFaceResponseBody struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *BlurFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s BlurFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BlurFaceResponseBody) GoString() string {
	return s.String()
}

func (s *BlurFaceResponseBody) SetRequestId(v string) *BlurFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BlurFaceResponseBody) SetData(v *BlurFaceResponseBodyData) *BlurFaceResponseBody {
	s.Data = v
	return s
}

type BlurFaceResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BlurFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BlurFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *BlurFaceResponseBodyData) SetImageURL(v string) *BlurFaceResponseBodyData {
	s.ImageURL = &v
	return s
}

type BlurFaceResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BlurFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BlurFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s BlurFaceResponse) GoString() string {
	return s.String()
}

func (s *BlurFaceResponse) SetHeaders(v map[string]*string) *BlurFaceResponse {
	s.Headers = v
	return s
}

func (s *BlurFaceResponse) SetBody(v *BlurFaceResponseBody) *BlurFaceResponse {
	s.Body = v
	return s
}

type FaceMakeupRequest struct {
	ImageURL     *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	MakeupType   *string  `json:"MakeupType,omitempty" xml:"MakeupType,omitempty"`
	ResourceType *string  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Strength     *float32 `json:"Strength,omitempty" xml:"Strength,omitempty"`
}

func (s FaceMakeupRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceMakeupRequest) GoString() string {
	return s.String()
}

func (s *FaceMakeupRequest) SetImageURL(v string) *FaceMakeupRequest {
	s.ImageURL = &v
	return s
}

func (s *FaceMakeupRequest) SetMakeupType(v string) *FaceMakeupRequest {
	s.MakeupType = &v
	return s
}

func (s *FaceMakeupRequest) SetResourceType(v string) *FaceMakeupRequest {
	s.ResourceType = &v
	return s
}

func (s *FaceMakeupRequest) SetStrength(v float32) *FaceMakeupRequest {
	s.Strength = &v
	return s
}

type FaceMakeupAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	MakeupType     *string   `json:"MakeupType,omitempty" xml:"MakeupType,omitempty"`
	ResourceType   *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Strength       *float32  `json:"Strength,omitempty" xml:"Strength,omitempty"`
}

func (s FaceMakeupAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceMakeupAdvanceRequest) GoString() string {
	return s.String()
}

func (s *FaceMakeupAdvanceRequest) SetImageURLObject(v io.Reader) *FaceMakeupAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *FaceMakeupAdvanceRequest) SetMakeupType(v string) *FaceMakeupAdvanceRequest {
	s.MakeupType = &v
	return s
}

func (s *FaceMakeupAdvanceRequest) SetResourceType(v string) *FaceMakeupAdvanceRequest {
	s.ResourceType = &v
	return s
}

func (s *FaceMakeupAdvanceRequest) SetStrength(v float32) *FaceMakeupAdvanceRequest {
	s.Strength = &v
	return s
}

type FaceMakeupResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *FaceMakeupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s FaceMakeupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceMakeupResponseBody) GoString() string {
	return s.String()
}

func (s *FaceMakeupResponseBody) SetRequestId(v string) *FaceMakeupResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceMakeupResponseBody) SetData(v *FaceMakeupResponseBodyData) *FaceMakeupResponseBody {
	s.Data = v
	return s
}

type FaceMakeupResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s FaceMakeupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FaceMakeupResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceMakeupResponseBodyData) SetImageURL(v string) *FaceMakeupResponseBodyData {
	s.ImageURL = &v
	return s
}

type FaceMakeupResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FaceMakeupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FaceMakeupResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceMakeupResponse) GoString() string {
	return s.String()
}

func (s *FaceMakeupResponse) SetHeaders(v map[string]*string) *FaceMakeupResponse {
	s.Headers = v
	return s
}

func (s *FaceMakeupResponse) SetBody(v *FaceMakeupResponseBody) *FaceMakeupResponse {
	s.Body = v
	return s
}

type CreateBodyPersonRequest struct {
	// 数据库ID
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// 实体ID，可以包含数字、字母和下划线，相同db下不可以重复，长度1-64。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateBodyPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBodyPersonRequest) GoString() string {
	return s.String()
}

func (s *CreateBodyPersonRequest) SetDbId(v int64) *CreateBodyPersonRequest {
	s.DbId = &v
	return s
}

func (s *CreateBodyPersonRequest) SetName(v string) *CreateBodyPersonRequest {
	s.Name = &v
	return s
}

type CreateBodyPersonResponseBody struct {
	// RequestId
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateBodyPersonResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CreateBodyPersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBodyPersonResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBodyPersonResponseBody) SetRequestId(v string) *CreateBodyPersonResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBodyPersonResponseBody) SetData(v *CreateBodyPersonResponseBodyData) *CreateBodyPersonResponseBody {
	s.Data = v
	return s
}

type CreateBodyPersonResponseBodyData struct {
	// 数据库ID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateBodyPersonResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateBodyPersonResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateBodyPersonResponseBodyData) SetId(v int64) *CreateBodyPersonResponseBodyData {
	s.Id = &v
	return s
}

type CreateBodyPersonResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBodyPersonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBodyPersonResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBodyPersonResponse) GoString() string {
	return s.String()
}

func (s *CreateBodyPersonResponse) SetHeaders(v map[string]*string) *CreateBodyPersonResponse {
	s.Headers = v
	return s
}

func (s *CreateBodyPersonResponse) SetBody(v *CreateBodyPersonResponseBody) *CreateBodyPersonResponse {
	s.Body = v
	return s
}

type AddFaceRequest struct {
	DbName    *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ImageUrl  *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
}

func (s AddFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceRequest) GoString() string {
	return s.String()
}

func (s *AddFaceRequest) SetDbName(v string) *AddFaceRequest {
	s.DbName = &v
	return s
}

func (s *AddFaceRequest) SetImageUrl(v string) *AddFaceRequest {
	s.ImageUrl = &v
	return s
}

func (s *AddFaceRequest) SetEntityId(v string) *AddFaceRequest {
	s.EntityId = &v
	return s
}

func (s *AddFaceRequest) SetExtraData(v string) *AddFaceRequest {
	s.ExtraData = &v
	return s
}

type AddFaceAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	DbName         *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId       *string   `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	ExtraData      *string   `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
}

func (s AddFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddFaceAdvanceRequest) SetImageUrlObject(v io.Reader) *AddFaceAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *AddFaceAdvanceRequest) SetDbName(v string) *AddFaceAdvanceRequest {
	s.DbName = &v
	return s
}

func (s *AddFaceAdvanceRequest) SetEntityId(v string) *AddFaceAdvanceRequest {
	s.EntityId = &v
	return s
}

func (s *AddFaceAdvanceRequest) SetExtraData(v string) *AddFaceAdvanceRequest {
	s.ExtraData = &v
	return s
}

type AddFaceResponseBody struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AddFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s AddFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceResponseBody) SetRequestId(v string) *AddFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceResponseBody) SetData(v *AddFaceResponseBodyData) *AddFaceResponseBody {
	s.Data = v
	return s
}

type AddFaceResponseBodyData struct {
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
}

func (s AddFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFaceResponseBodyData) SetFaceId(v string) *AddFaceResponseBodyData {
	s.FaceId = &v
	return s
}

type AddFaceResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceResponse) GoString() string {
	return s.String()
}

func (s *AddFaceResponse) SetHeaders(v map[string]*string) *AddFaceResponse {
	s.Headers = v
	return s
}

func (s *AddFaceResponse) SetBody(v *AddFaceResponseBody) *AddFaceResponse {
	s.Body = v
	return s
}

type GenerateHumanSketchStyleRequest struct {
	// A short description of struct
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s GenerateHumanSketchStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanSketchStyleRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleRequest) SetImageURL(v string) *GenerateHumanSketchStyleRequest {
	s.ImageURL = &v
	return s
}

type GenerateHumanSketchStyleAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s GenerateHumanSketchStyleAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanSketchStyleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleAdvanceRequest) SetImageURLObject(v io.Reader) *GenerateHumanSketchStyleAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type GenerateHumanSketchStyleResponseBody struct {
	// Id of the request
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GenerateHumanSketchStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GenerateHumanSketchStyleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanSketchStyleResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleResponseBody) SetRequestId(v string) *GenerateHumanSketchStyleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateHumanSketchStyleResponseBody) SetData(v *GenerateHumanSketchStyleResponseBodyData) *GenerateHumanSketchStyleResponseBody {
	s.Data = v
	return s
}

type GenerateHumanSketchStyleResponseBodyData struct {
	// 出参图片地址
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s GenerateHumanSketchStyleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanSketchStyleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleResponseBodyData) SetImageURL(v string) *GenerateHumanSketchStyleResponseBodyData {
	s.ImageURL = &v
	return s
}

type GenerateHumanSketchStyleResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateHumanSketchStyleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateHumanSketchStyleResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanSketchStyleResponse) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleResponse) SetHeaders(v map[string]*string) *GenerateHumanSketchStyleResponse {
	s.Headers = v
	return s
}

func (s *GenerateHumanSketchStyleResponse) SetBody(v *GenerateHumanSketchStyleResponseBody) *GenerateHumanSketchStyleResponse {
	s.Body = v
	return s
}

type DeleteBodyDbRequest struct {
	// 数据库ID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteBodyDbRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBodyDbRequest) GoString() string {
	return s.String()
}

func (s *DeleteBodyDbRequest) SetId(v int64) *DeleteBodyDbRequest {
	s.Id = &v
	return s
}

type DeleteBodyDbResponseBody struct {
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBodyDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBodyDbResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBodyDbResponseBody) SetRequestId(v string) *DeleteBodyDbResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBodyDbResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBodyDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBodyDbResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBodyDbResponse) GoString() string {
	return s.String()
}

func (s *DeleteBodyDbResponse) SetHeaders(v map[string]*string) *DeleteBodyDbResponse {
	s.Headers = v
	return s
}

func (s *DeleteBodyDbResponse) SetBody(v *DeleteBodyDbResponseBody) *DeleteBodyDbResponse {
	s.Body = v
	return s
}

type DetectPedestrianIntrusionRequest struct {
	ImageURL     *string                                         `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	DetectRegion []*DetectPedestrianIntrusionRequestDetectRegion `json:"DetectRegion,omitempty" xml:"DetectRegion,omitempty" type:"Repeated"`
	RegionType   *string                                         `json:"RegionType,omitempty" xml:"RegionType,omitempty"`
}

func (s DetectPedestrianIntrusionRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionRequest) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionRequest) SetImageURL(v string) *DetectPedestrianIntrusionRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectPedestrianIntrusionRequest) SetDetectRegion(v []*DetectPedestrianIntrusionRequestDetectRegion) *DetectPedestrianIntrusionRequest {
	s.DetectRegion = v
	return s
}

func (s *DetectPedestrianIntrusionRequest) SetRegionType(v string) *DetectPedestrianIntrusionRequest {
	s.RegionType = &v
	return s
}

type DetectPedestrianIntrusionRequestDetectRegion struct {
	Rect *DetectPedestrianIntrusionRequestDetectRegionRect `json:"Rect,omitempty" xml:"Rect,omitempty" type:"Struct"`
	Line *DetectPedestrianIntrusionRequestDetectRegionLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
}

func (s DetectPedestrianIntrusionRequestDetectRegion) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionRequestDetectRegion) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionRequestDetectRegion) SetRect(v *DetectPedestrianIntrusionRequestDetectRegionRect) *DetectPedestrianIntrusionRequestDetectRegion {
	s.Rect = v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegion) SetLine(v *DetectPedestrianIntrusionRequestDetectRegionLine) *DetectPedestrianIntrusionRequestDetectRegion {
	s.Line = v
	return s
}

type DetectPedestrianIntrusionRequestDetectRegionRect struct {
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	Right  *int64 `json:"Right,omitempty" xml:"Right,omitempty"`
	Bottom *int64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
}

func (s DetectPedestrianIntrusionRequestDetectRegionRect) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionRequestDetectRegionRect) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionRequestDetectRegionRect) SetLeft(v int64) *DetectPedestrianIntrusionRequestDetectRegionRect {
	s.Left = &v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegionRect) SetTop(v int64) *DetectPedestrianIntrusionRequestDetectRegionRect {
	s.Top = &v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegionRect) SetRight(v int64) *DetectPedestrianIntrusionRequestDetectRegionRect {
	s.Right = &v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegionRect) SetBottom(v int64) *DetectPedestrianIntrusionRequestDetectRegionRect {
	s.Bottom = &v
	return s
}

type DetectPedestrianIntrusionRequestDetectRegionLine struct {
	X1 *int64 `json:"X1,omitempty" xml:"X1,omitempty"`
	Y1 *int64 `json:"Y1,omitempty" xml:"Y1,omitempty"`
	X2 *int64 `json:"X2,omitempty" xml:"X2,omitempty"`
	Y2 *int64 `json:"Y2,omitempty" xml:"Y2,omitempty"`
}

func (s DetectPedestrianIntrusionRequestDetectRegionLine) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionRequestDetectRegionLine) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionRequestDetectRegionLine) SetX1(v int64) *DetectPedestrianIntrusionRequestDetectRegionLine {
	s.X1 = &v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegionLine) SetY1(v int64) *DetectPedestrianIntrusionRequestDetectRegionLine {
	s.Y1 = &v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegionLine) SetX2(v int64) *DetectPedestrianIntrusionRequestDetectRegionLine {
	s.X2 = &v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegionLine) SetY2(v int64) *DetectPedestrianIntrusionRequestDetectRegionLine {
	s.Y2 = &v
	return s
}

type DetectPedestrianIntrusionAdvanceRequest struct {
	ImageURLObject io.Reader                                              `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	DetectRegion   []*DetectPedestrianIntrusionAdvanceRequestDetectRegion `json:"DetectRegion,omitempty" xml:"DetectRegion,omitempty" type:"Repeated"`
	RegionType     *string                                                `json:"RegionType,omitempty" xml:"RegionType,omitempty"`
}

func (s DetectPedestrianIntrusionAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionAdvanceRequest) SetImageURLObject(v io.Reader) *DetectPedestrianIntrusionAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequest) SetDetectRegion(v []*DetectPedestrianIntrusionAdvanceRequestDetectRegion) *DetectPedestrianIntrusionAdvanceRequest {
	s.DetectRegion = v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequest) SetRegionType(v string) *DetectPedestrianIntrusionAdvanceRequest {
	s.RegionType = &v
	return s
}

type DetectPedestrianIntrusionAdvanceRequestDetectRegion struct {
	Rect *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect `json:"Rect,omitempty" xml:"Rect,omitempty" type:"Struct"`
	Line *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
}

func (s DetectPedestrianIntrusionAdvanceRequestDetectRegion) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionAdvanceRequestDetectRegion) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegion) SetRect(v *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) *DetectPedestrianIntrusionAdvanceRequestDetectRegion {
	s.Rect = v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegion) SetLine(v *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) *DetectPedestrianIntrusionAdvanceRequestDetectRegion {
	s.Line = v
	return s
}

type DetectPedestrianIntrusionAdvanceRequestDetectRegionRect struct {
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	Right  *int64 `json:"Right,omitempty" xml:"Right,omitempty"`
	Bottom *int64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
}

func (s DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) SetLeft(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect {
	s.Left = &v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) SetTop(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect {
	s.Top = &v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) SetRight(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect {
	s.Right = &v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) SetBottom(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect {
	s.Bottom = &v
	return s
}

type DetectPedestrianIntrusionAdvanceRequestDetectRegionLine struct {
	X1 *int64 `json:"X1,omitempty" xml:"X1,omitempty"`
	Y1 *int64 `json:"Y1,omitempty" xml:"Y1,omitempty"`
	X2 *int64 `json:"X2,omitempty" xml:"X2,omitempty"`
	Y2 *int64 `json:"Y2,omitempty" xml:"Y2,omitempty"`
}

func (s DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) SetX1(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine {
	s.X1 = &v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) SetY1(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine {
	s.Y1 = &v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) SetX2(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine {
	s.X2 = &v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) SetY2(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine {
	s.Y2 = &v
	return s
}

type DetectPedestrianIntrusionShrinkRequest struct {
	ImageURL           *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	DetectRegionShrink *string `json:"DetectRegion,omitempty" xml:"DetectRegion,omitempty"`
	RegionType         *string `json:"RegionType,omitempty" xml:"RegionType,omitempty"`
}

func (s DetectPedestrianIntrusionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionShrinkRequest) SetImageURL(v string) *DetectPedestrianIntrusionShrinkRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectPedestrianIntrusionShrinkRequest) SetDetectRegionShrink(v string) *DetectPedestrianIntrusionShrinkRequest {
	s.DetectRegionShrink = &v
	return s
}

func (s *DetectPedestrianIntrusionShrinkRequest) SetRegionType(v string) *DetectPedestrianIntrusionShrinkRequest {
	s.RegionType = &v
	return s
}

type DetectPedestrianIntrusionResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectPedestrianIntrusionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectPedestrianIntrusionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionResponseBody) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionResponseBody) SetRequestId(v string) *DetectPedestrianIntrusionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBody) SetData(v *DetectPedestrianIntrusionResponseBodyData) *DetectPedestrianIntrusionResponseBody {
	s.Data = v
	return s
}

type DetectPedestrianIntrusionResponseBodyData struct {
	ImageWidth  *int64                                               `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	ImageHeight *int64                                               `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	Elements    []*DetectPedestrianIntrusionResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectPedestrianIntrusionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionResponseBodyData) SetImageWidth(v int64) *DetectPedestrianIntrusionResponseBodyData {
	s.ImageWidth = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyData) SetImageHeight(v int64) *DetectPedestrianIntrusionResponseBodyData {
	s.ImageHeight = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyData) SetElements(v []*DetectPedestrianIntrusionResponseBodyDataElements) *DetectPedestrianIntrusionResponseBodyData {
	s.Elements = v
	return s
}

type DetectPedestrianIntrusionResponseBodyDataElements struct {
	Score     *int64                                                `json:"Score,omitempty" xml:"Score,omitempty"`
	Type      *string                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	IsIntrude *bool                                                 `json:"IsIntrude,omitempty" xml:"IsIntrude,omitempty"`
	Box       *DetectPedestrianIntrusionResponseBodyDataElementsBox `json:"Box,omitempty" xml:"Box,omitempty" type:"Struct"`
	BoxId     *int64                                                `json:"BoxId,omitempty" xml:"BoxId,omitempty"`
}

func (s DetectPedestrianIntrusionResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionResponseBodyDataElements) SetScore(v int64) *DetectPedestrianIntrusionResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElements) SetType(v string) *DetectPedestrianIntrusionResponseBodyDataElements {
	s.Type = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElements) SetIsIntrude(v bool) *DetectPedestrianIntrusionResponseBodyDataElements {
	s.IsIntrude = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElements) SetBox(v *DetectPedestrianIntrusionResponseBodyDataElementsBox) *DetectPedestrianIntrusionResponseBodyDataElements {
	s.Box = v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElements) SetBoxId(v int64) *DetectPedestrianIntrusionResponseBodyDataElements {
	s.BoxId = &v
	return s
}

type DetectPedestrianIntrusionResponseBodyDataElementsBox struct {
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	Right  *int64 `json:"Right,omitempty" xml:"Right,omitempty"`
	Bottom *int64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
}

func (s DetectPedestrianIntrusionResponseBodyDataElementsBox) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionResponseBodyDataElementsBox) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionResponseBodyDataElementsBox) SetLeft(v int64) *DetectPedestrianIntrusionResponseBodyDataElementsBox {
	s.Left = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElementsBox) SetTop(v int64) *DetectPedestrianIntrusionResponseBodyDataElementsBox {
	s.Top = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElementsBox) SetRight(v int64) *DetectPedestrianIntrusionResponseBodyDataElementsBox {
	s.Right = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElementsBox) SetBottom(v int64) *DetectPedestrianIntrusionResponseBodyDataElementsBox {
	s.Bottom = &v
	return s
}

type DetectPedestrianIntrusionResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectPedestrianIntrusionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectPedestrianIntrusionResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionResponse) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionResponse) SetHeaders(v map[string]*string) *DetectPedestrianIntrusionResponse {
	s.Headers = v
	return s
}

func (s *DetectPedestrianIntrusionResponse) SetBody(v *DetectPedestrianIntrusionResponseBody) *DetectPedestrianIntrusionResponse {
	s.Body = v
	return s
}

type HandPostureRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s HandPostureRequest) String() string {
	return tea.Prettify(s)
}

func (s HandPostureRequest) GoString() string {
	return s.String()
}

func (s *HandPostureRequest) SetImageURL(v string) *HandPostureRequest {
	s.ImageURL = &v
	return s
}

type HandPostureAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s HandPostureAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s HandPostureAdvanceRequest) GoString() string {
	return s.String()
}

func (s *HandPostureAdvanceRequest) SetImageURLObject(v io.Reader) *HandPostureAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type HandPostureResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *HandPostureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s HandPostureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBody) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBody) SetRequestId(v string) *HandPostureResponseBody {
	s.RequestId = &v
	return s
}

func (s *HandPostureResponseBody) SetData(v *HandPostureResponseBodyData) *HandPostureResponseBody {
	s.Data = v
	return s
}

type HandPostureResponseBodyData struct {
	Outputs    []*HandPostureResponseBodyDataOutputs  `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	MetaObject *HandPostureResponseBodyDataMetaObject `json:"MetaObject,omitempty" xml:"MetaObject,omitempty" type:"Struct"`
}

func (s HandPostureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyData) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyData) SetOutputs(v []*HandPostureResponseBodyDataOutputs) *HandPostureResponseBodyData {
	s.Outputs = v
	return s
}

func (s *HandPostureResponseBodyData) SetMetaObject(v *HandPostureResponseBodyDataMetaObject) *HandPostureResponseBodyData {
	s.MetaObject = v
	return s
}

type HandPostureResponseBodyDataOutputs struct {
	HandCount *int32                                       `json:"HandCount,omitempty" xml:"HandCount,omitempty"`
	Results   []*HandPostureResponseBodyDataOutputsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s HandPostureResponseBodyDataOutputs) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputs) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputs) SetHandCount(v int32) *HandPostureResponseBodyDataOutputs {
	s.HandCount = &v
	return s
}

func (s *HandPostureResponseBodyDataOutputs) SetResults(v []*HandPostureResponseBodyDataOutputsResults) *HandPostureResponseBodyDataOutputs {
	s.Results = v
	return s
}

type HandPostureResponseBodyDataOutputsResults struct {
	Hands *HandPostureResponseBodyDataOutputsResultsHands `json:"Hands,omitempty" xml:"Hands,omitempty" type:"Struct"`
	Box   *HandPostureResponseBodyDataOutputsResultsBox   `json:"Box,omitempty" xml:"Box,omitempty" type:"Struct"`
}

func (s HandPostureResponseBodyDataOutputsResults) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputsResults) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputsResults) SetHands(v *HandPostureResponseBodyDataOutputsResultsHands) *HandPostureResponseBodyDataOutputsResults {
	s.Hands = v
	return s
}

func (s *HandPostureResponseBodyDataOutputsResults) SetBox(v *HandPostureResponseBodyDataOutputsResultsBox) *HandPostureResponseBodyDataOutputsResults {
	s.Box = v
	return s
}

type HandPostureResponseBodyDataOutputsResultsHands struct {
	KeyPoints []*HandPostureResponseBodyDataOutputsResultsHandsKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Confident *float32                                                   `json:"Confident,omitempty" xml:"Confident,omitempty"`
}

func (s HandPostureResponseBodyDataOutputsResultsHands) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputsResultsHands) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputsResultsHands) SetKeyPoints(v []*HandPostureResponseBodyDataOutputsResultsHandsKeyPoints) *HandPostureResponseBodyDataOutputsResultsHands {
	s.KeyPoints = v
	return s
}

func (s *HandPostureResponseBodyDataOutputsResultsHands) SetConfident(v float32) *HandPostureResponseBodyDataOutputsResultsHands {
	s.Confident = &v
	return s
}

type HandPostureResponseBodyDataOutputsResultsHandsKeyPoints struct {
	Positions []*HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions `json:"Positions,omitempty" xml:"Positions,omitempty" type:"Repeated"`
	Label     *string                                                             `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s HandPostureResponseBodyDataOutputsResultsHandsKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputsResultsHandsKeyPoints) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputsResultsHandsKeyPoints) SetPositions(v []*HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions) *HandPostureResponseBodyDataOutputsResultsHandsKeyPoints {
	s.Positions = v
	return s
}

func (s *HandPostureResponseBodyDataOutputsResultsHandsKeyPoints) SetLabel(v string) *HandPostureResponseBodyDataOutputsResultsHandsKeyPoints {
	s.Label = &v
	return s
}

type HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions struct {
	Points []*float32 `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions) SetPoints(v []*float32) *HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions {
	s.Points = v
	return s
}

type HandPostureResponseBodyDataOutputsResultsBox struct {
	Positions []*HandPostureResponseBodyDataOutputsResultsBoxPositions `json:"Positions,omitempty" xml:"Positions,omitempty" type:"Repeated"`
	Confident *float32                                                 `json:"Confident,omitempty" xml:"Confident,omitempty"`
}

func (s HandPostureResponseBodyDataOutputsResultsBox) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputsResultsBox) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputsResultsBox) SetPositions(v []*HandPostureResponseBodyDataOutputsResultsBoxPositions) *HandPostureResponseBodyDataOutputsResultsBox {
	s.Positions = v
	return s
}

func (s *HandPostureResponseBodyDataOutputsResultsBox) SetConfident(v float32) *HandPostureResponseBodyDataOutputsResultsBox {
	s.Confident = &v
	return s
}

type HandPostureResponseBodyDataOutputsResultsBoxPositions struct {
	Points []*float32 `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s HandPostureResponseBodyDataOutputsResultsBoxPositions) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputsResultsBoxPositions) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputsResultsBoxPositions) SetPoints(v []*float32) *HandPostureResponseBodyDataOutputsResultsBoxPositions {
	s.Points = v
	return s
}

type HandPostureResponseBodyDataMetaObject struct {
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s HandPostureResponseBodyDataMetaObject) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataMetaObject) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataMetaObject) SetWidth(v int32) *HandPostureResponseBodyDataMetaObject {
	s.Width = &v
	return s
}

func (s *HandPostureResponseBodyDataMetaObject) SetHeight(v int32) *HandPostureResponseBodyDataMetaObject {
	s.Height = &v
	return s
}

type HandPostureResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HandPostureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HandPostureResponse) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponse) GoString() string {
	return s.String()
}

func (s *HandPostureResponse) SetHeaders(v map[string]*string) *HandPostureResponse {
	s.Headers = v
	return s
}

func (s *HandPostureResponse) SetBody(v *HandPostureResponseBody) *HandPostureResponse {
	s.Body = v
	return s
}

type EnhanceFaceRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s EnhanceFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnhanceFaceRequest) GoString() string {
	return s.String()
}

func (s *EnhanceFaceRequest) SetImageURL(v string) *EnhanceFaceRequest {
	s.ImageURL = &v
	return s
}

type EnhanceFaceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s EnhanceFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnhanceFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EnhanceFaceAdvanceRequest) SetImageURLObject(v io.Reader) *EnhanceFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type EnhanceFaceResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *EnhanceFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s EnhanceFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnhanceFaceResponseBody) GoString() string {
	return s.String()
}

func (s *EnhanceFaceResponseBody) SetRequestId(v string) *EnhanceFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnhanceFaceResponseBody) SetData(v *EnhanceFaceResponseBodyData) *EnhanceFaceResponseBody {
	s.Data = v
	return s
}

type EnhanceFaceResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s EnhanceFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnhanceFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnhanceFaceResponseBodyData) SetImageURL(v string) *EnhanceFaceResponseBodyData {
	s.ImageURL = &v
	return s
}

type EnhanceFaceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnhanceFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnhanceFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s EnhanceFaceResponse) GoString() string {
	return s.String()
}

func (s *EnhanceFaceResponse) SetHeaders(v map[string]*string) *EnhanceFaceResponse {
	s.Headers = v
	return s
}

func (s *EnhanceFaceResponse) SetBody(v *EnhanceFaceResponseBody) *EnhanceFaceResponse {
	s.Body = v
	return s
}

type GetBodyPersonRequest struct {
	// 数据库ID
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// 人员ID
	PersonId *int64 `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s GetBodyPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBodyPersonRequest) GoString() string {
	return s.String()
}

func (s *GetBodyPersonRequest) SetDbId(v int64) *GetBodyPersonRequest {
	s.DbId = &v
	return s
}

func (s *GetBodyPersonRequest) SetPersonId(v int64) *GetBodyPersonRequest {
	s.PersonId = &v
	return s
}

type GetBodyPersonResponseBody struct {
	// RequestId
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetBodyPersonResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetBodyPersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBodyPersonResponseBody) GoString() string {
	return s.String()
}

func (s *GetBodyPersonResponseBody) SetRequestId(v string) *GetBodyPersonResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBodyPersonResponseBody) SetData(v *GetBodyPersonResponseBodyData) *GetBodyPersonResponseBody {
	s.Data = v
	return s
}

type GetBodyPersonResponseBodyData struct {
	// 数据库ID
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// 人员ID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 人员名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Trace数量
	TraceCount *int64 `json:"TraceCount,omitempty" xml:"TraceCount,omitempty"`
	// Trace列表
	TraceList []*GetBodyPersonResponseBodyDataTraceList `json:"TraceList,omitempty" xml:"TraceList,omitempty" type:"Repeated"`
}

func (s GetBodyPersonResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetBodyPersonResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBodyPersonResponseBodyData) SetDbId(v int64) *GetBodyPersonResponseBodyData {
	s.DbId = &v
	return s
}

func (s *GetBodyPersonResponseBodyData) SetId(v int64) *GetBodyPersonResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetBodyPersonResponseBodyData) SetName(v string) *GetBodyPersonResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetBodyPersonResponseBodyData) SetTraceCount(v int64) *GetBodyPersonResponseBodyData {
	s.TraceCount = &v
	return s
}

func (s *GetBodyPersonResponseBodyData) SetTraceList(v []*GetBodyPersonResponseBodyDataTraceList) *GetBodyPersonResponseBodyData {
	s.TraceList = v
	return s
}

type GetBodyPersonResponseBodyDataTraceList struct {
	// TraceId
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 自定义Trace信息
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
}

func (s GetBodyPersonResponseBodyDataTraceList) String() string {
	return tea.Prettify(s)
}

func (s GetBodyPersonResponseBodyDataTraceList) GoString() string {
	return s.String()
}

func (s *GetBodyPersonResponseBodyDataTraceList) SetId(v int64) *GetBodyPersonResponseBodyDataTraceList {
	s.Id = &v
	return s
}

func (s *GetBodyPersonResponseBodyDataTraceList) SetExtraData(v string) *GetBodyPersonResponseBodyDataTraceList {
	s.ExtraData = &v
	return s
}

type GetBodyPersonResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBodyPersonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBodyPersonResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBodyPersonResponse) GoString() string {
	return s.String()
}

func (s *GetBodyPersonResponse) SetHeaders(v map[string]*string) *GetBodyPersonResponse {
	s.Headers = v
	return s
}

func (s *GetBodyPersonResponse) SetBody(v *GetBodyPersonResponseBody) *GetBodyPersonResponse {
	s.Body = v
	return s
}

type DeleteFaceDbRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteFaceDbRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceDbRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceDbRequest) SetName(v string) *DeleteFaceDbRequest {
	s.Name = &v
	return s
}

type DeleteFaceDbResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceDbResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceDbResponseBody) SetRequestId(v string) *DeleteFaceDbResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFaceDbResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFaceDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFaceDbResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceDbResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceDbResponse) SetHeaders(v map[string]*string) *DeleteFaceDbResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceDbResponse) SetBody(v *DeleteFaceDbResponseBody) *DeleteFaceDbResponse {
	s.Body = v
	return s
}

type ListBodyPersonRequest struct {
	// 数据库ID
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// 起始位置(不含)
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// 分页数量
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s ListBodyPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBodyPersonRequest) GoString() string {
	return s.String()
}

func (s *ListBodyPersonRequest) SetDbId(v int64) *ListBodyPersonRequest {
	s.DbId = &v
	return s
}

func (s *ListBodyPersonRequest) SetOffset(v int64) *ListBodyPersonRequest {
	s.Offset = &v
	return s
}

func (s *ListBodyPersonRequest) SetLimit(v int64) *ListBodyPersonRequest {
	s.Limit = &v
	return s
}

type ListBodyPersonResponseBody struct {
	// RequestId
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListBodyPersonResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListBodyPersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBodyPersonResponseBody) GoString() string {
	return s.String()
}

func (s *ListBodyPersonResponseBody) SetRequestId(v string) *ListBodyPersonResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBodyPersonResponseBody) SetData(v *ListBodyPersonResponseBodyData) *ListBodyPersonResponseBody {
	s.Data = v
	return s
}

type ListBodyPersonResponseBodyData struct {
	// 数据总量
	Total      *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
	PersonList []*ListBodyPersonResponseBodyDataPersonList `json:"PersonList,omitempty" xml:"PersonList,omitempty" type:"Repeated"`
}

func (s ListBodyPersonResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListBodyPersonResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBodyPersonResponseBodyData) SetTotal(v int64) *ListBodyPersonResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListBodyPersonResponseBodyData) SetPersonList(v []*ListBodyPersonResponseBodyDataPersonList) *ListBodyPersonResponseBodyData {
	s.PersonList = v
	return s
}

type ListBodyPersonResponseBodyDataPersonList struct {
	// 数据库ID
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// 人员名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Trace数量
	TraceCount *int64 `json:"TraceCount,omitempty" xml:"TraceCount,omitempty"`
	// 人员ID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListBodyPersonResponseBodyDataPersonList) String() string {
	return tea.Prettify(s)
}

func (s ListBodyPersonResponseBodyDataPersonList) GoString() string {
	return s.String()
}

func (s *ListBodyPersonResponseBodyDataPersonList) SetDbId(v int64) *ListBodyPersonResponseBodyDataPersonList {
	s.DbId = &v
	return s
}

func (s *ListBodyPersonResponseBodyDataPersonList) SetName(v string) *ListBodyPersonResponseBodyDataPersonList {
	s.Name = &v
	return s
}

func (s *ListBodyPersonResponseBodyDataPersonList) SetTraceCount(v int64) *ListBodyPersonResponseBodyDataPersonList {
	s.TraceCount = &v
	return s
}

func (s *ListBodyPersonResponseBodyDataPersonList) SetId(v int64) *ListBodyPersonResponseBodyDataPersonList {
	s.Id = &v
	return s
}

type ListBodyPersonResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBodyPersonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBodyPersonResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBodyPersonResponse) GoString() string {
	return s.String()
}

func (s *ListBodyPersonResponse) SetHeaders(v map[string]*string) *ListBodyPersonResponse {
	s.Headers = v
	return s
}

func (s *ListBodyPersonResponse) SetBody(v *ListBodyPersonResponseBody) *ListBodyPersonResponse {
	s.Body = v
	return s
}

type ListBodyDbsRequest struct {
	// 起始位置(不含)
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// 分页数量
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s ListBodyDbsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBodyDbsRequest) GoString() string {
	return s.String()
}

func (s *ListBodyDbsRequest) SetOffset(v int64) *ListBodyDbsRequest {
	s.Offset = &v
	return s
}

func (s *ListBodyDbsRequest) SetLimit(v int64) *ListBodyDbsRequest {
	s.Limit = &v
	return s
}

type ListBodyDbsResponseBody struct {
	// RequestId
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListBodyDbsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListBodyDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBodyDbsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBodyDbsResponseBody) SetRequestId(v string) *ListBodyDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBodyDbsResponseBody) SetData(v *ListBodyDbsResponseBodyData) *ListBodyDbsResponseBody {
	s.Data = v
	return s
}

type ListBodyDbsResponseBodyData struct {
	// 数据库ID
	Total  *int64                               `json:"Total,omitempty" xml:"Total,omitempty"`
	DbList []*ListBodyDbsResponseBodyDataDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
}

func (s ListBodyDbsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListBodyDbsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBodyDbsResponseBodyData) SetTotal(v int64) *ListBodyDbsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListBodyDbsResponseBodyData) SetDbList(v []*ListBodyDbsResponseBodyDataDbList) *ListBodyDbsResponseBodyData {
	s.DbList = v
	return s
}

type ListBodyDbsResponseBodyDataDbList struct {
	// 数据库ID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 数据库名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListBodyDbsResponseBodyDataDbList) String() string {
	return tea.Prettify(s)
}

func (s ListBodyDbsResponseBodyDataDbList) GoString() string {
	return s.String()
}

func (s *ListBodyDbsResponseBodyDataDbList) SetId(v int64) *ListBodyDbsResponseBodyDataDbList {
	s.Id = &v
	return s
}

func (s *ListBodyDbsResponseBodyDataDbList) SetName(v string) *ListBodyDbsResponseBodyDataDbList {
	s.Name = &v
	return s
}

type ListBodyDbsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBodyDbsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBodyDbsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBodyDbsResponse) GoString() string {
	return s.String()
}

func (s *ListBodyDbsResponse) SetHeaders(v map[string]*string) *ListBodyDbsResponse {
	s.Headers = v
	return s
}

func (s *ListBodyDbsResponse) SetBody(v *ListBodyDbsResponseBody) *ListBodyDbsResponse {
	s.Body = v
	return s
}

type ListFaceEntitiesRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Offset         *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Limit          *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Token          *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Labels         *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	EntityIdPrefix *string `json:"EntityIdPrefix,omitempty" xml:"EntityIdPrefix,omitempty"`
	Order          *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s ListFaceEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFaceEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesRequest) SetDbName(v string) *ListFaceEntitiesRequest {
	s.DbName = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetOffset(v int32) *ListFaceEntitiesRequest {
	s.Offset = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetLimit(v int32) *ListFaceEntitiesRequest {
	s.Limit = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetToken(v string) *ListFaceEntitiesRequest {
	s.Token = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetLabels(v string) *ListFaceEntitiesRequest {
	s.Labels = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetEntityIdPrefix(v string) *ListFaceEntitiesRequest {
	s.EntityIdPrefix = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetOrder(v string) *ListFaceEntitiesRequest {
	s.Order = &v
	return s
}

type ListFaceEntitiesResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListFaceEntitiesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListFaceEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFaceEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesResponseBody) SetRequestId(v string) *ListFaceEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFaceEntitiesResponseBody) SetData(v *ListFaceEntitiesResponseBodyData) *ListFaceEntitiesResponseBody {
	s.Data = v
	return s
}

type ListFaceEntitiesResponseBodyData struct {
	Token      *string                                     `json:"Token,omitempty" xml:"Token,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Entities   []*ListFaceEntitiesResponseBodyDataEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
}

func (s ListFaceEntitiesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFaceEntitiesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesResponseBodyData) SetToken(v string) *ListFaceEntitiesResponseBodyData {
	s.Token = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyData) SetTotalCount(v int32) *ListFaceEntitiesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyData) SetEntities(v []*ListFaceEntitiesResponseBodyDataEntities) *ListFaceEntitiesResponseBodyData {
	s.Entities = v
	return s
}

type ListFaceEntitiesResponseBodyDataEntities struct {
	DbName    *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Labels    *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	CreatedAt *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	FaceCount *int32  `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListFaceEntitiesResponseBodyDataEntities) String() string {
	return tea.Prettify(s)
}

func (s ListFaceEntitiesResponseBodyDataEntities) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetDbName(v string) *ListFaceEntitiesResponseBodyDataEntities {
	s.DbName = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetEntityId(v string) *ListFaceEntitiesResponseBodyDataEntities {
	s.EntityId = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetLabels(v string) *ListFaceEntitiesResponseBodyDataEntities {
	s.Labels = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetCreatedAt(v int64) *ListFaceEntitiesResponseBodyDataEntities {
	s.CreatedAt = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetFaceCount(v int32) *ListFaceEntitiesResponseBodyDataEntities {
	s.FaceCount = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetUpdatedAt(v int64) *ListFaceEntitiesResponseBodyDataEntities {
	s.UpdatedAt = &v
	return s
}

type ListFaceEntitiesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFaceEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFaceEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFaceEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesResponse) SetHeaders(v map[string]*string) *ListFaceEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListFaceEntitiesResponse) SetBody(v *ListFaceEntitiesResponseBody) *ListFaceEntitiesResponse {
	s.Body = v
	return s
}

type RecognizePublicFaceRequest struct {
	Task []*RecognizePublicFaceRequestTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s RecognizePublicFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceRequest) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceRequest) SetTask(v []*RecognizePublicFaceRequestTask) *RecognizePublicFaceRequest {
	s.Task = v
	return s
}

type RecognizePublicFaceRequestTask struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizePublicFaceRequestTask) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceRequestTask) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceRequestTask) SetImageURL(v string) *RecognizePublicFaceRequestTask {
	s.ImageURL = &v
	return s
}

type RecognizePublicFaceResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizePublicFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizePublicFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBody) SetRequestId(v string) *RecognizePublicFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizePublicFaceResponseBody) SetData(v *RecognizePublicFaceResponseBodyData) *RecognizePublicFaceResponseBody {
	s.Data = v
	return s
}

type RecognizePublicFaceResponseBodyData struct {
	Elements []*RecognizePublicFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizePublicFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyData) SetElements(v []*RecognizePublicFaceResponseBodyDataElements) *RecognizePublicFaceResponseBodyData {
	s.Elements = v
	return s
}

type RecognizePublicFaceResponseBodyDataElements struct {
	ImageURL *string                                               `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	TaskId   *string                                               `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Results  []*RecognizePublicFaceResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RecognizePublicFaceResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyDataElements) SetImageURL(v string) *RecognizePublicFaceResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElements) SetTaskId(v string) *RecognizePublicFaceResponseBodyDataElements {
	s.TaskId = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElements) SetResults(v []*RecognizePublicFaceResponseBodyDataElementsResults) *RecognizePublicFaceResponseBodyDataElements {
	s.Results = v
	return s
}

type RecognizePublicFaceResponseBodyDataElementsResults struct {
	Suggestion *string                                                         `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	Label      *string                                                         `json:"Label,omitempty" xml:"Label,omitempty"`
	SubResults []*RecognizePublicFaceResponseBodyDataElementsResultsSubResults `json:"SubResults,omitempty" xml:"SubResults,omitempty" type:"Repeated"`
	Rate       *float32                                                        `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s RecognizePublicFaceResponseBodyDataElementsResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) SetSuggestion(v string) *RecognizePublicFaceResponseBodyDataElementsResults {
	s.Suggestion = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) SetLabel(v string) *RecognizePublicFaceResponseBodyDataElementsResults {
	s.Label = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) SetSubResults(v []*RecognizePublicFaceResponseBodyDataElementsResultsSubResults) *RecognizePublicFaceResponseBodyDataElementsResults {
	s.SubResults = v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) SetRate(v float32) *RecognizePublicFaceResponseBodyDataElementsResults {
	s.Rate = &v
	return s
}

type RecognizePublicFaceResponseBodyDataElementsResultsSubResults struct {
	W     *float32                                                             `json:"W,omitempty" xml:"W,omitempty"`
	Faces []*RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	H     *float32                                                             `json:"H,omitempty" xml:"H,omitempty"`
	Y     *float32                                                             `json:"Y,omitempty" xml:"Y,omitempty"`
	X     *float32                                                             `json:"X,omitempty" xml:"X,omitempty"`
}

func (s RecognizePublicFaceResponseBodyDataElementsResultsSubResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyDataElementsResultsSubResults) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetW(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.W = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetFaces(v []*RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.Faces = v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetH(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.H = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetY(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.Y = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetX(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.X = &v
	return s
}

type RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces struct {
	Name *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *string  `json:"Id,omitempty" xml:"Id,omitempty"`
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) SetName(v string) *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces {
	s.Name = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) SetId(v string) *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces {
	s.Id = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) SetRate(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces {
	s.Rate = &v
	return s
}

type RecognizePublicFaceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizePublicFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizePublicFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponse) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponse) SetHeaders(v map[string]*string) *RecognizePublicFaceResponse {
	s.Headers = v
	return s
}

func (s *RecognizePublicFaceResponse) SetBody(v *RecognizePublicFaceResponseBody) *RecognizePublicFaceResponse {
	s.Body = v
	return s
}

type DeleteFaceImageTemplateRequest struct {
	// A short description of struct
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteFaceImageTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceImageTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceImageTemplateRequest) SetUserId(v string) *DeleteFaceImageTemplateRequest {
	s.UserId = &v
	return s
}

func (s *DeleteFaceImageTemplateRequest) SetTemplateId(v string) *DeleteFaceImageTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteFaceImageTemplateResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceImageTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceImageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceImageTemplateResponseBody) SetRequestId(v string) *DeleteFaceImageTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFaceImageTemplateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFaceImageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFaceImageTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceImageTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceImageTemplateResponse) SetHeaders(v map[string]*string) *DeleteFaceImageTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceImageTemplateResponse) SetBody(v *DeleteFaceImageTemplateResponseBody) *DeleteFaceImageTemplateResponse {
	s.Body = v
	return s
}

type CreateFaceDbRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateFaceDbRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceDbRequest) GoString() string {
	return s.String()
}

func (s *CreateFaceDbRequest) SetName(v string) *CreateFaceDbRequest {
	s.Name = &v
	return s
}

type CreateFaceDbResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFaceDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceDbResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFaceDbResponseBody) SetRequestId(v string) *CreateFaceDbResponseBody {
	s.RequestId = &v
	return s
}

type CreateFaceDbResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFaceDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFaceDbResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceDbResponse) GoString() string {
	return s.String()
}

func (s *CreateFaceDbResponse) SetHeaders(v map[string]*string) *CreateFaceDbResponse {
	s.Headers = v
	return s
}

func (s *CreateFaceDbResponse) SetBody(v *CreateFaceDbResponseBody) *CreateFaceDbResponse {
	s.Body = v
	return s
}

type AddBodyTraceRequest struct {
	// 数据库ID
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// 实体ID，可以包含数字、字母和下划线，相同db下不可以重复，长度1-64。
	PersonId *int64 `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	// Trace图片信息列表
	Images []*AddBodyTraceRequestImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// 自定义信息。支持字母、数字、标点符号和汉字。不超过4096个字符。
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
}

func (s AddBodyTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddBodyTraceRequest) GoString() string {
	return s.String()
}

func (s *AddBodyTraceRequest) SetDbId(v int64) *AddBodyTraceRequest {
	s.DbId = &v
	return s
}

func (s *AddBodyTraceRequest) SetPersonId(v int64) *AddBodyTraceRequest {
	s.PersonId = &v
	return s
}

func (s *AddBodyTraceRequest) SetImages(v []*AddBodyTraceRequestImages) *AddBodyTraceRequest {
	s.Images = v
	return s
}

func (s *AddBodyTraceRequest) SetExtraData(v string) *AddBodyTraceRequest {
	s.ExtraData = &v
	return s
}

type AddBodyTraceRequestImages struct {
	// Trace图片URL
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AddBodyTraceRequestImages) String() string {
	return tea.Prettify(s)
}

func (s AddBodyTraceRequestImages) GoString() string {
	return s.String()
}

func (s *AddBodyTraceRequestImages) SetImageURL(v string) *AddBodyTraceRequestImages {
	s.ImageURL = &v
	return s
}

type AddBodyTraceShrinkRequest struct {
	// 数据库ID
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// 实体ID，可以包含数字、字母和下划线，相同db下不可以重复，长度1-64。
	PersonId *int64 `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	// Trace图片信息列表
	ImagesShrink *string `json:"Images,omitempty" xml:"Images,omitempty"`
	// 自定义信息。支持字母、数字、标点符号和汉字。不超过4096个字符。
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
}

func (s AddBodyTraceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddBodyTraceShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddBodyTraceShrinkRequest) SetDbId(v int64) *AddBodyTraceShrinkRequest {
	s.DbId = &v
	return s
}

func (s *AddBodyTraceShrinkRequest) SetPersonId(v int64) *AddBodyTraceShrinkRequest {
	s.PersonId = &v
	return s
}

func (s *AddBodyTraceShrinkRequest) SetImagesShrink(v string) *AddBodyTraceShrinkRequest {
	s.ImagesShrink = &v
	return s
}

func (s *AddBodyTraceShrinkRequest) SetExtraData(v string) *AddBodyTraceShrinkRequest {
	s.ExtraData = &v
	return s
}

type AddBodyTraceResponseBody struct {
	// RequestId
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AddBodyTraceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s AddBodyTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddBodyTraceResponseBody) GoString() string {
	return s.String()
}

func (s *AddBodyTraceResponseBody) SetRequestId(v string) *AddBodyTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBodyTraceResponseBody) SetData(v *AddBodyTraceResponseBodyData) *AddBodyTraceResponseBody {
	s.Data = v
	return s
}

type AddBodyTraceResponseBodyData struct {
	// TraceID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AddBodyTraceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddBodyTraceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddBodyTraceResponseBodyData) SetId(v int64) *AddBodyTraceResponseBodyData {
	s.Id = &v
	return s
}

type AddBodyTraceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddBodyTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddBodyTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddBodyTraceResponse) GoString() string {
	return s.String()
}

func (s *AddBodyTraceResponse) SetHeaders(v map[string]*string) *AddBodyTraceResponse {
	s.Headers = v
	return s
}

func (s *AddBodyTraceResponse) SetBody(v *AddBodyTraceResponseBody) *AddBodyTraceResponse {
	s.Body = v
	return s
}

type SearchBodyTraceRequest struct {
	// 数据库ID
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Trace图片信息列表
	Images []*SearchBodyTraceRequestImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// 结果Trace数量上限，默认10，取值范围[1, 100]
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// score阈值，只返回大于等于该score的数据，取值范围[-1.0, 1.0]，默认为空
	MinScore *float32 `json:"MinScore,omitempty" xml:"MinScore,omitempty"`
}

func (s SearchBodyTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyTraceRequest) GoString() string {
	return s.String()
}

func (s *SearchBodyTraceRequest) SetDbId(v int64) *SearchBodyTraceRequest {
	s.DbId = &v
	return s
}

func (s *SearchBodyTraceRequest) SetImages(v []*SearchBodyTraceRequestImages) *SearchBodyTraceRequest {
	s.Images = v
	return s
}

func (s *SearchBodyTraceRequest) SetLimit(v int64) *SearchBodyTraceRequest {
	s.Limit = &v
	return s
}

func (s *SearchBodyTraceRequest) SetMinScore(v float32) *SearchBodyTraceRequest {
	s.MinScore = &v
	return s
}

type SearchBodyTraceRequestImages struct {
	// Trace图片URL
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SearchBodyTraceRequestImages) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyTraceRequestImages) GoString() string {
	return s.String()
}

func (s *SearchBodyTraceRequestImages) SetImageURL(v string) *SearchBodyTraceRequestImages {
	s.ImageURL = &v
	return s
}

type SearchBodyTraceShrinkRequest struct {
	// 数据库ID
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Trace图片信息列表
	ImagesShrink *string `json:"Images,omitempty" xml:"Images,omitempty"`
	// 结果Trace数量上限，默认10，取值范围[1, 100]
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// score阈值，只返回大于等于该score的数据，取值范围[-1.0, 1.0]，默认为空
	MinScore *float32 `json:"MinScore,omitempty" xml:"MinScore,omitempty"`
}

func (s SearchBodyTraceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyTraceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchBodyTraceShrinkRequest) SetDbId(v int64) *SearchBodyTraceShrinkRequest {
	s.DbId = &v
	return s
}

func (s *SearchBodyTraceShrinkRequest) SetImagesShrink(v string) *SearchBodyTraceShrinkRequest {
	s.ImagesShrink = &v
	return s
}

func (s *SearchBodyTraceShrinkRequest) SetLimit(v int64) *SearchBodyTraceShrinkRequest {
	s.Limit = &v
	return s
}

func (s *SearchBodyTraceShrinkRequest) SetMinScore(v float32) *SearchBodyTraceShrinkRequest {
	s.MinScore = &v
	return s
}

type SearchBodyTraceResponseBody struct {
	// RequestId
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SearchBodyTraceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s SearchBodyTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyTraceResponseBody) GoString() string {
	return s.String()
}

func (s *SearchBodyTraceResponseBody) SetRequestId(v string) *SearchBodyTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchBodyTraceResponseBody) SetData(v *SearchBodyTraceResponseBodyData) *SearchBodyTraceResponseBody {
	s.Data = v
	return s
}

type SearchBodyTraceResponseBodyData struct {
	// 匹配的Trace列表
	MatchList []*SearchBodyTraceResponseBodyDataMatchList `json:"MatchList,omitempty" xml:"MatchList,omitempty" type:"Repeated"`
}

func (s SearchBodyTraceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyTraceResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchBodyTraceResponseBodyData) SetMatchList(v []*SearchBodyTraceResponseBodyDataMatchList) *SearchBodyTraceResponseBodyData {
	s.MatchList = v
	return s
}

type SearchBodyTraceResponseBodyDataMatchList struct {
	// 数据库ID
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// 人员ID
	PersonId *int64 `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	// TraceId
	TraceId *int64 `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// 匹配度分值，越大越相似，取值范围[-1.0, 1.0]
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// 自定义数据信息
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
}

func (s SearchBodyTraceResponseBodyDataMatchList) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyTraceResponseBodyDataMatchList) GoString() string {
	return s.String()
}

func (s *SearchBodyTraceResponseBodyDataMatchList) SetDbId(v int64) *SearchBodyTraceResponseBodyDataMatchList {
	s.DbId = &v
	return s
}

func (s *SearchBodyTraceResponseBodyDataMatchList) SetPersonId(v int64) *SearchBodyTraceResponseBodyDataMatchList {
	s.PersonId = &v
	return s
}

func (s *SearchBodyTraceResponseBodyDataMatchList) SetTraceId(v int64) *SearchBodyTraceResponseBodyDataMatchList {
	s.TraceId = &v
	return s
}

func (s *SearchBodyTraceResponseBodyDataMatchList) SetScore(v float32) *SearchBodyTraceResponseBodyDataMatchList {
	s.Score = &v
	return s
}

func (s *SearchBodyTraceResponseBodyDataMatchList) SetExtraData(v string) *SearchBodyTraceResponseBodyDataMatchList {
	s.ExtraData = &v
	return s
}

type SearchBodyTraceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchBodyTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchBodyTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyTraceResponse) GoString() string {
	return s.String()
}

func (s *SearchBodyTraceResponse) SetHeaders(v map[string]*string) *SearchBodyTraceResponse {
	s.Headers = v
	return s
}

func (s *SearchBodyTraceResponse) SetBody(v *SearchBodyTraceResponseBody) *SearchBodyTraceResponse {
	s.Body = v
	return s
}

type AddFaceImageTemplateRequest struct {
	// A short description of struct
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AddFaceImageTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceImageTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateRequest) SetUserId(v string) *AddFaceImageTemplateRequest {
	s.UserId = &v
	return s
}

func (s *AddFaceImageTemplateRequest) SetImageURL(v string) *AddFaceImageTemplateRequest {
	s.ImageURL = &v
	return s
}

type AddFaceImageTemplateAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	// A short description of struct
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddFaceImageTemplateAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceImageTemplateAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateAdvanceRequest) SetImageURLObject(v io.Reader) *AddFaceImageTemplateAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *AddFaceImageTemplateAdvanceRequest) SetUserId(v string) *AddFaceImageTemplateAdvanceRequest {
	s.UserId = &v
	return s
}

type AddFaceImageTemplateResponseBody struct {
	// Id of the request
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AddFaceImageTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s AddFaceImageTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceImageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateResponseBody) SetRequestId(v string) *AddFaceImageTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceImageTemplateResponseBody) SetData(v *AddFaceImageTemplateResponseBodyData) *AddFaceImageTemplateResponseBody {
	s.Data = v
	return s
}

type AddFaceImageTemplateResponseBodyData struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AddFaceImageTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddFaceImageTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateResponseBodyData) SetTemplateId(v string) *AddFaceImageTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

type AddFaceImageTemplateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFaceImageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddFaceImageTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceImageTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateResponse) SetHeaders(v map[string]*string) *AddFaceImageTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddFaceImageTemplateResponse) SetBody(v *AddFaceImageTemplateResponseBody) *AddFaceImageTemplateResponse {
	s.Body = v
	return s
}

type CountCrowdRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	IsShow   *bool   `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
}

func (s CountCrowdRequest) String() string {
	return tea.Prettify(s)
}

func (s CountCrowdRequest) GoString() string {
	return s.String()
}

func (s *CountCrowdRequest) SetImageURL(v string) *CountCrowdRequest {
	s.ImageURL = &v
	return s
}

func (s *CountCrowdRequest) SetIsShow(v bool) *CountCrowdRequest {
	s.IsShow = &v
	return s
}

type CountCrowdAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	IsShow         *bool     `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
}

func (s CountCrowdAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CountCrowdAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CountCrowdAdvanceRequest) SetImageURLObject(v io.Reader) *CountCrowdAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *CountCrowdAdvanceRequest) SetIsShow(v bool) *CountCrowdAdvanceRequest {
	s.IsShow = &v
	return s
}

type CountCrowdResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CountCrowdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CountCrowdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *CountCrowdResponseBody) SetRequestId(v string) *CountCrowdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountCrowdResponseBody) SetData(v *CountCrowdResponseBodyData) *CountCrowdResponseBody {
	s.Data = v
	return s
}

type CountCrowdResponseBodyData struct {
	PeopleNumber *int32  `json:"PeopleNumber,omitempty" xml:"PeopleNumber,omitempty"`
	HotMap       *string `json:"HotMap,omitempty" xml:"HotMap,omitempty"`
}

func (s CountCrowdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CountCrowdResponseBodyData) GoString() string {
	return s.String()
}

func (s *CountCrowdResponseBodyData) SetPeopleNumber(v int32) *CountCrowdResponseBodyData {
	s.PeopleNumber = &v
	return s
}

func (s *CountCrowdResponseBodyData) SetHotMap(v string) *CountCrowdResponseBodyData {
	s.HotMap = &v
	return s
}

type CountCrowdResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CountCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountCrowdResponse) String() string {
	return tea.Prettify(s)
}

func (s CountCrowdResponse) GoString() string {
	return s.String()
}

func (s *CountCrowdResponse) SetHeaders(v map[string]*string) *CountCrowdResponse {
	s.Headers = v
	return s
}

func (s *CountCrowdResponse) SetBody(v *CountCrowdResponseBody) *CountCrowdResponse {
	s.Body = v
	return s
}

type AddFaceEntityRequest struct {
	DbName   *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Labels   *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s AddFaceEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceEntityRequest) GoString() string {
	return s.String()
}

func (s *AddFaceEntityRequest) SetDbName(v string) *AddFaceEntityRequest {
	s.DbName = &v
	return s
}

func (s *AddFaceEntityRequest) SetEntityId(v string) *AddFaceEntityRequest {
	s.EntityId = &v
	return s
}

func (s *AddFaceEntityRequest) SetLabels(v string) *AddFaceEntityRequest {
	s.Labels = &v
	return s
}

type AddFaceEntityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFaceEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceEntityResponseBody) SetRequestId(v string) *AddFaceEntityResponseBody {
	s.RequestId = &v
	return s
}

type AddFaceEntityResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFaceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddFaceEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceEntityResponse) GoString() string {
	return s.String()
}

func (s *AddFaceEntityResponse) SetHeaders(v map[string]*string) *AddFaceEntityResponse {
	s.Headers = v
	return s
}

func (s *AddFaceEntityResponse) SetBody(v *AddFaceEntityResponseBody) *AddFaceEntityResponse {
	s.Body = v
	return s
}

type DeleteFaceEntityRequest struct {
	DbName   *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s DeleteFaceEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceEntityRequest) SetDbName(v string) *DeleteFaceEntityRequest {
	s.DbName = &v
	return s
}

func (s *DeleteFaceEntityRequest) SetEntityId(v string) *DeleteFaceEntityRequest {
	s.EntityId = &v
	return s
}

type DeleteFaceEntityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceEntityResponseBody) SetRequestId(v string) *DeleteFaceEntityResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFaceEntityResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFaceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFaceEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceEntityResponse) SetHeaders(v map[string]*string) *DeleteFaceEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceEntityResponse) SetBody(v *DeleteFaceEntityResponseBody) *DeleteFaceEntityResponse {
	s.Body = v
	return s
}

type FaceTidyupRequest struct {
	ImageURL  *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ShapeType *int32   `json:"ShapeType,omitempty" xml:"ShapeType,omitempty"`
	Strength  *float32 `json:"Strength,omitempty" xml:"Strength,omitempty"`
}

func (s FaceTidyupRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceTidyupRequest) GoString() string {
	return s.String()
}

func (s *FaceTidyupRequest) SetImageURL(v string) *FaceTidyupRequest {
	s.ImageURL = &v
	return s
}

func (s *FaceTidyupRequest) SetShapeType(v int32) *FaceTidyupRequest {
	s.ShapeType = &v
	return s
}

func (s *FaceTidyupRequest) SetStrength(v float32) *FaceTidyupRequest {
	s.Strength = &v
	return s
}

type FaceTidyupAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ShapeType      *int32    `json:"ShapeType,omitempty" xml:"ShapeType,omitempty"`
	Strength       *float32  `json:"Strength,omitempty" xml:"Strength,omitempty"`
}

func (s FaceTidyupAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceTidyupAdvanceRequest) GoString() string {
	return s.String()
}

func (s *FaceTidyupAdvanceRequest) SetImageURLObject(v io.Reader) *FaceTidyupAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *FaceTidyupAdvanceRequest) SetShapeType(v int32) *FaceTidyupAdvanceRequest {
	s.ShapeType = &v
	return s
}

func (s *FaceTidyupAdvanceRequest) SetStrength(v float32) *FaceTidyupAdvanceRequest {
	s.Strength = &v
	return s
}

type FaceTidyupResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *FaceTidyupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s FaceTidyupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceTidyupResponseBody) GoString() string {
	return s.String()
}

func (s *FaceTidyupResponseBody) SetRequestId(v string) *FaceTidyupResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceTidyupResponseBody) SetData(v *FaceTidyupResponseBodyData) *FaceTidyupResponseBody {
	s.Data = v
	return s
}

type FaceTidyupResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s FaceTidyupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FaceTidyupResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceTidyupResponseBodyData) SetImageURL(v string) *FaceTidyupResponseBodyData {
	s.ImageURL = &v
	return s
}

type FaceTidyupResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FaceTidyupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FaceTidyupResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceTidyupResponse) GoString() string {
	return s.String()
}

func (s *FaceTidyupResponse) SetHeaders(v map[string]*string) *FaceTidyupResponse {
	s.Headers = v
	return s
}

func (s *FaceTidyupResponse) SetBody(v *FaceTidyupResponseBody) *FaceTidyupResponse {
	s.Body = v
	return s
}

type BodyPostureRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BodyPostureRequest) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureRequest) GoString() string {
	return s.String()
}

func (s *BodyPostureRequest) SetImageURL(v string) *BodyPostureRequest {
	s.ImageURL = &v
	return s
}

type BodyPostureAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s BodyPostureAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BodyPostureAdvanceRequest) SetImageURLObject(v io.Reader) *BodyPostureAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type BodyPostureResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *BodyPostureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s BodyPostureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBody) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBody) SetRequestId(v string) *BodyPostureResponseBody {
	s.RequestId = &v
	return s
}

func (s *BodyPostureResponseBody) SetData(v *BodyPostureResponseBodyData) *BodyPostureResponseBody {
	s.Data = v
	return s
}

type BodyPostureResponseBodyData struct {
	Outputs    []*BodyPostureResponseBodyDataOutputs  `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	MetaObject *BodyPostureResponseBodyDataMetaObject `json:"MetaObject,omitempty" xml:"MetaObject,omitempty" type:"Struct"`
}

func (s BodyPostureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBodyData) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyData) SetOutputs(v []*BodyPostureResponseBodyDataOutputs) *BodyPostureResponseBodyData {
	s.Outputs = v
	return s
}

func (s *BodyPostureResponseBodyData) SetMetaObject(v *BodyPostureResponseBodyDataMetaObject) *BodyPostureResponseBodyData {
	s.MetaObject = v
	return s
}

type BodyPostureResponseBodyDataOutputs struct {
	HumanCount *int32                                       `json:"HumanCount,omitempty" xml:"HumanCount,omitempty"`
	Results    []*BodyPostureResponseBodyDataOutputsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BodyPostureResponseBodyDataOutputs) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBodyDataOutputs) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataOutputs) SetHumanCount(v int32) *BodyPostureResponseBodyDataOutputs {
	s.HumanCount = &v
	return s
}

func (s *BodyPostureResponseBodyDataOutputs) SetResults(v []*BodyPostureResponseBodyDataOutputsResults) *BodyPostureResponseBodyDataOutputs {
	s.Results = v
	return s
}

type BodyPostureResponseBodyDataOutputsResults struct {
	Bodies []*BodyPostureResponseBodyDataOutputsResultsBodies `json:"Bodies,omitempty" xml:"Bodies,omitempty" type:"Repeated"`
}

func (s BodyPostureResponseBodyDataOutputsResults) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBodyDataOutputsResults) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataOutputsResults) SetBodies(v []*BodyPostureResponseBodyDataOutputsResultsBodies) *BodyPostureResponseBodyDataOutputsResults {
	s.Bodies = v
	return s
}

type BodyPostureResponseBodyDataOutputsResultsBodies struct {
	Positions []*BodyPostureResponseBodyDataOutputsResultsBodiesPositions `json:"Positions,omitempty" xml:"Positions,omitempty" type:"Repeated"`
	Confident *float32                                                    `json:"Confident,omitempty" xml:"Confident,omitempty"`
	Label     *string                                                     `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s BodyPostureResponseBodyDataOutputsResultsBodies) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBodyDataOutputsResultsBodies) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodies) SetPositions(v []*BodyPostureResponseBodyDataOutputsResultsBodiesPositions) *BodyPostureResponseBodyDataOutputsResultsBodies {
	s.Positions = v
	return s
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodies) SetConfident(v float32) *BodyPostureResponseBodyDataOutputsResultsBodies {
	s.Confident = &v
	return s
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodies) SetLabel(v string) *BodyPostureResponseBodyDataOutputsResultsBodies {
	s.Label = &v
	return s
}

type BodyPostureResponseBodyDataOutputsResultsBodiesPositions struct {
	Points []*float32 `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s BodyPostureResponseBodyDataOutputsResultsBodiesPositions) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBodyDataOutputsResultsBodiesPositions) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodiesPositions) SetPoints(v []*float32) *BodyPostureResponseBodyDataOutputsResultsBodiesPositions {
	s.Points = v
	return s
}

type BodyPostureResponseBodyDataMetaObject struct {
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s BodyPostureResponseBodyDataMetaObject) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBodyDataMetaObject) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataMetaObject) SetWidth(v int32) *BodyPostureResponseBodyDataMetaObject {
	s.Width = &v
	return s
}

func (s *BodyPostureResponseBodyDataMetaObject) SetHeight(v int32) *BodyPostureResponseBodyDataMetaObject {
	s.Height = &v
	return s
}

type BodyPostureResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BodyPostureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BodyPostureResponse) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponse) GoString() string {
	return s.String()
}

func (s *BodyPostureResponse) SetHeaders(v map[string]*string) *BodyPostureResponse {
	s.Headers = v
	return s
}

func (s *BodyPostureResponse) SetBody(v *BodyPostureResponseBody) *BodyPostureResponse {
	s.Body = v
	return s
}

type CreateBodyDbRequest struct {
	// 数据库名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateBodyDbRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBodyDbRequest) GoString() string {
	return s.String()
}

func (s *CreateBodyDbRequest) SetName(v string) *CreateBodyDbRequest {
	s.Name = &v
	return s
}

type CreateBodyDbResponseBody struct {
	// RequestId
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateBodyDbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CreateBodyDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBodyDbResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBodyDbResponseBody) SetRequestId(v string) *CreateBodyDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBodyDbResponseBody) SetData(v *CreateBodyDbResponseBodyData) *CreateBodyDbResponseBody {
	s.Data = v
	return s
}

type CreateBodyDbResponseBodyData struct {
	// 数据库ID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateBodyDbResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateBodyDbResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateBodyDbResponseBodyData) SetId(v int64) *CreateBodyDbResponseBodyData {
	s.Id = &v
	return s
}

type CreateBodyDbResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBodyDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBodyDbResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBodyDbResponse) GoString() string {
	return s.String()
}

func (s *CreateBodyDbResponse) SetHeaders(v map[string]*string) *CreateBodyDbResponse {
	s.Headers = v
	return s
}

func (s *CreateBodyDbResponse) SetBody(v *CreateBodyDbResponseBody) *CreateBodyDbResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("facebody"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtractPedestrianFeatureAttrWithOptions(request *ExtractPedestrianFeatureAttrRequest, runtime *util.RuntimeOptions) (_result *ExtractPedestrianFeatureAttrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExtractPedestrianFeatureAttrResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExtractPedestrianFeatureAttr"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtractPedestrianFeatureAttr(request *ExtractPedestrianFeatureAttrRequest) (_result *ExtractPedestrianFeatureAttrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtractPedestrianFeatureAttrResponse{}
	_body, _err := client.ExtractPedestrianFeatureAttrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtractPedestrianFeatureAttrAdvance(request *ExtractPedestrianFeatureAttrAdvanceRequest, runtime *util.RuntimeOptions) (_result *ExtractPedestrianFeatureAttrResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	extractPedestrianFeatureAttrReq := &ExtractPedestrianFeatureAttrRequest{}
	openapiutil.Convert(request, extractPedestrianFeatureAttrReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	extractPedestrianFeatureAttrReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	extractPedestrianFeatureAttrResp, _err := client.ExtractPedestrianFeatureAttrWithOptions(extractPedestrianFeatureAttrReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = extractPedestrianFeatureAttrResp
	return _result, _err
}

func (client *Client) DetectBodyCountWithOptions(request *DetectBodyCountRequest, runtime *util.RuntimeOptions) (_result *DetectBodyCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectBodyCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectBodyCount"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectBodyCount(request *DetectBodyCountRequest) (_result *DetectBodyCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectBodyCountResponse{}
	_body, _err := client.DetectBodyCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectBodyCountAdvance(request *DetectBodyCountAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectBodyCountResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectBodyCountReq := &DetectBodyCountRequest{}
	openapiutil.Convert(request, detectBodyCountReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectBodyCountReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectBodyCountResp, _err := client.DetectBodyCountWithOptions(detectBodyCountReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectBodyCountResp
	return _result, _err
}

func (client *Client) DetectVideoLivingFaceWithOptions(request *DetectVideoLivingFaceRequest, runtime *util.RuntimeOptions) (_result *DetectVideoLivingFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectVideoLivingFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectVideoLivingFace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectVideoLivingFace(request *DetectVideoLivingFaceRequest) (_result *DetectVideoLivingFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectVideoLivingFaceResponse{}
	_body, _err := client.DetectVideoLivingFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectVideoLivingFaceAdvance(request *DetectVideoLivingFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectVideoLivingFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectVideoLivingFaceReq := &DetectVideoLivingFaceRequest{}
	openapiutil.Convert(request, detectVideoLivingFaceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.VideoUrlObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectVideoLivingFaceReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectVideoLivingFaceResp, _err := client.DetectVideoLivingFaceWithOptions(detectVideoLivingFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectVideoLivingFaceResp
	return _result, _err
}

func (client *Client) RecognizeFaceWithOptions(request *RecognizeFaceRequest, runtime *util.RuntimeOptions) (_result *RecognizeFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeFace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeFace(request *RecognizeFaceRequest) (_result *RecognizeFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeFaceResponse{}
	_body, _err := client.RecognizeFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeFaceAdvance(request *RecognizeFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeFaceReq := &RecognizeFaceRequest{}
	openapiutil.Convert(request, recognizeFaceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeFaceResp, _err := client.RecognizeFaceWithOptions(recognizeFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeFaceResp
	return _result, _err
}

func (client *Client) VerifyFaceMaskWithOptions(request *VerifyFaceMaskRequest, runtime *util.RuntimeOptions) (_result *VerifyFaceMaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyFaceMaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyFaceMask"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyFaceMask(request *VerifyFaceMaskRequest) (_result *VerifyFaceMaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyFaceMaskResponse{}
	_body, _err := client.VerifyFaceMaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyFaceMaskAdvance(request *VerifyFaceMaskAdvanceRequest, runtime *util.RuntimeOptions) (_result *VerifyFaceMaskResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	verifyFaceMaskReq := &VerifyFaceMaskRequest{}
	openapiutil.Convert(request, verifyFaceMaskReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	verifyFaceMaskReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	verifyFaceMaskResp, _err := client.VerifyFaceMaskWithOptions(verifyFaceMaskReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = verifyFaceMaskResp
	return _result, _err
}

func (client *Client) DetectIPCPedestrianWithOptions(request *DetectIPCPedestrianRequest, runtime *util.RuntimeOptions) (_result *DetectIPCPedestrianResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectIPCPedestrianResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectIPCPedestrian"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectIPCPedestrian(request *DetectIPCPedestrianRequest) (_result *DetectIPCPedestrianResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectIPCPedestrianResponse{}
	_body, _err := client.DetectIPCPedestrianWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFaceEntityWithOptions(request *GetFaceEntityRequest, runtime *util.RuntimeOptions) (_result *GetFaceEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetFaceEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetFaceEntity"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFaceEntity(request *GetFaceEntityRequest) (_result *GetFaceEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFaceEntityResponse{}
	_body, _err := client.GetFaceEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompareFaceWithOptions(request *CompareFaceRequest, runtime *util.RuntimeOptions) (_result *CompareFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CompareFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CompareFace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompareFace(request *CompareFaceRequest) (_result *CompareFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareFaceResponse{}
	_body, _err := client.CompareFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PedestrianDetectAttributeWithOptions(request *PedestrianDetectAttributeRequest, runtime *util.RuntimeOptions) (_result *PedestrianDetectAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PedestrianDetectAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PedestrianDetectAttribute"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PedestrianDetectAttribute(request *PedestrianDetectAttributeRequest) (_result *PedestrianDetectAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PedestrianDetectAttributeResponse{}
	_body, _err := client.PedestrianDetectAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PedestrianDetectAttributeAdvance(request *PedestrianDetectAttributeAdvanceRequest, runtime *util.RuntimeOptions) (_result *PedestrianDetectAttributeResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	pedestrianDetectAttributeReq := &PedestrianDetectAttributeRequest{}
	openapiutil.Convert(request, pedestrianDetectAttributeReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	pedestrianDetectAttributeReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	pedestrianDetectAttributeResp, _err := client.PedestrianDetectAttributeWithOptions(pedestrianDetectAttributeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = pedestrianDetectAttributeResp
	return _result, _err
}

func (client *Client) FaceFilterWithOptions(request *FaceFilterRequest, runtime *util.RuntimeOptions) (_result *FaceFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FaceFilterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FaceFilter"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceFilter(request *FaceFilterRequest) (_result *FaceFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceFilterResponse{}
	_body, _err := client.FaceFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceFilterAdvance(request *FaceFilterAdvanceRequest, runtime *util.RuntimeOptions) (_result *FaceFilterResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	faceFilterReq := &FaceFilterRequest{}
	openapiutil.Convert(request, faceFilterReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	faceFilterReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	faceFilterResp, _err := client.FaceFilterWithOptions(faceFilterReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = faceFilterResp
	return _result, _err
}

func (client *Client) FaceBeautyWithOptions(request *FaceBeautyRequest, runtime *util.RuntimeOptions) (_result *FaceBeautyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FaceBeautyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FaceBeauty"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceBeauty(request *FaceBeautyRequest) (_result *FaceBeautyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceBeautyResponse{}
	_body, _err := client.FaceBeautyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceBeautyAdvance(request *FaceBeautyAdvanceRequest, runtime *util.RuntimeOptions) (_result *FaceBeautyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	faceBeautyReq := &FaceBeautyRequest{}
	openapiutil.Convert(request, faceBeautyReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	faceBeautyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	faceBeautyResp, _err := client.FaceBeautyWithOptions(faceBeautyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = faceBeautyResp
	return _result, _err
}

func (client *Client) GenerateHumanAnimeStyleWithOptions(request *GenerateHumanAnimeStyleRequest, runtime *util.RuntimeOptions) (_result *GenerateHumanAnimeStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateHumanAnimeStyleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateHumanAnimeStyle"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateHumanAnimeStyle(request *GenerateHumanAnimeStyleRequest) (_result *GenerateHumanAnimeStyleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateHumanAnimeStyleResponse{}
	_body, _err := client.GenerateHumanAnimeStyleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateHumanAnimeStyleAdvance(request *GenerateHumanAnimeStyleAdvanceRequest, runtime *util.RuntimeOptions) (_result *GenerateHumanAnimeStyleResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	generateHumanAnimeStyleReq := &GenerateHumanAnimeStyleRequest{}
	openapiutil.Convert(request, generateHumanAnimeStyleReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	generateHumanAnimeStyleReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	generateHumanAnimeStyleResp, _err := client.GenerateHumanAnimeStyleWithOptions(generateHumanAnimeStyleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generateHumanAnimeStyleResp
	return _result, _err
}

func (client *Client) QueryFaceImageTemplateWithOptions(request *QueryFaceImageTemplateRequest, runtime *util.RuntimeOptions) (_result *QueryFaceImageTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &QueryFaceImageTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryFaceImageTemplate"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceImageTemplate(request *QueryFaceImageTemplateRequest) (_result *QueryFaceImageTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceImageTemplateResponse{}
	_body, _err := client.QueryFaceImageTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectFaceWithOptions(request *DetectFaceRequest, runtime *util.RuntimeOptions) (_result *DetectFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectFace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectFace(request *DetectFaceRequest) (_result *DetectFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectFaceResponse{}
	_body, _err := client.DetectFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectFaceAdvance(request *DetectFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectFaceReq := &DetectFaceRequest{}
	openapiutil.Convert(request, detectFaceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectFaceResp, _err := client.DetectFaceWithOptions(detectFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectFaceResp
	return _result, _err
}

func (client *Client) DetectMaskWithOptions(request *DetectMaskRequest, runtime *util.RuntimeOptions) (_result *DetectMaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectMaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectMask"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectMask(request *DetectMaskRequest) (_result *DetectMaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectMaskResponse{}
	_body, _err := client.DetectMaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectMaskAdvance(request *DetectMaskAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectMaskResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectMaskReq := &DetectMaskRequest{}
	openapiutil.Convert(request, detectMaskReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectMaskReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectMaskResp, _err := client.DetectMaskWithOptions(detectMaskReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectMaskResp
	return _result, _err
}

func (client *Client) GenRealPersonVerificationTokenWithOptions(request *GenRealPersonVerificationTokenRequest, runtime *util.RuntimeOptions) (_result *GenRealPersonVerificationTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenRealPersonVerificationTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenRealPersonVerificationToken"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenRealPersonVerificationToken(request *GenRealPersonVerificationTokenRequest) (_result *GenRealPersonVerificationTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenRealPersonVerificationTokenResponse{}
	_body, _err := client.GenRealPersonVerificationTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFaceDbsWithOptions(runtime *util.RuntimeOptions) (_result *ListFaceDbsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListFaceDbsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFaceDbs"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFaceDbs() (_result *ListFaceDbsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFaceDbsResponse{}
	_body, _err := client.ListFaceDbsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeActionWithOptions(request *RecognizeActionRequest, runtime *util.RuntimeOptions) (_result *RecognizeActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeActionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeAction"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeAction(request *RecognizeActionRequest) (_result *RecognizeActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeActionResponse{}
	_body, _err := client.RecognizeActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectChefCapWithOptions(request *DetectChefCapRequest, runtime *util.RuntimeOptions) (_result *DetectChefCapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectChefCapResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectChefCap"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectChefCap(request *DetectChefCapRequest) (_result *DetectChefCapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectChefCapResponse{}
	_body, _err := client.DetectChefCapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectChefCapAdvance(request *DetectChefCapAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectChefCapResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectChefCapReq := &DetectChefCapRequest{}
	openapiutil.Convert(request, detectChefCapReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectChefCapReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectChefCapResp, _err := client.DetectChefCapWithOptions(detectChefCapReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectChefCapResp
	return _result, _err
}

func (client *Client) DetectLivingFaceWithOptions(request *DetectLivingFaceRequest, runtime *util.RuntimeOptions) (_result *DetectLivingFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectLivingFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectLivingFace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectLivingFace(request *DetectLivingFaceRequest) (_result *DetectLivingFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectLivingFaceResponse{}
	_body, _err := client.DetectLivingFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectCelebrityWithOptions(request *DetectCelebrityRequest, runtime *util.RuntimeOptions) (_result *DetectCelebrityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectCelebrityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectCelebrity"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectCelebrity(request *DetectCelebrityRequest) (_result *DetectCelebrityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectCelebrityResponse{}
	_body, _err := client.DetectCelebrityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectCelebrityAdvance(request *DetectCelebrityAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectCelebrityResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectCelebrityReq := &DetectCelebrityRequest{}
	openapiutil.Convert(request, detectCelebrityReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectCelebrityReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectCelebrityResp, _err := client.DetectCelebrityWithOptions(detectCelebrityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectCelebrityResp
	return _result, _err
}

func (client *Client) GetRealPersonVerificationResultWithOptions(request *GetRealPersonVerificationResultRequest, runtime *util.RuntimeOptions) (_result *GetRealPersonVerificationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRealPersonVerificationResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRealPersonVerificationResult"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRealPersonVerificationResult(request *GetRealPersonVerificationResultRequest) (_result *GetRealPersonVerificationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRealPersonVerificationResultResponse{}
	_body, _err := client.GetRealPersonVerificationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceWithOptions(request *DeleteFaceRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFace(request *DeleteFaceRequest) (_result *DeleteFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceResponse{}
	_body, _err := client.DeleteFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtractPedestrianFeatureAttributeWithOptions(request *ExtractPedestrianFeatureAttributeRequest, runtime *util.RuntimeOptions) (_result *ExtractPedestrianFeatureAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExtractPedestrianFeatureAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExtractPedestrianFeatureAttribute"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtractPedestrianFeatureAttribute(request *ExtractPedestrianFeatureAttributeRequest) (_result *ExtractPedestrianFeatureAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtractPedestrianFeatureAttributeResponse{}
	_body, _err := client.ExtractPedestrianFeatureAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeExpressionWithOptions(request *RecognizeExpressionRequest, runtime *util.RuntimeOptions) (_result *RecognizeExpressionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeExpressionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeExpression"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeExpression(request *RecognizeExpressionRequest) (_result *RecognizeExpressionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeExpressionResponse{}
	_body, _err := client.RecognizeExpressionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeExpressionAdvance(request *RecognizeExpressionAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeExpressionResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeExpressionReq := &RecognizeExpressionRequest{}
	openapiutil.Convert(request, recognizeExpressionReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeExpressionReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeExpressionResp, _err := client.RecognizeExpressionWithOptions(recognizeExpressionReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeExpressionResp
	return _result, _err
}

func (client *Client) MergeImageFaceWithOptions(request *MergeImageFaceRequest, runtime *util.RuntimeOptions) (_result *MergeImageFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MergeImageFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MergeImageFace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MergeImageFace(request *MergeImageFaceRequest) (_result *MergeImageFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MergeImageFaceResponse{}
	_body, _err := client.MergeImageFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MergeImageFaceAdvance(request *MergeImageFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *MergeImageFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	mergeImageFaceReq := &MergeImageFaceRequest{}
	openapiutil.Convert(request, mergeImageFaceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	mergeImageFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	mergeImageFaceResp, _err := client.MergeImageFaceWithOptions(mergeImageFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = mergeImageFaceResp
	return _result, _err
}

func (client *Client) DeleteBodyPersonWithOptions(request *DeleteBodyPersonRequest, runtime *util.RuntimeOptions) (_result *DeleteBodyPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteBodyPersonResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteBodyPerson"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBodyPerson(request *DeleteBodyPersonRequest) (_result *DeleteBodyPersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBodyPersonResponse{}
	_body, _err := client.DeleteBodyPersonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectPedestrianWithOptions(request *DetectPedestrianRequest, runtime *util.RuntimeOptions) (_result *DetectPedestrianResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectPedestrianResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectPedestrian"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectPedestrian(request *DetectPedestrianRequest) (_result *DetectPedestrianResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectPedestrianResponse{}
	_body, _err := client.DetectPedestrianWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectPedestrianAdvance(request *DetectPedestrianAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectPedestrianResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectPedestrianReq := &DetectPedestrianRequest{}
	openapiutil.Convert(request, detectPedestrianReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectPedestrianReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectPedestrianResp, _err := client.DetectPedestrianWithOptions(detectPedestrianReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectPedestrianResp
	return _result, _err
}

func (client *Client) SwapFacialFeaturesWithOptions(request *SwapFacialFeaturesRequest, runtime *util.RuntimeOptions) (_result *SwapFacialFeaturesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SwapFacialFeaturesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SwapFacialFeatures"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwapFacialFeatures(request *SwapFacialFeaturesRequest) (_result *SwapFacialFeaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwapFacialFeaturesResponse{}
	_body, _err := client.SwapFacialFeaturesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwapFacialFeaturesAdvance(request *SwapFacialFeaturesAdvanceRequest, runtime *util.RuntimeOptions) (_result *SwapFacialFeaturesResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	swapFacialFeaturesReq := &SwapFacialFeaturesRequest{}
	openapiutil.Convert(request, swapFacialFeaturesReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.SourceImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	swapFacialFeaturesReq.SourceImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	swapFacialFeaturesResp, _err := client.SwapFacialFeaturesWithOptions(swapFacialFeaturesReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = swapFacialFeaturesResp
	return _result, _err
}

func (client *Client) SearchFaceWithOptions(request *SearchFaceRequest, runtime *util.RuntimeOptions) (_result *SearchFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchFace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchFace(request *SearchFaceRequest) (_result *SearchFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchFaceResponse{}
	_body, _err := client.SearchFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchFaceAdvance(request *SearchFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *SearchFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	searchFaceReq := &SearchFaceRequest{}
	openapiutil.Convert(request, searchFaceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageUrlObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	searchFaceReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	searchFaceResp, _err := client.SearchFaceWithOptions(searchFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = searchFaceResp
	return _result, _err
}

func (client *Client) UpdateFaceEntityWithOptions(request *UpdateFaceEntityRequest, runtime *util.RuntimeOptions) (_result *UpdateFaceEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFaceEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFaceEntity"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFaceEntity(request *UpdateFaceEntityRequest) (_result *UpdateFaceEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFaceEntityResponse{}
	_body, _err := client.UpdateFaceEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BlurFaceWithOptions(request *BlurFaceRequest, runtime *util.RuntimeOptions) (_result *BlurFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BlurFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BlurFace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BlurFace(request *BlurFaceRequest) (_result *BlurFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BlurFaceResponse{}
	_body, _err := client.BlurFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BlurFaceAdvance(request *BlurFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *BlurFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	blurFaceReq := &BlurFaceRequest{}
	openapiutil.Convert(request, blurFaceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	blurFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	blurFaceResp, _err := client.BlurFaceWithOptions(blurFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = blurFaceResp
	return _result, _err
}

func (client *Client) FaceMakeupWithOptions(request *FaceMakeupRequest, runtime *util.RuntimeOptions) (_result *FaceMakeupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FaceMakeupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FaceMakeup"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceMakeup(request *FaceMakeupRequest) (_result *FaceMakeupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceMakeupResponse{}
	_body, _err := client.FaceMakeupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceMakeupAdvance(request *FaceMakeupAdvanceRequest, runtime *util.RuntimeOptions) (_result *FaceMakeupResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	faceMakeupReq := &FaceMakeupRequest{}
	openapiutil.Convert(request, faceMakeupReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	faceMakeupReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	faceMakeupResp, _err := client.FaceMakeupWithOptions(faceMakeupReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = faceMakeupResp
	return _result, _err
}

func (client *Client) CreateBodyPersonWithOptions(request *CreateBodyPersonRequest, runtime *util.RuntimeOptions) (_result *CreateBodyPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBodyPersonResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBodyPerson"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBodyPerson(request *CreateBodyPersonRequest) (_result *CreateBodyPersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBodyPersonResponse{}
	_body, _err := client.CreateBodyPersonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceWithOptions(request *AddFaceRequest, runtime *util.RuntimeOptions) (_result *AddFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddFace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFace(request *AddFaceRequest) (_result *AddFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceResponse{}
	_body, _err := client.AddFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceAdvance(request *AddFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *AddFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	addFaceReq := &AddFaceRequest{}
	openapiutil.Convert(request, addFaceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageUrlObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	addFaceReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	addFaceResp, _err := client.AddFaceWithOptions(addFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addFaceResp
	return _result, _err
}

func (client *Client) GenerateHumanSketchStyleWithOptions(request *GenerateHumanSketchStyleRequest, runtime *util.RuntimeOptions) (_result *GenerateHumanSketchStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateHumanSketchStyleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateHumanSketchStyle"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateHumanSketchStyle(request *GenerateHumanSketchStyleRequest) (_result *GenerateHumanSketchStyleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateHumanSketchStyleResponse{}
	_body, _err := client.GenerateHumanSketchStyleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateHumanSketchStyleAdvance(request *GenerateHumanSketchStyleAdvanceRequest, runtime *util.RuntimeOptions) (_result *GenerateHumanSketchStyleResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	generateHumanSketchStyleReq := &GenerateHumanSketchStyleRequest{}
	openapiutil.Convert(request, generateHumanSketchStyleReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	generateHumanSketchStyleReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	generateHumanSketchStyleResp, _err := client.GenerateHumanSketchStyleWithOptions(generateHumanSketchStyleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generateHumanSketchStyleResp
	return _result, _err
}

func (client *Client) DeleteBodyDbWithOptions(request *DeleteBodyDbRequest, runtime *util.RuntimeOptions) (_result *DeleteBodyDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteBodyDbResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteBodyDb"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBodyDb(request *DeleteBodyDbRequest) (_result *DeleteBodyDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBodyDbResponse{}
	_body, _err := client.DeleteBodyDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectPedestrianIntrusionWithOptions(tmpReq *DetectPedestrianIntrusionRequest, runtime *util.RuntimeOptions) (_result *DetectPedestrianIntrusionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectPedestrianIntrusionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DetectRegion)) {
		request.DetectRegionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DetectRegion, tea.String("DetectRegion"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectPedestrianIntrusionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectPedestrianIntrusion"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectPedestrianIntrusion(request *DetectPedestrianIntrusionRequest) (_result *DetectPedestrianIntrusionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectPedestrianIntrusionResponse{}
	_body, _err := client.DetectPedestrianIntrusionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectPedestrianIntrusionAdvance(request *DetectPedestrianIntrusionAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectPedestrianIntrusionResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectPedestrianIntrusionReq := &DetectPedestrianIntrusionRequest{}
	openapiutil.Convert(request, detectPedestrianIntrusionReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectPedestrianIntrusionReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectPedestrianIntrusionResp, _err := client.DetectPedestrianIntrusionWithOptions(detectPedestrianIntrusionReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectPedestrianIntrusionResp
	return _result, _err
}

func (client *Client) HandPostureWithOptions(request *HandPostureRequest, runtime *util.RuntimeOptions) (_result *HandPostureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &HandPostureResponse{}
	_body, _err := client.DoRPCRequest(tea.String("HandPosture"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HandPosture(request *HandPostureRequest) (_result *HandPostureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HandPostureResponse{}
	_body, _err := client.HandPostureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HandPostureAdvance(request *HandPostureAdvanceRequest, runtime *util.RuntimeOptions) (_result *HandPostureResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	handPostureReq := &HandPostureRequest{}
	openapiutil.Convert(request, handPostureReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	handPostureReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	handPostureResp, _err := client.HandPostureWithOptions(handPostureReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = handPostureResp
	return _result, _err
}

func (client *Client) EnhanceFaceWithOptions(request *EnhanceFaceRequest, runtime *util.RuntimeOptions) (_result *EnhanceFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnhanceFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnhanceFace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnhanceFace(request *EnhanceFaceRequest) (_result *EnhanceFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnhanceFaceResponse{}
	_body, _err := client.EnhanceFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnhanceFaceAdvance(request *EnhanceFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *EnhanceFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	enhanceFaceReq := &EnhanceFaceRequest{}
	openapiutil.Convert(request, enhanceFaceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	enhanceFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	enhanceFaceResp, _err := client.EnhanceFaceWithOptions(enhanceFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = enhanceFaceResp
	return _result, _err
}

func (client *Client) GetBodyPersonWithOptions(request *GetBodyPersonRequest, runtime *util.RuntimeOptions) (_result *GetBodyPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetBodyPersonResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetBodyPerson"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBodyPerson(request *GetBodyPersonRequest) (_result *GetBodyPersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBodyPersonResponse{}
	_body, _err := client.GetBodyPersonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceDbWithOptions(request *DeleteFaceDbRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFaceDbResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFaceDb"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaceDb(request *DeleteFaceDbRequest) (_result *DeleteFaceDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceDbResponse{}
	_body, _err := client.DeleteFaceDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBodyPersonWithOptions(request *ListBodyPersonRequest, runtime *util.RuntimeOptions) (_result *ListBodyPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListBodyPersonResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBodyPerson"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBodyPerson(request *ListBodyPersonRequest) (_result *ListBodyPersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBodyPersonResponse{}
	_body, _err := client.ListBodyPersonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBodyDbsWithOptions(request *ListBodyDbsRequest, runtime *util.RuntimeOptions) (_result *ListBodyDbsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListBodyDbsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBodyDbs"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBodyDbs(request *ListBodyDbsRequest) (_result *ListBodyDbsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBodyDbsResponse{}
	_body, _err := client.ListBodyDbsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFaceEntitiesWithOptions(request *ListFaceEntitiesRequest, runtime *util.RuntimeOptions) (_result *ListFaceEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFaceEntitiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFaceEntities"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFaceEntities(request *ListFaceEntitiesRequest) (_result *ListFaceEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFaceEntitiesResponse{}
	_body, _err := client.ListFaceEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizePublicFaceWithOptions(request *RecognizePublicFaceRequest, runtime *util.RuntimeOptions) (_result *RecognizePublicFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizePublicFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizePublicFace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizePublicFace(request *RecognizePublicFaceRequest) (_result *RecognizePublicFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizePublicFaceResponse{}
	_body, _err := client.RecognizePublicFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceImageTemplateWithOptions(request *DeleteFaceImageTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceImageTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFaceImageTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFaceImageTemplate"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaceImageTemplate(request *DeleteFaceImageTemplateRequest) (_result *DeleteFaceImageTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceImageTemplateResponse{}
	_body, _err := client.DeleteFaceImageTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFaceDbWithOptions(request *CreateFaceDbRequest, runtime *util.RuntimeOptions) (_result *CreateFaceDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFaceDbResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFaceDb"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFaceDb(request *CreateFaceDbRequest) (_result *CreateFaceDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFaceDbResponse{}
	_body, _err := client.CreateFaceDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddBodyTraceWithOptions(tmpReq *AddBodyTraceRequest, runtime *util.RuntimeOptions) (_result *AddBodyTraceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddBodyTraceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Images)) {
		request.ImagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Images, tea.String("Images"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddBodyTraceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddBodyTrace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddBodyTrace(request *AddBodyTraceRequest) (_result *AddBodyTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddBodyTraceResponse{}
	_body, _err := client.AddBodyTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchBodyTraceWithOptions(tmpReq *SearchBodyTraceRequest, runtime *util.RuntimeOptions) (_result *SearchBodyTraceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchBodyTraceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Images)) {
		request.ImagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Images, tea.String("Images"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchBodyTraceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchBodyTrace"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchBodyTrace(request *SearchBodyTraceRequest) (_result *SearchBodyTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchBodyTraceResponse{}
	_body, _err := client.SearchBodyTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceImageTemplateWithOptions(request *AddFaceImageTemplateRequest, runtime *util.RuntimeOptions) (_result *AddFaceImageTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddFaceImageTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddFaceImageTemplate"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFaceImageTemplate(request *AddFaceImageTemplateRequest) (_result *AddFaceImageTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceImageTemplateResponse{}
	_body, _err := client.AddFaceImageTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceImageTemplateAdvance(request *AddFaceImageTemplateAdvanceRequest, runtime *util.RuntimeOptions) (_result *AddFaceImageTemplateResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	addFaceImageTemplateReq := &AddFaceImageTemplateRequest{}
	openapiutil.Convert(request, addFaceImageTemplateReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	addFaceImageTemplateReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	addFaceImageTemplateResp, _err := client.AddFaceImageTemplateWithOptions(addFaceImageTemplateReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addFaceImageTemplateResp
	return _result, _err
}

func (client *Client) CountCrowdWithOptions(request *CountCrowdRequest, runtime *util.RuntimeOptions) (_result *CountCrowdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CountCrowdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CountCrowd"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountCrowd(request *CountCrowdRequest) (_result *CountCrowdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountCrowdResponse{}
	_body, _err := client.CountCrowdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountCrowdAdvance(request *CountCrowdAdvanceRequest, runtime *util.RuntimeOptions) (_result *CountCrowdResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	countCrowdReq := &CountCrowdRequest{}
	openapiutil.Convert(request, countCrowdReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	countCrowdReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	countCrowdResp, _err := client.CountCrowdWithOptions(countCrowdReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = countCrowdResp
	return _result, _err
}

func (client *Client) AddFaceEntityWithOptions(request *AddFaceEntityRequest, runtime *util.RuntimeOptions) (_result *AddFaceEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddFaceEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddFaceEntity"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFaceEntity(request *AddFaceEntityRequest) (_result *AddFaceEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceEntityResponse{}
	_body, _err := client.AddFaceEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceEntityWithOptions(request *DeleteFaceEntityRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFaceEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFaceEntity"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaceEntity(request *DeleteFaceEntityRequest) (_result *DeleteFaceEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceEntityResponse{}
	_body, _err := client.DeleteFaceEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceTidyupWithOptions(request *FaceTidyupRequest, runtime *util.RuntimeOptions) (_result *FaceTidyupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FaceTidyupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FaceTidyup"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceTidyup(request *FaceTidyupRequest) (_result *FaceTidyupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceTidyupResponse{}
	_body, _err := client.FaceTidyupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceTidyupAdvance(request *FaceTidyupAdvanceRequest, runtime *util.RuntimeOptions) (_result *FaceTidyupResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	faceTidyupReq := &FaceTidyupRequest{}
	openapiutil.Convert(request, faceTidyupReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	faceTidyupReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	faceTidyupResp, _err := client.FaceTidyupWithOptions(faceTidyupReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = faceTidyupResp
	return _result, _err
}

func (client *Client) BodyPostureWithOptions(request *BodyPostureRequest, runtime *util.RuntimeOptions) (_result *BodyPostureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BodyPostureResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BodyPosture"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BodyPosture(request *BodyPostureRequest) (_result *BodyPostureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BodyPostureResponse{}
	_body, _err := client.BodyPostureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BodyPostureAdvance(request *BodyPostureAdvanceRequest, runtime *util.RuntimeOptions) (_result *BodyPostureResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	bodyPostureReq := &BodyPostureRequest{}
	openapiutil.Convert(request, bodyPostureReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	bodyPostureReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	bodyPostureResp, _err := client.BodyPostureWithOptions(bodyPostureReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = bodyPostureResp
	return _result, _err
}

func (client *Client) CreateBodyDbWithOptions(request *CreateBodyDbRequest, runtime *util.RuntimeOptions) (_result *CreateBodyDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBodyDbResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBodyDb"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBodyDb(request *CreateBodyDbRequest) (_result *CreateBodyDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBodyDbResponse{}
	_body, _err := client.CreateBodyDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
