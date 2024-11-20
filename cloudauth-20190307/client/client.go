// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type AIGCFaceVerifyRequest struct {
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceContrastPicture *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// example:
	//
	// LR_FR_AIGC
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 100000xxxx
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s AIGCFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s AIGCFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *AIGCFaceVerifyRequest) SetFaceContrastPicture(v string) *AIGCFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *AIGCFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *AIGCFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *AIGCFaceVerifyRequest) SetOssBucketName(v string) *AIGCFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *AIGCFaceVerifyRequest) SetOssObjectName(v string) *AIGCFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *AIGCFaceVerifyRequest) SetOuterOrderNo(v string) *AIGCFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *AIGCFaceVerifyRequest) SetProductCode(v string) *AIGCFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *AIGCFaceVerifyRequest) SetSceneId(v int64) *AIGCFaceVerifyRequest {
	s.SceneId = &v
	return s
}

type AIGCFaceVerifyResponseBody struct {
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *AIGCFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s AIGCFaceVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AIGCFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *AIGCFaceVerifyResponseBody) SetCode(v string) *AIGCFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *AIGCFaceVerifyResponseBody) SetMessage(v string) *AIGCFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *AIGCFaceVerifyResponseBody) SetRequestId(v string) *AIGCFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AIGCFaceVerifyResponseBody) SetResultObject(v *AIGCFaceVerifyResponseBodyResultObject) *AIGCFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

type AIGCFaceVerifyResponseBodyResultObject struct {
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// Y
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 1.0000
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s AIGCFaceVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s AIGCFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *AIGCFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *AIGCFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *AIGCFaceVerifyResponseBodyResultObject) SetResult(v string) *AIGCFaceVerifyResponseBodyResultObject {
	s.Result = &v
	return s
}

func (s *AIGCFaceVerifyResponseBodyResultObject) SetScore(v string) *AIGCFaceVerifyResponseBodyResultObject {
	s.Score = &v
	return s
}

type AIGCFaceVerifyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AIGCFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AIGCFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s AIGCFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *AIGCFaceVerifyResponse) SetHeaders(v map[string]*string) *AIGCFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *AIGCFaceVerifyResponse) SetStatusCode(v int32) *AIGCFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *AIGCFaceVerifyResponse) SetBody(v *AIGCFaceVerifyResponseBody) *AIGCFaceVerifyResponse {
	s.Body = v
	return s
}

type BankMetaVerifyRequest struct {
	// example:
	//
	// 610*************1181
	BankCard *string `json:"BankCard,omitempty" xml:"BankCard,omitempty"`
	// example:
	//
	// 429001********8211
	IdentifyNum  *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// example:
	//
	// 138******11
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// example:
	//
	// BANK_CARD_2_META
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// VERIFY_BANK_CARD
	VerifyMode *string `json:"VerifyMode,omitempty" xml:"VerifyMode,omitempty"`
}

func (s BankMetaVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s BankMetaVerifyRequest) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyRequest) SetBankCard(v string) *BankMetaVerifyRequest {
	s.BankCard = &v
	return s
}

func (s *BankMetaVerifyRequest) SetIdentifyNum(v string) *BankMetaVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *BankMetaVerifyRequest) SetIdentityType(v string) *BankMetaVerifyRequest {
	s.IdentityType = &v
	return s
}

func (s *BankMetaVerifyRequest) SetMobile(v string) *BankMetaVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *BankMetaVerifyRequest) SetParamType(v string) *BankMetaVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *BankMetaVerifyRequest) SetProductType(v string) *BankMetaVerifyRequest {
	s.ProductType = &v
	return s
}

func (s *BankMetaVerifyRequest) SetUserName(v string) *BankMetaVerifyRequest {
	s.UserName = &v
	return s
}

func (s *BankMetaVerifyRequest) SetVerifyMode(v string) *BankMetaVerifyRequest {
	s.VerifyMode = &v
	return s
}

type BankMetaVerifyResponseBody struct {
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
	// 473469C7-A***B-A3DC0DE3C83E
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *BankMetaVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s BankMetaVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BankMetaVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyResponseBody) SetCode(v string) *BankMetaVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *BankMetaVerifyResponseBody) SetMessage(v string) *BankMetaVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *BankMetaVerifyResponseBody) SetRequestId(v string) *BankMetaVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *BankMetaVerifyResponseBody) SetResultObject(v *BankMetaVerifyResponseBodyResultObject) *BankMetaVerifyResponseBody {
	s.ResultObject = v
	return s
}

type BankMetaVerifyResponseBodyResultObject struct {
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s BankMetaVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s BankMetaVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyResponseBodyResultObject) SetBizCode(v string) *BankMetaVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *BankMetaVerifyResponseBodyResultObject) SetSubCode(v string) *BankMetaVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

type BankMetaVerifyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BankMetaVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BankMetaVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s BankMetaVerifyResponse) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyResponse) SetHeaders(v map[string]*string) *BankMetaVerifyResponse {
	s.Headers = v
	return s
}

func (s *BankMetaVerifyResponse) SetStatusCode(v int32) *BankMetaVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *BankMetaVerifyResponse) SetBody(v *BankMetaVerifyResponseBody) *BankMetaVerifyResponse {
	s.Body = v
	return s
}

type CompareFaceVerifyRequest struct {
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// example:
	//
	// PV_FC
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 1000000006
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 0bfa7c493f850e5178b9f8613634c9xx
	SourceCertifyId *string `json:"SourceCertifyId,omitempty" xml:"SourceCertifyId,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	SourceFaceContrastPicture *string `json:"SourceFaceContrastPicture,omitempty" xml:"SourceFaceContrastPicture,omitempty"`
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	SourceFaceContrastPictureUrl *string `json:"SourceFaceContrastPictureUrl,omitempty" xml:"SourceFaceContrastPictureUrl,omitempty"`
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	SourceOssBucketName *string `json:"SourceOssBucketName,omitempty" xml:"SourceOssBucketName,omitempty"`
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	SourceOssObjectName *string `json:"SourceOssObjectName,omitempty" xml:"SourceOssObjectName,omitempty"`
	// example:
	//
	// 0bfa7c493f850e5178b9f8613634c9xx
	TargetCertifyId *string `json:"TargetCertifyId,omitempty" xml:"TargetCertifyId,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	TargetFaceContrastPicture *string `json:"TargetFaceContrastPicture,omitempty" xml:"TargetFaceContrastPicture,omitempty"`
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	TargetFaceContrastPictureUrl *string `json:"TargetFaceContrastPictureUrl,omitempty" xml:"TargetFaceContrastPictureUrl,omitempty"`
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	TargetOssBucketName *string `json:"TargetOssBucketName,omitempty" xml:"TargetOssBucketName,omitempty"`
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	TargetOssObjectName *string `json:"TargetOssObjectName,omitempty" xml:"TargetOssObjectName,omitempty"`
}

func (s CompareFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyRequest) SetCrop(v string) *CompareFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetOuterOrderNo(v string) *CompareFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetProductCode(v string) *CompareFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSceneId(v int64) *CompareFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceCertifyId(v string) *CompareFaceVerifyRequest {
	s.SourceCertifyId = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceFaceContrastPicture(v string) *CompareFaceVerifyRequest {
	s.SourceFaceContrastPicture = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceFaceContrastPictureUrl(v string) *CompareFaceVerifyRequest {
	s.SourceFaceContrastPictureUrl = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceOssBucketName(v string) *CompareFaceVerifyRequest {
	s.SourceOssBucketName = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceOssObjectName(v string) *CompareFaceVerifyRequest {
	s.SourceOssObjectName = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetCertifyId(v string) *CompareFaceVerifyRequest {
	s.TargetCertifyId = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetFaceContrastPicture(v string) *CompareFaceVerifyRequest {
	s.TargetFaceContrastPicture = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetFaceContrastPictureUrl(v string) *CompareFaceVerifyRequest {
	s.TargetFaceContrastPictureUrl = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetOssBucketName(v string) *CompareFaceVerifyRequest {
	s.TargetOssBucketName = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetOssObjectName(v string) *CompareFaceVerifyRequest {
	s.TargetOssObjectName = &v
	return s
}

type CompareFaceVerifyResponseBody struct {
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
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *CompareFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CompareFaceVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyResponseBody) SetCode(v string) *CompareFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *CompareFaceVerifyResponseBody) SetMessage(v string) *CompareFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *CompareFaceVerifyResponseBody) SetRequestId(v string) *CompareFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareFaceVerifyResponseBody) SetResultObject(v *CompareFaceVerifyResponseBodyResultObject) *CompareFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

type CompareFaceVerifyResponseBodyResultObject struct {
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// T
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 99.60875
	VerifyScore *float32 `json:"VerifyScore,omitempty" xml:"VerifyScore,omitempty"`
}

func (s CompareFaceVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *CompareFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *CompareFaceVerifyResponseBodyResultObject) SetPassed(v string) *CompareFaceVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *CompareFaceVerifyResponseBodyResultObject) SetVerifyScore(v float32) *CompareFaceVerifyResponseBodyResultObject {
	s.VerifyScore = &v
	return s
}

type CompareFaceVerifyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompareFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompareFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyResponse) SetHeaders(v map[string]*string) *CompareFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *CompareFaceVerifyResponse) SetStatusCode(v int32) *CompareFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareFaceVerifyResponse) SetBody(v *CompareFaceVerifyResponseBody) *CompareFaceVerifyResponse {
	s.Body = v
	return s
}

type CompareFacesRequest struct {
	// example:
	//
	// FacePic
	SourceImageType *string `json:"SourceImageType,omitempty" xml:"SourceImageType,omitempty"`
	// example:
	//
	// http%3A%2F%2Fjiangsu.china.com.cn%2Fuploadfile%2F2015%2F0114%2F1421221304095989.jpg
	SourceImageValue *string `json:"SourceImageValue,omitempty" xml:"SourceImageValue,omitempty"`
	// example:
	//
	// FacePic
	TargetImageType *string `json:"TargetImageType,omitempty" xml:"TargetImageType,omitempty"`
	// example:
	//
	// http%3A%2F%2Fjiangsu.china.com.cn%2Fuploadfile%2F2015%2F0114%2F1421221304095989.jpg
	TargetImageValue *string `json:"TargetImageValue,omitempty" xml:"TargetImageValue,omitempty"`
}

func (s CompareFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesRequest) GoString() string {
	return s.String()
}

func (s *CompareFacesRequest) SetSourceImageType(v string) *CompareFacesRequest {
	s.SourceImageType = &v
	return s
}

func (s *CompareFacesRequest) SetSourceImageValue(v string) *CompareFacesRequest {
	s.SourceImageValue = &v
	return s
}

func (s *CompareFacesRequest) SetTargetImageType(v string) *CompareFacesRequest {
	s.TargetImageType = &v
	return s
}

func (s *CompareFacesRequest) SetTargetImageValue(v string) *CompareFacesRequest {
	s.TargetImageValue = &v
	return s
}

type CompareFacesResponseBody struct {
	// example:
	//
	// 200
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CompareFacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Error.InternalError
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CompareFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesResponseBody) GoString() string {
	return s.String()
}

func (s *CompareFacesResponseBody) SetCode(v string) *CompareFacesResponseBody {
	s.Code = &v
	return s
}

func (s *CompareFacesResponseBody) SetData(v *CompareFacesResponseBodyData) *CompareFacesResponseBody {
	s.Data = v
	return s
}

func (s *CompareFacesResponseBody) SetMessage(v string) *CompareFacesResponseBody {
	s.Message = &v
	return s
}

func (s *CompareFacesResponseBody) SetRequestId(v string) *CompareFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareFacesResponseBody) SetSuccess(v bool) *CompareFacesResponseBody {
	s.Success = &v
	return s
}

type CompareFacesResponseBodyData struct {
	// example:
	//
	// {"0.0001":"90.07","0.001":"80.01","0.01":"70.02"}
	ConfidenceThresholds *string `json:"ConfidenceThresholds,omitempty" xml:"ConfidenceThresholds,omitempty"`
	// example:
	//
	// 98.7913
	SimilarityScore *float32 `json:"SimilarityScore,omitempty" xml:"SimilarityScore,omitempty"`
}

func (s CompareFacesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompareFacesResponseBodyData) SetConfidenceThresholds(v string) *CompareFacesResponseBodyData {
	s.ConfidenceThresholds = &v
	return s
}

func (s *CompareFacesResponseBodyData) SetSimilarityScore(v float32) *CompareFacesResponseBodyData {
	s.SimilarityScore = &v
	return s
}

type CompareFacesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompareFacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompareFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesResponse) GoString() string {
	return s.String()
}

func (s *CompareFacesResponse) SetHeaders(v map[string]*string) *CompareFacesResponse {
	s.Headers = v
	return s
}

func (s *CompareFacesResponse) SetStatusCode(v int32) *CompareFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareFacesResponse) SetBody(v *CompareFacesResponseBody) *CompareFacesResponse {
	s.Body = v
	return s
}

type ContrastFaceVerifyRequest struct {
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// example:
	//
	// IDENTITY_CARD
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// example:
	//
	// 0bfa7c493f850e5178b9f8613634c9xx
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Crop      *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// example:
	//
	// McozS1ZWRcRZStlERcZZo_QOytx5jcgZoZJEoRLOxxxxxxx
	DeviceToken      *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	EncryptType      *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	FaceContrastFile *string `json:"FaceContrastFile,omitempty" xml:"FaceContrastFile,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceContrastPicture *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	// example:
	//
	// 114.xxx.xxx.xxx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 130xxxxxxxx
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Model  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// example:
	//
	// ID_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 1000000006
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ContrastFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyRequest) SetCertName(v string) *ContrastFaceVerifyRequest {
	s.CertName = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertNo(v string) *ContrastFaceVerifyRequest {
	s.CertNo = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertType(v string) *ContrastFaceVerifyRequest {
	s.CertType = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertifyId(v string) *ContrastFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCrop(v string) *ContrastFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetDeviceToken(v string) *ContrastFaceVerifyRequest {
	s.DeviceToken = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetEncryptType(v string) *ContrastFaceVerifyRequest {
	s.EncryptType = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastFile(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastFile = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastPicture(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetIp(v string) *ContrastFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetMobile(v string) *ContrastFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetModel(v string) *ContrastFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetOssBucketName(v string) *ContrastFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetOssObjectName(v string) *ContrastFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetOuterOrderNo(v string) *ContrastFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetProductCode(v string) *ContrastFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetSceneId(v int64) *ContrastFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetUserId(v string) *ContrastFaceVerifyRequest {
	s.UserId = &v
	return s
}

type ContrastFaceVerifyAdvanceRequest struct {
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// example:
	//
	// IDENTITY_CARD
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// example:
	//
	// 0bfa7c493f850e5178b9f8613634c9xx
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Crop      *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// example:
	//
	// McozS1ZWRcRZStlERcZZo_QOytx5jcgZoZJEoRLOxxxxxxx
	DeviceToken            *string   `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	EncryptType            *string   `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	FaceContrastFileObject io.Reader `json:"FaceContrastFile,omitempty" xml:"FaceContrastFile,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceContrastPicture *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	// example:
	//
	// 114.xxx.xxx.xxx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 130xxxxxxxx
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Model  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// example:
	//
	// ID_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 1000000006
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ContrastFaceVerifyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertName(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertName = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertNo(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertNo = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertType(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertType = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertifyId(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertifyId = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCrop(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Crop = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetDeviceToken(v string) *ContrastFaceVerifyAdvanceRequest {
	s.DeviceToken = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetEncryptType(v string) *ContrastFaceVerifyAdvanceRequest {
	s.EncryptType = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastFileObject(v io.Reader) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastFileObject = v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastPicture(v string) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastPictureUrl(v string) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetIp(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Ip = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetMobile(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Mobile = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetModel(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Model = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetOssBucketName(v string) *ContrastFaceVerifyAdvanceRequest {
	s.OssBucketName = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetOssObjectName(v string) *ContrastFaceVerifyAdvanceRequest {
	s.OssObjectName = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetOuterOrderNo(v string) *ContrastFaceVerifyAdvanceRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetProductCode(v string) *ContrastFaceVerifyAdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetSceneId(v int64) *ContrastFaceVerifyAdvanceRequest {
	s.SceneId = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetUserId(v string) *ContrastFaceVerifyAdvanceRequest {
	s.UserId = &v
	return s
}

type ContrastFaceVerifyResponseBody struct {
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
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *ContrastFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s ContrastFaceVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyResponseBody) SetCode(v string) *ContrastFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *ContrastFaceVerifyResponseBody) SetMessage(v string) *ContrastFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *ContrastFaceVerifyResponseBody) SetRequestId(v string) *ContrastFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContrastFaceVerifyResponseBody) SetResultObject(v *ContrastFaceVerifyResponseBodyResultObject) *ContrastFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

type ContrastFaceVerifyResponseBodyResultObject struct {
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// null
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty"`
	// example:
	//
	// {"faceAttack": "F","facialPictureFront": {"qualityScore": 88.3615493774414,"verifyScore": 50.28594166529785}}
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	// example:
	//
	// T
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s ContrastFaceVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetIdentityInfo(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.IdentityInfo = &v
	return s
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetMaterialInfo(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetPassed(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetSubCode(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

type ContrastFaceVerifyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContrastFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContrastFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyResponse) SetHeaders(v map[string]*string) *ContrastFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *ContrastFaceVerifyResponse) SetStatusCode(v int32) *ContrastFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *ContrastFaceVerifyResponse) SetBody(v *ContrastFaceVerifyResponseBody) *ContrastFaceVerifyResponse {
	s.Body = v
	return s
}

type CreateAuthKeyRequest struct {
	// example:
	//
	// 1
	AuthYears *int32 `json:"AuthYears,omitempty" xml:"AuthYears,omitempty"`
	// example:
	//
	// FACE_TEST
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// false
	Test *bool `json:"Test,omitempty" xml:"Test,omitempty"`
	// example:
	//
	// 3iJ1AY$oHcu7mC69
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s CreateAuthKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthKeyRequest) SetAuthYears(v int32) *CreateAuthKeyRequest {
	s.AuthYears = &v
	return s
}

func (s *CreateAuthKeyRequest) SetBizType(v string) *CreateAuthKeyRequest {
	s.BizType = &v
	return s
}

func (s *CreateAuthKeyRequest) SetTest(v bool) *CreateAuthKeyRequest {
	s.Test = &v
	return s
}

func (s *CreateAuthKeyRequest) SetUserDeviceId(v string) *CreateAuthKeyRequest {
	s.UserDeviceId = &v
	return s
}

type CreateAuthKeyResponseBody struct {
	// example:
	//
	// auth.1KQMcnLd4m37LN2D0F0WCD-1qtQI$
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAuthKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuthKeyResponseBody) SetAuthKey(v string) *CreateAuthKeyResponseBody {
	s.AuthKey = &v
	return s
}

func (s *CreateAuthKeyResponseBody) SetRequestId(v string) *CreateAuthKeyResponseBody {
	s.RequestId = &v
	return s
}

type CreateAuthKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAuthKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAuthKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateAuthKeyResponse) SetHeaders(v map[string]*string) *CreateAuthKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateAuthKeyResponse) SetStatusCode(v int32) *CreateAuthKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAuthKeyResponse) SetBody(v *CreateAuthKeyResponseBody) *CreateAuthKeyResponse {
	s.Body = v
	return s
}

type CreateVerifySettingRequest struct {
	// This parameter is required.
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UserRegister
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// false
	GuideStep *bool `json:"GuideStep,omitempty" xml:"GuideStep,omitempty"`
	// example:
	//
	// true
	PrivacyStep *bool `json:"PrivacyStep,omitempty" xml:"PrivacyStep,omitempty"`
	// example:
	//
	// false
	ResultStep *bool `json:"ResultStep,omitempty" xml:"ResultStep,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RPBasic
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
}

func (s CreateVerifySettingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySettingRequest) GoString() string {
	return s.String()
}

func (s *CreateVerifySettingRequest) SetBizName(v string) *CreateVerifySettingRequest {
	s.BizName = &v
	return s
}

func (s *CreateVerifySettingRequest) SetBizType(v string) *CreateVerifySettingRequest {
	s.BizType = &v
	return s
}

func (s *CreateVerifySettingRequest) SetGuideStep(v bool) *CreateVerifySettingRequest {
	s.GuideStep = &v
	return s
}

func (s *CreateVerifySettingRequest) SetPrivacyStep(v bool) *CreateVerifySettingRequest {
	s.PrivacyStep = &v
	return s
}

func (s *CreateVerifySettingRequest) SetResultStep(v bool) *CreateVerifySettingRequest {
	s.ResultStep = &v
	return s
}

func (s *CreateVerifySettingRequest) SetSolution(v string) *CreateVerifySettingRequest {
	s.Solution = &v
	return s
}

type CreateVerifySettingResponseBody struct {
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// example:
	//
	// UserRegister
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// D6163397-15C5-419C-9ACC-B7C83E0B4C10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// RPBasic
	Solution *string   `json:"Solution,omitempty" xml:"Solution,omitempty"`
	StepList []*string `json:"StepList,omitempty" xml:"StepList,omitempty" type:"Repeated"`
}

func (s CreateVerifySettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySettingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVerifySettingResponseBody) SetBizName(v string) *CreateVerifySettingResponseBody {
	s.BizName = &v
	return s
}

func (s *CreateVerifySettingResponseBody) SetBizType(v string) *CreateVerifySettingResponseBody {
	s.BizType = &v
	return s
}

func (s *CreateVerifySettingResponseBody) SetRequestId(v string) *CreateVerifySettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVerifySettingResponseBody) SetSolution(v string) *CreateVerifySettingResponseBody {
	s.Solution = &v
	return s
}

func (s *CreateVerifySettingResponseBody) SetStepList(v []*string) *CreateVerifySettingResponseBody {
	s.StepList = v
	return s
}

type CreateVerifySettingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVerifySettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVerifySettingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySettingResponse) GoString() string {
	return s.String()
}

func (s *CreateVerifySettingResponse) SetHeaders(v map[string]*string) *CreateVerifySettingResponse {
	s.Headers = v
	return s
}

func (s *CreateVerifySettingResponse) SetStatusCode(v int32) *CreateVerifySettingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVerifySettingResponse) SetBody(v *CreateVerifySettingResponseBody) *CreateVerifySettingResponse {
	s.Body = v
	return s
}

type CredentialVerifyRequest struct {
	// example:
	//
	// 4601*****
	CertNum *string `json:"CertNum,omitempty" xml:"CertNum,omitempty"`
	// example:
	//
	// 0104
	CredName *string `json:"CredName,omitempty" xml:"CredName,omitempty"`
	// example:
	//
	// 01
	CredType *string `json:"CredType,omitempty" xml:"CredType,omitempty"`
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// example:
	//
	// base64
	ImageContext *string `json:"ImageContext,omitempty" xml:"ImageContext,omitempty"`
	// example:
	//
	// http://marry.momocdn.com/avatar/3B/B6/3BB6527E-7467-926E-1048-B43614F20CC420230803_L.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// 0
	IsCheck *string `json:"IsCheck,omitempty" xml:"IsCheck,omitempty"`
	// example:
	//
	// 1
	IsOCR          *string                                  `json:"IsOCR,omitempty" xml:"IsOCR,omitempty"`
	MerchantDetail []*CredentialVerifyRequestMerchantDetail `json:"MerchantDetail,omitempty" xml:"MerchantDetail,omitempty" type:"Repeated"`
	MerchantId     *string                                  `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
	ProductCode    *string                                  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Prompt         *string                                  `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	PromptModel    *string                                  `json:"PromptModel,omitempty" xml:"PromptModel,omitempty"`
	UserName       *string                                  `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CredentialVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s CredentialVerifyRequest) GoString() string {
	return s.String()
}

func (s *CredentialVerifyRequest) SetCertNum(v string) *CredentialVerifyRequest {
	s.CertNum = &v
	return s
}

func (s *CredentialVerifyRequest) SetCredName(v string) *CredentialVerifyRequest {
	s.CredName = &v
	return s
}

func (s *CredentialVerifyRequest) SetCredType(v string) *CredentialVerifyRequest {
	s.CredType = &v
	return s
}

func (s *CredentialVerifyRequest) SetIdentifyNum(v string) *CredentialVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *CredentialVerifyRequest) SetImageContext(v string) *CredentialVerifyRequest {
	s.ImageContext = &v
	return s
}

func (s *CredentialVerifyRequest) SetImageUrl(v string) *CredentialVerifyRequest {
	s.ImageUrl = &v
	return s
}

func (s *CredentialVerifyRequest) SetIsCheck(v string) *CredentialVerifyRequest {
	s.IsCheck = &v
	return s
}

func (s *CredentialVerifyRequest) SetIsOCR(v string) *CredentialVerifyRequest {
	s.IsOCR = &v
	return s
}

func (s *CredentialVerifyRequest) SetMerchantDetail(v []*CredentialVerifyRequestMerchantDetail) *CredentialVerifyRequest {
	s.MerchantDetail = v
	return s
}

func (s *CredentialVerifyRequest) SetMerchantId(v string) *CredentialVerifyRequest {
	s.MerchantId = &v
	return s
}

func (s *CredentialVerifyRequest) SetProductCode(v string) *CredentialVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialVerifyRequest) SetPrompt(v string) *CredentialVerifyRequest {
	s.Prompt = &v
	return s
}

func (s *CredentialVerifyRequest) SetPromptModel(v string) *CredentialVerifyRequest {
	s.PromptModel = &v
	return s
}

func (s *CredentialVerifyRequest) SetUserName(v string) *CredentialVerifyRequest {
	s.UserName = &v
	return s
}

type CredentialVerifyRequestMerchantDetail struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CredentialVerifyRequestMerchantDetail) String() string {
	return tea.Prettify(s)
}

func (s CredentialVerifyRequestMerchantDetail) GoString() string {
	return s.String()
}

func (s *CredentialVerifyRequestMerchantDetail) SetKey(v string) *CredentialVerifyRequestMerchantDetail {
	s.Key = &v
	return s
}

func (s *CredentialVerifyRequestMerchantDetail) SetValue(v string) *CredentialVerifyRequestMerchantDetail {
	s.Value = &v
	return s
}

type CredentialVerifyShrinkRequest struct {
	// example:
	//
	// 4601*****
	CertNum *string `json:"CertNum,omitempty" xml:"CertNum,omitempty"`
	// example:
	//
	// 0104
	CredName *string `json:"CredName,omitempty" xml:"CredName,omitempty"`
	// example:
	//
	// 01
	CredType *string `json:"CredType,omitempty" xml:"CredType,omitempty"`
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// example:
	//
	// base64
	ImageContext *string `json:"ImageContext,omitempty" xml:"ImageContext,omitempty"`
	// example:
	//
	// http://marry.momocdn.com/avatar/3B/B6/3BB6527E-7467-926E-1048-B43614F20CC420230803_L.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// 0
	IsCheck *string `json:"IsCheck,omitempty" xml:"IsCheck,omitempty"`
	// example:
	//
	// 1
	IsOCR                *string `json:"IsOCR,omitempty" xml:"IsOCR,omitempty"`
	MerchantDetailShrink *string `json:"MerchantDetail,omitempty" xml:"MerchantDetail,omitempty"`
	MerchantId           *string `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
	ProductCode          *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Prompt               *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	PromptModel          *string `json:"PromptModel,omitempty" xml:"PromptModel,omitempty"`
	UserName             *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CredentialVerifyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CredentialVerifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CredentialVerifyShrinkRequest) SetCertNum(v string) *CredentialVerifyShrinkRequest {
	s.CertNum = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetCredName(v string) *CredentialVerifyShrinkRequest {
	s.CredName = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetCredType(v string) *CredentialVerifyShrinkRequest {
	s.CredType = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetIdentifyNum(v string) *CredentialVerifyShrinkRequest {
	s.IdentifyNum = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetImageContext(v string) *CredentialVerifyShrinkRequest {
	s.ImageContext = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetImageUrl(v string) *CredentialVerifyShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetIsCheck(v string) *CredentialVerifyShrinkRequest {
	s.IsCheck = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetIsOCR(v string) *CredentialVerifyShrinkRequest {
	s.IsOCR = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetMerchantDetailShrink(v string) *CredentialVerifyShrinkRequest {
	s.MerchantDetailShrink = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetMerchantId(v string) *CredentialVerifyShrinkRequest {
	s.MerchantId = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetProductCode(v string) *CredentialVerifyShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetPrompt(v string) *CredentialVerifyShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetPromptModel(v string) *CredentialVerifyShrinkRequest {
	s.PromptModel = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetUserName(v string) *CredentialVerifyShrinkRequest {
	s.UserName = &v
	return s
}

type CredentialVerifyResponseBody struct {
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
	// D6163397-15C5-419C-9ACC-B7C83E0B4C10
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *CredentialVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CredentialVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CredentialVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *CredentialVerifyResponseBody) SetCode(v string) *CredentialVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *CredentialVerifyResponseBody) SetMessage(v string) *CredentialVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *CredentialVerifyResponseBody) SetRequestId(v string) *CredentialVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CredentialVerifyResponseBody) SetResultObject(v *CredentialVerifyResponseBodyResultObject) *CredentialVerifyResponseBody {
	s.ResultObject = v
	return s
}

type CredentialVerifyResponseBodyResultObject struct {
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	OcrInfo      *string `json:"OcrInfo,omitempty" xml:"OcrInfo,omitempty"`
	// example:
	//
	// 1
	Result    *string            `json:"Result,omitempty" xml:"Result,omitempty"`
	RiskScore map[string]*string `json:"RiskScore,omitempty" xml:"RiskScore,omitempty"`
	// example:
	//
	// PS,SCREEN_PHOTO
	RiskTag *string `json:"RiskTag,omitempty" xml:"RiskTag,omitempty"`
	// example:
	//
	// **
	VerifyDetail *string `json:"VerifyDetail,omitempty" xml:"VerifyDetail,omitempty"`
	// example:
	//
	// *
	VerifyResult *string                                           `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
	VlResult     *CredentialVerifyResponseBodyResultObjectVlResult `json:"VlResult,omitempty" xml:"VlResult,omitempty" type:"Struct"`
}

func (s CredentialVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s CredentialVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CredentialVerifyResponseBodyResultObject) SetMaterialInfo(v string) *CredentialVerifyResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetOcrInfo(v string) *CredentialVerifyResponseBodyResultObject {
	s.OcrInfo = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetResult(v string) *CredentialVerifyResponseBodyResultObject {
	s.Result = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetRiskScore(v map[string]*string) *CredentialVerifyResponseBodyResultObject {
	s.RiskScore = v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetRiskTag(v string) *CredentialVerifyResponseBodyResultObject {
	s.RiskTag = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetVerifyDetail(v string) *CredentialVerifyResponseBodyResultObject {
	s.VerifyDetail = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetVerifyResult(v string) *CredentialVerifyResponseBodyResultObject {
	s.VerifyResult = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetVlResult(v *CredentialVerifyResponseBodyResultObjectVlResult) *CredentialVerifyResponseBodyResultObject {
	s.VlResult = v
	return s
}

type CredentialVerifyResponseBodyResultObjectVlResult struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	VlContent *string `json:"VlContent,omitempty" xml:"VlContent,omitempty"`
}

func (s CredentialVerifyResponseBodyResultObjectVlResult) String() string {
	return tea.Prettify(s)
}

func (s CredentialVerifyResponseBodyResultObjectVlResult) GoString() string {
	return s.String()
}

func (s *CredentialVerifyResponseBodyResultObjectVlResult) SetSuccess(v bool) *CredentialVerifyResponseBodyResultObjectVlResult {
	s.Success = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObjectVlResult) SetVlContent(v string) *CredentialVerifyResponseBodyResultObjectVlResult {
	s.VlContent = &v
	return s
}

type CredentialVerifyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CredentialVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s CredentialVerifyResponse) GoString() string {
	return s.String()
}

func (s *CredentialVerifyResponse) SetHeaders(v map[string]*string) *CredentialVerifyResponse {
	s.Headers = v
	return s
}

func (s *CredentialVerifyResponse) SetStatusCode(v int32) *CredentialVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *CredentialVerifyResponse) SetBody(v *CredentialVerifyResponseBody) *CredentialVerifyResponse {
	s.Body = v
	return s
}

type DeepfakeDetectRequest struct {
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
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c******
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
}

func (s DeepfakeDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeepfakeDetectRequest) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectRequest) SetFaceBase64(v string) *DeepfakeDetectRequest {
	s.FaceBase64 = &v
	return s
}

func (s *DeepfakeDetectRequest) SetFaceInputType(v string) *DeepfakeDetectRequest {
	s.FaceInputType = &v
	return s
}

func (s *DeepfakeDetectRequest) SetFaceUrl(v string) *DeepfakeDetectRequest {
	s.FaceUrl = &v
	return s
}

func (s *DeepfakeDetectRequest) SetOuterOrderNo(v string) *DeepfakeDetectRequest {
	s.OuterOrderNo = &v
	return s
}

type DeepfakeDetectResponseBody struct {
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
	// 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DeepfakeDetectResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DeepfakeDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeepfakeDetectResponseBody) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectResponseBody) SetCode(v string) *DeepfakeDetectResponseBody {
	s.Code = &v
	return s
}

func (s *DeepfakeDetectResponseBody) SetMessage(v string) *DeepfakeDetectResponseBody {
	s.Message = &v
	return s
}

func (s *DeepfakeDetectResponseBody) SetRequestId(v string) *DeepfakeDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeepfakeDetectResponseBody) SetResultObject(v *DeepfakeDetectResponseBodyResultObject) *DeepfakeDetectResponseBody {
	s.ResultObject = v
	return s
}

type DeepfakeDetectResponseBodyResultObject struct {
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

func (s DeepfakeDetectResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s DeepfakeDetectResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectResponseBodyResultObject) SetResult(v string) *DeepfakeDetectResponseBodyResultObject {
	s.Result = &v
	return s
}

func (s *DeepfakeDetectResponseBodyResultObject) SetRiskScore(v map[string]*string) *DeepfakeDetectResponseBodyResultObject {
	s.RiskScore = v
	return s
}

func (s *DeepfakeDetectResponseBodyResultObject) SetRiskTag(v string) *DeepfakeDetectResponseBodyResultObject {
	s.RiskTag = &v
	return s
}

type DeepfakeDetectResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeepfakeDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeepfakeDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeepfakeDetectResponse) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectResponse) SetHeaders(v map[string]*string) *DeepfakeDetectResponse {
	s.Headers = v
	return s
}

func (s *DeepfakeDetectResponse) SetStatusCode(v int32) *DeepfakeDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeepfakeDetectResponse) SetBody(v *DeepfakeDetectResponseBody) *DeepfakeDetectResponse {
	s.Body = v
	return s
}

type DeleteFaceVerifyResultRequest struct {
	// example:
	//
	// shae18209d29ce4e8ba252caae98ab15
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// Y
	DeleteAfterQuery *string `json:"DeleteAfterQuery,omitempty" xml:"DeleteAfterQuery,omitempty"`
}

func (s DeleteFaceVerifyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceVerifyResultRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceVerifyResultRequest) SetCertifyId(v string) *DeleteFaceVerifyResultRequest {
	s.CertifyId = &v
	return s
}

func (s *DeleteFaceVerifyResultRequest) SetDeleteAfterQuery(v string) *DeleteFaceVerifyResultRequest {
	s.DeleteAfterQuery = &v
	return s
}

type DeleteFaceVerifyResultResponseBody struct {
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
	// 5A6229C0-E156-48E4-B6EC-0F528BDF60D2
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DeleteFaceVerifyResultResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DeleteFaceVerifyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceVerifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceVerifyResultResponseBody) SetCode(v string) *DeleteFaceVerifyResultResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceVerifyResultResponseBody) SetMessage(v string) *DeleteFaceVerifyResultResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFaceVerifyResultResponseBody) SetRequestId(v string) *DeleteFaceVerifyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceVerifyResultResponseBody) SetResultObject(v *DeleteFaceVerifyResultResponseBodyResultObject) *DeleteFaceVerifyResultResponseBody {
	s.ResultObject = v
	return s
}

type DeleteFaceVerifyResultResponseBodyResultObject struct {
	// example:
	//
	// sha58aeae7ea2f5ed069530f58df4e6d
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// N
	DeleteResult *string `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty"`
	// example:
	//
	// NOT_DELETE_REPEATEDLY
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
}

func (s DeleteFaceVerifyResultResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceVerifyResultResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DeleteFaceVerifyResultResponseBodyResultObject) SetCertifyId(v string) *DeleteFaceVerifyResultResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *DeleteFaceVerifyResultResponseBodyResultObject) SetDeleteResult(v string) *DeleteFaceVerifyResultResponseBodyResultObject {
	s.DeleteResult = &v
	return s
}

func (s *DeleteFaceVerifyResultResponseBodyResultObject) SetFailReason(v string) *DeleteFaceVerifyResultResponseBodyResultObject {
	s.FailReason = &v
	return s
}

type DeleteFaceVerifyResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceVerifyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceVerifyResultResponse) SetHeaders(v map[string]*string) *DeleteFaceVerifyResultResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceVerifyResultResponse) SetStatusCode(v int32) *DeleteFaceVerifyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceVerifyResultResponse) SetBody(v *DeleteFaceVerifyResultResponseBody) *DeleteFaceVerifyResultResponse {
	s.Body = v
	return s
}

type DescribeDeviceInfoRequest struct {
	// example:
	//
	// FACE_TEST
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// wd.6ziUffspAeW5FVYbaqmexR-1qwNjM
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 20200330
	ExpiredEndDay *string `json:"ExpiredEndDay,omitempty" xml:"ExpiredEndDay,omitempty"`
	// example:
	//
	// 20190401
	ExpiredStartDay *string `json:"ExpiredStartDay,omitempty" xml:"ExpiredStartDay,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3iJ1AY$oHcu7mC69
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s DescribeDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoRequest) SetBizType(v string) *DescribeDeviceInfoRequest {
	s.BizType = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetCurrentPage(v int32) *DescribeDeviceInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetDeviceId(v string) *DescribeDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetExpiredEndDay(v string) *DescribeDeviceInfoRequest {
	s.ExpiredEndDay = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetExpiredStartDay(v string) *DescribeDeviceInfoRequest {
	s.ExpiredStartDay = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetPageSize(v int32) *DescribeDeviceInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetUserDeviceId(v string) *DescribeDeviceInfoRequest {
	s.UserDeviceId = &v
	return s
}

type DescribeDeviceInfoResponseBody struct {
	// example:
	//
	// 1
	CurrentPage    *int32                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceInfoList *DescribeDeviceInfoResponseBodyDeviceInfoList `json:"DeviceInfoList,omitempty" xml:"DeviceInfoList,omitempty" type:"Struct"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeviceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBody) SetCurrentPage(v int32) *DescribeDeviceInfoResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetDeviceInfoList(v *DescribeDeviceInfoResponseBodyDeviceInfoList) *DescribeDeviceInfoResponseBody {
	s.DeviceInfoList = v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetPageSize(v int32) *DescribeDeviceInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetRequestId(v string) *DescribeDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetTotalCount(v int32) *DescribeDeviceInfoResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDeviceInfoResponseBodyDeviceInfoList struct {
	DeviceInfo []*DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Repeated"`
}

func (s DescribeDeviceInfoResponseBodyDeviceInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseBodyDeviceInfoList) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoList) SetDeviceInfo(v []*DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) *DescribeDeviceInfoResponseBodyDeviceInfoList {
	s.DeviceInfo = v
	return s
}

type DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo struct {
	// example:
	//
	// 20180101
	BeginDay *string `json:"BeginDay,omitempty" xml:"BeginDay,omitempty"`
	// example:
	//
	// FACE_TEST
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// wd.6ziUffspAeW5FVYbaqmexR-1qwNjM
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 20180101
	ExpiredDay *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
	// example:
	//
	// 3iJ1AY$oHcu7mC69
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetBeginDay(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.BeginDay = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetBizType(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.BizType = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetDeviceId(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetExpiredDay(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.ExpiredDay = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetUserDeviceId(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.UserDeviceId = &v
	return s
}

type DescribeDeviceInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponse) SetHeaders(v map[string]*string) *DescribeDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceInfoResponse) SetStatusCode(v int32) *DescribeDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceInfoResponse) SetBody(v *DescribeDeviceInfoResponseBody) *DescribeDeviceInfoResponse {
	s.Body = v
	return s
}

type DescribeFaceGuardRiskRequest struct {
	// example:
	//
	// aba9830f471a4335af4612c8adaa91b0
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// McozS1ZWRcRZStlERcZZo_QOytx5jcgZoZJEoRLOxxxxxxx
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35xxxx
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// example:
	//
	// FACE_GUARD
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeFaceGuardRiskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceGuardRiskRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaceGuardRiskRequest) SetBizId(v string) *DescribeFaceGuardRiskRequest {
	s.BizId = &v
	return s
}

func (s *DescribeFaceGuardRiskRequest) SetDeviceToken(v string) *DescribeFaceGuardRiskRequest {
	s.DeviceToken = &v
	return s
}

func (s *DescribeFaceGuardRiskRequest) SetOuterOrderNo(v string) *DescribeFaceGuardRiskRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *DescribeFaceGuardRiskRequest) SetProductCode(v string) *DescribeFaceGuardRiskRequest {
	s.ProductCode = &v
	return s
}

type DescribeFaceGuardRiskResponseBody struct {
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
	// D6163397-15C5-419C-9ACC-B7C83E0B4C10
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DescribeFaceGuardRiskResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeFaceGuardRiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceGuardRiskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaceGuardRiskResponseBody) SetCode(v string) *DescribeFaceGuardRiskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFaceGuardRiskResponseBody) SetMessage(v string) *DescribeFaceGuardRiskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFaceGuardRiskResponseBody) SetRequestId(v string) *DescribeFaceGuardRiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaceGuardRiskResponseBody) SetResultObject(v *DescribeFaceGuardRiskResponseBodyResultObject) *DescribeFaceGuardRiskResponseBody {
	s.ResultObject = v
	return s
}

type DescribeFaceGuardRiskResponseBodyResultObject struct {
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// {
	//
	//   "code": 200
	//
	//   "badNet":false,
	//
	//   "umid":"74e37355171ab62230063569350d368e",
	//
	//   "fileTags":"basic_root,basic_hook",
	//
	//   "queryCount":1,
	//
	//   "querySessionCount":1,
	//
	//   "queryUmidCount":1
	//
	//   "platform":"Android"
	//
	// }
	RiskExtends *string `json:"RiskExtends,omitempty" xml:"RiskExtends,omitempty"`
	RiskTags    *string `json:"RiskTags,omitempty" xml:"RiskTags,omitempty"`
}

func (s DescribeFaceGuardRiskResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceGuardRiskResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeFaceGuardRiskResponseBodyResultObject) SetCertifyId(v string) *DescribeFaceGuardRiskResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *DescribeFaceGuardRiskResponseBodyResultObject) SetRiskExtends(v string) *DescribeFaceGuardRiskResponseBodyResultObject {
	s.RiskExtends = &v
	return s
}

func (s *DescribeFaceGuardRiskResponseBodyResultObject) SetRiskTags(v string) *DescribeFaceGuardRiskResponseBodyResultObject {
	s.RiskTags = &v
	return s
}

type DescribeFaceGuardRiskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaceGuardRiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaceGuardRiskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceGuardRiskResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaceGuardRiskResponse) SetHeaders(v map[string]*string) *DescribeFaceGuardRiskResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaceGuardRiskResponse) SetStatusCode(v int32) *DescribeFaceGuardRiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaceGuardRiskResponse) SetBody(v *DescribeFaceGuardRiskResponseBody) *DescribeFaceGuardRiskResponse {
	s.Body = v
	return s
}

type DescribeFaceVerifyRequest struct {
	// example:
	//
	// 91707dc296d469ad38e4c5efa6a0f24b
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// JPG
	PictureReturnType *string `json:"PictureReturnType,omitempty" xml:"PictureReturnType,omitempty"`
	// example:
	//
	// 1000000006
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DescribeFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyRequest) SetCertifyId(v string) *DescribeFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *DescribeFaceVerifyRequest) SetPictureReturnType(v string) *DescribeFaceVerifyRequest {
	s.PictureReturnType = &v
	return s
}

func (s *DescribeFaceVerifyRequest) SetSceneId(v int64) *DescribeFaceVerifyRequest {
	s.SceneId = &v
	return s
}

type DescribeFaceVerifyResponseBody struct {
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
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DescribeFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeFaceVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyResponseBody) SetCode(v string) *DescribeFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFaceVerifyResponseBody) SetMessage(v string) *DescribeFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFaceVerifyResponseBody) SetRequestId(v string) *DescribeFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaceVerifyResponseBody) SetResultObject(v *DescribeFaceVerifyResponseBodyResultObject) *DescribeFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

type DescribeFaceVerifyResponseBodyResultObject struct {
	DeviceRisk *string `json:"DeviceRisk,omitempty" xml:"DeviceRisk,omitempty"`
	// example:
	//
	// McozS1ZWRcRZStlERcZZo_QOytx5jcgZoZJEoRLOxxxxxxx
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	// example:
	//
	// null
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty"`
	// example:
	//
	// {"faceAttack": "F","facialPictureFront": {"qualityScore": 88.3615493774414,"pictureUrl": "https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg","ossBucketName": "cn-shanghai-aliyun-cloudauth-1260051251634779","ossObjectName": "verify/1260051251634779/6ba7bcfccf33f56cdb44ed086f36ce3e0.jpeg"}}
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	// example:
	//
	// T
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 200
	SubCode  *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	Success  *string `json:"Success,omitempty" xml:"Success,omitempty"`
	UserInfo *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s DescribeFaceVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetDeviceRisk(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.DeviceRisk = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetDeviceToken(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.DeviceToken = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetIdentityInfo(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.IdentityInfo = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetMaterialInfo(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetPassed(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetSubCode(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetSuccess(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.Success = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetUserInfo(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.UserInfo = &v
	return s
}

type DescribeFaceVerifyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyResponse) SetHeaders(v map[string]*string) *DescribeFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaceVerifyResponse) SetStatusCode(v int32) *DescribeFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaceVerifyResponse) SetBody(v *DescribeFaceVerifyResponseBody) *DescribeFaceVerifyResponse {
	s.Body = v
	return s
}

type DescribeOssUploadTokenResponseBody struct {
	OssUploadToken *DescribeOssUploadTokenResponseBodyOssUploadToken `json:"OssUploadToken,omitempty" xml:"OssUploadToken,omitempty" type:"Struct"`
	// example:
	//
	// 2FA2C773-47DB-4156-B1EE-5B047321A939
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOssUploadTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssUploadTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssUploadTokenResponseBody) SetOssUploadToken(v *DescribeOssUploadTokenResponseBodyOssUploadToken) *DescribeOssUploadTokenResponseBody {
	s.OssUploadToken = v
	return s
}

func (s *DescribeOssUploadTokenResponseBody) SetRequestId(v string) *DescribeOssUploadTokenResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOssUploadTokenResponseBodyOssUploadToken struct {
	// example:
	//
	// cloudauth-zhangjiakou-external
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// https://oss-cn-zhangjiakou.aliyuncs.com
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// example:
	//
	// 1582636610000
	Expired *int64 `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// STS.NU8rUBj****
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// prod/RdNLC@Ox2n-1s7NMt
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// FwmnyoqT8dHj7nJLuM67T****
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// example:
	//
	// uWia500nTS5knZaDzq4/KqpvhcLnO****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeOssUploadTokenResponseBodyOssUploadToken) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssUploadTokenResponseBodyOssUploadToken) GoString() string {
	return s.String()
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetBucket(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Bucket = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetEndPoint(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.EndPoint = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetExpired(v int64) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Expired = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetKey(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Key = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetPath(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Path = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetSecret(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Secret = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetToken(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Token = &v
	return s
}

type DescribeOssUploadTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssUploadTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssUploadTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssUploadTokenResponse) SetHeaders(v map[string]*string) *DescribeOssUploadTokenResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssUploadTokenResponse) SetStatusCode(v int32) *DescribeOssUploadTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssUploadTokenResponse) SetBody(v *DescribeOssUploadTokenResponseBody) *DescribeOssUploadTokenResponse {
	s.Body = v
	return s
}

type DescribePageFaceVerifyDataRequest struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-04-30
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ID_PLUS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 36**01
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 2023-04-10
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribePageFaceVerifyDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePageFaceVerifyDataRequest) GoString() string {
	return s.String()
}

func (s *DescribePageFaceVerifyDataRequest) SetCurrentPage(v int64) *DescribePageFaceVerifyDataRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePageFaceVerifyDataRequest) SetEndDate(v string) *DescribePageFaceVerifyDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribePageFaceVerifyDataRequest) SetPageSize(v int64) *DescribePageFaceVerifyDataRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePageFaceVerifyDataRequest) SetProductCode(v string) *DescribePageFaceVerifyDataRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribePageFaceVerifyDataRequest) SetSceneId(v int64) *DescribePageFaceVerifyDataRequest {
	s.SceneId = &v
	return s
}

func (s *DescribePageFaceVerifyDataRequest) SetStartDate(v string) *DescribePageFaceVerifyDataRequest {
	s.StartDate = &v
	return s
}

type DescribePageFaceVerifyDataResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	CurrentPage *string                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*DescribePageFaceVerifyDataResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 473469C7-A***B-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 5
	TotalPage *string `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribePageFaceVerifyDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePageFaceVerifyDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePageFaceVerifyDataResponseBody) SetCode(v string) *DescribePageFaceVerifyDataResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetCurrentPage(v string) *DescribePageFaceVerifyDataResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetItems(v []*DescribePageFaceVerifyDataResponseBodyItems) *DescribePageFaceVerifyDataResponseBody {
	s.Items = v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetMessage(v string) *DescribePageFaceVerifyDataResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetPageSize(v string) *DescribePageFaceVerifyDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetRequestId(v string) *DescribePageFaceVerifyDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetSuccess(v string) *DescribePageFaceVerifyDataResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetTotalCount(v string) *DescribePageFaceVerifyDataResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetTotalPage(v string) *DescribePageFaceVerifyDataResponseBody {
	s.TotalPage = &v
	return s
}

type DescribePageFaceVerifyDataResponseBodyItems struct {
	// example:
	//
	// 2024-03-24T00:00:00.000Z
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// ID_PLUS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 20**40
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// 1
	SuccessCount *string `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// 19
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePageFaceVerifyDataResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribePageFaceVerifyDataResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) SetDate(v string) *DescribePageFaceVerifyDataResponseBodyItems {
	s.Date = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) SetProductCode(v string) *DescribePageFaceVerifyDataResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) SetSceneId(v string) *DescribePageFaceVerifyDataResponseBodyItems {
	s.SceneId = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) SetSceneName(v string) *DescribePageFaceVerifyDataResponseBodyItems {
	s.SceneName = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) SetSuccessCount(v string) *DescribePageFaceVerifyDataResponseBodyItems {
	s.SuccessCount = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) SetTotalCount(v string) *DescribePageFaceVerifyDataResponseBodyItems {
	s.TotalCount = &v
	return s
}

type DescribePageFaceVerifyDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePageFaceVerifyDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePageFaceVerifyDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePageFaceVerifyDataResponse) GoString() string {
	return s.String()
}

func (s *DescribePageFaceVerifyDataResponse) SetHeaders(v map[string]*string) *DescribePageFaceVerifyDataResponse {
	s.Headers = v
	return s
}

func (s *DescribePageFaceVerifyDataResponse) SetStatusCode(v int32) *DescribePageFaceVerifyDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponse) SetBody(v *DescribePageFaceVerifyDataResponseBody) *DescribePageFaceVerifyDataResponse {
	s.Body = v
	return s
}

type DescribeSmartStatisticsPageListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-16 23:59:59 +0800
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 36**01
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// cloudauthst
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-11-01 00:00:00 +0800
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeSmartStatisticsPageListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmartStatisticsPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmartStatisticsPageListRequest) SetCurrentPage(v string) *DescribeSmartStatisticsPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSmartStatisticsPageListRequest) SetEndDate(v string) *DescribeSmartStatisticsPageListRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeSmartStatisticsPageListRequest) SetPageSize(v string) *DescribeSmartStatisticsPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSmartStatisticsPageListRequest) SetSceneId(v string) *DescribeSmartStatisticsPageListRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeSmartStatisticsPageListRequest) SetServiceCode(v string) *DescribeSmartStatisticsPageListRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeSmartStatisticsPageListRequest) SetStartDate(v string) *DescribeSmartStatisticsPageListRequest {
	s.StartDate = &v
	return s
}

type DescribeSmartStatisticsPageListResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*DescribeSmartStatisticsPageListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 96943***4E39F805
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 29
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 3
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeSmartStatisticsPageListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmartStatisticsPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmartStatisticsPageListResponseBody) SetCurrentPage(v int32) *DescribeSmartStatisticsPageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBody) SetItems(v []*DescribeSmartStatisticsPageListResponseBodyItems) *DescribeSmartStatisticsPageListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBody) SetPageSize(v int32) *DescribeSmartStatisticsPageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBody) SetRequestId(v string) *DescribeSmartStatisticsPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBody) SetTotalCount(v int32) *DescribeSmartStatisticsPageListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBody) SetTotalPage(v int32) *DescribeSmartStatisticsPageListResponseBody {
	s.TotalPage = &v
	return s
}

type DescribeSmartStatisticsPageListResponseBodyItems struct {
	// example:
	//
	// 11/8
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 25
	PassRate *string `json:"PassRate,omitempty" xml:"PassRate,omitempty"`
	// example:
	//
	// SMART_VERIFY
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 20**40
	SceneId   *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// 1
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSmartStatisticsPageListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmartStatisticsPageListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetDate(v string) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.Date = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetPassRate(v string) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.PassRate = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetProductCode(v string) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetSceneId(v int64) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.SceneId = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetSceneName(v string) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.SceneName = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetSuccessCount(v int32) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.SuccessCount = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetTotalCount(v int32) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.TotalCount = &v
	return s
}

type DescribeSmartStatisticsPageListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSmartStatisticsPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSmartStatisticsPageListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmartStatisticsPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmartStatisticsPageListResponse) SetHeaders(v map[string]*string) *DescribeSmartStatisticsPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmartStatisticsPageListResponse) SetStatusCode(v int32) *DescribeSmartStatisticsPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponse) SetBody(v *DescribeSmartStatisticsPageListResponseBody) *DescribeSmartStatisticsPageListResponse {
	s.Body = v
	return s
}

type DescribeVerifyResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 39ecf51e-2f81-4dc5-90ee-ff86125b****
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FVBioOnlyTest
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s DescribeVerifyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultRequest) SetBizId(v string) *DescribeVerifyResultRequest {
	s.BizId = &v
	return s
}

func (s *DescribeVerifyResultRequest) SetBizType(v string) *DescribeVerifyResultRequest {
	s.BizType = &v
	return s
}

type DescribeVerifyResultResponseBody struct {
	// example:
	//
	// 97
	AuthorityComparisionScore *float32 `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty"`
	// example:
	//
	// 97
	FaceComparisonScore *float32 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	// example:
	//
	// 97
	IdCardFaceComparisonScore *float32                                  `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	Material                  *DescribeVerifyResultResponseBodyMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	VerifyStatus *int32 `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
}

func (s DescribeVerifyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseBody) SetAuthorityComparisionScore(v float32) *DescribeVerifyResultResponseBody {
	s.AuthorityComparisionScore = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetFaceComparisonScore(v float32) *DescribeVerifyResultResponseBody {
	s.FaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetIdCardFaceComparisonScore(v float32) *DescribeVerifyResultResponseBody {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetMaterial(v *DescribeVerifyResultResponseBodyMaterial) *DescribeVerifyResultResponseBody {
	s.Material = v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetRequestId(v string) *DescribeVerifyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetVerifyStatus(v int32) *DescribeVerifyResultResponseBody {
	s.VerifyStatus = &v
	return s
}

type DescribeVerifyResultResponseBodyMaterial struct {
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/face-global-example.jpg
	FaceGlobalUrl *string `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/face-image-example.jpg
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	// example:
	//
	// false
	FaceMask *bool `json:"FaceMask,omitempty" xml:"FaceMask,omitempty"`
	// example:
	//
	// NORMAL
	FaceQuality *string                                             `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	IdCardInfo  *DescribeVerifyResultResponseBodyMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	IdCardName  *string                                             `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	// example:
	//
	// 02343218901123****
	IdCardNumber *string   `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	VideoUrls    []*string `json:"VideoUrls,omitempty" xml:"VideoUrls,omitempty" type:"Repeated"`
}

func (s DescribeVerifyResultResponseBodyMaterial) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseBodyMaterial) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetFaceGlobalUrl(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.FaceGlobalUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetFaceImageUrl(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetFaceMask(v bool) *DescribeVerifyResultResponseBodyMaterial {
	s.FaceMask = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetFaceQuality(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.FaceQuality = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetIdCardInfo(v *DescribeVerifyResultResponseBodyMaterialIdCardInfo) *DescribeVerifyResultResponseBodyMaterial {
	s.IdCardInfo = v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetIdCardName(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.IdCardName = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetIdCardNumber(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetVideoUrls(v []*string) *DescribeVerifyResultResponseBodyMaterial {
	s.VideoUrls = v
	return s
}

type DescribeVerifyResultResponseBodyMaterialIdCardInfo struct {
	Address   *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Authority *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example3.jpg
	BackImageUrl *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty"`
	// example:
	//
	// 19900101
	Birth *string `json:"Birth,omitempty" xml:"Birth,omitempty"`
	// example:
	//
	// 20201101
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example2.jpg
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	// example:
	//
	// 02343218901123****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 20201101
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeVerifyResultResponseBodyMaterialIdCardInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseBodyMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetAddress(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetAuthority(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetBackImageUrl(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetBirth(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetEndDate(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetFrontImageUrl(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetName(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetNationality(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetNumber(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetStartDate(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

type DescribeVerifyResultResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponse) SetHeaders(v map[string]*string) *DescribeVerifyResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyResultResponse) SetStatusCode(v int32) *DescribeVerifyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyResultResponse) SetBody(v *DescribeVerifyResultResponseBody) *DescribeVerifyResultResponse {
	s.Body = v
	return s
}

type DescribeVerifySDKRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1KQMcnLd4m37LN2D0F0WCD
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeVerifySDKRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySDKRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifySDKRequest) SetTaskId(v string) *DescribeVerifySDKRequest {
	s.TaskId = &v
	return s
}

type DescribeVerifySDKResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// https://www.xxx.com
	SdkUrl *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
}

func (s DescribeVerifySDKResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySDKResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifySDKResponseBody) SetRequestId(v string) *DescribeVerifySDKResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifySDKResponseBody) SetSdkUrl(v string) *DescribeVerifySDKResponseBody {
	s.SdkUrl = &v
	return s
}

type DescribeVerifySDKResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifySDKResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifySDKResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySDKResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifySDKResponse) SetHeaders(v map[string]*string) *DescribeVerifySDKResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifySDKResponse) SetStatusCode(v int32) *DescribeVerifySDKResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifySDKResponse) SetBody(v *DescribeVerifySDKResponseBody) *DescribeVerifySDKResponse {
	s.Body = v
	return s
}

type DescribeVerifyTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 39ecf51e-2f81-4dc5-90ee-ff86125be683
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RPBasicTest
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// -
	CallbackSeed *string `json:"CallbackSeed,omitempty" xml:"CallbackSeed,omitempty"`
	// example:
	//
	// -
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// example:
	//
	// http%3A%2F%2Fjiangsu.china.com.cn%2Fuploadfile%2F2015%2F0114%2F1421221304095989.jpg
	FaceRetainedImageUrl *string `json:"FaceRetainedImageUrl,omitempty" xml:"FaceRetainedImageUrl,omitempty"`
	// example:
	//
	// -
	FailedRedirectUrl *string `json:"FailedRedirectUrl,omitempty" xml:"FailedRedirectUrl,omitempty"`
	// example:
	//
	// http%3A%2F%2Fjiangsu.china.com.cn%2Fuploadfile%2F2015%2F0114%2F1421221304095989.jpg
	IdCardBackImageUrl *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	// example:
	//
	// http%3A%2F%2Fjiangsu.china.com.cn%2Fuploadfile%2F2015%2F0114%2F1421221304095989.jpg
	IdCardFrontImageUrl *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	// example:
	//
	// 330100xxxxxxxxxxxx
	IdCardNumber *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// -
	PassedRedirectUrl *string `json:"PassedRedirectUrl,omitempty" xml:"PassedRedirectUrl,omitempty"`
	// example:
	//
	// user111
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 192.168.***.***
	UserIp *string `json:"UserIp,omitempty" xml:"UserIp,omitempty"`
	// example:
	//
	// 187********
	UserPhoneNumber *string `json:"UserPhoneNumber,omitempty" xml:"UserPhoneNumber,omitempty"`
	// example:
	//
	// 1577808000000
	UserRegistTime *int64 `json:"UserRegistTime,omitempty" xml:"UserRegistTime,omitempty"`
}

func (s DescribeVerifyTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenRequest) SetBizId(v string) *DescribeVerifyTokenRequest {
	s.BizId = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetBizType(v string) *DescribeVerifyTokenRequest {
	s.BizType = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetCallbackSeed(v string) *DescribeVerifyTokenRequest {
	s.CallbackSeed = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetCallbackUrl(v string) *DescribeVerifyTokenRequest {
	s.CallbackUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetFaceRetainedImageUrl(v string) *DescribeVerifyTokenRequest {
	s.FaceRetainedImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetFailedRedirectUrl(v string) *DescribeVerifyTokenRequest {
	s.FailedRedirectUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardBackImageUrl(v string) *DescribeVerifyTokenRequest {
	s.IdCardBackImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardFrontImageUrl(v string) *DescribeVerifyTokenRequest {
	s.IdCardFrontImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardNumber(v string) *DescribeVerifyTokenRequest {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetName(v string) *DescribeVerifyTokenRequest {
	s.Name = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetPassedRedirectUrl(v string) *DescribeVerifyTokenRequest {
	s.PassedRedirectUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserId(v string) *DescribeVerifyTokenRequest {
	s.UserId = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserIp(v string) *DescribeVerifyTokenRequest {
	s.UserIp = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserPhoneNumber(v string) *DescribeVerifyTokenRequest {
	s.UserPhoneNumber = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserRegistTime(v int64) *DescribeVerifyTokenRequest {
	s.UserRegistTime = &v
	return s
}

type DescribeVerifyTokenResponseBody struct {
	OssUploadToken *DescribeVerifyTokenResponseBodyOssUploadToken `json:"OssUploadToken,omitempty" xml:"OssUploadToken,omitempty" type:"Struct"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// http%3A%2F%2Fjiangsu.china.com.cn%2Fuploadfile%2F2015%2F0114%2F1421221304095989.jpg
	VerifyPageUrl *string `json:"VerifyPageUrl,omitempty" xml:"VerifyPageUrl,omitempty"`
	// example:
	//
	// c302c0797679457685410ee51a5ba375
	VerifyToken *string `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
}

func (s DescribeVerifyTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponseBody) SetOssUploadToken(v *DescribeVerifyTokenResponseBodyOssUploadToken) *DescribeVerifyTokenResponseBody {
	s.OssUploadToken = v
	return s
}

func (s *DescribeVerifyTokenResponseBody) SetRequestId(v string) *DescribeVerifyTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyTokenResponseBody) SetVerifyPageUrl(v string) *DescribeVerifyTokenResponseBody {
	s.VerifyPageUrl = &v
	return s
}

func (s *DescribeVerifyTokenResponseBody) SetVerifyToken(v string) *DescribeVerifyTokenResponseBody {
	s.VerifyToken = &v
	return s
}

type DescribeVerifyTokenResponseBodyOssUploadToken struct {
	// example:
	//
	// cloudauth-zhangjiakou-external
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// https://oss-cn-zhangjiakou.aliyuncs.com
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// example:
	//
	// 1582636610000
	Expired *int64 `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// STS.NU8rUBj****
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// prod/RdNLC@Ox2n-1s7NMt
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// FwmnyoqT8dHj7nJLuM67T****
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// example:
	//
	// uWia500nTS5knZaDzq4/KqpvhcLnO****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeVerifyTokenResponseBodyOssUploadToken) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenResponseBodyOssUploadToken) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetBucket(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Bucket = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetEndPoint(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.EndPoint = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetExpired(v int64) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Expired = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetKey(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Key = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetPath(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Path = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetSecret(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Secret = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetToken(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Token = &v
	return s
}

type DescribeVerifyTokenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifyTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponse) SetHeaders(v map[string]*string) *DescribeVerifyTokenResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyTokenResponse) SetStatusCode(v int32) *DescribeVerifyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyTokenResponse) SetBody(v *DescribeVerifyTokenResponseBody) *DescribeVerifyTokenResponse {
	s.Body = v
	return s
}

type DetectFaceAttributesRequest struct {
	// example:
	//
	// RPBasicTest
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg
	MaterialValue *string `json:"MaterialValue,omitempty" xml:"MaterialValue,omitempty"`
}

func (s DetectFaceAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesRequest) SetBizType(v string) *DetectFaceAttributesRequest {
	s.BizType = &v
	return s
}

func (s *DetectFaceAttributesRequest) SetMaterialValue(v string) *DetectFaceAttributesRequest {
	s.MaterialValue = &v
	return s
}

type DetectFaceAttributesResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DetectFaceAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Error.InternalError
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetectFaceAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBody) SetCode(v string) *DetectFaceAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetData(v *DetectFaceAttributesResponseBodyData) *DetectFaceAttributesResponseBody {
	s.Data = v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetMessage(v string) *DetectFaceAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetRequestId(v string) *DetectFaceAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetSuccess(v bool) *DetectFaceAttributesResponseBody {
	s.Success = &v
	return s
}

type DetectFaceAttributesResponseBodyData struct {
	FaceInfos *DetectFaceAttributesResponseBodyDataFaceInfos `json:"FaceInfos,omitempty" xml:"FaceInfos,omitempty" type:"Struct"`
	// example:
	//
	// 1920
	ImgHeight *int32 `json:"ImgHeight,omitempty" xml:"ImgHeight,omitempty"`
	// example:
	//
	// 1080
	ImgWidth *int32 `json:"ImgWidth,omitempty" xml:"ImgWidth,omitempty"`
}

func (s DetectFaceAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyData) SetFaceInfos(v *DetectFaceAttributesResponseBodyDataFaceInfos) *DetectFaceAttributesResponseBodyData {
	s.FaceInfos = v
	return s
}

func (s *DetectFaceAttributesResponseBodyData) SetImgHeight(v int32) *DetectFaceAttributesResponseBodyData {
	s.ImgHeight = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyData) SetImgWidth(v int32) *DetectFaceAttributesResponseBodyData {
	s.ImgWidth = &v
	return s
}

type DetectFaceAttributesResponseBodyDataFaceInfos struct {
	FaceAttributesDetectInfo []*DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo `json:"FaceAttributesDetectInfo,omitempty" xml:"FaceAttributesDetectInfo,omitempty" type:"Repeated"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfos) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfos) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfos) SetFaceAttributesDetectInfo(v []*DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) *DetectFaceAttributesResponseBodyDataFaceInfos {
	s.FaceAttributesDetectInfo = v
	return s
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo struct {
	FaceAttributes *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	FaceRect       *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect       `json:"FaceRect,omitempty" xml:"FaceRect,omitempty" type:"Struct"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) SetFaceAttributes(v *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo {
	s.FaceAttributes = v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) SetFaceRect(v *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo {
	s.FaceRect = v
	return s
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes struct {
	// example:
	//
	// 0.1419367790222168
	Blur *float32 `json:"Blur,omitempty" xml:"Blur,omitempty"`
	// example:
	//
	// 60
	Facequal *float32 `json:"Facequal,omitempty" xml:"Facequal,omitempty"`
	// example:
	//
	// Face
	Facetype *string `json:"Facetype,omitempty" xml:"Facetype,omitempty"`
	// example:
	//
	// None
	Glasses  *string                                                                                      `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	Headpose *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose `json:"Headpose,omitempty" xml:"Headpose,omitempty" type:"Struct"`
	// example:
	//
	// 70
	Integrity *int32 `json:"Integrity,omitempty" xml:"Integrity,omitempty"`
	// example:
	//
	// Wear
	Respirator *string                                                                                     `json:"Respirator,omitempty" xml:"Respirator,omitempty"`
	Smiling    *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling `json:"Smiling,omitempty" xml:"Smiling,omitempty" type:"Struct"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetBlur(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Blur = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetFacequal(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Facequal = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetFacetype(v string) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Facetype = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetGlasses(v string) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetHeadpose(v *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Headpose = v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetIntegrity(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Integrity = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetRespirator(v string) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Respirator = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetSmiling(v *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Smiling = v
	return s
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose struct {
	// example:
	//
	// -1.5683923959732056
	PitchAngle *float32 `json:"PitchAngle,omitempty" xml:"PitchAngle,omitempty"`
	// example:
	//
	// 7.163370132446289
	RollAngle *float32 `json:"RollAngle,omitempty" xml:"RollAngle,omitempty"`
	// example:
	//
	// -6.925303936004639
	YawAngle *float32 `json:"YawAngle,omitempty" xml:"YawAngle,omitempty"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetPitchAngle(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.PitchAngle = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetRollAngle(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.RollAngle = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetYawAngle(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.YawAngle = &v
	return s
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling struct {
	// example:
	//
	// 95
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// example:
	//
	// 97
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) SetThreshold(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling {
	s.Threshold = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) SetValue(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling {
	s.Value = &v
	return s
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect struct {
	// example:
	//
	// 473
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 354
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 453
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 473
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) SetHeight(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Height = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) SetLeft(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Left = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) SetTop(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Top = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) SetWidth(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Width = &v
	return s
}

type DetectFaceAttributesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectFaceAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectFaceAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponse) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponse) SetHeaders(v map[string]*string) *DetectFaceAttributesResponse {
	s.Headers = v
	return s
}

func (s *DetectFaceAttributesResponse) SetStatusCode(v int32) *DetectFaceAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectFaceAttributesResponse) SetBody(v *DetectFaceAttributesResponseBody) *DetectFaceAttributesResponse {
	s.Body = v
	return s
}

type Id2MetaVerifyRequest struct {
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	ParamType   *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Id2MetaVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyRequest) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyRequest) SetIdentifyNum(v string) *Id2MetaVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Id2MetaVerifyRequest) SetParamType(v string) *Id2MetaVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Id2MetaVerifyRequest) SetUserName(v string) *Id2MetaVerifyRequest {
	s.UserName = &v
	return s
}

type Id2MetaVerifyResponseBody struct {
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
	// D6163397-15C5-419C-9ACC-B7C83E0B4C10
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *Id2MetaVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Id2MetaVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyResponseBody) SetCode(v string) *Id2MetaVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Id2MetaVerifyResponseBody) SetMessage(v string) *Id2MetaVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Id2MetaVerifyResponseBody) SetRequestId(v string) *Id2MetaVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id2MetaVerifyResponseBody) SetResultObject(v *Id2MetaVerifyResponseBodyResultObject) *Id2MetaVerifyResponseBody {
	s.ResultObject = v
	return s
}

type Id2MetaVerifyResponseBodyResultObject struct {
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
}

func (s Id2MetaVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyResponseBodyResultObject) SetBizCode(v string) *Id2MetaVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

type Id2MetaVerifyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Id2MetaVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Id2MetaVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyResponse) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyResponse) SetHeaders(v map[string]*string) *Id2MetaVerifyResponse {
	s.Headers = v
	return s
}

func (s *Id2MetaVerifyResponse) SetStatusCode(v int32) *Id2MetaVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Id2MetaVerifyResponse) SetBody(v *Id2MetaVerifyResponseBody) *Id2MetaVerifyResponse {
	s.Body = v
	return s
}

type InitFaceVerifyRequest struct {
	AppQualityCheck *string `json:"AppQualityCheck,omitempty" xml:"AppQualityCheck,omitempty"`
	AuthId          *string `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	Birthday        *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	CallbackToken   *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	CallbackUrl     *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	CertName        *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// example:
	//
	// IDENTITY_CARD
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// example:
	//
	// 0bfa7c493f850e5178b9f8613634c9xx
	CertifyId       *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	CertifyUrlStyle *string `json:"CertifyUrlStyle,omitempty" xml:"CertifyUrlStyle,omitempty"`
	CertifyUrlType  *string `json:"CertifyUrlType,omitempty" xml:"CertifyUrlType,omitempty"`
	Crop            *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	EncryptType     *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceContrastPicture *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	FaceGuardOutput        *string `json:"FaceGuardOutput,omitempty" xml:"FaceGuardOutput,omitempty"`
	// example:
	//
	// 114.xxx.xxx.xxx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// {"zimVer":"3.0.0","appVersion": "1","bioMetaInfo": "4.1.0:11501568,0","appName": "com.aliyun.antcloudauth","deviceType": "ios","osVersion": "iOS 10.3.2","apdidToken": "","deviceModel": "iPhone9,1"}
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// example:
	//
	// 130xxxxxxxx
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Mode   *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Model  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo      *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProcedurePriority *string `json:"ProcedurePriority,omitempty" xml:"ProcedurePriority,omitempty"`
	// example:
	//
	// ID_PRO
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RarelyCharacters *string `json:"RarelyCharacters,omitempty" xml:"RarelyCharacters,omitempty"`
	ReadImg          *string `json:"ReadImg,omitempty" xml:"ReadImg,omitempty"`
	// example:
	//
	// www.aliyun.com
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	// example:
	//
	// 1000000006
	SceneId      *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SuitableType *string `json:"SuitableType,omitempty" xml:"SuitableType,omitempty"`
	UiCustomUrl  *string `json:"UiCustomUrl,omitempty" xml:"UiCustomUrl,omitempty"`
	// example:
	//
	// 123456789
	UserId                     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ValidityDate               *string `json:"ValidityDate,omitempty" xml:"ValidityDate,omitempty"`
	VideoEvidence              *string `json:"VideoEvidence,omitempty" xml:"VideoEvidence,omitempty"`
	VoluntaryCustomizedContent *string `json:"VoluntaryCustomizedContent,omitempty" xml:"VoluntaryCustomizedContent,omitempty"`
}

func (s InitFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyRequest) SetAppQualityCheck(v string) *InitFaceVerifyRequest {
	s.AppQualityCheck = &v
	return s
}

func (s *InitFaceVerifyRequest) SetAuthId(v string) *InitFaceVerifyRequest {
	s.AuthId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetBirthday(v string) *InitFaceVerifyRequest {
	s.Birthday = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCallbackToken(v string) *InitFaceVerifyRequest {
	s.CallbackToken = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCallbackUrl(v string) *InitFaceVerifyRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertName(v string) *InitFaceVerifyRequest {
	s.CertName = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertNo(v string) *InitFaceVerifyRequest {
	s.CertNo = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertType(v string) *InitFaceVerifyRequest {
	s.CertType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertifyId(v string) *InitFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertifyUrlStyle(v string) *InitFaceVerifyRequest {
	s.CertifyUrlStyle = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertifyUrlType(v string) *InitFaceVerifyRequest {
	s.CertifyUrlType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCrop(v string) *InitFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *InitFaceVerifyRequest) SetEncryptType(v string) *InitFaceVerifyRequest {
	s.EncryptType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetFaceContrastPicture(v string) *InitFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *InitFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *InitFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetFaceGuardOutput(v string) *InitFaceVerifyRequest {
	s.FaceGuardOutput = &v
	return s
}

func (s *InitFaceVerifyRequest) SetIp(v string) *InitFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *InitFaceVerifyRequest) SetMetaInfo(v string) *InitFaceVerifyRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitFaceVerifyRequest) SetMobile(v string) *InitFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *InitFaceVerifyRequest) SetMode(v string) *InitFaceVerifyRequest {
	s.Mode = &v
	return s
}

func (s *InitFaceVerifyRequest) SetModel(v string) *InitFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *InitFaceVerifyRequest) SetOssBucketName(v string) *InitFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *InitFaceVerifyRequest) SetOssObjectName(v string) *InitFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *InitFaceVerifyRequest) SetOuterOrderNo(v string) *InitFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *InitFaceVerifyRequest) SetProcedurePriority(v string) *InitFaceVerifyRequest {
	s.ProcedurePriority = &v
	return s
}

func (s *InitFaceVerifyRequest) SetProductCode(v string) *InitFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *InitFaceVerifyRequest) SetRarelyCharacters(v string) *InitFaceVerifyRequest {
	s.RarelyCharacters = &v
	return s
}

func (s *InitFaceVerifyRequest) SetReadImg(v string) *InitFaceVerifyRequest {
	s.ReadImg = &v
	return s
}

func (s *InitFaceVerifyRequest) SetReturnUrl(v string) *InitFaceVerifyRequest {
	s.ReturnUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetSceneId(v int64) *InitFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetSuitableType(v string) *InitFaceVerifyRequest {
	s.SuitableType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetUiCustomUrl(v string) *InitFaceVerifyRequest {
	s.UiCustomUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetUserId(v string) *InitFaceVerifyRequest {
	s.UserId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetValidityDate(v string) *InitFaceVerifyRequest {
	s.ValidityDate = &v
	return s
}

func (s *InitFaceVerifyRequest) SetVideoEvidence(v string) *InitFaceVerifyRequest {
	s.VideoEvidence = &v
	return s
}

func (s *InitFaceVerifyRequest) SetVoluntaryCustomizedContent(v string) *InitFaceVerifyRequest {
	s.VoluntaryCustomizedContent = &v
	return s
}

type InitFaceVerifyResponseBody struct {
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
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *InitFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s InitFaceVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponseBody) SetCode(v string) *InitFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *InitFaceVerifyResponseBody) SetMessage(v string) *InitFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *InitFaceVerifyResponseBody) SetRequestId(v string) *InitFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitFaceVerifyResponseBody) SetResultObject(v *InitFaceVerifyResponseBodyResultObject) *InitFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

type InitFaceVerifyResponseBodyResultObject struct {
	// example:
	//
	// 91707dc296d469ad38e4c5efa6a0f24b
	CertifyId  *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	CertifyUrl *string `json:"CertifyUrl,omitempty" xml:"CertifyUrl,omitempty"`
}

func (s InitFaceVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *InitFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *InitFaceVerifyResponseBodyResultObject) SetCertifyUrl(v string) *InitFaceVerifyResponseBodyResultObject {
	s.CertifyUrl = &v
	return s
}

type InitFaceVerifyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponse) SetHeaders(v map[string]*string) *InitFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *InitFaceVerifyResponse) SetStatusCode(v int32) *InitFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *InitFaceVerifyResponse) SetBody(v *InitFaceVerifyResponseBody) *InitFaceVerifyResponse {
	s.Body = v
	return s
}

type InsertWhiteListSettingRequest struct {
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// example:
	//
	// shsf57a4e0d9981c3bd66dc754f3d3cd
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// xxxxxx
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 100000xxxx
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// 30
	ValidDay *int32 `json:"ValidDay,omitempty" xml:"ValidDay,omitempty"`
}

func (s InsertWhiteListSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertWhiteListSettingRequest) GoString() string {
	return s.String()
}

func (s *InsertWhiteListSettingRequest) SetCertNo(v string) *InsertWhiteListSettingRequest {
	s.CertNo = &v
	return s
}

func (s *InsertWhiteListSettingRequest) SetCertifyId(v string) *InsertWhiteListSettingRequest {
	s.CertifyId = &v
	return s
}

func (s *InsertWhiteListSettingRequest) SetRemark(v string) *InsertWhiteListSettingRequest {
	s.Remark = &v
	return s
}

func (s *InsertWhiteListSettingRequest) SetSceneId(v int64) *InsertWhiteListSettingRequest {
	s.SceneId = &v
	return s
}

func (s *InsertWhiteListSettingRequest) SetServiceCode(v string) *InsertWhiteListSettingRequest {
	s.ServiceCode = &v
	return s
}

func (s *InsertWhiteListSettingRequest) SetValidDay(v int32) *InsertWhiteListSettingRequest {
	s.ValidDay = &v
	return s
}

type InsertWhiteListSettingResponseBody struct {
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertWhiteListSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertWhiteListSettingResponseBody) GoString() string {
	return s.String()
}

func (s *InsertWhiteListSettingResponseBody) SetCode(v string) *InsertWhiteListSettingResponseBody {
	s.Code = &v
	return s
}

func (s *InsertWhiteListSettingResponseBody) SetMessage(v string) *InsertWhiteListSettingResponseBody {
	s.Message = &v
	return s
}

func (s *InsertWhiteListSettingResponseBody) SetRequestId(v string) *InsertWhiteListSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertWhiteListSettingResponseBody) SetResultObject(v bool) *InsertWhiteListSettingResponseBody {
	s.ResultObject = &v
	return s
}

func (s *InsertWhiteListSettingResponseBody) SetSuccess(v bool) *InsertWhiteListSettingResponseBody {
	s.Success = &v
	return s
}

type InsertWhiteListSettingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertWhiteListSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertWhiteListSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertWhiteListSettingResponse) GoString() string {
	return s.String()
}

func (s *InsertWhiteListSettingResponse) SetHeaders(v map[string]*string) *InsertWhiteListSettingResponse {
	s.Headers = v
	return s
}

func (s *InsertWhiteListSettingResponse) SetStatusCode(v int32) *InsertWhiteListSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertWhiteListSettingResponse) SetBody(v *InsertWhiteListSettingResponseBody) *InsertWhiteListSettingResponse {
	s.Body = v
	return s
}

type LivenessFaceVerifyRequest struct {
	CertifyId              *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Crop                   *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DeviceToken            *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	FaceContrastPicture    *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	Ip                     *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Mobile                 *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Model                  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OssBucketName          *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	OuterOrderNo           *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SceneId                *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LivenessFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyRequest) SetCertifyId(v string) *LivenessFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetCrop(v string) *LivenessFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetDeviceToken(v string) *LivenessFaceVerifyRequest {
	s.DeviceToken = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetFaceContrastPicture(v string) *LivenessFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *LivenessFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetIp(v string) *LivenessFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetMobile(v string) *LivenessFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetModel(v string) *LivenessFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetOssBucketName(v string) *LivenessFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetOssObjectName(v string) *LivenessFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetOuterOrderNo(v string) *LivenessFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetProductCode(v string) *LivenessFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetSceneId(v int64) *LivenessFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetUserId(v string) *LivenessFaceVerifyRequest {
	s.UserId = &v
	return s
}

type LivenessFaceVerifyResponseBody struct {
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *LivenessFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s LivenessFaceVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyResponseBody) SetCode(v string) *LivenessFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *LivenessFaceVerifyResponseBody) SetMessage(v string) *LivenessFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *LivenessFaceVerifyResponseBody) SetRequestId(v string) *LivenessFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *LivenessFaceVerifyResponseBody) SetResultObject(v *LivenessFaceVerifyResponseBodyResultObject) *LivenessFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

type LivenessFaceVerifyResponseBodyResultObject struct {
	CertifyId    *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s LivenessFaceVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *LivenessFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *LivenessFaceVerifyResponseBodyResultObject) SetMaterialInfo(v string) *LivenessFaceVerifyResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *LivenessFaceVerifyResponseBodyResultObject) SetPassed(v string) *LivenessFaceVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *LivenessFaceVerifyResponseBodyResultObject) SetSubCode(v string) *LivenessFaceVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

type LivenessFaceVerifyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LivenessFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LivenessFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyResponse) SetHeaders(v map[string]*string) *LivenessFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *LivenessFaceVerifyResponse) SetStatusCode(v int32) *LivenessFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *LivenessFaceVerifyResponse) SetBody(v *LivenessFaceVerifyResponseBody) *LivenessFaceVerifyResponse {
	s.Body = v
	return s
}

type Mobile3MetaDetailVerifyRequest struct {
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	ParamType   *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Mobile3MetaDetailVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaDetailVerifyRequest) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailVerifyRequest) SetIdentifyNum(v string) *Mobile3MetaDetailVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Mobile3MetaDetailVerifyRequest) SetMobile(v string) *Mobile3MetaDetailVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *Mobile3MetaDetailVerifyRequest) SetParamType(v string) *Mobile3MetaDetailVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Mobile3MetaDetailVerifyRequest) SetUserName(v string) *Mobile3MetaDetailVerifyRequest {
	s.UserName = &v
	return s
}

type Mobile3MetaDetailVerifyResponseBody struct {
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
	// 5A6229C0-E156-48E4-B6EC-0F528BDF60D2
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *Mobile3MetaDetailVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Mobile3MetaDetailVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaDetailVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailVerifyResponseBody) SetCode(v string) *Mobile3MetaDetailVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBody) SetMessage(v string) *Mobile3MetaDetailVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBody) SetRequestId(v string) *Mobile3MetaDetailVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBody) SetResultObject(v *Mobile3MetaDetailVerifyResponseBodyResultObject) *Mobile3MetaDetailVerifyResponseBody {
	s.ResultObject = v
	return s
}

type Mobile3MetaDetailVerifyResponseBodyResultObject struct {
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

func (s Mobile3MetaDetailVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaDetailVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailVerifyResponseBodyResultObject) SetBizCode(v string) *Mobile3MetaDetailVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBodyResultObject) SetIspName(v string) *Mobile3MetaDetailVerifyResponseBodyResultObject {
	s.IspName = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBodyResultObject) SetSubCode(v string) *Mobile3MetaDetailVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

type Mobile3MetaDetailVerifyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Mobile3MetaDetailVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Mobile3MetaDetailVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaDetailVerifyResponse) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailVerifyResponse) SetHeaders(v map[string]*string) *Mobile3MetaDetailVerifyResponse {
	s.Headers = v
	return s
}

func (s *Mobile3MetaDetailVerifyResponse) SetStatusCode(v int32) *Mobile3MetaDetailVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponse) SetBody(v *Mobile3MetaDetailVerifyResponseBody) *Mobile3MetaDetailVerifyResponse {
	s.Body = v
	return s
}

type Mobile3MetaSimpleVerifyRequest struct {
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	ParamType   *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Mobile3MetaSimpleVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaSimpleVerifyRequest) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleVerifyRequest) SetIdentifyNum(v string) *Mobile3MetaSimpleVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyRequest) SetMobile(v string) *Mobile3MetaSimpleVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyRequest) SetParamType(v string) *Mobile3MetaSimpleVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyRequest) SetUserName(v string) *Mobile3MetaSimpleVerifyRequest {
	s.UserName = &v
	return s
}

type Mobile3MetaSimpleVerifyResponseBody struct {
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *Mobile3MetaSimpleVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Mobile3MetaSimpleVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaSimpleVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleVerifyResponseBody) SetCode(v string) *Mobile3MetaSimpleVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBody) SetMessage(v string) *Mobile3MetaSimpleVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBody) SetRequestId(v string) *Mobile3MetaSimpleVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBody) SetResultObject(v *Mobile3MetaSimpleVerifyResponseBodyResultObject) *Mobile3MetaSimpleVerifyResponseBody {
	s.ResultObject = v
	return s
}

type Mobile3MetaSimpleVerifyResponseBodyResultObject struct {
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s Mobile3MetaSimpleVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaSimpleVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleVerifyResponseBodyResultObject) SetBizCode(v string) *Mobile3MetaSimpleVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBodyResultObject) SetIspName(v string) *Mobile3MetaSimpleVerifyResponseBodyResultObject {
	s.IspName = &v
	return s
}

type Mobile3MetaSimpleVerifyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Mobile3MetaSimpleVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Mobile3MetaSimpleVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaSimpleVerifyResponse) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleVerifyResponse) SetHeaders(v map[string]*string) *Mobile3MetaSimpleVerifyResponse {
	s.Headers = v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponse) SetStatusCode(v int32) *Mobile3MetaSimpleVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponse) SetBody(v *Mobile3MetaSimpleVerifyResponseBody) *Mobile3MetaSimpleVerifyResponse {
	s.Body = v
	return s
}

type MobileDetectRequest struct {
	Mobiles   *string `json:"Mobiles,omitempty" xml:"Mobiles,omitempty"`
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s MobileDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s MobileDetectRequest) GoString() string {
	return s.String()
}

func (s *MobileDetectRequest) SetMobiles(v string) *MobileDetectRequest {
	s.Mobiles = &v
	return s
}

func (s *MobileDetectRequest) SetParamType(v string) *MobileDetectRequest {
	s.ParamType = &v
	return s
}

type MobileDetectResponseBody struct {
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
	// 969434DF-926B-4997-9881-4DE94E39F805
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *MobileDetectResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s MobileDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MobileDetectResponseBody) GoString() string {
	return s.String()
}

func (s *MobileDetectResponseBody) SetCode(v string) *MobileDetectResponseBody {
	s.Code = &v
	return s
}

func (s *MobileDetectResponseBody) SetMessage(v string) *MobileDetectResponseBody {
	s.Message = &v
	return s
}

func (s *MobileDetectResponseBody) SetRequestId(v string) *MobileDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *MobileDetectResponseBody) SetResultObject(v *MobileDetectResponseBodyResultObject) *MobileDetectResponseBody {
	s.ResultObject = v
	return s
}

type MobileDetectResponseBodyResultObject struct {
	// example:
	//
	// 2
	ChargeCount *string                                      `json:"ChargeCount,omitempty" xml:"ChargeCount,omitempty"`
	Items       []*MobileDetectResponseBodyResultObjectItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s MobileDetectResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s MobileDetectResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *MobileDetectResponseBodyResultObject) SetChargeCount(v string) *MobileDetectResponseBodyResultObject {
	s.ChargeCount = &v
	return s
}

func (s *MobileDetectResponseBodyResultObject) SetItems(v []*MobileDetectResponseBodyResultObjectItems) *MobileDetectResponseBodyResultObject {
	s.Items = v
	return s
}

type MobileDetectResponseBodyResultObjectItems struct {
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
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
	// 131********
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s MobileDetectResponseBodyResultObjectItems) String() string {
	return tea.Prettify(s)
}

func (s MobileDetectResponseBodyResultObjectItems) GoString() string {
	return s.String()
}

func (s *MobileDetectResponseBodyResultObjectItems) SetArea(v string) *MobileDetectResponseBodyResultObjectItems {
	s.Area = &v
	return s
}

func (s *MobileDetectResponseBodyResultObjectItems) SetBizCode(v string) *MobileDetectResponseBodyResultObjectItems {
	s.BizCode = &v
	return s
}

func (s *MobileDetectResponseBodyResultObjectItems) SetIspName(v string) *MobileDetectResponseBodyResultObjectItems {
	s.IspName = &v
	return s
}

func (s *MobileDetectResponseBodyResultObjectItems) SetMobile(v string) *MobileDetectResponseBodyResultObjectItems {
	s.Mobile = &v
	return s
}

func (s *MobileDetectResponseBodyResultObjectItems) SetSubCode(v string) *MobileDetectResponseBodyResultObjectItems {
	s.SubCode = &v
	return s
}

type MobileDetectResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MobileDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MobileDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s MobileDetectResponse) GoString() string {
	return s.String()
}

func (s *MobileDetectResponse) SetHeaders(v map[string]*string) *MobileDetectResponse {
	s.Headers = v
	return s
}

func (s *MobileDetectResponse) SetStatusCode(v int32) *MobileDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *MobileDetectResponse) SetBody(v *MobileDetectResponseBody) *MobileDetectResponse {
	s.Body = v
	return s
}

type MobileOnlineStatusRequest struct {
	Mobile    *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s MobileOnlineStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s MobileOnlineStatusRequest) GoString() string {
	return s.String()
}

func (s *MobileOnlineStatusRequest) SetMobile(v string) *MobileOnlineStatusRequest {
	s.Mobile = &v
	return s
}

func (s *MobileOnlineStatusRequest) SetParamType(v string) *MobileOnlineStatusRequest {
	s.ParamType = &v
	return s
}

type MobileOnlineStatusResponseBody struct {
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
	// B506328A-D84B-4750-82C7-6A207C585CF1
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *MobileOnlineStatusResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s MobileOnlineStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MobileOnlineStatusResponseBody) GoString() string {
	return s.String()
}

func (s *MobileOnlineStatusResponseBody) SetCode(v string) *MobileOnlineStatusResponseBody {
	s.Code = &v
	return s
}

func (s *MobileOnlineStatusResponseBody) SetMessage(v string) *MobileOnlineStatusResponseBody {
	s.Message = &v
	return s
}

func (s *MobileOnlineStatusResponseBody) SetRequestId(v string) *MobileOnlineStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *MobileOnlineStatusResponseBody) SetResultObject(v *MobileOnlineStatusResponseBodyResultObject) *MobileOnlineStatusResponseBody {
	s.ResultObject = v
	return s
}

type MobileOnlineStatusResponseBodyResultObject struct {
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

func (s MobileOnlineStatusResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s MobileOnlineStatusResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *MobileOnlineStatusResponseBodyResultObject) SetBizCode(v string) *MobileOnlineStatusResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *MobileOnlineStatusResponseBodyResultObject) SetIspName(v string) *MobileOnlineStatusResponseBodyResultObject {
	s.IspName = &v
	return s
}

func (s *MobileOnlineStatusResponseBodyResultObject) SetSubCode(v string) *MobileOnlineStatusResponseBodyResultObject {
	s.SubCode = &v
	return s
}

type MobileOnlineStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MobileOnlineStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MobileOnlineStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s MobileOnlineStatusResponse) GoString() string {
	return s.String()
}

func (s *MobileOnlineStatusResponse) SetHeaders(v map[string]*string) *MobileOnlineStatusResponse {
	s.Headers = v
	return s
}

func (s *MobileOnlineStatusResponse) SetStatusCode(v int32) *MobileOnlineStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *MobileOnlineStatusResponse) SetBody(v *MobileOnlineStatusResponseBody) *MobileOnlineStatusResponse {
	s.Body = v
	return s
}

type MobileOnlineTimeRequest struct {
	Mobile    *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s MobileOnlineTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s MobileOnlineTimeRequest) GoString() string {
	return s.String()
}

func (s *MobileOnlineTimeRequest) SetMobile(v string) *MobileOnlineTimeRequest {
	s.Mobile = &v
	return s
}

func (s *MobileOnlineTimeRequest) SetParamType(v string) *MobileOnlineTimeRequest {
	s.ParamType = &v
	return s
}

type MobileOnlineTimeResponseBody struct {
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
	// B506328A-D84B-4750-82C7-6A207C585CF1
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *MobileOnlineTimeResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s MobileOnlineTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MobileOnlineTimeResponseBody) GoString() string {
	return s.String()
}

func (s *MobileOnlineTimeResponseBody) SetCode(v string) *MobileOnlineTimeResponseBody {
	s.Code = &v
	return s
}

func (s *MobileOnlineTimeResponseBody) SetMessage(v string) *MobileOnlineTimeResponseBody {
	s.Message = &v
	return s
}

func (s *MobileOnlineTimeResponseBody) SetRequestId(v string) *MobileOnlineTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *MobileOnlineTimeResponseBody) SetResultObject(v *MobileOnlineTimeResponseBodyResultObject) *MobileOnlineTimeResponseBody {
	s.ResultObject = v
	return s
}

type MobileOnlineTimeResponseBodyResultObject struct {
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
	// 5
	TimeCode *string `json:"TimeCode,omitempty" xml:"TimeCode,omitempty"`
}

func (s MobileOnlineTimeResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s MobileOnlineTimeResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *MobileOnlineTimeResponseBodyResultObject) SetBizCode(v string) *MobileOnlineTimeResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *MobileOnlineTimeResponseBodyResultObject) SetIspName(v string) *MobileOnlineTimeResponseBodyResultObject {
	s.IspName = &v
	return s
}

func (s *MobileOnlineTimeResponseBodyResultObject) SetTimeCode(v string) *MobileOnlineTimeResponseBodyResultObject {
	s.TimeCode = &v
	return s
}

type MobileOnlineTimeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MobileOnlineTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MobileOnlineTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s MobileOnlineTimeResponse) GoString() string {
	return s.String()
}

func (s *MobileOnlineTimeResponse) SetHeaders(v map[string]*string) *MobileOnlineTimeResponse {
	s.Headers = v
	return s
}

func (s *MobileOnlineTimeResponse) SetStatusCode(v int32) *MobileOnlineTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *MobileOnlineTimeResponse) SetBody(v *MobileOnlineTimeResponseBody) *MobileOnlineTimeResponse {
	s.Body = v
	return s
}

type ModifyDeviceInfoRequest struct {
	// example:
	//
	// FACE_TEST
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// wd.6ziUffspAeW5FVYbaqmexR-1qwNjM
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 20190401
	ExpiredDay *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
	// example:
	//
	// 3iJ1AY$oHcu7mC69
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s ModifyDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceInfoRequest) SetBizType(v string) *ModifyDeviceInfoRequest {
	s.BizType = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetDeviceId(v string) *ModifyDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetDuration(v string) *ModifyDeviceInfoRequest {
	s.Duration = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetExpiredDay(v string) *ModifyDeviceInfoRequest {
	s.ExpiredDay = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetUserDeviceId(v string) *ModifyDeviceInfoRequest {
	s.UserDeviceId = &v
	return s
}

type ModifyDeviceInfoResponseBody struct {
	// example:
	//
	// 20190401
	BeginDay *string `json:"BeginDay,omitempty" xml:"BeginDay,omitempty"`
	// example:
	//
	// FACE_TEST
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// wd.6ziUffspAeW5FVYbaqmexR-1qwNjM
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 20200330
	ExpiredDay *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3iJ1AY$oHcu7mC69
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s ModifyDeviceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeviceInfoResponseBody) SetBeginDay(v string) *ModifyDeviceInfoResponseBody {
	s.BeginDay = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetBizType(v string) *ModifyDeviceInfoResponseBody {
	s.BizType = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetDeviceId(v string) *ModifyDeviceInfoResponseBody {
	s.DeviceId = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetExpiredDay(v string) *ModifyDeviceInfoResponseBody {
	s.ExpiredDay = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetRequestId(v string) *ModifyDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetUserDeviceId(v string) *ModifyDeviceInfoResponseBody {
	s.UserDeviceId = &v
	return s
}

type ModifyDeviceInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceInfoResponse) SetHeaders(v map[string]*string) *ModifyDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeviceInfoResponse) SetStatusCode(v int32) *ModifyDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDeviceInfoResponse) SetBody(v *ModifyDeviceInfoResponseBody) *ModifyDeviceInfoResponse {
	s.Body = v
	return s
}

type PageQueryWhiteListSettingRequest struct {
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// example:
	//
	// sha75b4e19a1ddda059b920757b0e12b
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1000000xxx
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1725379200000
	ValidEndDate *string `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	// example:
	//
	// 1725120000000
	ValidStartDate *string `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
}

func (s PageQueryWhiteListSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s PageQueryWhiteListSettingRequest) GoString() string {
	return s.String()
}

func (s *PageQueryWhiteListSettingRequest) SetCertNo(v string) *PageQueryWhiteListSettingRequest {
	s.CertNo = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetCertifyId(v string) *PageQueryWhiteListSettingRequest {
	s.CertifyId = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetCurrentPage(v int32) *PageQueryWhiteListSettingRequest {
	s.CurrentPage = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetPageSize(v int32) *PageQueryWhiteListSettingRequest {
	s.PageSize = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetSceneId(v int64) *PageQueryWhiteListSettingRequest {
	s.SceneId = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetServiceCode(v string) *PageQueryWhiteListSettingRequest {
	s.ServiceCode = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetStatus(v string) *PageQueryWhiteListSettingRequest {
	s.Status = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetValidEndDate(v string) *PageQueryWhiteListSettingRequest {
	s.ValidEndDate = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetValidStartDate(v string) *PageQueryWhiteListSettingRequest {
	s.ValidStartDate = &v
	return s
}

type PageQueryWhiteListSettingResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5A6229C0-E156-48E4-B6EC-0F528BDF60D2
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject []*PageQueryWhiteListSettingResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 28
	TotalItem *int32 `json:"TotalItem,omitempty" xml:"TotalItem,omitempty"`
	// example:
	//
	// 3
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s PageQueryWhiteListSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageQueryWhiteListSettingResponseBody) GoString() string {
	return s.String()
}

func (s *PageQueryWhiteListSettingResponseBody) SetCode(v string) *PageQueryWhiteListSettingResponseBody {
	s.Code = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetCurrentPage(v int32) *PageQueryWhiteListSettingResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetMessage(v string) *PageQueryWhiteListSettingResponseBody {
	s.Message = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetPageSize(v int32) *PageQueryWhiteListSettingResponseBody {
	s.PageSize = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetRequestId(v string) *PageQueryWhiteListSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetResultObject(v []*PageQueryWhiteListSettingResponseBodyResultObject) *PageQueryWhiteListSettingResponseBody {
	s.ResultObject = v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetSuccess(v bool) *PageQueryWhiteListSettingResponseBody {
	s.Success = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetTotalItem(v int32) *PageQueryWhiteListSettingResponseBody {
	s.TotalItem = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetTotalPage(v int32) *PageQueryWhiteListSettingResponseBody {
	s.TotalPage = &v
	return s
}

type PageQueryWhiteListSettingResponseBodyResultObject struct {
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// example:
	//
	// sha43d9cabd52d370d9f4cca9468f71e
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// 2024-08-30 14:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-08-30 14:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 234822
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 1000000332
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-09-02 13:57:51
	ValidEndDate *string `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	// example:
	//
	// 2024-08-30 13:57:51
	ValidStartDate *string `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
}

func (s PageQueryWhiteListSettingResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s PageQueryWhiteListSettingResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetCertNo(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.CertNo = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetCertifyId(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetGmtCreate(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetGmtModified(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetId(v int64) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetRemark(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.Remark = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetSceneId(v int64) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.SceneId = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetServiceCode(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.ServiceCode = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetStatus(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetValidEndDate(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.ValidEndDate = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetValidStartDate(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.ValidStartDate = &v
	return s
}

type PageQueryWhiteListSettingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageQueryWhiteListSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageQueryWhiteListSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s PageQueryWhiteListSettingResponse) GoString() string {
	return s.String()
}

func (s *PageQueryWhiteListSettingResponse) SetHeaders(v map[string]*string) *PageQueryWhiteListSettingResponse {
	s.Headers = v
	return s
}

func (s *PageQueryWhiteListSettingResponse) SetStatusCode(v int32) *PageQueryWhiteListSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *PageQueryWhiteListSettingResponse) SetBody(v *PageQueryWhiteListSettingResponseBody) *PageQueryWhiteListSettingResponse {
	s.Body = v
	return s
}

type RemoveWhiteListSettingRequest struct {
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s RemoveWhiteListSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveWhiteListSettingRequest) GoString() string {
	return s.String()
}

func (s *RemoveWhiteListSettingRequest) SetIds(v []*int64) *RemoveWhiteListSettingRequest {
	s.Ids = v
	return s
}

func (s *RemoveWhiteListSettingRequest) SetServiceCode(v string) *RemoveWhiteListSettingRequest {
	s.ServiceCode = &v
	return s
}

type RemoveWhiteListSettingShrinkRequest struct {
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s RemoveWhiteListSettingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveWhiteListSettingShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveWhiteListSettingShrinkRequest) SetIdsShrink(v string) *RemoveWhiteListSettingShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *RemoveWhiteListSettingShrinkRequest) SetServiceCode(v string) *RemoveWhiteListSettingShrinkRequest {
	s.ServiceCode = &v
	return s
}

type RemoveWhiteListSettingResponseBody struct {
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveWhiteListSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveWhiteListSettingResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveWhiteListSettingResponseBody) SetCode(v string) *RemoveWhiteListSettingResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveWhiteListSettingResponseBody) SetMessage(v string) *RemoveWhiteListSettingResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveWhiteListSettingResponseBody) SetRequestId(v string) *RemoveWhiteListSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveWhiteListSettingResponseBody) SetResultObject(v bool) *RemoveWhiteListSettingResponseBody {
	s.ResultObject = &v
	return s
}

func (s *RemoveWhiteListSettingResponseBody) SetSuccess(v bool) *RemoveWhiteListSettingResponseBody {
	s.Success = &v
	return s
}

type RemoveWhiteListSettingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveWhiteListSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveWhiteListSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveWhiteListSettingResponse) GoString() string {
	return s.String()
}

func (s *RemoveWhiteListSettingResponse) SetHeaders(v map[string]*string) *RemoveWhiteListSettingResponse {
	s.Headers = v
	return s
}

func (s *RemoveWhiteListSettingResponse) SetStatusCode(v int32) *RemoveWhiteListSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveWhiteListSettingResponse) SetBody(v *RemoveWhiteListSettingResponseBody) *RemoveWhiteListSettingResponse {
	s.Body = v
	return s
}

type Vehicle5ItemQueryRequest struct {
	// example:
	//
	// normal
	ParamType  *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
	// example:
	//
	// 02
	VehicleType *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
}

func (s Vehicle5ItemQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s Vehicle5ItemQueryRequest) GoString() string {
	return s.String()
}

func (s *Vehicle5ItemQueryRequest) SetParamType(v string) *Vehicle5ItemQueryRequest {
	s.ParamType = &v
	return s
}

func (s *Vehicle5ItemQueryRequest) SetVehicleNum(v string) *Vehicle5ItemQueryRequest {
	s.VehicleNum = &v
	return s
}

func (s *Vehicle5ItemQueryRequest) SetVehicleType(v string) *Vehicle5ItemQueryRequest {
	s.VehicleType = &v
	return s
}

type Vehicle5ItemQueryResponseBody struct {
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0D******
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *Vehicle5ItemQueryResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Vehicle5ItemQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Vehicle5ItemQueryResponseBody) GoString() string {
	return s.String()
}

func (s *Vehicle5ItemQueryResponseBody) SetCode(v string) *Vehicle5ItemQueryResponseBody {
	s.Code = &v
	return s
}

func (s *Vehicle5ItemQueryResponseBody) SetMessage(v string) *Vehicle5ItemQueryResponseBody {
	s.Message = &v
	return s
}

func (s *Vehicle5ItemQueryResponseBody) SetRequestId(v string) *Vehicle5ItemQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *Vehicle5ItemQueryResponseBody) SetResultObject(v *Vehicle5ItemQueryResponseBodyResultObject) *Vehicle5ItemQueryResponseBody {
	s.ResultObject = v
	return s
}

type Vehicle5ItemQueryResponseBodyResultObject struct {
	// example:
	//
	// 1
	BizCode     *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	VehicleInfo *string `json:"VehicleInfo,omitempty" xml:"VehicleInfo,omitempty"`
}

func (s Vehicle5ItemQueryResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s Vehicle5ItemQueryResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Vehicle5ItemQueryResponseBodyResultObject) SetBizCode(v string) *Vehicle5ItemQueryResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Vehicle5ItemQueryResponseBodyResultObject) SetVehicleInfo(v string) *Vehicle5ItemQueryResponseBodyResultObject {
	s.VehicleInfo = &v
	return s
}

type Vehicle5ItemQueryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Vehicle5ItemQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Vehicle5ItemQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s Vehicle5ItemQueryResponse) GoString() string {
	return s.String()
}

func (s *Vehicle5ItemQueryResponse) SetHeaders(v map[string]*string) *Vehicle5ItemQueryResponse {
	s.Headers = v
	return s
}

func (s *Vehicle5ItemQueryResponse) SetStatusCode(v int32) *Vehicle5ItemQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *Vehicle5ItemQueryResponse) SetBody(v *Vehicle5ItemQueryResponseBody) *Vehicle5ItemQueryResponse {
	s.Body = v
	return s
}

type VehicleInsureQueryRequest struct {
	// example:
	//
	// normal
	ParamType  *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
	// example:
	//
	// 02
	VehicleType *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
	// example:
	//
	// LB**************
	Vin *string `json:"Vin,omitempty" xml:"Vin,omitempty"`
}

func (s VehicleInsureQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s VehicleInsureQueryRequest) GoString() string {
	return s.String()
}

func (s *VehicleInsureQueryRequest) SetParamType(v string) *VehicleInsureQueryRequest {
	s.ParamType = &v
	return s
}

func (s *VehicleInsureQueryRequest) SetVehicleNum(v string) *VehicleInsureQueryRequest {
	s.VehicleNum = &v
	return s
}

func (s *VehicleInsureQueryRequest) SetVehicleType(v string) *VehicleInsureQueryRequest {
	s.VehicleType = &v
	return s
}

func (s *VehicleInsureQueryRequest) SetVin(v string) *VehicleInsureQueryRequest {
	s.Vin = &v
	return s
}

type VehicleInsureQueryResponseBody struct {
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
	// 5A6229C0-E156-48E4-B6EC-0F52********
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *VehicleInsureQueryResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s VehicleInsureQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VehicleInsureQueryResponseBody) GoString() string {
	return s.String()
}

func (s *VehicleInsureQueryResponseBody) SetCode(v string) *VehicleInsureQueryResponseBody {
	s.Code = &v
	return s
}

func (s *VehicleInsureQueryResponseBody) SetMessage(v string) *VehicleInsureQueryResponseBody {
	s.Message = &v
	return s
}

func (s *VehicleInsureQueryResponseBody) SetRequestId(v string) *VehicleInsureQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *VehicleInsureQueryResponseBody) SetResultObject(v *VehicleInsureQueryResponseBodyResultObject) *VehicleInsureQueryResponseBody {
	s.ResultObject = v
	return s
}

type VehicleInsureQueryResponseBodyResultObject struct {
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {
	//
	//     "firstInsuranceDate": "****-**-**",
	//
	//     "lastInsuranceDate": "****-**",
	//
	//     "latestInsuranceDate": "****-**",
	//
	//     "latestInsuranceDateStart": "****-**"
	//
	//   }
	VehicleInfo *string `json:"VehicleInfo,omitempty" xml:"VehicleInfo,omitempty"`
}

func (s VehicleInsureQueryResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s VehicleInsureQueryResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *VehicleInsureQueryResponseBodyResultObject) SetBizCode(v string) *VehicleInsureQueryResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *VehicleInsureQueryResponseBodyResultObject) SetVehicleInfo(v string) *VehicleInsureQueryResponseBodyResultObject {
	s.VehicleInfo = &v
	return s
}

type VehicleInsureQueryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VehicleInsureQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VehicleInsureQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s VehicleInsureQueryResponse) GoString() string {
	return s.String()
}

func (s *VehicleInsureQueryResponse) SetHeaders(v map[string]*string) *VehicleInsureQueryResponse {
	s.Headers = v
	return s
}

func (s *VehicleInsureQueryResponse) SetStatusCode(v int32) *VehicleInsureQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *VehicleInsureQueryResponse) SetBody(v *VehicleInsureQueryResponseBody) *VehicleInsureQueryResponse {
	s.Body = v
	return s
}

type VehicleMetaVerifyRequest struct {
	// example:
	//
	// 4****************1
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// example:
	//
	// normal
	ParamType  *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
	// example:
	//
	// 02
	VehicleType *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
	// example:
	//
	// VEHICLE_2_META
	VerifyMetaType *string `json:"VerifyMetaType,omitempty" xml:"VerifyMetaType,omitempty"`
}

func (s VehicleMetaVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s VehicleMetaVerifyRequest) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyRequest) SetIdentifyNum(v string) *VehicleMetaVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *VehicleMetaVerifyRequest) SetParamType(v string) *VehicleMetaVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *VehicleMetaVerifyRequest) SetUserName(v string) *VehicleMetaVerifyRequest {
	s.UserName = &v
	return s
}

func (s *VehicleMetaVerifyRequest) SetVehicleNum(v string) *VehicleMetaVerifyRequest {
	s.VehicleNum = &v
	return s
}

func (s *VehicleMetaVerifyRequest) SetVehicleType(v string) *VehicleMetaVerifyRequest {
	s.VehicleType = &v
	return s
}

func (s *VehicleMetaVerifyRequest) SetVerifyMetaType(v string) *VehicleMetaVerifyRequest {
	s.VerifyMetaType = &v
	return s
}

type VehicleMetaVerifyResponseBody struct {
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
	// 8FC3D6AC-9FED-4311-8DA7-C4BF4*****
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *VehicleMetaVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s VehicleMetaVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VehicleMetaVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyResponseBody) SetCode(v string) *VehicleMetaVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *VehicleMetaVerifyResponseBody) SetMessage(v string) *VehicleMetaVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *VehicleMetaVerifyResponseBody) SetRequestId(v string) *VehicleMetaVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *VehicleMetaVerifyResponseBody) SetResultObject(v *VehicleMetaVerifyResponseBodyResultObject) *VehicleMetaVerifyResponseBody {
	s.ResultObject = v
	return s
}

type VehicleMetaVerifyResponseBodyResultObject struct {
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
}

func (s VehicleMetaVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s VehicleMetaVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyResponseBodyResultObject) SetBizCode(v string) *VehicleMetaVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

type VehicleMetaVerifyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VehicleMetaVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VehicleMetaVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s VehicleMetaVerifyResponse) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyResponse) SetHeaders(v map[string]*string) *VehicleMetaVerifyResponse {
	s.Headers = v
	return s
}

func (s *VehicleMetaVerifyResponse) SetStatusCode(v int32) *VehicleMetaVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *VehicleMetaVerifyResponse) SetBody(v *VehicleMetaVerifyResponseBody) *VehicleMetaVerifyResponse {
	s.Body = v
	return s
}

type VehicleMetaVerifyV2Request struct {
	// example:
	//
	// 4****************1
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// example:
	//
	// normal
	ParamType  *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
	// example:
	//
	// 02
	VehicleType *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
	// example:
	//
	// VEHICLE_3_META
	VerifyMetaType *string `json:"VerifyMetaType,omitempty" xml:"VerifyMetaType,omitempty"`
}

func (s VehicleMetaVerifyV2Request) String() string {
	return tea.Prettify(s)
}

func (s VehicleMetaVerifyV2Request) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyV2Request) SetIdentifyNum(v string) *VehicleMetaVerifyV2Request {
	s.IdentifyNum = &v
	return s
}

func (s *VehicleMetaVerifyV2Request) SetParamType(v string) *VehicleMetaVerifyV2Request {
	s.ParamType = &v
	return s
}

func (s *VehicleMetaVerifyV2Request) SetUserName(v string) *VehicleMetaVerifyV2Request {
	s.UserName = &v
	return s
}

func (s *VehicleMetaVerifyV2Request) SetVehicleNum(v string) *VehicleMetaVerifyV2Request {
	s.VehicleNum = &v
	return s
}

func (s *VehicleMetaVerifyV2Request) SetVehicleType(v string) *VehicleMetaVerifyV2Request {
	s.VehicleType = &v
	return s
}

func (s *VehicleMetaVerifyV2Request) SetVerifyMetaType(v string) *VehicleMetaVerifyV2Request {
	s.VerifyMetaType = &v
	return s
}

type VehicleMetaVerifyV2ResponseBody struct {
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
	// 5A6229C0-E156-48E4-B6EC-0F528B******
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *VehicleMetaVerifyV2ResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s VehicleMetaVerifyV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VehicleMetaVerifyV2ResponseBody) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyV2ResponseBody) SetCode(v string) *VehicleMetaVerifyV2ResponseBody {
	s.Code = &v
	return s
}

func (s *VehicleMetaVerifyV2ResponseBody) SetMessage(v string) *VehicleMetaVerifyV2ResponseBody {
	s.Message = &v
	return s
}

func (s *VehicleMetaVerifyV2ResponseBody) SetRequestId(v string) *VehicleMetaVerifyV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *VehicleMetaVerifyV2ResponseBody) SetResultObject(v *VehicleMetaVerifyV2ResponseBodyResultObject) *VehicleMetaVerifyV2ResponseBody {
	s.ResultObject = v
	return s
}

type VehicleMetaVerifyV2ResponseBodyResultObject struct {
	// example:
	//
	// 1
	BizCode     *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	VehicleInfo *string `json:"VehicleInfo,omitempty" xml:"VehicleInfo,omitempty"`
}

func (s VehicleMetaVerifyV2ResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s VehicleMetaVerifyV2ResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyV2ResponseBodyResultObject) SetBizCode(v string) *VehicleMetaVerifyV2ResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *VehicleMetaVerifyV2ResponseBodyResultObject) SetVehicleInfo(v string) *VehicleMetaVerifyV2ResponseBodyResultObject {
	s.VehicleInfo = &v
	return s
}

type VehicleMetaVerifyV2Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VehicleMetaVerifyV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VehicleMetaVerifyV2Response) String() string {
	return tea.Prettify(s)
}

func (s VehicleMetaVerifyV2Response) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyV2Response) SetHeaders(v map[string]*string) *VehicleMetaVerifyV2Response {
	s.Headers = v
	return s
}

func (s *VehicleMetaVerifyV2Response) SetStatusCode(v int32) *VehicleMetaVerifyV2Response {
	s.StatusCode = &v
	return s
}

func (s *VehicleMetaVerifyV2Response) SetBody(v *VehicleMetaVerifyV2ResponseBody) *VehicleMetaVerifyV2Response {
	s.Body = v
	return s
}

type VehicleQueryRequest struct {
	// example:
	//
	// normal
	ParamType  *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
	// example:
	//
	// 02
	VehicleType *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
}

func (s VehicleQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s VehicleQueryRequest) GoString() string {
	return s.String()
}

func (s *VehicleQueryRequest) SetParamType(v string) *VehicleQueryRequest {
	s.ParamType = &v
	return s
}

func (s *VehicleQueryRequest) SetVehicleNum(v string) *VehicleQueryRequest {
	s.VehicleNum = &v
	return s
}

func (s *VehicleQueryRequest) SetVehicleType(v string) *VehicleQueryRequest {
	s.VehicleType = &v
	return s
}

type VehicleQueryResponseBody struct {
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
	// D6163397-15C5-419C-9ACC-B7C83*******
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *VehicleQueryResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s VehicleQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VehicleQueryResponseBody) GoString() string {
	return s.String()
}

func (s *VehicleQueryResponseBody) SetCode(v string) *VehicleQueryResponseBody {
	s.Code = &v
	return s
}

func (s *VehicleQueryResponseBody) SetMessage(v string) *VehicleQueryResponseBody {
	s.Message = &v
	return s
}

func (s *VehicleQueryResponseBody) SetRequestId(v string) *VehicleQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *VehicleQueryResponseBody) SetResultObject(v *VehicleQueryResponseBodyResultObject) *VehicleQueryResponseBody {
	s.ResultObject = v
	return s
}

type VehicleQueryResponseBodyResultObject struct {
	// example:
	//
	// 1
	BizCode     *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	VehicleInfo *string `json:"VehicleInfo,omitempty" xml:"VehicleInfo,omitempty"`
}

func (s VehicleQueryResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s VehicleQueryResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *VehicleQueryResponseBodyResultObject) SetBizCode(v string) *VehicleQueryResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *VehicleQueryResponseBodyResultObject) SetVehicleInfo(v string) *VehicleQueryResponseBodyResultObject {
	s.VehicleInfo = &v
	return s
}

type VehicleQueryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VehicleQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VehicleQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s VehicleQueryResponse) GoString() string {
	return s.String()
}

func (s *VehicleQueryResponse) SetHeaders(v map[string]*string) *VehicleQueryResponse {
	s.Headers = v
	return s
}

func (s *VehicleQueryResponse) SetStatusCode(v int32) *VehicleQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *VehicleQueryResponse) SetBody(v *VehicleQueryResponseBody) *VehicleQueryResponse {
	s.Body = v
	return s
}

type VerifyMaterialRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 39ecf51e-2f81-4dc5-90ee-ff86125b****
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RPMinTest
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example3.jpg
	IdCardBackImageUrl *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example2.jpg
	IdCardFrontImageUrl *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 02343218901123****
	IdCardNumber *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 54sdj
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s VerifyMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialRequest) GoString() string {
	return s.String()
}

func (s *VerifyMaterialRequest) SetBizId(v string) *VerifyMaterialRequest {
	s.BizId = &v
	return s
}

func (s *VerifyMaterialRequest) SetBizType(v string) *VerifyMaterialRequest {
	s.BizType = &v
	return s
}

func (s *VerifyMaterialRequest) SetFaceImageUrl(v string) *VerifyMaterialRequest {
	s.FaceImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardBackImageUrl(v string) *VerifyMaterialRequest {
	s.IdCardBackImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardFrontImageUrl(v string) *VerifyMaterialRequest {
	s.IdCardFrontImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardNumber(v string) *VerifyMaterialRequest {
	s.IdCardNumber = &v
	return s
}

func (s *VerifyMaterialRequest) SetName(v string) *VerifyMaterialRequest {
	s.Name = &v
	return s
}

func (s *VerifyMaterialRequest) SetUserId(v string) *VerifyMaterialRequest {
	s.UserId = &v
	return s
}

type VerifyMaterialResponseBody struct {
	// example:
	//
	// 97
	AuthorityComparisionScore *float32 `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty"`
	// example:
	//
	// 97
	IdCardFaceComparisonScore *float32                            `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	Material                  *VerifyMaterialResponseBodyMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	VerifyStatus *int32 `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
	// example:
	//
	// c302c0797679457685410ee51a5ba375
	VerifyToken *string `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
}

func (s VerifyMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseBody) SetAuthorityComparisionScore(v float32) *VerifyMaterialResponseBody {
	s.AuthorityComparisionScore = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetIdCardFaceComparisonScore(v float32) *VerifyMaterialResponseBody {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetMaterial(v *VerifyMaterialResponseBodyMaterial) *VerifyMaterialResponseBody {
	s.Material = v
	return s
}

func (s *VerifyMaterialResponseBody) SetRequestId(v string) *VerifyMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetVerifyStatus(v int32) *VerifyMaterialResponseBody {
	s.VerifyStatus = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetVerifyToken(v string) *VerifyMaterialResponseBody {
	s.VerifyToken = &v
	return s
}

type VerifyMaterialResponseBodyMaterial struct {
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/face-global-example.jpg
	FaceGlobalUrl *string `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	// example:
	//
	// false
	FaceMask *string `json:"FaceMask,omitempty" xml:"FaceMask,omitempty"`
	// example:
	//
	// NORMAL
	FaceQuality *string                                       `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	IdCardInfo  *VerifyMaterialResponseBodyMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	IdCardName  *string                                       `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	// example:
	//
	// 02343218901123****
	IdCardNumber *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
}

func (s VerifyMaterialResponseBodyMaterial) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseBodyMaterial) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseBodyMaterial) SetFaceGlobalUrl(v string) *VerifyMaterialResponseBodyMaterial {
	s.FaceGlobalUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetFaceImageUrl(v string) *VerifyMaterialResponseBodyMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetFaceMask(v string) *VerifyMaterialResponseBodyMaterial {
	s.FaceMask = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetFaceQuality(v string) *VerifyMaterialResponseBodyMaterial {
	s.FaceQuality = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetIdCardInfo(v *VerifyMaterialResponseBodyMaterialIdCardInfo) *VerifyMaterialResponseBodyMaterial {
	s.IdCardInfo = v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetIdCardName(v string) *VerifyMaterialResponseBodyMaterial {
	s.IdCardName = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetIdCardNumber(v string) *VerifyMaterialResponseBodyMaterial {
	s.IdCardNumber = &v
	return s
}

type VerifyMaterialResponseBodyMaterialIdCardInfo struct {
	Address   *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Authority *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example3.jpg
	BackImageUrl *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty"`
	// example:
	//
	// 19900101
	Birth *string `json:"Birth,omitempty" xml:"Birth,omitempty"`
	// example:
	//
	// 20201101
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example2.jpg
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	// example:
	//
	// 02343218901123****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 20201101
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s VerifyMaterialResponseBodyMaterialIdCardInfo) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseBodyMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetAddress(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetAuthority(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetBackImageUrl(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetBirth(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetEndDate(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetFrontImageUrl(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetName(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetNationality(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetNumber(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetStartDate(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

type VerifyMaterialResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponse) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponse) SetHeaders(v map[string]*string) *VerifyMaterialResponse {
	s.Headers = v
	return s
}

func (s *VerifyMaterialResponse) SetStatusCode(v int32) *VerifyMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyMaterialResponse) SetBody(v *VerifyMaterialResponseBody) *VerifyMaterialResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudauth"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 新增AIGC人脸检测能力
//
// @param request - AIGCFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AIGCFaceVerifyResponse
func (client *Client) AIGCFaceVerifyWithOptions(request *AIGCFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *AIGCFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceContrastPictureUrl)) {
		query["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssObjectName)) {
		query["OssObjectName"] = request.OssObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.OuterOrderNo)) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceContrastPicture)) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AIGCFaceVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AIGCFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增AIGC人脸检测能力
//
// @param request - AIGCFaceVerifyRequest
//
// @return AIGCFaceVerifyResponse
func (client *Client) AIGCFaceVerify(request *AIGCFaceVerifyRequest) (_result *AIGCFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AIGCFaceVerifyResponse{}
	_body, _err := client.AIGCFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 银行卡要素核验接口
//
// @param request - BankMetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BankMetaVerifyResponse
func (client *Client) BankMetaVerifyWithOptions(request *BankMetaVerifyRequest, runtime *util.RuntimeOptions) (_result *BankMetaVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BankCard)) {
		query["BankCard"] = request.BankCard
	}

	if !tea.BoolValue(util.IsUnset(request.IdentifyNum)) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityType)) {
		query["IdentityType"] = request.IdentityType
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		query["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyMode)) {
		query["VerifyMode"] = request.VerifyMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BankMetaVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BankMetaVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 银行卡要素核验接口
//
// @param request - BankMetaVerifyRequest
//
// @return BankMetaVerifyResponse
func (client *Client) BankMetaVerify(request *BankMetaVerifyRequest) (_result *BankMetaVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BankMetaVerifyResponse{}
	_body, _err := client.BankMetaVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CompareFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareFaceVerifyResponse
func (client *Client) CompareFaceVerifyWithOptions(request *CompareFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *CompareFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.OuterOrderNo)) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCertifyId)) {
		body["SourceCertifyId"] = request.SourceCertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFaceContrastPicture)) {
		body["SourceFaceContrastPicture"] = request.SourceFaceContrastPicture
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFaceContrastPictureUrl)) {
		body["SourceFaceContrastPictureUrl"] = request.SourceFaceContrastPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOssBucketName)) {
		body["SourceOssBucketName"] = request.SourceOssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOssObjectName)) {
		body["SourceOssObjectName"] = request.SourceOssObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCertifyId)) {
		body["TargetCertifyId"] = request.TargetCertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFaceContrastPicture)) {
		body["TargetFaceContrastPicture"] = request.TargetFaceContrastPicture
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFaceContrastPictureUrl)) {
		body["TargetFaceContrastPictureUrl"] = request.TargetFaceContrastPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TargetOssBucketName)) {
		body["TargetOssBucketName"] = request.TargetOssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetOssObjectName)) {
		body["TargetOssObjectName"] = request.TargetOssObjectName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CompareFaceVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CompareFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CompareFaceVerifyRequest
//
// @return CompareFaceVerifyResponse
func (client *Client) CompareFaceVerify(request *CompareFaceVerifyRequest) (_result *CompareFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareFaceVerifyResponse{}
	_body, _err := client.CompareFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CompareFacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareFacesResponse
func (client *Client) CompareFacesWithOptions(request *CompareFacesRequest, runtime *util.RuntimeOptions) (_result *CompareFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceImageType)) {
		body["SourceImageType"] = request.SourceImageType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceImageValue)) {
		body["SourceImageValue"] = request.SourceImageValue
	}

	if !tea.BoolValue(util.IsUnset(request.TargetImageType)) {
		body["TargetImageType"] = request.TargetImageType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetImageValue)) {
		body["TargetImageValue"] = request.TargetImageValue
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CompareFaces"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CompareFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CompareFacesRequest
//
// @return CompareFacesResponse
func (client *Client) CompareFaces(request *CompareFacesRequest) (_result *CompareFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareFacesResponse{}
	_body, _err := client.CompareFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ContrastFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContrastFaceVerifyResponse
func (client *Client) ContrastFaceVerifyWithOptions(request *ContrastFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *ContrastFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Model)) {
		query["Model"] = request.Model
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		body["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.CertNo)) {
		body["CertNo"] = request.CertNo
	}

	if !tea.BoolValue(util.IsUnset(request.CertType)) {
		body["CertType"] = request.CertType
	}

	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		body["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceToken)) {
		body["DeviceToken"] = request.DeviceToken
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptType)) {
		body["EncryptType"] = request.EncryptType
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastFile)) {
		body["FaceContrastFile"] = request.FaceContrastFile
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastPicture)) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastPictureUrl)) {
		body["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		body["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		body["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssObjectName)) {
		body["OssObjectName"] = request.OssObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.OuterOrderNo)) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ContrastFaceVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ContrastFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ContrastFaceVerifyRequest
//
// @return ContrastFaceVerifyResponse
func (client *Client) ContrastFaceVerify(request *ContrastFaceVerifyRequest) (_result *ContrastFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContrastFaceVerifyResponse{}
	_body, _err := client.ContrastFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContrastFaceVerifyAdvance(request *ContrastFaceVerifyAdvanceRequest, runtime *util.RuntimeOptions) (_result *ContrastFaceVerifyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("Cloudauth"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	contrastFaceVerifyReq := &ContrastFaceVerifyRequest{}
	openapiutil.Convert(request, contrastFaceVerifyReq)
	if !tea.BoolValue(util.IsUnset(request.FaceContrastFileObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.FaceContrastFileObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		contrastFaceVerifyReq.FaceContrastFile = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	contrastFaceVerifyResp, _err := client.ContrastFaceVerifyWithOptions(contrastFaceVerifyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = contrastFaceVerifyResp
	return _result, _err
}

// @param request - CreateAuthKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAuthKeyResponse
func (client *Client) CreateAuthKeyWithOptions(request *CreateAuthKeyRequest, runtime *util.RuntimeOptions) (_result *CreateAuthKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthYears)) {
		query["AuthYears"] = request.AuthYears
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Test)) {
		query["Test"] = request.Test
	}

	if !tea.BoolValue(util.IsUnset(request.UserDeviceId)) {
		query["UserDeviceId"] = request.UserDeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAuthKey"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAuthKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAuthKeyRequest
//
// @return CreateAuthKeyResponse
func (client *Client) CreateAuthKey(request *CreateAuthKeyRequest) (_result *CreateAuthKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAuthKeyResponse{}
	_body, _err := client.CreateAuthKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateVerifySettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVerifySettingResponse
func (client *Client) CreateVerifySettingWithOptions(request *CreateVerifySettingRequest, runtime *util.RuntimeOptions) (_result *CreateVerifySettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		query["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.GuideStep)) {
		query["GuideStep"] = request.GuideStep
	}

	if !tea.BoolValue(util.IsUnset(request.PrivacyStep)) {
		query["PrivacyStep"] = request.PrivacyStep
	}

	if !tea.BoolValue(util.IsUnset(request.ResultStep)) {
		query["ResultStep"] = request.ResultStep
	}

	if !tea.BoolValue(util.IsUnset(request.Solution)) {
		query["Solution"] = request.Solution
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVerifySetting"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVerifySettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateVerifySettingRequest
//
// @return CreateVerifySettingResponse
func (client *Client) CreateVerifySetting(request *CreateVerifySettingRequest) (_result *CreateVerifySettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVerifySettingResponse{}
	_body, _err := client.CreateVerifySettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 凭证核验
//
// @param tmpReq - CredentialVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CredentialVerifyResponse
func (client *Client) CredentialVerifyWithOptions(tmpReq *CredentialVerifyRequest, runtime *util.RuntimeOptions) (_result *CredentialVerifyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CredentialVerifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MerchantDetail)) {
		request.MerchantDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MerchantDetail, tea.String("MerchantDetail"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertNum)) {
		query["CertNum"] = request.CertNum
	}

	if !tea.BoolValue(util.IsUnset(request.CredName)) {
		query["CredName"] = request.CredName
	}

	if !tea.BoolValue(util.IsUnset(request.CredType)) {
		query["CredType"] = request.CredType
	}

	if !tea.BoolValue(util.IsUnset(request.IdentifyNum)) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IsCheck)) {
		query["IsCheck"] = request.IsCheck
	}

	if !tea.BoolValue(util.IsUnset(request.IsOCR)) {
		query["IsOCR"] = request.IsOCR
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantDetailShrink)) {
		query["MerchantDetail"] = request.MerchantDetailShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		query["MerchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		query["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.PromptModel)) {
		query["PromptModel"] = request.PromptModel
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageContext)) {
		body["ImageContext"] = request.ImageContext
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CredentialVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CredentialVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 凭证核验
//
// @param request - CredentialVerifyRequest
//
// @return CredentialVerifyResponse
func (client *Client) CredentialVerify(request *CredentialVerifyRequest) (_result *CredentialVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CredentialVerifyResponse{}
	_body, _err := client.CredentialVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人脸凭证核验服务
//
// @param request - DeepfakeDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeepfakeDetectResponse
func (client *Client) DeepfakeDetectWithOptions(request *DeepfakeDetectRequest, runtime *util.RuntimeOptions) (_result *DeepfakeDetectResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OuterOrderNo)) {
		query["OuterOrderNo"] = request.OuterOrderNo
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
		Action:      tea.String("DeepfakeDetect"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeepfakeDetectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人脸凭证核验服务
//
// @param request - DeepfakeDetectRequest
//
// @return DeepfakeDetectResponse
func (client *Client) DeepfakeDetect(request *DeepfakeDetectRequest) (_result *DeepfakeDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeepfakeDetectResponse{}
	_body, _err := client.DeepfakeDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 金融级服务敏感数据删除接口
//
// @param request - DeleteFaceVerifyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceVerifyResultResponse
func (client *Client) DeleteFaceVerifyResultWithOptions(request *DeleteFaceVerifyResultRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceVerifyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		query["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteAfterQuery)) {
		query["DeleteAfterQuery"] = request.DeleteAfterQuery
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFaceVerifyResult"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFaceVerifyResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 金融级服务敏感数据删除接口
//
// @param request - DeleteFaceVerifyResultRequest
//
// @return DeleteFaceVerifyResultResponse
func (client *Client) DeleteFaceVerifyResult(request *DeleteFaceVerifyResultRequest) (_result *DeleteFaceVerifyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceVerifyResultResponse{}
	_body, _err := client.DeleteFaceVerifyResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceInfoResponse
func (client *Client) DescribeDeviceInfoWithOptions(request *DescribeDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredEndDay)) {
		query["ExpiredEndDay"] = request.ExpiredEndDay
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredStartDay)) {
		query["ExpiredStartDay"] = request.ExpiredStartDay
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserDeviceId)) {
		query["UserDeviceId"] = request.UserDeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceInfo"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDeviceInfoRequest
//
// @return DescribeDeviceInfoResponse
func (client *Client) DescribeDeviceInfo(request *DescribeDeviceInfoRequest) (_result *DescribeDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.DescribeDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 金融级人脸保镖服务
//
// @param request - DescribeFaceGuardRiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaceGuardRiskResponse
func (client *Client) DescribeFaceGuardRiskWithOptions(request *DescribeFaceGuardRiskRequest, runtime *util.RuntimeOptions) (_result *DescribeFaceGuardRiskResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OuterOrderNo)) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFaceGuardRisk"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFaceGuardRiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 金融级人脸保镖服务
//
// @param request - DescribeFaceGuardRiskRequest
//
// @return DescribeFaceGuardRiskResponse
func (client *Client) DescribeFaceGuardRisk(request *DescribeFaceGuardRiskRequest) (_result *DescribeFaceGuardRiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaceGuardRiskResponse{}
	_body, _err := client.DescribeFaceGuardRiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaceVerifyResponse
func (client *Client) DescribeFaceVerifyWithOptions(request *DescribeFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *DescribeFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		query["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.PictureReturnType)) {
		query["PictureReturnType"] = request.PictureReturnType
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFaceVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFaceVerifyRequest
//
// @return DescribeFaceVerifyResponse
func (client *Client) DescribeFaceVerify(request *DescribeFaceVerifyRequest) (_result *DescribeFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaceVerifyResponse{}
	_body, _err := client.DescribeFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeOssUploadTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssUploadTokenResponse
func (client *Client) DescribeOssUploadTokenWithOptions(runtime *util.RuntimeOptions) (_result *DescribeOssUploadTokenResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssUploadToken"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssUploadTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return DescribeOssUploadTokenResponse
func (client *Client) DescribeOssUploadToken() (_result *DescribeOssUploadTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssUploadTokenResponse{}
	_body, _err := client.DescribeOssUploadTokenWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Open API新增金融级数据统计API
//
// @param request - DescribePageFaceVerifyDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePageFaceVerifyDataResponse
func (client *Client) DescribePageFaceVerifyDataWithOptions(request *DescribePageFaceVerifyDataRequest, runtime *util.RuntimeOptions) (_result *DescribePageFaceVerifyDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePageFaceVerifyData"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePageFaceVerifyDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Open API新增金融级数据统计API
//
// @param request - DescribePageFaceVerifyDataRequest
//
// @return DescribePageFaceVerifyDataResponse
func (client *Client) DescribePageFaceVerifyData(request *DescribePageFaceVerifyDataRequest) (_result *DescribePageFaceVerifyDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePageFaceVerifyDataResponse{}
	_body, _err := client.DescribePageFaceVerifyDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeSmartStatisticsPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSmartStatisticsPageListResponse
func (client *Client) DescribeSmartStatisticsPageListWithOptions(request *DescribeSmartStatisticsPageListRequest, runtime *util.RuntimeOptions) (_result *DescribeSmartStatisticsPageListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSmartStatisticsPageList"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSmartStatisticsPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeSmartStatisticsPageListRequest
//
// @return DescribeSmartStatisticsPageListResponse
func (client *Client) DescribeSmartStatisticsPageList(request *DescribeSmartStatisticsPageListRequest) (_result *DescribeSmartStatisticsPageListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSmartStatisticsPageListResponse{}
	_body, _err := client.DescribeSmartStatisticsPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeVerifyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyResultResponse
func (client *Client) DescribeVerifyResultWithOptions(request *DescribeVerifyResultRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVerifyResult"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVerifyResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeVerifyResultRequest
//
// @return DescribeVerifyResultResponse
func (client *Client) DescribeVerifyResult(request *DescribeVerifyResultRequest) (_result *DescribeVerifyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifyResultResponse{}
	_body, _err := client.DescribeVerifyResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeVerifySDKRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifySDKResponse
func (client *Client) DescribeVerifySDKWithOptions(request *DescribeVerifySDKRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifySDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVerifySDK"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVerifySDKResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeVerifySDKRequest
//
// @return DescribeVerifySDKResponse
func (client *Client) DescribeVerifySDK(request *DescribeVerifySDKRequest) (_result *DescribeVerifySDKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifySDKResponse{}
	_body, _err := client.DescribeVerifySDKWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeVerifyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyTokenResponse
func (client *Client) DescribeVerifyTokenWithOptions(request *DescribeVerifyTokenRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifyTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackSeed)) {
		query["CallbackSeed"] = request.CallbackSeed
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FaceRetainedImageUrl)) {
		query["FaceRetainedImageUrl"] = request.FaceRetainedImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FailedRedirectUrl)) {
		query["FailedRedirectUrl"] = request.FailedRedirectUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardBackImageUrl)) {
		query["IdCardBackImageUrl"] = request.IdCardBackImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardFrontImageUrl)) {
		query["IdCardFrontImageUrl"] = request.IdCardFrontImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardNumber)) {
		query["IdCardNumber"] = request.IdCardNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PassedRedirectUrl)) {
		query["PassedRedirectUrl"] = request.PassedRedirectUrl
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIp)) {
		query["UserIp"] = request.UserIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserPhoneNumber)) {
		query["UserPhoneNumber"] = request.UserPhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.UserRegistTime)) {
		query["UserRegistTime"] = request.UserRegistTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVerifyToken"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVerifyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeVerifyTokenRequest
//
// @return DescribeVerifyTokenResponse
func (client *Client) DescribeVerifyToken(request *DescribeVerifyTokenRequest) (_result *DescribeVerifyTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifyTokenResponse{}
	_body, _err := client.DescribeVerifyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DetectFaceAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectFaceAttributesResponse
func (client *Client) DetectFaceAttributesWithOptions(request *DetectFaceAttributesRequest, runtime *util.RuntimeOptions) (_result *DetectFaceAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialValue)) {
		body["MaterialValue"] = request.MaterialValue
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectFaceAttributes"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectFaceAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DetectFaceAttributesRequest
//
// @return DetectFaceAttributesResponse
func (client *Client) DetectFaceAttributes(request *DetectFaceAttributesRequest) (_result *DetectFaceAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectFaceAttributesResponse{}
	_body, _err := client.DetectFaceAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 身份二要素接口
//
// @param request - Id2MetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaVerifyResponse
func (client *Client) Id2MetaVerifyWithOptions(request *Id2MetaVerifyRequest, runtime *util.RuntimeOptions) (_result *Id2MetaVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentifyNum)) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		body["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Id2MetaVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Id2MetaVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 身份二要素接口
//
// @param request - Id2MetaVerifyRequest
//
// @return Id2MetaVerifyResponse
func (client *Client) Id2MetaVerify(request *Id2MetaVerifyRequest) (_result *Id2MetaVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Id2MetaVerifyResponse{}
	_body, _err := client.Id2MetaVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - InitFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitFaceVerifyResponse
func (client *Client) InitFaceVerifyWithOptions(request *InitFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *InitFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppQualityCheck)) {
		query["AppQualityCheck"] = request.AppQualityCheck
	}

	if !tea.BoolValue(util.IsUnset(request.Birthday)) {
		query["Birthday"] = request.Birthday
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackToken)) {
		query["CallbackToken"] = request.CallbackToken
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.CertNo)) {
		query["CertNo"] = request.CertNo
	}

	if !tea.BoolValue(util.IsUnset(request.CertType)) {
		query["CertType"] = request.CertType
	}

	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		query["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.CertifyUrlStyle)) {
		query["CertifyUrlStyle"] = request.CertifyUrlStyle
	}

	if !tea.BoolValue(util.IsUnset(request.CertifyUrlType)) {
		query["CertifyUrlType"] = request.CertifyUrlType
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptType)) {
		query["EncryptType"] = request.EncryptType
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastPictureUrl)) {
		query["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FaceGuardOutput)) {
		query["FaceGuardOutput"] = request.FaceGuardOutput
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.MetaInfo)) {
		query["MetaInfo"] = request.MetaInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssObjectName)) {
		query["OssObjectName"] = request.OssObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.OuterOrderNo)) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.ProcedurePriority)) {
		query["ProcedurePriority"] = request.ProcedurePriority
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.RarelyCharacters)) {
		query["RarelyCharacters"] = request.RarelyCharacters
	}

	if !tea.BoolValue(util.IsUnset(request.ReadImg)) {
		query["ReadImg"] = request.ReadImg
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnUrl)) {
		query["ReturnUrl"] = request.ReturnUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.SuitableType)) {
		query["SuitableType"] = request.SuitableType
	}

	if !tea.BoolValue(util.IsUnset(request.UiCustomUrl)) {
		query["UiCustomUrl"] = request.UiCustomUrl
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.ValidityDate)) {
		query["ValidityDate"] = request.ValidityDate
	}

	if !tea.BoolValue(util.IsUnset(request.VideoEvidence)) {
		query["VideoEvidence"] = request.VideoEvidence
	}

	if !tea.BoolValue(util.IsUnset(request.VoluntaryCustomizedContent)) {
		query["VoluntaryCustomizedContent"] = request.VoluntaryCustomizedContent
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthId)) {
		body["AuthId"] = request.AuthId
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastPicture)) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["Model"] = request.Model
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InitFaceVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - InitFaceVerifyRequest
//
// @return InitFaceVerifyResponse
func (client *Client) InitFaceVerify(request *InitFaceVerifyRequest) (_result *InitFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.InitFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增实人白名单
//
// @param request - InsertWhiteListSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertWhiteListSettingResponse
func (client *Client) InsertWhiteListSettingWithOptions(request *InsertWhiteListSettingRequest, runtime *util.RuntimeOptions) (_result *InsertWhiteListSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertNo)) {
		query["CertNo"] = request.CertNo
	}

	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		query["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.ValidDay)) {
		query["ValidDay"] = request.ValidDay
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertWhiteListSetting"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertWhiteListSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增实人白名单
//
// @param request - InsertWhiteListSettingRequest
//
// @return InsertWhiteListSettingResponse
func (client *Client) InsertWhiteListSetting(request *InsertWhiteListSettingRequest) (_result *InsertWhiteListSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertWhiteListSettingResponse{}
	_body, _err := client.InsertWhiteListSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - LivenessFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LivenessFaceVerifyResponse
func (client *Client) LivenessFaceVerifyWithOptions(request *LivenessFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *LivenessFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Model)) {
		query["Model"] = request.Model
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		body["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceToken)) {
		body["DeviceToken"] = request.DeviceToken
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastPicture)) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastPictureUrl)) {
		body["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		body["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		body["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssObjectName)) {
		body["OssObjectName"] = request.OssObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.OuterOrderNo)) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LivenessFaceVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LivenessFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - LivenessFaceVerifyRequest
//
// @return LivenessFaceVerifyResponse
func (client *Client) LivenessFaceVerify(request *LivenessFaceVerifyRequest) (_result *LivenessFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LivenessFaceVerifyResponse{}
	_body, _err := client.LivenessFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 手机三要素详版接口
//
// @param request - Mobile3MetaDetailVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile3MetaDetailVerifyResponse
func (client *Client) Mobile3MetaDetailVerifyWithOptions(request *Mobile3MetaDetailVerifyRequest, runtime *util.RuntimeOptions) (_result *Mobile3MetaDetailVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentifyNum)) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		body["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Mobile3MetaDetailVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Mobile3MetaDetailVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 手机三要素详版接口
//
// @param request - Mobile3MetaDetailVerifyRequest
//
// @return Mobile3MetaDetailVerifyResponse
func (client *Client) Mobile3MetaDetailVerify(request *Mobile3MetaDetailVerifyRequest) (_result *Mobile3MetaDetailVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Mobile3MetaDetailVerifyResponse{}
	_body, _err := client.Mobile3MetaDetailVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 手机号三要素简版接口
//
// @param request - Mobile3MetaSimpleVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile3MetaSimpleVerifyResponse
func (client *Client) Mobile3MetaSimpleVerifyWithOptions(request *Mobile3MetaSimpleVerifyRequest, runtime *util.RuntimeOptions) (_result *Mobile3MetaSimpleVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentifyNum)) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		body["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Mobile3MetaSimpleVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Mobile3MetaSimpleVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 手机号三要素简版接口
//
// @param request - Mobile3MetaSimpleVerifyRequest
//
// @return Mobile3MetaSimpleVerifyResponse
func (client *Client) Mobile3MetaSimpleVerify(request *Mobile3MetaSimpleVerifyRequest) (_result *Mobile3MetaSimpleVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Mobile3MetaSimpleVerifyResponse{}
	_body, _err := client.Mobile3MetaSimpleVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 号码检测
//
// @param request - MobileDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MobileDetectResponse
func (client *Client) MobileDetectWithOptions(request *MobileDetectRequest, runtime *util.RuntimeOptions) (_result *MobileDetectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mobiles)) {
		body["Mobiles"] = request.Mobiles
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		body["ParamType"] = request.ParamType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MobileDetect"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MobileDetectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 号码检测
//
// @param request - MobileDetectRequest
//
// @return MobileDetectResponse
func (client *Client) MobileDetect(request *MobileDetectRequest) (_result *MobileDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MobileDetectResponse{}
	_body, _err := client.MobileDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询手机号在网状态
//
// @param request - MobileOnlineStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MobileOnlineStatusResponse
func (client *Client) MobileOnlineStatusWithOptions(request *MobileOnlineStatusRequest, runtime *util.RuntimeOptions) (_result *MobileOnlineStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		body["ParamType"] = request.ParamType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MobileOnlineStatus"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MobileOnlineStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询手机号在网状态
//
// @param request - MobileOnlineStatusRequest
//
// @return MobileOnlineStatusResponse
func (client *Client) MobileOnlineStatus(request *MobileOnlineStatusRequest) (_result *MobileOnlineStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MobileOnlineStatusResponse{}
	_body, _err := client.MobileOnlineStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询手机号在网时长
//
// @param request - MobileOnlineTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MobileOnlineTimeResponse
func (client *Client) MobileOnlineTimeWithOptions(request *MobileOnlineTimeRequest, runtime *util.RuntimeOptions) (_result *MobileOnlineTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		body["ParamType"] = request.ParamType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MobileOnlineTime"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MobileOnlineTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询手机号在网时长
//
// @param request - MobileOnlineTimeRequest
//
// @return MobileOnlineTimeResponse
func (client *Client) MobileOnlineTime(request *MobileOnlineTimeRequest) (_result *MobileOnlineTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MobileOnlineTimeResponse{}
	_body, _err := client.MobileOnlineTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeviceInfoResponse
func (client *Client) ModifyDeviceInfoWithOptions(request *ModifyDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *ModifyDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredDay)) {
		query["ExpiredDay"] = request.ExpiredDay
	}

	if !tea.BoolValue(util.IsUnset(request.UserDeviceId)) {
		query["UserDeviceId"] = request.UserDeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDeviceInfo"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyDeviceInfoRequest
//
// @return ModifyDeviceInfoResponse
func (client *Client) ModifyDeviceInfo(request *ModifyDeviceInfoRequest) (_result *ModifyDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDeviceInfoResponse{}
	_body, _err := client.ModifyDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询实人白名单配置
//
// @param request - PageQueryWhiteListSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageQueryWhiteListSettingResponse
func (client *Client) PageQueryWhiteListSettingWithOptions(request *PageQueryWhiteListSettingRequest, runtime *util.RuntimeOptions) (_result *PageQueryWhiteListSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertNo)) {
		query["CertNo"] = request.CertNo
	}

	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		query["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ValidEndDate)) {
		query["ValidEndDate"] = request.ValidEndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ValidStartDate)) {
		query["ValidStartDate"] = request.ValidStartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PageQueryWhiteListSetting"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageQueryWhiteListSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询实人白名单配置
//
// @param request - PageQueryWhiteListSettingRequest
//
// @return PageQueryWhiteListSettingResponse
func (client *Client) PageQueryWhiteListSetting(request *PageQueryWhiteListSettingRequest) (_result *PageQueryWhiteListSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageQueryWhiteListSettingResponse{}
	_body, _err := client.PageQueryWhiteListSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除实人白名单
//
// @param tmpReq - RemoveWhiteListSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveWhiteListSettingResponse
func (client *Client) RemoveWhiteListSettingWithOptions(tmpReq *RemoveWhiteListSettingRequest, runtime *util.RuntimeOptions) (_result *RemoveWhiteListSettingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RemoveWhiteListSettingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Ids)) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, tea.String("Ids"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdsShrink)) {
		query["Ids"] = request.IdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		query["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveWhiteListSetting"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveWhiteListSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实人白名单
//
// @param request - RemoveWhiteListSettingRequest
//
// @return RemoveWhiteListSettingResponse
func (client *Client) RemoveWhiteListSetting(request *RemoveWhiteListSettingRequest) (_result *RemoveWhiteListSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveWhiteListSettingResponse{}
	_body, _err := client.RemoveWhiteListSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 车五项信息识别
//
// @param request - Vehicle5ItemQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Vehicle5ItemQueryResponse
func (client *Client) Vehicle5ItemQueryWithOptions(request *Vehicle5ItemQueryRequest, runtime *util.RuntimeOptions) (_result *Vehicle5ItemQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		query["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.VehicleNum)) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !tea.BoolValue(util.IsUnset(request.VehicleType)) {
		query["VehicleType"] = request.VehicleType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Vehicle5ItemQuery"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Vehicle5ItemQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 车五项信息识别
//
// @param request - Vehicle5ItemQueryRequest
//
// @return Vehicle5ItemQueryResponse
func (client *Client) Vehicle5ItemQuery(request *Vehicle5ItemQueryRequest) (_result *Vehicle5ItemQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Vehicle5ItemQueryResponse{}
	_body, _err := client.Vehicle5ItemQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 车辆投保日期查询
//
// @param request - VehicleInsureQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VehicleInsureQueryResponse
func (client *Client) VehicleInsureQueryWithOptions(request *VehicleInsureQueryRequest, runtime *util.RuntimeOptions) (_result *VehicleInsureQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		query["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.VehicleNum)) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !tea.BoolValue(util.IsUnset(request.VehicleType)) {
		query["VehicleType"] = request.VehicleType
	}

	if !tea.BoolValue(util.IsUnset(request.Vin)) {
		query["Vin"] = request.Vin
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VehicleInsureQuery"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VehicleInsureQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 车辆投保日期查询
//
// @param request - VehicleInsureQueryRequest
//
// @return VehicleInsureQueryResponse
func (client *Client) VehicleInsureQuery(request *VehicleInsureQueryRequest) (_result *VehicleInsureQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VehicleInsureQueryResponse{}
	_body, _err := client.VehicleInsureQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 车辆要素核验
//
// @param request - VehicleMetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VehicleMetaVerifyResponse
func (client *Client) VehicleMetaVerifyWithOptions(request *VehicleMetaVerifyRequest, runtime *util.RuntimeOptions) (_result *VehicleMetaVerifyResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.VehicleNum)) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !tea.BoolValue(util.IsUnset(request.VehicleType)) {
		query["VehicleType"] = request.VehicleType
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyMetaType)) {
		query["VerifyMetaType"] = request.VerifyMetaType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VehicleMetaVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VehicleMetaVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 车辆要素核验
//
// @param request - VehicleMetaVerifyRequest
//
// @return VehicleMetaVerifyResponse
func (client *Client) VehicleMetaVerify(request *VehicleMetaVerifyRequest) (_result *VehicleMetaVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VehicleMetaVerifyResponse{}
	_body, _err := client.VehicleMetaVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 车辆要素核验增强版
//
// @param request - VehicleMetaVerifyV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VehicleMetaVerifyV2Response
func (client *Client) VehicleMetaVerifyV2WithOptions(request *VehicleMetaVerifyV2Request, runtime *util.RuntimeOptions) (_result *VehicleMetaVerifyV2Response, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.VehicleNum)) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !tea.BoolValue(util.IsUnset(request.VehicleType)) {
		query["VehicleType"] = request.VehicleType
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyMetaType)) {
		query["VerifyMetaType"] = request.VerifyMetaType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VehicleMetaVerifyV2"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VehicleMetaVerifyV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 车辆要素核验增强版
//
// @param request - VehicleMetaVerifyV2Request
//
// @return VehicleMetaVerifyV2Response
func (client *Client) VehicleMetaVerifyV2(request *VehicleMetaVerifyV2Request) (_result *VehicleMetaVerifyV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VehicleMetaVerifyV2Response{}
	_body, _err := client.VehicleMetaVerifyV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 车辆信息识别
//
// @param request - VehicleQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VehicleQueryResponse
func (client *Client) VehicleQueryWithOptions(request *VehicleQueryRequest, runtime *util.RuntimeOptions) (_result *VehicleQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		query["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.VehicleNum)) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !tea.BoolValue(util.IsUnset(request.VehicleType)) {
		query["VehicleType"] = request.VehicleType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VehicleQuery"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VehicleQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 车辆信息识别
//
// @param request - VehicleQueryRequest
//
// @return VehicleQueryResponse
func (client *Client) VehicleQuery(request *VehicleQueryRequest) (_result *VehicleQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VehicleQueryResponse{}
	_body, _err := client.VehicleQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - VerifyMaterialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyMaterialResponse
func (client *Client) VerifyMaterialWithOptions(request *VerifyMaterialRequest, runtime *util.RuntimeOptions) (_result *VerifyMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.FaceImageUrl)) {
		query["FaceImageUrl"] = request.FaceImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardBackImageUrl)) {
		query["IdCardBackImageUrl"] = request.IdCardBackImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardFrontImageUrl)) {
		query["IdCardFrontImageUrl"] = request.IdCardFrontImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardNumber)) {
		query["IdCardNumber"] = request.IdCardNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyMaterial"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyMaterialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - VerifyMaterialRequest
//
// @return VerifyMaterialResponse
func (client *Client) VerifyMaterial(request *VerifyMaterialRequest) (_result *VerifyMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyMaterialResponse{}
	_body, _err := client.VerifyMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
