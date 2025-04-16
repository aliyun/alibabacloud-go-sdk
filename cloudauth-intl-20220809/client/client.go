// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CardOcrRequest struct {
	// example:
	//
	// 00000006
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// F
	IdFaceQuality      *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// example:
	//
	// https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	// example:
	//
	// dso9322***dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// T
	Ocr *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	// example:
	//
	// ID_OCR_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// F
	Spoof *string `json:"Spoof,omitempty" xml:"Spoof,omitempty"`
}

func (s CardOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s CardOcrRequest) GoString() string {
	return s.String()
}

func (s *CardOcrRequest) SetDocType(v string) *CardOcrRequest {
	s.DocType = &v
	return s
}

func (s *CardOcrRequest) SetIdFaceQuality(v string) *CardOcrRequest {
	s.IdFaceQuality = &v
	return s
}

func (s *CardOcrRequest) SetIdOcrPictureBase64(v string) *CardOcrRequest {
	s.IdOcrPictureBase64 = &v
	return s
}

func (s *CardOcrRequest) SetIdOcrPictureUrl(v string) *CardOcrRequest {
	s.IdOcrPictureUrl = &v
	return s
}

func (s *CardOcrRequest) SetMerchantBizId(v string) *CardOcrRequest {
	s.MerchantBizId = &v
	return s
}

func (s *CardOcrRequest) SetMerchantUserId(v string) *CardOcrRequest {
	s.MerchantUserId = &v
	return s
}

func (s *CardOcrRequest) SetOcr(v string) *CardOcrRequest {
	s.Ocr = &v
	return s
}

func (s *CardOcrRequest) SetProductCode(v string) *CardOcrRequest {
	s.ProductCode = &v
	return s
}

func (s *CardOcrRequest) SetSpoof(v string) *CardOcrRequest {
	s.Spoof = &v
	return s
}

type CardOcrResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB356FE-BB6A-5DCC-B4C5-E8051787EBA1
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CardOcrResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CardOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CardOcrResponseBody) GoString() string {
	return s.String()
}

func (s *CardOcrResponseBody) SetCode(v string) *CardOcrResponseBody {
	s.Code = &v
	return s
}

func (s *CardOcrResponseBody) SetMessage(v string) *CardOcrResponseBody {
	s.Message = &v
	return s
}

func (s *CardOcrResponseBody) SetRequestId(v string) *CardOcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CardOcrResponseBody) SetResult(v *CardOcrResponseBodyResult) *CardOcrResponseBody {
	s.Result = v
	return s
}

type CardOcrResponseBodyResult struct {
	ExtCardInfo *string `json:"ExtCardInfo,omitempty" xml:"ExtCardInfo,omitempty"`
	ExtIdInfo   *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CardOcrResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CardOcrResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CardOcrResponseBodyResult) SetExtCardInfo(v string) *CardOcrResponseBodyResult {
	s.ExtCardInfo = &v
	return s
}

func (s *CardOcrResponseBodyResult) SetExtIdInfo(v string) *CardOcrResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *CardOcrResponseBodyResult) SetPassed(v string) *CardOcrResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *CardOcrResponseBodyResult) SetSubCode(v string) *CardOcrResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *CardOcrResponseBodyResult) SetTransactionId(v string) *CardOcrResponseBodyResult {
	s.TransactionId = &v
	return s
}

type CardOcrResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CardOcrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CardOcrResponse) String() string {
	return tea.Prettify(s)
}

func (s CardOcrResponse) GoString() string {
	return s.String()
}

func (s *CardOcrResponse) SetHeaders(v map[string]*string) *CardOcrResponse {
	s.Headers = v
	return s
}

func (s *CardOcrResponse) SetStatusCode(v int32) *CardOcrResponse {
	s.StatusCode = &v
	return s
}

func (s *CardOcrResponse) SetBody(v *CardOcrResponseBody) *CardOcrResponse {
	s.Body = v
	return s
}

type CheckResultRequest struct {
	// example:
	//
	// ***
	ExtraImageControlList *string `json:"ExtraImageControlList,omitempty" xml:"ExtraImageControlList,omitempty"`
	// example:
	//
	// N
	IsReturnImage *string `json:"IsReturnImage,omitempty" xml:"IsReturnImage,omitempty"`
	// example:
	//
	// djs20d***9-dsskc
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// Y
	ReturnFiveCategorySpoofResult *string `json:"ReturnFiveCategorySpoofResult,omitempty" xml:"ReturnFiveCategorySpoofResult,omitempty"`
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CheckResultRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckResultRequest) GoString() string {
	return s.String()
}

func (s *CheckResultRequest) SetExtraImageControlList(v string) *CheckResultRequest {
	s.ExtraImageControlList = &v
	return s
}

func (s *CheckResultRequest) SetIsReturnImage(v string) *CheckResultRequest {
	s.IsReturnImage = &v
	return s
}

func (s *CheckResultRequest) SetMerchantBizId(v string) *CheckResultRequest {
	s.MerchantBizId = &v
	return s
}

func (s *CheckResultRequest) SetReturnFiveCategorySpoofResult(v string) *CheckResultRequest {
	s.ReturnFiveCategorySpoofResult = &v
	return s
}

func (s *CheckResultRequest) SetTransactionId(v string) *CheckResultRequest {
	s.TransactionId = &v
	return s
}

type CheckResultResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CheckResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CheckResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *CheckResultResponseBody) SetCode(v string) *CheckResultResponseBody {
	s.Code = &v
	return s
}

func (s *CheckResultResponseBody) SetMessage(v string) *CheckResultResponseBody {
	s.Message = &v
	return s
}

func (s *CheckResultResponseBody) SetRequestId(v string) *CheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckResultResponseBody) SetResult(v *CheckResultResponseBodyResult) *CheckResultResponseBody {
	s.Result = v
	return s
}

type CheckResultResponseBodyResult struct {
	// example:
	//
	// **
	EkycResult *string `json:"EkycResult,omitempty" xml:"EkycResult,omitempty"`
	// example:
	//
	// **
	ExtBasicInfo *string `json:"ExtBasicInfo,omitempty" xml:"ExtBasicInfo,omitempty"`
	// example:
	//
	// **
	ExtFaceInfo *string `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty"`
	// example:
	//
	// **
	ExtIdInfo *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	ExtInfo   *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	// example:
	//
	// **
	ExtRiskInfo *string `json:"ExtRiskInfo,omitempty" xml:"ExtRiskInfo,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// ***
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s CheckResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CheckResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckResultResponseBodyResult) SetEkycResult(v string) *CheckResultResponseBodyResult {
	s.EkycResult = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtBasicInfo(v string) *CheckResultResponseBodyResult {
	s.ExtBasicInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtFaceInfo(v string) *CheckResultResponseBodyResult {
	s.ExtFaceInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtIdInfo(v string) *CheckResultResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtInfo(v string) *CheckResultResponseBodyResult {
	s.ExtInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtRiskInfo(v string) *CheckResultResponseBodyResult {
	s.ExtRiskInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetPassed(v string) *CheckResultResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetSubCode(v string) *CheckResultResponseBodyResult {
	s.SubCode = &v
	return s
}

type CheckResultResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckResultResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckResultResponse) GoString() string {
	return s.String()
}

func (s *CheckResultResponse) SetHeaders(v map[string]*string) *CheckResultResponse {
	s.Headers = v
	return s
}

func (s *CheckResultResponse) SetStatusCode(v int32) *CheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckResultResponse) SetBody(v *CheckResultResponseBody) *CheckResultResponse {
	s.Body = v
	return s
}

type CheckVerifyLogRequest struct {
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// hksb7ba1b*********015d694361bee4
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CheckVerifyLogRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckVerifyLogRequest) GoString() string {
	return s.String()
}

func (s *CheckVerifyLogRequest) SetMerchantBizId(v string) *CheckVerifyLogRequest {
	s.MerchantBizId = &v
	return s
}

func (s *CheckVerifyLogRequest) SetTransactionId(v string) *CheckVerifyLogRequest {
	s.TransactionId = &v
	return s
}

type CheckVerifyLogResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CheckVerifyLogResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CheckVerifyLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckVerifyLogResponseBody) GoString() string {
	return s.String()
}

func (s *CheckVerifyLogResponseBody) SetCode(v string) *CheckVerifyLogResponseBody {
	s.Code = &v
	return s
}

func (s *CheckVerifyLogResponseBody) SetMessage(v string) *CheckVerifyLogResponseBody {
	s.Message = &v
	return s
}

func (s *CheckVerifyLogResponseBody) SetRequestId(v string) *CheckVerifyLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckVerifyLogResponseBody) SetResult(v *CheckVerifyLogResponseBodyResult) *CheckVerifyLogResponseBody {
	s.Result = v
	return s
}

type CheckVerifyLogResponseBodyResult struct {
	// example:
	//
	// {}
	ExtInfo       *string   `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	InterruptPage *string   `json:"InterruptPage,omitempty" xml:"InterruptPage,omitempty"`
	LogInfo       []*string `json:"LogInfo,omitempty" xml:"LogInfo,omitempty" type:"Repeated"`
	// example:
	//
	// {
	//
	//           "faceOverTimes": 0,
	//
	//           "hasFaceOverTimes": false,
	//
	//           "hasFacePermissionRefuse": false,
	//
	//           "hasOcrEdit": true,
	//
	//           "hasOcrEditOverTimes": false,
	//
	//           "hasOcrOverTimes": true,
	//
	//           "hasOcrPermissionRefuse": false,
	//
	//           "ocrEditOverTimes": 0,
	//
	//           "ocrEditTimes": 1,
	//
	//           "ocrOverTimes": 1,
	//
	//           "pageStayTimeInfo": {
	//
	//             "LOADING": "1615",
	//
	//             "OCR_SCAN": "37446",
	//
	//             "OCR_RESULT": "1338",
	//
	//             "FACE": "8707"
	//
	//           }
	//
	//         }
	LogStatisticsInfo *string `json:"LogStatisticsInfo,omitempty" xml:"LogStatisticsInfo,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// example:
	//
	// 1001
	VerifyErrorCode *string `json:"VerifyErrorCode,omitempty" xml:"VerifyErrorCode,omitempty"`
	// example:
	//
	// 1
	VerifyStatus *string `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
}

func (s CheckVerifyLogResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CheckVerifyLogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckVerifyLogResponseBodyResult) SetExtInfo(v string) *CheckVerifyLogResponseBodyResult {
	s.ExtInfo = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetInterruptPage(v string) *CheckVerifyLogResponseBodyResult {
	s.InterruptPage = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetLogInfo(v []*string) *CheckVerifyLogResponseBodyResult {
	s.LogInfo = v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetLogStatisticsInfo(v string) *CheckVerifyLogResponseBodyResult {
	s.LogStatisticsInfo = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetPassed(v string) *CheckVerifyLogResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetSubCode(v string) *CheckVerifyLogResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetVerifyErrorCode(v string) *CheckVerifyLogResponseBodyResult {
	s.VerifyErrorCode = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetVerifyStatus(v string) *CheckVerifyLogResponseBodyResult {
	s.VerifyStatus = &v
	return s
}

type CheckVerifyLogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckVerifyLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckVerifyLogResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckVerifyLogResponse) GoString() string {
	return s.String()
}

func (s *CheckVerifyLogResponse) SetHeaders(v map[string]*string) *CheckVerifyLogResponse {
	s.Headers = v
	return s
}

func (s *CheckVerifyLogResponse) SetStatusCode(v int32) *CheckVerifyLogResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckVerifyLogResponse) SetBody(v *CheckVerifyLogResponseBody) *CheckVerifyLogResponse {
	s.Body = v
	return s
}

type DeepfakeDetectIntlRequest struct {
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceBase64 *string `json:"FaceBase64,omitempty" xml:"FaceBase64,omitempty"`
	// example:
	//
	// IMAGE
	FaceInputType *string `json:"FaceInputType,omitempty" xml:"FaceInputType,omitempty"`
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceUrl *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c******
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FACE_DEEPFAKE
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s DeepfakeDetectIntlRequest) String() string {
	return tea.Prettify(s)
}

func (s DeepfakeDetectIntlRequest) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectIntlRequest) SetFaceBase64(v string) *DeepfakeDetectIntlRequest {
	s.FaceBase64 = &v
	return s
}

func (s *DeepfakeDetectIntlRequest) SetFaceInputType(v string) *DeepfakeDetectIntlRequest {
	s.FaceInputType = &v
	return s
}

func (s *DeepfakeDetectIntlRequest) SetFaceUrl(v string) *DeepfakeDetectIntlRequest {
	s.FaceUrl = &v
	return s
}

func (s *DeepfakeDetectIntlRequest) SetMerchantBizId(v string) *DeepfakeDetectIntlRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DeepfakeDetectIntlRequest) SetProductCode(v string) *DeepfakeDetectIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *DeepfakeDetectIntlRequest) SetSceneCode(v string) *DeepfakeDetectIntlRequest {
	s.SceneCode = &v
	return s
}

type DeepfakeDetectIntlResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DeepfakeDetectIntlResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DeepfakeDetectIntlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeepfakeDetectIntlResponseBody) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectIntlResponseBody) SetCode(v string) *DeepfakeDetectIntlResponseBody {
	s.Code = &v
	return s
}

func (s *DeepfakeDetectIntlResponseBody) SetMessage(v string) *DeepfakeDetectIntlResponseBody {
	s.Message = &v
	return s
}

func (s *DeepfakeDetectIntlResponseBody) SetRequestId(v string) *DeepfakeDetectIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeepfakeDetectIntlResponseBody) SetResultObject(v *DeepfakeDetectIntlResponseBodyResultObject) *DeepfakeDetectIntlResponseBody {
	s.ResultObject = v
	return s
}

type DeepfakeDetectIntlResponseBodyResultObject struct {
	// example:
	//
	// 1
	Result    *string            `json:"Result,omitempty" xml:"Result,omitempty"`
	RiskScore map[string]*string `json:"RiskScore,omitempty" xml:"RiskScore,omitempty"`
	// example:
	//
	// SuspectDeepForgery,SuspectWarterMark
	RiskTag *string `json:"RiskTag,omitempty" xml:"RiskTag,omitempty"`
}

func (s DeepfakeDetectIntlResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s DeepfakeDetectIntlResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectIntlResponseBodyResultObject) SetResult(v string) *DeepfakeDetectIntlResponseBodyResultObject {
	s.Result = &v
	return s
}

func (s *DeepfakeDetectIntlResponseBodyResultObject) SetRiskScore(v map[string]*string) *DeepfakeDetectIntlResponseBodyResultObject {
	s.RiskScore = v
	return s
}

func (s *DeepfakeDetectIntlResponseBodyResultObject) SetRiskTag(v string) *DeepfakeDetectIntlResponseBodyResultObject {
	s.RiskTag = &v
	return s
}

type DeepfakeDetectIntlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeepfakeDetectIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeepfakeDetectIntlResponse) String() string {
	return tea.Prettify(s)
}

func (s DeepfakeDetectIntlResponse) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectIntlResponse) SetHeaders(v map[string]*string) *DeepfakeDetectIntlResponse {
	s.Headers = v
	return s
}

func (s *DeepfakeDetectIntlResponse) SetStatusCode(v int32) *DeepfakeDetectIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeepfakeDetectIntlResponse) SetBody(v *DeepfakeDetectIntlResponseBody) *DeepfakeDetectIntlResponse {
	s.Body = v
	return s
}

type DeleteVerifyResultRequest struct {
	// example:
	//
	// Y / N
	DeleteAfterQuery *string `json:"DeleteAfterQuery,omitempty" xml:"DeleteAfterQuery,omitempty"`
	// example:
	//
	// Img / Text / All
	DeleteType *string `json:"DeleteType,omitempty" xml:"DeleteType,omitempty"`
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s DeleteVerifyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifyResultRequest) GoString() string {
	return s.String()
}

func (s *DeleteVerifyResultRequest) SetDeleteAfterQuery(v string) *DeleteVerifyResultRequest {
	s.DeleteAfterQuery = &v
	return s
}

func (s *DeleteVerifyResultRequest) SetDeleteType(v string) *DeleteVerifyResultRequest {
	s.DeleteType = &v
	return s
}

func (s *DeleteVerifyResultRequest) SetTransactionId(v string) *DeleteVerifyResultRequest {
	s.TransactionId = &v
	return s
}

type DeleteVerifyResultResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeleteVerifyResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteVerifyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVerifyResultResponseBody) SetCode(v string) *DeleteVerifyResultResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVerifyResultResponseBody) SetMessage(v string) *DeleteVerifyResultResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVerifyResultResponseBody) SetRequestId(v string) *DeleteVerifyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVerifyResultResponseBody) SetResult(v *DeleteVerifyResultResponseBodyResult) *DeleteVerifyResultResponseBody {
	s.Result = v
	return s
}

type DeleteVerifyResultResponseBodyResult struct {
	// example:
	//
	// Y/N
	DeleteResult *string `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty"`
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s DeleteVerifyResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifyResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteVerifyResultResponseBodyResult) SetDeleteResult(v string) *DeleteVerifyResultResponseBodyResult {
	s.DeleteResult = &v
	return s
}

func (s *DeleteVerifyResultResponseBodyResult) SetTransactionId(v string) *DeleteVerifyResultResponseBodyResult {
	s.TransactionId = &v
	return s
}

type DeleteVerifyResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVerifyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *DeleteVerifyResultResponse) SetHeaders(v map[string]*string) *DeleteVerifyResultResponse {
	s.Headers = v
	return s
}

func (s *DeleteVerifyResultResponse) SetStatusCode(v int32) *DeleteVerifyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVerifyResultResponse) SetBody(v *DeleteVerifyResultResponseBody) *DeleteVerifyResultResponse {
	s.Body = v
	return s
}

type DocOcrRequest struct {
	CardSide *string `json:"CardSide,omitempty" xml:"CardSide,omitempty"`
	// example:
	//
	// 00000006
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// F
	IdFaceQuality      *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// example:
	//
	// https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	IdThreshold     *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// example:
	//
	// dso9322***dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// T
	Ocr         *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// F
	Spoof *string `json:"Spoof,omitempty" xml:"Spoof,omitempty"`
}

func (s DocOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s DocOcrRequest) GoString() string {
	return s.String()
}

func (s *DocOcrRequest) SetCardSide(v string) *DocOcrRequest {
	s.CardSide = &v
	return s
}

func (s *DocOcrRequest) SetDocType(v string) *DocOcrRequest {
	s.DocType = &v
	return s
}

func (s *DocOcrRequest) SetIdFaceQuality(v string) *DocOcrRequest {
	s.IdFaceQuality = &v
	return s
}

func (s *DocOcrRequest) SetIdOcrPictureBase64(v string) *DocOcrRequest {
	s.IdOcrPictureBase64 = &v
	return s
}

func (s *DocOcrRequest) SetIdOcrPictureUrl(v string) *DocOcrRequest {
	s.IdOcrPictureUrl = &v
	return s
}

func (s *DocOcrRequest) SetIdThreshold(v string) *DocOcrRequest {
	s.IdThreshold = &v
	return s
}

func (s *DocOcrRequest) SetMerchantBizId(v string) *DocOcrRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DocOcrRequest) SetMerchantUserId(v string) *DocOcrRequest {
	s.MerchantUserId = &v
	return s
}

func (s *DocOcrRequest) SetOcr(v string) *DocOcrRequest {
	s.Ocr = &v
	return s
}

func (s *DocOcrRequest) SetProductCode(v string) *DocOcrRequest {
	s.ProductCode = &v
	return s
}

func (s *DocOcrRequest) SetSpoof(v string) *DocOcrRequest {
	s.Spoof = &v
	return s
}

type DocOcrResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 86C40EC3-5940-5F47-995C-BFE90B70E540
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DocOcrResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DocOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DocOcrResponseBody) GoString() string {
	return s.String()
}

func (s *DocOcrResponseBody) SetCode(v string) *DocOcrResponseBody {
	s.Code = &v
	return s
}

func (s *DocOcrResponseBody) SetMessage(v string) *DocOcrResponseBody {
	s.Message = &v
	return s
}

func (s *DocOcrResponseBody) SetRequestId(v string) *DocOcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocOcrResponseBody) SetResult(v *DocOcrResponseBodyResult) *DocOcrResponseBody {
	s.Result = v
	return s
}

type DocOcrResponseBodyResult struct {
	ExtIdInfo *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s DocOcrResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DocOcrResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DocOcrResponseBodyResult) SetExtIdInfo(v string) *DocOcrResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *DocOcrResponseBodyResult) SetPassed(v string) *DocOcrResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *DocOcrResponseBodyResult) SetSubCode(v string) *DocOcrResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *DocOcrResponseBodyResult) SetTransactionId(v string) *DocOcrResponseBodyResult {
	s.TransactionId = &v
	return s
}

type DocOcrResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocOcrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocOcrResponse) String() string {
	return tea.Prettify(s)
}

func (s DocOcrResponse) GoString() string {
	return s.String()
}

func (s *DocOcrResponse) SetHeaders(v map[string]*string) *DocOcrResponse {
	s.Headers = v
	return s
}

func (s *DocOcrResponse) SetStatusCode(v int32) *DocOcrResponse {
	s.StatusCode = &v
	return s
}

func (s *DocOcrResponse) SetBody(v *DocOcrResponseBody) *DocOcrResponse {
	s.Body = v
	return s
}

type EkycVerifyRequest struct {
	// example:
	//
	// T
	Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	// example:
	//
	// F
	Crop    *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// 410***************
	DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
	// example:
	//
	// 00000001
	DocType           *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// example:
	//
	// https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	FacePictureUrl     *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// example:
	//
	// https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	IdThreshold     *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 123456
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// eKYC_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s EkycVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s EkycVerifyRequest) GoString() string {
	return s.String()
}

func (s *EkycVerifyRequest) SetAuthorize(v string) *EkycVerifyRequest {
	s.Authorize = &v
	return s
}

func (s *EkycVerifyRequest) SetCrop(v string) *EkycVerifyRequest {
	s.Crop = &v
	return s
}

func (s *EkycVerifyRequest) SetDocName(v string) *EkycVerifyRequest {
	s.DocName = &v
	return s
}

func (s *EkycVerifyRequest) SetDocNo(v string) *EkycVerifyRequest {
	s.DocNo = &v
	return s
}

func (s *EkycVerifyRequest) SetDocType(v string) *EkycVerifyRequest {
	s.DocType = &v
	return s
}

func (s *EkycVerifyRequest) SetFacePictureBase64(v string) *EkycVerifyRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *EkycVerifyRequest) SetFacePictureUrl(v string) *EkycVerifyRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *EkycVerifyRequest) SetIdOcrPictureBase64(v string) *EkycVerifyRequest {
	s.IdOcrPictureBase64 = &v
	return s
}

func (s *EkycVerifyRequest) SetIdOcrPictureUrl(v string) *EkycVerifyRequest {
	s.IdOcrPictureUrl = &v
	return s
}

func (s *EkycVerifyRequest) SetIdThreshold(v string) *EkycVerifyRequest {
	s.IdThreshold = &v
	return s
}

func (s *EkycVerifyRequest) SetMerchantBizId(v string) *EkycVerifyRequest {
	s.MerchantBizId = &v
	return s
}

func (s *EkycVerifyRequest) SetMerchantUserId(v string) *EkycVerifyRequest {
	s.MerchantUserId = &v
	return s
}

func (s *EkycVerifyRequest) SetProductCode(v string) *EkycVerifyRequest {
	s.ProductCode = &v
	return s
}

type EkycVerifyResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB356FE-BB6A-5DCC-B4C5-E8051787EBA1
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *EkycVerifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EkycVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EkycVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *EkycVerifyResponseBody) SetCode(v string) *EkycVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *EkycVerifyResponseBody) SetMessage(v string) *EkycVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *EkycVerifyResponseBody) SetRequestId(v string) *EkycVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *EkycVerifyResponseBody) SetResult(v *EkycVerifyResponseBodyResult) *EkycVerifyResponseBody {
	s.Result = v
	return s
}

type EkycVerifyResponseBodyResult struct {
	// example:
	//
	// {
	//
	// "faceAttack": "N",
	//
	// "faceComparisonScore": 52.57,
	//
	// "facePassed": "N",
	//
	// "authorityComparisonScore": 80.39
	//
	// }
	ExtFaceInfo *string `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty"`
	ExtIdInfo   *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 205
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s EkycVerifyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EkycVerifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EkycVerifyResponseBodyResult) SetExtFaceInfo(v string) *EkycVerifyResponseBodyResult {
	s.ExtFaceInfo = &v
	return s
}

func (s *EkycVerifyResponseBodyResult) SetExtIdInfo(v string) *EkycVerifyResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *EkycVerifyResponseBodyResult) SetPassed(v string) *EkycVerifyResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *EkycVerifyResponseBodyResult) SetSubCode(v string) *EkycVerifyResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *EkycVerifyResponseBodyResult) SetTransactionId(v string) *EkycVerifyResponseBodyResult {
	s.TransactionId = &v
	return s
}

type EkycVerifyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EkycVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EkycVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s EkycVerifyResponse) GoString() string {
	return s.String()
}

func (s *EkycVerifyResponse) SetHeaders(v map[string]*string) *EkycVerifyResponse {
	s.Headers = v
	return s
}

func (s *EkycVerifyResponse) SetStatusCode(v int32) *EkycVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *EkycVerifyResponse) SetBody(v *EkycVerifyResponseBody) *EkycVerifyResponse {
	s.Body = v
	return s
}

type FaceCompareRequest struct {
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId     *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	TargetFacePicture    *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// example:
	//
	// https://***face2.jpeg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
}

func (s FaceCompareRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareRequest) GoString() string {
	return s.String()
}

func (s *FaceCompareRequest) SetMerchantBizId(v string) *FaceCompareRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceCompareRequest) SetSourceFacePicture(v string) *FaceCompareRequest {
	s.SourceFacePicture = &v
	return s
}

func (s *FaceCompareRequest) SetSourceFacePictureUrl(v string) *FaceCompareRequest {
	s.SourceFacePictureUrl = &v
	return s
}

func (s *FaceCompareRequest) SetTargetFacePicture(v string) *FaceCompareRequest {
	s.TargetFacePicture = &v
	return s
}

func (s *FaceCompareRequest) SetTargetFacePictureUrl(v string) *FaceCompareRequest {
	s.TargetFacePictureUrl = &v
	return s
}

type FaceCompareResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB356FE-BB6A-5DCC-B4C5-E8051787EBA1
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *FaceCompareResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceCompareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareResponseBody) GoString() string {
	return s.String()
}

func (s *FaceCompareResponseBody) SetCode(v string) *FaceCompareResponseBody {
	s.Code = &v
	return s
}

func (s *FaceCompareResponseBody) SetMessage(v string) *FaceCompareResponseBody {
	s.Message = &v
	return s
}

func (s *FaceCompareResponseBody) SetRequestId(v string) *FaceCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceCompareResponseBody) SetResult(v *FaceCompareResponseBodyResult) *FaceCompareResponseBody {
	s.Result = v
	return s
}

type FaceCompareResponseBodyResult struct {
	// example:
	//
	// 98
	FaceComparisonScore *float64 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceCompareResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceCompareResponseBodyResult) SetFaceComparisonScore(v float64) *FaceCompareResponseBodyResult {
	s.FaceComparisonScore = &v
	return s
}

func (s *FaceCompareResponseBodyResult) SetPassed(v string) *FaceCompareResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *FaceCompareResponseBodyResult) SetTransactionId(v string) *FaceCompareResponseBodyResult {
	s.TransactionId = &v
	return s
}

type FaceCompareResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceCompareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceCompareResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareResponse) GoString() string {
	return s.String()
}

func (s *FaceCompareResponse) SetHeaders(v map[string]*string) *FaceCompareResponse {
	s.Headers = v
	return s
}

func (s *FaceCompareResponse) SetStatusCode(v int32) *FaceCompareResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceCompareResponse) SetBody(v *FaceCompareResponseBody) *FaceCompareResponse {
	s.Body = v
	return s
}

type FaceGuardRiskRequest struct {
	// example:
	//
	// LMALL20******001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// Tk9SSUQuMS*****************ZDNmNWY5NzQxOW1o
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	// example:
	//
	// 0c83ce0101d34eff886b1f7d1cdef67f
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// FACE_GUARD
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s FaceGuardRiskRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceGuardRiskRequest) GoString() string {
	return s.String()
}

func (s *FaceGuardRiskRequest) SetBizId(v string) *FaceGuardRiskRequest {
	s.BizId = &v
	return s
}

func (s *FaceGuardRiskRequest) SetDeviceToken(v string) *FaceGuardRiskRequest {
	s.DeviceToken = &v
	return s
}

func (s *FaceGuardRiskRequest) SetMerchantBizId(v string) *FaceGuardRiskRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceGuardRiskRequest) SetProductCode(v string) *FaceGuardRiskRequest {
	s.ProductCode = &v
	return s
}

type FaceGuardRiskResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 595E387B-3F0E-5C52-BD02-8EFE63D41FD5
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *FaceGuardRiskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceGuardRiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceGuardRiskResponseBody) GoString() string {
	return s.String()
}

func (s *FaceGuardRiskResponseBody) SetCode(v string) *FaceGuardRiskResponseBody {
	s.Code = &v
	return s
}

func (s *FaceGuardRiskResponseBody) SetMessage(v string) *FaceGuardRiskResponseBody {
	s.Message = &v
	return s
}

func (s *FaceGuardRiskResponseBody) SetRequestId(v string) *FaceGuardRiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceGuardRiskResponseBody) SetResult(v *FaceGuardRiskResponseBodyResult) *FaceGuardRiskResponseBody {
	s.Result = v
	return s
}

type FaceGuardRiskResponseBodyResult struct {
	GuardRiskScore *float64 `json:"GuardRiskScore,omitempty" xml:"GuardRiskScore,omitempty"`
	RiskExtends    *string  `json:"RiskExtends,omitempty" xml:"RiskExtends,omitempty"`
	// example:
	//
	// ROOT,VPN,HOOK
	RiskTags *string `json:"RiskTags,omitempty" xml:"RiskTags,omitempty"`
	// example:
	//
	// hk573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceGuardRiskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s FaceGuardRiskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceGuardRiskResponseBodyResult) SetGuardRiskScore(v float64) *FaceGuardRiskResponseBodyResult {
	s.GuardRiskScore = &v
	return s
}

func (s *FaceGuardRiskResponseBodyResult) SetRiskExtends(v string) *FaceGuardRiskResponseBodyResult {
	s.RiskExtends = &v
	return s
}

func (s *FaceGuardRiskResponseBodyResult) SetRiskTags(v string) *FaceGuardRiskResponseBodyResult {
	s.RiskTags = &v
	return s
}

func (s *FaceGuardRiskResponseBodyResult) SetTransactionId(v string) *FaceGuardRiskResponseBodyResult {
	s.TransactionId = &v
	return s
}

type FaceGuardRiskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceGuardRiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceGuardRiskResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceGuardRiskResponse) GoString() string {
	return s.String()
}

func (s *FaceGuardRiskResponse) SetHeaders(v map[string]*string) *FaceGuardRiskResponse {
	s.Headers = v
	return s
}

func (s *FaceGuardRiskResponse) SetStatusCode(v int32) *FaceGuardRiskResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceGuardRiskResponse) SetBody(v *FaceGuardRiskResponseBody) *FaceGuardRiskResponse {
	s.Body = v
	return s
}

type FaceLivenessRequest struct {
	// example:
	//
	// T
	Crop              *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// example:
	//
	// https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	// example:
	//
	// T
	FaceQuality *string `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// T
	Occlusion *string `json:"Occlusion,omitempty" xml:"Occlusion,omitempty"`
	// example:
	//
	// FACE_LIVENESS_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s FaceLivenessRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessRequest) GoString() string {
	return s.String()
}

func (s *FaceLivenessRequest) SetCrop(v string) *FaceLivenessRequest {
	s.Crop = &v
	return s
}

func (s *FaceLivenessRequest) SetFacePictureBase64(v string) *FaceLivenessRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *FaceLivenessRequest) SetFacePictureUrl(v string) *FaceLivenessRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *FaceLivenessRequest) SetFaceQuality(v string) *FaceLivenessRequest {
	s.FaceQuality = &v
	return s
}

func (s *FaceLivenessRequest) SetMerchantBizId(v string) *FaceLivenessRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceLivenessRequest) SetMerchantUserId(v string) *FaceLivenessRequest {
	s.MerchantUserId = &v
	return s
}

func (s *FaceLivenessRequest) SetOcclusion(v string) *FaceLivenessRequest {
	s.Occlusion = &v
	return s
}

func (s *FaceLivenessRequest) SetProductCode(v string) *FaceLivenessRequest {
	s.ProductCode = &v
	return s
}

type FaceLivenessResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 42EA58CA-5DF4-55D5-82C4-5E7A40DA62BA
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *FaceLivenessResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceLivenessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessResponseBody) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponseBody) SetCode(v string) *FaceLivenessResponseBody {
	s.Code = &v
	return s
}

func (s *FaceLivenessResponseBody) SetMessage(v string) *FaceLivenessResponseBody {
	s.Message = &v
	return s
}

func (s *FaceLivenessResponseBody) SetRequestId(v string) *FaceLivenessResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceLivenessResponseBody) SetResult(v *FaceLivenessResponseBodyResult) *FaceLivenessResponseBody {
	s.Result = v
	return s
}

type FaceLivenessResponseBodyResult struct {
	ExtFaceInfo *FaceLivenessResponseBodyResultExtFaceInfo `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty" type:"Struct"`
	// example:
	//
	// N
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 205
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceLivenessResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponseBodyResult) SetExtFaceInfo(v *FaceLivenessResponseBodyResultExtFaceInfo) *FaceLivenessResponseBodyResult {
	s.ExtFaceInfo = v
	return s
}

func (s *FaceLivenessResponseBodyResult) SetPassed(v string) *FaceLivenessResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *FaceLivenessResponseBodyResult) SetSubCode(v string) *FaceLivenessResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *FaceLivenessResponseBodyResult) SetTransactionId(v string) *FaceLivenessResponseBodyResult {
	s.TransactionId = &v
	return s
}

type FaceLivenessResponseBodyResultExtFaceInfo struct {
	FaceAge *int32 `json:"FaceAge,omitempty" xml:"FaceAge,omitempty"`
	// example:
	//
	// Y
	FaceAttack *string `json:"FaceAttack,omitempty" xml:"FaceAttack,omitempty"`
	FaceGender *string `json:"FaceGender,omitempty" xml:"FaceGender,omitempty"`
	// example:
	//
	// 87.19
	FaceQualityScore *float64 `json:"FaceQualityScore,omitempty" xml:"FaceQualityScore,omitempty"`
	// example:
	//
	// Y
	OcclusionResult *string `json:"OcclusionResult,omitempty" xml:"OcclusionResult,omitempty"`
}

func (s FaceLivenessResponseBodyResultExtFaceInfo) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessResponseBodyResultExtFaceInfo) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetFaceAge(v int32) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.FaceAge = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetFaceAttack(v string) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.FaceAttack = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetFaceGender(v string) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.FaceGender = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetFaceQualityScore(v float64) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.FaceQualityScore = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetOcclusionResult(v string) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.OcclusionResult = &v
	return s
}

type FaceLivenessResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceLivenessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceLivenessResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessResponse) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponse) SetHeaders(v map[string]*string) *FaceLivenessResponse {
	s.Headers = v
	return s
}

func (s *FaceLivenessResponse) SetStatusCode(v int32) *FaceLivenessResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceLivenessResponse) SetBody(v *FaceLivenessResponseBody) *FaceLivenessResponse {
	s.Body = v
	return s
}

type FraudResultCallBackRequest struct {
	// example:
	//
	// shs2b27333914876c01de4cb22f5841f
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	ExtParams *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	// example:
	//
	// PASS
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// production
	VerifyDeployEnv *string `json:"VerifyDeployEnv,omitempty" xml:"VerifyDeployEnv,omitempty"`
}

func (s FraudResultCallBackRequest) String() string {
	return tea.Prettify(s)
}

func (s FraudResultCallBackRequest) GoString() string {
	return s.String()
}

func (s *FraudResultCallBackRequest) SetCertifyId(v string) *FraudResultCallBackRequest {
	s.CertifyId = &v
	return s
}

func (s *FraudResultCallBackRequest) SetExtParams(v string) *FraudResultCallBackRequest {
	s.ExtParams = &v
	return s
}

func (s *FraudResultCallBackRequest) SetResultCode(v string) *FraudResultCallBackRequest {
	s.ResultCode = &v
	return s
}

func (s *FraudResultCallBackRequest) SetVerifyDeployEnv(v string) *FraudResultCallBackRequest {
	s.VerifyDeployEnv = &v
	return s
}

type FraudResultCallBackResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FraudResultCallBackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FraudResultCallBackResponseBody) GoString() string {
	return s.String()
}

func (s *FraudResultCallBackResponseBody) SetCode(v string) *FraudResultCallBackResponseBody {
	s.Code = &v
	return s
}

func (s *FraudResultCallBackResponseBody) SetMessage(v string) *FraudResultCallBackResponseBody {
	s.Message = &v
	return s
}

func (s *FraudResultCallBackResponseBody) SetRequestId(v string) *FraudResultCallBackResponseBody {
	s.RequestId = &v
	return s
}

func (s *FraudResultCallBackResponseBody) SetSuccess(v bool) *FraudResultCallBackResponseBody {
	s.Success = &v
	return s
}

type FraudResultCallBackResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FraudResultCallBackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FraudResultCallBackResponse) String() string {
	return tea.Prettify(s)
}

func (s FraudResultCallBackResponse) GoString() string {
	return s.String()
}

func (s *FraudResultCallBackResponse) SetHeaders(v map[string]*string) *FraudResultCallBackResponse {
	s.Headers = v
	return s
}

func (s *FraudResultCallBackResponse) SetStatusCode(v int32) *FraudResultCallBackResponse {
	s.StatusCode = &v
	return s
}

func (s *FraudResultCallBackResponse) SetBody(v *FraudResultCallBackResponseBody) *FraudResultCallBackResponse {
	s.Body = v
	return s
}

type Id2MetaPeriodVerifyIntlRequest struct {
	// This parameter is required.
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 411xxxxxxxxxxx0001
	DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ​00000001
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 1234567890
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eKYC_Date_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20301001
	ValidityEndDate *string `json:"ValidityEndDate,omitempty" xml:"ValidityEndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20201001
	ValidityStartDate *string `json:"ValidityStartDate,omitempty" xml:"ValidityStartDate,omitempty"`
}

func (s Id2MetaPeriodVerifyIntlRequest) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaPeriodVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetDocName(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.DocName = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetDocNo(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.DocNo = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetDocType(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.DocType = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetMerchantBizId(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.MerchantBizId = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetMerchantUserId(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.MerchantUserId = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetProductCode(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetSceneCode(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.SceneCode = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetValidityEndDate(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.ValidityEndDate = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetValidityStartDate(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.ValidityStartDate = &v
	return s
}

type Id2MetaPeriodVerifyIntlResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7B97D932-7FF5-517D-BF39-7CA1BEE3CDD9
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *Id2MetaPeriodVerifyIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s Id2MetaPeriodVerifyIntlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaPeriodVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *Id2MetaPeriodVerifyIntlResponseBody) SetCode(v string) *Id2MetaPeriodVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponseBody) SetMessage(v string) *Id2MetaPeriodVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponseBody) SetRequestId(v string) *Id2MetaPeriodVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponseBody) SetResult(v *Id2MetaPeriodVerifyIntlResponseBodyResult) *Id2MetaPeriodVerifyIntlResponseBody {
	s.Result = v
	return s
}

type Id2MetaPeriodVerifyIntlResponseBodyResult struct {
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s Id2MetaPeriodVerifyIntlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaPeriodVerifyIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *Id2MetaPeriodVerifyIntlResponseBodyResult) SetPassed(v string) *Id2MetaPeriodVerifyIntlResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponseBodyResult) SetSubCode(v string) *Id2MetaPeriodVerifyIntlResponseBodyResult {
	s.SubCode = &v
	return s
}

type Id2MetaPeriodVerifyIntlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Id2MetaPeriodVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Id2MetaPeriodVerifyIntlResponse) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaPeriodVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *Id2MetaPeriodVerifyIntlResponse) SetHeaders(v map[string]*string) *Id2MetaPeriodVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponse) SetStatusCode(v int32) *Id2MetaPeriodVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponse) SetBody(v *Id2MetaPeriodVerifyIntlResponseBody) *Id2MetaPeriodVerifyIntlResponse {
	s.Body = v
	return s
}

type Id2MetaVerifyIntlRequest struct {
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// example:
	//
	// ID_2META
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Id2MetaVerifyIntlRequest) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyIntlRequest) SetIdentifyNum(v string) *Id2MetaVerifyIntlRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Id2MetaVerifyIntlRequest) SetParamType(v string) *Id2MetaVerifyIntlRequest {
	s.ParamType = &v
	return s
}

func (s *Id2MetaVerifyIntlRequest) SetProductCode(v string) *Id2MetaVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *Id2MetaVerifyIntlRequest) SetUserName(v string) *Id2MetaVerifyIntlRequest {
	s.UserName = &v
	return s
}

type Id2MetaVerifyIntlResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EFA11401-C961-5E89-A2D3-BF9883E5CC3D
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *Id2MetaVerifyIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s Id2MetaVerifyIntlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyIntlResponseBody) SetCode(v string) *Id2MetaVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *Id2MetaVerifyIntlResponseBody) SetMessage(v string) *Id2MetaVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *Id2MetaVerifyIntlResponseBody) SetRequestId(v string) *Id2MetaVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id2MetaVerifyIntlResponseBody) SetResult(v *Id2MetaVerifyIntlResponseBodyResult) *Id2MetaVerifyIntlResponseBody {
	s.Result = v
	return s
}

type Id2MetaVerifyIntlResponseBodyResult struct {
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
}

func (s Id2MetaVerifyIntlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyIntlResponseBodyResult) SetBizCode(v string) *Id2MetaVerifyIntlResponseBodyResult {
	s.BizCode = &v
	return s
}

type Id2MetaVerifyIntlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Id2MetaVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Id2MetaVerifyIntlResponse) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyIntlResponse) SetHeaders(v map[string]*string) *Id2MetaVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *Id2MetaVerifyIntlResponse) SetStatusCode(v int32) *Id2MetaVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *Id2MetaVerifyIntlResponse) SetBody(v *Id2MetaVerifyIntlResponseBody) *Id2MetaVerifyIntlResponse {
	s.Body = v
	return s
}

type InitializeRequest struct {
	AppQualityCheck *string `json:"AppQualityCheck,omitempty" xml:"AppQualityCheck,omitempty"`
	Authorize       *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	CallbackToken   *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	CallbackUrl     *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// example:
	//
	// *
	Crop        *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DocScanMode *string `json:"DocScanMode,omitempty" xml:"DocScanMode,omitempty"`
	// example:
	//
	// 01000000
	DocType           *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	DocVideo          *string `json:"DocVideo,omitempty" xml:"DocVideo,omitempty"`
	ExperienceCode    *string `json:"ExperienceCode,omitempty" xml:"ExperienceCode,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// example:
	//
	// ***
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	// example:
	//
	// *
	IdFaceQuality *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	// example:
	//
	// Y
	IdSpoof        *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	IdThreshold    *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	LanguageConfig *string `json:"LanguageConfig,omitempty" xml:"LanguageConfig,omitempty"`
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 1221****6543
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// {\\"bioMetaInfo\\":\\"4.1.0:2916352,0\\",\\"deviceType\\":\\"web\\",\\"ua\\":\\"Mozilla/5.0 (Macintosh
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	Model    *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// OCR。
	//
	// example:
	//
	// *
	Ocr               *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	ProcedurePriority *string `json:"ProcedurePriority,omitempty" xml:"ProcedurePriority,omitempty"`
	// example:
	//
	// eKYC
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductFlow *string `json:"ProductFlow,omitempty" xml:"ProductFlow,omitempty"`
	// example:
	//
	// http*****
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	// example:
	//
	// PAY**
	SceneCode     *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	ShowAlbumIcon *string `json:"ShowAlbumIcon,omitempty" xml:"ShowAlbumIcon,omitempty"`
	ShowGuidePage *string `json:"ShowGuidePage,omitempty" xml:"ShowGuidePage,omitempty"`
	ShowOcrResult *string `json:"ShowOcrResult,omitempty" xml:"ShowOcrResult,omitempty"`
	StyleConfig   *string `json:"StyleConfig,omitempty" xml:"StyleConfig,omitempty"`
}

func (s InitializeRequest) String() string {
	return tea.Prettify(s)
}

func (s InitializeRequest) GoString() string {
	return s.String()
}

func (s *InitializeRequest) SetAppQualityCheck(v string) *InitializeRequest {
	s.AppQualityCheck = &v
	return s
}

func (s *InitializeRequest) SetAuthorize(v string) *InitializeRequest {
	s.Authorize = &v
	return s
}

func (s *InitializeRequest) SetCallbackToken(v string) *InitializeRequest {
	s.CallbackToken = &v
	return s
}

func (s *InitializeRequest) SetCallbackUrl(v string) *InitializeRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitializeRequest) SetCrop(v string) *InitializeRequest {
	s.Crop = &v
	return s
}

func (s *InitializeRequest) SetDocScanMode(v string) *InitializeRequest {
	s.DocScanMode = &v
	return s
}

func (s *InitializeRequest) SetDocType(v string) *InitializeRequest {
	s.DocType = &v
	return s
}

func (s *InitializeRequest) SetDocVideo(v string) *InitializeRequest {
	s.DocVideo = &v
	return s
}

func (s *InitializeRequest) SetExperienceCode(v string) *InitializeRequest {
	s.ExperienceCode = &v
	return s
}

func (s *InitializeRequest) SetFacePictureBase64(v string) *InitializeRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *InitializeRequest) SetFacePictureUrl(v string) *InitializeRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *InitializeRequest) SetIdFaceQuality(v string) *InitializeRequest {
	s.IdFaceQuality = &v
	return s
}

func (s *InitializeRequest) SetIdSpoof(v string) *InitializeRequest {
	s.IdSpoof = &v
	return s
}

func (s *InitializeRequest) SetIdThreshold(v string) *InitializeRequest {
	s.IdThreshold = &v
	return s
}

func (s *InitializeRequest) SetLanguageConfig(v string) *InitializeRequest {
	s.LanguageConfig = &v
	return s
}

func (s *InitializeRequest) SetMerchantBizId(v string) *InitializeRequest {
	s.MerchantBizId = &v
	return s
}

func (s *InitializeRequest) SetMerchantUserId(v string) *InitializeRequest {
	s.MerchantUserId = &v
	return s
}

func (s *InitializeRequest) SetMetaInfo(v string) *InitializeRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitializeRequest) SetModel(v string) *InitializeRequest {
	s.Model = &v
	return s
}

func (s *InitializeRequest) SetOcr(v string) *InitializeRequest {
	s.Ocr = &v
	return s
}

func (s *InitializeRequest) SetProcedurePriority(v string) *InitializeRequest {
	s.ProcedurePriority = &v
	return s
}

func (s *InitializeRequest) SetProductCode(v string) *InitializeRequest {
	s.ProductCode = &v
	return s
}

func (s *InitializeRequest) SetProductFlow(v string) *InitializeRequest {
	s.ProductFlow = &v
	return s
}

func (s *InitializeRequest) SetReturnUrl(v string) *InitializeRequest {
	s.ReturnUrl = &v
	return s
}

func (s *InitializeRequest) SetSceneCode(v string) *InitializeRequest {
	s.SceneCode = &v
	return s
}

func (s *InitializeRequest) SetSecurityLevel(v string) *InitializeRequest {
	s.SecurityLevel = &v
	return s
}

func (s *InitializeRequest) SetShowAlbumIcon(v string) *InitializeRequest {
	s.ShowAlbumIcon = &v
	return s
}

func (s *InitializeRequest) SetShowGuidePage(v string) *InitializeRequest {
	s.ShowGuidePage = &v
	return s
}

func (s *InitializeRequest) SetShowOcrResult(v string) *InitializeRequest {
	s.ShowOcrResult = &v
	return s
}

func (s *InitializeRequest) SetStyleConfig(v string) *InitializeRequest {
	s.StyleConfig = &v
	return s
}

type InitializeResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *InitializeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s InitializeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeResponseBody) SetCode(v string) *InitializeResponseBody {
	s.Code = &v
	return s
}

func (s *InitializeResponseBody) SetMessage(v string) *InitializeResponseBody {
	s.Message = &v
	return s
}

func (s *InitializeResponseBody) SetRequestId(v string) *InitializeResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeResponseBody) SetResult(v *InitializeResponseBodyResult) *InitializeResponseBody {
	s.Result = v
	return s
}

type InitializeResponseBodyResult struct {
	// example:
	//
	// ***
	ClientCfg *string `json:"ClientCfg,omitempty" xml:"ClientCfg,omitempty"`
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
	// example:
	//
	// http****
	TransactionUrl *string `json:"TransactionUrl,omitempty" xml:"TransactionUrl,omitempty"`
}

func (s InitializeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s InitializeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InitializeResponseBodyResult) SetClientCfg(v string) *InitializeResponseBodyResult {
	s.ClientCfg = &v
	return s
}

func (s *InitializeResponseBodyResult) SetTransactionId(v string) *InitializeResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *InitializeResponseBodyResult) SetTransactionUrl(v string) *InitializeResponseBodyResult {
	s.TransactionUrl = &v
	return s
}

type InitializeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeResponse) String() string {
	return tea.Prettify(s)
}

func (s InitializeResponse) GoString() string {
	return s.String()
}

func (s *InitializeResponse) SetHeaders(v map[string]*string) *InitializeResponse {
	s.Headers = v
	return s
}

func (s *InitializeResponse) SetStatusCode(v int32) *InitializeResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeResponse) SetBody(v *InitializeResponseBody) *InitializeResponse {
	s.Body = v
	return s
}

type Mobile3MetaVerifyIntlRequest struct {
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// example:
	//
	// 186****1234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// example:
	//
	// MOBILE_3META
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Mobile3MetaVerifyIntlRequest) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *Mobile3MetaVerifyIntlRequest) SetIdentifyNum(v string) *Mobile3MetaVerifyIntlRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Mobile3MetaVerifyIntlRequest) SetMobile(v string) *Mobile3MetaVerifyIntlRequest {
	s.Mobile = &v
	return s
}

func (s *Mobile3MetaVerifyIntlRequest) SetParamType(v string) *Mobile3MetaVerifyIntlRequest {
	s.ParamType = &v
	return s
}

func (s *Mobile3MetaVerifyIntlRequest) SetProductCode(v string) *Mobile3MetaVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *Mobile3MetaVerifyIntlRequest) SetUserName(v string) *Mobile3MetaVerifyIntlRequest {
	s.UserName = &v
	return s
}

type Mobile3MetaVerifyIntlResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D241532C-4EE9-5A2A-A5A5-C1FD98CE2EDD
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *Mobile3MetaVerifyIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s Mobile3MetaVerifyIntlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *Mobile3MetaVerifyIntlResponseBody) SetCode(v string) *Mobile3MetaVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBody) SetMessage(v string) *Mobile3MetaVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBody) SetRequestId(v string) *Mobile3MetaVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBody) SetResult(v *Mobile3MetaVerifyIntlResponseBodyResult) *Mobile3MetaVerifyIntlResponseBody {
	s.Result = v
	return s
}

type Mobile3MetaVerifyIntlResponseBodyResult struct {
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s Mobile3MetaVerifyIntlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaVerifyIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *Mobile3MetaVerifyIntlResponseBodyResult) SetBizCode(v string) *Mobile3MetaVerifyIntlResponseBodyResult {
	s.BizCode = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBodyResult) SetIspName(v string) *Mobile3MetaVerifyIntlResponseBodyResult {
	s.IspName = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBodyResult) SetSubCode(v string) *Mobile3MetaVerifyIntlResponseBodyResult {
	s.SubCode = &v
	return s
}

type Mobile3MetaVerifyIntlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Mobile3MetaVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Mobile3MetaVerifyIntlResponse) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *Mobile3MetaVerifyIntlResponse) SetHeaders(v map[string]*string) *Mobile3MetaVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *Mobile3MetaVerifyIntlResponse) SetStatusCode(v int32) *Mobile3MetaVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponse) SetBody(v *Mobile3MetaVerifyIntlResponseBody) *Mobile3MetaVerifyIntlResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudauth-intl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Deprecated: OpenAPI CardOcr is deprecated, please use Cloudauth-intl::2022-08-09::DocOcr instead.
//
// Summary:
//
// 证件OCR识别纯服务端接口
//
// @param request - CardOcrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CardOcrResponse
// Deprecated
func (client *Client) CardOcrWithOptions(request *CardOcrRequest, runtime *util.RuntimeOptions) (_result *CardOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		query["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.IdFaceQuality)) {
		query["IdFaceQuality"] = request.IdFaceQuality
	}

	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureUrl)) {
		query["IdOcrPictureUrl"] = request.IdOcrPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Ocr)) {
		query["Ocr"] = request.Ocr
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.Spoof)) {
		query["Spoof"] = request.Spoof
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureBase64)) {
		body["IdOcrPictureBase64"] = request.IdOcrPictureBase64
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CardOcr"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CardOcrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CardOcr is deprecated, please use Cloudauth-intl::2022-08-09::DocOcr instead.
//
// Summary:
//
// 证件OCR识别纯服务端接口
//
// @param request - CardOcrRequest
//
// @return CardOcrResponse
// Deprecated
func (client *Client) CardOcr(request *CardOcrRequest) (_result *CardOcrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CardOcrResponse{}
	_body, _err := client.CardOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 结果查询
//
// @param request - CheckResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckResultResponse
func (client *Client) CheckResultWithOptions(request *CheckResultRequest, runtime *util.RuntimeOptions) (_result *CheckResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraImageControlList)) {
		query["ExtraImageControlList"] = request.ExtraImageControlList
	}

	if !tea.BoolValue(util.IsUnset(request.IsReturnImage)) {
		query["IsReturnImage"] = request.IsReturnImage
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnFiveCategorySpoofResult)) {
		query["ReturnFiveCategorySpoofResult"] = request.ReturnFiveCategorySpoofResult
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		query["TransactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckResult"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 结果查询
//
// @param request - CheckResultRequest
//
// @return CheckResultResponse
func (client *Client) CheckResult(request *CheckResultRequest) (_result *CheckResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckResultResponse{}
	_body, _err := client.CheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 认证日志查询接口
//
// @param request - CheckVerifyLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckVerifyLogResponse
func (client *Client) CheckVerifyLogWithOptions(request *CheckVerifyLogRequest, runtime *util.RuntimeOptions) (_result *CheckVerifyLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		body["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		body["TransactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckVerifyLog"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckVerifyLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 认证日志查询接口
//
// @param request - CheckVerifyLogRequest
//
// @return CheckVerifyLogResponse
func (client *Client) CheckVerifyLog(request *CheckVerifyLogRequest) (_result *CheckVerifyLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckVerifyLogResponse{}
	_body, _err := client.CheckVerifyLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人脸凭证核验
//
// @param request - DeepfakeDetectIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeepfakeDetectIntlResponse
func (client *Client) DeepfakeDetectIntlWithOptions(request *DeepfakeDetectIntlRequest, runtime *util.RuntimeOptions) (_result *DeepfakeDetectIntlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceInputType)) {
		query["FaceInputType"] = request.FaceInputType
	}

	if !tea.BoolValue(util.IsUnset(request.FaceUrl)) {
		query["FaceUrl"] = request.FaceUrl
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceBase64)) {
		body["FaceBase64"] = request.FaceBase64
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeepfakeDetectIntl"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeepfakeDetectIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人脸凭证核验
//
// @param request - DeepfakeDetectIntlRequest
//
// @return DeepfakeDetectIntlResponse
func (client *Client) DeepfakeDetectIntl(request *DeepfakeDetectIntlRequest) (_result *DeepfakeDetectIntlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeepfakeDetectIntlResponse{}
	_body, _err := client.DeepfakeDetectIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除用户认证记录结果
//
// @param request - DeleteVerifyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVerifyResultResponse
func (client *Client) DeleteVerifyResultWithOptions(request *DeleteVerifyResultRequest, runtime *util.RuntimeOptions) (_result *DeleteVerifyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeleteAfterQuery)) {
		query["DeleteAfterQuery"] = request.DeleteAfterQuery
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteType)) {
		query["DeleteType"] = request.DeleteType
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		query["TransactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVerifyResult"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVerifyResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户认证记录结果
//
// @param request - DeleteVerifyResultRequest
//
// @return DeleteVerifyResultResponse
func (client *Client) DeleteVerifyResult(request *DeleteVerifyResultRequest) (_result *DeleteVerifyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVerifyResultResponse{}
	_body, _err := client.DeleteVerifyResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 卡证ocr纯服务端
//
// @param request - DocOcrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocOcrResponse
func (client *Client) DocOcrWithOptions(request *DocOcrRequest, runtime *util.RuntimeOptions) (_result *DocOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardSide)) {
		query["CardSide"] = request.CardSide
	}

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		query["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.IdFaceQuality)) {
		query["IdFaceQuality"] = request.IdFaceQuality
	}

	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureUrl)) {
		query["IdOcrPictureUrl"] = request.IdOcrPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdThreshold)) {
		query["IdThreshold"] = request.IdThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Ocr)) {
		query["Ocr"] = request.Ocr
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.Spoof)) {
		query["Spoof"] = request.Spoof
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureBase64)) {
		body["IdOcrPictureBase64"] = request.IdOcrPictureBase64
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DocOcr"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DocOcrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 卡证ocr纯服务端
//
// @param request - DocOcrRequest
//
// @return DocOcrResponse
func (client *Client) DocOcr(request *DocOcrRequest) (_result *DocOcrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DocOcrResponse{}
	_body, _err := client.DocOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ekyc纯服务端接口
//
// @param request - EkycVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EkycVerifyResponse
func (client *Client) EkycVerifyWithOptions(request *EkycVerifyRequest, runtime *util.RuntimeOptions) (_result *EkycVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Authorize)) {
		query["Authorize"] = request.Authorize
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		query["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.DocName)) {
		query["DocName"] = request.DocName
	}

	if !tea.BoolValue(util.IsUnset(request.DocNo)) {
		query["DocNo"] = request.DocNo
	}

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		query["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.FacePictureUrl)) {
		query["FacePictureUrl"] = request.FacePictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureUrl)) {
		query["IdOcrPictureUrl"] = request.IdOcrPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdThreshold)) {
		query["IdThreshold"] = request.IdThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FacePictureBase64)) {
		body["FacePictureBase64"] = request.FacePictureBase64
	}

	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureBase64)) {
		body["IdOcrPictureBase64"] = request.IdOcrPictureBase64
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EkycVerify"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EkycVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ekyc纯服务端接口
//
// @param request - EkycVerifyRequest
//
// @return EkycVerifyResponse
func (client *Client) EkycVerify(request *EkycVerifyRequest) (_result *EkycVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EkycVerifyResponse{}
	_body, _err := client.EkycVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人脸比对
//
// @param request - FaceCompareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceCompareResponse
func (client *Client) FaceCompareWithOptions(request *FaceCompareRequest, runtime *util.RuntimeOptions) (_result *FaceCompareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFacePictureUrl)) {
		query["SourceFacePictureUrl"] = request.SourceFacePictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFacePictureUrl)) {
		query["TargetFacePictureUrl"] = request.TargetFacePictureUrl
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceFacePicture)) {
		body["SourceFacePicture"] = request.SourceFacePicture
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFacePicture)) {
		body["TargetFacePicture"] = request.TargetFacePicture
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceCompare"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceCompareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人脸比对
//
// @param request - FaceCompareRequest
//
// @return FaceCompareResponse
func (client *Client) FaceCompare(request *FaceCompareRequest) (_result *FaceCompareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceCompareResponse{}
	_body, _err := client.FaceCompareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 国际人脸保镖纯服务端接口
//
// @param request - FaceGuardRiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceGuardRiskResponse
func (client *Client) FaceGuardRiskWithOptions(request *FaceGuardRiskRequest, runtime *util.RuntimeOptions) (_result *FaceGuardRiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceToken)) {
		query["DeviceToken"] = request.DeviceToken
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceGuardRisk"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceGuardRiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际人脸保镖纯服务端接口
//
// @param request - FaceGuardRiskRequest
//
// @return FaceGuardRiskResponse
func (client *Client) FaceGuardRisk(request *FaceGuardRiskRequest) (_result *FaceGuardRiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceGuardRiskResponse{}
	_body, _err := client.FaceGuardRiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 静默活体API 纯服务端
//
// @param request - FaceLivenessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceLivenessResponse
func (client *Client) FaceLivenessWithOptions(request *FaceLivenessRequest, runtime *util.RuntimeOptions) (_result *FaceLivenessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		query["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.FacePictureUrl)) {
		query["FacePictureUrl"] = request.FacePictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FaceQuality)) {
		query["FaceQuality"] = request.FaceQuality
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Occlusion)) {
		query["Occlusion"] = request.Occlusion
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FacePictureBase64)) {
		body["FacePictureBase64"] = request.FacePictureBase64
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceLiveness"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceLivenessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 静默活体API 纯服务端
//
// @param request - FaceLivenessRequest
//
// @return FaceLivenessResponse
func (client *Client) FaceLiveness(request *FaceLivenessRequest) (_result *FaceLivenessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceLivenessResponse{}
	_body, _err := client.FaceLivenessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 防伪回调接口
//
// @param request - FraudResultCallBackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FraudResultCallBackResponse
func (client *Client) FraudResultCallBackWithOptions(request *FraudResultCallBackRequest, runtime *util.RuntimeOptions) (_result *FraudResultCallBackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		query["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtParams)) {
		query["ExtParams"] = request.ExtParams
	}

	if !tea.BoolValue(util.IsUnset(request.ResultCode)) {
		query["ResultCode"] = request.ResultCode
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyDeployEnv)) {
		query["VerifyDeployEnv"] = request.VerifyDeployEnv
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FraudResultCallBack"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FraudResultCallBackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 防伪回调接口
//
// @param request - FraudResultCallBackRequest
//
// @return FraudResultCallBackResponse
func (client *Client) FraudResultCallBack(request *FraudResultCallBackRequest) (_result *FraudResultCallBackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FraudResultCallBackResponse{}
	_body, _err := client.FraudResultCallBackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 身份二要素有效期核验
//
// @param request - Id2MetaPeriodVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaPeriodVerifyIntlResponse
func (client *Client) Id2MetaPeriodVerifyIntlWithOptions(request *Id2MetaPeriodVerifyIntlRequest, runtime *util.RuntimeOptions) (_result *Id2MetaPeriodVerifyIntlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocName)) {
		body["DocName"] = request.DocName
	}

	if !tea.BoolValue(util.IsUnset(request.DocNo)) {
		body["DocNo"] = request.DocNo
	}

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		body["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		body["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		body["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		body["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.ValidityEndDate)) {
		body["ValidityEndDate"] = request.ValidityEndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ValidityStartDate)) {
		body["ValidityStartDate"] = request.ValidityStartDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Id2MetaPeriodVerifyIntl"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Id2MetaPeriodVerifyIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 身份二要素有效期核验
//
// @param request - Id2MetaPeriodVerifyIntlRequest
//
// @return Id2MetaPeriodVerifyIntlResponse
func (client *Client) Id2MetaPeriodVerifyIntl(request *Id2MetaPeriodVerifyIntlRequest) (_result *Id2MetaPeriodVerifyIntlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Id2MetaPeriodVerifyIntlResponse{}
	_body, _err := client.Id2MetaPeriodVerifyIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 身份二要素国际版接口
//
// @param request - Id2MetaVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaVerifyIntlResponse
func (client *Client) Id2MetaVerifyIntlWithOptions(request *Id2MetaVerifyIntlRequest, runtime *util.RuntimeOptions) (_result *Id2MetaVerifyIntlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentifyNum)) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		query["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Id2MetaVerifyIntl"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Id2MetaVerifyIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 身份二要素国际版接口
//
// @param request - Id2MetaVerifyIntlRequest
//
// @return Id2MetaVerifyIntlResponse
func (client *Client) Id2MetaVerifyIntl(request *Id2MetaVerifyIntlRequest) (_result *Id2MetaVerifyIntlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Id2MetaVerifyIntlResponse{}
	_body, _err := client.Id2MetaVerifyIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 认证初始化
//
// @param request - InitializeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeResponse
func (client *Client) InitializeWithOptions(request *InitializeRequest, runtime *util.RuntimeOptions) (_result *InitializeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppQualityCheck)) {
		query["AppQualityCheck"] = request.AppQualityCheck
	}

	if !tea.BoolValue(util.IsUnset(request.Authorize)) {
		query["Authorize"] = request.Authorize
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackToken)) {
		query["CallbackToken"] = request.CallbackToken
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		query["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.DocScanMode)) {
		query["DocScanMode"] = request.DocScanMode
	}

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		query["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.DocVideo)) {
		query["DocVideo"] = request.DocVideo
	}

	if !tea.BoolValue(util.IsUnset(request.ExperienceCode)) {
		query["ExperienceCode"] = request.ExperienceCode
	}

	if !tea.BoolValue(util.IsUnset(request.FacePictureUrl)) {
		query["FacePictureUrl"] = request.FacePictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdFaceQuality)) {
		query["IdFaceQuality"] = request.IdFaceQuality
	}

	if !tea.BoolValue(util.IsUnset(request.IdSpoof)) {
		query["IdSpoof"] = request.IdSpoof
	}

	if !tea.BoolValue(util.IsUnset(request.IdThreshold)) {
		query["IdThreshold"] = request.IdThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageConfig)) {
		query["LanguageConfig"] = request.LanguageConfig
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MetaInfo)) {
		query["MetaInfo"] = request.MetaInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		query["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.Ocr)) {
		query["Ocr"] = request.Ocr
	}

	if !tea.BoolValue(util.IsUnset(request.ProcedurePriority)) {
		query["ProcedurePriority"] = request.ProcedurePriority
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductFlow)) {
		query["ProductFlow"] = request.ProductFlow
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnUrl)) {
		query["ReturnUrl"] = request.ReturnUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityLevel)) {
		query["SecurityLevel"] = request.SecurityLevel
	}

	if !tea.BoolValue(util.IsUnset(request.ShowAlbumIcon)) {
		query["ShowAlbumIcon"] = request.ShowAlbumIcon
	}

	if !tea.BoolValue(util.IsUnset(request.ShowGuidePage)) {
		query["ShowGuidePage"] = request.ShowGuidePage
	}

	if !tea.BoolValue(util.IsUnset(request.ShowOcrResult)) {
		query["ShowOcrResult"] = request.ShowOcrResult
	}

	if !tea.BoolValue(util.IsUnset(request.StyleConfig)) {
		query["StyleConfig"] = request.StyleConfig
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FacePictureBase64)) {
		body["FacePictureBase64"] = request.FacePictureBase64
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Initialize"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitializeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 认证初始化
//
// @param request - InitializeRequest
//
// @return InitializeResponse
func (client *Client) Initialize(request *InitializeRequest) (_result *InitializeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitializeResponse{}
	_body, _err := client.InitializeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 手机号三要素国际版接口
//
// @param request - Mobile3MetaVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile3MetaVerifyIntlResponse
func (client *Client) Mobile3MetaVerifyIntlWithOptions(request *Mobile3MetaVerifyIntlRequest, runtime *util.RuntimeOptions) (_result *Mobile3MetaVerifyIntlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentifyNum)) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		query["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Mobile3MetaVerifyIntl"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Mobile3MetaVerifyIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 手机号三要素国际版接口
//
// @param request - Mobile3MetaVerifyIntlRequest
//
// @return Mobile3MetaVerifyIntlResponse
func (client *Client) Mobile3MetaVerifyIntl(request *Mobile3MetaVerifyIntlRequest) (_result *Mobile3MetaVerifyIntlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Mobile3MetaVerifyIntlResponse{}
	_body, _err := client.Mobile3MetaVerifyIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
