// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type FaceRecognitionCompareRequest struct {
	IdCardImageData   *string `json:"idCardImageData,omitempty" xml:"idCardImageData,omitempty"`
	IdCardImageType   *string `json:"idCardImageType,omitempty" xml:"idCardImageType,omitempty"`
	IdCardImageUrl    *string `json:"idCardImageUrl,omitempty" xml:"idCardImageUrl,omitempty"`
	PortraitImageData *string `json:"portraitImageData,omitempty" xml:"portraitImageData,omitempty"`
	PortraitImageType *string `json:"portraitImageType,omitempty" xml:"portraitImageType,omitempty"`
	PortraitImageUrl  *string `json:"portraitImageUrl,omitempty" xml:"portraitImageUrl,omitempty"`
	QualityControl    *string `json:"qualityControl,omitempty" xml:"qualityControl,omitempty"`
}

func (s FaceRecognitionCompareRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceRecognitionCompareRequest) GoString() string {
	return s.String()
}

func (s *FaceRecognitionCompareRequest) SetIdCardImageData(v string) *FaceRecognitionCompareRequest {
	s.IdCardImageData = &v
	return s
}

func (s *FaceRecognitionCompareRequest) SetIdCardImageType(v string) *FaceRecognitionCompareRequest {
	s.IdCardImageType = &v
	return s
}

func (s *FaceRecognitionCompareRequest) SetIdCardImageUrl(v string) *FaceRecognitionCompareRequest {
	s.IdCardImageUrl = &v
	return s
}

func (s *FaceRecognitionCompareRequest) SetPortraitImageData(v string) *FaceRecognitionCompareRequest {
	s.PortraitImageData = &v
	return s
}

func (s *FaceRecognitionCompareRequest) SetPortraitImageType(v string) *FaceRecognitionCompareRequest {
	s.PortraitImageType = &v
	return s
}

func (s *FaceRecognitionCompareRequest) SetPortraitImageUrl(v string) *FaceRecognitionCompareRequest {
	s.PortraitImageUrl = &v
	return s
}

func (s *FaceRecognitionCompareRequest) SetQualityControl(v string) *FaceRecognitionCompareRequest {
	s.QualityControl = &v
	return s
}

type FaceRecognitionCompareResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *FaceRecognitionCompareResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpCode  *int64                                  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	Ok        *bool                                   `json:"Ok,omitempty" xml:"Ok,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s FaceRecognitionCompareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceRecognitionCompareResponseBody) GoString() string {
	return s.String()
}

func (s *FaceRecognitionCompareResponseBody) SetCode(v string) *FaceRecognitionCompareResponseBody {
	s.Code = &v
	return s
}

func (s *FaceRecognitionCompareResponseBody) SetData(v *FaceRecognitionCompareResponseBodyData) *FaceRecognitionCompareResponseBody {
	s.Data = v
	return s
}

func (s *FaceRecognitionCompareResponseBody) SetHttpCode(v int64) *FaceRecognitionCompareResponseBody {
	s.HttpCode = &v
	return s
}

func (s *FaceRecognitionCompareResponseBody) SetMessage(v string) *FaceRecognitionCompareResponseBody {
	s.Message = &v
	return s
}

func (s *FaceRecognitionCompareResponseBody) SetOk(v bool) *FaceRecognitionCompareResponseBody {
	s.Ok = &v
	return s
}

func (s *FaceRecognitionCompareResponseBody) SetRequestId(v string) *FaceRecognitionCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceRecognitionCompareResponseBody) SetStatus(v string) *FaceRecognitionCompareResponseBody {
	s.Status = &v
	return s
}

type FaceRecognitionCompareResponseBodyData struct {
	Match *string  `json:"Match,omitempty" xml:"Match,omitempty"`
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s FaceRecognitionCompareResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FaceRecognitionCompareResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceRecognitionCompareResponseBodyData) SetMatch(v string) *FaceRecognitionCompareResponseBodyData {
	s.Match = &v
	return s
}

func (s *FaceRecognitionCompareResponseBodyData) SetScore(v float64) *FaceRecognitionCompareResponseBodyData {
	s.Score = &v
	return s
}

type FaceRecognitionCompareResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FaceRecognitionCompareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FaceRecognitionCompareResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceRecognitionCompareResponse) GoString() string {
	return s.String()
}

func (s *FaceRecognitionCompareResponse) SetHeaders(v map[string]*string) *FaceRecognitionCompareResponse {
	s.Headers = v
	return s
}

func (s *FaceRecognitionCompareResponse) SetStatusCode(v int32) *FaceRecognitionCompareResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceRecognitionCompareResponse) SetBody(v *FaceRecognitionCompareResponseBody) *FaceRecognitionCompareResponse {
	s.Body = v
	return s
}

type FaceToFaceCompareRequest struct {
	PortraitImageData1 *string `json:"portraitImageData1,omitempty" xml:"portraitImageData1,omitempty"`
	PortraitImageData2 *string `json:"portraitImageData2,omitempty" xml:"portraitImageData2,omitempty"`
	PortraitImageType1 *string `json:"portraitImageType1,omitempty" xml:"portraitImageType1,omitempty"`
	PortraitImageType2 *string `json:"portraitImageType2,omitempty" xml:"portraitImageType2,omitempty"`
	PortraitImageUrl1  *string `json:"portraitImageUrl1,omitempty" xml:"portraitImageUrl1,omitempty"`
	PortraitImageUrl2  *string `json:"portraitImageUrl2,omitempty" xml:"portraitImageUrl2,omitempty"`
	QualityControl     *string `json:"qualityControl,omitempty" xml:"qualityControl,omitempty"`
}

func (s FaceToFaceCompareRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceToFaceCompareRequest) GoString() string {
	return s.String()
}

func (s *FaceToFaceCompareRequest) SetPortraitImageData1(v string) *FaceToFaceCompareRequest {
	s.PortraitImageData1 = &v
	return s
}

func (s *FaceToFaceCompareRequest) SetPortraitImageData2(v string) *FaceToFaceCompareRequest {
	s.PortraitImageData2 = &v
	return s
}

func (s *FaceToFaceCompareRequest) SetPortraitImageType1(v string) *FaceToFaceCompareRequest {
	s.PortraitImageType1 = &v
	return s
}

func (s *FaceToFaceCompareRequest) SetPortraitImageType2(v string) *FaceToFaceCompareRequest {
	s.PortraitImageType2 = &v
	return s
}

func (s *FaceToFaceCompareRequest) SetPortraitImageUrl1(v string) *FaceToFaceCompareRequest {
	s.PortraitImageUrl1 = &v
	return s
}

func (s *FaceToFaceCompareRequest) SetPortraitImageUrl2(v string) *FaceToFaceCompareRequest {
	s.PortraitImageUrl2 = &v
	return s
}

func (s *FaceToFaceCompareRequest) SetQualityControl(v string) *FaceToFaceCompareRequest {
	s.QualityControl = &v
	return s
}

type FaceToFaceCompareResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *FaceToFaceCompareResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpCode  *int64                             `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Ok        *bool                              `json:"Ok,omitempty" xml:"Ok,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                            `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s FaceToFaceCompareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceToFaceCompareResponseBody) GoString() string {
	return s.String()
}

func (s *FaceToFaceCompareResponseBody) SetCode(v string) *FaceToFaceCompareResponseBody {
	s.Code = &v
	return s
}

func (s *FaceToFaceCompareResponseBody) SetData(v *FaceToFaceCompareResponseBodyData) *FaceToFaceCompareResponseBody {
	s.Data = v
	return s
}

func (s *FaceToFaceCompareResponseBody) SetHttpCode(v int64) *FaceToFaceCompareResponseBody {
	s.HttpCode = &v
	return s
}

func (s *FaceToFaceCompareResponseBody) SetMessage(v string) *FaceToFaceCompareResponseBody {
	s.Message = &v
	return s
}

func (s *FaceToFaceCompareResponseBody) SetOk(v bool) *FaceToFaceCompareResponseBody {
	s.Ok = &v
	return s
}

func (s *FaceToFaceCompareResponseBody) SetRequestId(v string) *FaceToFaceCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceToFaceCompareResponseBody) SetStatus(v string) *FaceToFaceCompareResponseBody {
	s.Status = &v
	return s
}

type FaceToFaceCompareResponseBodyData struct {
	Match *string  `json:"Match,omitempty" xml:"Match,omitempty"`
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s FaceToFaceCompareResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FaceToFaceCompareResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceToFaceCompareResponseBodyData) SetMatch(v string) *FaceToFaceCompareResponseBodyData {
	s.Match = &v
	return s
}

func (s *FaceToFaceCompareResponseBodyData) SetScore(v float64) *FaceToFaceCompareResponseBodyData {
	s.Score = &v
	return s
}

type FaceToFaceCompareResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FaceToFaceCompareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FaceToFaceCompareResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceToFaceCompareResponse) GoString() string {
	return s.String()
}

func (s *FaceToFaceCompareResponse) SetHeaders(v map[string]*string) *FaceToFaceCompareResponse {
	s.Headers = v
	return s
}

func (s *FaceToFaceCompareResponse) SetStatusCode(v int32) *FaceToFaceCompareResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceToFaceCompareResponse) SetBody(v *FaceToFaceCompareResponseBody) *FaceToFaceCompareResponse {
	s.Body = v
	return s
}

type IdDetectOcrCompareFacesRequest struct {
	OCR               *bool   `json:"OCR,omitempty" xml:"OCR,omitempty"`
	Unrepeat          *bool   `json:"Unrepeat,omitempty" xml:"Unrepeat,omitempty"`
	CardDetect        *bool   `json:"cardDetect,omitempty" xml:"cardDetect,omitempty"`
	CountryCode       *string `json:"countryCode,omitempty" xml:"countryCode,omitempty"`
	DocumentType      *string `json:"documentType,omitempty" xml:"documentType,omitempty"`
	FaceCompare       *bool   `json:"faceCompare,omitempty" xml:"faceCompare,omitempty"`
	ImageAType        *string `json:"imageAType,omitempty" xml:"imageAType,omitempty"`
	ImageBType        *string `json:"imageBType,omitempty" xml:"imageBType,omitempty"`
	ImageDataA        *string `json:"imageDataA,omitempty" xml:"imageDataA,omitempty"`
	ImageDataB        *string `json:"imageDataB,omitempty" xml:"imageDataB,omitempty"`
	ImageUrlA         *string `json:"imageUrlA,omitempty" xml:"imageUrlA,omitempty"`
	ImageUrlB         *string `json:"imageUrlB,omitempty" xml:"imageUrlB,omitempty"`
	PortraitImageData *string `json:"portraitImageData,omitempty" xml:"portraitImageData,omitempty"`
	PortraitImageUrl  *string `json:"portraitImageUrl,omitempty" xml:"portraitImageUrl,omitempty"`
	QualityControl    *string `json:"qualityControl,omitempty" xml:"qualityControl,omitempty"`
}

func (s IdDetectOcrCompareFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesRequest) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesRequest) SetOCR(v bool) *IdDetectOcrCompareFacesRequest {
	s.OCR = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetUnrepeat(v bool) *IdDetectOcrCompareFacesRequest {
	s.Unrepeat = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetCardDetect(v bool) *IdDetectOcrCompareFacesRequest {
	s.CardDetect = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetCountryCode(v string) *IdDetectOcrCompareFacesRequest {
	s.CountryCode = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetDocumentType(v string) *IdDetectOcrCompareFacesRequest {
	s.DocumentType = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetFaceCompare(v bool) *IdDetectOcrCompareFacesRequest {
	s.FaceCompare = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetImageAType(v string) *IdDetectOcrCompareFacesRequest {
	s.ImageAType = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetImageBType(v string) *IdDetectOcrCompareFacesRequest {
	s.ImageBType = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetImageDataA(v string) *IdDetectOcrCompareFacesRequest {
	s.ImageDataA = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetImageDataB(v string) *IdDetectOcrCompareFacesRequest {
	s.ImageDataB = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetImageUrlA(v string) *IdDetectOcrCompareFacesRequest {
	s.ImageUrlA = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetImageUrlB(v string) *IdDetectOcrCompareFacesRequest {
	s.ImageUrlB = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetPortraitImageData(v string) *IdDetectOcrCompareFacesRequest {
	s.PortraitImageData = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetPortraitImageUrl(v string) *IdDetectOcrCompareFacesRequest {
	s.PortraitImageUrl = &v
	return s
}

func (s *IdDetectOcrCompareFacesRequest) SetQualityControl(v string) *IdDetectOcrCompareFacesRequest {
	s.QualityControl = &v
	return s
}

type IdDetectOcrCompareFacesResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *IdDetectOcrCompareFacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpCode  *int64                                   `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	Ok        *bool                                    `json:"Ok,omitempty" xml:"Ok,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                                  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s IdDetectOcrCompareFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBody) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBody) SetCode(v string) *IdDetectOcrCompareFacesResponseBody {
	s.Code = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBody) SetData(v *IdDetectOcrCompareFacesResponseBodyData) *IdDetectOcrCompareFacesResponseBody {
	s.Data = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBody) SetHttpCode(v int64) *IdDetectOcrCompareFacesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBody) SetMessage(v string) *IdDetectOcrCompareFacesResponseBody {
	s.Message = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBody) SetOk(v bool) *IdDetectOcrCompareFacesResponseBody {
	s.Ok = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBody) SetRequestId(v string) *IdDetectOcrCompareFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBody) SetStatus(v string) *IdDetectOcrCompareFacesResponseBody {
	s.Status = &v
	return s
}

type IdDetectOcrCompareFacesResponseBodyData struct {
	CompareResult  *IdDetectOcrCompareFacesResponseBodyDataCompareResult  `json:"CompareResult,omitempty" xml:"CompareResult,omitempty" type:"Struct"`
	IdDetectResult *IdDetectOcrCompareFacesResponseBodyDataIdDetectResult `json:"IdDetectResult,omitempty" xml:"IdDetectResult,omitempty" type:"Struct"`
	LivenessResult *IdDetectOcrCompareFacesResponseBodyDataLivenessResult `json:"LivenessResult,omitempty" xml:"LivenessResult,omitempty" type:"Struct"`
	OcrResult      *IdDetectOcrCompareFacesResponseBodyDataOcrResult      `json:"OcrResult,omitempty" xml:"OcrResult,omitempty" type:"Struct"`
	UnrepeatResult *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResult `json:"UnrepeatResult,omitempty" xml:"UnrepeatResult,omitempty" type:"Struct"`
}

func (s IdDetectOcrCompareFacesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyData) SetCompareResult(v *IdDetectOcrCompareFacesResponseBodyDataCompareResult) *IdDetectOcrCompareFacesResponseBodyData {
	s.CompareResult = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyData) SetIdDetectResult(v *IdDetectOcrCompareFacesResponseBodyDataIdDetectResult) *IdDetectOcrCompareFacesResponseBodyData {
	s.IdDetectResult = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyData) SetLivenessResult(v *IdDetectOcrCompareFacesResponseBodyDataLivenessResult) *IdDetectOcrCompareFacesResponseBodyData {
	s.LivenessResult = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyData) SetOcrResult(v *IdDetectOcrCompareFacesResponseBodyDataOcrResult) *IdDetectOcrCompareFacesResponseBodyData {
	s.OcrResult = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyData) SetUnrepeatResult(v *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResult) *IdDetectOcrCompareFacesResponseBodyData {
	s.UnrepeatResult = v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataCompareResult struct {
	Code              *string                                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	CompareResultData *IdDetectOcrCompareFacesResponseBodyDataCompareResultCompareResultData `json:"CompareResultData,omitempty" xml:"CompareResultData,omitempty" type:"Struct"`
	Message           *string                                                                `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataCompareResult) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataCompareResult) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataCompareResult) SetCode(v string) *IdDetectOcrCompareFacesResponseBodyDataCompareResult {
	s.Code = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataCompareResult) SetCompareResultData(v *IdDetectOcrCompareFacesResponseBodyDataCompareResultCompareResultData) *IdDetectOcrCompareFacesResponseBodyDataCompareResult {
	s.CompareResultData = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataCompareResult) SetMessage(v string) *IdDetectOcrCompareFacesResponseBodyDataCompareResult {
	s.Message = &v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataCompareResultCompareResultData struct {
	Match *string   `json:"Match,omitempty" xml:"Match,omitempty"`
	Risks []*string `json:"Risks,omitempty" xml:"Risks,omitempty" type:"Repeated"`
	Score *float64  `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataCompareResultCompareResultData) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataCompareResultCompareResultData) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataCompareResultCompareResultData) SetMatch(v string) *IdDetectOcrCompareFacesResponseBodyDataCompareResultCompareResultData {
	s.Match = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataCompareResultCompareResultData) SetRisks(v []*string) *IdDetectOcrCompareFacesResponseBodyDataCompareResultCompareResultData {
	s.Risks = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataCompareResultCompareResultData) SetScore(v float64) *IdDetectOcrCompareFacesResponseBodyDataCompareResultCompareResultData {
	s.Score = &v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataIdDetectResult struct {
	Code               *string                                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	IdDetectResultData *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData `json:"IdDetectResultData,omitempty" xml:"IdDetectResultData,omitempty" type:"Struct"`
	Message            *string                                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataIdDetectResult) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataIdDetectResult) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResult) SetCode(v string) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResult {
	s.Code = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResult) SetIdDetectResultData(v *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResult {
	s.IdDetectResultData = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResult) SetMessage(v string) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResult {
	s.Message = &v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData struct {
	AntiSpoofingResult *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataAntiSpoofingResult `json:"AntiSpoofingResult,omitempty" xml:"AntiSpoofingResult,omitempty" type:"Struct"`
	CountryCode        *string                                                                                    `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	DetectionResult    *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataDetectionResult    `json:"DetectionResult,omitempty" xml:"DetectionResult,omitempty" type:"Struct"`
	DocumentType       *string                                                                                    `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	Passed             *bool                                                                                      `json:"Passed,omitempty" xml:"Passed,omitempty"`
	Warning            []*string                                                                                  `json:"Warning,omitempty" xml:"Warning,omitempty" type:"Repeated"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData) SetAntiSpoofingResult(v *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataAntiSpoofingResult) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData {
	s.AntiSpoofingResult = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData) SetCountryCode(v string) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData {
	s.CountryCode = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData) SetDetectionResult(v *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataDetectionResult) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData {
	s.DetectionResult = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData) SetDocumentType(v string) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData {
	s.DocumentType = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData) SetPassed(v bool) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData {
	s.Passed = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData) SetWarning(v []*string) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultData {
	s.Warning = v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataAntiSpoofingResult struct {
	Passed *bool     `json:"Passed,omitempty" xml:"Passed,omitempty"`
	Risks  []*string `json:"Risks,omitempty" xml:"Risks,omitempty" type:"Repeated"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataAntiSpoofingResult) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataAntiSpoofingResult) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataAntiSpoofingResult) SetPassed(v bool) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataAntiSpoofingResult {
	s.Passed = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataAntiSpoofingResult) SetRisks(v []*string) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataAntiSpoofingResult {
	s.Risks = v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataDetectionResult struct {
	CardRectangle     []*int64 `json:"CardRectangle,omitempty" xml:"CardRectangle,omitempty" type:"Repeated"`
	Height            *int64   `json:"Height,omitempty" xml:"Height,omitempty"`
	PortraitRectangle []*int64 `json:"PortraitRectangle,omitempty" xml:"PortraitRectangle,omitempty" type:"Repeated"`
	Width             *int64   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataDetectionResult) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataDetectionResult) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataDetectionResult) SetCardRectangle(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataDetectionResult {
	s.CardRectangle = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataDetectionResult) SetHeight(v int64) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataDetectionResult {
	s.Height = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataDetectionResult) SetPortraitRectangle(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataDetectionResult {
	s.PortraitRectangle = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataDetectionResult) SetWidth(v int64) *IdDetectOcrCompareFacesResponseBodyDataIdDetectResultIdDetectResultDataDetectionResult {
	s.Width = &v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataLivenessResult struct {
	Code               *string                                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	LivenessResultData *IdDetectOcrCompareFacesResponseBodyDataLivenessResultLivenessResultData `json:"LivenessResultData,omitempty" xml:"LivenessResultData,omitempty" type:"Struct"`
	Message            *string                                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataLivenessResult) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataLivenessResult) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataLivenessResult) SetCode(v string) *IdDetectOcrCompareFacesResponseBodyDataLivenessResult {
	s.Code = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataLivenessResult) SetLivenessResultData(v *IdDetectOcrCompareFacesResponseBodyDataLivenessResultLivenessResultData) *IdDetectOcrCompareFacesResponseBodyDataLivenessResult {
	s.LivenessResultData = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataLivenessResult) SetMessage(v string) *IdDetectOcrCompareFacesResponseBodyDataLivenessResult {
	s.Message = &v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataLivenessResultLivenessResultData struct {
	Liveness *string   `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	Risks    []*string `json:"Risks,omitempty" xml:"Risks,omitempty" type:"Repeated"`
	Score    *float64  `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataLivenessResultLivenessResultData) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataLivenessResultLivenessResultData) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataLivenessResultLivenessResultData) SetLiveness(v string) *IdDetectOcrCompareFacesResponseBodyDataLivenessResultLivenessResultData {
	s.Liveness = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataLivenessResultLivenessResultData) SetRisks(v []*string) *IdDetectOcrCompareFacesResponseBodyDataLivenessResultLivenessResultData {
	s.Risks = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataLivenessResultLivenessResultData) SetScore(v float64) *IdDetectOcrCompareFacesResponseBodyDataLivenessResultLivenessResultData {
	s.Score = &v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataOcrResult struct {
	Code          *string                                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string                                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	OcrResultData *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData `json:"OcrResultData,omitempty" xml:"OcrResultData,omitempty" type:"Struct"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataOcrResult) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataOcrResult) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResult) SetCode(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResult {
	s.Code = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResult) SetMessage(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResult {
	s.Message = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResult) SetOcrResultData(v *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) *IdDetectOcrCompareFacesResponseBodyDataOcrResult {
	s.OcrResultData = v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData struct {
	Address                    *string  `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressConfidence          *float64 `json:"AddressConfidence,omitempty" xml:"AddressConfidence,omitempty"`
	AddressPosition            []*int64 `json:"AddressPosition,omitempty" xml:"AddressPosition,omitempty" type:"Repeated"`
	DateOfBirth                *string  `json:"DateOfBirth,omitempty" xml:"DateOfBirth,omitempty"`
	DateOfBirthConfidence      *float64 `json:"DateOfBirthConfidence,omitempty" xml:"DateOfBirthConfidence,omitempty"`
	DateOfBirthPosition        []*int64 `json:"DateOfBirthPosition,omitempty" xml:"DateOfBirthPosition,omitempty" type:"Repeated"`
	DocumentNumber             *string  `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	DocumentNumber2            *string  `json:"DocumentNumber2,omitempty" xml:"DocumentNumber2,omitempty"`
	DocumentNumber2Confidence  *float64 `json:"DocumentNumber2Confidence,omitempty" xml:"DocumentNumber2Confidence,omitempty"`
	DocumentNumber2Position    []*int64 `json:"DocumentNumber2Position,omitempty" xml:"DocumentNumber2Position,omitempty" type:"Repeated"`
	DocumentNumberConfidence   *float64 `json:"DocumentNumberConfidence,omitempty" xml:"DocumentNumberConfidence,omitempty"`
	DocumentNumberPosition     []*int64 `json:"DocumentNumberPosition,omitempty" xml:"DocumentNumberPosition,omitempty" type:"Repeated"`
	ExpirationDate             *string  `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	ExpirationDateConfidence   *float64 `json:"ExpirationDateConfidence,omitempty" xml:"ExpirationDateConfidence,omitempty"`
	ExpirationDatePosition     []*int64 `json:"ExpirationDatePosition,omitempty" xml:"ExpirationDatePosition,omitempty" type:"Repeated"`
	FirstName                  *string  `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	FirstNameConfidence        *float64 `json:"FirstNameConfidence,omitempty" xml:"FirstNameConfidence,omitempty"`
	FirstNamePosition          []*int64 `json:"FirstNamePosition,omitempty" xml:"FirstNamePosition,omitempty" type:"Repeated"`
	FullName                   *string  `json:"FullName,omitempty" xml:"FullName,omitempty"`
	FullNameConfidence         *float64 `json:"FullNameConfidence,omitempty" xml:"FullNameConfidence,omitempty"`
	FullNamePosition           []*int64 `json:"FullNamePosition,omitempty" xml:"FullNamePosition,omitempty" type:"Repeated"`
	IssuedDate                 *string  `json:"IssuedDate,omitempty" xml:"IssuedDate,omitempty"`
	IssuedDateConfidence       *float64 `json:"IssuedDateConfidence,omitempty" xml:"IssuedDateConfidence,omitempty"`
	IssuedDatePosition         []*int64 `json:"IssuedDatePosition,omitempty" xml:"IssuedDatePosition,omitempty" type:"Repeated"`
	IssuingAuthority           *string  `json:"IssuingAuthority,omitempty" xml:"IssuingAuthority,omitempty"`
	IssuingAuthorityConfidence *float64 `json:"IssuingAuthorityConfidence,omitempty" xml:"IssuingAuthorityConfidence,omitempty"`
	IssuingAuthorityPosition   []*int64 `json:"IssuingAuthorityPosition,omitempty" xml:"IssuingAuthorityPosition,omitempty" type:"Repeated"`
	LastName                   *string  `json:"LastName,omitempty" xml:"LastName,omitempty"`
	LastNameConfidence         *float64 `json:"LastNameConfidence,omitempty" xml:"LastNameConfidence,omitempty"`
	LastNamePosition           []*int64 `json:"LastNamePosition,omitempty" xml:"LastNamePosition,omitempty" type:"Repeated"`
	NationalityCode            *string  `json:"NationalityCode,omitempty" xml:"NationalityCode,omitempty"`
	NationalityCodeConfidence  *float64 `json:"NationalityCodeConfidence,omitempty" xml:"NationalityCodeConfidence,omitempty"`
	NationalityCodePosition    []*int64 `json:"NationalityCodePosition,omitempty" xml:"NationalityCodePosition,omitempty" type:"Repeated"`
	NormalizedDateOfBirth      *string  `json:"NormalizedDateOfBirth,omitempty" xml:"NormalizedDateOfBirth,omitempty"`
	NormalizedExpirationDate   *string  `json:"NormalizedExpirationDate,omitempty" xml:"NormalizedExpirationDate,omitempty"`
	NormalizedIssuedDate       *string  `json:"NormalizedIssuedDate,omitempty" xml:"NormalizedIssuedDate,omitempty"`
	NormalizedNationalityCode  *string  `json:"NormalizedNationalityCode,omitempty" xml:"NormalizedNationalityCode,omitempty"`
	NormalizedSex              *string  `json:"NormalizedSex,omitempty" xml:"NormalizedSex,omitempty"`
	Sex                        *string  `json:"Sex,omitempty" xml:"Sex,omitempty"`
	SexConfidence              *float64 `json:"SexConfidence,omitempty" xml:"SexConfidence,omitempty"`
	SexPosition                []*int64 `json:"SexPosition,omitempty" xml:"SexPosition,omitempty" type:"Repeated"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetAddress(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.Address = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetAddressConfidence(v float64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.AddressConfidence = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetAddressPosition(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.AddressPosition = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetDateOfBirth(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.DateOfBirth = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetDateOfBirthConfidence(v float64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.DateOfBirthConfidence = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetDateOfBirthPosition(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.DateOfBirthPosition = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetDocumentNumber(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.DocumentNumber = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetDocumentNumber2(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.DocumentNumber2 = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetDocumentNumber2Confidence(v float64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.DocumentNumber2Confidence = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetDocumentNumber2Position(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.DocumentNumber2Position = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetDocumentNumberConfidence(v float64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.DocumentNumberConfidence = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetDocumentNumberPosition(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.DocumentNumberPosition = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetExpirationDate(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.ExpirationDate = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetExpirationDateConfidence(v float64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.ExpirationDateConfidence = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetExpirationDatePosition(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.ExpirationDatePosition = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetFirstName(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.FirstName = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetFirstNameConfidence(v float64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.FirstNameConfidence = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetFirstNamePosition(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.FirstNamePosition = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetFullName(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.FullName = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetFullNameConfidence(v float64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.FullNameConfidence = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetFullNamePosition(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.FullNamePosition = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetIssuedDate(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.IssuedDate = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetIssuedDateConfidence(v float64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.IssuedDateConfidence = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetIssuedDatePosition(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.IssuedDatePosition = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetIssuingAuthority(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.IssuingAuthority = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetIssuingAuthorityConfidence(v float64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.IssuingAuthorityConfidence = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetIssuingAuthorityPosition(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.IssuingAuthorityPosition = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetLastName(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.LastName = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetLastNameConfidence(v float64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.LastNameConfidence = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetLastNamePosition(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.LastNamePosition = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetNationalityCode(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.NationalityCode = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetNationalityCodeConfidence(v float64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.NationalityCodeConfidence = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetNationalityCodePosition(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.NationalityCodePosition = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetNormalizedDateOfBirth(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.NormalizedDateOfBirth = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetNormalizedExpirationDate(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.NormalizedExpirationDate = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetNormalizedIssuedDate(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.NormalizedIssuedDate = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetNormalizedNationalityCode(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.NormalizedNationalityCode = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetNormalizedSex(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.NormalizedSex = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetSex(v string) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.Sex = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetSexConfidence(v float64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.SexConfidence = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData) SetSexPosition(v []*int64) *IdDetectOcrCompareFacesResponseBodyDataOcrResultOcrResultData {
	s.SexPosition = v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataUnrepeatResult struct {
	CardImageResult     *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResult     `json:"CardImageResult,omitempty" xml:"CardImageResult,omitempty" type:"Struct"`
	PortraitImageResult *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResult `json:"PortraitImageResult,omitempty" xml:"PortraitImageResult,omitempty" type:"Struct"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataUnrepeatResult) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataUnrepeatResult) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResult) SetCardImageResult(v *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResult) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResult {
	s.CardImageResult = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResult) SetPortraitImageResult(v *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResult) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResult {
	s.PortraitImageResult = v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResult struct {
	CardImageResultData *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResultCardImageResultData `json:"CardImageResultData,omitempty" xml:"CardImageResultData,omitempty" type:"Struct"`
	Code                *string                                                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message             *string                                                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResult) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResult) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResult) SetCardImageResultData(v *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResultCardImageResultData) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResult {
	s.CardImageResultData = v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResult) SetCode(v string) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResult {
	s.Code = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResult) SetMessage(v string) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResult {
	s.Message = &v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResultCardImageResultData struct {
	Dbname    *string  `json:"Dbname,omitempty" xml:"Dbname,omitempty"`
	EntityId  *string  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	ExtraData *string  `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	Score     *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResultCardImageResultData) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResultCardImageResultData) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResultCardImageResultData) SetDbname(v string) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResultCardImageResultData {
	s.Dbname = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResultCardImageResultData) SetEntityId(v string) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResultCardImageResultData {
	s.EntityId = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResultCardImageResultData) SetExtraData(v string) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResultCardImageResultData {
	s.ExtraData = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResultCardImageResultData) SetScore(v float64) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultCardImageResultCardImageResultData {
	s.Score = &v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResult struct {
	Code                    *string                                                                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                 *string                                                                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PortraitImageResultData *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResultPortraitImageResultData `json:"PortraitImageResultData,omitempty" xml:"PortraitImageResultData,omitempty" type:"Struct"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResult) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResult) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResult) SetCode(v string) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResult {
	s.Code = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResult) SetMessage(v string) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResult {
	s.Message = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResult) SetPortraitImageResultData(v *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResultPortraitImageResultData) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResult {
	s.PortraitImageResultData = v
	return s
}

type IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResultPortraitImageResultData struct {
	Dbname    *string  `json:"Dbname,omitempty" xml:"Dbname,omitempty"`
	EntityId  *string  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	ExtraData *string  `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	Score     *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResultPortraitImageResultData) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResultPortraitImageResultData) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResultPortraitImageResultData) SetDbname(v string) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResultPortraitImageResultData {
	s.Dbname = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResultPortraitImageResultData) SetEntityId(v string) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResultPortraitImageResultData {
	s.EntityId = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResultPortraitImageResultData) SetExtraData(v string) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResultPortraitImageResultData {
	s.ExtraData = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResultPortraitImageResultData) SetScore(v float64) *IdDetectOcrCompareFacesResponseBodyDataUnrepeatResultPortraitImageResultPortraitImageResultData {
	s.Score = &v
	return s
}

type IdDetectOcrCompareFacesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IdDetectOcrCompareFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IdDetectOcrCompareFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s IdDetectOcrCompareFacesResponse) GoString() string {
	return s.String()
}

func (s *IdDetectOcrCompareFacesResponse) SetHeaders(v map[string]*string) *IdDetectOcrCompareFacesResponse {
	s.Headers = v
	return s
}

func (s *IdDetectOcrCompareFacesResponse) SetStatusCode(v int32) *IdDetectOcrCompareFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *IdDetectOcrCompareFacesResponse) SetBody(v *IdDetectOcrCompareFacesResponseBody) *IdDetectOcrCompareFacesResponse {
	s.Body = v
	return s
}

type IdDetectionAndScanDocumentRequest struct {
	CountryCode  *string `json:"countryCode,omitempty" xml:"countryCode,omitempty"`
	DocumentType *string `json:"documentType,omitempty" xml:"documentType,omitempty"`
	ImageAType   *string `json:"imageAType,omitempty" xml:"imageAType,omitempty"`
	ImageBType   *string `json:"imageBType,omitempty" xml:"imageBType,omitempty"`
	ImageDataA   *string `json:"imageDataA,omitempty" xml:"imageDataA,omitempty"`
	ImageDataB   *string `json:"imageDataB,omitempty" xml:"imageDataB,omitempty"`
	ImageUrlA    *string `json:"imageUrlA,omitempty" xml:"imageUrlA,omitempty"`
	ImageUrlB    *string `json:"imageUrlB,omitempty" xml:"imageUrlB,omitempty"`
}

func (s IdDetectionAndScanDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s IdDetectionAndScanDocumentRequest) GoString() string {
	return s.String()
}

func (s *IdDetectionAndScanDocumentRequest) SetCountryCode(v string) *IdDetectionAndScanDocumentRequest {
	s.CountryCode = &v
	return s
}

func (s *IdDetectionAndScanDocumentRequest) SetDocumentType(v string) *IdDetectionAndScanDocumentRequest {
	s.DocumentType = &v
	return s
}

func (s *IdDetectionAndScanDocumentRequest) SetImageAType(v string) *IdDetectionAndScanDocumentRequest {
	s.ImageAType = &v
	return s
}

func (s *IdDetectionAndScanDocumentRequest) SetImageBType(v string) *IdDetectionAndScanDocumentRequest {
	s.ImageBType = &v
	return s
}

func (s *IdDetectionAndScanDocumentRequest) SetImageDataA(v string) *IdDetectionAndScanDocumentRequest {
	s.ImageDataA = &v
	return s
}

func (s *IdDetectionAndScanDocumentRequest) SetImageDataB(v string) *IdDetectionAndScanDocumentRequest {
	s.ImageDataB = &v
	return s
}

func (s *IdDetectionAndScanDocumentRequest) SetImageUrlA(v string) *IdDetectionAndScanDocumentRequest {
	s.ImageUrlA = &v
	return s
}

func (s *IdDetectionAndScanDocumentRequest) SetImageUrlB(v string) *IdDetectionAndScanDocumentRequest {
	s.ImageUrlB = &v
	return s
}

type IdDetectionAndScanDocumentResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *IdDetectionAndScanDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpCode  *int64                                      `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Ok        *bool                                       `json:"Ok,omitempty" xml:"Ok,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s IdDetectionAndScanDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IdDetectionAndScanDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *IdDetectionAndScanDocumentResponseBody) SetCode(v string) *IdDetectionAndScanDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBody) SetData(v *IdDetectionAndScanDocumentResponseBodyData) *IdDetectionAndScanDocumentResponseBody {
	s.Data = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBody) SetHttpCode(v int64) *IdDetectionAndScanDocumentResponseBody {
	s.HttpCode = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBody) SetMessage(v string) *IdDetectionAndScanDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBody) SetOk(v bool) *IdDetectionAndScanDocumentResponseBody {
	s.Ok = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBody) SetRequestId(v string) *IdDetectionAndScanDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBody) SetStatus(v string) *IdDetectionAndScanDocumentResponseBody {
	s.Status = &v
	return s
}

type IdDetectionAndScanDocumentResponseBodyData struct {
	IdDetectResult *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult `json:"IdDetectResult,omitempty" xml:"IdDetectResult,omitempty" type:"Struct"`
	OcrResult      *IdDetectionAndScanDocumentResponseBodyDataOcrResult      `json:"OcrResult,omitempty" xml:"OcrResult,omitempty" type:"Struct"`
}

func (s IdDetectionAndScanDocumentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s IdDetectionAndScanDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *IdDetectionAndScanDocumentResponseBodyData) SetIdDetectResult(v *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult) *IdDetectionAndScanDocumentResponseBodyData {
	s.IdDetectResult = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyData) SetOcrResult(v *IdDetectionAndScanDocumentResponseBodyDataOcrResult) *IdDetectionAndScanDocumentResponseBodyData {
	s.OcrResult = v
	return s
}

type IdDetectionAndScanDocumentResponseBodyDataIdDetectResult struct {
	AntiSpoofingResult *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultAntiSpoofingResult `json:"AntiSpoofingResult,omitempty" xml:"AntiSpoofingResult,omitempty" type:"Struct"`
	CountryCode        *string                                                                     `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	DetectionResult    *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultDetectionResult    `json:"DetectionResult,omitempty" xml:"DetectionResult,omitempty" type:"Struct"`
	DocumentType       *string                                                                     `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	Passed             *bool                                                                       `json:"Passed,omitempty" xml:"Passed,omitempty"`
	Warning            []*string                                                                   `json:"Warning,omitempty" xml:"Warning,omitempty" type:"Repeated"`
}

func (s IdDetectionAndScanDocumentResponseBodyDataIdDetectResult) String() string {
	return tea.Prettify(s)
}

func (s IdDetectionAndScanDocumentResponseBodyDataIdDetectResult) GoString() string {
	return s.String()
}

func (s *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult) SetAntiSpoofingResult(v *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultAntiSpoofingResult) *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult {
	s.AntiSpoofingResult = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult) SetCountryCode(v string) *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult {
	s.CountryCode = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult) SetDetectionResult(v *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultDetectionResult) *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult {
	s.DetectionResult = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult) SetDocumentType(v string) *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult {
	s.DocumentType = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult) SetPassed(v bool) *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult {
	s.Passed = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult) SetWarning(v []*string) *IdDetectionAndScanDocumentResponseBodyDataIdDetectResult {
	s.Warning = v
	return s
}

type IdDetectionAndScanDocumentResponseBodyDataIdDetectResultAntiSpoofingResult struct {
	Passed *bool     `json:"Passed,omitempty" xml:"Passed,omitempty"`
	Risks  []*string `json:"Risks,omitempty" xml:"Risks,omitempty" type:"Repeated"`
}

func (s IdDetectionAndScanDocumentResponseBodyDataIdDetectResultAntiSpoofingResult) String() string {
	return tea.Prettify(s)
}

func (s IdDetectionAndScanDocumentResponseBodyDataIdDetectResultAntiSpoofingResult) GoString() string {
	return s.String()
}

func (s *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultAntiSpoofingResult) SetPassed(v bool) *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultAntiSpoofingResult {
	s.Passed = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultAntiSpoofingResult) SetRisks(v []*string) *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultAntiSpoofingResult {
	s.Risks = v
	return s
}

type IdDetectionAndScanDocumentResponseBodyDataIdDetectResultDetectionResult struct {
	CardRectangle     []*int64 `json:"CardRectangle,omitempty" xml:"CardRectangle,omitempty" type:"Repeated"`
	Height            *int64   `json:"Height,omitempty" xml:"Height,omitempty"`
	PortraitRectangle []*int64 `json:"PortraitRectangle,omitempty" xml:"PortraitRectangle,omitempty" type:"Repeated"`
	Width             *int64   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s IdDetectionAndScanDocumentResponseBodyDataIdDetectResultDetectionResult) String() string {
	return tea.Prettify(s)
}

func (s IdDetectionAndScanDocumentResponseBodyDataIdDetectResultDetectionResult) GoString() string {
	return s.String()
}

func (s *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultDetectionResult) SetCardRectangle(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultDetectionResult {
	s.CardRectangle = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultDetectionResult) SetHeight(v int64) *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultDetectionResult {
	s.Height = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultDetectionResult) SetPortraitRectangle(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultDetectionResult {
	s.PortraitRectangle = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultDetectionResult) SetWidth(v int64) *IdDetectionAndScanDocumentResponseBodyDataIdDetectResultDetectionResult {
	s.Width = &v
	return s
}

type IdDetectionAndScanDocumentResponseBodyDataOcrResult struct {
	Address                    *string  `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressConfidence          *float64 `json:"AddressConfidence,omitempty" xml:"AddressConfidence,omitempty"`
	AddressPosition            []*int64 `json:"AddressPosition,omitempty" xml:"AddressPosition,omitempty" type:"Repeated"`
	DateOfBirth                *string  `json:"DateOfBirth,omitempty" xml:"DateOfBirth,omitempty"`
	DateOfBirthConfidence      *float64 `json:"DateOfBirthConfidence,omitempty" xml:"DateOfBirthConfidence,omitempty"`
	DateOfBirthPosition        []*int64 `json:"DateOfBirthPosition,omitempty" xml:"DateOfBirthPosition,omitempty" type:"Repeated"`
	DocumentNumber             *string  `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	DocumentNumber2            *string  `json:"DocumentNumber2,omitempty" xml:"DocumentNumber2,omitempty"`
	DocumentNumber2Confidence  *float64 `json:"DocumentNumber2Confidence,omitempty" xml:"DocumentNumber2Confidence,omitempty"`
	DocumentNumber2Position    []*int64 `json:"DocumentNumber2Position,omitempty" xml:"DocumentNumber2Position,omitempty" type:"Repeated"`
	DocumentNumberConfidence   *float64 `json:"DocumentNumberConfidence,omitempty" xml:"DocumentNumberConfidence,omitempty"`
	DocumentNumberPosition     []*int64 `json:"DocumentNumberPosition,omitempty" xml:"DocumentNumberPosition,omitempty" type:"Repeated"`
	ExpirationDate             *string  `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	ExpirationDateConfidence   *float64 `json:"ExpirationDateConfidence,omitempty" xml:"ExpirationDateConfidence,omitempty"`
	ExpirationDatePosition     []*int64 `json:"ExpirationDatePosition,omitempty" xml:"ExpirationDatePosition,omitempty" type:"Repeated"`
	FirstName                  *string  `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	FirstNameConfidence        *float64 `json:"FirstNameConfidence,omitempty" xml:"FirstNameConfidence,omitempty"`
	FirstNamePosition          []*int64 `json:"FirstNamePosition,omitempty" xml:"FirstNamePosition,omitempty" type:"Repeated"`
	FullName                   *string  `json:"FullName,omitempty" xml:"FullName,omitempty"`
	FullNameConfidence         *float64 `json:"FullNameConfidence,omitempty" xml:"FullNameConfidence,omitempty"`
	FullNamePosition           []*int64 `json:"FullNamePosition,omitempty" xml:"FullNamePosition,omitempty" type:"Repeated"`
	IssuedDate                 *string  `json:"IssuedDate,omitempty" xml:"IssuedDate,omitempty"`
	IssuedDateConfidence       *float64 `json:"IssuedDateConfidence,omitempty" xml:"IssuedDateConfidence,omitempty"`
	IssuedDatePosition         []*int64 `json:"IssuedDatePosition,omitempty" xml:"IssuedDatePosition,omitempty" type:"Repeated"`
	IssuingAuthority           *string  `json:"IssuingAuthority,omitempty" xml:"IssuingAuthority,omitempty"`
	IssuingAuthorityConfidence *float64 `json:"IssuingAuthorityConfidence,omitempty" xml:"IssuingAuthorityConfidence,omitempty"`
	IssuingAuthorityPosition   []*int64 `json:"IssuingAuthorityPosition,omitempty" xml:"IssuingAuthorityPosition,omitempty" type:"Repeated"`
	LastName                   *string  `json:"LastName,omitempty" xml:"LastName,omitempty"`
	LastNameConfidence         *float64 `json:"LastNameConfidence,omitempty" xml:"LastNameConfidence,omitempty"`
	LastNamePosition           []*int64 `json:"LastNamePosition,omitempty" xml:"LastNamePosition,omitempty" type:"Repeated"`
	NationalityCode            *string  `json:"NationalityCode,omitempty" xml:"NationalityCode,omitempty"`
	NationalityCodeConfidence  *float64 `json:"NationalityCodeConfidence,omitempty" xml:"NationalityCodeConfidence,omitempty"`
	NationalityCodePosition    []*int64 `json:"NationalityCodePosition,omitempty" xml:"NationalityCodePosition,omitempty" type:"Repeated"`
	NormalizedDateOfBirth      *string  `json:"NormalizedDateOfBirth,omitempty" xml:"NormalizedDateOfBirth,omitempty"`
	NormalizedExpirationDate   *string  `json:"NormalizedExpirationDate,omitempty" xml:"NormalizedExpirationDate,omitempty"`
	NormalizedIssuedDate       *string  `json:"NormalizedIssuedDate,omitempty" xml:"NormalizedIssuedDate,omitempty"`
	NormalizedNationalityCode  *string  `json:"NormalizedNationalityCode,omitempty" xml:"NormalizedNationalityCode,omitempty"`
	NormalizedSex              *string  `json:"NormalizedSex,omitempty" xml:"NormalizedSex,omitempty"`
	Sex                        *string  `json:"Sex,omitempty" xml:"Sex,omitempty"`
	SexConfidence              *float64 `json:"SexConfidence,omitempty" xml:"SexConfidence,omitempty"`
	SexPosition                []*int64 `json:"SexPosition,omitempty" xml:"SexPosition,omitempty" type:"Repeated"`
}

func (s IdDetectionAndScanDocumentResponseBodyDataOcrResult) String() string {
	return tea.Prettify(s)
}

func (s IdDetectionAndScanDocumentResponseBodyDataOcrResult) GoString() string {
	return s.String()
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetAddress(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.Address = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetAddressConfidence(v float64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.AddressConfidence = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetAddressPosition(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.AddressPosition = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetDateOfBirth(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.DateOfBirth = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetDateOfBirthConfidence(v float64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.DateOfBirthConfidence = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetDateOfBirthPosition(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.DateOfBirthPosition = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetDocumentNumber(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.DocumentNumber = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetDocumentNumber2(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.DocumentNumber2 = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetDocumentNumber2Confidence(v float64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.DocumentNumber2Confidence = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetDocumentNumber2Position(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.DocumentNumber2Position = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetDocumentNumberConfidence(v float64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.DocumentNumberConfidence = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetDocumentNumberPosition(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.DocumentNumberPosition = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetExpirationDate(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.ExpirationDate = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetExpirationDateConfidence(v float64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.ExpirationDateConfidence = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetExpirationDatePosition(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.ExpirationDatePosition = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetFirstName(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.FirstName = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetFirstNameConfidence(v float64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.FirstNameConfidence = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetFirstNamePosition(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.FirstNamePosition = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetFullName(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.FullName = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetFullNameConfidence(v float64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.FullNameConfidence = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetFullNamePosition(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.FullNamePosition = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetIssuedDate(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.IssuedDate = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetIssuedDateConfidence(v float64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.IssuedDateConfidence = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetIssuedDatePosition(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.IssuedDatePosition = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetIssuingAuthority(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.IssuingAuthority = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetIssuingAuthorityConfidence(v float64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.IssuingAuthorityConfidence = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetIssuingAuthorityPosition(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.IssuingAuthorityPosition = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetLastName(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.LastName = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetLastNameConfidence(v float64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.LastNameConfidence = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetLastNamePosition(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.LastNamePosition = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetNationalityCode(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.NationalityCode = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetNationalityCodeConfidence(v float64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.NationalityCodeConfidence = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetNationalityCodePosition(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.NationalityCodePosition = v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetNormalizedDateOfBirth(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.NormalizedDateOfBirth = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetNormalizedExpirationDate(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.NormalizedExpirationDate = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetNormalizedIssuedDate(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.NormalizedIssuedDate = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetNormalizedNationalityCode(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.NormalizedNationalityCode = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetNormalizedSex(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.NormalizedSex = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetSex(v string) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.Sex = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetSexConfidence(v float64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.SexConfidence = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponseBodyDataOcrResult) SetSexPosition(v []*int64) *IdDetectionAndScanDocumentResponseBodyDataOcrResult {
	s.SexPosition = v
	return s
}

type IdDetectionAndScanDocumentResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IdDetectionAndScanDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IdDetectionAndScanDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s IdDetectionAndScanDocumentResponse) GoString() string {
	return s.String()
}

func (s *IdDetectionAndScanDocumentResponse) SetHeaders(v map[string]*string) *IdDetectionAndScanDocumentResponse {
	s.Headers = v
	return s
}

func (s *IdDetectionAndScanDocumentResponse) SetStatusCode(v int32) *IdDetectionAndScanDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *IdDetectionAndScanDocumentResponse) SetBody(v *IdDetectionAndScanDocumentResponseBody) *IdDetectionAndScanDocumentResponse {
	s.Body = v
	return s
}

type ScanDocumentRequest struct {
	// The country or region code of the certificate. Specify the code in the ISO 3166 alpha-3 format.
	CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty"`
	// The type of the certificate.
	//
	// Valid values:
	//
	// *   IDENTITY_CARD
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   DRIVER_LICENSE
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   RESIDENCE_CARD
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   PASSPORT
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	DocumentType *string `json:"documentType,omitempty" xml:"documentType,omitempty"`
	// Indicates whether the uploaded front-side image is in an image format or PDF format. The default value is pic.
	ImageAType *string `json:"imageAType,omitempty" xml:"imageAType,omitempty"`
	// Indicates whether the uploaded back-side image is in an image format or PDF format. The default value is pic.
	ImageBType *string `json:"imageBType,omitempty" xml:"imageBType,omitempty"`
	// The Base64-encoded front-side image of the certificate. Either this parameter or the imageUrlA parameter must be specified.
	ImageDataA *string `json:"imageDataA,omitempty" xml:"imageDataA,omitempty"`
	// The Base64-encoded back-side image of the certificate.
	ImageDataB *string `json:"imageDataB,omitempty" xml:"imageDataB,omitempty"`
	// The URL of the front-side image of the certificate. Either this parameter or the imageDataA parameter must be specified.
	ImageUrlA *string `json:"imageUrlA,omitempty" xml:"imageUrlA,omitempty"`
	// The URL of the back-side image of the certificate.
	ImageUrlB *string `json:"imageUrlB,omitempty" xml:"imageUrlB,omitempty"`
}

func (s ScanDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s ScanDocumentRequest) GoString() string {
	return s.String()
}

func (s *ScanDocumentRequest) SetCountryCode(v string) *ScanDocumentRequest {
	s.CountryCode = &v
	return s
}

func (s *ScanDocumentRequest) SetDocumentType(v string) *ScanDocumentRequest {
	s.DocumentType = &v
	return s
}

func (s *ScanDocumentRequest) SetImageAType(v string) *ScanDocumentRequest {
	s.ImageAType = &v
	return s
}

func (s *ScanDocumentRequest) SetImageBType(v string) *ScanDocumentRequest {
	s.ImageBType = &v
	return s
}

func (s *ScanDocumentRequest) SetImageDataA(v string) *ScanDocumentRequest {
	s.ImageDataA = &v
	return s
}

func (s *ScanDocumentRequest) SetImageDataB(v string) *ScanDocumentRequest {
	s.ImageDataB = &v
	return s
}

func (s *ScanDocumentRequest) SetImageUrlA(v string) *ScanDocumentRequest {
	s.ImageUrlA = &v
	return s
}

func (s *ScanDocumentRequest) SetImageUrlB(v string) *ScanDocumentRequest {
	s.ImageUrlB = &v
	return s
}

type ScanDocumentResponseBody struct {
	// The error code. A value of 0 indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ScanDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The error messages.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request is successful.
	//
	// Valid values:
	//
	// *   true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ScanDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScanDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *ScanDocumentResponseBody) SetCode(v string) *ScanDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *ScanDocumentResponseBody) SetData(v *ScanDocumentResponseBodyData) *ScanDocumentResponseBody {
	s.Data = v
	return s
}

func (s *ScanDocumentResponseBody) SetHttpCode(v int64) *ScanDocumentResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ScanDocumentResponseBody) SetMessage(v string) *ScanDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *ScanDocumentResponseBody) SetOk(v bool) *ScanDocumentResponseBody {
	s.Ok = &v
	return s
}

func (s *ScanDocumentResponseBody) SetRequestId(v string) *ScanDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScanDocumentResponseBody) SetStatus(v string) *ScanDocumentResponseBody {
	s.Status = &v
	return s
}

type ScanDocumentResponseBodyData struct {
	// The address.
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The confidence level of the recognized Address.
	AddressConfidence *float64 `json:"AddressConfidence,omitempty" xml:"AddressConfidence,omitempty"`
	// The coordinates of four vertices of the Address field area in the clockwise sequence: the upper left, upper right, lower right, and lower left.
	AddressPosition []*int64 `json:"AddressPosition,omitempty" xml:"AddressPosition,omitempty" type:"Repeated"`
	// The date of birth.
	DateOfBirth *string `json:"DateOfBirth,omitempty" xml:"DateOfBirth,omitempty"`
	// The confidence level of the recognized DateOfBirth.
	DateOfBirthConfidence *float64 `json:"DateOfBirthConfidence,omitempty" xml:"DateOfBirthConfidence,omitempty"`
	// The coordinates of four vertices of the DateOfBirth field area in the clockwise sequence: the upper left, upper right, lower right, and lower left.
	DateOfBirthPosition []*int64 `json:"DateOfBirthPosition,omitempty" xml:"DateOfBirthPosition,omitempty" type:"Repeated"`
	// The certificate number.
	DocumentNumber *string `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	// The secondary certificate number.
	DocumentNumber2 *string `json:"DocumentNumber2,omitempty" xml:"DocumentNumber2,omitempty"`
	// The confidence level of the recognized DocumentNumber2.
	DocumentNumber2Confidence *float64 `json:"DocumentNumber2Confidence,omitempty" xml:"DocumentNumber2Confidence,omitempty"`
	// The coordinates of four vertices of the DocumentNumber2 field area in the clockwise sequence: the upper left, upper right, lower right, and lower left.
	DocumentNumber2Position []*int64 `json:"DocumentNumber2Position,omitempty" xml:"DocumentNumber2Position,omitempty" type:"Repeated"`
	// The confidence level of the recognized DocumentNumber.
	DocumentNumberConfidence *float64 `json:"DocumentNumberConfidence,omitempty" xml:"DocumentNumberConfidence,omitempty"`
	// The coordinates of four vertices of the DocumentNumber field area in the clockwise sequence: the upper left, upper right, lower right, and lower left.
	DocumentNumberPosition []*int64 `json:"DocumentNumberPosition,omitempty" xml:"DocumentNumberPosition,omitempty" type:"Repeated"`
	// The expiration date of the certificate.
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// The confidence level of the recognized ExpirationDate.
	ExpirationDateConfidence *float64 `json:"ExpirationDateConfidence,omitempty" xml:"ExpirationDateConfidence,omitempty"`
	// The coordinates of four vertices of the ExpirationDate field area in the clockwise sequence: the upper left, upper right, lower right, and lower left.
	ExpirationDatePosition []*int64 `json:"ExpirationDatePosition,omitempty" xml:"ExpirationDatePosition,omitempty" type:"Repeated"`
	// The first name.
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// The confidence level of the recognized FirstName field.
	FirstNameConfidence *float64 `json:"FirstNameConfidence,omitempty" xml:"FirstNameConfidence,omitempty"`
	// The coordinates of four vertices of the FirstName field area in the clockwise sequence: the upper left, upper right, lower right, and lower left.
	FirstNamePosition []*int64 `json:"FirstNamePosition,omitempty" xml:"FirstNamePosition,omitempty" type:"Repeated"`
	// The full name.
	FullName *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	// The confidence level of the recognized FullName.
	FullNameConfidence *float64 `json:"FullNameConfidence,omitempty" xml:"FullNameConfidence,omitempty"`
	// The coordinates of four vertices of the FullName field area in the clockwise sequence: the upper left, upper right, lower right, and lower left.
	FullNamePosition []*int64 `json:"FullNamePosition,omitempty" xml:"FullNamePosition,omitempty" type:"Repeated"`
	// The date of issue.
	IssuedDate *string `json:"IssuedDate,omitempty" xml:"IssuedDate,omitempty"`
	// The confidence level of the recognized IssuedDate.
	IssuedDateConfidence *float64 `json:"IssuedDateConfidence,omitempty" xml:"IssuedDateConfidence,omitempty"`
	// The coordinates of four vertices of the IssuedDate field area in the clockwise sequence: the upper left, upper right, lower right, and lower left.
	IssuedDatePosition []*int64 `json:"IssuedDatePosition,omitempty" xml:"IssuedDatePosition,omitempty" type:"Repeated"`
	// The organization that issued the certificate.
	IssuingAuthority *string `json:"IssuingAuthority,omitempty" xml:"IssuingAuthority,omitempty"`
	// The confidence level of the recognized IssuingAuthority.
	IssuingAuthorityConfidence *float64 `json:"IssuingAuthorityConfidence,omitempty" xml:"IssuingAuthorityConfidence,omitempty"`
	// The coordinates of four vertices of the IssuingAuthority field area in the clockwise sequence: the upper left, upper right, lower right, and lower left.
	IssuingAuthorityPosition []*int64 `json:"IssuingAuthorityPosition,omitempty" xml:"IssuingAuthorityPosition,omitempty" type:"Repeated"`
	// The last name.
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// The confidence level of the recognized LastName.
	LastNameConfidence *float64 `json:"LastNameConfidence,omitempty" xml:"LastNameConfidence,omitempty"`
	// The coordinates of four vertices of the LastName field area in the clockwise sequence: the upper left, upper right, lower right, and lower left.
	LastNamePosition []*int64 `json:"LastNamePosition,omitempty" xml:"LastNamePosition,omitempty" type:"Repeated"`
	// The country or region code on the certificate.
	NationalityCode *string `json:"NationalityCode,omitempty" xml:"NationalityCode,omitempty"`
	// The confidence level of the recognized NationalityCode.
	NationalityCodeConfidence *float64 `json:"NationalityCodeConfidence,omitempty" xml:"NationalityCodeConfidence,omitempty"`
	// The coordinates of four vertices of the NationalityCode field area in the clockwise sequence: the upper left, upper right, lower right, and lower left.
	NationalityCodePosition []*int64 `json:"NationalityCodePosition,omitempty" xml:"NationalityCodePosition,omitempty" type:"Repeated"`
	// The date of birth in the format of YYYY/MM/dd.
	NormalizedDateOfBirth *string `json:"NormalizedDateOfBirth,omitempty" xml:"NormalizedDateOfBirth,omitempty"`
	// The expiration date in the format of YYYY/MM/dd.
	NormalizedExpirationDate *string `json:"NormalizedExpirationDate,omitempty" xml:"NormalizedExpirationDate,omitempty"`
	// The date of issue in the format of YYYY/MM/dd.
	NormalizedIssuedDate *string `json:"NormalizedIssuedDate,omitempty" xml:"NormalizedIssuedDate,omitempty"`
	// The country or region code in the ISO 3166 alpha-3 format.
	NormalizedNationalityCode *string `json:"NormalizedNationalityCode,omitempty" xml:"NormalizedNationalityCode,omitempty"`
	// The gender.
	//
	// Valid values:
	//
	// *   null
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   F
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     female
	//
	//     <!-- -->
	//
	// *   M
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     male
	//
	//     <!-- -->
	NormalizedSex *string `json:"NormalizedSex,omitempty" xml:"NormalizedSex,omitempty"`
	// The gender.
	Sex *string `json:"Sex,omitempty" xml:"Sex,omitempty"`
	// The confidence level of the recognized Sex.
	SexConfidence *float64 `json:"SexConfidence,omitempty" xml:"SexConfidence,omitempty"`
	// The coordinates of four vertices of the Sex field area in the clockwise sequence: the upper left, upper right, lower right, and lower left.
	SexPosition []*int64 `json:"SexPosition,omitempty" xml:"SexPosition,omitempty" type:"Repeated"`
}

func (s ScanDocumentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ScanDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ScanDocumentResponseBodyData) SetAddress(v string) *ScanDocumentResponseBodyData {
	s.Address = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetAddressConfidence(v float64) *ScanDocumentResponseBodyData {
	s.AddressConfidence = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetAddressPosition(v []*int64) *ScanDocumentResponseBodyData {
	s.AddressPosition = v
	return s
}

func (s *ScanDocumentResponseBodyData) SetDateOfBirth(v string) *ScanDocumentResponseBodyData {
	s.DateOfBirth = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetDateOfBirthConfidence(v float64) *ScanDocumentResponseBodyData {
	s.DateOfBirthConfidence = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetDateOfBirthPosition(v []*int64) *ScanDocumentResponseBodyData {
	s.DateOfBirthPosition = v
	return s
}

func (s *ScanDocumentResponseBodyData) SetDocumentNumber(v string) *ScanDocumentResponseBodyData {
	s.DocumentNumber = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetDocumentNumber2(v string) *ScanDocumentResponseBodyData {
	s.DocumentNumber2 = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetDocumentNumber2Confidence(v float64) *ScanDocumentResponseBodyData {
	s.DocumentNumber2Confidence = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetDocumentNumber2Position(v []*int64) *ScanDocumentResponseBodyData {
	s.DocumentNumber2Position = v
	return s
}

func (s *ScanDocumentResponseBodyData) SetDocumentNumberConfidence(v float64) *ScanDocumentResponseBodyData {
	s.DocumentNumberConfidence = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetDocumentNumberPosition(v []*int64) *ScanDocumentResponseBodyData {
	s.DocumentNumberPosition = v
	return s
}

func (s *ScanDocumentResponseBodyData) SetExpirationDate(v string) *ScanDocumentResponseBodyData {
	s.ExpirationDate = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetExpirationDateConfidence(v float64) *ScanDocumentResponseBodyData {
	s.ExpirationDateConfidence = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetExpirationDatePosition(v []*int64) *ScanDocumentResponseBodyData {
	s.ExpirationDatePosition = v
	return s
}

func (s *ScanDocumentResponseBodyData) SetFirstName(v string) *ScanDocumentResponseBodyData {
	s.FirstName = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetFirstNameConfidence(v float64) *ScanDocumentResponseBodyData {
	s.FirstNameConfidence = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetFirstNamePosition(v []*int64) *ScanDocumentResponseBodyData {
	s.FirstNamePosition = v
	return s
}

func (s *ScanDocumentResponseBodyData) SetFullName(v string) *ScanDocumentResponseBodyData {
	s.FullName = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetFullNameConfidence(v float64) *ScanDocumentResponseBodyData {
	s.FullNameConfidence = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetFullNamePosition(v []*int64) *ScanDocumentResponseBodyData {
	s.FullNamePosition = v
	return s
}

func (s *ScanDocumentResponseBodyData) SetIssuedDate(v string) *ScanDocumentResponseBodyData {
	s.IssuedDate = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetIssuedDateConfidence(v float64) *ScanDocumentResponseBodyData {
	s.IssuedDateConfidence = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetIssuedDatePosition(v []*int64) *ScanDocumentResponseBodyData {
	s.IssuedDatePosition = v
	return s
}

func (s *ScanDocumentResponseBodyData) SetIssuingAuthority(v string) *ScanDocumentResponseBodyData {
	s.IssuingAuthority = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetIssuingAuthorityConfidence(v float64) *ScanDocumentResponseBodyData {
	s.IssuingAuthorityConfidence = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetIssuingAuthorityPosition(v []*int64) *ScanDocumentResponseBodyData {
	s.IssuingAuthorityPosition = v
	return s
}

func (s *ScanDocumentResponseBodyData) SetLastName(v string) *ScanDocumentResponseBodyData {
	s.LastName = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetLastNameConfidence(v float64) *ScanDocumentResponseBodyData {
	s.LastNameConfidence = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetLastNamePosition(v []*int64) *ScanDocumentResponseBodyData {
	s.LastNamePosition = v
	return s
}

func (s *ScanDocumentResponseBodyData) SetNationalityCode(v string) *ScanDocumentResponseBodyData {
	s.NationalityCode = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetNationalityCodeConfidence(v float64) *ScanDocumentResponseBodyData {
	s.NationalityCodeConfidence = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetNationalityCodePosition(v []*int64) *ScanDocumentResponseBodyData {
	s.NationalityCodePosition = v
	return s
}

func (s *ScanDocumentResponseBodyData) SetNormalizedDateOfBirth(v string) *ScanDocumentResponseBodyData {
	s.NormalizedDateOfBirth = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetNormalizedExpirationDate(v string) *ScanDocumentResponseBodyData {
	s.NormalizedExpirationDate = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetNormalizedIssuedDate(v string) *ScanDocumentResponseBodyData {
	s.NormalizedIssuedDate = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetNormalizedNationalityCode(v string) *ScanDocumentResponseBodyData {
	s.NormalizedNationalityCode = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetNormalizedSex(v string) *ScanDocumentResponseBodyData {
	s.NormalizedSex = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetSex(v string) *ScanDocumentResponseBodyData {
	s.Sex = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetSexConfidence(v float64) *ScanDocumentResponseBodyData {
	s.SexConfidence = &v
	return s
}

func (s *ScanDocumentResponseBodyData) SetSexPosition(v []*int64) *ScanDocumentResponseBodyData {
	s.SexPosition = v
	return s
}

type ScanDocumentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScanDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScanDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s ScanDocumentResponse) GoString() string {
	return s.String()
}

func (s *ScanDocumentResponse) SetHeaders(v map[string]*string) *ScanDocumentResponse {
	s.Headers = v
	return s
}

func (s *ScanDocumentResponse) SetStatusCode(v int32) *ScanDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *ScanDocumentResponse) SetBody(v *ScanDocumentResponseBody) *ScanDocumentResponse {
	s.Body = v
	return s
}

type VerifyDocumentRequest struct {
	CountryCode  *string `json:"countryCode,omitempty" xml:"countryCode,omitempty"`
	DetThickness *bool   `json:"detThickness,omitempty" xml:"detThickness,omitempty"`
	DocumentType *string `json:"documentType,omitempty" xml:"documentType,omitempty"`
	ImageAType   *string `json:"imageAType,omitempty" xml:"imageAType,omitempty"`
	ImageBType   *string `json:"imageBType,omitempty" xml:"imageBType,omitempty"`
	ImageCType   *string `json:"imageCType,omitempty" xml:"imageCType,omitempty"`
	ImageDataA   *string `json:"imageDataA,omitempty" xml:"imageDataA,omitempty"`
	ImageDataB   *string `json:"imageDataB,omitempty" xml:"imageDataB,omitempty"`
	ImageDataC   *string `json:"imageDataC,omitempty" xml:"imageDataC,omitempty"`
	ImageUrlA    *string `json:"imageUrlA,omitempty" xml:"imageUrlA,omitempty"`
	ImageUrlB    *string `json:"imageUrlB,omitempty" xml:"imageUrlB,omitempty"`
	ImageUrlC    *string `json:"imageUrlC,omitempty" xml:"imageUrlC,omitempty"`
}

func (s VerifyDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyDocumentRequest) GoString() string {
	return s.String()
}

func (s *VerifyDocumentRequest) SetCountryCode(v string) *VerifyDocumentRequest {
	s.CountryCode = &v
	return s
}

func (s *VerifyDocumentRequest) SetDetThickness(v bool) *VerifyDocumentRequest {
	s.DetThickness = &v
	return s
}

func (s *VerifyDocumentRequest) SetDocumentType(v string) *VerifyDocumentRequest {
	s.DocumentType = &v
	return s
}

func (s *VerifyDocumentRequest) SetImageAType(v string) *VerifyDocumentRequest {
	s.ImageAType = &v
	return s
}

func (s *VerifyDocumentRequest) SetImageBType(v string) *VerifyDocumentRequest {
	s.ImageBType = &v
	return s
}

func (s *VerifyDocumentRequest) SetImageCType(v string) *VerifyDocumentRequest {
	s.ImageCType = &v
	return s
}

func (s *VerifyDocumentRequest) SetImageDataA(v string) *VerifyDocumentRequest {
	s.ImageDataA = &v
	return s
}

func (s *VerifyDocumentRequest) SetImageDataB(v string) *VerifyDocumentRequest {
	s.ImageDataB = &v
	return s
}

func (s *VerifyDocumentRequest) SetImageDataC(v string) *VerifyDocumentRequest {
	s.ImageDataC = &v
	return s
}

func (s *VerifyDocumentRequest) SetImageUrlA(v string) *VerifyDocumentRequest {
	s.ImageUrlA = &v
	return s
}

func (s *VerifyDocumentRequest) SetImageUrlB(v string) *VerifyDocumentRequest {
	s.ImageUrlB = &v
	return s
}

func (s *VerifyDocumentRequest) SetImageUrlC(v string) *VerifyDocumentRequest {
	s.ImageUrlC = &v
	return s
}

type VerifyDocumentResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *VerifyDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpCode  *int64                          `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Ok        *bool                           `json:"Ok,omitempty" xml:"Ok,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                         `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s VerifyDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyDocumentResponseBody) SetCode(v string) *VerifyDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyDocumentResponseBody) SetData(v *VerifyDocumentResponseBodyData) *VerifyDocumentResponseBody {
	s.Data = v
	return s
}

func (s *VerifyDocumentResponseBody) SetHttpCode(v int64) *VerifyDocumentResponseBody {
	s.HttpCode = &v
	return s
}

func (s *VerifyDocumentResponseBody) SetMessage(v string) *VerifyDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyDocumentResponseBody) SetOk(v bool) *VerifyDocumentResponseBody {
	s.Ok = &v
	return s
}

func (s *VerifyDocumentResponseBody) SetRequestId(v string) *VerifyDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyDocumentResponseBody) SetStatus(v string) *VerifyDocumentResponseBody {
	s.Status = &v
	return s
}

type VerifyDocumentResponseBodyData struct {
	AntiSpoofingResult *VerifyDocumentResponseBodyDataAntiSpoofingResult `json:"AntiSpoofingResult,omitempty" xml:"AntiSpoofingResult,omitempty" type:"Struct"`
	CountryCode        *string                                           `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	DetectionResult    *VerifyDocumentResponseBodyDataDetectionResult    `json:"DetectionResult,omitempty" xml:"DetectionResult,omitempty" type:"Struct"`
	DocumentType       *string                                           `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	Passed             *bool                                             `json:"Passed,omitempty" xml:"Passed,omitempty"`
}

func (s VerifyDocumentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VerifyDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyDocumentResponseBodyData) SetAntiSpoofingResult(v *VerifyDocumentResponseBodyDataAntiSpoofingResult) *VerifyDocumentResponseBodyData {
	s.AntiSpoofingResult = v
	return s
}

func (s *VerifyDocumentResponseBodyData) SetCountryCode(v string) *VerifyDocumentResponseBodyData {
	s.CountryCode = &v
	return s
}

func (s *VerifyDocumentResponseBodyData) SetDetectionResult(v *VerifyDocumentResponseBodyDataDetectionResult) *VerifyDocumentResponseBodyData {
	s.DetectionResult = v
	return s
}

func (s *VerifyDocumentResponseBodyData) SetDocumentType(v string) *VerifyDocumentResponseBodyData {
	s.DocumentType = &v
	return s
}

func (s *VerifyDocumentResponseBodyData) SetPassed(v bool) *VerifyDocumentResponseBodyData {
	s.Passed = &v
	return s
}

type VerifyDocumentResponseBodyDataAntiSpoofingResult struct {
	Passed *bool     `json:"Passed,omitempty" xml:"Passed,omitempty"`
	Risks  []*string `json:"Risks,omitempty" xml:"Risks,omitempty" type:"Repeated"`
}

func (s VerifyDocumentResponseBodyDataAntiSpoofingResult) String() string {
	return tea.Prettify(s)
}

func (s VerifyDocumentResponseBodyDataAntiSpoofingResult) GoString() string {
	return s.String()
}

func (s *VerifyDocumentResponseBodyDataAntiSpoofingResult) SetPassed(v bool) *VerifyDocumentResponseBodyDataAntiSpoofingResult {
	s.Passed = &v
	return s
}

func (s *VerifyDocumentResponseBodyDataAntiSpoofingResult) SetRisks(v []*string) *VerifyDocumentResponseBodyDataAntiSpoofingResult {
	s.Risks = v
	return s
}

type VerifyDocumentResponseBodyDataDetectionResult struct {
	CardRectangle     []*int64 `json:"CardRectangle,omitempty" xml:"CardRectangle,omitempty" type:"Repeated"`
	Height            *int64   `json:"Height,omitempty" xml:"Height,omitempty"`
	PortraitRectangle []*int64 `json:"PortraitRectangle,omitempty" xml:"PortraitRectangle,omitempty" type:"Repeated"`
	Width             *int64   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s VerifyDocumentResponseBodyDataDetectionResult) String() string {
	return tea.Prettify(s)
}

func (s VerifyDocumentResponseBodyDataDetectionResult) GoString() string {
	return s.String()
}

func (s *VerifyDocumentResponseBodyDataDetectionResult) SetCardRectangle(v []*int64) *VerifyDocumentResponseBodyDataDetectionResult {
	s.CardRectangle = v
	return s
}

func (s *VerifyDocumentResponseBodyDataDetectionResult) SetHeight(v int64) *VerifyDocumentResponseBodyDataDetectionResult {
	s.Height = &v
	return s
}

func (s *VerifyDocumentResponseBodyDataDetectionResult) SetPortraitRectangle(v []*int64) *VerifyDocumentResponseBodyDataDetectionResult {
	s.PortraitRectangle = v
	return s
}

func (s *VerifyDocumentResponseBodyDataDetectionResult) SetWidth(v int64) *VerifyDocumentResponseBodyDataDetectionResult {
	s.Width = &v
	return s
}

type VerifyDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyDocumentResponse) GoString() string {
	return s.String()
}

func (s *VerifyDocumentResponse) SetHeaders(v map[string]*string) *VerifyDocumentResponse {
	s.Headers = v
	return s
}

func (s *VerifyDocumentResponse) SetStatusCode(v int32) *VerifyDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyDocumentResponse) SetBody(v *VerifyDocumentResponseBody) *VerifyDocumentResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ekyc-saas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) FaceRecognitionCompareWithOptions(request *FaceRecognitionCompareRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FaceRecognitionCompareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdCardImageData)) {
		body["idCardImageData"] = request.IdCardImageData
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardImageType)) {
		body["idCardImageType"] = request.IdCardImageType
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardImageUrl)) {
		body["idCardImageUrl"] = request.IdCardImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PortraitImageData)) {
		body["portraitImageData"] = request.PortraitImageData
	}

	if !tea.BoolValue(util.IsUnset(request.PortraitImageType)) {
		body["portraitImageType"] = request.PortraitImageType
	}

	if !tea.BoolValue(util.IsUnset(request.PortraitImageUrl)) {
		body["portraitImageUrl"] = request.PortraitImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.QualityControl)) {
		body["qualityControl"] = request.QualityControl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceRecognitionCompare"),
		Version:     tea.String("2023-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/face_recognition_compare"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceRecognitionCompareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceRecognitionCompare(request *FaceRecognitionCompareRequest) (_result *FaceRecognitionCompareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FaceRecognitionCompareResponse{}
	_body, _err := client.FaceRecognitionCompareWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceToFaceCompareWithOptions(request *FaceToFaceCompareRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FaceToFaceCompareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PortraitImageData1)) {
		body["portraitImageData1"] = request.PortraitImageData1
	}

	if !tea.BoolValue(util.IsUnset(request.PortraitImageData2)) {
		body["portraitImageData2"] = request.PortraitImageData2
	}

	if !tea.BoolValue(util.IsUnset(request.PortraitImageType1)) {
		body["portraitImageType1"] = request.PortraitImageType1
	}

	if !tea.BoolValue(util.IsUnset(request.PortraitImageType2)) {
		body["portraitImageType2"] = request.PortraitImageType2
	}

	if !tea.BoolValue(util.IsUnset(request.PortraitImageUrl1)) {
		body["portraitImageUrl1"] = request.PortraitImageUrl1
	}

	if !tea.BoolValue(util.IsUnset(request.PortraitImageUrl2)) {
		body["portraitImageUrl2"] = request.PortraitImageUrl2
	}

	if !tea.BoolValue(util.IsUnset(request.QualityControl)) {
		body["qualityControl"] = request.QualityControl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceToFaceCompare"),
		Version:     tea.String("2023-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/face_to_face_compare"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceToFaceCompareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceToFaceCompare(request *FaceToFaceCompareRequest) (_result *FaceToFaceCompareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FaceToFaceCompareResponse{}
	_body, _err := client.FaceToFaceCompareWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IdDetectOcrCompareFacesWithOptions(request *IdDetectOcrCompareFacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *IdDetectOcrCompareFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OCR)) {
		body["OCR"] = request.OCR
	}

	if !tea.BoolValue(util.IsUnset(request.Unrepeat)) {
		body["Unrepeat"] = request.Unrepeat
	}

	if !tea.BoolValue(util.IsUnset(request.CardDetect)) {
		body["cardDetect"] = request.CardDetect
	}

	if !tea.BoolValue(util.IsUnset(request.CountryCode)) {
		body["countryCode"] = request.CountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentType)) {
		body["documentType"] = request.DocumentType
	}

	if !tea.BoolValue(util.IsUnset(request.FaceCompare)) {
		body["faceCompare"] = request.FaceCompare
	}

	if !tea.BoolValue(util.IsUnset(request.ImageAType)) {
		body["imageAType"] = request.ImageAType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageBType)) {
		body["imageBType"] = request.ImageBType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageDataA)) {
		body["imageDataA"] = request.ImageDataA
	}

	if !tea.BoolValue(util.IsUnset(request.ImageDataB)) {
		body["imageDataB"] = request.ImageDataB
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrlA)) {
		body["imageUrlA"] = request.ImageUrlA
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrlB)) {
		body["imageUrlB"] = request.ImageUrlB
	}

	if !tea.BoolValue(util.IsUnset(request.PortraitImageData)) {
		body["portraitImageData"] = request.PortraitImageData
	}

	if !tea.BoolValue(util.IsUnset(request.PortraitImageUrl)) {
		body["portraitImageUrl"] = request.PortraitImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.QualityControl)) {
		body["qualityControl"] = request.QualityControl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("IdDetectOcrCompareFaces"),
		Version:     tea.String("2023-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/id_detect_ocr_compare_faces"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &IdDetectOcrCompareFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IdDetectOcrCompareFaces(request *IdDetectOcrCompareFacesRequest) (_result *IdDetectOcrCompareFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &IdDetectOcrCompareFacesResponse{}
	_body, _err := client.IdDetectOcrCompareFacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IdDetectionAndScanDocumentWithOptions(request *IdDetectionAndScanDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *IdDetectionAndScanDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CountryCode)) {
		body["countryCode"] = request.CountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentType)) {
		body["documentType"] = request.DocumentType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageAType)) {
		body["imageAType"] = request.ImageAType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageBType)) {
		body["imageBType"] = request.ImageBType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageDataA)) {
		body["imageDataA"] = request.ImageDataA
	}

	if !tea.BoolValue(util.IsUnset(request.ImageDataB)) {
		body["imageDataB"] = request.ImageDataB
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrlA)) {
		body["imageUrlA"] = request.ImageUrlA
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrlB)) {
		body["imageUrlB"] = request.ImageUrlB
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("IdDetectionAndScanDocument"),
		Version:     tea.String("2023-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/id_detection_and_scan_document"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &IdDetectionAndScanDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IdDetectionAndScanDocument(request *IdDetectionAndScanDocumentRequest) (_result *IdDetectionAndScanDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &IdDetectionAndScanDocumentResponse{}
	_body, _err := client.IdDetectionAndScanDocumentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After you upload the front-side image and back-side image (optional) of the certificate, the system returns the structured OCR results.
 * ### Limits
 * *   Format: JPEG or PNG.
 * *   Resolution: The length of an image is no greater than 4000 pixels, and the width is no smaller than 400 pixels. Recommended size: 1024\\*768.
 * *   Size: the size of a single image. Valid values: \\[4KB, 6MB].
 * ### Data protocol
 * *   Request: application/json.
 * *   Response: application/json.
 * *   Binary data: Base64.
 *
 * @param request ScanDocumentRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ScanDocumentResponse
 */
func (client *Client) ScanDocumentWithOptions(request *ScanDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ScanDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CountryCode)) {
		body["countryCode"] = request.CountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentType)) {
		body["documentType"] = request.DocumentType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageAType)) {
		body["imageAType"] = request.ImageAType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageBType)) {
		body["imageBType"] = request.ImageBType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageDataA)) {
		body["imageDataA"] = request.ImageDataA
	}

	if !tea.BoolValue(util.IsUnset(request.ImageDataB)) {
		body["imageDataB"] = request.ImageDataB
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrlA)) {
		body["imageUrlA"] = request.ImageUrlA
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrlB)) {
		body["imageUrlB"] = request.ImageUrlB
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ScanDocument"),
		Version:     tea.String("2023-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/scan_document"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ScanDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After you upload the front-side image and back-side image (optional) of the certificate, the system returns the structured OCR results.
 * ### Limits
 * *   Format: JPEG or PNG.
 * *   Resolution: The length of an image is no greater than 4000 pixels, and the width is no smaller than 400 pixels. Recommended size: 1024\\*768.
 * *   Size: the size of a single image. Valid values: \\[4KB, 6MB].
 * ### Data protocol
 * *   Request: application/json.
 * *   Response: application/json.
 * *   Binary data: Base64.
 *
 * @param request ScanDocumentRequest
 * @return ScanDocumentResponse
 */
func (client *Client) ScanDocument(request *ScanDocumentRequest) (_result *ScanDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScanDocumentResponse{}
	_body, _err := client.ScanDocumentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyDocumentWithOptions(request *VerifyDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VerifyDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CountryCode)) {
		body["countryCode"] = request.CountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.DetThickness)) {
		body["detThickness"] = request.DetThickness
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentType)) {
		body["documentType"] = request.DocumentType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageAType)) {
		body["imageAType"] = request.ImageAType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageBType)) {
		body["imageBType"] = request.ImageBType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCType)) {
		body["imageCType"] = request.ImageCType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageDataA)) {
		body["imageDataA"] = request.ImageDataA
	}

	if !tea.BoolValue(util.IsUnset(request.ImageDataB)) {
		body["imageDataB"] = request.ImageDataB
	}

	if !tea.BoolValue(util.IsUnset(request.ImageDataC)) {
		body["imageDataC"] = request.ImageDataC
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrlA)) {
		body["imageUrlA"] = request.ImageUrlA
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrlB)) {
		body["imageUrlB"] = request.ImageUrlB
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrlC)) {
		body["imageUrlC"] = request.ImageUrlC
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyDocument"),
		Version:     tea.String("2023-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/verify_document"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyDocument(request *VerifyDocumentRequest) (_result *VerifyDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifyDocumentResponse{}
	_body, _err := client.VerifyDocumentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
