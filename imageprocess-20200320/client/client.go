// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type DetectRibFractureRequest struct {
	URLList    []*DetectRibFractureRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" require:"true" type:"Repeated"`
	DataFormat *string                            `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgName    *string                            `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	OrgId      *string                            `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	SourceType *string                            `json:"SourceType,omitempty" xml:"SourceType,omitempty" require:"true"`
}

func (s DetectRibFractureRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectRibFractureRequest) GoString() string {
	return s.String()
}

func (s *DetectRibFractureRequest) SetURLList(v []*DetectRibFractureRequestURLList) *DetectRibFractureRequest {
	s.URLList = v
	return s
}

func (s *DetectRibFractureRequest) SetDataFormat(v string) *DetectRibFractureRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectRibFractureRequest) SetOrgName(v string) *DetectRibFractureRequest {
	s.OrgName = &v
	return s
}

func (s *DetectRibFractureRequest) SetOrgId(v string) *DetectRibFractureRequest {
	s.OrgId = &v
	return s
}

func (s *DetectRibFractureRequest) SetSourceType(v string) *DetectRibFractureRequest {
	s.SourceType = &v
	return s
}

type DetectRibFractureRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty" require:"true"`
}

func (s DetectRibFractureRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectRibFractureRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectRibFractureRequestURLList) SetURL(v string) *DetectRibFractureRequestURLList {
	s.URL = &v
	return s
}

type DetectRibFractureResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *DetectRibFractureResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s DetectRibFractureResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectRibFractureResponse) GoString() string {
	return s.String()
}

func (s *DetectRibFractureResponse) SetRequestId(v string) *DetectRibFractureResponse {
	s.RequestId = &v
	return s
}

func (s *DetectRibFractureResponse) SetData(v *DetectRibFractureResponseData) *DetectRibFractureResponse {
	s.Data = v
	return s
}

type DetectRibFractureResponseData struct {
	ResultURL  *string                                    `json:"ResultURL,omitempty" xml:"ResultURL,omitempty" require:"true"`
	Detections []*DetectRibFractureResponseDataDetections `json:"Detections,omitempty" xml:"Detections,omitempty" require:"true" type:"Repeated"`
	Spacing    []*float32                                 `json:"Spacing,omitempty" xml:"Spacing,omitempty" require:"true" type:"Repeated"`
	Origin     []*float32                                 `json:"Origin,omitempty" xml:"Origin,omitempty" require:"true" type:"Repeated"`
}

func (s DetectRibFractureResponseData) String() string {
	return tea.Prettify(s)
}

func (s DetectRibFractureResponseData) GoString() string {
	return s.String()
}

func (s *DetectRibFractureResponseData) SetResultURL(v string) *DetectRibFractureResponseData {
	s.ResultURL = &v
	return s
}

func (s *DetectRibFractureResponseData) SetDetections(v []*DetectRibFractureResponseDataDetections) *DetectRibFractureResponseData {
	s.Detections = v
	return s
}

func (s *DetectRibFractureResponseData) SetSpacing(v []*float32) *DetectRibFractureResponseData {
	s.Spacing = v
	return s
}

func (s *DetectRibFractureResponseData) SetOrigin(v []*float32) *DetectRibFractureResponseData {
	s.Origin = v
	return s
}

type DetectRibFractureResponseDataDetections struct {
	FractureId         *int     `json:"FractureId,omitempty" xml:"FractureId,omitempty" require:"true"`
	FractureConfidence *float32 `json:"FractureConfidence,omitempty" xml:"FractureConfidence,omitempty" require:"true"`
	FractureCategory   *string  `json:"FractureCategory,omitempty" xml:"FractureCategory,omitempty" require:"true"`
	Coordinates        []*int   `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" require:"true" type:"Repeated"`
	CoordinateImage    []*int   `json:"CoordinateImage,omitempty" xml:"CoordinateImage,omitempty" require:"true" type:"Repeated"`
}

func (s DetectRibFractureResponseDataDetections) String() string {
	return tea.Prettify(s)
}

func (s DetectRibFractureResponseDataDetections) GoString() string {
	return s.String()
}

func (s *DetectRibFractureResponseDataDetections) SetFractureId(v int) *DetectRibFractureResponseDataDetections {
	s.FractureId = &v
	return s
}

func (s *DetectRibFractureResponseDataDetections) SetFractureConfidence(v float32) *DetectRibFractureResponseDataDetections {
	s.FractureConfidence = &v
	return s
}

func (s *DetectRibFractureResponseDataDetections) SetFractureCategory(v string) *DetectRibFractureResponseDataDetections {
	s.FractureCategory = &v
	return s
}

func (s *DetectRibFractureResponseDataDetections) SetCoordinates(v []*int) *DetectRibFractureResponseDataDetections {
	s.Coordinates = v
	return s
}

func (s *DetectRibFractureResponseDataDetections) SetCoordinateImage(v []*int) *DetectRibFractureResponseDataDetections {
	s.CoordinateImage = v
	return s
}

type ScreenChestCTRequest struct {
	DataFormat *string                        `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgName    *string                        `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	OrgId      *string                        `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	URLList    []*ScreenChestCTRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" require:"true" type:"Repeated"`
}

func (s ScreenChestCTRequest) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTRequest) GoString() string {
	return s.String()
}

func (s *ScreenChestCTRequest) SetDataFormat(v string) *ScreenChestCTRequest {
	s.DataFormat = &v
	return s
}

func (s *ScreenChestCTRequest) SetOrgName(v string) *ScreenChestCTRequest {
	s.OrgName = &v
	return s
}

func (s *ScreenChestCTRequest) SetOrgId(v string) *ScreenChestCTRequest {
	s.OrgId = &v
	return s
}

func (s *ScreenChestCTRequest) SetURLList(v []*ScreenChestCTRequestURLList) *ScreenChestCTRequest {
	s.URLList = v
	return s
}

type ScreenChestCTRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty" require:"true"`
}

func (s ScreenChestCTRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTRequestURLList) GoString() string {
	return s.String()
}

func (s *ScreenChestCTRequestURLList) SetURL(v string) *ScreenChestCTRequestURLList {
	s.URL = &v
	return s
}

type ScreenChestCTResponse struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ScreenChestCTResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ScreenChestCTResponse) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponse) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponse) SetRequestId(v string) *ScreenChestCTResponse {
	s.RequestId = &v
	return s
}

func (s *ScreenChestCTResponse) SetData(v *ScreenChestCTResponseData) *ScreenChestCTResponse {
	s.Data = v
	return s
}

type ScreenChestCTResponseData struct {
	LungNodule *ScreenChestCTResponseDataLungNodule `json:"LungNodule,omitempty" xml:"LungNodule,omitempty" require:"true" type:"Struct"`
	CACS       *ScreenChestCTResponseDataCACS       `json:"CACS,omitempty" xml:"CACS,omitempty" require:"true" type:"Struct"`
	Covid      *ScreenChestCTResponseDataCovid      `json:"Covid,omitempty" xml:"Covid,omitempty" require:"true" type:"Struct"`
}

func (s ScreenChestCTResponseData) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseData) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseData) SetLungNodule(v *ScreenChestCTResponseDataLungNodule) *ScreenChestCTResponseData {
	s.LungNodule = v
	return s
}

func (s *ScreenChestCTResponseData) SetCACS(v *ScreenChestCTResponseDataCACS) *ScreenChestCTResponseData {
	s.CACS = v
	return s
}

func (s *ScreenChestCTResponseData) SetCovid(v *ScreenChestCTResponseDataCovid) *ScreenChestCTResponseData {
	s.Covid = v
	return s
}

type ScreenChestCTResponseDataLungNodule struct {
	Series []*ScreenChestCTResponseDataLungNoduleSeries `json:"Series,omitempty" xml:"Series,omitempty" require:"true" type:"Repeated"`
}

func (s ScreenChestCTResponseDataLungNodule) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseDataLungNodule) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseDataLungNodule) SetSeries(v []*ScreenChestCTResponseDataLungNoduleSeries) *ScreenChestCTResponseDataLungNodule {
	s.Series = v
	return s
}

type ScreenChestCTResponseDataLungNoduleSeries struct {
	SeriesInstanceUid *string                                              `json:"SeriesInstanceUid,omitempty" xml:"SeriesInstanceUid,omitempty" require:"true"`
	Report            *string                                              `json:"Report,omitempty" xml:"Report,omitempty" require:"true"`
	Elements          []*ScreenChestCTResponseDataLungNoduleSeriesElements `json:"Elements,omitempty" xml:"Elements,omitempty" require:"true" type:"Repeated"`
	Origin            []*float32                                           `json:"Origin,omitempty" xml:"Origin,omitempty" require:"true" type:"Repeated"`
	Spacing           []*float32                                           `json:"Spacing,omitempty" xml:"Spacing,omitempty" require:"true" type:"Repeated"`
}

func (s ScreenChestCTResponseDataLungNoduleSeries) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseDataLungNoduleSeries) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseDataLungNoduleSeries) SetSeriesInstanceUid(v string) *ScreenChestCTResponseDataLungNoduleSeries {
	s.SeriesInstanceUid = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeries) SetReport(v string) *ScreenChestCTResponseDataLungNoduleSeries {
	s.Report = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeries) SetElements(v []*ScreenChestCTResponseDataLungNoduleSeriesElements) *ScreenChestCTResponseDataLungNoduleSeries {
	s.Elements = v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeries) SetOrigin(v []*float32) *ScreenChestCTResponseDataLungNoduleSeries {
	s.Origin = v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeries) SetSpacing(v []*float32) *ScreenChestCTResponseDataLungNoduleSeries {
	s.Spacing = v
	return s
}

type ScreenChestCTResponseDataLungNoduleSeriesElements struct {
	Category       *string  `json:"Category,omitempty" xml:"Category,omitempty" require:"true"`
	Confidence     *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty" require:"true"`
	Diameter       *float32 `json:"Diameter,omitempty" xml:"Diameter,omitempty" require:"true"`
	Lobe           *string  `json:"Lobe,omitempty" xml:"Lobe,omitempty" require:"true"`
	Lung           *string  `json:"Lung,omitempty" xml:"Lung,omitempty" require:"true"`
	X              *float32 `json:"X,omitempty" xml:"X,omitempty" require:"true"`
	Z              *float32 `json:"Z,omitempty" xml:"Z,omitempty" require:"true"`
	Y              *float32 `json:"Y,omitempty" xml:"Y,omitempty" require:"true"`
	ImageX         *float32 `json:"ImageX,omitempty" xml:"ImageX,omitempty" require:"true"`
	ImageY         *float32 `json:"ImageY,omitempty" xml:"ImageY,omitempty" require:"true"`
	ImageZ         *float32 `json:"ImageZ,omitempty" xml:"ImageZ,omitempty" require:"true"`
	SOPInstanceUID *string  `json:"SOPInstanceUID,omitempty" xml:"SOPInstanceUID,omitempty" require:"true"`
	Volume         *float32 `json:"Volume,omitempty" xml:"Volume,omitempty" require:"true"`
	MeanValue      *float32 `json:"MeanValue,omitempty" xml:"MeanValue,omitempty" require:"true"`
}

func (s ScreenChestCTResponseDataLungNoduleSeriesElements) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseDataLungNoduleSeriesElements) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetCategory(v string) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.Category = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetConfidence(v float32) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.Confidence = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetDiameter(v float32) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.Diameter = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetLobe(v string) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.Lobe = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetLung(v string) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.Lung = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetX(v float32) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.X = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetZ(v float32) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.Z = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetY(v float32) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.Y = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetImageX(v float32) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.ImageX = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetImageY(v float32) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.ImageY = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetImageZ(v float32) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.ImageZ = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetSOPInstanceUID(v string) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.SOPInstanceUID = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetVolume(v float32) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.Volume = &v
	return s
}

func (s *ScreenChestCTResponseDataLungNoduleSeriesElements) SetMeanValue(v float32) *ScreenChestCTResponseDataLungNoduleSeriesElements {
	s.MeanValue = &v
	return s
}

type ScreenChestCTResponseDataCACS struct {
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty" require:"true"`
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty" require:"true"`
}

func (s ScreenChestCTResponseDataCACS) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseDataCACS) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseDataCACS) SetScore(v string) *ScreenChestCTResponseDataCACS {
	s.Score = &v
	return s
}

func (s *ScreenChestCTResponseDataCACS) SetResultUrl(v string) *ScreenChestCTResponseDataCACS {
	s.ResultUrl = &v
	return s
}

type ScreenChestCTResponseDataCovid struct {
	NewProbability    *string `json:"NewProbability,omitempty" xml:"NewProbability,omitempty" require:"true"`
	NormalProbability *string `json:"NormalProbability,omitempty" xml:"NormalProbability,omitempty" require:"true"`
	OtherProbability  *string `json:"OtherProbability,omitempty" xml:"OtherProbability,omitempty" require:"true"`
	LesionRatio       *string `json:"LesionRatio,omitempty" xml:"LesionRatio,omitempty" require:"true"`
	Mask              *string `json:"Mask,omitempty" xml:"Mask,omitempty" require:"true"`
}

func (s ScreenChestCTResponseDataCovid) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseDataCovid) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseDataCovid) SetNewProbability(v string) *ScreenChestCTResponseDataCovid {
	s.NewProbability = &v
	return s
}

func (s *ScreenChestCTResponseDataCovid) SetNormalProbability(v string) *ScreenChestCTResponseDataCovid {
	s.NormalProbability = &v
	return s
}

func (s *ScreenChestCTResponseDataCovid) SetOtherProbability(v string) *ScreenChestCTResponseDataCovid {
	s.OtherProbability = &v
	return s
}

func (s *ScreenChestCTResponseDataCovid) SetLesionRatio(v string) *ScreenChestCTResponseDataCovid {
	s.LesionRatio = &v
	return s
}

func (s *ScreenChestCTResponseDataCovid) SetMask(v string) *ScreenChestCTResponseDataCovid {
	s.Mask = &v
	return s
}

type DetectSkinDiseaseRequest struct {
	Url     *string `json:"Url,omitempty" xml:"Url,omitempty" require:"true"`
	OrgId   *string `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
}

func (s DetectSkinDiseaseRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectSkinDiseaseRequest) GoString() string {
	return s.String()
}

func (s *DetectSkinDiseaseRequest) SetUrl(v string) *DetectSkinDiseaseRequest {
	s.Url = &v
	return s
}

func (s *DetectSkinDiseaseRequest) SetOrgId(v string) *DetectSkinDiseaseRequest {
	s.OrgId = &v
	return s
}

func (s *DetectSkinDiseaseRequest) SetOrgName(v string) *DetectSkinDiseaseRequest {
	s.OrgName = &v
	return s
}

type DetectSkinDiseaseResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *DetectSkinDiseaseResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s DetectSkinDiseaseResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectSkinDiseaseResponse) GoString() string {
	return s.String()
}

func (s *DetectSkinDiseaseResponse) SetRequestId(v string) *DetectSkinDiseaseResponse {
	s.RequestId = &v
	return s
}

func (s *DetectSkinDiseaseResponse) SetData(v *DetectSkinDiseaseResponseData) *DetectSkinDiseaseResponse {
	s.Data = v
	return s
}

type DetectSkinDiseaseResponseData struct {
	Results map[string]interface{} `json:"Results,omitempty" xml:"Results,omitempty" require:"true"`
}

func (s DetectSkinDiseaseResponseData) String() string {
	return tea.Prettify(s)
}

func (s DetectSkinDiseaseResponseData) GoString() string {
	return s.String()
}

func (s *DetectSkinDiseaseResponseData) SetResults(v map[string]interface{}) *DetectSkinDiseaseResponseData {
	s.Results = v
	return s
}

type DetectSkinDiseaseAdvanceRequest struct {
	UrlObject io.Reader `json:"UrlObject,omitempty" xml:"UrlObject,omitempty" require:"true"`
	OrgId     *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	OrgName   *string   `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
}

func (s DetectSkinDiseaseAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectSkinDiseaseAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectSkinDiseaseAdvanceRequest) SetUrlObject(v io.Reader) *DetectSkinDiseaseAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *DetectSkinDiseaseAdvanceRequest) SetOrgId(v string) *DetectSkinDiseaseAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *DetectSkinDiseaseAdvanceRequest) SetOrgName(v string) *DetectSkinDiseaseAdvanceRequest {
	s.OrgName = &v
	return s
}

type RunMedQARequest struct {
	Question *string `json:"Question,omitempty" xml:"Question,omitempty" require:"true"`
	OrgId    *string `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	OrgName  *string `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
}

func (s RunMedQARequest) String() string {
	return tea.Prettify(s)
}

func (s RunMedQARequest) GoString() string {
	return s.String()
}

func (s *RunMedQARequest) SetQuestion(v string) *RunMedQARequest {
	s.Question = &v
	return s
}

func (s *RunMedQARequest) SetOrgId(v string) *RunMedQARequest {
	s.OrgId = &v
	return s
}

func (s *RunMedQARequest) SetOrgName(v string) *RunMedQARequest {
	s.OrgName = &v
	return s
}

type RunMedQAResponse struct {
	RequestId *string               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *RunMedQAResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s RunMedQAResponse) String() string {
	return tea.Prettify(s)
}

func (s RunMedQAResponse) GoString() string {
	return s.String()
}

func (s *RunMedQAResponse) SetRequestId(v string) *RunMedQAResponse {
	s.RequestId = &v
	return s
}

func (s *RunMedQAResponse) SetData(v *RunMedQAResponseData) *RunMedQAResponse {
	s.Data = v
	return s
}

type RunMedQAResponseData struct {
	Answer          *string   `json:"Answer,omitempty" xml:"Answer,omitempty" require:"true"`
	SimilarQuestion []*string `json:"SimilarQuestion,omitempty" xml:"SimilarQuestion,omitempty" require:"true" type:"Repeated"`
}

func (s RunMedQAResponseData) String() string {
	return tea.Prettify(s)
}

func (s RunMedQAResponseData) GoString() string {
	return s.String()
}

func (s *RunMedQAResponseData) SetAnswer(v string) *RunMedQAResponseData {
	s.Answer = &v
	return s
}

func (s *RunMedQAResponseData) SetSimilarQuestion(v []*string) *RunMedQAResponseData {
	s.SimilarQuestion = v
	return s
}

type DetectKneeKeypointXRayRequest struct {
	ImageUrl   *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" require:"true"`
	DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	OrgName    *string `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	TracerId   *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DetectKneeKeypointXRayRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeKeypointXRayRequest) GoString() string {
	return s.String()
}

func (s *DetectKneeKeypointXRayRequest) SetImageUrl(v string) *DetectKneeKeypointXRayRequest {
	s.ImageUrl = &v
	return s
}

func (s *DetectKneeKeypointXRayRequest) SetDataFormat(v string) *DetectKneeKeypointXRayRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectKneeKeypointXRayRequest) SetOrgId(v string) *DetectKneeKeypointXRayRequest {
	s.OrgId = &v
	return s
}

func (s *DetectKneeKeypointXRayRequest) SetOrgName(v string) *DetectKneeKeypointXRayRequest {
	s.OrgName = &v
	return s
}

func (s *DetectKneeKeypointXRayRequest) SetTracerId(v string) *DetectKneeKeypointXRayRequest {
	s.TracerId = &v
	return s
}

type DetectKneeKeypointXRayResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *DetectKneeKeypointXRayResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s DetectKneeKeypointXRayResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeKeypointXRayResponse) GoString() string {
	return s.String()
}

func (s *DetectKneeKeypointXRayResponse) SetRequestId(v string) *DetectKneeKeypointXRayResponse {
	s.RequestId = &v
	return s
}

func (s *DetectKneeKeypointXRayResponse) SetData(v *DetectKneeKeypointXRayResponseData) *DetectKneeKeypointXRayResponse {
	s.Data = v
	return s
}

type DetectKneeKeypointXRayResponseData struct {
	ImageUrl  *string                                        `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" require:"true"`
	OrgId     *string                                        `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	OrgName   *string                                        `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	KeyPoints []*DetectKneeKeypointXRayResponseDataKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" require:"true" type:"Repeated"`
}

func (s DetectKneeKeypointXRayResponseData) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeKeypointXRayResponseData) GoString() string {
	return s.String()
}

func (s *DetectKneeKeypointXRayResponseData) SetImageUrl(v string) *DetectKneeKeypointXRayResponseData {
	s.ImageUrl = &v
	return s
}

func (s *DetectKneeKeypointXRayResponseData) SetOrgId(v string) *DetectKneeKeypointXRayResponseData {
	s.OrgId = &v
	return s
}

func (s *DetectKneeKeypointXRayResponseData) SetOrgName(v string) *DetectKneeKeypointXRayResponseData {
	s.OrgName = &v
	return s
}

func (s *DetectKneeKeypointXRayResponseData) SetKeyPoints(v []*DetectKneeKeypointXRayResponseDataKeyPoints) *DetectKneeKeypointXRayResponseData {
	s.KeyPoints = v
	return s
}

type DetectKneeKeypointXRayResponseDataKeyPoints struct {
	Value       *float32                                        `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
	Tag         *DetectKneeKeypointXRayResponseDataKeyPointsTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Struct"`
	Coordinates []*int                                          `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" require:"true" type:"Repeated"`
}

func (s DetectKneeKeypointXRayResponseDataKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeKeypointXRayResponseDataKeyPoints) GoString() string {
	return s.String()
}

func (s *DetectKneeKeypointXRayResponseDataKeyPoints) SetValue(v float32) *DetectKneeKeypointXRayResponseDataKeyPoints {
	s.Value = &v
	return s
}

func (s *DetectKneeKeypointXRayResponseDataKeyPoints) SetTag(v *DetectKneeKeypointXRayResponseDataKeyPointsTag) *DetectKneeKeypointXRayResponseDataKeyPoints {
	s.Tag = v
	return s
}

func (s *DetectKneeKeypointXRayResponseDataKeyPoints) SetCoordinates(v []*int) *DetectKneeKeypointXRayResponseDataKeyPoints {
	s.Coordinates = v
	return s
}

type DetectKneeKeypointXRayResponseDataKeyPointsTag struct {
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty" require:"true"`
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty" require:"true"`
}

func (s DetectKneeKeypointXRayResponseDataKeyPointsTag) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeKeypointXRayResponseDataKeyPointsTag) GoString() string {
	return s.String()
}

func (s *DetectKneeKeypointXRayResponseDataKeyPointsTag) SetDirection(v string) *DetectKneeKeypointXRayResponseDataKeyPointsTag {
	s.Direction = &v
	return s
}

func (s *DetectKneeKeypointXRayResponseDataKeyPointsTag) SetLabel(v string) *DetectKneeKeypointXRayResponseDataKeyPointsTag {
	s.Label = &v
	return s
}

type DetectKneeKeypointXRayAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	DataFormat     *string   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgId          *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	OrgName        *string   `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	TracerId       *string   `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DetectKneeKeypointXRayAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeKeypointXRayAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectKneeKeypointXRayAdvanceRequest) SetImageUrlObject(v io.Reader) *DetectKneeKeypointXRayAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *DetectKneeKeypointXRayAdvanceRequest) SetDataFormat(v string) *DetectKneeKeypointXRayAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectKneeKeypointXRayAdvanceRequest) SetOrgId(v string) *DetectKneeKeypointXRayAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *DetectKneeKeypointXRayAdvanceRequest) SetOrgName(v string) *DetectKneeKeypointXRayAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *DetectKneeKeypointXRayAdvanceRequest) SetTracerId(v string) *DetectKneeKeypointXRayAdvanceRequest {
	s.TracerId = &v
	return s
}

type ClassifyFNFRequest struct {
	ImageUrl   *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" require:"true"`
	DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	OrgName    *string `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	TracerId   *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ClassifyFNFRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyFNFRequest) GoString() string {
	return s.String()
}

func (s *ClassifyFNFRequest) SetImageUrl(v string) *ClassifyFNFRequest {
	s.ImageUrl = &v
	return s
}

func (s *ClassifyFNFRequest) SetDataFormat(v string) *ClassifyFNFRequest {
	s.DataFormat = &v
	return s
}

func (s *ClassifyFNFRequest) SetOrgId(v string) *ClassifyFNFRequest {
	s.OrgId = &v
	return s
}

func (s *ClassifyFNFRequest) SetOrgName(v string) *ClassifyFNFRequest {
	s.OrgName = &v
	return s
}

func (s *ClassifyFNFRequest) SetTracerId(v string) *ClassifyFNFRequest {
	s.TracerId = &v
	return s
}

type ClassifyFNFResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ClassifyFNFResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ClassifyFNFResponse) String() string {
	return tea.Prettify(s)
}

func (s ClassifyFNFResponse) GoString() string {
	return s.String()
}

func (s *ClassifyFNFResponse) SetRequestId(v string) *ClassifyFNFResponse {
	s.RequestId = &v
	return s
}

func (s *ClassifyFNFResponse) SetData(v *ClassifyFNFResponseData) *ClassifyFNFResponse {
	s.Data = v
	return s
}

type ClassifyFNFResponseData struct {
	ImageUrl  *string                             `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" require:"true"`
	OrgId     *string                             `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	OrgName   *string                             `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	Fractures []*ClassifyFNFResponseDataFractures `json:"Fractures,omitempty" xml:"Fractures,omitempty" require:"true" type:"Repeated"`
}

func (s ClassifyFNFResponseData) String() string {
	return tea.Prettify(s)
}

func (s ClassifyFNFResponseData) GoString() string {
	return s.String()
}

func (s *ClassifyFNFResponseData) SetImageUrl(v string) *ClassifyFNFResponseData {
	s.ImageUrl = &v
	return s
}

func (s *ClassifyFNFResponseData) SetOrgId(v string) *ClassifyFNFResponseData {
	s.OrgId = &v
	return s
}

func (s *ClassifyFNFResponseData) SetOrgName(v string) *ClassifyFNFResponseData {
	s.OrgName = &v
	return s
}

func (s *ClassifyFNFResponseData) SetFractures(v []*ClassifyFNFResponseDataFractures) *ClassifyFNFResponseData {
	s.Fractures = v
	return s
}

type ClassifyFNFResponseDataFractures struct {
	Value *float32                             `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
	Tag   *ClassifyFNFResponseDataFracturesTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Struct"`
	Boxes []*int                               `json:"Boxes,omitempty" xml:"Boxes,omitempty" require:"true" type:"Repeated"`
}

func (s ClassifyFNFResponseDataFractures) String() string {
	return tea.Prettify(s)
}

func (s ClassifyFNFResponseDataFractures) GoString() string {
	return s.String()
}

func (s *ClassifyFNFResponseDataFractures) SetValue(v float32) *ClassifyFNFResponseDataFractures {
	s.Value = &v
	return s
}

func (s *ClassifyFNFResponseDataFractures) SetTag(v *ClassifyFNFResponseDataFracturesTag) *ClassifyFNFResponseDataFractures {
	s.Tag = v
	return s
}

func (s *ClassifyFNFResponseDataFractures) SetBoxes(v []*int) *ClassifyFNFResponseDataFractures {
	s.Boxes = v
	return s
}

type ClassifyFNFResponseDataFracturesTag struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty" require:"true"`
}

func (s ClassifyFNFResponseDataFracturesTag) String() string {
	return tea.Prettify(s)
}

func (s ClassifyFNFResponseDataFracturesTag) GoString() string {
	return s.String()
}

func (s *ClassifyFNFResponseDataFracturesTag) SetLabel(v string) *ClassifyFNFResponseDataFracturesTag {
	s.Label = &v
	return s
}

type ClassifyFNFAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	DataFormat     *string   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgId          *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	OrgName        *string   `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	TracerId       *string   `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ClassifyFNFAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyFNFAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ClassifyFNFAdvanceRequest) SetImageUrlObject(v io.Reader) *ClassifyFNFAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *ClassifyFNFAdvanceRequest) SetDataFormat(v string) *ClassifyFNFAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *ClassifyFNFAdvanceRequest) SetOrgId(v string) *ClassifyFNFAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *ClassifyFNFAdvanceRequest) SetOrgName(v string) *ClassifyFNFAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *ClassifyFNFAdvanceRequest) SetTracerId(v string) *ClassifyFNFAdvanceRequest {
	s.TracerId = &v
	return s
}

type RunCTRegistrationRequest struct {
	ReferenceList  []*RunCTRegistrationRequestReferenceList `json:"ReferenceList,omitempty" xml:"ReferenceList,omitempty" require:"true" type:"Repeated"`
	DataFormat     *string                                  `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgName        *string                                  `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	OrgId          *string                                  `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	DataSourceType *string                                  `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty" require:"true"`
	FloatingList   []*RunCTRegistrationRequestFloatingList  `json:"FloatingList,omitempty" xml:"FloatingList,omitempty" require:"true" type:"Repeated"`
}

func (s RunCTRegistrationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationRequest) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationRequest) SetReferenceList(v []*RunCTRegistrationRequestReferenceList) *RunCTRegistrationRequest {
	s.ReferenceList = v
	return s
}

func (s *RunCTRegistrationRequest) SetDataFormat(v string) *RunCTRegistrationRequest {
	s.DataFormat = &v
	return s
}

func (s *RunCTRegistrationRequest) SetOrgName(v string) *RunCTRegistrationRequest {
	s.OrgName = &v
	return s
}

func (s *RunCTRegistrationRequest) SetOrgId(v string) *RunCTRegistrationRequest {
	s.OrgId = &v
	return s
}

func (s *RunCTRegistrationRequest) SetDataSourceType(v string) *RunCTRegistrationRequest {
	s.DataSourceType = &v
	return s
}

func (s *RunCTRegistrationRequest) SetFloatingList(v []*RunCTRegistrationRequestFloatingList) *RunCTRegistrationRequest {
	s.FloatingList = v
	return s
}

type RunCTRegistrationRequestReferenceList struct {
	ReferenceURL *string `json:"ReferenceURL,omitempty" xml:"ReferenceURL,omitempty" require:"true"`
}

func (s RunCTRegistrationRequestReferenceList) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationRequestReferenceList) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationRequestReferenceList) SetReferenceURL(v string) *RunCTRegistrationRequestReferenceList {
	s.ReferenceURL = &v
	return s
}

type RunCTRegistrationRequestFloatingList struct {
	FloatingURL *string `json:"FloatingURL,omitempty" xml:"FloatingURL,omitempty" require:"true"`
}

func (s RunCTRegistrationRequestFloatingList) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationRequestFloatingList) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationRequestFloatingList) SetFloatingURL(v string) *RunCTRegistrationRequestFloatingList {
	s.FloatingURL = &v
	return s
}

type RunCTRegistrationResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *RunCTRegistrationResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s RunCTRegistrationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationResponse) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationResponse) SetRequestId(v string) *RunCTRegistrationResponse {
	s.RequestId = &v
	return s
}

func (s *RunCTRegistrationResponse) SetData(v *RunCTRegistrationResponseData) *RunCTRegistrationResponse {
	s.Data = v
	return s
}

type RunCTRegistrationResponseData struct {
	DUrl *string `json:"DUrl,omitempty" xml:"DUrl,omitempty" require:"true"`
	NUrl *string `json:"NUrl,omitempty" xml:"NUrl,omitempty" require:"true"`
}

func (s RunCTRegistrationResponseData) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationResponseData) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationResponseData) SetDUrl(v string) *RunCTRegistrationResponseData {
	s.DUrl = &v
	return s
}

func (s *RunCTRegistrationResponseData) SetNUrl(v string) *RunCTRegistrationResponseData {
	s.NUrl = &v
	return s
}

type DetectHipKeypointXRayRequest struct {
	ImageUrl   *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" require:"true"`
	DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	OrgName    *string `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	TracerId   *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DetectHipKeypointXRayRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectHipKeypointXRayRequest) GoString() string {
	return s.String()
}

func (s *DetectHipKeypointXRayRequest) SetImageUrl(v string) *DetectHipKeypointXRayRequest {
	s.ImageUrl = &v
	return s
}

func (s *DetectHipKeypointXRayRequest) SetDataFormat(v string) *DetectHipKeypointXRayRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectHipKeypointXRayRequest) SetOrgId(v string) *DetectHipKeypointXRayRequest {
	s.OrgId = &v
	return s
}

func (s *DetectHipKeypointXRayRequest) SetOrgName(v string) *DetectHipKeypointXRayRequest {
	s.OrgName = &v
	return s
}

func (s *DetectHipKeypointXRayRequest) SetTracerId(v string) *DetectHipKeypointXRayRequest {
	s.TracerId = &v
	return s
}

type DetectHipKeypointXRayResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *DetectHipKeypointXRayResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s DetectHipKeypointXRayResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectHipKeypointXRayResponse) GoString() string {
	return s.String()
}

func (s *DetectHipKeypointXRayResponse) SetRequestId(v string) *DetectHipKeypointXRayResponse {
	s.RequestId = &v
	return s
}

func (s *DetectHipKeypointXRayResponse) SetData(v *DetectHipKeypointXRayResponseData) *DetectHipKeypointXRayResponse {
	s.Data = v
	return s
}

type DetectHipKeypointXRayResponseData struct {
	ImageUrl  *string                                       `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" require:"true"`
	OrgId     *string                                       `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	OrgName   *string                                       `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	KeyPoints []*DetectHipKeypointXRayResponseDataKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" require:"true" type:"Repeated"`
}

func (s DetectHipKeypointXRayResponseData) String() string {
	return tea.Prettify(s)
}

func (s DetectHipKeypointXRayResponseData) GoString() string {
	return s.String()
}

func (s *DetectHipKeypointXRayResponseData) SetImageUrl(v string) *DetectHipKeypointXRayResponseData {
	s.ImageUrl = &v
	return s
}

func (s *DetectHipKeypointXRayResponseData) SetOrgId(v string) *DetectHipKeypointXRayResponseData {
	s.OrgId = &v
	return s
}

func (s *DetectHipKeypointXRayResponseData) SetOrgName(v string) *DetectHipKeypointXRayResponseData {
	s.OrgName = &v
	return s
}

func (s *DetectHipKeypointXRayResponseData) SetKeyPoints(v []*DetectHipKeypointXRayResponseDataKeyPoints) *DetectHipKeypointXRayResponseData {
	s.KeyPoints = v
	return s
}

type DetectHipKeypointXRayResponseDataKeyPoints struct {
	Value       *float32                                       `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
	Tag         *DetectHipKeypointXRayResponseDataKeyPointsTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Struct"`
	Coordinates []*int                                         `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" require:"true" type:"Repeated"`
}

func (s DetectHipKeypointXRayResponseDataKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s DetectHipKeypointXRayResponseDataKeyPoints) GoString() string {
	return s.String()
}

func (s *DetectHipKeypointXRayResponseDataKeyPoints) SetValue(v float32) *DetectHipKeypointXRayResponseDataKeyPoints {
	s.Value = &v
	return s
}

func (s *DetectHipKeypointXRayResponseDataKeyPoints) SetTag(v *DetectHipKeypointXRayResponseDataKeyPointsTag) *DetectHipKeypointXRayResponseDataKeyPoints {
	s.Tag = v
	return s
}

func (s *DetectHipKeypointXRayResponseDataKeyPoints) SetCoordinates(v []*int) *DetectHipKeypointXRayResponseDataKeyPoints {
	s.Coordinates = v
	return s
}

type DetectHipKeypointXRayResponseDataKeyPointsTag struct {
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty" require:"true"`
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty" require:"true"`
}

func (s DetectHipKeypointXRayResponseDataKeyPointsTag) String() string {
	return tea.Prettify(s)
}

func (s DetectHipKeypointXRayResponseDataKeyPointsTag) GoString() string {
	return s.String()
}

func (s *DetectHipKeypointXRayResponseDataKeyPointsTag) SetDirection(v string) *DetectHipKeypointXRayResponseDataKeyPointsTag {
	s.Direction = &v
	return s
}

func (s *DetectHipKeypointXRayResponseDataKeyPointsTag) SetLabel(v string) *DetectHipKeypointXRayResponseDataKeyPointsTag {
	s.Label = &v
	return s
}

type DetectHipKeypointXRayAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	DataFormat     *string   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgId          *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	OrgName        *string   `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	TracerId       *string   `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DetectHipKeypointXRayAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectHipKeypointXRayAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectHipKeypointXRayAdvanceRequest) SetImageUrlObject(v io.Reader) *DetectHipKeypointXRayAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *DetectHipKeypointXRayAdvanceRequest) SetDataFormat(v string) *DetectHipKeypointXRayAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectHipKeypointXRayAdvanceRequest) SetOrgId(v string) *DetectHipKeypointXRayAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *DetectHipKeypointXRayAdvanceRequest) SetOrgName(v string) *DetectHipKeypointXRayAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *DetectHipKeypointXRayAdvanceRequest) SetTracerId(v string) *DetectHipKeypointXRayAdvanceRequest {
	s.TracerId = &v
	return s
}

type CalcCACSRequest struct {
	URLList        []*CalcCACSRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" require:"true" type:"Repeated"`
	DataFormat     *string                   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgName        *string                   `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	OrgId          *string                   `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
	DataSourceType *string                   `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty" require:"true"`
}

func (s CalcCACSRequest) String() string {
	return tea.Prettify(s)
}

func (s CalcCACSRequest) GoString() string {
	return s.String()
}

func (s *CalcCACSRequest) SetURLList(v []*CalcCACSRequestURLList) *CalcCACSRequest {
	s.URLList = v
	return s
}

func (s *CalcCACSRequest) SetDataFormat(v string) *CalcCACSRequest {
	s.DataFormat = &v
	return s
}

func (s *CalcCACSRequest) SetOrgName(v string) *CalcCACSRequest {
	s.OrgName = &v
	return s
}

func (s *CalcCACSRequest) SetOrgId(v string) *CalcCACSRequest {
	s.OrgId = &v
	return s
}

func (s *CalcCACSRequest) SetDataSourceType(v string) *CalcCACSRequest {
	s.DataSourceType = &v
	return s
}

type CalcCACSRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty" require:"true"`
}

func (s CalcCACSRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s CalcCACSRequestURLList) GoString() string {
	return s.String()
}

func (s *CalcCACSRequestURLList) SetURL(v string) *CalcCACSRequestURLList {
	s.URL = &v
	return s
}

type CalcCACSResponse struct {
	RequestId *string               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *CalcCACSResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s CalcCACSResponse) String() string {
	return tea.Prettify(s)
}

func (s CalcCACSResponse) GoString() string {
	return s.String()
}

func (s *CalcCACSResponse) SetRequestId(v string) *CalcCACSResponse {
	s.RequestId = &v
	return s
}

func (s *CalcCACSResponse) SetData(v *CalcCACSResponseData) *CalcCACSResponse {
	s.Data = v
	return s
}

type CalcCACSResponseData struct {
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty" require:"true"`
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty" require:"true"`
}

func (s CalcCACSResponseData) String() string {
	return tea.Prettify(s)
}

func (s CalcCACSResponseData) GoString() string {
	return s.String()
}

func (s *CalcCACSResponseData) SetScore(v string) *CalcCACSResponseData {
	s.Score = &v
	return s
}

func (s *CalcCACSResponseData) SetResultUrl(v string) *CalcCACSResponseData {
	s.ResultUrl = &v
	return s
}

type DetectKneeXRayRequest struct {
	Url        *string `json:"Url,omitempty" xml:"Url,omitempty" require:"true"`
	DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgName    *string `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
}

func (s DetectKneeXRayRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeXRayRequest) GoString() string {
	return s.String()
}

func (s *DetectKneeXRayRequest) SetUrl(v string) *DetectKneeXRayRequest {
	s.Url = &v
	return s
}

func (s *DetectKneeXRayRequest) SetDataFormat(v string) *DetectKneeXRayRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectKneeXRayRequest) SetOrgName(v string) *DetectKneeXRayRequest {
	s.OrgName = &v
	return s
}

func (s *DetectKneeXRayRequest) SetOrgId(v string) *DetectKneeXRayRequest {
	s.OrgId = &v
	return s
}

type DetectKneeXRayResponse struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *DetectKneeXRayResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s DetectKneeXRayResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeXRayResponse) GoString() string {
	return s.String()
}

func (s *DetectKneeXRayResponse) SetRequestId(v string) *DetectKneeXRayResponse {
	s.RequestId = &v
	return s
}

func (s *DetectKneeXRayResponse) SetData(v *DetectKneeXRayResponseData) *DetectKneeXRayResponse {
	s.Data = v
	return s
}

type DetectKneeXRayResponseData struct {
	KLDetections []*DetectKneeXRayResponseDataKLDetections `json:"KLDetections,omitempty" xml:"KLDetections,omitempty" require:"true" type:"Repeated"`
}

func (s DetectKneeXRayResponseData) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeXRayResponseData) GoString() string {
	return s.String()
}

func (s *DetectKneeXRayResponseData) SetKLDetections(v []*DetectKneeXRayResponseDataKLDetections) *DetectKneeXRayResponseData {
	s.KLDetections = v
	return s
}

type DetectKneeXRayResponseDataKLDetections struct {
	Detections []*float32 `json:"Detections,omitempty" xml:"Detections,omitempty" require:"true" type:"Repeated"`
}

func (s DetectKneeXRayResponseDataKLDetections) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeXRayResponseDataKLDetections) GoString() string {
	return s.String()
}

func (s *DetectKneeXRayResponseDataKLDetections) SetDetections(v []*float32) *DetectKneeXRayResponseDataKLDetections {
	s.Detections = v
	return s
}

type DetectKneeXRayAdvanceRequest struct {
	UrlObject  io.Reader `json:"UrlObject,omitempty" xml:"UrlObject,omitempty" require:"true"`
	DataFormat *string   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgName    *string   `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	OrgId      *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
}

func (s DetectKneeXRayAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeXRayAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectKneeXRayAdvanceRequest) SetUrlObject(v io.Reader) *DetectKneeXRayAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *DetectKneeXRayAdvanceRequest) SetDataFormat(v string) *DetectKneeXRayAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectKneeXRayAdvanceRequest) SetOrgName(v string) *DetectKneeXRayAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *DetectKneeXRayAdvanceRequest) SetOrgId(v string) *DetectKneeXRayAdvanceRequest {
	s.OrgId = &v
	return s
}

type DetectSpineMRIRequest struct {
	URLList    []*DetectSpineMRIRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" require:"true" type:"Repeated"`
	DataFormat *string                         `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgName    *string                         `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	OrgId      *string                         `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
}

func (s DetectSpineMRIRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIRequest) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIRequest) SetURLList(v []*DetectSpineMRIRequestURLList) *DetectSpineMRIRequest {
	s.URLList = v
	return s
}

func (s *DetectSpineMRIRequest) SetDataFormat(v string) *DetectSpineMRIRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectSpineMRIRequest) SetOrgName(v string) *DetectSpineMRIRequest {
	s.OrgName = &v
	return s
}

func (s *DetectSpineMRIRequest) SetOrgId(v string) *DetectSpineMRIRequest {
	s.OrgId = &v
	return s
}

type DetectSpineMRIRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty" require:"true"`
}

func (s DetectSpineMRIRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIRequestURLList) SetURL(v string) *DetectSpineMRIRequestURLList {
	s.URL = &v
	return s
}

type DetectSpineMRIResponse struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *DetectSpineMRIResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s DetectSpineMRIResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIResponse) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIResponse) SetRequestId(v string) *DetectSpineMRIResponse {
	s.RequestId = &v
	return s
}

func (s *DetectSpineMRIResponse) SetData(v *DetectSpineMRIResponseData) *DetectSpineMRIResponse {
	s.Data = v
	return s
}

type DetectSpineMRIResponseData struct {
	Discs     []*DetectSpineMRIResponseDataDiscs     `json:"Discs,omitempty" xml:"Discs,omitempty" require:"true" type:"Repeated"`
	Vertebras []*DetectSpineMRIResponseDataVertebras `json:"Vertebras,omitempty" xml:"Vertebras,omitempty" require:"true" type:"Repeated"`
}

func (s DetectSpineMRIResponseData) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIResponseData) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIResponseData) SetDiscs(v []*DetectSpineMRIResponseDataDiscs) *DetectSpineMRIResponseData {
	s.Discs = v
	return s
}

func (s *DetectSpineMRIResponseData) SetVertebras(v []*DetectSpineMRIResponseDataVertebras) *DetectSpineMRIResponseData {
	s.Vertebras = v
	return s
}

type DetectSpineMRIResponseDataDiscs struct {
	Disease        *string    `json:"Disease,omitempty" xml:"Disease,omitempty" require:"true"`
	Identification *string    `json:"Identification,omitempty" xml:"Identification,omitempty" require:"true"`
	Location       []*float32 `json:"Location,omitempty" xml:"Location,omitempty" require:"true" type:"Repeated"`
}

func (s DetectSpineMRIResponseDataDiscs) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIResponseDataDiscs) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIResponseDataDiscs) SetDisease(v string) *DetectSpineMRIResponseDataDiscs {
	s.Disease = &v
	return s
}

func (s *DetectSpineMRIResponseDataDiscs) SetIdentification(v string) *DetectSpineMRIResponseDataDiscs {
	s.Identification = &v
	return s
}

func (s *DetectSpineMRIResponseDataDiscs) SetLocation(v []*float32) *DetectSpineMRIResponseDataDiscs {
	s.Location = v
	return s
}

type DetectSpineMRIResponseDataVertebras struct {
	Disease        *string    `json:"Disease,omitempty" xml:"Disease,omitempty" require:"true"`
	Identification *string    `json:"Identification,omitempty" xml:"Identification,omitempty" require:"true"`
	Location       []*float32 `json:"Location,omitempty" xml:"Location,omitempty" require:"true" type:"Repeated"`
}

func (s DetectSpineMRIResponseDataVertebras) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIResponseDataVertebras) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIResponseDataVertebras) SetDisease(v string) *DetectSpineMRIResponseDataVertebras {
	s.Disease = &v
	return s
}

func (s *DetectSpineMRIResponseDataVertebras) SetIdentification(v string) *DetectSpineMRIResponseDataVertebras {
	s.Identification = &v
	return s
}

func (s *DetectSpineMRIResponseDataVertebras) SetLocation(v []*float32) *DetectSpineMRIResponseDataVertebras {
	s.Location = v
	return s
}

type TranslateMedRequest struct {
	FromLanguage *string `json:"FromLanguage,omitempty" xml:"FromLanguage,omitempty" require:"true"`
	ToLanguage   *string `json:"ToLanguage,omitempty" xml:"ToLanguage,omitempty" require:"true"`
	Text         *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
}

func (s TranslateMedRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateMedRequest) GoString() string {
	return s.String()
}

func (s *TranslateMedRequest) SetFromLanguage(v string) *TranslateMedRequest {
	s.FromLanguage = &v
	return s
}

func (s *TranslateMedRequest) SetToLanguage(v string) *TranslateMedRequest {
	s.ToLanguage = &v
	return s
}

func (s *TranslateMedRequest) SetText(v string) *TranslateMedRequest {
	s.Text = &v
	return s
}

type TranslateMedResponse struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *TranslateMedResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s TranslateMedResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateMedResponse) GoString() string {
	return s.String()
}

func (s *TranslateMedResponse) SetRequestId(v string) *TranslateMedResponse {
	s.RequestId = &v
	return s
}

func (s *TranslateMedResponse) SetData(v *TranslateMedResponseData) *TranslateMedResponse {
	s.Data = v
	return s
}

type TranslateMedResponseData struct {
	Text  *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	Words *int64  `json:"Words,omitempty" xml:"Words,omitempty" require:"true"`
}

func (s TranslateMedResponseData) String() string {
	return tea.Prettify(s)
}

func (s TranslateMedResponseData) GoString() string {
	return s.String()
}

func (s *TranslateMedResponseData) SetText(v string) *TranslateMedResponseData {
	s.Text = &v
	return s
}

func (s *TranslateMedResponseData) SetWords(v int64) *TranslateMedResponseData {
	s.Words = &v
	return s
}

type DetectLungNoduleRequest struct {
	URLList    []*DetectLungNoduleRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" require:"true" type:"Repeated"`
	DataFormat *string                           `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgName    *string                           `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	OrgId      *string                           `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
}

func (s DetectLungNoduleRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleRequest) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleRequest) SetURLList(v []*DetectLungNoduleRequestURLList) *DetectLungNoduleRequest {
	s.URLList = v
	return s
}

func (s *DetectLungNoduleRequest) SetDataFormat(v string) *DetectLungNoduleRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectLungNoduleRequest) SetOrgName(v string) *DetectLungNoduleRequest {
	s.OrgName = &v
	return s
}

func (s *DetectLungNoduleRequest) SetOrgId(v string) *DetectLungNoduleRequest {
	s.OrgId = &v
	return s
}

type DetectLungNoduleRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty" require:"true"`
}

func (s DetectLungNoduleRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleRequestURLList) SetURL(v string) *DetectLungNoduleRequestURLList {
	s.URL = &v
	return s
}

type DetectLungNoduleResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *DetectLungNoduleResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s DetectLungNoduleResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleResponse) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleResponse) SetRequestId(v string) *DetectLungNoduleResponse {
	s.RequestId = &v
	return s
}

func (s *DetectLungNoduleResponse) SetData(v *DetectLungNoduleResponseData) *DetectLungNoduleResponse {
	s.Data = v
	return s
}

type DetectLungNoduleResponseData struct {
	Series []*DetectLungNoduleResponseDataSeries `json:"Series,omitempty" xml:"Series,omitempty" require:"true" type:"Repeated"`
}

func (s DetectLungNoduleResponseData) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleResponseData) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleResponseData) SetSeries(v []*DetectLungNoduleResponseDataSeries) *DetectLungNoduleResponseData {
	s.Series = v
	return s
}

type DetectLungNoduleResponseDataSeries struct {
	SeriesInstanceUid *string                                       `json:"SeriesInstanceUid,omitempty" xml:"SeriesInstanceUid,omitempty" require:"true"`
	Report            *string                                       `json:"Report,omitempty" xml:"Report,omitempty" require:"true"`
	Elements          []*DetectLungNoduleResponseDataSeriesElements `json:"Elements,omitempty" xml:"Elements,omitempty" require:"true" type:"Repeated"`
	Origin            []*float32                                    `json:"Origin,omitempty" xml:"Origin,omitempty" require:"true" type:"Repeated"`
	Spacing           []*float32                                    `json:"Spacing,omitempty" xml:"Spacing,omitempty" require:"true" type:"Repeated"`
}

func (s DetectLungNoduleResponseDataSeries) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleResponseDataSeries) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleResponseDataSeries) SetSeriesInstanceUid(v string) *DetectLungNoduleResponseDataSeries {
	s.SeriesInstanceUid = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeries) SetReport(v string) *DetectLungNoduleResponseDataSeries {
	s.Report = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeries) SetElements(v []*DetectLungNoduleResponseDataSeriesElements) *DetectLungNoduleResponseDataSeries {
	s.Elements = v
	return s
}

func (s *DetectLungNoduleResponseDataSeries) SetOrigin(v []*float32) *DetectLungNoduleResponseDataSeries {
	s.Origin = v
	return s
}

func (s *DetectLungNoduleResponseDataSeries) SetSpacing(v []*float32) *DetectLungNoduleResponseDataSeries {
	s.Spacing = v
	return s
}

type DetectLungNoduleResponseDataSeriesElements struct {
	Category       *string  `json:"Category,omitempty" xml:"Category,omitempty" require:"true"`
	Confidence     *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty" require:"true"`
	Diameter       *float32 `json:"Diameter,omitempty" xml:"Diameter,omitempty" require:"true"`
	Lobe           *string  `json:"Lobe,omitempty" xml:"Lobe,omitempty" require:"true"`
	Lung           *string  `json:"Lung,omitempty" xml:"Lung,omitempty" require:"true"`
	X              *float32 `json:"X,omitempty" xml:"X,omitempty" require:"true"`
	Z              *float32 `json:"Z,omitempty" xml:"Z,omitempty" require:"true"`
	Y              *float32 `json:"Y,omitempty" xml:"Y,omitempty" require:"true"`
	ImageX         *float32 `json:"ImageX,omitempty" xml:"ImageX,omitempty" require:"true"`
	ImageY         *float32 `json:"ImageY,omitempty" xml:"ImageY,omitempty" require:"true"`
	ImageZ         *float32 `json:"ImageZ,omitempty" xml:"ImageZ,omitempty" require:"true"`
	SOPInstanceUID *string  `json:"SOPInstanceUID,omitempty" xml:"SOPInstanceUID,omitempty" require:"true"`
	Volume         *float32 `json:"Volume,omitempty" xml:"Volume,omitempty" require:"true"`
	MeanValue      *float32 `json:"MeanValue,omitempty" xml:"MeanValue,omitempty" require:"true"`
}

func (s DetectLungNoduleResponseDataSeriesElements) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleResponseDataSeriesElements) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetCategory(v string) *DetectLungNoduleResponseDataSeriesElements {
	s.Category = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetConfidence(v float32) *DetectLungNoduleResponseDataSeriesElements {
	s.Confidence = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetDiameter(v float32) *DetectLungNoduleResponseDataSeriesElements {
	s.Diameter = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetLobe(v string) *DetectLungNoduleResponseDataSeriesElements {
	s.Lobe = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetLung(v string) *DetectLungNoduleResponseDataSeriesElements {
	s.Lung = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetX(v float32) *DetectLungNoduleResponseDataSeriesElements {
	s.X = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetZ(v float32) *DetectLungNoduleResponseDataSeriesElements {
	s.Z = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetY(v float32) *DetectLungNoduleResponseDataSeriesElements {
	s.Y = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetImageX(v float32) *DetectLungNoduleResponseDataSeriesElements {
	s.ImageX = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetImageY(v float32) *DetectLungNoduleResponseDataSeriesElements {
	s.ImageY = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetImageZ(v float32) *DetectLungNoduleResponseDataSeriesElements {
	s.ImageZ = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetSOPInstanceUID(v string) *DetectLungNoduleResponseDataSeriesElements {
	s.SOPInstanceUID = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetVolume(v float32) *DetectLungNoduleResponseDataSeriesElements {
	s.Volume = &v
	return s
}

func (s *DetectLungNoduleResponseDataSeriesElements) SetMeanValue(v float32) *DetectLungNoduleResponseDataSeriesElements {
	s.MeanValue = &v
	return s
}

type DetectCovid19CadRequest struct {
	URLList    []*DetectCovid19CadRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" require:"true" type:"Repeated"`
	DataFormat *string                           `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" require:"true"`
	OrgName    *string                           `json:"OrgName,omitempty" xml:"OrgName,omitempty" require:"true"`
	OrgId      *string                           `json:"OrgId,omitempty" xml:"OrgId,omitempty" require:"true"`
}

func (s DetectCovid19CadRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectCovid19CadRequest) GoString() string {
	return s.String()
}

func (s *DetectCovid19CadRequest) SetURLList(v []*DetectCovid19CadRequestURLList) *DetectCovid19CadRequest {
	s.URLList = v
	return s
}

func (s *DetectCovid19CadRequest) SetDataFormat(v string) *DetectCovid19CadRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectCovid19CadRequest) SetOrgName(v string) *DetectCovid19CadRequest {
	s.OrgName = &v
	return s
}

func (s *DetectCovid19CadRequest) SetOrgId(v string) *DetectCovid19CadRequest {
	s.OrgId = &v
	return s
}

type DetectCovid19CadRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty" require:"true"`
}

func (s DetectCovid19CadRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectCovid19CadRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectCovid19CadRequestURLList) SetURL(v string) *DetectCovid19CadRequestURLList {
	s.URL = &v
	return s
}

type DetectCovid19CadResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *DetectCovid19CadResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s DetectCovid19CadResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectCovid19CadResponse) GoString() string {
	return s.String()
}

func (s *DetectCovid19CadResponse) SetRequestId(v string) *DetectCovid19CadResponse {
	s.RequestId = &v
	return s
}

func (s *DetectCovid19CadResponse) SetData(v *DetectCovid19CadResponseData) *DetectCovid19CadResponse {
	s.Data = v
	return s
}

type DetectCovid19CadResponseData struct {
	NewProbability    *string `json:"NewProbability,omitempty" xml:"NewProbability,omitempty" require:"true"`
	NormalProbability *string `json:"NormalProbability,omitempty" xml:"NormalProbability,omitempty" require:"true"`
	OtherProbability  *string `json:"OtherProbability,omitempty" xml:"OtherProbability,omitempty" require:"true"`
	LesionRatio       *string `json:"LesionRatio,omitempty" xml:"LesionRatio,omitempty" require:"true"`
	Mask              *string `json:"Mask,omitempty" xml:"Mask,omitempty" require:"true"`
}

func (s DetectCovid19CadResponseData) String() string {
	return tea.Prettify(s)
}

func (s DetectCovid19CadResponseData) GoString() string {
	return s.String()
}

func (s *DetectCovid19CadResponseData) SetNewProbability(v string) *DetectCovid19CadResponseData {
	s.NewProbability = &v
	return s
}

func (s *DetectCovid19CadResponseData) SetNormalProbability(v string) *DetectCovid19CadResponseData {
	s.NormalProbability = &v
	return s
}

func (s *DetectCovid19CadResponseData) SetOtherProbability(v string) *DetectCovid19CadResponseData {
	s.OtherProbability = &v
	return s
}

func (s *DetectCovid19CadResponseData) SetLesionRatio(v string) *DetectCovid19CadResponseData {
	s.LesionRatio = &v
	return s
}

func (s *DetectCovid19CadResponseData) SetMask(v string) *DetectCovid19CadResponseData {
	s.Mask = &v
	return s
}

type GetAsyncJobResultRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty" require:"true"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetAsyncJobResultResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetAsyncJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponse) SetRequestId(v string) *GetAsyncJobResultResponse {
	s.RequestId = &v
	return s
}

func (s *GetAsyncJobResultResponse) SetData(v *GetAsyncJobResultResponseData) *GetAsyncJobResultResponse {
	s.Data = v
	return s
}

type GetAsyncJobResultResponseData struct {
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty" require:"true"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty" require:"true"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty" require:"true"`
}

func (s GetAsyncJobResultResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseData) SetJobId(v string) *GetAsyncJobResultResponseData {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultResponseData) SetStatus(v string) *GetAsyncJobResultResponseData {
	s.Status = &v
	return s
}

func (s *GetAsyncJobResultResponseData) SetResult(v string) *GetAsyncJobResultResponseData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseData) SetErrorCode(v string) *GetAsyncJobResultResponseData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseData) SetErrorMessage(v string) *GetAsyncJobResultResponseData {
	s.ErrorMessage = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("imageprocess"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) DetectRibFracture(request *DetectRibFractureRequest, runtime *util.RuntimeOptions) (_result *DetectRibFractureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DetectRibFractureResponse{}
	_body, _err := client.DoRequest(tea.String("DetectRibFracture"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectRibFractureSimply(request *DetectRibFractureRequest) (_result *DetectRibFractureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectRibFractureResponse{}
	_body, _err := client.DetectRibFracture(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScreenChestCT(request *ScreenChestCTRequest, runtime *util.RuntimeOptions) (_result *ScreenChestCTResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ScreenChestCTResponse{}
	_body, _err := client.DoRequest(tea.String("ScreenChestCT"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScreenChestCTSimply(request *ScreenChestCTRequest) (_result *ScreenChestCTResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScreenChestCTResponse{}
	_body, _err := client.ScreenChestCT(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectSkinDisease(request *DetectSkinDiseaseRequest, runtime *util.RuntimeOptions) (_result *DetectSkinDiseaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DetectSkinDiseaseResponse{}
	_body, _err := client.DoRequest(tea.String("DetectSkinDisease"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectSkinDiseaseSimply(request *DetectSkinDiseaseRequest) (_result *DetectSkinDiseaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectSkinDiseaseResponse{}
	_body, _err := client.DetectSkinDisease(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectSkinDiseaseAdvance(request *DetectSkinDiseaseAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectSkinDiseaseResponse, _err error) {
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
		Product:  tea.String("imageprocess"),
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
	rpcutil.Convert(runtime, ossRuntime)
	detectSkinDiseaseReq := &DetectSkinDiseaseRequest{}
	rpcutil.Convert(request, detectSkinDiseaseReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.UrlObject,
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
	detectSkinDiseaseReq.Url = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectSkinDiseaseResp, _err := client.DetectSkinDisease(detectSkinDiseaseReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectSkinDiseaseResp
	return _result, _err
}

func (client *Client) RunMedQA(request *RunMedQARequest, runtime *util.RuntimeOptions) (_result *RunMedQAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RunMedQAResponse{}
	_body, _err := client.DoRequest(tea.String("RunMedQA"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunMedQASimply(request *RunMedQARequest) (_result *RunMedQAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunMedQAResponse{}
	_body, _err := client.RunMedQA(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectKneeKeypointXRay(request *DetectKneeKeypointXRayRequest, runtime *util.RuntimeOptions) (_result *DetectKneeKeypointXRayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DetectKneeKeypointXRayResponse{}
	_body, _err := client.DoRequest(tea.String("DetectKneeKeypointXRay"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectKneeKeypointXRaySimply(request *DetectKneeKeypointXRayRequest) (_result *DetectKneeKeypointXRayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectKneeKeypointXRayResponse{}
	_body, _err := client.DetectKneeKeypointXRay(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectKneeKeypointXRayAdvance(request *DetectKneeKeypointXRayAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectKneeKeypointXRayResponse, _err error) {
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
		Product:  tea.String("imageprocess"),
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
	rpcutil.Convert(runtime, ossRuntime)
	detectKneeKeypointXRayReq := &DetectKneeKeypointXRayRequest{}
	rpcutil.Convert(request, detectKneeKeypointXRayReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	detectKneeKeypointXRayReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectKneeKeypointXRayResp, _err := client.DetectKneeKeypointXRay(detectKneeKeypointXRayReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectKneeKeypointXRayResp
	return _result, _err
}

func (client *Client) ClassifyFNF(request *ClassifyFNFRequest, runtime *util.RuntimeOptions) (_result *ClassifyFNFResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ClassifyFNFResponse{}
	_body, _err := client.DoRequest(tea.String("ClassifyFNF"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClassifyFNFSimply(request *ClassifyFNFRequest) (_result *ClassifyFNFResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClassifyFNFResponse{}
	_body, _err := client.ClassifyFNF(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClassifyFNFAdvance(request *ClassifyFNFAdvanceRequest, runtime *util.RuntimeOptions) (_result *ClassifyFNFResponse, _err error) {
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
		Product:  tea.String("imageprocess"),
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
	rpcutil.Convert(runtime, ossRuntime)
	classifyFNFReq := &ClassifyFNFRequest{}
	rpcutil.Convert(request, classifyFNFReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	classifyFNFReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	classifyFNFResp, _err := client.ClassifyFNF(classifyFNFReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = classifyFNFResp
	return _result, _err
}

func (client *Client) RunCTRegistration(request *RunCTRegistrationRequest, runtime *util.RuntimeOptions) (_result *RunCTRegistrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RunCTRegistrationResponse{}
	_body, _err := client.DoRequest(tea.String("RunCTRegistration"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunCTRegistrationSimply(request *RunCTRegistrationRequest) (_result *RunCTRegistrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCTRegistrationResponse{}
	_body, _err := client.RunCTRegistration(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectHipKeypointXRay(request *DetectHipKeypointXRayRequest, runtime *util.RuntimeOptions) (_result *DetectHipKeypointXRayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DetectHipKeypointXRayResponse{}
	_body, _err := client.DoRequest(tea.String("DetectHipKeypointXRay"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectHipKeypointXRaySimply(request *DetectHipKeypointXRayRequest) (_result *DetectHipKeypointXRayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectHipKeypointXRayResponse{}
	_body, _err := client.DetectHipKeypointXRay(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectHipKeypointXRayAdvance(request *DetectHipKeypointXRayAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectHipKeypointXRayResponse, _err error) {
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
		Product:  tea.String("imageprocess"),
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
	rpcutil.Convert(runtime, ossRuntime)
	detectHipKeypointXRayReq := &DetectHipKeypointXRayRequest{}
	rpcutil.Convert(request, detectHipKeypointXRayReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	detectHipKeypointXRayReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectHipKeypointXRayResp, _err := client.DetectHipKeypointXRay(detectHipKeypointXRayReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectHipKeypointXRayResp
	return _result, _err
}

func (client *Client) CalcCACS(request *CalcCACSRequest, runtime *util.RuntimeOptions) (_result *CalcCACSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CalcCACSResponse{}
	_body, _err := client.DoRequest(tea.String("CalcCACS"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CalcCACSSimply(request *CalcCACSRequest) (_result *CalcCACSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CalcCACSResponse{}
	_body, _err := client.CalcCACS(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectKneeXRay(request *DetectKneeXRayRequest, runtime *util.RuntimeOptions) (_result *DetectKneeXRayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DetectKneeXRayResponse{}
	_body, _err := client.DoRequest(tea.String("DetectKneeXRay"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectKneeXRaySimply(request *DetectKneeXRayRequest) (_result *DetectKneeXRayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectKneeXRayResponse{}
	_body, _err := client.DetectKneeXRay(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectKneeXRayAdvance(request *DetectKneeXRayAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectKneeXRayResponse, _err error) {
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
		Product:  tea.String("imageprocess"),
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
	rpcutil.Convert(runtime, ossRuntime)
	detectKneeXRayReq := &DetectKneeXRayRequest{}
	rpcutil.Convert(request, detectKneeXRayReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.UrlObject,
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
	detectKneeXRayReq.Url = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectKneeXRayResp, _err := client.DetectKneeXRay(detectKneeXRayReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectKneeXRayResp
	return _result, _err
}

func (client *Client) DetectSpineMRI(request *DetectSpineMRIRequest, runtime *util.RuntimeOptions) (_result *DetectSpineMRIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DetectSpineMRIResponse{}
	_body, _err := client.DoRequest(tea.String("DetectSpineMRI"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectSpineMRISimply(request *DetectSpineMRIRequest) (_result *DetectSpineMRIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectSpineMRIResponse{}
	_body, _err := client.DetectSpineMRI(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TranslateMed(request *TranslateMedRequest, runtime *util.RuntimeOptions) (_result *TranslateMedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &TranslateMedResponse{}
	_body, _err := client.DoRequest(tea.String("TranslateMed"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TranslateMedSimply(request *TranslateMedRequest) (_result *TranslateMedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TranslateMedResponse{}
	_body, _err := client.TranslateMed(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectLungNodule(request *DetectLungNoduleRequest, runtime *util.RuntimeOptions) (_result *DetectLungNoduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DetectLungNoduleResponse{}
	_body, _err := client.DoRequest(tea.String("DetectLungNodule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectLungNoduleSimply(request *DetectLungNoduleRequest) (_result *DetectLungNoduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectLungNoduleResponse{}
	_body, _err := client.DetectLungNodule(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectCovid19Cad(request *DetectCovid19CadRequest, runtime *util.RuntimeOptions) (_result *DetectCovid19CadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DetectCovid19CadResponse{}
	_body, _err := client.DoRequest(tea.String("DetectCovid19Cad"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectCovid19CadSimply(request *DetectCovid19CadRequest) (_result *DetectCovid19CadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectCovid19CadResponse{}
	_body, _err := client.DetectCovid19Cad(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncJobResult(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.DoRequest(tea.String("GetAsyncJobResult"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncJobResultSimply(request *GetAsyncJobResultRequest) (_result *GetAsyncJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.GetAsyncJobResult(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
