// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
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

type CompareFacesRequest struct {
	BizId             *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType           *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	SourceImageBase64 *string `json:"SourceImageBase64,omitempty" xml:"SourceImageBase64,omitempty"`
	SourceImageUrl    *string `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty"`
	TargetImageBase64 *string `json:"TargetImageBase64,omitempty" xml:"TargetImageBase64,omitempty"`
	TargetImageUrl    *string `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
}

func (s CompareFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesRequest) GoString() string {
	return s.String()
}

func (s *CompareFacesRequest) SetBizId(v string) *CompareFacesRequest {
	s.BizId = &v
	return s
}

func (s *CompareFacesRequest) SetBizType(v string) *CompareFacesRequest {
	s.BizType = &v
	return s
}

func (s *CompareFacesRequest) SetSourceImageBase64(v string) *CompareFacesRequest {
	s.SourceImageBase64 = &v
	return s
}

func (s *CompareFacesRequest) SetSourceImageUrl(v string) *CompareFacesRequest {
	s.SourceImageUrl = &v
	return s
}

func (s *CompareFacesRequest) SetTargetImageBase64(v string) *CompareFacesRequest {
	s.TargetImageBase64 = &v
	return s
}

func (s *CompareFacesRequest) SetTargetImageUrl(v string) *CompareFacesRequest {
	s.TargetImageUrl = &v
	return s
}

type CompareFacesResponseBody struct {
	Code         *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *CompareFacesResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *CompareFacesResponseBody) SetMessage(v string) *CompareFacesResponseBody {
	s.Message = &v
	return s
}

func (s *CompareFacesResponseBody) SetRequestId(v string) *CompareFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareFacesResponseBody) SetResultObject(v *CompareFacesResponseBodyResultObject) *CompareFacesResponseBody {
	s.ResultObject = v
	return s
}

func (s *CompareFacesResponseBody) SetSuccess(v bool) *CompareFacesResponseBody {
	s.Success = &v
	return s
}

type CompareFacesResponseBodyResultObject struct {
	ConfidenceThresholds *string  `json:"ConfidenceThresholds,omitempty" xml:"ConfidenceThresholds,omitempty"`
	SimilarityScore      *float32 `json:"SimilarityScore,omitempty" xml:"SimilarityScore,omitempty"`
}

func (s CompareFacesResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CompareFacesResponseBodyResultObject) SetConfidenceThresholds(v string) *CompareFacesResponseBodyResultObject {
	s.ConfidenceThresholds = &v
	return s
}

func (s *CompareFacesResponseBodyResultObject) SetSimilarityScore(v float32) *CompareFacesResponseBodyResultObject {
	s.SimilarityScore = &v
	return s
}

type CompareFacesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CompareFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeVerifyResultRequest struct {
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
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
	Code         *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DescribeVerifyResultResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeVerifyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseBody) SetCode(v string) *DescribeVerifyResultResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetMessage(v string) *DescribeVerifyResultResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetRequestId(v string) *DescribeVerifyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetResultObject(v *DescribeVerifyResultResponseBodyResultObject) *DescribeVerifyResultResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetSuccess(v bool) *DescribeVerifyResultResponseBody {
	s.Success = &v
	return s
}

type DescribeVerifyResultResponseBodyResultObject struct {
	AuthorityComparisionScore *float32                                              `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty"`
	FaceComparisonScore       *float32                                              `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	IdCardFaceComparisonScore *float32                                              `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	Material                  *DescribeVerifyResultResponseBodyResultObjectMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	VerifyStatus              *int32                                                `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
}

func (s DescribeVerifyResultResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseBodyResultObject) SetAuthorityComparisionScore(v float32) *DescribeVerifyResultResponseBodyResultObject {
	s.AuthorityComparisionScore = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObject) SetFaceComparisonScore(v float32) *DescribeVerifyResultResponseBodyResultObject {
	s.FaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObject) SetIdCardFaceComparisonScore(v float32) *DescribeVerifyResultResponseBodyResultObject {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObject) SetMaterial(v *DescribeVerifyResultResponseBodyResultObjectMaterial) *DescribeVerifyResultResponseBodyResultObject {
	s.Material = v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObject) SetVerifyStatus(v int32) *DescribeVerifyResultResponseBodyResultObject {
	s.VerifyStatus = &v
	return s
}

type DescribeVerifyResultResponseBodyResultObjectMaterial struct {
	FaceGlobalUrl *string                                                         `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty"`
	FaceImageUrl  *string                                                         `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	FaceMask      *bool                                                           `json:"FaceMask,omitempty" xml:"FaceMask,omitempty"`
	FaceQuality   *string                                                         `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	IdCardInfo    *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	IdCardName    *string                                                         `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	IdCardNumber  *string                                                         `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	VideoUrls     []*string                                                       `json:"VideoUrls,omitempty" xml:"VideoUrls,omitempty" type:"Repeated"`
}

func (s DescribeVerifyResultResponseBodyResultObjectMaterial) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseBodyResultObjectMaterial) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterial) SetFaceGlobalUrl(v string) *DescribeVerifyResultResponseBodyResultObjectMaterial {
	s.FaceGlobalUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterial) SetFaceImageUrl(v string) *DescribeVerifyResultResponseBodyResultObjectMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterial) SetFaceMask(v bool) *DescribeVerifyResultResponseBodyResultObjectMaterial {
	s.FaceMask = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterial) SetFaceQuality(v string) *DescribeVerifyResultResponseBodyResultObjectMaterial {
	s.FaceQuality = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterial) SetIdCardInfo(v *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo) *DescribeVerifyResultResponseBodyResultObjectMaterial {
	s.IdCardInfo = v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterial) SetIdCardName(v string) *DescribeVerifyResultResponseBodyResultObjectMaterial {
	s.IdCardName = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterial) SetIdCardNumber(v string) *DescribeVerifyResultResponseBodyResultObjectMaterial {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterial) SetVideoUrls(v []*string) *DescribeVerifyResultResponseBodyResultObjectMaterial {
	s.VideoUrls = v
	return s
}

type DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo struct {
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Authority     *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	BackImageUrl  *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty"`
	Birth         *string `json:"Birth,omitempty" xml:"Birth,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	Number        *string `json:"Number,omitempty" xml:"Number,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo) SetAddress(v string) *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo) SetAuthority(v string) *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo) SetBackImageUrl(v string) *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo) SetBirth(v string) *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo) SetEndDate(v string) *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo) SetFrontImageUrl(v string) *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo) SetName(v string) *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo) SetNationality(v string) *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo) SetNumber(v string) *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo) SetStartDate(v string) *DescribeVerifyResultResponseBodyResultObjectMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

type DescribeVerifyResultResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeVerifyTokenRequest struct {
	BizId                *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	FaceRetainedImageUrl *string `json:"FaceRetainedImageUrl,omitempty" xml:"FaceRetainedImageUrl,omitempty"`
	IdCardBackImageUrl   *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	IdCardFrontImageUrl  *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	IdCardNumber         *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserIp               *string `json:"UserIp,omitempty" xml:"UserIp,omitempty"`
	UserPhoneNumber      *string `json:"UserPhoneNumber,omitempty" xml:"UserPhoneNumber,omitempty"`
	UserRegistTime       *int64  `json:"UserRegistTime,omitempty" xml:"UserRegistTime,omitempty"`
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

func (s *DescribeVerifyTokenRequest) SetFaceRetainedImageUrl(v string) *DescribeVerifyTokenRequest {
	s.FaceRetainedImageUrl = &v
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
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DescribeVerifyTokenResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeVerifyTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponseBody) SetCode(v string) *DescribeVerifyTokenResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeVerifyTokenResponseBody) SetMessage(v string) *DescribeVerifyTokenResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVerifyTokenResponseBody) SetRequestId(v string) *DescribeVerifyTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyTokenResponseBody) SetResultObject(v *DescribeVerifyTokenResponseBodyResultObject) *DescribeVerifyTokenResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVerifyTokenResponseBody) SetSuccess(v bool) *DescribeVerifyTokenResponseBody {
	s.Success = &v
	return s
}

type DescribeVerifyTokenResponseBodyResultObject struct {
	VerifyPageUrl *string `json:"VerifyPageUrl,omitempty" xml:"VerifyPageUrl,omitempty"`
	VerifyToken   *string `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
}

func (s DescribeVerifyTokenResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponseBodyResultObject) SetVerifyPageUrl(v string) *DescribeVerifyTokenResponseBodyResultObject {
	s.VerifyPageUrl = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyResultObject) SetVerifyToken(v string) *DescribeVerifyTokenResponseBodyResultObject {
	s.VerifyToken = &v
	return s
}

type DescribeVerifyTokenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVerifyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ImageFile *string `json:"ImageFile,omitempty" xml:"ImageFile,omitempty"`
	ImageUrl  *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s DetectFaceAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesRequest) SetBizId(v string) *DetectFaceAttributesRequest {
	s.BizId = &v
	return s
}

func (s *DetectFaceAttributesRequest) SetBizType(v string) *DetectFaceAttributesRequest {
	s.BizType = &v
	return s
}

func (s *DetectFaceAttributesRequest) SetImageFile(v string) *DetectFaceAttributesRequest {
	s.ImageFile = &v
	return s
}

func (s *DetectFaceAttributesRequest) SetImageUrl(v string) *DetectFaceAttributesRequest {
	s.ImageUrl = &v
	return s
}

type DetectFaceAttributesAdvanceRequest struct {
	BizId           *string   `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType         *string   `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ImageFileObject io.Reader `json:"ImageFile,omitempty" xml:"ImageFile,omitempty"`
	ImageUrl        *string   `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s DetectFaceAttributesAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesAdvanceRequest) SetBizId(v string) *DetectFaceAttributesAdvanceRequest {
	s.BizId = &v
	return s
}

func (s *DetectFaceAttributesAdvanceRequest) SetBizType(v string) *DetectFaceAttributesAdvanceRequest {
	s.BizType = &v
	return s
}

func (s *DetectFaceAttributesAdvanceRequest) SetImageFileObject(v io.Reader) *DetectFaceAttributesAdvanceRequest {
	s.ImageFileObject = v
	return s
}

func (s *DetectFaceAttributesAdvanceRequest) SetImageUrl(v string) *DetectFaceAttributesAdvanceRequest {
	s.ImageUrl = &v
	return s
}

type DetectFaceAttributesResponseBody struct {
	Code         *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DetectFaceAttributesResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *DetectFaceAttributesResponseBody) SetMessage(v string) *DetectFaceAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetRequestId(v string) *DetectFaceAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetResultObject(v *DetectFaceAttributesResponseBodyResultObject) *DetectFaceAttributesResponseBody {
	s.ResultObject = v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetSuccess(v bool) *DetectFaceAttributesResponseBody {
	s.Success = &v
	return s
}

type DetectFaceAttributesResponseBodyResultObject struct {
	FaceInfos *DetectFaceAttributesResponseBodyResultObjectFaceInfos `json:"FaceInfos,omitempty" xml:"FaceInfos,omitempty" type:"Struct"`
	ImgHeight *int32                                                 `json:"ImgHeight,omitempty" xml:"ImgHeight,omitempty"`
	ImgWidth  *int32                                                 `json:"ImgWidth,omitempty" xml:"ImgWidth,omitempty"`
}

func (s DetectFaceAttributesResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyResultObject) SetFaceInfos(v *DetectFaceAttributesResponseBodyResultObjectFaceInfos) *DetectFaceAttributesResponseBodyResultObject {
	s.FaceInfos = v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObject) SetImgHeight(v int32) *DetectFaceAttributesResponseBodyResultObject {
	s.ImgHeight = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObject) SetImgWidth(v int32) *DetectFaceAttributesResponseBodyResultObject {
	s.ImgWidth = &v
	return s
}

type DetectFaceAttributesResponseBodyResultObjectFaceInfos struct {
	FaceAttributesDetectInfo []*DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfo `json:"FaceAttributesDetectInfo,omitempty" xml:"FaceAttributesDetectInfo,omitempty" type:"Repeated"`
}

func (s DetectFaceAttributesResponseBodyResultObjectFaceInfos) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyResultObjectFaceInfos) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfos) SetFaceAttributesDetectInfo(v []*DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfo) *DetectFaceAttributesResponseBodyResultObjectFaceInfos {
	s.FaceAttributesDetectInfo = v
	return s
}

type DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfo struct {
	FaceAttributes *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	FaceRect       *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceRect       `json:"FaceRect,omitempty" xml:"FaceRect,omitempty" type:"Struct"`
}

func (s DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfo) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfo) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfo) SetFaceAttributes(v *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfo {
	s.FaceAttributes = v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfo) SetFaceRect(v *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceRect) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfo {
	s.FaceRect = v
	return s
}

type DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes struct {
	Blur       *float32                                                                                             `json:"Blur,omitempty" xml:"Blur,omitempty"`
	Facequal   *float32                                                                                             `json:"Facequal,omitempty" xml:"Facequal,omitempty"`
	Facetype   *string                                                                                              `json:"Facetype,omitempty" xml:"Facetype,omitempty"`
	Glasses    *string                                                                                              `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	Headpose   *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose `json:"Headpose,omitempty" xml:"Headpose,omitempty" type:"Struct"`
	Integrity  *int32                                                                                               `json:"Integrity,omitempty" xml:"Integrity,omitempty"`
	Respirator *string                                                                                              `json:"Respirator,omitempty" xml:"Respirator,omitempty"`
	Smiling    *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling  `json:"Smiling,omitempty" xml:"Smiling,omitempty" type:"Struct"`
}

func (s DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes) SetBlur(v float32) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Blur = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes) SetFacequal(v float32) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Facequal = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes) SetFacetype(v string) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Facetype = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes) SetGlasses(v string) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes) SetHeadpose(v *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Headpose = v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes) SetIntegrity(v int32) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Integrity = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes) SetRespirator(v string) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Respirator = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes) SetSmiling(v *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Smiling = v
	return s
}

type DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose struct {
	PitchAngle *float32 `json:"PitchAngle,omitempty" xml:"PitchAngle,omitempty"`
	RollAngle  *float32 `json:"RollAngle,omitempty" xml:"RollAngle,omitempty"`
	YawAngle   *float32 `json:"YawAngle,omitempty" xml:"YawAngle,omitempty"`
}

func (s DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetPitchAngle(v float32) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.PitchAngle = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetRollAngle(v float32) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.RollAngle = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetYawAngle(v float32) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.YawAngle = &v
	return s
}

type DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling struct {
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Value     *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) SetThreshold(v float32) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling {
	s.Threshold = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) SetValue(v float32) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling {
	s.Value = &v
	return s
}

type DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceRect struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceRect) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceRect) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceRect) SetHeight(v int32) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Height = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceRect) SetLeft(v int32) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Left = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceRect) SetTop(v int32) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Top = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceRect) SetWidth(v int32) *DetectFaceAttributesResponseBodyResultObjectFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Width = &v
	return s
}

type DetectFaceAttributesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectFaceAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type LivenessDetectRequest struct {
	BizId         *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	MediaCategory *string `json:"MediaCategory,omitempty" xml:"MediaCategory,omitempty"`
	MediaFile     *string `json:"MediaFile,omitempty" xml:"MediaFile,omitempty"`
	MediaUrl      *string `json:"MediaUrl,omitempty" xml:"MediaUrl,omitempty"`
}

func (s LivenessDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s LivenessDetectRequest) GoString() string {
	return s.String()
}

func (s *LivenessDetectRequest) SetBizId(v string) *LivenessDetectRequest {
	s.BizId = &v
	return s
}

func (s *LivenessDetectRequest) SetBizType(v string) *LivenessDetectRequest {
	s.BizType = &v
	return s
}

func (s *LivenessDetectRequest) SetMediaCategory(v string) *LivenessDetectRequest {
	s.MediaCategory = &v
	return s
}

func (s *LivenessDetectRequest) SetMediaFile(v string) *LivenessDetectRequest {
	s.MediaFile = &v
	return s
}

func (s *LivenessDetectRequest) SetMediaUrl(v string) *LivenessDetectRequest {
	s.MediaUrl = &v
	return s
}

type LivenessDetectAdvanceRequest struct {
	BizId           *string   `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType         *string   `json:"BizType,omitempty" xml:"BizType,omitempty"`
	MediaCategory   *string   `json:"MediaCategory,omitempty" xml:"MediaCategory,omitempty"`
	MediaFileObject io.Reader `json:"MediaFile,omitempty" xml:"MediaFile,omitempty"`
	MediaUrl        *string   `json:"MediaUrl,omitempty" xml:"MediaUrl,omitempty"`
}

func (s LivenessDetectAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s LivenessDetectAdvanceRequest) GoString() string {
	return s.String()
}

func (s *LivenessDetectAdvanceRequest) SetBizId(v string) *LivenessDetectAdvanceRequest {
	s.BizId = &v
	return s
}

func (s *LivenessDetectAdvanceRequest) SetBizType(v string) *LivenessDetectAdvanceRequest {
	s.BizType = &v
	return s
}

func (s *LivenessDetectAdvanceRequest) SetMediaCategory(v string) *LivenessDetectAdvanceRequest {
	s.MediaCategory = &v
	return s
}

func (s *LivenessDetectAdvanceRequest) SetMediaFileObject(v io.Reader) *LivenessDetectAdvanceRequest {
	s.MediaFileObject = v
	return s
}

func (s *LivenessDetectAdvanceRequest) SetMediaUrl(v string) *LivenessDetectAdvanceRequest {
	s.MediaUrl = &v
	return s
}

type LivenessDetectResponseBody struct {
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *LivenessDetectResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LivenessDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LivenessDetectResponseBody) GoString() string {
	return s.String()
}

func (s *LivenessDetectResponseBody) SetCode(v string) *LivenessDetectResponseBody {
	s.Code = &v
	return s
}

func (s *LivenessDetectResponseBody) SetMessage(v string) *LivenessDetectResponseBody {
	s.Message = &v
	return s
}

func (s *LivenessDetectResponseBody) SetRequestId(v string) *LivenessDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *LivenessDetectResponseBody) SetResultObject(v *LivenessDetectResponseBodyResultObject) *LivenessDetectResponseBody {
	s.ResultObject = v
	return s
}

func (s *LivenessDetectResponseBody) SetSuccess(v bool) *LivenessDetectResponseBody {
	s.Success = &v
	return s
}

type LivenessDetectResponseBodyResultObject struct {
	FrameUrl *string  `json:"FrameUrl,omitempty" xml:"FrameUrl,omitempty"`
	Passed   *string  `json:"Passed,omitempty" xml:"Passed,omitempty"`
	Score    *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s LivenessDetectResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s LivenessDetectResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *LivenessDetectResponseBodyResultObject) SetFrameUrl(v string) *LivenessDetectResponseBodyResultObject {
	s.FrameUrl = &v
	return s
}

func (s *LivenessDetectResponseBodyResultObject) SetPassed(v string) *LivenessDetectResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *LivenessDetectResponseBodyResultObject) SetScore(v float32) *LivenessDetectResponseBodyResultObject {
	s.Score = &v
	return s
}

type LivenessDetectResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LivenessDetectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LivenessDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s LivenessDetectResponse) GoString() string {
	return s.String()
}

func (s *LivenessDetectResponse) SetHeaders(v map[string]*string) *LivenessDetectResponse {
	s.Headers = v
	return s
}

func (s *LivenessDetectResponse) SetStatusCode(v int32) *LivenessDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *LivenessDetectResponse) SetBody(v *LivenessDetectResponseBody) *LivenessDetectResponse {
	s.Body = v
	return s
}

type VerifyMaterialRequest struct {
	BizId               *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType             *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	FaceImageUrl        *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	IdCardBackImageUrl  *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	IdCardFrontImageUrl *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	IdCardNumber        *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *VerifyMaterialResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseBody) SetCode(v string) *VerifyMaterialResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetMessage(v string) *VerifyMaterialResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetRequestId(v string) *VerifyMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetResultObject(v *VerifyMaterialResponseBodyResultObject) *VerifyMaterialResponseBody {
	s.ResultObject = v
	return s
}

func (s *VerifyMaterialResponseBody) SetSuccess(v bool) *VerifyMaterialResponseBody {
	s.Success = &v
	return s
}

type VerifyMaterialResponseBodyResultObject struct {
	AuthorityComparisionScore *float32                                        `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty"`
	IdCardFaceComparisonScore *float32                                        `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	Material                  *VerifyMaterialResponseBodyResultObjectMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	VerifyStatus              *int32                                          `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
	VerifyToken               *string                                         `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
}

func (s VerifyMaterialResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseBodyResultObject) SetAuthorityComparisionScore(v float32) *VerifyMaterialResponseBodyResultObject {
	s.AuthorityComparisionScore = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObject) SetIdCardFaceComparisonScore(v float32) *VerifyMaterialResponseBodyResultObject {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObject) SetMaterial(v *VerifyMaterialResponseBodyResultObjectMaterial) *VerifyMaterialResponseBodyResultObject {
	s.Material = v
	return s
}

func (s *VerifyMaterialResponseBodyResultObject) SetVerifyStatus(v int32) *VerifyMaterialResponseBodyResultObject {
	s.VerifyStatus = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObject) SetVerifyToken(v string) *VerifyMaterialResponseBodyResultObject {
	s.VerifyToken = &v
	return s
}

type VerifyMaterialResponseBodyResultObjectMaterial struct {
	FaceGlobalUrl *string                                                   `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty"`
	FaceImageUrl  *string                                                   `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	FaceMask      *string                                                   `json:"FaceMask,omitempty" xml:"FaceMask,omitempty"`
	FaceQuality   *string                                                   `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	IdCardInfo    *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	IdCardName    *string                                                   `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	IdCardNumber  *string                                                   `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
}

func (s VerifyMaterialResponseBodyResultObjectMaterial) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseBodyResultObjectMaterial) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseBodyResultObjectMaterial) SetFaceGlobalUrl(v string) *VerifyMaterialResponseBodyResultObjectMaterial {
	s.FaceGlobalUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterial) SetFaceImageUrl(v string) *VerifyMaterialResponseBodyResultObjectMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterial) SetFaceMask(v string) *VerifyMaterialResponseBodyResultObjectMaterial {
	s.FaceMask = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterial) SetFaceQuality(v string) *VerifyMaterialResponseBodyResultObjectMaterial {
	s.FaceQuality = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterial) SetIdCardInfo(v *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo) *VerifyMaterialResponseBodyResultObjectMaterial {
	s.IdCardInfo = v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterial) SetIdCardName(v string) *VerifyMaterialResponseBodyResultObjectMaterial {
	s.IdCardName = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterial) SetIdCardNumber(v string) *VerifyMaterialResponseBodyResultObjectMaterial {
	s.IdCardNumber = &v
	return s
}

type VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo struct {
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Authority     *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	BackImageUrl  *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty"`
	Birth         *string `json:"Birth,omitempty" xml:"Birth,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	Number        *string `json:"Number,omitempty" xml:"Number,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo) SetAddress(v string) *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo) SetAuthority(v string) *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo) SetBackImageUrl(v string) *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo) SetBirth(v string) *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo) SetEndDate(v string) *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo) SetFrontImageUrl(v string) *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo) SetName(v string) *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo) SetNationality(v string) *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo) SetNumber(v string) *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo) SetStartDate(v string) *VerifyMaterialResponseBodyResultObjectMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

type VerifyMaterialResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) CompareFacesWithOptions(request *CompareFacesRequest, runtime *util.RuntimeOptions) (_result *CompareFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceImageBase64)) {
		body["SourceImageBase64"] = request.SourceImageBase64
	}

	if !tea.BoolValue(util.IsUnset(request.SourceImageUrl)) {
		body["SourceImageUrl"] = request.SourceImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TargetImageBase64)) {
		body["TargetImageBase64"] = request.TargetImageBase64
	}

	if !tea.BoolValue(util.IsUnset(request.TargetImageUrl)) {
		body["TargetImageUrl"] = request.TargetImageUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CompareFaces"),
		Version:     tea.String("2020-11-12"),
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
		Version:     tea.String("2020-11-12"),
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

	if !tea.BoolValue(util.IsUnset(request.FaceRetainedImageUrl)) {
		query["FaceRetainedImageUrl"] = request.FaceRetainedImageUrl
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
		Version:     tea.String("2020-11-12"),
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

func (client *Client) DetectFaceAttributesWithOptions(request *DetectFaceAttributesRequest, runtime *util.RuntimeOptions) (_result *DetectFaceAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageFile)) {
		body["ImageFile"] = request.ImageFile
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectFaceAttributes"),
		Version:     tea.String("2020-11-12"),
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

func (client *Client) DetectFaceAttributesAdvance(request *DetectFaceAttributesAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectFaceAttributesResponse, _err error) {
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
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
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
	detectFaceAttributesReq := &DetectFaceAttributesRequest{}
	openapiutil.Convert(request, detectFaceAttributesReq)
	if !tea.BoolValue(util.IsUnset(request.ImageFileObject)) {
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
			Content:     request.ImageFileObject,
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
		detectFaceAttributesReq.ImageFile = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectFaceAttributesResp, _err := client.DetectFaceAttributesWithOptions(detectFaceAttributesReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectFaceAttributesResp
	return _result, _err
}

func (client *Client) LivenessDetectWithOptions(request *LivenessDetectRequest, runtime *util.RuntimeOptions) (_result *LivenessDetectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.MediaCategory)) {
		body["MediaCategory"] = request.MediaCategory
	}

	if !tea.BoolValue(util.IsUnset(request.MediaFile)) {
		body["MediaFile"] = request.MediaFile
	}

	if !tea.BoolValue(util.IsUnset(request.MediaUrl)) {
		body["MediaUrl"] = request.MediaUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LivenessDetect"),
		Version:     tea.String("2020-11-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LivenessDetectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LivenessDetect(request *LivenessDetectRequest) (_result *LivenessDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LivenessDetectResponse{}
	_body, _err := client.LivenessDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LivenessDetectAdvance(request *LivenessDetectAdvanceRequest, runtime *util.RuntimeOptions) (_result *LivenessDetectResponse, _err error) {
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
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
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
	livenessDetectReq := &LivenessDetectRequest{}
	openapiutil.Convert(request, livenessDetectReq)
	if !tea.BoolValue(util.IsUnset(request.MediaFileObject)) {
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
			Content:     request.MediaFileObject,
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
		livenessDetectReq.MediaFile = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	livenessDetectResp, _err := client.LivenessDetectWithOptions(livenessDetectReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = livenessDetectResp
	return _result, _err
}

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
		Version:     tea.String("2020-11-12"),
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
