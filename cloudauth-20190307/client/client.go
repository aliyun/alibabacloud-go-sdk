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

type DescribeWhitelistRequest struct {
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizType        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	IdCardNum      *string `json:"IdCardNum,omitempty" xml:"IdCardNum,omitempty"`
	ValidStartDate *string `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
	ValidEndDate   *string `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	Valid          *string `json:"Valid,omitempty" xml:"Valid,omitempty"`
	PageSize       *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CurrentPage    *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
}

func (s DescribeWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistRequest) SetSourceIp(v string) *DescribeWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWhitelistRequest) SetLang(v string) *DescribeWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWhitelistRequest) SetBizType(v string) *DescribeWhitelistRequest {
	s.BizType = &v
	return s
}

func (s *DescribeWhitelistRequest) SetBizId(v string) *DescribeWhitelistRequest {
	s.BizId = &v
	return s
}

func (s *DescribeWhitelistRequest) SetIdCardNum(v string) *DescribeWhitelistRequest {
	s.IdCardNum = &v
	return s
}

func (s *DescribeWhitelistRequest) SetValidStartDate(v string) *DescribeWhitelistRequest {
	s.ValidStartDate = &v
	return s
}

func (s *DescribeWhitelistRequest) SetValidEndDate(v string) *DescribeWhitelistRequest {
	s.ValidEndDate = &v
	return s
}

func (s *DescribeWhitelistRequest) SetValid(v string) *DescribeWhitelistRequest {
	s.Valid = &v
	return s
}

func (s *DescribeWhitelistRequest) SetPageSize(v int) *DescribeWhitelistRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWhitelistRequest) SetCurrentPage(v int) *DescribeWhitelistRequest {
	s.CurrentPage = &v
	return s
}

type DescribeWhitelistResponse struct {
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount  *int                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	CurrentPage *int                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize    *int                              `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Items       []*DescribeWhitelistResponseItems `json:"Items,omitempty" xml:"Items,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistResponse) SetRequestId(v string) *DescribeWhitelistResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeWhitelistResponse) SetTotalCount(v int) *DescribeWhitelistResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeWhitelistResponse) SetCurrentPage(v int) *DescribeWhitelistResponse {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhitelistResponse) SetPageSize(v int) *DescribeWhitelistResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeWhitelistResponse) SetItems(v []*DescribeWhitelistResponseItems) *DescribeWhitelistResponse {
	s.Items = v
	return s
}

type DescribeWhitelistResponseItems struct {
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	Uid         *int64  `json:"Uid,omitempty" xml:"Uid,omitempty" require:"true"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	StartDate   *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate     *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
	IdCardNum   *string `json:"IdCardNum,omitempty" xml:"IdCardNum,omitempty" require:"true"`
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	Valid       *int    `json:"Valid,omitempty" xml:"Valid,omitempty" require:"true"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty" require:"true"`
}

func (s DescribeWhitelistResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhitelistResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistResponseItems) SetId(v int64) *DescribeWhitelistResponseItems {
	s.Id = &v
	return s
}

func (s *DescribeWhitelistResponseItems) SetUid(v int64) *DescribeWhitelistResponseItems {
	s.Uid = &v
	return s
}

func (s *DescribeWhitelistResponseItems) SetBizType(v string) *DescribeWhitelistResponseItems {
	s.BizType = &v
	return s
}

func (s *DescribeWhitelistResponseItems) SetStartDate(v int64) *DescribeWhitelistResponseItems {
	s.StartDate = &v
	return s
}

func (s *DescribeWhitelistResponseItems) SetEndDate(v int64) *DescribeWhitelistResponseItems {
	s.EndDate = &v
	return s
}

func (s *DescribeWhitelistResponseItems) SetIdCardNum(v string) *DescribeWhitelistResponseItems {
	s.IdCardNum = &v
	return s
}

func (s *DescribeWhitelistResponseItems) SetBizId(v string) *DescribeWhitelistResponseItems {
	s.BizId = &v
	return s
}

func (s *DescribeWhitelistResponseItems) SetValid(v int) *DescribeWhitelistResponseItems {
	s.Valid = &v
	return s
}

func (s *DescribeWhitelistResponseItems) SetGmtCreate(v int64) *DescribeWhitelistResponseItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeWhitelistResponseItems) SetGmtModified(v int64) *DescribeWhitelistResponseItems {
	s.GmtModified = &v
	return s
}

type DeleteWhitelistRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Ids      *string `json:"Ids,omitempty" xml:"Ids,omitempty" require:"true"`
}

func (s DeleteWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistRequest) SetSourceIp(v string) *DeleteWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteWhitelistRequest) SetLang(v string) *DeleteWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *DeleteWhitelistRequest) SetIds(v string) *DeleteWhitelistRequest {
	s.Ids = &v
	return s
}

type DeleteWhitelistResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistResponse) SetRequestId(v string) *DeleteWhitelistResponse {
	s.RequestId = &v
	return s
}

type CreateWhitelistRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	IdCardNum *string `json:"IdCardNum,omitempty" xml:"IdCardNum,omitempty" require:"true"`
	ValidDay  *string `json:"ValidDay,omitempty" xml:"ValidDay,omitempty" require:"true"`
}

func (s CreateWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWhitelistRequest) GoString() string {
	return s.String()
}

func (s *CreateWhitelistRequest) SetSourceIp(v string) *CreateWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateWhitelistRequest) SetLang(v string) *CreateWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *CreateWhitelistRequest) SetBizType(v string) *CreateWhitelistRequest {
	s.BizType = &v
	return s
}

func (s *CreateWhitelistRequest) SetBizId(v string) *CreateWhitelistRequest {
	s.BizId = &v
	return s
}

func (s *CreateWhitelistRequest) SetIdCardNum(v string) *CreateWhitelistRequest {
	s.IdCardNum = &v
	return s
}

func (s *CreateWhitelistRequest) SetValidDay(v string) *CreateWhitelistRequest {
	s.ValidDay = &v
	return s
}

type CreateWhitelistResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWhitelistResponse) GoString() string {
	return s.String()
}

func (s *CreateWhitelistResponse) SetRequestId(v string) *CreateWhitelistResponse {
	s.RequestId = &v
	return s
}

type DescribeFaceConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeFaceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaceConfigRequest) SetSourceIp(v string) *DescribeFaceConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFaceConfigRequest) SetLang(v string) *DescribeFaceConfigRequest {
	s.Lang = &v
	return s
}

type DescribeFaceConfigResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Items     []*DescribeFaceConfigResponseItems `json:"Items,omitempty" xml:"Items,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeFaceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaceConfigResponse) SetRequestId(v string) *DescribeFaceConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeFaceConfigResponse) SetItems(v []*DescribeFaceConfigResponseItems) *DescribeFaceConfigResponse {
	s.Items = v
	return s
}

type DescribeFaceConfigResponseItems struct {
	BizType    *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	BizName    *string `json:"BizName,omitempty" xml:"BizName,omitempty" require:"true"`
	GmtUpdated *int64  `json:"GmtUpdated,omitempty" xml:"GmtUpdated,omitempty" require:"true"`
}

func (s DescribeFaceConfigResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceConfigResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeFaceConfigResponseItems) SetBizType(v string) *DescribeFaceConfigResponseItems {
	s.BizType = &v
	return s
}

func (s *DescribeFaceConfigResponseItems) SetBizName(v string) *DescribeFaceConfigResponseItems {
	s.BizName = &v
	return s
}

func (s *DescribeFaceConfigResponseItems) SetGmtUpdated(v int64) *DescribeFaceConfigResponseItems {
	s.GmtUpdated = &v
	return s
}

type UpdateFaceConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizType  *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	BizName  *string `json:"BizName,omitempty" xml:"BizName,omitempty" require:"true"`
}

func (s UpdateFaceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaceConfigRequest) SetSourceIp(v string) *UpdateFaceConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateFaceConfigRequest) SetLang(v string) *UpdateFaceConfigRequest {
	s.Lang = &v
	return s
}

func (s *UpdateFaceConfigRequest) SetBizType(v string) *UpdateFaceConfigRequest {
	s.BizType = &v
	return s
}

func (s *UpdateFaceConfigRequest) SetBizName(v string) *UpdateFaceConfigRequest {
	s.BizName = &v
	return s
}

type UpdateFaceConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateFaceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateFaceConfigResponse) SetRequestId(v string) *UpdateFaceConfigResponse {
	s.RequestId = &v
	return s
}

type CreateFaceConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizType  *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	BizName  *string `json:"BizName,omitempty" xml:"BizName,omitempty" require:"true"`
}

func (s CreateFaceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateFaceConfigRequest) SetSourceIp(v string) *CreateFaceConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateFaceConfigRequest) SetLang(v string) *CreateFaceConfigRequest {
	s.Lang = &v
	return s
}

func (s *CreateFaceConfigRequest) SetBizType(v string) *CreateFaceConfigRequest {
	s.BizType = &v
	return s
}

func (s *CreateFaceConfigRequest) SetBizName(v string) *CreateFaceConfigRequest {
	s.BizName = &v
	return s
}

type CreateFaceConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateFaceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateFaceConfigResponse) SetRequestId(v string) *CreateFaceConfigResponse {
	s.RequestId = &v
	return s
}

type LivenessFaceVerifyRequest struct {
	SceneId                *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	OuterOrderNo           *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	FaceContrastPicture    *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	DeviceToken            *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	Mobile                 *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Ip                     *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	CertifyId              *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	OssBucketName          *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	Model                  *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s LivenessFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyRequest) SetSceneId(v int64) *LivenessFaceVerifyRequest {
	s.SceneId = &v
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

func (s *LivenessFaceVerifyRequest) SetFaceContrastPicture(v string) *LivenessFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetDeviceToken(v string) *LivenessFaceVerifyRequest {
	s.DeviceToken = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetMobile(v string) *LivenessFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetIp(v string) *LivenessFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetUserId(v string) *LivenessFaceVerifyRequest {
	s.UserId = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *LivenessFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetCertifyId(v string) *LivenessFaceVerifyRequest {
	s.CertifyId = &v
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

func (s *LivenessFaceVerifyRequest) SetModel(v string) *LivenessFaceVerifyRequest {
	s.Model = &v
	return s
}

type LivenessFaceVerifyResponse struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResultObject *LivenessFaceVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s LivenessFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyResponse) SetRequestId(v string) *LivenessFaceVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *LivenessFaceVerifyResponse) SetMessage(v string) *LivenessFaceVerifyResponse {
	s.Message = &v
	return s
}

func (s *LivenessFaceVerifyResponse) SetCode(v string) *LivenessFaceVerifyResponse {
	s.Code = &v
	return s
}

func (s *LivenessFaceVerifyResponse) SetResultObject(v *LivenessFaceVerifyResponseResultObject) *LivenessFaceVerifyResponse {
	s.ResultObject = v
	return s
}

type LivenessFaceVerifyResponseResultObject struct {
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty" require:"true"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty" require:"true"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty" require:"true"`
}

func (s LivenessFaceVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyResponseResultObject) SetPassed(v string) *LivenessFaceVerifyResponseResultObject {
	s.Passed = &v
	return s
}

func (s *LivenessFaceVerifyResponseResultObject) SetMaterialInfo(v string) *LivenessFaceVerifyResponseResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *LivenessFaceVerifyResponseResultObject) SetSubCode(v string) *LivenessFaceVerifyResponseResultObject {
	s.SubCode = &v
	return s
}

type CompareFaceVerifyRequest struct {
	SceneId                      *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	OuterOrderNo                 *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode                  *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SourceFaceContrastPicture    *string `json:"SourceFaceContrastPicture,omitempty" xml:"SourceFaceContrastPicture,omitempty"`
	SourceFaceContrastPictureUrl *string `json:"SourceFaceContrastPictureUrl,omitempty" xml:"SourceFaceContrastPictureUrl,omitempty"`
	SourceCertifyId              *string `json:"SourceCertifyId,omitempty" xml:"SourceCertifyId,omitempty"`
	SourceOssBucketName          *string `json:"SourceOssBucketName,omitempty" xml:"SourceOssBucketName,omitempty"`
	SourceOssObjectName          *string `json:"SourceOssObjectName,omitempty" xml:"SourceOssObjectName,omitempty"`
	TargetFaceContrastPicture    *string `json:"TargetFaceContrastPicture,omitempty" xml:"TargetFaceContrastPicture,omitempty"`
	TargetFaceContrastPictureUrl *string `json:"TargetFaceContrastPictureUrl,omitempty" xml:"TargetFaceContrastPictureUrl,omitempty"`
	TargetCertifyId              *string `json:"TargetCertifyId,omitempty" xml:"TargetCertifyId,omitempty"`
	TargetOssBucketName          *string `json:"TargetOssBucketName,omitempty" xml:"TargetOssBucketName,omitempty"`
	TargetOssObjectName          *string `json:"TargetOssObjectName,omitempty" xml:"TargetOssObjectName,omitempty"`
}

func (s CompareFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyRequest) SetSceneId(v int64) *CompareFaceVerifyRequest {
	s.SceneId = &v
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

func (s *CompareFaceVerifyRequest) SetSourceFaceContrastPicture(v string) *CompareFaceVerifyRequest {
	s.SourceFaceContrastPicture = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceFaceContrastPictureUrl(v string) *CompareFaceVerifyRequest {
	s.SourceFaceContrastPictureUrl = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceCertifyId(v string) *CompareFaceVerifyRequest {
	s.SourceCertifyId = &v
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

func (s *CompareFaceVerifyRequest) SetTargetFaceContrastPicture(v string) *CompareFaceVerifyRequest {
	s.TargetFaceContrastPicture = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetFaceContrastPictureUrl(v string) *CompareFaceVerifyRequest {
	s.TargetFaceContrastPictureUrl = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetCertifyId(v string) *CompareFaceVerifyRequest {
	s.TargetCertifyId = &v
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

type CompareFaceVerifyResponse struct {
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message      *string                                `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResultObject *CompareFaceVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s CompareFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyResponse) SetRequestId(v string) *CompareFaceVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *CompareFaceVerifyResponse) SetMessage(v string) *CompareFaceVerifyResponse {
	s.Message = &v
	return s
}

func (s *CompareFaceVerifyResponse) SetCode(v string) *CompareFaceVerifyResponse {
	s.Code = &v
	return s
}

func (s *CompareFaceVerifyResponse) SetResultObject(v *CompareFaceVerifyResponseResultObject) *CompareFaceVerifyResponse {
	s.ResultObject = v
	return s
}

type CompareFaceVerifyResponseResultObject struct {
	Passed      *string  `json:"Passed,omitempty" xml:"Passed,omitempty" require:"true"`
	VerifyScore *float32 `json:"VerifyScore,omitempty" xml:"VerifyScore,omitempty" require:"true"`
}

func (s CompareFaceVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyResponseResultObject) SetPassed(v string) *CompareFaceVerifyResponseResultObject {
	s.Passed = &v
	return s
}

func (s *CompareFaceVerifyResponseResultObject) SetVerifyScore(v float32) *CompareFaceVerifyResponseResultObject {
	s.VerifyScore = &v
	return s
}

type DescribeSdkUrlRequest struct {
	Id    *int64 `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	Debug *bool  `json:"Debug,omitempty" xml:"Debug,omitempty"`
}

func (s DescribeSdkUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSdkUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeSdkUrlRequest) SetId(v int64) *DescribeSdkUrlRequest {
	s.Id = &v
	return s
}

func (s *DescribeSdkUrlRequest) SetDebug(v bool) *DescribeSdkUrlRequest {
	s.Debug = &v
	return s
}

type DescribeSdkUrlResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SdkUrl    *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty" require:"true"`
}

func (s DescribeSdkUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSdkUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeSdkUrlResponse) SetRequestId(v string) *DescribeSdkUrlResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSdkUrlResponse) SetSdkUrl(v string) *DescribeSdkUrlResponse {
	s.SdkUrl = &v
	return s
}

type DescribeUpdatePackageResultRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
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

type DescribeUpdatePackageResultResponse struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AppInfo   *DescribeUpdatePackageResultResponseAppInfo `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribeUpdatePackageResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponse) SetRequestId(v string) *DescribeUpdatePackageResultResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUpdatePackageResultResponse) SetAppInfo(v *DescribeUpdatePackageResultResponseAppInfo) *DescribeUpdatePackageResultResponse {
	s.AppInfo = v
	return s
}

type DescribeUpdatePackageResultResponseAppInfo struct {
	Id               *int64                                                      `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	Name             *string                                                     `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	PackageName      *string                                                     `json:"PackageName,omitempty" xml:"PackageName,omitempty" require:"true"`
	Icon             *string                                                     `json:"Icon,omitempty" xml:"Icon,omitempty" require:"true"`
	StartDate        *string                                                     `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate          *string                                                     `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
	Type             *int                                                        `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	PackageInfo      *DescribeUpdatePackageResultResponseAppInfoPackageInfo      `json:"PackageInfo,omitempty" xml:"PackageInfo,omitempty" require:"true" type:"Struct"`
	DebugPackageInfo *DescribeUpdatePackageResultResponseAppInfoDebugPackageInfo `json:"DebugPackageInfo,omitempty" xml:"DebugPackageInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribeUpdatePackageResultResponseAppInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponseAppInfo) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponseAppInfo) SetId(v int64) *DescribeUpdatePackageResultResponseAppInfo {
	s.Id = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseAppInfo) SetName(v string) *DescribeUpdatePackageResultResponseAppInfo {
	s.Name = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseAppInfo) SetPackageName(v string) *DescribeUpdatePackageResultResponseAppInfo {
	s.PackageName = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseAppInfo) SetIcon(v string) *DescribeUpdatePackageResultResponseAppInfo {
	s.Icon = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseAppInfo) SetStartDate(v string) *DescribeUpdatePackageResultResponseAppInfo {
	s.StartDate = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseAppInfo) SetEndDate(v string) *DescribeUpdatePackageResultResponseAppInfo {
	s.EndDate = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseAppInfo) SetType(v int) *DescribeUpdatePackageResultResponseAppInfo {
	s.Type = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseAppInfo) SetPackageInfo(v *DescribeUpdatePackageResultResponseAppInfoPackageInfo) *DescribeUpdatePackageResultResponseAppInfo {
	s.PackageInfo = v
	return s
}

func (s *DescribeUpdatePackageResultResponseAppInfo) SetDebugPackageInfo(v *DescribeUpdatePackageResultResponseAppInfoDebugPackageInfo) *DescribeUpdatePackageResultResponseAppInfo {
	s.DebugPackageInfo = v
	return s
}

type DescribeUpdatePackageResultResponseAppInfoPackageInfo struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s DescribeUpdatePackageResultResponseAppInfoPackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponseAppInfoPackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponseAppInfoPackageInfo) SetVersion(v string) *DescribeUpdatePackageResultResponseAppInfoPackageInfo {
	s.Version = &v
	return s
}

type DescribeUpdatePackageResultResponseAppInfoDebugPackageInfo struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s DescribeUpdatePackageResultResponseAppInfoDebugPackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponseAppInfoDebugPackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponseAppInfoDebugPackageInfo) SetVersion(v string) *DescribeUpdatePackageResultResponseAppInfoDebugPackageInfo {
	s.Version = &v
	return s
}

type UpdateAppPackageRequest struct {
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty" require:"true"`
	Platform   *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Debug      *bool   `json:"Debug,omitempty" xml:"Debug,omitempty"`
}

func (s UpdateAppPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppPackageRequest) GoString() string {
	return s.String()
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

func (s *UpdateAppPackageRequest) SetDebug(v bool) *UpdateAppPackageRequest {
	s.Debug = &v
	return s
}

type UpdateAppPackageResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s UpdateAppPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppPackageResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppPackageResponse) SetRequestId(v string) *UpdateAppPackageResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateAppPackageResponse) SetTaskId(v string) *UpdateAppPackageResponse {
	s.TaskId = &v
	return s
}

type DescribeAppInfoRequest struct {
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Platform    *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s DescribeAppInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoRequest) SetPageSize(v int) *DescribeAppInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppInfoRequest) SetCurrentPage(v int) *DescribeAppInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAppInfoRequest) SetPlatform(v string) *DescribeAppInfoRequest {
	s.Platform = &v
	return s
}

type DescribeAppInfoResponse struct {
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageSize    *int                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CurrentPage *int                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	TotalCount  *int                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	AppInfoList []*DescribeAppInfoResponseAppInfoList `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAppInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponse) SetRequestId(v string) *DescribeAppInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeAppInfoResponse) SetPageSize(v int) *DescribeAppInfoResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeAppInfoResponse) SetCurrentPage(v int) *DescribeAppInfoResponse {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAppInfoResponse) SetTotalCount(v int) *DescribeAppInfoResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeAppInfoResponse) SetAppInfoList(v []*DescribeAppInfoResponseAppInfoList) *DescribeAppInfoResponse {
	s.AppInfoList = v
	return s
}

type DescribeAppInfoResponseAppInfoList struct {
	Id               *int64                                              `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	Name             *string                                             `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	PackageName      *string                                             `json:"PackageName,omitempty" xml:"PackageName,omitempty" require:"true"`
	Icon             *string                                             `json:"Icon,omitempty" xml:"Icon,omitempty" require:"true"`
	StartDate        *string                                             `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate          *string                                             `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
	Type             *int                                                `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	PackageInfo      *DescribeAppInfoResponseAppInfoListPackageInfo      `json:"PackageInfo,omitempty" xml:"PackageInfo,omitempty" require:"true" type:"Struct"`
	DebugPackageInfo *DescribeAppInfoResponseAppInfoListDebugPackageInfo `json:"DebugPackageInfo,omitempty" xml:"DebugPackageInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribeAppInfoResponseAppInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponseAppInfoList) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponseAppInfoList) SetId(v int64) *DescribeAppInfoResponseAppInfoList {
	s.Id = &v
	return s
}

func (s *DescribeAppInfoResponseAppInfoList) SetName(v string) *DescribeAppInfoResponseAppInfoList {
	s.Name = &v
	return s
}

func (s *DescribeAppInfoResponseAppInfoList) SetPackageName(v string) *DescribeAppInfoResponseAppInfoList {
	s.PackageName = &v
	return s
}

func (s *DescribeAppInfoResponseAppInfoList) SetIcon(v string) *DescribeAppInfoResponseAppInfoList {
	s.Icon = &v
	return s
}

func (s *DescribeAppInfoResponseAppInfoList) SetStartDate(v string) *DescribeAppInfoResponseAppInfoList {
	s.StartDate = &v
	return s
}

func (s *DescribeAppInfoResponseAppInfoList) SetEndDate(v string) *DescribeAppInfoResponseAppInfoList {
	s.EndDate = &v
	return s
}

func (s *DescribeAppInfoResponseAppInfoList) SetType(v int) *DescribeAppInfoResponseAppInfoList {
	s.Type = &v
	return s
}

func (s *DescribeAppInfoResponseAppInfoList) SetPackageInfo(v *DescribeAppInfoResponseAppInfoListPackageInfo) *DescribeAppInfoResponseAppInfoList {
	s.PackageInfo = v
	return s
}

func (s *DescribeAppInfoResponseAppInfoList) SetDebugPackageInfo(v *DescribeAppInfoResponseAppInfoListDebugPackageInfo) *DescribeAppInfoResponseAppInfoList {
	s.DebugPackageInfo = v
	return s
}

type DescribeAppInfoResponseAppInfoListPackageInfo struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s DescribeAppInfoResponseAppInfoListPackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponseAppInfoListPackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponseAppInfoListPackageInfo) SetVersion(v string) *DescribeAppInfoResponseAppInfoListPackageInfo {
	s.Version = &v
	return s
}

type DescribeAppInfoResponseAppInfoListDebugPackageInfo struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s DescribeAppInfoResponseAppInfoListDebugPackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponseAppInfoListDebugPackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponseAppInfoListDebugPackageInfo) SetVersion(v string) *DescribeAppInfoResponseAppInfoListDebugPackageInfo {
	s.Version = &v
	return s
}

type ContrastFaceVerifyRequest struct {
	SceneId                *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	OuterOrderNo           *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	CertType               *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertName               *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                 *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	FaceContrastPicture    *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	DeviceToken            *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	Mobile                 *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Ip                     *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	CertifyId              *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	OssBucketName          *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	Model                  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	FaceContrastFile       *string `json:"FaceContrastFile,omitempty" xml:"FaceContrastFile,omitempty"`
}

func (s ContrastFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyRequest) SetSceneId(v int64) *ContrastFaceVerifyRequest {
	s.SceneId = &v
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

func (s *ContrastFaceVerifyRequest) SetCertType(v string) *ContrastFaceVerifyRequest {
	s.CertType = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertName(v string) *ContrastFaceVerifyRequest {
	s.CertName = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertNo(v string) *ContrastFaceVerifyRequest {
	s.CertNo = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastPicture(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetDeviceToken(v string) *ContrastFaceVerifyRequest {
	s.DeviceToken = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetMobile(v string) *ContrastFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetIp(v string) *ContrastFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetUserId(v string) *ContrastFaceVerifyRequest {
	s.UserId = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertifyId(v string) *ContrastFaceVerifyRequest {
	s.CertifyId = &v
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

func (s *ContrastFaceVerifyRequest) SetModel(v string) *ContrastFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastFile(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastFile = &v
	return s
}

type ContrastFaceVerifyResponse struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResultObject *ContrastFaceVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s ContrastFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyResponse) SetRequestId(v string) *ContrastFaceVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *ContrastFaceVerifyResponse) SetMessage(v string) *ContrastFaceVerifyResponse {
	s.Message = &v
	return s
}

func (s *ContrastFaceVerifyResponse) SetCode(v string) *ContrastFaceVerifyResponse {
	s.Code = &v
	return s
}

func (s *ContrastFaceVerifyResponse) SetResultObject(v *ContrastFaceVerifyResponseResultObject) *ContrastFaceVerifyResponse {
	s.ResultObject = v
	return s
}

type ContrastFaceVerifyResponseResultObject struct {
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty" require:"true"`
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty" require:"true"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty" require:"true"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty" require:"true"`
}

func (s ContrastFaceVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyResponseResultObject) SetPassed(v string) *ContrastFaceVerifyResponseResultObject {
	s.Passed = &v
	return s
}

func (s *ContrastFaceVerifyResponseResultObject) SetIdentityInfo(v string) *ContrastFaceVerifyResponseResultObject {
	s.IdentityInfo = &v
	return s
}

func (s *ContrastFaceVerifyResponseResultObject) SetMaterialInfo(v string) *ContrastFaceVerifyResponseResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *ContrastFaceVerifyResponseResultObject) SetSubCode(v string) *ContrastFaceVerifyResponseResultObject {
	s.SubCode = &v
	return s
}

type ContrastFaceVerifyAdvanceRequest struct {
	FaceContrastFileObject io.Reader `json:"FaceContrastFileObject,omitempty" xml:"FaceContrastFileObject,omitempty" require:"true"`
	SceneId                *int64    `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	OuterOrderNo           *string   `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string   `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	CertType               *string   `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertName               *string   `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                 *string   `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	FaceContrastPicture    *string   `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	DeviceToken            *string   `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Ip                     *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	FaceContrastPictureUrl *string   `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	CertifyId              *string   `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	OssBucketName          *string   `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string   `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	Model                  *string   `json:"Model,omitempty" xml:"Model,omitempty"`
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

func (s *ContrastFaceVerifyAdvanceRequest) SetSceneId(v int64) *ContrastFaceVerifyAdvanceRequest {
	s.SceneId = &v
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

func (s *ContrastFaceVerifyAdvanceRequest) SetCertType(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertType = &v
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

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastPicture(v string) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetDeviceToken(v string) *ContrastFaceVerifyAdvanceRequest {
	s.DeviceToken = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetMobile(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Mobile = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetIp(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Ip = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetUserId(v string) *ContrastFaceVerifyAdvanceRequest {
	s.UserId = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastPictureUrl(v string) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertifyId(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertifyId = &v
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

func (s *ContrastFaceVerifyAdvanceRequest) SetModel(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Model = &v
	return s
}

type InitDeviceRequest struct {
	CertifyId        *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	OuterOrderNo     *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	Channel          *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Merchant         *string `json:"Merchant,omitempty" xml:"Merchant,omitempty"`
	ProductName      *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProduceNode      *string `json:"ProduceNode,omitempty" xml:"ProduceNode,omitempty"`
	BizData          *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
	MetaInfo         *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	CertifyPrincipal *string `json:"CertifyPrincipal,omitempty" xml:"CertifyPrincipal,omitempty"`
	AppVersion       *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	DeviceToken      *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	UaToken          *string `json:"UaToken,omitempty" xml:"UaToken,omitempty"`
	WebUmidToken     *string `json:"WebUmidToken,omitempty" xml:"WebUmidToken,omitempty"`
}

func (s InitDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s InitDeviceRequest) GoString() string {
	return s.String()
}

func (s *InitDeviceRequest) SetCertifyId(v string) *InitDeviceRequest {
	s.CertifyId = &v
	return s
}

func (s *InitDeviceRequest) SetOuterOrderNo(v string) *InitDeviceRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *InitDeviceRequest) SetChannel(v string) *InitDeviceRequest {
	s.Channel = &v
	return s
}

func (s *InitDeviceRequest) SetMerchant(v string) *InitDeviceRequest {
	s.Merchant = &v
	return s
}

func (s *InitDeviceRequest) SetProductName(v string) *InitDeviceRequest {
	s.ProductName = &v
	return s
}

func (s *InitDeviceRequest) SetProduceNode(v string) *InitDeviceRequest {
	s.ProduceNode = &v
	return s
}

func (s *InitDeviceRequest) SetBizData(v string) *InitDeviceRequest {
	s.BizData = &v
	return s
}

func (s *InitDeviceRequest) SetMetaInfo(v string) *InitDeviceRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitDeviceRequest) SetCertifyPrincipal(v string) *InitDeviceRequest {
	s.CertifyPrincipal = &v
	return s
}

func (s *InitDeviceRequest) SetAppVersion(v string) *InitDeviceRequest {
	s.AppVersion = &v
	return s
}

func (s *InitDeviceRequest) SetDeviceToken(v string) *InitDeviceRequest {
	s.DeviceToken = &v
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

type InitDeviceResponse struct {
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message      *string                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code         *string                         `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResultObject *InitDeviceResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s InitDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s InitDeviceResponse) GoString() string {
	return s.String()
}

func (s *InitDeviceResponse) SetRequestId(v string) *InitDeviceResponse {
	s.RequestId = &v
	return s
}

func (s *InitDeviceResponse) SetMessage(v string) *InitDeviceResponse {
	s.Message = &v
	return s
}

func (s *InitDeviceResponse) SetCode(v string) *InitDeviceResponse {
	s.Code = &v
	return s
}

func (s *InitDeviceResponse) SetResultObject(v *InitDeviceResponseResultObject) *InitDeviceResponse {
	s.ResultObject = v
	return s
}

type InitDeviceResponseResultObject struct {
	CertifyId       *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty" require:"true"`
	Protocol        *string `json:"Protocol,omitempty" xml:"Protocol,omitempty" require:"true"`
	ExtParams       *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty" require:"true"`
	RetCode         *string `json:"RetCode,omitempty" xml:"RetCode,omitempty" require:"true"`
	RetCodeSub      *string `json:"RetCodeSub,omitempty" xml:"RetCodeSub,omitempty" require:"true"`
	RetMessageSub   *string `json:"RetMessageSub,omitempty" xml:"RetMessageSub,omitempty" require:"true"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	OssEndPoint     *string `json:"OssEndPoint,omitempty" xml:"OssEndPoint,omitempty" require:"true"`
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty" require:"true"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty" require:"true"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty" require:"true"`
	BucketName      *string `json:"BucketName,omitempty" xml:"BucketName,omitempty" require:"true"`
	FileNamePrefix  *string `json:"FileNamePrefix,omitempty" xml:"FileNamePrefix,omitempty" require:"true"`
	FileName        *string `json:"FileName,omitempty" xml:"FileName,omitempty" require:"true"`
	PresignedUrl    *string `json:"PresignedUrl,omitempty" xml:"PresignedUrl,omitempty" require:"true"`
}

func (s InitDeviceResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s InitDeviceResponseResultObject) GoString() string {
	return s.String()
}

func (s *InitDeviceResponseResultObject) SetCertifyId(v string) *InitDeviceResponseResultObject {
	s.CertifyId = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetProtocol(v string) *InitDeviceResponseResultObject {
	s.Protocol = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetExtParams(v string) *InitDeviceResponseResultObject {
	s.ExtParams = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetRetCode(v string) *InitDeviceResponseResultObject {
	s.RetCode = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetRetCodeSub(v string) *InitDeviceResponseResultObject {
	s.RetCodeSub = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetRetMessageSub(v string) *InitDeviceResponseResultObject {
	s.RetMessageSub = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetMessage(v string) *InitDeviceResponseResultObject {
	s.Message = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetOssEndPoint(v string) *InitDeviceResponseResultObject {
	s.OssEndPoint = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetAccessKeyId(v string) *InitDeviceResponseResultObject {
	s.AccessKeyId = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetAccessKeySecret(v string) *InitDeviceResponseResultObject {
	s.AccessKeySecret = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetSecurityToken(v string) *InitDeviceResponseResultObject {
	s.SecurityToken = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetBucketName(v string) *InitDeviceResponseResultObject {
	s.BucketName = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetFileNamePrefix(v string) *InitDeviceResponseResultObject {
	s.FileNamePrefix = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetFileName(v string) *InitDeviceResponseResultObject {
	s.FileName = &v
	return s
}

func (s *InitDeviceResponseResultObject) SetPresignedUrl(v string) *InitDeviceResponseResultObject {
	s.PresignedUrl = &v
	return s
}

type InitFaceVerifyRequest struct {
	SceneId                *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	OuterOrderNo           *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	CertType               *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertName               *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                 *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	ReturnUrl              *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	FaceContrastPicture    *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	MetaInfo               *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	Mobile                 *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Ip                     *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	CertifyId              *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	OssBucketName          *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	Model                  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	CallbackUrl            *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	CallbackToken          *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
}

func (s InitFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyRequest) SetSceneId(v int64) *InitFaceVerifyRequest {
	s.SceneId = &v
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

func (s *InitFaceVerifyRequest) SetCertType(v string) *InitFaceVerifyRequest {
	s.CertType = &v
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

func (s *InitFaceVerifyRequest) SetReturnUrl(v string) *InitFaceVerifyRequest {
	s.ReturnUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetFaceContrastPicture(v string) *InitFaceVerifyRequest {
	s.FaceContrastPicture = &v
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

func (s *InitFaceVerifyRequest) SetIp(v string) *InitFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *InitFaceVerifyRequest) SetUserId(v string) *InitFaceVerifyRequest {
	s.UserId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *InitFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertifyId(v string) *InitFaceVerifyRequest {
	s.CertifyId = &v
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

func (s *InitFaceVerifyRequest) SetModel(v string) *InitFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCallbackUrl(v string) *InitFaceVerifyRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCallbackToken(v string) *InitFaceVerifyRequest {
	s.CallbackToken = &v
	return s
}

type InitFaceVerifyResponse struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message      *string                             `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResultObject *InitFaceVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s InitFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponse) SetRequestId(v string) *InitFaceVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *InitFaceVerifyResponse) SetMessage(v string) *InitFaceVerifyResponse {
	s.Message = &v
	return s
}

func (s *InitFaceVerifyResponse) SetCode(v string) *InitFaceVerifyResponse {
	s.Code = &v
	return s
}

func (s *InitFaceVerifyResponse) SetResultObject(v *InitFaceVerifyResponseResultObject) *InitFaceVerifyResponse {
	s.ResultObject = v
	return s
}

type InitFaceVerifyResponseResultObject struct {
	CertifyId  *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty" require:"true"`
	CertifyUrl *string `json:"CertifyUrl,omitempty" xml:"CertifyUrl,omitempty" require:"true"`
}

func (s InitFaceVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponseResultObject) SetCertifyId(v string) *InitFaceVerifyResponseResultObject {
	s.CertifyId = &v
	return s
}

func (s *InitFaceVerifyResponseResultObject) SetCertifyUrl(v string) *InitFaceVerifyResponseResultObject {
	s.CertifyUrl = &v
	return s
}

type DescribeFaceVerifyRequest struct {
	SceneId           *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	CertifyId         *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	PictureReturnType *string `json:"PictureReturnType,omitempty" xml:"PictureReturnType,omitempty"`
}

func (s DescribeFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyRequest) SetSceneId(v int64) *DescribeFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeFaceVerifyRequest) SetCertifyId(v string) *DescribeFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *DescribeFaceVerifyRequest) SetPictureReturnType(v string) *DescribeFaceVerifyRequest {
	s.PictureReturnType = &v
	return s
}

type DescribeFaceVerifyResponse struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResultObject *DescribeFaceVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s DescribeFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyResponse) SetRequestId(v string) *DescribeFaceVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeFaceVerifyResponse) SetMessage(v string) *DescribeFaceVerifyResponse {
	s.Message = &v
	return s
}

func (s *DescribeFaceVerifyResponse) SetCode(v string) *DescribeFaceVerifyResponse {
	s.Code = &v
	return s
}

func (s *DescribeFaceVerifyResponse) SetResultObject(v *DescribeFaceVerifyResponseResultObject) *DescribeFaceVerifyResponse {
	s.ResultObject = v
	return s
}

type DescribeFaceVerifyResponseResultObject struct {
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty" require:"true"`
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty" require:"true"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty" require:"true"`
	DeviceToken  *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty" require:"true"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty" require:"true"`
}

func (s DescribeFaceVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyResponseResultObject) SetPassed(v string) *DescribeFaceVerifyResponseResultObject {
	s.Passed = &v
	return s
}

func (s *DescribeFaceVerifyResponseResultObject) SetIdentityInfo(v string) *DescribeFaceVerifyResponseResultObject {
	s.IdentityInfo = &v
	return s
}

func (s *DescribeFaceVerifyResponseResultObject) SetMaterialInfo(v string) *DescribeFaceVerifyResponseResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *DescribeFaceVerifyResponseResultObject) SetDeviceToken(v string) *DescribeFaceVerifyResponseResultObject {
	s.DeviceToken = &v
	return s
}

func (s *DescribeFaceVerifyResponseResultObject) SetSubCode(v string) *DescribeFaceVerifyResponseResultObject {
	s.SubCode = &v
	return s
}

type VerifyDeviceRequest struct {
	CertifyId   *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	CertifyData *string `json:"CertifyData,omitempty" xml:"CertifyData,omitempty"`
	AppVersion  *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	ExtInfo     *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
}

func (s VerifyDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyDeviceRequest) GoString() string {
	return s.String()
}

func (s *VerifyDeviceRequest) SetCertifyId(v string) *VerifyDeviceRequest {
	s.CertifyId = &v
	return s
}

func (s *VerifyDeviceRequest) SetCertifyData(v string) *VerifyDeviceRequest {
	s.CertifyData = &v
	return s
}

func (s *VerifyDeviceRequest) SetAppVersion(v string) *VerifyDeviceRequest {
	s.AppVersion = &v
	return s
}

func (s *VerifyDeviceRequest) SetExtInfo(v string) *VerifyDeviceRequest {
	s.ExtInfo = &v
	return s
}

func (s *VerifyDeviceRequest) SetDeviceToken(v string) *VerifyDeviceRequest {
	s.DeviceToken = &v
	return s
}

type VerifyDeviceResponse struct {
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code         *string                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message      *string                           `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ResultObject *VerifyDeviceResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s VerifyDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyDeviceResponse) GoString() string {
	return s.String()
}

func (s *VerifyDeviceResponse) SetRequestId(v string) *VerifyDeviceResponse {
	s.RequestId = &v
	return s
}

func (s *VerifyDeviceResponse) SetCode(v string) *VerifyDeviceResponse {
	s.Code = &v
	return s
}

func (s *VerifyDeviceResponse) SetMessage(v string) *VerifyDeviceResponse {
	s.Message = &v
	return s
}

func (s *VerifyDeviceResponse) SetResultObject(v *VerifyDeviceResponseResultObject) *VerifyDeviceResponse {
	s.ResultObject = v
	return s
}

type VerifyDeviceResponseResultObject struct {
	ValidationRetCode *string `json:"ValidationRetCode,omitempty" xml:"ValidationRetCode,omitempty" require:"true"`
	ProductRetCode    *string `json:"ProductRetCode,omitempty" xml:"ProductRetCode,omitempty" require:"true"`
	RetCodeSub        *string `json:"RetCodeSub,omitempty" xml:"RetCodeSub,omitempty" require:"true"`
	RetMessageSub     *string `json:"RetMessageSub,omitempty" xml:"RetMessageSub,omitempty" require:"true"`
	HasNext           *string `json:"HasNext,omitempty" xml:"HasNext,omitempty" require:"true"`
	ExtParams         *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty" require:"true"`
}

func (s VerifyDeviceResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s VerifyDeviceResponseResultObject) GoString() string {
	return s.String()
}

func (s *VerifyDeviceResponseResultObject) SetValidationRetCode(v string) *VerifyDeviceResponseResultObject {
	s.ValidationRetCode = &v
	return s
}

func (s *VerifyDeviceResponseResultObject) SetProductRetCode(v string) *VerifyDeviceResponseResultObject {
	s.ProductRetCode = &v
	return s
}

func (s *VerifyDeviceResponseResultObject) SetRetCodeSub(v string) *VerifyDeviceResponseResultObject {
	s.RetCodeSub = &v
	return s
}

func (s *VerifyDeviceResponseResultObject) SetRetMessageSub(v string) *VerifyDeviceResponseResultObject {
	s.RetMessageSub = &v
	return s
}

func (s *VerifyDeviceResponseResultObject) SetHasNext(v string) *VerifyDeviceResponseResultObject {
	s.HasNext = &v
	return s
}

func (s *VerifyDeviceResponseResultObject) SetExtParams(v string) *VerifyDeviceResponseResultObject {
	s.ExtParams = &v
	return s
}

type ModifyDeviceInfoRequest struct {
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" require:"true"`
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ExpiredDay   *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
}

func (s ModifyDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceInfoRequest) SetDeviceId(v string) *ModifyDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetUserDeviceId(v string) *ModifyDeviceInfoRequest {
	s.UserDeviceId = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetBizType(v string) *ModifyDeviceInfoRequest {
	s.BizType = &v
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

type ModifyDeviceInfoResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" require:"true"`
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty" require:"true"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	BeginDay     *string `json:"BeginDay,omitempty" xml:"BeginDay,omitempty" require:"true"`
	ExpiredDay   *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty" require:"true"`
}

func (s ModifyDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceInfoResponse) SetRequestId(v string) *ModifyDeviceInfoResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyDeviceInfoResponse) SetDeviceId(v string) *ModifyDeviceInfoResponse {
	s.DeviceId = &v
	return s
}

func (s *ModifyDeviceInfoResponse) SetUserDeviceId(v string) *ModifyDeviceInfoResponse {
	s.UserDeviceId = &v
	return s
}

func (s *ModifyDeviceInfoResponse) SetBizType(v string) *ModifyDeviceInfoResponse {
	s.BizType = &v
	return s
}

func (s *ModifyDeviceInfoResponse) SetBeginDay(v string) *ModifyDeviceInfoResponse {
	s.BeginDay = &v
	return s
}

func (s *ModifyDeviceInfoResponse) SetExpiredDay(v string) *ModifyDeviceInfoResponse {
	s.ExpiredDay = &v
	return s
}

type DescribeVerifySDKRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
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

type DescribeVerifySDKResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SdkUrl    *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty" require:"true"`
}

func (s DescribeVerifySDKResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySDKResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifySDKResponse) SetRequestId(v string) *DescribeVerifySDKResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifySDKResponse) SetSdkUrl(v string) *DescribeVerifySDKResponse {
	s.SdkUrl = &v
	return s
}

type DescribeDeviceInfoRequest struct {
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage     *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceId        *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	BizType         *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	UserDeviceId    *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
	ExpiredStartDay *string `json:"ExpiredStartDay,omitempty" xml:"ExpiredStartDay,omitempty"`
	ExpiredEndDay   *string `json:"ExpiredEndDay,omitempty" xml:"ExpiredEndDay,omitempty"`
}

func (s DescribeDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoRequest) SetPageSize(v int) *DescribeDeviceInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetCurrentPage(v int) *DescribeDeviceInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetDeviceId(v string) *DescribeDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetBizType(v string) *DescribeDeviceInfoRequest {
	s.BizType = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetUserDeviceId(v string) *DescribeDeviceInfoRequest {
	s.UserDeviceId = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetExpiredStartDay(v string) *DescribeDeviceInfoRequest {
	s.ExpiredStartDay = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetExpiredEndDay(v string) *DescribeDeviceInfoRequest {
	s.ExpiredEndDay = &v
	return s
}

type DescribeDeviceInfoResponse struct {
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageSize       *int                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CurrentPage    *int                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	TotalCount     *int                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	DeviceInfoList *DescribeDeviceInfoResponseDeviceInfoList `json:"DeviceInfoList,omitempty" xml:"DeviceInfoList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponse) SetRequestId(v string) *DescribeDeviceInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceInfoResponse) SetPageSize(v int) *DescribeDeviceInfoResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceInfoResponse) SetCurrentPage(v int) *DescribeDeviceInfoResponse {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDeviceInfoResponse) SetTotalCount(v int) *DescribeDeviceInfoResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeDeviceInfoResponse) SetDeviceInfoList(v *DescribeDeviceInfoResponseDeviceInfoList) *DescribeDeviceInfoResponse {
	s.DeviceInfoList = v
	return s
}

type DescribeDeviceInfoResponseDeviceInfoList struct {
	DeviceInfo []*DescribeDeviceInfoResponseDeviceInfoListDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDeviceInfoResponseDeviceInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseDeviceInfoList) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseDeviceInfoList) SetDeviceInfo(v []*DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) *DescribeDeviceInfoResponseDeviceInfoList {
	s.DeviceInfo = v
	return s
}

type DescribeDeviceInfoResponseDeviceInfoListDeviceInfo struct {
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" require:"true"`
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty" require:"true"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	BeginDay     *string `json:"BeginDay,omitempty" xml:"BeginDay,omitempty" require:"true"`
	ExpiredDay   *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty" require:"true"`
}

func (s DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) SetDeviceId(v string) *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) SetUserDeviceId(v string) *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo {
	s.UserDeviceId = &v
	return s
}

func (s *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) SetBizType(v string) *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo {
	s.BizType = &v
	return s
}

func (s *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) SetBeginDay(v string) *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo {
	s.BeginDay = &v
	return s
}

func (s *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) SetExpiredDay(v string) *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo {
	s.ExpiredDay = &v
	return s
}

type CreateVerifySDKRequest struct {
	AppUrl   *string `json:"AppUrl,omitempty" xml:"AppUrl,omitempty" require:"true"`
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

type CreateVerifySDKResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s CreateVerifySDKResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySDKResponse) GoString() string {
	return s.String()
}

func (s *CreateVerifySDKResponse) SetRequestId(v string) *CreateVerifySDKResponse {
	s.RequestId = &v
	return s
}

func (s *CreateVerifySDKResponse) SetTaskId(v string) *CreateVerifySDKResponse {
	s.TaskId = &v
	return s
}

type CreateAuthKeyRequest struct {
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
	Test         *bool   `json:"Test,omitempty" xml:"Test,omitempty"`
	AuthYears    *int    `json:"AuthYears,omitempty" xml:"AuthYears,omitempty"`
}

func (s CreateAuthKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthKeyRequest) SetBizType(v string) *CreateAuthKeyRequest {
	s.BizType = &v
	return s
}

func (s *CreateAuthKeyRequest) SetUserDeviceId(v string) *CreateAuthKeyRequest {
	s.UserDeviceId = &v
	return s
}

func (s *CreateAuthKeyRequest) SetTest(v bool) *CreateAuthKeyRequest {
	s.Test = &v
	return s
}

func (s *CreateAuthKeyRequest) SetAuthYears(v int) *CreateAuthKeyRequest {
	s.AuthYears = &v
	return s
}

type CreateAuthKeyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AuthKey   *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty" require:"true"`
}

func (s CreateAuthKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateAuthKeyResponse) SetRequestId(v string) *CreateAuthKeyResponse {
	s.RequestId = &v
	return s
}

func (s *CreateAuthKeyResponse) SetAuthKey(v string) *CreateAuthKeyResponse {
	s.AuthKey = &v
	return s
}

type DetectFaceAttributesRequest struct {
	MaterialValue *string `json:"MaterialValue,omitempty" xml:"MaterialValue,omitempty" require:"true"`
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s DetectFaceAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesRequest) SetMaterialValue(v string) *DetectFaceAttributesRequest {
	s.MaterialValue = &v
	return s
}

func (s *DetectFaceAttributesRequest) SetBizType(v string) *DetectFaceAttributesRequest {
	s.BizType = &v
	return s
}

type DetectFaceAttributesResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *DetectFaceAttributesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s DetectFaceAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponse) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponse) SetRequestId(v string) *DetectFaceAttributesResponse {
	s.RequestId = &v
	return s
}

func (s *DetectFaceAttributesResponse) SetSuccess(v bool) *DetectFaceAttributesResponse {
	s.Success = &v
	return s
}

func (s *DetectFaceAttributesResponse) SetCode(v string) *DetectFaceAttributesResponse {
	s.Code = &v
	return s
}

func (s *DetectFaceAttributesResponse) SetMessage(v string) *DetectFaceAttributesResponse {
	s.Message = &v
	return s
}

func (s *DetectFaceAttributesResponse) SetData(v *DetectFaceAttributesResponseData) *DetectFaceAttributesResponse {
	s.Data = v
	return s
}

type DetectFaceAttributesResponseData struct {
	ImgWidth  *int                                       `json:"ImgWidth,omitempty" xml:"ImgWidth,omitempty" require:"true"`
	ImgHeight *int                                       `json:"ImgHeight,omitempty" xml:"ImgHeight,omitempty" require:"true"`
	FaceInfos *DetectFaceAttributesResponseDataFaceInfos `json:"FaceInfos,omitempty" xml:"FaceInfos,omitempty" require:"true" type:"Struct"`
}

func (s DetectFaceAttributesResponseData) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseData) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseData) SetImgWidth(v int) *DetectFaceAttributesResponseData {
	s.ImgWidth = &v
	return s
}

func (s *DetectFaceAttributesResponseData) SetImgHeight(v int) *DetectFaceAttributesResponseData {
	s.ImgHeight = &v
	return s
}

func (s *DetectFaceAttributesResponseData) SetFaceInfos(v *DetectFaceAttributesResponseDataFaceInfos) *DetectFaceAttributesResponseData {
	s.FaceInfos = v
	return s
}

type DetectFaceAttributesResponseDataFaceInfos struct {
	FaceAttributesDetectInfo []*DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo `json:"FaceAttributesDetectInfo,omitempty" xml:"FaceAttributesDetectInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DetectFaceAttributesResponseDataFaceInfos) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseDataFaceInfos) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseDataFaceInfos) SetFaceAttributesDetectInfo(v []*DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo) *DetectFaceAttributesResponseDataFaceInfos {
	s.FaceAttributesDetectInfo = v
	return s
}

type DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo struct {
	FaceRect       *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect       `json:"FaceRect,omitempty" xml:"FaceRect,omitempty" require:"true" type:"Struct"`
	FaceAttributes *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" require:"true" type:"Struct"`
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo) SetFaceRect(v *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo {
	s.FaceRect = v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo) SetFaceAttributes(v *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo {
	s.FaceAttributes = v
	return s
}

type DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect struct {
	Top    *int `json:"Top,omitempty" xml:"Top,omitempty" require:"true"`
	Left   *int `json:"Left,omitempty" xml:"Left,omitempty" require:"true"`
	Width  *int `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
	Height *int `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) SetTop(v int) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Top = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) SetLeft(v int) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Left = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) SetWidth(v int) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Width = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) SetHeight(v int) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Height = &v
	return s
}

type DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes struct {
	Glasses         *string                                                                                  `json:"Glasses,omitempty" xml:"Glasses,omitempty" require:"true"`
	Facetype        *string                                                                                  `json:"Facetype,omitempty" xml:"Facetype,omitempty" require:"true"`
	Blur            *float32                                                                                 `json:"Blur,omitempty" xml:"Blur,omitempty" require:"true"`
	Facequal        *float32                                                                                 `json:"Facequal,omitempty" xml:"Facequal,omitempty" require:"true"`
	Integrity       *int                                                                                     `json:"Integrity,omitempty" xml:"Integrity,omitempty" require:"true"`
	Respirator      *string                                                                                  `json:"Respirator,omitempty" xml:"Respirator,omitempty" require:"true"`
	AppearanceScore *float32                                                                                 `json:"AppearanceScore,omitempty" xml:"AppearanceScore,omitempty" require:"true"`
	Smiling         *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling  `json:"Smiling,omitempty" xml:"Smiling,omitempty" require:"true" type:"Struct"`
	Headpose        *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose `json:"Headpose,omitempty" xml:"Headpose,omitempty" require:"true" type:"Struct"`
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetGlasses(v string) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetFacetype(v string) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Facetype = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetBlur(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Blur = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetFacequal(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Facequal = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetIntegrity(v int) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Integrity = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetRespirator(v string) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Respirator = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetAppearanceScore(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.AppearanceScore = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetSmiling(v *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Smiling = v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetHeadpose(v *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Headpose = v
	return s
}

type DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling struct {
	Value     *float32 `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty" require:"true"`
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) SetValue(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling {
	s.Value = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) SetThreshold(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling {
	s.Threshold = &v
	return s
}

type DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose struct {
	PitchAngle *float32 `json:"PitchAngle,omitempty" xml:"PitchAngle,omitempty" require:"true"`
	RollAngle  *float32 `json:"RollAngle,omitempty" xml:"RollAngle,omitempty" require:"true"`
	YawAngle   *float32 `json:"YawAngle,omitempty" xml:"YawAngle,omitempty" require:"true"`
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetPitchAngle(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.PitchAngle = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetRollAngle(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.RollAngle = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetYawAngle(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.YawAngle = &v
	return s
}

type CompareFacesRequest struct {
	TargetImageType  *string `json:"TargetImageType,omitempty" xml:"TargetImageType,omitempty"`
	SourceImageType  *string `json:"SourceImageType,omitempty" xml:"SourceImageType,omitempty"`
	SourceImageValue *string `json:"SourceImageValue,omitempty" xml:"SourceImageValue,omitempty"`
	TargetImageValue *string `json:"TargetImageValue,omitempty" xml:"TargetImageValue,omitempty"`
	BizType          *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s CompareFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesRequest) GoString() string {
	return s.String()
}

func (s *CompareFacesRequest) SetTargetImageType(v string) *CompareFacesRequest {
	s.TargetImageType = &v
	return s
}

func (s *CompareFacesRequest) SetSourceImageType(v string) *CompareFacesRequest {
	s.SourceImageType = &v
	return s
}

func (s *CompareFacesRequest) SetSourceImageValue(v string) *CompareFacesRequest {
	s.SourceImageValue = &v
	return s
}

func (s *CompareFacesRequest) SetTargetImageValue(v string) *CompareFacesRequest {
	s.TargetImageValue = &v
	return s
}

func (s *CompareFacesRequest) SetBizType(v string) *CompareFacesRequest {
	s.BizType = &v
	return s
}

type CompareFacesResponse struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool                     `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Code      *string                   `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *CompareFacesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s CompareFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesResponse) GoString() string {
	return s.String()
}

func (s *CompareFacesResponse) SetRequestId(v string) *CompareFacesResponse {
	s.RequestId = &v
	return s
}

func (s *CompareFacesResponse) SetSuccess(v bool) *CompareFacesResponse {
	s.Success = &v
	return s
}

func (s *CompareFacesResponse) SetCode(v string) *CompareFacesResponse {
	s.Code = &v
	return s
}

func (s *CompareFacesResponse) SetMessage(v string) *CompareFacesResponse {
	s.Message = &v
	return s
}

func (s *CompareFacesResponse) SetData(v *CompareFacesResponseData) *CompareFacesResponse {
	s.Data = v
	return s
}

type CompareFacesResponseData struct {
	SimilarityScore      *float32 `json:"SimilarityScore,omitempty" xml:"SimilarityScore,omitempty" require:"true"`
	ConfidenceThresholds *string  `json:"ConfidenceThresholds,omitempty" xml:"ConfidenceThresholds,omitempty" require:"true"`
}

func (s CompareFacesResponseData) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesResponseData) GoString() string {
	return s.String()
}

func (s *CompareFacesResponseData) SetSimilarityScore(v float32) *CompareFacesResponseData {
	s.SimilarityScore = &v
	return s
}

func (s *CompareFacesResponseData) SetConfidenceThresholds(v string) *CompareFacesResponseData {
	s.ConfidenceThresholds = &v
	return s
}

type DescribeFaceUsageRequest struct {
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s DescribeFaceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaceUsageRequest) SetStartDate(v string) *DescribeFaceUsageRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeFaceUsageRequest) SetEndDate(v string) *DescribeFaceUsageRequest {
	s.EndDate = &v
	return s
}

type DescribeFaceUsageResponse struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount    *int                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	FaceUsageList []*DescribeFaceUsageResponseFaceUsageList `json:"FaceUsageList,omitempty" xml:"FaceUsageList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeFaceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaceUsageResponse) SetRequestId(v string) *DescribeFaceUsageResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeFaceUsageResponse) SetTotalCount(v int) *DescribeFaceUsageResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeFaceUsageResponse) SetFaceUsageList(v []*DescribeFaceUsageResponseFaceUsageList) *DescribeFaceUsageResponse {
	s.FaceUsageList = v
	return s
}

type DescribeFaceUsageResponseFaceUsageList struct {
	Date       *string `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
}

func (s DescribeFaceUsageResponseFaceUsageList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceUsageResponseFaceUsageList) GoString() string {
	return s.String()
}

func (s *DescribeFaceUsageResponseFaceUsageList) SetDate(v string) *DescribeFaceUsageResponseFaceUsageList {
	s.Date = &v
	return s
}

func (s *DescribeFaceUsageResponseFaceUsageList) SetTotalCount(v int64) *DescribeFaceUsageResponseFaceUsageList {
	s.TotalCount = &v
	return s
}

type DescribeVerifyRecordsRequest struct {
	TotalCount  *int    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CurrentPage *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate     *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	IdCardNum   *string `json:"IdCardNum,omitempty" xml:"IdCardNum,omitempty"`
	StatusList  *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	QueryId     *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s DescribeVerifyRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyRecordsRequest) SetTotalCount(v int) *DescribeVerifyRecordsRequest {
	s.TotalCount = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetPageSize(v int) *DescribeVerifyRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetCurrentPage(v int) *DescribeVerifyRecordsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetBizType(v string) *DescribeVerifyRecordsRequest {
	s.BizType = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetStartDate(v string) *DescribeVerifyRecordsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetEndDate(v string) *DescribeVerifyRecordsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetBizId(v string) *DescribeVerifyRecordsRequest {
	s.BizId = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetIdCardNum(v string) *DescribeVerifyRecordsRequest {
	s.IdCardNum = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetStatusList(v string) *DescribeVerifyRecordsRequest {
	s.StatusList = &v
	return s
}

func (s *DescribeVerifyRecordsRequest) SetQueryId(v string) *DescribeVerifyRecordsRequest {
	s.QueryId = &v
	return s
}

type DescribeVerifyRecordsResponse struct {
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount  *int                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize    *int                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CurrentPage *int                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	QueryId     *string                                     `json:"QueryId,omitempty" xml:"QueryId,omitempty" require:"true"`
	RecordsList []*DescribeVerifyRecordsResponseRecordsList `json:"RecordsList,omitempty" xml:"RecordsList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVerifyRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyRecordsResponse) SetRequestId(v string) *DescribeVerifyRecordsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyRecordsResponse) SetTotalCount(v int) *DescribeVerifyRecordsResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeVerifyRecordsResponse) SetPageSize(v int) *DescribeVerifyRecordsResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeVerifyRecordsResponse) SetCurrentPage(v int) *DescribeVerifyRecordsResponse {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVerifyRecordsResponse) SetQueryId(v string) *DescribeVerifyRecordsResponse {
	s.QueryId = &v
	return s
}

func (s *DescribeVerifyRecordsResponse) SetRecordsList(v []*DescribeVerifyRecordsResponseRecordsList) *DescribeVerifyRecordsResponse {
	s.RecordsList = v
	return s
}

type DescribeVerifyRecordsResponseRecordsList struct {
	BizType                   *string                                           `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	BizId                     *string                                           `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	DataStats                 *string                                           `json:"DataStats,omitempty" xml:"DataStats,omitempty" require:"true"`
	VerifyId                  *string                                           `json:"VerifyId,omitempty" xml:"VerifyId,omitempty" require:"true"`
	FinishTime                *int64                                            `json:"FinishTime,omitempty" xml:"FinishTime,omitempty" require:"true"`
	Status                    *int                                              `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	IdCardFaceComparisonScore *float32                                          `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty" require:"true"`
	AuthorityComparisonScore  *float32                                          `json:"AuthorityComparisonScore,omitempty" xml:"AuthorityComparisonScore,omitempty" require:"true"`
	Material                  *DescribeVerifyRecordsResponseRecordsListMaterial `json:"Material,omitempty" xml:"Material,omitempty" require:"true" type:"Struct"`
}

func (s DescribeVerifyRecordsResponseRecordsList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyRecordsResponseRecordsList) GoString() string {
	return s.String()
}

func (s *DescribeVerifyRecordsResponseRecordsList) SetBizType(v string) *DescribeVerifyRecordsResponseRecordsList {
	s.BizType = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsList) SetBizId(v string) *DescribeVerifyRecordsResponseRecordsList {
	s.BizId = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsList) SetDataStats(v string) *DescribeVerifyRecordsResponseRecordsList {
	s.DataStats = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsList) SetVerifyId(v string) *DescribeVerifyRecordsResponseRecordsList {
	s.VerifyId = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsList) SetFinishTime(v int64) *DescribeVerifyRecordsResponseRecordsList {
	s.FinishTime = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsList) SetStatus(v int) *DescribeVerifyRecordsResponseRecordsList {
	s.Status = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsList) SetIdCardFaceComparisonScore(v float32) *DescribeVerifyRecordsResponseRecordsList {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsList) SetAuthorityComparisonScore(v float32) *DescribeVerifyRecordsResponseRecordsList {
	s.AuthorityComparisonScore = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsList) SetMaterial(v *DescribeVerifyRecordsResponseRecordsListMaterial) *DescribeVerifyRecordsResponseRecordsList {
	s.Material = v
	return s
}

type DescribeVerifyRecordsResponseRecordsListMaterial struct {
	FaceImageUrl *string                                                     `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty" require:"true"`
	IdCardName   *string                                                     `json:"IdCardName,omitempty" xml:"IdCardName,omitempty" require:"true"`
	IdCardNumber *string                                                     `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty" require:"true"`
	IdCardInfo   *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribeVerifyRecordsResponseRecordsListMaterial) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyRecordsResponseRecordsListMaterial) GoString() string {
	return s.String()
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterial) SetFaceImageUrl(v string) *DescribeVerifyRecordsResponseRecordsListMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterial) SetIdCardName(v string) *DescribeVerifyRecordsResponseRecordsListMaterial {
	s.IdCardName = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterial) SetIdCardNumber(v string) *DescribeVerifyRecordsResponseRecordsListMaterial {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterial) SetIdCardInfo(v *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) *DescribeVerifyRecordsResponseRecordsListMaterial {
	s.IdCardInfo = v
	return s
}

type DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo struct {
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty" require:"true"`
	BackImageUrl  *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Number        *string `json:"Number,omitempty" xml:"Number,omitempty" require:"true"`
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty" require:"true"`
	Birth         *string `json:"Birth,omitempty" xml:"Birth,omitempty" require:"true"`
	Sex           *string `json:"Sex,omitempty" xml:"Sex,omitempty" require:"true"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty" require:"true"`
	Authority     *string `json:"Authority,omitempty" xml:"Authority,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) SetFrontImageUrl(v string) *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) SetBackImageUrl(v string) *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) SetName(v string) *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) SetNumber(v string) *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) SetAddress(v string) *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) SetBirth(v string) *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) SetSex(v string) *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo {
	s.Sex = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) SetNationality(v string) *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) SetAuthority(v string) *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) SetStartDate(v string) *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

func (s *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo) SetEndDate(v string) *DescribeVerifyRecordsResponseRecordsListMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

type UpdateVerifySettingRequest struct {
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	BizName     *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	Solution    *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	GuideStep   *bool   `json:"GuideStep,omitempty" xml:"GuideStep,omitempty"`
	PrivacyStep *bool   `json:"PrivacyStep,omitempty" xml:"PrivacyStep,omitempty"`
	ResultStep  *bool   `json:"ResultStep,omitempty" xml:"ResultStep,omitempty"`
}

func (s UpdateVerifySettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVerifySettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateVerifySettingRequest) SetBizType(v string) *UpdateVerifySettingRequest {
	s.BizType = &v
	return s
}

func (s *UpdateVerifySettingRequest) SetBizName(v string) *UpdateVerifySettingRequest {
	s.BizName = &v
	return s
}

func (s *UpdateVerifySettingRequest) SetSolution(v string) *UpdateVerifySettingRequest {
	s.Solution = &v
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

type UpdateVerifySettingResponse struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizType   *string   `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	BizName   *string   `json:"BizName,omitempty" xml:"BizName,omitempty" require:"true"`
	Solution  *string   `json:"Solution,omitempty" xml:"Solution,omitempty" require:"true"`
	StepList  []*string `json:"StepList,omitempty" xml:"StepList,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateVerifySettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVerifySettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateVerifySettingResponse) SetRequestId(v string) *UpdateVerifySettingResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateVerifySettingResponse) SetBizType(v string) *UpdateVerifySettingResponse {
	s.BizType = &v
	return s
}

func (s *UpdateVerifySettingResponse) SetBizName(v string) *UpdateVerifySettingResponse {
	s.BizName = &v
	return s
}

func (s *UpdateVerifySettingResponse) SetSolution(v string) *UpdateVerifySettingResponse {
	s.Solution = &v
	return s
}

func (s *UpdateVerifySettingResponse) SetStepList(v []*string) *UpdateVerifySettingResponse {
	s.StepList = v
	return s
}

type CreateVerifySettingRequest struct {
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	BizName     *string `json:"BizName,omitempty" xml:"BizName,omitempty" require:"true"`
	Solution    *string `json:"Solution,omitempty" xml:"Solution,omitempty" require:"true"`
	GuideStep   *bool   `json:"GuideStep,omitempty" xml:"GuideStep,omitempty"`
	PrivacyStep *bool   `json:"PrivacyStep,omitempty" xml:"PrivacyStep,omitempty"`
	ResultStep  *bool   `json:"ResultStep,omitempty" xml:"ResultStep,omitempty"`
}

func (s CreateVerifySettingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySettingRequest) GoString() string {
	return s.String()
}

func (s *CreateVerifySettingRequest) SetBizType(v string) *CreateVerifySettingRequest {
	s.BizType = &v
	return s
}

func (s *CreateVerifySettingRequest) SetBizName(v string) *CreateVerifySettingRequest {
	s.BizName = &v
	return s
}

func (s *CreateVerifySettingRequest) SetSolution(v string) *CreateVerifySettingRequest {
	s.Solution = &v
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

type CreateVerifySettingResponse struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizType   *string   `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	BizName   *string   `json:"BizName,omitempty" xml:"BizName,omitempty" require:"true"`
	Solution  *string   `json:"Solution,omitempty" xml:"Solution,omitempty" require:"true"`
	StepList  []*string `json:"StepList,omitempty" xml:"StepList,omitempty" require:"true" type:"Repeated"`
}

func (s CreateVerifySettingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySettingResponse) GoString() string {
	return s.String()
}

func (s *CreateVerifySettingResponse) SetRequestId(v string) *CreateVerifySettingResponse {
	s.RequestId = &v
	return s
}

func (s *CreateVerifySettingResponse) SetBizType(v string) *CreateVerifySettingResponse {
	s.BizType = &v
	return s
}

func (s *CreateVerifySettingResponse) SetBizName(v string) *CreateVerifySettingResponse {
	s.BizName = &v
	return s
}

func (s *CreateVerifySettingResponse) SetSolution(v string) *CreateVerifySettingResponse {
	s.Solution = &v
	return s
}

func (s *CreateVerifySettingResponse) SetStepList(v []*string) *CreateVerifySettingResponse {
	s.StepList = v
	return s
}

type DescribeVerifySettingRequest struct {
}

func (s DescribeVerifySettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySettingRequest) GoString() string {
	return s.String()
}

type DescribeVerifySettingResponse struct {
	RequestId         *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VerifySettingList []*DescribeVerifySettingResponseVerifySettingList `json:"VerifySettingList,omitempty" xml:"VerifySettingList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVerifySettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifySettingResponse) SetRequestId(v string) *DescribeVerifySettingResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifySettingResponse) SetVerifySettingList(v []*DescribeVerifySettingResponseVerifySettingList) *DescribeVerifySettingResponse {
	s.VerifySettingList = v
	return s
}

type DescribeVerifySettingResponseVerifySettingList struct {
	BizType  *string   `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	BizName  *string   `json:"BizName,omitempty" xml:"BizName,omitempty" require:"true"`
	Solution *string   `json:"Solution,omitempty" xml:"Solution,omitempty" require:"true"`
	StepList []*string `json:"StepList,omitempty" xml:"StepList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVerifySettingResponseVerifySettingList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySettingResponseVerifySettingList) GoString() string {
	return s.String()
}

func (s *DescribeVerifySettingResponseVerifySettingList) SetBizType(v string) *DescribeVerifySettingResponseVerifySettingList {
	s.BizType = &v
	return s
}

func (s *DescribeVerifySettingResponseVerifySettingList) SetBizName(v string) *DescribeVerifySettingResponseVerifySettingList {
	s.BizName = &v
	return s
}

func (s *DescribeVerifySettingResponseVerifySettingList) SetSolution(v string) *DescribeVerifySettingResponseVerifySettingList {
	s.Solution = &v
	return s
}

func (s *DescribeVerifySettingResponseVerifySettingList) SetStepList(v []*string) *DescribeVerifySettingResponseVerifySettingList {
	s.StepList = v
	return s
}

type DescribeVerifyUsageRequest struct {
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
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

func (s *DescribeVerifyUsageRequest) SetStartDate(v string) *DescribeVerifyUsageRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeVerifyUsageRequest) SetEndDate(v string) *DescribeVerifyUsageRequest {
	s.EndDate = &v
	return s
}

type DescribeVerifyUsageResponse struct {
	RequestId       *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount      *int                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	VerifyUsageList []*DescribeVerifyUsageResponseVerifyUsageList `json:"VerifyUsageList,omitempty" xml:"VerifyUsageList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVerifyUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyUsageResponse) SetRequestId(v string) *DescribeVerifyUsageResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyUsageResponse) SetTotalCount(v int) *DescribeVerifyUsageResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeVerifyUsageResponse) SetVerifyUsageList(v []*DescribeVerifyUsageResponseVerifyUsageList) *DescribeVerifyUsageResponse {
	s.VerifyUsageList = v
	return s
}

type DescribeVerifyUsageResponseVerifyUsageList struct {
	BizType    *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	Date       *string `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PassCount  *int64  `json:"PassCount,omitempty" xml:"PassCount,omitempty" require:"true"`
	FailCount  *int64  `json:"FailCount,omitempty" xml:"FailCount,omitempty" require:"true"`
}

func (s DescribeVerifyUsageResponseVerifyUsageList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyUsageResponseVerifyUsageList) GoString() string {
	return s.String()
}

func (s *DescribeVerifyUsageResponseVerifyUsageList) SetBizType(v string) *DescribeVerifyUsageResponseVerifyUsageList {
	s.BizType = &v
	return s
}

func (s *DescribeVerifyUsageResponseVerifyUsageList) SetDate(v string) *DescribeVerifyUsageResponseVerifyUsageList {
	s.Date = &v
	return s
}

func (s *DescribeVerifyUsageResponseVerifyUsageList) SetTotalCount(v int64) *DescribeVerifyUsageResponseVerifyUsageList {
	s.TotalCount = &v
	return s
}

func (s *DescribeVerifyUsageResponseVerifyUsageList) SetPassCount(v int64) *DescribeVerifyUsageResponseVerifyUsageList {
	s.PassCount = &v
	return s
}

func (s *DescribeVerifyUsageResponseVerifyUsageList) SetFailCount(v int64) *DescribeVerifyUsageResponseVerifyUsageList {
	s.FailCount = &v
	return s
}

type DescribeUserStatusRequest struct {
}

func (s DescribeUserStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserStatusRequest) GoString() string {
	return s.String()
}

type DescribeUserStatusResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Enabled   *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty" require:"true"`
}

func (s DescribeUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusResponse) SetRequestId(v string) *DescribeUserStatusResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUserStatusResponse) SetEnabled(v bool) *DescribeUserStatusResponse {
	s.Enabled = &v
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

type DescribeUploadInfoResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Accessid  *string `json:"Accessid,omitempty" xml:"Accessid,omitempty" require:"true"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty" require:"true"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty" require:"true"`
	Folder    *string `json:"Folder,omitempty" xml:"Folder,omitempty" require:"true"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty" require:"true"`
	Expire    *int64  `json:"Expire,omitempty" xml:"Expire,omitempty" require:"true"`
}

func (s DescribeUploadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeUploadInfoResponse) SetRequestId(v string) *DescribeUploadInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUploadInfoResponse) SetAccessid(v string) *DescribeUploadInfoResponse {
	s.Accessid = &v
	return s
}

func (s *DescribeUploadInfoResponse) SetPolicy(v string) *DescribeUploadInfoResponse {
	s.Policy = &v
	return s
}

func (s *DescribeUploadInfoResponse) SetSignature(v string) *DescribeUploadInfoResponse {
	s.Signature = &v
	return s
}

func (s *DescribeUploadInfoResponse) SetFolder(v string) *DescribeUploadInfoResponse {
	s.Folder = &v
	return s
}

func (s *DescribeUploadInfoResponse) SetHost(v string) *DescribeUploadInfoResponse {
	s.Host = &v
	return s
}

func (s *DescribeUploadInfoResponse) SetExpire(v int64) *DescribeUploadInfoResponse {
	s.Expire = &v
	return s
}

type DescribeRPSDKRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s DescribeRPSDKRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRPSDKRequest) GoString() string {
	return s.String()
}

func (s *DescribeRPSDKRequest) SetSourceIp(v string) *DescribeRPSDKRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRPSDKRequest) SetLang(v string) *DescribeRPSDKRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRPSDKRequest) SetTaskId(v string) *DescribeRPSDKRequest {
	s.TaskId = &v
	return s
}

type DescribeRPSDKResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SdkUrl    *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty" require:"true"`
}

func (s DescribeRPSDKResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRPSDKResponse) GoString() string {
	return s.String()
}

func (s *DescribeRPSDKResponse) SetRequestId(v string) *DescribeRPSDKResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRPSDKResponse) SetSdkUrl(v string) *DescribeRPSDKResponse {
	s.SdkUrl = &v
	return s
}

type CreateRPSDKRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppUrl   *string `json:"AppUrl,omitempty" xml:"AppUrl,omitempty" require:"true"`
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s CreateRPSDKRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRPSDKRequest) GoString() string {
	return s.String()
}

func (s *CreateRPSDKRequest) SetSourceIp(v string) *CreateRPSDKRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateRPSDKRequest) SetLang(v string) *CreateRPSDKRequest {
	s.Lang = &v
	return s
}

func (s *CreateRPSDKRequest) SetAppUrl(v string) *CreateRPSDKRequest {
	s.AppUrl = &v
	return s
}

func (s *CreateRPSDKRequest) SetPlatform(v string) *CreateRPSDKRequest {
	s.Platform = &v
	return s
}

type CreateRPSDKResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s CreateRPSDKResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRPSDKResponse) GoString() string {
	return s.String()
}

func (s *CreateRPSDKResponse) SetRequestId(v string) *CreateRPSDKResponse {
	s.RequestId = &v
	return s
}

func (s *CreateRPSDKResponse) SetTaskId(v string) *CreateRPSDKResponse {
	s.TaskId = &v
	return s
}

type VerifyMaterialRequest struct {
	IdCardBackImageUrl  *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	FaceImageUrl        *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty" require:"true"`
	BizType             *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	BizId               *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	IdCardNumber        *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty" require:"true"`
	IdCardFrontImageUrl *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s VerifyMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialRequest) GoString() string {
	return s.String()
}

func (s *VerifyMaterialRequest) SetIdCardBackImageUrl(v string) *VerifyMaterialRequest {
	s.IdCardBackImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetFaceImageUrl(v string) *VerifyMaterialRequest {
	s.FaceImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetBizType(v string) *VerifyMaterialRequest {
	s.BizType = &v
	return s
}

func (s *VerifyMaterialRequest) SetBizId(v string) *VerifyMaterialRequest {
	s.BizId = &v
	return s
}

func (s *VerifyMaterialRequest) SetName(v string) *VerifyMaterialRequest {
	s.Name = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardNumber(v string) *VerifyMaterialRequest {
	s.IdCardNumber = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardFrontImageUrl(v string) *VerifyMaterialRequest {
	s.IdCardFrontImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetUserId(v string) *VerifyMaterialRequest {
	s.UserId = &v
	return s
}

type VerifyMaterialResponse struct {
	RequestId                 *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VerifyToken               *string                         `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty" require:"true"`
	VerifyStatus              *int                            `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty" require:"true"`
	AuthorityComparisionScore *float32                        `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty" require:"true"`
	IdCardFaceComparisonScore *float32                        `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty" require:"true"`
	Material                  *VerifyMaterialResponseMaterial `json:"Material,omitempty" xml:"Material,omitempty" require:"true" type:"Struct"`
}

func (s VerifyMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponse) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponse) SetRequestId(v string) *VerifyMaterialResponse {
	s.RequestId = &v
	return s
}

func (s *VerifyMaterialResponse) SetVerifyToken(v string) *VerifyMaterialResponse {
	s.VerifyToken = &v
	return s
}

func (s *VerifyMaterialResponse) SetVerifyStatus(v int) *VerifyMaterialResponse {
	s.VerifyStatus = &v
	return s
}

func (s *VerifyMaterialResponse) SetAuthorityComparisionScore(v float32) *VerifyMaterialResponse {
	s.AuthorityComparisionScore = &v
	return s
}

func (s *VerifyMaterialResponse) SetIdCardFaceComparisonScore(v float32) *VerifyMaterialResponse {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *VerifyMaterialResponse) SetMaterial(v *VerifyMaterialResponseMaterial) *VerifyMaterialResponse {
	s.Material = v
	return s
}

type VerifyMaterialResponseMaterial struct {
	FaceImageUrl  *string                                   `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty" require:"true"`
	IdCardName    *string                                   `json:"IdCardName,omitempty" xml:"IdCardName,omitempty" require:"true"`
	IdCardNumber  *string                                   `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty" require:"true"`
	FaceQuality   *string                                   `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty" require:"true"`
	FaceGlobalUrl *string                                   `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty" require:"true"`
	FaceMask      *string                                   `json:"FaceMask,omitempty" xml:"FaceMask,omitempty" require:"true"`
	IdCardInfo    *VerifyMaterialResponseMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" require:"true" type:"Struct"`
}

func (s VerifyMaterialResponseMaterial) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseMaterial) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseMaterial) SetFaceImageUrl(v string) *VerifyMaterialResponseMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseMaterial) SetIdCardName(v string) *VerifyMaterialResponseMaterial {
	s.IdCardName = &v
	return s
}

func (s *VerifyMaterialResponseMaterial) SetIdCardNumber(v string) *VerifyMaterialResponseMaterial {
	s.IdCardNumber = &v
	return s
}

func (s *VerifyMaterialResponseMaterial) SetFaceQuality(v string) *VerifyMaterialResponseMaterial {
	s.FaceQuality = &v
	return s
}

func (s *VerifyMaterialResponseMaterial) SetFaceGlobalUrl(v string) *VerifyMaterialResponseMaterial {
	s.FaceGlobalUrl = &v
	return s
}

func (s *VerifyMaterialResponseMaterial) SetFaceMask(v string) *VerifyMaterialResponseMaterial {
	s.FaceMask = &v
	return s
}

func (s *VerifyMaterialResponseMaterial) SetIdCardInfo(v *VerifyMaterialResponseMaterialIdCardInfo) *VerifyMaterialResponseMaterial {
	s.IdCardInfo = v
	return s
}

type VerifyMaterialResponseMaterialIdCardInfo struct {
	Number        *string `json:"Number,omitempty" xml:"Number,omitempty" require:"true"`
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty" require:"true"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty" require:"true"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty" require:"true"`
	Authority     *string `json:"Authority,omitempty" xml:"Authority,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Birth         *string `json:"Birth,omitempty" xml:"Birth,omitempty" require:"true"`
	BackImageUrl  *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
}

func (s VerifyMaterialResponseMaterialIdCardInfo) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetNumber(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetAddress(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetNationality(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetEndDate(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetFrontImageUrl(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetAuthority(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetName(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetBirth(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetBackImageUrl(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetStartDate(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

type DescribeVerifyResultRequest struct {
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
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

type DescribeVerifyResultResponse struct {
	RequestId                 *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VerifyStatus              *int                                  `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty" require:"true"`
	AuthorityComparisionScore *float32                              `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty" require:"true"`
	FaceComparisonScore       *float32                              `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty" require:"true"`
	IdCardFaceComparisonScore *float32                              `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty" require:"true"`
	Material                  *DescribeVerifyResultResponseMaterial `json:"Material,omitempty" xml:"Material,omitempty" require:"true" type:"Struct"`
}

func (s DescribeVerifyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponse) SetRequestId(v string) *DescribeVerifyResultResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyResultResponse) SetVerifyStatus(v int) *DescribeVerifyResultResponse {
	s.VerifyStatus = &v
	return s
}

func (s *DescribeVerifyResultResponse) SetAuthorityComparisionScore(v float32) *DescribeVerifyResultResponse {
	s.AuthorityComparisionScore = &v
	return s
}

func (s *DescribeVerifyResultResponse) SetFaceComparisonScore(v float32) *DescribeVerifyResultResponse {
	s.FaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyResultResponse) SetIdCardFaceComparisonScore(v float32) *DescribeVerifyResultResponse {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyResultResponse) SetMaterial(v *DescribeVerifyResultResponseMaterial) *DescribeVerifyResultResponse {
	s.Material = v
	return s
}

type DescribeVerifyResultResponseMaterial struct {
	FaceImageUrl  *string                                         `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty" require:"true"`
	IdCardName    *string                                         `json:"IdCardName,omitempty" xml:"IdCardName,omitempty" require:"true"`
	IdCardNumber  *string                                         `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty" require:"true"`
	FaceQuality   *string                                         `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty" require:"true"`
	FaceGlobalUrl *string                                         `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty" require:"true"`
	FaceMask      *bool                                           `json:"FaceMask,omitempty" xml:"FaceMask,omitempty" require:"true"`
	IdCardInfo    *DescribeVerifyResultResponseMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" require:"true" type:"Struct"`
	VideoUrls     []*string                                       `json:"VideoUrls,omitempty" xml:"VideoUrls,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVerifyResultResponseMaterial) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseMaterial) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseMaterial) SetFaceImageUrl(v string) *DescribeVerifyResultResponseMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetIdCardName(v string) *DescribeVerifyResultResponseMaterial {
	s.IdCardName = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetIdCardNumber(v string) *DescribeVerifyResultResponseMaterial {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetFaceQuality(v string) *DescribeVerifyResultResponseMaterial {
	s.FaceQuality = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetFaceGlobalUrl(v string) *DescribeVerifyResultResponseMaterial {
	s.FaceGlobalUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetFaceMask(v bool) *DescribeVerifyResultResponseMaterial {
	s.FaceMask = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetIdCardInfo(v *DescribeVerifyResultResponseMaterialIdCardInfo) *DescribeVerifyResultResponseMaterial {
	s.IdCardInfo = v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetVideoUrls(v []*string) *DescribeVerifyResultResponseMaterial {
	s.VideoUrls = v
	return s
}

type DescribeVerifyResultResponseMaterialIdCardInfo struct {
	Number        *string `json:"Number,omitempty" xml:"Number,omitempty" require:"true"`
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty" require:"true"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty" require:"true"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty" require:"true"`
	Authority     *string `json:"Authority,omitempty" xml:"Authority,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Birth         *string `json:"Birth,omitempty" xml:"Birth,omitempty" require:"true"`
	BackImageUrl  *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
}

func (s DescribeVerifyResultResponseMaterialIdCardInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetNumber(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetAddress(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetNationality(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetEndDate(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetFrontImageUrl(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetAuthority(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetName(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetBirth(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetBackImageUrl(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetStartDate(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

type DescribeOssUploadTokenRequest struct {
}

func (s DescribeOssUploadTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssUploadTokenRequest) GoString() string {
	return s.String()
}

type DescribeOssUploadTokenResponse struct {
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OssUploadToken *DescribeOssUploadTokenResponseOssUploadToken `json:"OssUploadToken,omitempty" xml:"OssUploadToken,omitempty" require:"true" type:"Struct"`
}

func (s DescribeOssUploadTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssUploadTokenResponse) SetRequestId(v string) *DescribeOssUploadTokenResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeOssUploadTokenResponse) SetOssUploadToken(v *DescribeOssUploadTokenResponseOssUploadToken) *DescribeOssUploadTokenResponse {
	s.OssUploadToken = v
	return s
}

type DescribeOssUploadTokenResponseOssUploadToken struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty" require:"true"`
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty" require:"true"`
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty" require:"true"`
	Expired  *int64  `json:"Expired,omitempty" xml:"Expired,omitempty" require:"true"`
	Secret   *string `json:"Secret,omitempty" xml:"Secret,omitempty" require:"true"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Token    *string `json:"Token,omitempty" xml:"Token,omitempty" require:"true"`
}

func (s DescribeOssUploadTokenResponseOssUploadToken) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssUploadTokenResponseOssUploadToken) GoString() string {
	return s.String()
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetBucket(v string) *DescribeOssUploadTokenResponseOssUploadToken {
	s.Bucket = &v
	return s
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetEndPoint(v string) *DescribeOssUploadTokenResponseOssUploadToken {
	s.EndPoint = &v
	return s
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetPath(v string) *DescribeOssUploadTokenResponseOssUploadToken {
	s.Path = &v
	return s
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetExpired(v int64) *DescribeOssUploadTokenResponseOssUploadToken {
	s.Expired = &v
	return s
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetSecret(v string) *DescribeOssUploadTokenResponseOssUploadToken {
	s.Secret = &v
	return s
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetKey(v string) *DescribeOssUploadTokenResponseOssUploadToken {
	s.Key = &v
	return s
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetToken(v string) *DescribeOssUploadTokenResponseOssUploadToken {
	s.Token = &v
	return s
}

type DescribeVerifyTokenRequest struct {
	IdCardBackImageUrl   *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	FailedRedirectUrl    *string `json:"FailedRedirectUrl,omitempty" xml:"FailedRedirectUrl,omitempty"`
	FaceRetainedImageUrl *string `json:"FaceRetainedImageUrl,omitempty" xml:"FaceRetainedImageUrl,omitempty"`
	CallbackSeed         *string `json:"CallbackSeed,omitempty" xml:"CallbackSeed,omitempty"`
	IdCardFrontImageUrl  *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BizId                *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	IdCardNumber         *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	PassedRedirectUrl    *string `json:"PassedRedirectUrl,omitempty" xml:"PassedRedirectUrl,omitempty"`
	CallbackUrl          *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
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

func (s *DescribeVerifyTokenRequest) SetIdCardBackImageUrl(v string) *DescribeVerifyTokenRequest {
	s.IdCardBackImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetBizType(v string) *DescribeVerifyTokenRequest {
	s.BizType = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetFailedRedirectUrl(v string) *DescribeVerifyTokenRequest {
	s.FailedRedirectUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetFaceRetainedImageUrl(v string) *DescribeVerifyTokenRequest {
	s.FaceRetainedImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetCallbackSeed(v string) *DescribeVerifyTokenRequest {
	s.CallbackSeed = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardFrontImageUrl(v string) *DescribeVerifyTokenRequest {
	s.IdCardFrontImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserId(v string) *DescribeVerifyTokenRequest {
	s.UserId = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetBizId(v string) *DescribeVerifyTokenRequest {
	s.BizId = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetName(v string) *DescribeVerifyTokenRequest {
	s.Name = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardNumber(v string) *DescribeVerifyTokenRequest {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetPassedRedirectUrl(v string) *DescribeVerifyTokenRequest {
	s.PassedRedirectUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetCallbackUrl(v string) *DescribeVerifyTokenRequest {
	s.CallbackUrl = &v
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

type DescribeVerifyTokenResponse struct {
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VerifyPageUrl  *string                                    `json:"VerifyPageUrl,omitempty" xml:"VerifyPageUrl,omitempty" require:"true"`
	VerifyToken    *string                                    `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty" require:"true"`
	OssUploadToken *DescribeVerifyTokenResponseOssUploadToken `json:"OssUploadToken,omitempty" xml:"OssUploadToken,omitempty" require:"true" type:"Struct"`
}

func (s DescribeVerifyTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponse) SetRequestId(v string) *DescribeVerifyTokenResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyTokenResponse) SetVerifyPageUrl(v string) *DescribeVerifyTokenResponse {
	s.VerifyPageUrl = &v
	return s
}

func (s *DescribeVerifyTokenResponse) SetVerifyToken(v string) *DescribeVerifyTokenResponse {
	s.VerifyToken = &v
	return s
}

func (s *DescribeVerifyTokenResponse) SetOssUploadToken(v *DescribeVerifyTokenResponseOssUploadToken) *DescribeVerifyTokenResponse {
	s.OssUploadToken = v
	return s
}

type DescribeVerifyTokenResponseOssUploadToken struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty" require:"true"`
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty" require:"true"`
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty" require:"true"`
	Expired  *int64  `json:"Expired,omitempty" xml:"Expired,omitempty" require:"true"`
	Secret   *string `json:"Secret,omitempty" xml:"Secret,omitempty" require:"true"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Token    *string `json:"Token,omitempty" xml:"Token,omitempty" require:"true"`
}

func (s DescribeVerifyTokenResponseOssUploadToken) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenResponseOssUploadToken) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetBucket(v string) *DescribeVerifyTokenResponseOssUploadToken {
	s.Bucket = &v
	return s
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetEndPoint(v string) *DescribeVerifyTokenResponseOssUploadToken {
	s.EndPoint = &v
	return s
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetPath(v string) *DescribeVerifyTokenResponseOssUploadToken {
	s.Path = &v
	return s
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetExpired(v int64) *DescribeVerifyTokenResponseOssUploadToken {
	s.Expired = &v
	return s
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetSecret(v string) *DescribeVerifyTokenResponseOssUploadToken {
	s.Secret = &v
	return s
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetKey(v string) *DescribeVerifyTokenResponseOssUploadToken {
	s.Key = &v
	return s
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetToken(v string) *DescribeVerifyTokenResponseOssUploadToken {
	s.Token = &v
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

func (client *Client) DescribeWhitelist(request *DescribeWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeWhitelistResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeWhitelist"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWhitelistSimply(request *DescribeWhitelistRequest) (_result *DescribeWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWhitelistResponse{}
	_body, _err := client.DescribeWhitelist(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWhitelist(request *DeleteWhitelistRequest, runtime *util.RuntimeOptions) (_result *DeleteWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteWhitelistResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteWhitelist"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWhitelistSimply(request *DeleteWhitelistRequest) (_result *DeleteWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWhitelistResponse{}
	_body, _err := client.DeleteWhitelist(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWhitelist(request *CreateWhitelistRequest, runtime *util.RuntimeOptions) (_result *CreateWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateWhitelistResponse{}
	_body, _err := client.DoRequest(tea.String("CreateWhitelist"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWhitelistSimply(request *CreateWhitelistRequest) (_result *CreateWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWhitelistResponse{}
	_body, _err := client.CreateWhitelist(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFaceConfig(request *DescribeFaceConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeFaceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeFaceConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeFaceConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFaceConfigSimply(request *DescribeFaceConfigRequest) (_result *DescribeFaceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaceConfigResponse{}
	_body, _err := client.DescribeFaceConfig(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFaceConfig(request *UpdateFaceConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateFaceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateFaceConfigResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateFaceConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFaceConfigSimply(request *UpdateFaceConfigRequest) (_result *UpdateFaceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFaceConfigResponse{}
	_body, _err := client.UpdateFaceConfig(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFaceConfig(request *CreateFaceConfigRequest, runtime *util.RuntimeOptions) (_result *CreateFaceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateFaceConfigResponse{}
	_body, _err := client.DoRequest(tea.String("CreateFaceConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFaceConfigSimply(request *CreateFaceConfigRequest) (_result *CreateFaceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFaceConfigResponse{}
	_body, _err := client.CreateFaceConfig(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LivenessFaceVerify(request *LivenessFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *LivenessFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &LivenessFaceVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("LivenessFaceVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LivenessFaceVerifySimply(request *LivenessFaceVerifyRequest) (_result *LivenessFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LivenessFaceVerifyResponse{}
	_body, _err := client.LivenessFaceVerify(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompareFaceVerify(request *CompareFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *CompareFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CompareFaceVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("CompareFaceVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompareFaceVerifySimply(request *CompareFaceVerifyRequest) (_result *CompareFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareFaceVerifyResponse{}
	_body, _err := client.CompareFaceVerify(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSdkUrl(request *DescribeSdkUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeSdkUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSdkUrlResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSdkUrl"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSdkUrlSimply(request *DescribeSdkUrlRequest) (_result *DescribeSdkUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSdkUrlResponse{}
	_body, _err := client.DescribeSdkUrl(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUpdatePackageResult(request *DescribeUpdatePackageResultRequest, runtime *util.RuntimeOptions) (_result *DescribeUpdatePackageResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUpdatePackageResultResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeUpdatePackageResult"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUpdatePackageResultSimply(request *DescribeUpdatePackageResultRequest) (_result *DescribeUpdatePackageResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUpdatePackageResultResponse{}
	_body, _err := client.DescribeUpdatePackageResult(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppPackage(request *UpdateAppPackageRequest, runtime *util.RuntimeOptions) (_result *UpdateAppPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateAppPackageResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateAppPackage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppPackageSimply(request *UpdateAppPackageRequest) (_result *UpdateAppPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppPackageResponse{}
	_body, _err := client.UpdateAppPackage(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppInfo(request *DescribeAppInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeAppInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAppInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeAppInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppInfoSimply(request *DescribeAppInfoRequest) (_result *DescribeAppInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppInfoResponse{}
	_body, _err := client.DescribeAppInfo(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContrastFaceVerify(request *ContrastFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *ContrastFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ContrastFaceVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("ContrastFaceVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ContrastFaceVerifySimply(request *ContrastFaceVerifyRequest) (_result *ContrastFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContrastFaceVerifyResponse{}
	_body, _err := client.ContrastFaceVerify(request, runtime)
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

	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
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
	rpcutil.Convert(runtime, ossRuntime)
	contrastFaceVerifyReq := &ContrastFaceVerifyRequest{}
	rpcutil.Convert(request, contrastFaceVerifyReq)
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
	contrastFaceVerifyResp, _err := client.ContrastFaceVerify(contrastFaceVerifyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = contrastFaceVerifyResp
	return _result, _err
}

func (client *Client) InitDevice(request *InitDeviceRequest, runtime *util.RuntimeOptions) (_result *InitDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InitDeviceResponse{}
	_body, _err := client.DoRequest(tea.String("InitDevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitDeviceSimply(request *InitDeviceRequest) (_result *InitDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitDeviceResponse{}
	_body, _err := client.InitDevice(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitFaceVerify(request *InitFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *InitFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("InitFaceVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitFaceVerifySimply(request *InitFaceVerifyRequest) (_result *InitFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.InitFaceVerify(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFaceVerify(request *DescribeFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *DescribeFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeFaceVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeFaceVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFaceVerifySimply(request *DescribeFaceVerifyRequest) (_result *DescribeFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaceVerifyResponse{}
	_body, _err := client.DescribeFaceVerify(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyDevice(request *VerifyDeviceRequest, runtime *util.RuntimeOptions) (_result *VerifyDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &VerifyDeviceResponse{}
	_body, _err := client.DoRequest(tea.String("VerifyDevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyDeviceSimply(request *VerifyDeviceRequest) (_result *VerifyDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyDeviceResponse{}
	_body, _err := client.VerifyDevice(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDeviceInfo(request *ModifyDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *ModifyDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDeviceInfoResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDeviceInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDeviceInfoSimply(request *ModifyDeviceInfoRequest) (_result *ModifyDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDeviceInfoResponse{}
	_body, _err := client.ModifyDeviceInfo(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifySDK(request *DescribeVerifySDKRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifySDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVerifySDKResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVerifySDK"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifySDKSimply(request *DescribeVerifySDKRequest) (_result *DescribeVerifySDKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifySDKResponse{}
	_body, _err := client.DescribeVerifySDK(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceInfo(request *DescribeDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDeviceInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceInfoSimply(request *DescribeDeviceInfoRequest) (_result *DescribeDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.DescribeDeviceInfo(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVerifySDK(request *CreateVerifySDKRequest, runtime *util.RuntimeOptions) (_result *CreateVerifySDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateVerifySDKResponse{}
	_body, _err := client.DoRequest(tea.String("CreateVerifySDK"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVerifySDKSimply(request *CreateVerifySDKRequest) (_result *CreateVerifySDKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVerifySDKResponse{}
	_body, _err := client.CreateVerifySDK(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAuthKey(request *CreateAuthKeyRequest, runtime *util.RuntimeOptions) (_result *CreateAuthKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateAuthKeyResponse{}
	_body, _err := client.DoRequest(tea.String("CreateAuthKey"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAuthKeySimply(request *CreateAuthKeyRequest) (_result *CreateAuthKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAuthKeyResponse{}
	_body, _err := client.CreateAuthKey(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectFaceAttributes(request *DetectFaceAttributesRequest, runtime *util.RuntimeOptions) (_result *DetectFaceAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DetectFaceAttributesResponse{}
	_body, _err := client.DoRequest(tea.String("DetectFaceAttributes"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectFaceAttributesSimply(request *DetectFaceAttributesRequest) (_result *DetectFaceAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectFaceAttributesResponse{}
	_body, _err := client.DetectFaceAttributes(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompareFaces(request *CompareFacesRequest, runtime *util.RuntimeOptions) (_result *CompareFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CompareFacesResponse{}
	_body, _err := client.DoRequest(tea.String("CompareFaces"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompareFacesSimply(request *CompareFacesRequest) (_result *CompareFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareFacesResponse{}
	_body, _err := client.CompareFaces(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFaceUsage(request *DescribeFaceUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeFaceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeFaceUsageResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeFaceUsage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFaceUsageSimply(request *DescribeFaceUsageRequest) (_result *DescribeFaceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaceUsageResponse{}
	_body, _err := client.DescribeFaceUsage(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifyRecords(request *DescribeVerifyRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifyRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVerifyRecordsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVerifyRecords"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifyRecordsSimply(request *DescribeVerifyRecordsRequest) (_result *DescribeVerifyRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifyRecordsResponse{}
	_body, _err := client.DescribeVerifyRecords(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVerifySetting(request *UpdateVerifySettingRequest, runtime *util.RuntimeOptions) (_result *UpdateVerifySettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateVerifySettingResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateVerifySetting"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVerifySettingSimply(request *UpdateVerifySettingRequest) (_result *UpdateVerifySettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVerifySettingResponse{}
	_body, _err := client.UpdateVerifySetting(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVerifySetting(request *CreateVerifySettingRequest, runtime *util.RuntimeOptions) (_result *CreateVerifySettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateVerifySettingResponse{}
	_body, _err := client.DoRequest(tea.String("CreateVerifySetting"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVerifySettingSimply(request *CreateVerifySettingRequest) (_result *CreateVerifySettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVerifySettingResponse{}
	_body, _err := client.CreateVerifySetting(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifySetting(request *DescribeVerifySettingRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifySettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVerifySettingResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVerifySetting"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifySettingSimply(request *DescribeVerifySettingRequest) (_result *DescribeVerifySettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifySettingResponse{}
	_body, _err := client.DescribeVerifySetting(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifyUsage(request *DescribeVerifyUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifyUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVerifyUsageResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVerifyUsage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifyUsageSimply(request *DescribeVerifyUsageRequest) (_result *DescribeVerifyUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifyUsageResponse{}
	_body, _err := client.DescribeVerifyUsage(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserStatus(request *DescribeUserStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeUserStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUserStatusResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeUserStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserStatusSimply(request *DescribeUserStatusRequest) (_result *DescribeUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserStatusResponse{}
	_body, _err := client.DescribeUserStatus(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUploadInfo(request *DescribeUploadInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeUploadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUploadInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeUploadInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUploadInfoSimply(request *DescribeUploadInfoRequest) (_result *DescribeUploadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUploadInfoResponse{}
	_body, _err := client.DescribeUploadInfo(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRPSDK(request *DescribeRPSDKRequest, runtime *util.RuntimeOptions) (_result *DescribeRPSDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRPSDKResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRPSDK"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRPSDKSimply(request *DescribeRPSDKRequest) (_result *DescribeRPSDKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRPSDKResponse{}
	_body, _err := client.DescribeRPSDK(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRPSDK(request *CreateRPSDKRequest, runtime *util.RuntimeOptions) (_result *CreateRPSDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateRPSDKResponse{}
	_body, _err := client.DoRequest(tea.String("CreateRPSDK"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRPSDKSimply(request *CreateRPSDKRequest) (_result *CreateRPSDKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRPSDKResponse{}
	_body, _err := client.CreateRPSDK(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyMaterial(request *VerifyMaterialRequest, runtime *util.RuntimeOptions) (_result *VerifyMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &VerifyMaterialResponse{}
	_body, _err := client.DoRequest(tea.String("VerifyMaterial"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyMaterialSimply(request *VerifyMaterialRequest) (_result *VerifyMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyMaterialResponse{}
	_body, _err := client.VerifyMaterial(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifyResult(request *DescribeVerifyResultRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVerifyResultResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVerifyResult"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifyResultSimply(request *DescribeVerifyResultRequest) (_result *DescribeVerifyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifyResultResponse{}
	_body, _err := client.DescribeVerifyResult(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssUploadToken(request *DescribeOssUploadTokenRequest, runtime *util.RuntimeOptions) (_result *DescribeOssUploadTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeOssUploadTokenResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeOssUploadToken"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssUploadTokenSimply(request *DescribeOssUploadTokenRequest) (_result *DescribeOssUploadTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssUploadTokenResponse{}
	_body, _err := client.DescribeOssUploadToken(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifyToken(request *DescribeVerifyTokenRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifyTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVerifyTokenResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVerifyToken"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifyTokenSimply(request *DescribeVerifyTokenRequest) (_result *DescribeVerifyTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifyTokenResponse{}
	_body, _err := client.DescribeVerifyToken(request, runtime)
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
