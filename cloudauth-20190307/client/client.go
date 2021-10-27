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

type CompareFaceVerifyRequest struct {
	Crop                         *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	OuterOrderNo                 *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode                  *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SceneId                      *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SourceCertifyId              *string `json:"SourceCertifyId,omitempty" xml:"SourceCertifyId,omitempty"`
	SourceFaceContrastPicture    *string `json:"SourceFaceContrastPicture,omitempty" xml:"SourceFaceContrastPicture,omitempty"`
	SourceFaceContrastPictureUrl *string `json:"SourceFaceContrastPictureUrl,omitempty" xml:"SourceFaceContrastPictureUrl,omitempty"`
	SourceOssBucketName          *string `json:"SourceOssBucketName,omitempty" xml:"SourceOssBucketName,omitempty"`
	SourceOssObjectName          *string `json:"SourceOssObjectName,omitempty" xml:"SourceOssObjectName,omitempty"`
	TargetCertifyId              *string `json:"TargetCertifyId,omitempty" xml:"TargetCertifyId,omitempty"`
	TargetFaceContrastPicture    *string `json:"TargetFaceContrastPicture,omitempty" xml:"TargetFaceContrastPicture,omitempty"`
	TargetFaceContrastPictureUrl *string `json:"TargetFaceContrastPictureUrl,omitempty" xml:"TargetFaceContrastPictureUrl,omitempty"`
	TargetOssBucketName          *string `json:"TargetOssBucketName,omitempty" xml:"TargetOssBucketName,omitempty"`
	TargetOssObjectName          *string `json:"TargetOssObjectName,omitempty" xml:"TargetOssObjectName,omitempty"`
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
	Code         *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
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
	CertifyId   *string  `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Passed      *string  `json:"Passed,omitempty" xml:"Passed,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CompareFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CompareFaceVerifyResponse) SetBody(v *CompareFaceVerifyResponseBody) *CompareFaceVerifyResponse {
	s.Body = v
	return s
}

type CompareFacesRequest struct {
	SourceImageType  *string `json:"SourceImageType,omitempty" xml:"SourceImageType,omitempty"`
	SourceImageValue *string `json:"SourceImageValue,omitempty" xml:"SourceImageValue,omitempty"`
	TargetImageType  *string `json:"TargetImageType,omitempty" xml:"TargetImageType,omitempty"`
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
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CompareFacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
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
	ConfidenceThresholds *string  `json:"ConfidenceThresholds,omitempty" xml:"ConfidenceThresholds,omitempty"`
	SimilarityScore      *float32 `json:"SimilarityScore,omitempty" xml:"SimilarityScore,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CompareFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CompareFacesResponse) SetBody(v *CompareFacesResponseBody) *CompareFacesResponse {
	s.Body = v
	return s
}

type ContrastFaceVerifyRequest struct {
	CertName               *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                 *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertType               *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertifyId              *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Crop                   *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DeviceToken            *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	FaceContrastFile       *string `json:"FaceContrastFile,omitempty" xml:"FaceContrastFile,omitempty"`
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
	FaceContrastFileObject io.Reader `json:"FaceContrastFileObject,omitempty" xml:"FaceContrastFileObject,omitempty" require:"true"`
	CertName               *string   `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                 *string   `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertType               *string   `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertifyId              *string   `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Crop                   *string   `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DeviceToken            *string   `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	FaceContrastPicture    *string   `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	FaceContrastPictureUrl *string   `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	Ip                     *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Model                  *string   `json:"Model,omitempty" xml:"Model,omitempty"`
	OssBucketName          *string   `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string   `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	OuterOrderNo           *string   `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string   `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SceneId                *int64    `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ContrastFaceVerifyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastFileObject(v io.Reader) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastFileObject = v
	return s
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
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
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
	CertifyId    *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ContrastFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ContrastFaceVerifyResponse) SetBody(v *ContrastFaceVerifyResponseBody) *ContrastFaceVerifyResponse {
	s.Body = v
	return s
}

type CreateAuthKeyRequest struct {
	AuthYears    *int32  `json:"AuthYears,omitempty" xml:"AuthYears,omitempty"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Test         *bool   `json:"Test,omitempty" xml:"Test,omitempty"`
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
	AuthKey   *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAuthKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAuthKeyResponse) SetBody(v *CreateAuthKeyResponseBody) *CreateAuthKeyResponse {
	s.Body = v
	return s
}

type CreateFaceConfigRequest struct {
	BizName  *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizType  *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateFaceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateFaceConfigRequest) SetBizName(v string) *CreateFaceConfigRequest {
	s.BizName = &v
	return s
}

func (s *CreateFaceConfigRequest) SetBizType(v string) *CreateFaceConfigRequest {
	s.BizType = &v
	return s
}

func (s *CreateFaceConfigRequest) SetLang(v string) *CreateFaceConfigRequest {
	s.Lang = &v
	return s
}

func (s *CreateFaceConfigRequest) SetSourceIp(v string) *CreateFaceConfigRequest {
	s.SourceIp = &v
	return s
}

type CreateFaceConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFaceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFaceConfigResponseBody) SetRequestId(v string) *CreateFaceConfigResponseBody {
	s.RequestId = &v
	return s
}

type CreateFaceConfigResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFaceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFaceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateFaceConfigResponse) SetHeaders(v map[string]*string) *CreateFaceConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateFaceConfigResponse) SetBody(v *CreateFaceConfigResponseBody) *CreateFaceConfigResponse {
	s.Body = v
	return s
}

type CreateRPSDKRequest struct {
	AppUrl   *string `json:"AppUrl,omitempty" xml:"AppUrl,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateRPSDKRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRPSDKRequest) GoString() string {
	return s.String()
}

func (s *CreateRPSDKRequest) SetAppUrl(v string) *CreateRPSDKRequest {
	s.AppUrl = &v
	return s
}

func (s *CreateRPSDKRequest) SetLang(v string) *CreateRPSDKRequest {
	s.Lang = &v
	return s
}

func (s *CreateRPSDKRequest) SetPlatform(v string) *CreateRPSDKRequest {
	s.Platform = &v
	return s
}

func (s *CreateRPSDKRequest) SetSourceIp(v string) *CreateRPSDKRequest {
	s.SourceIp = &v
	return s
}

type CreateRPSDKResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateRPSDKResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRPSDKResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRPSDKResponseBody) SetRequestId(v string) *CreateRPSDKResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRPSDKResponseBody) SetTaskId(v string) *CreateRPSDKResponseBody {
	s.TaskId = &v
	return s
}

type CreateRPSDKResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRPSDKResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRPSDKResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRPSDKResponse) GoString() string {
	return s.String()
}

func (s *CreateRPSDKResponse) SetHeaders(v map[string]*string) *CreateRPSDKResponse {
	s.Headers = v
	return s
}

func (s *CreateRPSDKResponse) SetBody(v *CreateRPSDKResponseBody) *CreateRPSDKResponse {
	s.Body = v
	return s
}

type CreateVerifySDKRequest struct {
	AppUrl   *string `json:"AppUrl,omitempty" xml:"AppUrl,omitempty"`
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s CreateVerifySDKRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySDKRequest) GoString() string {
	return s.String()
}

func (s *CreateVerifySDKRequest) SetAppUrl(v string) *CreateVerifySDKRequest {
	s.AppUrl = &v
	return s
}

func (s *CreateVerifySDKRequest) SetPlatform(v string) *CreateVerifySDKRequest {
	s.Platform = &v
	return s
}

type CreateVerifySDKResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateVerifySDKResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySDKResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVerifySDKResponseBody) SetRequestId(v string) *CreateVerifySDKResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVerifySDKResponseBody) SetTaskId(v string) *CreateVerifySDKResponseBody {
	s.TaskId = &v
	return s
}

type CreateVerifySDKResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVerifySDKResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVerifySDKResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySDKResponse) GoString() string {
	return s.String()
}

func (s *CreateVerifySDKResponse) SetHeaders(v map[string]*string) *CreateVerifySDKResponse {
	s.Headers = v
	return s
}

func (s *CreateVerifySDKResponse) SetBody(v *CreateVerifySDKResponseBody) *CreateVerifySDKResponse {
	s.Body = v
	return s
}

type CreateVerifySettingRequest struct {
	BizName     *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	GuideStep   *bool   `json:"GuideStep,omitempty" xml:"GuideStep,omitempty"`
	PrivacyStep *bool   `json:"PrivacyStep,omitempty" xml:"PrivacyStep,omitempty"`
	ResultStep  *bool   `json:"ResultStep,omitempty" xml:"ResultStep,omitempty"`
	Solution    *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
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
	BizName   *string   `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizType   *string   `json:"BizType,omitempty" xml:"BizType,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Solution  *string   `json:"Solution,omitempty" xml:"Solution,omitempty"`
	StepList  []*string `json:"StepList,omitempty" xml:"StepList,omitempty" type:"Repeated"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVerifySettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateVerifySettingResponse) SetBody(v *CreateVerifySettingResponseBody) *CreateVerifySettingResponse {
	s.Body = v
	return s
}

type CreateWhitelistRequest struct {
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	IdCardNum *string `json:"IdCardNum,omitempty" xml:"IdCardNum,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ValidDay  *string `json:"ValidDay,omitempty" xml:"ValidDay,omitempty"`
}

func (s CreateWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWhitelistRequest) GoString() string {
	return s.String()
}

func (s *CreateWhitelistRequest) SetBizId(v string) *CreateWhitelistRequest {
	s.BizId = &v
	return s
}

func (s *CreateWhitelistRequest) SetBizType(v string) *CreateWhitelistRequest {
	s.BizType = &v
	return s
}

func (s *CreateWhitelistRequest) SetIdCardNum(v string) *CreateWhitelistRequest {
	s.IdCardNum = &v
	return s
}

func (s *CreateWhitelistRequest) SetLang(v string) *CreateWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *CreateWhitelistRequest) SetSourceIp(v string) *CreateWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateWhitelistRequest) SetValidDay(v string) *CreateWhitelistRequest {
	s.ValidDay = &v
	return s
}

type CreateWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWhitelistResponseBody) SetRequestId(v string) *CreateWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type CreateWhitelistResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWhitelistResponse) GoString() string {
	return s.String()
}

func (s *CreateWhitelistResponse) SetHeaders(v map[string]*string) *CreateWhitelistResponse {
	s.Headers = v
	return s
}

func (s *CreateWhitelistResponse) SetBody(v *CreateWhitelistResponseBody) *CreateWhitelistResponse {
	s.Body = v
	return s
}

type CreateWhitelistSettingRequest struct {
	CertNo      *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertifyId   *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SceneId     *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ValidDay    *int32  `json:"ValidDay,omitempty" xml:"ValidDay,omitempty"`
}

func (s CreateWhitelistSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWhitelistSettingRequest) GoString() string {
	return s.String()
}

func (s *CreateWhitelistSettingRequest) SetCertNo(v string) *CreateWhitelistSettingRequest {
	s.CertNo = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetCertifyId(v string) *CreateWhitelistSettingRequest {
	s.CertifyId = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetLang(v string) *CreateWhitelistSettingRequest {
	s.Lang = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetSceneId(v int64) *CreateWhitelistSettingRequest {
	s.SceneId = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetServiceCode(v string) *CreateWhitelistSettingRequest {
	s.ServiceCode = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetSourceIp(v string) *CreateWhitelistSettingRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetValidDay(v int32) *CreateWhitelistSettingRequest {
	s.ValidDay = &v
	return s
}

type CreateWhitelistSettingResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *bool   `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s CreateWhitelistSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWhitelistSettingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWhitelistSettingResponseBody) SetRequestId(v string) *CreateWhitelistSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWhitelistSettingResponseBody) SetResultObject(v bool) *CreateWhitelistSettingResponseBody {
	s.ResultObject = &v
	return s
}

type CreateWhitelistSettingResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWhitelistSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWhitelistSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWhitelistSettingResponse) GoString() string {
	return s.String()
}

func (s *CreateWhitelistSettingResponse) SetHeaders(v map[string]*string) *CreateWhitelistSettingResponse {
	s.Headers = v
	return s
}

func (s *CreateWhitelistSettingResponse) SetBody(v *CreateWhitelistSettingResponseBody) *CreateWhitelistSettingResponse {
	s.Body = v
	return s
}

type DeleteWhitelistRequest struct {
	Ids      *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistRequest) SetIds(v string) *DeleteWhitelistRequest {
	s.Ids = &v
	return s
}

func (s *DeleteWhitelistRequest) SetLang(v string) *DeleteWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *DeleteWhitelistRequest) SetSourceIp(v string) *DeleteWhitelistRequest {
	s.SourceIp = &v
	return s
}

type DeleteWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistResponseBody) SetRequestId(v string) *DeleteWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWhitelistResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistResponse) SetHeaders(v map[string]*string) *DeleteWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DeleteWhitelistResponse) SetBody(v *DeleteWhitelistResponseBody) *DeleteWhitelistResponse {
	s.Body = v
	return s
}

type DeleteWhitelistSettingRequest struct {
	Ids         *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteWhitelistSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhitelistSettingRequest) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistSettingRequest) SetIds(v string) *DeleteWhitelistSettingRequest {
	s.Ids = &v
	return s
}

func (s *DeleteWhitelistSettingRequest) SetLang(v string) *DeleteWhitelistSettingRequest {
	s.Lang = &v
	return s
}

func (s *DeleteWhitelistSettingRequest) SetServiceCode(v string) *DeleteWhitelistSettingRequest {
	s.ServiceCode = &v
	return s
}

func (s *DeleteWhitelistSettingRequest) SetSourceIp(v string) *DeleteWhitelistSettingRequest {
	s.SourceIp = &v
	return s
}

type DeleteWhitelistSettingResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *bool   `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s DeleteWhitelistSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhitelistSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistSettingResponseBody) SetRequestId(v string) *DeleteWhitelistSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWhitelistSettingResponseBody) SetResultObject(v bool) *DeleteWhitelistSettingResponseBody {
	s.ResultObject = &v
	return s
}

type DeleteWhitelistSettingResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWhitelistSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWhitelistSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhitelistSettingResponse) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistSettingResponse) SetHeaders(v map[string]*string) *DeleteWhitelistSettingResponse {
	s.Headers = v
	return s
}

func (s *DeleteWhitelistSettingResponse) SetBody(v *DeleteWhitelistSettingResponseBody) *DeleteWhitelistSettingResponse {
	s.Body = v
	return s
}

type DescribeAppInfoRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Platform    *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s DescribeAppInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoRequest) SetCurrentPage(v int32) *DescribeAppInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAppInfoRequest) SetPageSize(v int32) *DescribeAppInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppInfoRequest) SetPlatform(v string) *DescribeAppInfoRequest {
	s.Platform = &v
	return s
}

type DescribeAppInfoResponseBody struct {
	AppInfoList []*DescribeAppInfoResponseBodyAppInfoList `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
	CurrentPage *int32                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponseBody) SetAppInfoList(v []*DescribeAppInfoResponseBodyAppInfoList) *DescribeAppInfoResponseBody {
	s.AppInfoList = v
	return s
}

func (s *DescribeAppInfoResponseBody) SetCurrentPage(v int32) *DescribeAppInfoResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAppInfoResponseBody) SetPageSize(v int32) *DescribeAppInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppInfoResponseBody) SetRequestId(v string) *DescribeAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppInfoResponseBody) SetTotalCount(v int32) *DescribeAppInfoResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAppInfoResponseBodyAppInfoList struct {
	DebugPackageInfo *DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo `json:"DebugPackageInfo,omitempty" xml:"DebugPackageInfo,omitempty" type:"Struct"`
	EndDate          *string                                                 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Icon             *string                                                 `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Id               *int64                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name             *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	PackageInfo      *DescribeAppInfoResponseBodyAppInfoListPackageInfo      `json:"PackageInfo,omitempty" xml:"PackageInfo,omitempty" type:"Struct"`
	PackageName      *string                                                 `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	StartDate        *string                                                 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Type             *int32                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAppInfoResponseBodyAppInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponseBodyAppInfoList) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetDebugPackageInfo(v *DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo) *DescribeAppInfoResponseBodyAppInfoList {
	s.DebugPackageInfo = v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetEndDate(v string) *DescribeAppInfoResponseBodyAppInfoList {
	s.EndDate = &v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetIcon(v string) *DescribeAppInfoResponseBodyAppInfoList {
	s.Icon = &v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetId(v int64) *DescribeAppInfoResponseBodyAppInfoList {
	s.Id = &v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetName(v string) *DescribeAppInfoResponseBodyAppInfoList {
	s.Name = &v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetPackageInfo(v *DescribeAppInfoResponseBodyAppInfoListPackageInfo) *DescribeAppInfoResponseBodyAppInfoList {
	s.PackageInfo = v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetPackageName(v string) *DescribeAppInfoResponseBodyAppInfoList {
	s.PackageName = &v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetStartDate(v string) *DescribeAppInfoResponseBodyAppInfoList {
	s.StartDate = &v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetType(v int32) *DescribeAppInfoResponseBodyAppInfoList {
	s.Type = &v
	return s
}

type DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo) SetVersion(v string) *DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo {
	s.Version = &v
	return s
}

type DescribeAppInfoResponseBodyAppInfoListPackageInfo struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAppInfoResponseBodyAppInfoListPackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponseBodyAppInfoListPackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponseBodyAppInfoListPackageInfo) SetVersion(v string) *DescribeAppInfoResponseBodyAppInfoListPackageInfo {
	s.Version = &v
	return s
}

type DescribeAppInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponse) SetHeaders(v map[string]*string) *DescribeAppInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppInfoResponse) SetBody(v *DescribeAppInfoResponseBody) *DescribeAppInfoResponse {
	s.Body = v
	return s
}

type DescribeDeviceInfoRequest struct {
	BizType         *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceId        *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ExpiredEndDay   *string `json:"ExpiredEndDay,omitempty" xml:"ExpiredEndDay,omitempty"`
	ExpiredStartDay *string `json:"ExpiredStartDay,omitempty" xml:"ExpiredStartDay,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserDeviceId    *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
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
	CurrentPage    *int32                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceInfoList *DescribeDeviceInfoResponseBodyDeviceInfoList `json:"DeviceInfoList,omitempty" xml:"DeviceInfoList,omitempty" type:"Struct"`
	PageSize       *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	BeginDay     *string `json:"BeginDay,omitempty" xml:"BeginDay,omitempty"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ExpiredDay   *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDeviceInfoResponse) SetBody(v *DescribeDeviceInfoResponseBody) *DescribeDeviceInfoResponse {
	s.Body = v
	return s
}

type DescribeFaceConfigRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeFaceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaceConfigRequest) SetLang(v string) *DescribeFaceConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFaceConfigRequest) SetSourceIp(v string) *DescribeFaceConfigRequest {
	s.SourceIp = &v
	return s
}

type DescribeFaceConfigResponseBody struct {
	Items     []*DescribeFaceConfigResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFaceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaceConfigResponseBody) SetItems(v []*DescribeFaceConfigResponseBodyItems) *DescribeFaceConfigResponseBody {
	s.Items = v
	return s
}

func (s *DescribeFaceConfigResponseBody) SetRequestId(v string) *DescribeFaceConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFaceConfigResponseBodyItems struct {
	BizName    *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizType    *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	GmtUpdated *int64  `json:"GmtUpdated,omitempty" xml:"GmtUpdated,omitempty"`
}

func (s DescribeFaceConfigResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceConfigResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeFaceConfigResponseBodyItems) SetBizName(v string) *DescribeFaceConfigResponseBodyItems {
	s.BizName = &v
	return s
}

func (s *DescribeFaceConfigResponseBodyItems) SetBizType(v string) *DescribeFaceConfigResponseBodyItems {
	s.BizType = &v
	return s
}

func (s *DescribeFaceConfigResponseBodyItems) SetGmtUpdated(v int64) *DescribeFaceConfigResponseBodyItems {
	s.GmtUpdated = &v
	return s
}

type DescribeFaceConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFaceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFaceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaceConfigResponse) SetHeaders(v map[string]*string) *DescribeFaceConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaceConfigResponse) SetBody(v *DescribeFaceConfigResponseBody) *DescribeFaceConfigResponse {
	s.Body = v
	return s
}

type DescribeFaceUsageRequest struct {
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeFaceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaceUsageRequest) SetEndDate(v string) *DescribeFaceUsageRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeFaceUsageRequest) SetStartDate(v string) *DescribeFaceUsageRequest {
	s.StartDate = &v
	return s
}

type DescribeFaceUsageResponseBody struct {
	FaceUsageList []*DescribeFaceUsageResponseBodyFaceUsageList `json:"FaceUsageList,omitempty" xml:"FaceUsageList,omitempty" type:"Repeated"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFaceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaceUsageResponseBody) SetFaceUsageList(v []*DescribeFaceUsageResponseBodyFaceUsageList) *DescribeFaceUsageResponseBody {
	s.FaceUsageList = v
	return s
}

func (s *DescribeFaceUsageResponseBody) SetRequestId(v string) *DescribeFaceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaceUsageResponseBody) SetTotalCount(v int32) *DescribeFaceUsageResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeFaceUsageResponseBodyFaceUsageList struct {
	Date       *string `json:"Date,omitempty" xml:"Date,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFaceUsageResponseBodyFaceUsageList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceUsageResponseBodyFaceUsageList) GoString() string {
	return s.String()
}

func (s *DescribeFaceUsageResponseBodyFaceUsageList) SetDate(v string) *DescribeFaceUsageResponseBodyFaceUsageList {
	s.Date = &v
	return s
}

func (s *DescribeFaceUsageResponseBodyFaceUsageList) SetTotalCount(v int64) *DescribeFaceUsageResponseBodyFaceUsageList {
	s.TotalCount = &v
	return s
}

type DescribeFaceUsageResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFaceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFaceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaceUsageResponse) SetHeaders(v map[string]*string) *DescribeFaceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaceUsageResponse) SetBody(v *DescribeFaceUsageResponseBody) *DescribeFaceUsageResponse {
	s.Body = v
	return s
}

type DescribeFaceVerifyRequest struct {
	CertifyId         *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	PictureReturnType *string `json:"PictureReturnType,omitempty" xml:"PictureReturnType,omitempty"`
	SceneId           *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
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
	DeviceToken  *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s DescribeFaceVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
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

type DescribeFaceVerifyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeFaceVerifyResponse) SetBody(v *DescribeFaceVerifyResponseBody) *DescribeFaceVerifyResponse {
	s.Body = v
	return s
}

type DescribeOssUploadTokenResponseBody struct {
	OssUploadToken *DescribeOssUploadTokenResponseBodyOssUploadToken `json:"OssUploadToken,omitempty" xml:"OssUploadToken,omitempty" type:"Struct"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	Expired  *int64  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Secret   *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	Token    *string `json:"Token,omitempty" xml:"Token,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOssUploadTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeOssUploadTokenResponse) SetBody(v *DescribeOssUploadTokenResponseBody) *DescribeOssUploadTokenResponse {
	s.Body = v
	return s
}

type DescribeRPSDKRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRPSDKRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRPSDKRequest) GoString() string {
	return s.String()
}

func (s *DescribeRPSDKRequest) SetLang(v string) *DescribeRPSDKRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRPSDKRequest) SetSourceIp(v string) *DescribeRPSDKRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRPSDKRequest) SetTaskId(v string) *DescribeRPSDKRequest {
	s.TaskId = &v
	return s
}

type DescribeRPSDKResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SdkUrl    *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
}

func (s DescribeRPSDKResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRPSDKResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRPSDKResponseBody) SetRequestId(v string) *DescribeRPSDKResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRPSDKResponseBody) SetSdkUrl(v string) *DescribeRPSDKResponseBody {
	s.SdkUrl = &v
	return s
}

type DescribeRPSDKResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRPSDKResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRPSDKResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRPSDKResponse) GoString() string {
	return s.String()
}

func (s *DescribeRPSDKResponse) SetHeaders(v map[string]*string) *DescribeRPSDKResponse {
	s.Headers = v
	return s
}

func (s *DescribeRPSDKResponse) SetBody(v *DescribeRPSDKResponseBody) *DescribeRPSDKResponse {
	s.Body = v
	return s
}

type DescribeSdkUrlRequest struct {
	Debug *bool  `json:"Debug,omitempty" xml:"Debug,omitempty"`
	Id    *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeSdkUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSdkUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeSdkUrlRequest) SetDebug(v bool) *DescribeSdkUrlRequest {
	s.Debug = &v
	return s
}

func (s *DescribeSdkUrlRequest) SetId(v int64) *DescribeSdkUrlRequest {
	s.Id = &v
	return s
}

type DescribeSdkUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SdkUrl    *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
}

func (s DescribeSdkUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSdkUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSdkUrlResponseBody) SetRequestId(v string) *DescribeSdkUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSdkUrlResponseBody) SetSdkUrl(v string) *DescribeSdkUrlResponseBody {
	s.SdkUrl = &v
	return s
}

type DescribeSdkUrlResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSdkUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSdkUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSdkUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeSdkUrlResponse) SetHeaders(v map[string]*string) *DescribeSdkUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeSdkUrlResponse) SetBody(v *DescribeSdkUrlResponseBody) *DescribeSdkUrlResponse {
	s.Body = v
	return s
}

type DescribeUpdatePackageResultRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeUpdatePackageResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultRequest) SetTaskId(v string) *DescribeUpdatePackageResultRequest {
	s.TaskId = &v
	return s
}

type DescribeUpdatePackageResultResponseBody struct {
	AppInfo   *DescribeUpdatePackageResultResponseBodyAppInfo `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUpdatePackageResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponseBody) SetAppInfo(v *DescribeUpdatePackageResultResponseBodyAppInfo) *DescribeUpdatePackageResultResponseBody {
	s.AppInfo = v
	return s
}

func (s *DescribeUpdatePackageResultResponseBody) SetRequestId(v string) *DescribeUpdatePackageResultResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUpdatePackageResultResponseBodyAppInfo struct {
	DebugPackageInfo *DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo `json:"DebugPackageInfo,omitempty" xml:"DebugPackageInfo,omitempty" type:"Struct"`
	EndDate          *string                                                         `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Icon             *string                                                         `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Id               *int64                                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	Name             *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	PackageInfo      *DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo      `json:"PackageInfo,omitempty" xml:"PackageInfo,omitempty" type:"Struct"`
	PackageName      *string                                                         `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	StartDate        *string                                                         `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Type             *int32                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUpdatePackageResultResponseBodyAppInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponseBodyAppInfo) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetDebugPackageInfo(v *DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.DebugPackageInfo = v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetEndDate(v string) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.EndDate = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetIcon(v string) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.Icon = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetId(v int64) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.Id = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetName(v string) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.Name = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetPackageInfo(v *DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.PackageInfo = v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetPackageName(v string) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.PackageName = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetStartDate(v string) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.StartDate = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetType(v int32) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.Type = &v
	return s
}

type DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo) SetVersion(v string) *DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo {
	s.Version = &v
	return s
}

type DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo) SetVersion(v string) *DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo {
	s.Version = &v
	return s
}

type DescribeUpdatePackageResultResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUpdatePackageResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUpdatePackageResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponse) SetHeaders(v map[string]*string) *DescribeUpdatePackageResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeUpdatePackageResultResponse) SetBody(v *DescribeUpdatePackageResultResponseBody) *DescribeUpdatePackageResultResponse {
	s.Body = v
	return s
}

type DescribeUploadInfoRequest struct {
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
}

func (s DescribeUploadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeUploadInfoRequest) SetBiz(v string) *DescribeUploadInfoRequest {
	s.Biz = &v
	return s
}

type DescribeUploadInfoResponseBody struct {
	Accessid  *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	Expire    *int64  `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Folder    *string `json:"Folder,omitempty" xml:"Folder,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s DescribeUploadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUploadInfoResponseBody) SetAccessid(v string) *DescribeUploadInfoResponseBody {
	s.Accessid = &v
	return s
}

func (s *DescribeUploadInfoResponseBody) SetExpire(v int64) *DescribeUploadInfoResponseBody {
	s.Expire = &v
	return s
}

func (s *DescribeUploadInfoResponseBody) SetFolder(v string) *DescribeUploadInfoResponseBody {
	s.Folder = &v
	return s
}

func (s *DescribeUploadInfoResponseBody) SetHost(v string) *DescribeUploadInfoResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeUploadInfoResponseBody) SetPolicy(v string) *DescribeUploadInfoResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeUploadInfoResponseBody) SetRequestId(v string) *DescribeUploadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUploadInfoResponseBody) SetSignature(v string) *DescribeUploadInfoResponseBody {
	s.Signature = &v
	return s
}

type DescribeUploadInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUploadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeUploadInfoResponse) SetHeaders(v map[string]*string) *DescribeUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeUploadInfoResponse) SetBody(v *DescribeUploadInfoResponseBody) *DescribeUploadInfoResponse {
	s.Body = v
	return s
}

type DescribeUserStatusResponseBody struct {
	Enabled   *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusResponseBody) SetEnabled(v bool) *DescribeUserStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeUserStatusResponseBody) SetRequestId(v string) *DescribeUserStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserStatusResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusResponse) SetHeaders(v map[string]*string) *DescribeUserStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserStatusResponse) SetBody(v *DescribeUserStatusResponseBody) *DescribeUserStatusResponse {
	s.Body = v
	return s
}

type DescribeVerifyRecordsRequest struct {
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndDate     *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	IdCardNum   *string `json:"IdCardNum,omitempty" xml:"IdCardNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryId     *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StatusList  *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	TotalCount  *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVerifyRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyRecordsRequest) SetBizId(v string) *DescribeVerifyRecordsRequest {
	s.BizId = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetBizType(v string) *DescribeVerifyRecordsRequest {
	s.BizType = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetCurrentPage(v int32) *DescribeVerifyRecordsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetEndDate(v string) *DescribeVerifyRecordsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetIdCardNum(v string) *DescribeVerifyRecordsRequest {
	s.IdCardNum = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetPageSize(v int32) *DescribeVerifyRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetQueryId(v string) *DescribeVerifyRecordsRequest {
	s.QueryId = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetStartDate(v string) *DescribeVerifyRecordsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetStatusList(v string) *DescribeVerifyRecordsRequest {
	s.StatusList = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetTotalCount(v int32) *DescribeVerifyRecordsRequest {
	s.TotalCount = &v
	return s
}

type DescribeVerifyRecordsResponseBody struct {
	CurrentPage *int32                                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryId     *string                                         `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	RecordsList []*DescribeVerifyRecordsResponseBodyRecordsList `json:"RecordsList,omitempty" xml:"RecordsList,omitempty" type:"Repeated"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVerifyRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyRecordsResponseBody) SetCurrentPage(v int32) *DescribeVerifyRecordsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBody) SetPageSize(v int32) *DescribeVerifyRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBody) SetQueryId(v string) *DescribeVerifyRecordsResponseBody {
	s.QueryId = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBody) SetRecordsList(v []*DescribeVerifyRecordsResponseBodyRecordsList) *DescribeVerifyRecordsResponseBody {
	s.RecordsList = v
	return s
}

func (s *DescribeVerifyRecordsResponseBody) SetRequestId(v string) *DescribeVerifyRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBody) SetTotalCount(v int32) *DescribeVerifyRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeVerifyRecordsResponseBodyRecordsList struct {
	AuthorityComparisonScore  *float32                                              `json:"AuthorityComparisonScore,omitempty" xml:"AuthorityComparisonScore,omitempty"`
	BizId                     *string                                               `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType                   *string                                               `json:"BizType,omitempty" xml:"BizType,omitempty"`
	DataStats                 *string                                               `json:"DataStats,omitempty" xml:"DataStats,omitempty"`
	FinishTime                *int64                                                `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	IdCardFaceComparisonScore *float32                                              `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	Material                  *DescribeVerifyRecordsResponseBodyRecordsListMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	Status                    *int32                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	VerifyId                  *string                                               `json:"VerifyId,omitempty" xml:"VerifyId,omitempty"`
}

func (s DescribeVerifyRecordsResponseBodyRecordsList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyRecordsResponseBodyRecordsList) GoString() string {
	return s.String()
}

func (s *DescribeVerifyRecordsResponseBodyRecordsList) SetAuthorityComparisonScore(v float32) *DescribeVerifyRecordsResponseBodyRecordsList {
	s.AuthorityComparisonScore = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsList) SetBizId(v string) *DescribeVerifyRecordsResponseBodyRecordsList {
	s.BizId = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsList) SetBizType(v string) *DescribeVerifyRecordsResponseBodyRecordsList {
	s.BizType = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsList) SetDataStats(v string) *DescribeVerifyRecordsResponseBodyRecordsList {
	s.DataStats = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsList) SetFinishTime(v int64) *DescribeVerifyRecordsResponseBodyRecordsList {
	s.FinishTime = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsList) SetIdCardFaceComparisonScore(v float32) *DescribeVerifyRecordsResponseBodyRecordsList {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsList) SetMaterial(v *DescribeVerifyRecordsResponseBodyRecordsListMaterial) *DescribeVerifyRecordsResponseBodyRecordsList {
	s.Material = v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsList) SetStatus(v int32) *DescribeVerifyRecordsResponseBodyRecordsList {
	s.Status = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsList) SetVerifyId(v string) *DescribeVerifyRecordsResponseBodyRecordsList {
	s.VerifyId = &v
	return s
}

type DescribeVerifyRecordsResponseBodyRecordsListMaterial struct {
	FaceImageUrl *string                                                         `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	IdCardInfo   *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	IdCardName   *string                                                         `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	IdCardNumber *string                                                         `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
}

func (s DescribeVerifyRecordsResponseBodyRecordsListMaterial) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyRecordsResponseBodyRecordsListMaterial) GoString() string {
	return s.String()
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterial) SetFaceImageUrl(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterial) SetIdCardInfo(v *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) *DescribeVerifyRecordsResponseBodyRecordsListMaterial {
	s.IdCardInfo = v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterial) SetIdCardName(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterial {
	s.IdCardName = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterial) SetIdCardNumber(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterial {
	s.IdCardNumber = &v
	return s
}

type DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo struct {
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Authority     *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	BackImageUrl  *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty"`
	Birth         *string `json:"Birth,omitempty" xml:"Birth,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	Number        *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Sex           *string `json:"Sex,omitempty" xml:"Sex,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) SetAddress(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) SetAuthority(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) SetBackImageUrl(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) SetBirth(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) SetEndDate(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) SetFrontImageUrl(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) SetName(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) SetNationality(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) SetNumber(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) SetSex(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo {
	s.Sex = &v
	return s
}

func (s *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo) SetStartDate(v string) *DescribeVerifyRecordsResponseBodyRecordsListMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

type DescribeVerifyRecordsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVerifyRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVerifyRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyRecordsResponse) SetHeaders(v map[string]*string) *DescribeVerifyRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyRecordsResponse) SetBody(v *DescribeVerifyRecordsResponseBody) *DescribeVerifyRecordsResponse {
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
	AuthorityComparisionScore *float32                                  `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty"`
	FaceComparisonScore       *float32                                  `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	IdCardFaceComparisonScore *float32                                  `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	Material                  *DescribeVerifyResultResponseBodyMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	RequestId                 *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VerifyStatus              *int32                                    `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
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
	FaceGlobalUrl *string                                             `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty"`
	FaceImageUrl  *string                                             `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	FaceMask      *bool                                               `json:"FaceMask,omitempty" xml:"FaceMask,omitempty"`
	FaceQuality   *string                                             `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	IdCardInfo    *DescribeVerifyResultResponseBodyMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	IdCardName    *string                                             `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	IdCardNumber  *string                                             `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	VideoUrls     []*string                                           `json:"VideoUrls,omitempty" xml:"VideoUrls,omitempty" type:"Repeated"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeVerifyResultResponse) SetBody(v *DescribeVerifyResultResponseBody) *DescribeVerifyResultResponse {
	s.Body = v
	return s
}

type DescribeVerifySDKRequest struct {
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SdkUrl    *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVerifySDKResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeVerifySDKResponse) SetBody(v *DescribeVerifySDKResponseBody) *DescribeVerifySDKResponse {
	s.Body = v
	return s
}

type DescribeVerifySettingResponseBody struct {
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VerifySettingList []*DescribeVerifySettingResponseBodyVerifySettingList `json:"VerifySettingList,omitempty" xml:"VerifySettingList,omitempty" type:"Repeated"`
}

func (s DescribeVerifySettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifySettingResponseBody) SetRequestId(v string) *DescribeVerifySettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifySettingResponseBody) SetVerifySettingList(v []*DescribeVerifySettingResponseBodyVerifySettingList) *DescribeVerifySettingResponseBody {
	s.VerifySettingList = v
	return s
}

type DescribeVerifySettingResponseBodyVerifySettingList struct {
	BizName  *string   `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizType  *string   `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Solution *string   `json:"Solution,omitempty" xml:"Solution,omitempty"`
	StepList []*string `json:"StepList,omitempty" xml:"StepList,omitempty" type:"Repeated"`
}

func (s DescribeVerifySettingResponseBodyVerifySettingList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySettingResponseBodyVerifySettingList) GoString() string {
	return s.String()
}

func (s *DescribeVerifySettingResponseBodyVerifySettingList) SetBizName(v string) *DescribeVerifySettingResponseBodyVerifySettingList {
	s.BizName = &v
	return s
}

func (s *DescribeVerifySettingResponseBodyVerifySettingList) SetBizType(v string) *DescribeVerifySettingResponseBodyVerifySettingList {
	s.BizType = &v
	return s
}

func (s *DescribeVerifySettingResponseBodyVerifySettingList) SetSolution(v string) *DescribeVerifySettingResponseBodyVerifySettingList {
	s.Solution = &v
	return s
}

func (s *DescribeVerifySettingResponseBodyVerifySettingList) SetStepList(v []*string) *DescribeVerifySettingResponseBodyVerifySettingList {
	s.StepList = v
	return s
}

type DescribeVerifySettingResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVerifySettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVerifySettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifySettingResponse) SetHeaders(v map[string]*string) *DescribeVerifySettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifySettingResponse) SetBody(v *DescribeVerifySettingResponseBody) *DescribeVerifySettingResponse {
	s.Body = v
	return s
}

type DescribeVerifyTokenRequest struct {
	BizId                *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CallbackSeed         *string `json:"CallbackSeed,omitempty" xml:"CallbackSeed,omitempty"`
	CallbackUrl          *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	FaceRetainedImageUrl *string `json:"FaceRetainedImageUrl,omitempty" xml:"FaceRetainedImageUrl,omitempty"`
	FailedRedirectUrl    *string `json:"FailedRedirectUrl,omitempty" xml:"FailedRedirectUrl,omitempty"`
	IdCardBackImageUrl   *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	IdCardFrontImageUrl  *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	IdCardNumber         *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PassedRedirectUrl    *string `json:"PassedRedirectUrl,omitempty" xml:"PassedRedirectUrl,omitempty"`
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
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VerifyPageUrl  *string                                        `json:"VerifyPageUrl,omitempty" xml:"VerifyPageUrl,omitempty"`
	VerifyToken    *string                                        `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	Expired  *int64  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Secret   *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	Token    *string `json:"Token,omitempty" xml:"Token,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVerifyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeVerifyTokenResponse) SetBody(v *DescribeVerifyTokenResponseBody) *DescribeVerifyTokenResponse {
	s.Body = v
	return s
}

type DescribeVerifyUsageRequest struct {
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeVerifyUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyUsageRequest) SetBizType(v string) *DescribeVerifyUsageRequest {
	s.BizType = &v
	return s
}

func (s *DescribeVerifyUsageRequest) SetEndDate(v string) *DescribeVerifyUsageRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifyUsageRequest) SetStartDate(v string) *DescribeVerifyUsageRequest {
	s.StartDate = &v
	return s
}

type DescribeVerifyUsageResponseBody struct {
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VerifyUsageList []*DescribeVerifyUsageResponseBodyVerifyUsageList `json:"VerifyUsageList,omitempty" xml:"VerifyUsageList,omitempty" type:"Repeated"`
}

func (s DescribeVerifyUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyUsageResponseBody) SetRequestId(v string) *DescribeVerifyUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyUsageResponseBody) SetTotalCount(v int32) *DescribeVerifyUsageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVerifyUsageResponseBody) SetVerifyUsageList(v []*DescribeVerifyUsageResponseBodyVerifyUsageList) *DescribeVerifyUsageResponseBody {
	s.VerifyUsageList = v
	return s
}

type DescribeVerifyUsageResponseBodyVerifyUsageList struct {
	BizType    *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Date       *string `json:"Date,omitempty" xml:"Date,omitempty"`
	FailCount  *int64  `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	PassCount  *int64  `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVerifyUsageResponseBodyVerifyUsageList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyUsageResponseBodyVerifyUsageList) GoString() string {
	return s.String()
}

func (s *DescribeVerifyUsageResponseBodyVerifyUsageList) SetBizType(v string) *DescribeVerifyUsageResponseBodyVerifyUsageList {
	s.BizType = &v
	return s
}

func (s *DescribeVerifyUsageResponseBodyVerifyUsageList) SetDate(v string) *DescribeVerifyUsageResponseBodyVerifyUsageList {
	s.Date = &v
	return s
}

func (s *DescribeVerifyUsageResponseBodyVerifyUsageList) SetFailCount(v int64) *DescribeVerifyUsageResponseBodyVerifyUsageList {
	s.FailCount = &v
	return s
}

func (s *DescribeVerifyUsageResponseBodyVerifyUsageList) SetPassCount(v int64) *DescribeVerifyUsageResponseBodyVerifyUsageList {
	s.PassCount = &v
	return s
}

func (s *DescribeVerifyUsageResponseBodyVerifyUsageList) SetTotalCount(v int64) *DescribeVerifyUsageResponseBodyVerifyUsageList {
	s.TotalCount = &v
	return s
}

type DescribeVerifyUsageResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVerifyUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVerifyUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyUsageResponse) SetHeaders(v map[string]*string) *DescribeVerifyUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyUsageResponse) SetBody(v *DescribeVerifyUsageResponseBody) *DescribeVerifyUsageResponse {
	s.Body = v
	return s
}

type DescribeWhitelistRequest struct {
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CurrentPage    *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	IdCardNum      *string `json:"IdCardNum,omitempty" xml:"IdCardNum,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Valid          *string `json:"Valid,omitempty" xml:"Valid,omitempty"`
	ValidEndDate   *string `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	ValidStartDate *string `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
}

func (s DescribeWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistRequest) SetBizId(v string) *DescribeWhitelistRequest {
	s.BizId = &v
	return s
}

func (s *DescribeWhitelistRequest) SetBizType(v string) *DescribeWhitelistRequest {
	s.BizType = &v
	return s
}

func (s *DescribeWhitelistRequest) SetCurrentPage(v int32) *DescribeWhitelistRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhitelistRequest) SetIdCardNum(v string) *DescribeWhitelistRequest {
	s.IdCardNum = &v
	return s
}

func (s *DescribeWhitelistRequest) SetLang(v string) *DescribeWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWhitelistRequest) SetPageSize(v int32) *DescribeWhitelistRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWhitelistRequest) SetSourceIp(v string) *DescribeWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWhitelistRequest) SetValid(v string) *DescribeWhitelistRequest {
	s.Valid = &v
	return s
}

func (s *DescribeWhitelistRequest) SetValidEndDate(v string) *DescribeWhitelistRequest {
	s.ValidEndDate = &v
	return s
}

func (s *DescribeWhitelistRequest) SetValidStartDate(v string) *DescribeWhitelistRequest {
	s.ValidStartDate = &v
	return s
}

type DescribeWhitelistResponseBody struct {
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*DescribeWhitelistResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageSize    *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistResponseBody) SetCurrentPage(v int32) *DescribeWhitelistResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhitelistResponseBody) SetItems(v []*DescribeWhitelistResponseBodyItems) *DescribeWhitelistResponseBody {
	s.Items = v
	return s
}

func (s *DescribeWhitelistResponseBody) SetPageSize(v int32) *DescribeWhitelistResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWhitelistResponseBody) SetRequestId(v string) *DescribeWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhitelistResponseBody) SetTotalCount(v int32) *DescribeWhitelistResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeWhitelistResponseBodyItems struct {
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	EndDate     *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IdCardNum   *string `json:"IdCardNum,omitempty" xml:"IdCardNum,omitempty"`
	StartDate   *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Uid         *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
	Valid       *int32  `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s DescribeWhitelistResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhitelistResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistResponseBodyItems) SetBizId(v string) *DescribeWhitelistResponseBodyItems {
	s.BizId = &v
	return s
}

func (s *DescribeWhitelistResponseBodyItems) SetBizType(v string) *DescribeWhitelistResponseBodyItems {
	s.BizType = &v
	return s
}

func (s *DescribeWhitelistResponseBodyItems) SetEndDate(v int64) *DescribeWhitelistResponseBodyItems {
	s.EndDate = &v
	return s
}

func (s *DescribeWhitelistResponseBodyItems) SetGmtCreate(v int64) *DescribeWhitelistResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeWhitelistResponseBodyItems) SetGmtModified(v int64) *DescribeWhitelistResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeWhitelistResponseBodyItems) SetId(v int64) *DescribeWhitelistResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeWhitelistResponseBodyItems) SetIdCardNum(v string) *DescribeWhitelistResponseBodyItems {
	s.IdCardNum = &v
	return s
}

func (s *DescribeWhitelistResponseBodyItems) SetStartDate(v int64) *DescribeWhitelistResponseBodyItems {
	s.StartDate = &v
	return s
}

func (s *DescribeWhitelistResponseBodyItems) SetUid(v int64) *DescribeWhitelistResponseBodyItems {
	s.Uid = &v
	return s
}

func (s *DescribeWhitelistResponseBodyItems) SetValid(v int32) *DescribeWhitelistResponseBodyItems {
	s.Valid = &v
	return s
}

type DescribeWhitelistResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistResponse) SetHeaders(v map[string]*string) *DescribeWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhitelistResponse) SetBody(v *DescribeWhitelistResponseBody) *DescribeWhitelistResponse {
	s.Body = v
	return s
}

type DescribeWhitelistSettingRequest struct {
	CertNo         *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertifyId      *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	CurrentPage    *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SceneId        *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	ServiceCode    *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ValidEndDate   *int64  `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	ValidStartDate *int64  `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
}

func (s DescribeWhitelistSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhitelistSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistSettingRequest) SetCertNo(v string) *DescribeWhitelistSettingRequest {
	s.CertNo = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetCertifyId(v string) *DescribeWhitelistSettingRequest {
	s.CertifyId = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetCurrentPage(v int32) *DescribeWhitelistSettingRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetLang(v string) *DescribeWhitelistSettingRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetPageSize(v int32) *DescribeWhitelistSettingRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetSceneId(v int64) *DescribeWhitelistSettingRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetServiceCode(v string) *DescribeWhitelistSettingRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetSourceIp(v string) *DescribeWhitelistSettingRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetStatus(v string) *DescribeWhitelistSettingRequest {
	s.Status = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetValidEndDate(v int64) *DescribeWhitelistSettingRequest {
	s.ValidEndDate = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetValidStartDate(v int64) *DescribeWhitelistSettingRequest {
	s.ValidStartDate = &v
	return s
}

type DescribeWhitelistSettingResponseBody struct {
	CurrentPage *int32                                       `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*DescribeWhitelistSettingResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageSize    *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWhitelistSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhitelistSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistSettingResponseBody) SetCurrentPage(v int32) *DescribeWhitelistSettingResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBody) SetItems(v []*DescribeWhitelistSettingResponseBodyItems) *DescribeWhitelistSettingResponseBody {
	s.Items = v
	return s
}

func (s *DescribeWhitelistSettingResponseBody) SetPageSize(v int32) *DescribeWhitelistSettingResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBody) SetRequestId(v string) *DescribeWhitelistSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBody) SetTotalCount(v int32) *DescribeWhitelistSettingResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeWhitelistSettingResponseBodyItems struct {
	CertNo         *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertifyId      *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	GmtCreate      *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified    *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	SceneId        *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ValidEndDate   *string `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	ValidStartDate *string `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
}

func (s DescribeWhitelistSettingResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhitelistSettingResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetCertNo(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.CertNo = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetCertifyId(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.CertifyId = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetGmtCreate(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetGmtModified(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetId(v int64) *DescribeWhitelistSettingResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetSceneId(v int64) *DescribeWhitelistSettingResponseBodyItems {
	s.SceneId = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetStatus(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetValidEndDate(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.ValidEndDate = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetValidStartDate(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.ValidStartDate = &v
	return s
}

type DescribeWhitelistSettingResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWhitelistSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWhitelistSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhitelistSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistSettingResponse) SetHeaders(v map[string]*string) *DescribeWhitelistSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhitelistSettingResponse) SetBody(v *DescribeWhitelistSettingResponseBody) *DescribeWhitelistSettingResponse {
	s.Body = v
	return s
}

type DetectFaceAttributesRequest struct {
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
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
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DetectFaceAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
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
	ImgHeight *int32                                         `json:"ImgHeight,omitempty" xml:"ImgHeight,omitempty"`
	ImgWidth  *int32                                         `json:"ImgWidth,omitempty" xml:"ImgWidth,omitempty"`
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
	Blur       *float32                                                                                     `json:"Blur,omitempty" xml:"Blur,omitempty"`
	Facequal   *float32                                                                                     `json:"Facequal,omitempty" xml:"Facequal,omitempty"`
	Facetype   *string                                                                                      `json:"Facetype,omitempty" xml:"Facetype,omitempty"`
	Glasses    *string                                                                                      `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	Headpose   *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose `json:"Headpose,omitempty" xml:"Headpose,omitempty" type:"Struct"`
	Integrity  *int32                                                                                       `json:"Integrity,omitempty" xml:"Integrity,omitempty"`
	Respirator *string                                                                                      `json:"Respirator,omitempty" xml:"Respirator,omitempty"`
	Smiling    *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling  `json:"Smiling,omitempty" xml:"Smiling,omitempty" type:"Struct"`
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
	PitchAngle *float32 `json:"PitchAngle,omitempty" xml:"PitchAngle,omitempty"`
	RollAngle  *float32 `json:"RollAngle,omitempty" xml:"RollAngle,omitempty"`
	YawAngle   *float32 `json:"YawAngle,omitempty" xml:"YawAngle,omitempty"`
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
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Value     *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectFaceAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectFaceAttributesResponse) SetBody(v *DetectFaceAttributesResponseBody) *DetectFaceAttributesResponse {
	s.Body = v
	return s
}

type InitDeviceRequest struct {
	AppVersion       *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BizData          *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
	CertifyId        *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	CertifyPrincipal *string `json:"CertifyPrincipal,omitempty" xml:"CertifyPrincipal,omitempty"`
	Channel          *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	DeviceToken      *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	Merchant         *string `json:"Merchant,omitempty" xml:"Merchant,omitempty"`
	MetaInfo         *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	OuterOrderNo     *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProduceNode      *string `json:"ProduceNode,omitempty" xml:"ProduceNode,omitempty"`
	ProductName      *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	UaToken          *string `json:"UaToken,omitempty" xml:"UaToken,omitempty"`
	WebUmidToken     *string `json:"WebUmidToken,omitempty" xml:"WebUmidToken,omitempty"`
}

func (s InitDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s InitDeviceRequest) GoString() string {
	return s.String()
}

func (s *InitDeviceRequest) SetAppVersion(v string) *InitDeviceRequest {
	s.AppVersion = &v
	return s
}

func (s *InitDeviceRequest) SetBizData(v string) *InitDeviceRequest {
	s.BizData = &v
	return s
}

func (s *InitDeviceRequest) SetCertifyId(v string) *InitDeviceRequest {
	s.CertifyId = &v
	return s
}

func (s *InitDeviceRequest) SetCertifyPrincipal(v string) *InitDeviceRequest {
	s.CertifyPrincipal = &v
	return s
}

func (s *InitDeviceRequest) SetChannel(v string) *InitDeviceRequest {
	s.Channel = &v
	return s
}

func (s *InitDeviceRequest) SetDeviceToken(v string) *InitDeviceRequest {
	s.DeviceToken = &v
	return s
}

func (s *InitDeviceRequest) SetMerchant(v string) *InitDeviceRequest {
	s.Merchant = &v
	return s
}

func (s *InitDeviceRequest) SetMetaInfo(v string) *InitDeviceRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitDeviceRequest) SetOuterOrderNo(v string) *InitDeviceRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *InitDeviceRequest) SetProduceNode(v string) *InitDeviceRequest {
	s.ProduceNode = &v
	return s
}

func (s *InitDeviceRequest) SetProductName(v string) *InitDeviceRequest {
	s.ProductName = &v
	return s
}

func (s *InitDeviceRequest) SetUaToken(v string) *InitDeviceRequest {
	s.UaToken = &v
	return s
}

func (s *InitDeviceRequest) SetWebUmidToken(v string) *InitDeviceRequest {
	s.WebUmidToken = &v
	return s
}

type InitDeviceResponseBody struct {
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *InitDeviceResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s InitDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *InitDeviceResponseBody) SetCode(v string) *InitDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *InitDeviceResponseBody) SetMessage(v string) *InitDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *InitDeviceResponseBody) SetRequestId(v string) *InitDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitDeviceResponseBody) SetResultObject(v *InitDeviceResponseBodyResultObject) *InitDeviceResponseBody {
	s.ResultObject = v
	return s
}

type InitDeviceResponseBodyResultObject struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	BucketName      *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	CertifyId       *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	ExtParams       *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	FileName        *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNamePrefix  *string `json:"FileNamePrefix,omitempty" xml:"FileNamePrefix,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OssEndPoint     *string `json:"OssEndPoint,omitempty" xml:"OssEndPoint,omitempty"`
	PresignedUrl    *string `json:"PresignedUrl,omitempty" xml:"PresignedUrl,omitempty"`
	Protocol        *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RetCode         *string `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	RetCodeSub      *string `json:"RetCodeSub,omitempty" xml:"RetCodeSub,omitempty"`
	RetMessageSub   *string `json:"RetMessageSub,omitempty" xml:"RetMessageSub,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s InitDeviceResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s InitDeviceResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *InitDeviceResponseBodyResultObject) SetAccessKeyId(v string) *InitDeviceResponseBodyResultObject {
	s.AccessKeyId = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetAccessKeySecret(v string) *InitDeviceResponseBodyResultObject {
	s.AccessKeySecret = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetBucketName(v string) *InitDeviceResponseBodyResultObject {
	s.BucketName = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetCertifyId(v string) *InitDeviceResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetExtParams(v string) *InitDeviceResponseBodyResultObject {
	s.ExtParams = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetFileName(v string) *InitDeviceResponseBodyResultObject {
	s.FileName = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetFileNamePrefix(v string) *InitDeviceResponseBodyResultObject {
	s.FileNamePrefix = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetMessage(v string) *InitDeviceResponseBodyResultObject {
	s.Message = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetOssEndPoint(v string) *InitDeviceResponseBodyResultObject {
	s.OssEndPoint = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetPresignedUrl(v string) *InitDeviceResponseBodyResultObject {
	s.PresignedUrl = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetProtocol(v string) *InitDeviceResponseBodyResultObject {
	s.Protocol = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetRetCode(v string) *InitDeviceResponseBodyResultObject {
	s.RetCode = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetRetCodeSub(v string) *InitDeviceResponseBodyResultObject {
	s.RetCodeSub = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetRetMessageSub(v string) *InitDeviceResponseBodyResultObject {
	s.RetMessageSub = &v
	return s
}

func (s *InitDeviceResponseBodyResultObject) SetSecurityToken(v string) *InitDeviceResponseBodyResultObject {
	s.SecurityToken = &v
	return s
}

type InitDeviceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s InitDeviceResponse) GoString() string {
	return s.String()
}

func (s *InitDeviceResponse) SetHeaders(v map[string]*string) *InitDeviceResponse {
	s.Headers = v
	return s
}

func (s *InitDeviceResponse) SetBody(v *InitDeviceResponseBody) *InitDeviceResponse {
	s.Body = v
	return s
}

type InitFaceVerifyRequest struct {
	CallbackToken          *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	CallbackUrl            *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	CertName               *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                 *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertType               *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertifyId              *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	CertifyUrlType         *string `json:"CertifyUrlType,omitempty" xml:"CertifyUrlType,omitempty"`
	Crop                   *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	FaceContrastPicture    *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	Ip                     *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	MetaInfo               *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	Mobile                 *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Model                  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OssBucketName          *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	OuterOrderNo           *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ReturnUrl              *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	SceneId                *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s InitFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyRequest) GoString() string {
	return s.String()
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

func (s *InitFaceVerifyRequest) SetCertifyUrlType(v string) *InitFaceVerifyRequest {
	s.CertifyUrlType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCrop(v string) *InitFaceVerifyRequest {
	s.Crop = &v
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

func (s *InitFaceVerifyRequest) SetProductCode(v string) *InitFaceVerifyRequest {
	s.ProductCode = &v
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

func (s *InitFaceVerifyRequest) SetUserId(v string) *InitFaceVerifyRequest {
	s.UserId = &v
	return s
}

type InitFaceVerifyResponseBody struct {
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *InitFaceVerifyResponse) SetBody(v *InitFaceVerifyResponseBody) *InitFaceVerifyResponse {
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LivenessFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *LivenessFaceVerifyResponse) SetBody(v *LivenessFaceVerifyResponseBody) *LivenessFaceVerifyResponse {
	s.Body = v
	return s
}

type ModifyDeviceInfoRequest struct {
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ExpiredDay   *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
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
	BeginDay     *string `json:"BeginDay,omitempty" xml:"BeginDay,omitempty"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ExpiredDay   *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDeviceInfoResponse) SetBody(v *ModifyDeviceInfoResponseBody) *ModifyDeviceInfoResponse {
	s.Body = v
	return s
}

type UpdateAppPackageRequest struct {
	Debug      *bool   `json:"Debug,omitempty" xml:"Debug,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	Platform   *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s UpdateAppPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppPackageRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppPackageRequest) SetDebug(v bool) *UpdateAppPackageRequest {
	s.Debug = &v
	return s
}

func (s *UpdateAppPackageRequest) SetId(v int64) *UpdateAppPackageRequest {
	s.Id = &v
	return s
}

func (s *UpdateAppPackageRequest) SetPackageUrl(v string) *UpdateAppPackageRequest {
	s.PackageUrl = &v
	return s
}

func (s *UpdateAppPackageRequest) SetPlatform(v string) *UpdateAppPackageRequest {
	s.Platform = &v
	return s
}

type UpdateAppPackageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateAppPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppPackageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppPackageResponseBody) SetRequestId(v string) *UpdateAppPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppPackageResponseBody) SetTaskId(v string) *UpdateAppPackageResponseBody {
	s.TaskId = &v
	return s
}

type UpdateAppPackageResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppPackageResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppPackageResponse) SetHeaders(v map[string]*string) *UpdateAppPackageResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppPackageResponse) SetBody(v *UpdateAppPackageResponseBody) *UpdateAppPackageResponse {
	s.Body = v
	return s
}

type UpdateFaceConfigRequest struct {
	BizName  *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizType  *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s UpdateFaceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaceConfigRequest) SetBizName(v string) *UpdateFaceConfigRequest {
	s.BizName = &v
	return s
}

func (s *UpdateFaceConfigRequest) SetBizType(v string) *UpdateFaceConfigRequest {
	s.BizType = &v
	return s
}

func (s *UpdateFaceConfigRequest) SetLang(v string) *UpdateFaceConfigRequest {
	s.Lang = &v
	return s
}

func (s *UpdateFaceConfigRequest) SetSourceIp(v string) *UpdateFaceConfigRequest {
	s.SourceIp = &v
	return s
}

type UpdateFaceConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFaceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFaceConfigResponseBody) SetRequestId(v string) *UpdateFaceConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFaceConfigResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFaceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFaceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateFaceConfigResponse) SetHeaders(v map[string]*string) *UpdateFaceConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateFaceConfigResponse) SetBody(v *UpdateFaceConfigResponseBody) *UpdateFaceConfigResponse {
	s.Body = v
	return s
}

type UpdateVerifySettingRequest struct {
	BizName     *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	GuideStep   *bool   `json:"GuideStep,omitempty" xml:"GuideStep,omitempty"`
	PrivacyStep *bool   `json:"PrivacyStep,omitempty" xml:"PrivacyStep,omitempty"`
	ResultStep  *bool   `json:"ResultStep,omitempty" xml:"ResultStep,omitempty"`
	Solution    *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
}

func (s UpdateVerifySettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVerifySettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateVerifySettingRequest) SetBizName(v string) *UpdateVerifySettingRequest {
	s.BizName = &v
	return s
}

func (s *UpdateVerifySettingRequest) SetBizType(v string) *UpdateVerifySettingRequest {
	s.BizType = &v
	return s
}

func (s *UpdateVerifySettingRequest) SetGuideStep(v bool) *UpdateVerifySettingRequest {
	s.GuideStep = &v
	return s
}

func (s *UpdateVerifySettingRequest) SetPrivacyStep(v bool) *UpdateVerifySettingRequest {
	s.PrivacyStep = &v
	return s
}

func (s *UpdateVerifySettingRequest) SetResultStep(v bool) *UpdateVerifySettingRequest {
	s.ResultStep = &v
	return s
}

func (s *UpdateVerifySettingRequest) SetSolution(v string) *UpdateVerifySettingRequest {
	s.Solution = &v
	return s
}

type UpdateVerifySettingResponseBody struct {
	BizName   *string   `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizType   *string   `json:"BizType,omitempty" xml:"BizType,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Solution  *string   `json:"Solution,omitempty" xml:"Solution,omitempty"`
	StepList  []*string `json:"StepList,omitempty" xml:"StepList,omitempty" type:"Repeated"`
}

func (s UpdateVerifySettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVerifySettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVerifySettingResponseBody) SetBizName(v string) *UpdateVerifySettingResponseBody {
	s.BizName = &v
	return s
}

func (s *UpdateVerifySettingResponseBody) SetBizType(v string) *UpdateVerifySettingResponseBody {
	s.BizType = &v
	return s
}

func (s *UpdateVerifySettingResponseBody) SetRequestId(v string) *UpdateVerifySettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVerifySettingResponseBody) SetSolution(v string) *UpdateVerifySettingResponseBody {
	s.Solution = &v
	return s
}

func (s *UpdateVerifySettingResponseBody) SetStepList(v []*string) *UpdateVerifySettingResponseBody {
	s.StepList = v
	return s
}

type UpdateVerifySettingResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVerifySettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVerifySettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVerifySettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateVerifySettingResponse) SetHeaders(v map[string]*string) *UpdateVerifySettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateVerifySettingResponse) SetBody(v *UpdateVerifySettingResponseBody) *UpdateVerifySettingResponse {
	s.Body = v
	return s
}

type VerifyDeviceRequest struct {
	AppVersion  *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	CertifyData *string `json:"CertifyData,omitempty" xml:"CertifyData,omitempty"`
	CertifyId   *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	ExtInfo     *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s VerifyDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyDeviceRequest) GoString() string {
	return s.String()
}

func (s *VerifyDeviceRequest) SetAppVersion(v string) *VerifyDeviceRequest {
	s.AppVersion = &v
	return s
}

func (s *VerifyDeviceRequest) SetCertifyData(v string) *VerifyDeviceRequest {
	s.CertifyData = &v
	return s
}

func (s *VerifyDeviceRequest) SetCertifyId(v string) *VerifyDeviceRequest {
	s.CertifyId = &v
	return s
}

func (s *VerifyDeviceRequest) SetDeviceToken(v string) *VerifyDeviceRequest {
	s.DeviceToken = &v
	return s
}

func (s *VerifyDeviceRequest) SetExtInfo(v string) *VerifyDeviceRequest {
	s.ExtInfo = &v
	return s
}

type VerifyDeviceResponseBody struct {
	Code         *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *VerifyDeviceResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s VerifyDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyDeviceResponseBody) SetCode(v string) *VerifyDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyDeviceResponseBody) SetMessage(v string) *VerifyDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyDeviceResponseBody) SetRequestId(v string) *VerifyDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyDeviceResponseBody) SetResultObject(v *VerifyDeviceResponseBodyResultObject) *VerifyDeviceResponseBody {
	s.ResultObject = v
	return s
}

type VerifyDeviceResponseBodyResultObject struct {
	ExtParams         *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	HasNext           *string `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	ProductRetCode    *string `json:"ProductRetCode,omitempty" xml:"ProductRetCode,omitempty"`
	RetCodeSub        *string `json:"RetCodeSub,omitempty" xml:"RetCodeSub,omitempty"`
	RetMessageSub     *string `json:"RetMessageSub,omitempty" xml:"RetMessageSub,omitempty"`
	ValidationRetCode *string `json:"ValidationRetCode,omitempty" xml:"ValidationRetCode,omitempty"`
}

func (s VerifyDeviceResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s VerifyDeviceResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *VerifyDeviceResponseBodyResultObject) SetExtParams(v string) *VerifyDeviceResponseBodyResultObject {
	s.ExtParams = &v
	return s
}

func (s *VerifyDeviceResponseBodyResultObject) SetHasNext(v string) *VerifyDeviceResponseBodyResultObject {
	s.HasNext = &v
	return s
}

func (s *VerifyDeviceResponseBodyResultObject) SetProductRetCode(v string) *VerifyDeviceResponseBodyResultObject {
	s.ProductRetCode = &v
	return s
}

func (s *VerifyDeviceResponseBodyResultObject) SetRetCodeSub(v string) *VerifyDeviceResponseBodyResultObject {
	s.RetCodeSub = &v
	return s
}

func (s *VerifyDeviceResponseBodyResultObject) SetRetMessageSub(v string) *VerifyDeviceResponseBodyResultObject {
	s.RetMessageSub = &v
	return s
}

func (s *VerifyDeviceResponseBodyResultObject) SetValidationRetCode(v string) *VerifyDeviceResponseBodyResultObject {
	s.ValidationRetCode = &v
	return s
}

type VerifyDeviceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyDeviceResponse) GoString() string {
	return s.String()
}

func (s *VerifyDeviceResponse) SetHeaders(v map[string]*string) *VerifyDeviceResponse {
	s.Headers = v
	return s
}

func (s *VerifyDeviceResponse) SetBody(v *VerifyDeviceResponseBody) *VerifyDeviceResponse {
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
	AuthorityComparisionScore *float32                            `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty"`
	IdCardFaceComparisonScore *float32                            `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	Material                  *VerifyMaterialResponseBodyMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	RequestId                 *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VerifyStatus              *int32                              `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
	VerifyToken               *string                             `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
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
	FaceGlobalUrl *string                                       `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty"`
	FaceImageUrl  *string                                       `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	FaceMask      *string                                       `json:"FaceMask,omitempty" xml:"FaceMask,omitempty"`
	FaceQuality   *string                                       `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	IdCardInfo    *VerifyMaterialResponseBodyMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	IdCardName    *string                                       `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	IdCardNumber  *string                                       `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) CompareFaceVerifyWithOptions(request *CompareFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *CompareFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CompareFaceVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CompareFaceVerify"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CompareFacesWithOptions(request *CompareFacesRequest, runtime *util.RuntimeOptions) (_result *CompareFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CompareFacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CompareFaces"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ContrastFaceVerifyWithOptions(request *ContrastFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *ContrastFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ContrastFaceVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ContrastFaceVerify"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &rpc.Config{
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
	contrastFaceVerifyReq := &ContrastFaceVerifyRequest{}
	openapiutil.Convert(request, contrastFaceVerifyReq)
	if !tea.BoolValue(util.IsUnset(request.FaceContrastFileObject)) {
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
			Content:     request.FaceContrastFileObject,
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
		contrastFaceVerifyReq.FaceContrastFile = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	contrastFaceVerifyResp, _err := client.ContrastFaceVerifyWithOptions(contrastFaceVerifyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = contrastFaceVerifyResp
	return _result, _err
}

func (client *Client) CreateAuthKeyWithOptions(request *CreateAuthKeyRequest, runtime *util.RuntimeOptions) (_result *CreateAuthKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAuthKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAuthKey"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateFaceConfigWithOptions(request *CreateFaceConfigRequest, runtime *util.RuntimeOptions) (_result *CreateFaceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFaceConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFaceConfig"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFaceConfig(request *CreateFaceConfigRequest) (_result *CreateFaceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFaceConfigResponse{}
	_body, _err := client.CreateFaceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRPSDKWithOptions(request *CreateRPSDKRequest, runtime *util.RuntimeOptions) (_result *CreateRPSDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRPSDKResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRPSDK"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRPSDK(request *CreateRPSDKRequest) (_result *CreateRPSDKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRPSDKResponse{}
	_body, _err := client.CreateRPSDKWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVerifySDKWithOptions(request *CreateVerifySDKRequest, runtime *util.RuntimeOptions) (_result *CreateVerifySDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVerifySDKResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVerifySDK"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVerifySDK(request *CreateVerifySDKRequest) (_result *CreateVerifySDKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVerifySDKResponse{}
	_body, _err := client.CreateVerifySDKWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVerifySettingWithOptions(request *CreateVerifySettingRequest, runtime *util.RuntimeOptions) (_result *CreateVerifySettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVerifySettingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVerifySetting"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateWhitelistWithOptions(request *CreateWhitelistRequest, runtime *util.RuntimeOptions) (_result *CreateWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateWhitelist"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWhitelist(request *CreateWhitelistRequest) (_result *CreateWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWhitelistResponse{}
	_body, _err := client.CreateWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWhitelistSettingWithOptions(request *CreateWhitelistSettingRequest, runtime *util.RuntimeOptions) (_result *CreateWhitelistSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateWhitelistSettingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateWhitelistSetting"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWhitelistSetting(request *CreateWhitelistSettingRequest) (_result *CreateWhitelistSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWhitelistSettingResponse{}
	_body, _err := client.CreateWhitelistSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWhitelistWithOptions(request *DeleteWhitelistRequest, runtime *util.RuntimeOptions) (_result *DeleteWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteWhitelist"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWhitelist(request *DeleteWhitelistRequest) (_result *DeleteWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWhitelistResponse{}
	_body, _err := client.DeleteWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWhitelistSettingWithOptions(request *DeleteWhitelistSettingRequest, runtime *util.RuntimeOptions) (_result *DeleteWhitelistSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteWhitelistSettingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteWhitelistSetting"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWhitelistSetting(request *DeleteWhitelistSettingRequest) (_result *DeleteWhitelistSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWhitelistSettingResponse{}
	_body, _err := client.DeleteWhitelistSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppInfoWithOptions(request *DescribeAppInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeAppInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAppInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAppInfo"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppInfo(request *DescribeAppInfoRequest) (_result *DescribeAppInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppInfoResponse{}
	_body, _err := client.DescribeAppInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceInfoWithOptions(request *DescribeDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDeviceInfo"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeFaceConfigWithOptions(request *DescribeFaceConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeFaceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFaceConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFaceConfig"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFaceConfig(request *DescribeFaceConfigRequest) (_result *DescribeFaceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaceConfigResponse{}
	_body, _err := client.DescribeFaceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFaceUsageWithOptions(request *DescribeFaceUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeFaceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFaceUsageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFaceUsage"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFaceUsage(request *DescribeFaceUsageRequest) (_result *DescribeFaceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaceUsageResponse{}
	_body, _err := client.DescribeFaceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFaceVerifyWithOptions(request *DescribeFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *DescribeFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFaceVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFaceVerify"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeOssUploadTokenWithOptions(runtime *util.RuntimeOptions) (_result *DescribeOssUploadTokenResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeOssUploadTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOssUploadToken"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeRPSDKWithOptions(request *DescribeRPSDKRequest, runtime *util.RuntimeOptions) (_result *DescribeRPSDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRPSDKResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRPSDK"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRPSDK(request *DescribeRPSDKRequest) (_result *DescribeRPSDKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRPSDKResponse{}
	_body, _err := client.DescribeRPSDKWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSdkUrlWithOptions(request *DescribeSdkUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeSdkUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSdkUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSdkUrl"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSdkUrl(request *DescribeSdkUrlRequest) (_result *DescribeSdkUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSdkUrlResponse{}
	_body, _err := client.DescribeSdkUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUpdatePackageResultWithOptions(request *DescribeUpdatePackageResultRequest, runtime *util.RuntimeOptions) (_result *DescribeUpdatePackageResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUpdatePackageResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUpdatePackageResult"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUpdatePackageResult(request *DescribeUpdatePackageResultRequest) (_result *DescribeUpdatePackageResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUpdatePackageResultResponse{}
	_body, _err := client.DescribeUpdatePackageResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUploadInfoWithOptions(request *DescribeUploadInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeUploadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUploadInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUploadInfo"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUploadInfo(request *DescribeUploadInfoRequest) (_result *DescribeUploadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUploadInfoResponse{}
	_body, _err := client.DescribeUploadInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserStatusWithOptions(runtime *util.RuntimeOptions) (_result *DescribeUserStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeUserStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserStatus"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserStatus() (_result *DescribeUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserStatusResponse{}
	_body, _err := client.DescribeUserStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifyRecordsWithOptions(request *DescribeVerifyRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifyRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVerifyRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVerifyRecords"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifyRecords(request *DescribeVerifyRecordsRequest) (_result *DescribeVerifyRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifyRecordsResponse{}
	_body, _err := client.DescribeVerifyRecordsWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVerifyResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVerifyResult"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeVerifySDKWithOptions(request *DescribeVerifySDKRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifySDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVerifySDKResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVerifySDK"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeVerifySettingWithOptions(runtime *util.RuntimeOptions) (_result *DescribeVerifySettingResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeVerifySettingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVerifySetting"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifySetting() (_result *DescribeVerifySettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifySettingResponse{}
	_body, _err := client.DescribeVerifySettingWithOptions(runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVerifyTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVerifyToken"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeVerifyUsageWithOptions(request *DescribeVerifyUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifyUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVerifyUsageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVerifyUsage"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifyUsage(request *DescribeVerifyUsageRequest) (_result *DescribeVerifyUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifyUsageResponse{}
	_body, _err := client.DescribeVerifyUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWhitelistWithOptions(request *DescribeWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWhitelist"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWhitelist(request *DescribeWhitelistRequest) (_result *DescribeWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWhitelistResponse{}
	_body, _err := client.DescribeWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWhitelistSettingWithOptions(request *DescribeWhitelistSettingRequest, runtime *util.RuntimeOptions) (_result *DescribeWhitelistSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWhitelistSettingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWhitelistSetting"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWhitelistSetting(request *DescribeWhitelistSettingRequest) (_result *DescribeWhitelistSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWhitelistSettingResponse{}
	_body, _err := client.DescribeWhitelistSettingWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectFaceAttributesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectFaceAttributes"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) InitDeviceWithOptions(request *InitDeviceRequest, runtime *util.RuntimeOptions) (_result *InitDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InitDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InitDevice"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitDevice(request *InitDeviceRequest) (_result *InitDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitDeviceResponse{}
	_body, _err := client.InitDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitFaceVerifyWithOptions(request *InitFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *InitFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InitFaceVerify"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) LivenessFaceVerifyWithOptions(request *LivenessFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *LivenessFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LivenessFaceVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LivenessFaceVerify"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifyDeviceInfoWithOptions(request *ModifyDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *ModifyDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDeviceInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDeviceInfo"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateAppPackageWithOptions(request *UpdateAppPackageRequest, runtime *util.RuntimeOptions) (_result *UpdateAppPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppPackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAppPackage"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppPackage(request *UpdateAppPackageRequest) (_result *UpdateAppPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppPackageResponse{}
	_body, _err := client.UpdateAppPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFaceConfigWithOptions(request *UpdateFaceConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateFaceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFaceConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFaceConfig"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFaceConfig(request *UpdateFaceConfigRequest) (_result *UpdateFaceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFaceConfigResponse{}
	_body, _err := client.UpdateFaceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVerifySettingWithOptions(request *UpdateVerifySettingRequest, runtime *util.RuntimeOptions) (_result *UpdateVerifySettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateVerifySettingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateVerifySetting"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVerifySetting(request *UpdateVerifySettingRequest) (_result *UpdateVerifySettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVerifySettingResponse{}
	_body, _err := client.UpdateVerifySettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyDeviceWithOptions(request *VerifyDeviceRequest, runtime *util.RuntimeOptions) (_result *VerifyDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyDevice"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyDevice(request *VerifyDeviceRequest) (_result *VerifyDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyDeviceResponse{}
	_body, _err := client.VerifyDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyMaterialWithOptions(request *VerifyMaterialRequest, runtime *util.RuntimeOptions) (_result *VerifyMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyMaterialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyMaterial"), tea.String("2019-03-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
