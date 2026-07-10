// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifySearchPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeVerifySearchPageListResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeVerifySearchPageListResponseBodyItems) *DescribeVerifySearchPageListResponseBody
	GetItems() []*DescribeVerifySearchPageListResponseBodyItems
	SetPageSize(v int32) *DescribeVerifySearchPageListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVerifySearchPageListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVerifySearchPageListResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *DescribeVerifySearchPageListResponseBody
	GetTotalPage() *int32
}

type DescribeVerifySearchPageListResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The paginated list data.
	Items []*DescribeVerifySearchPageListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 47D87BC1-D956-573A-8A15-A9007A76F56C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 53
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 4
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeVerifySearchPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifySearchPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifySearchPageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVerifySearchPageListResponseBody) GetItems() []*DescribeVerifySearchPageListResponseBodyItems {
	return s.Items
}

func (s *DescribeVerifySearchPageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVerifySearchPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifySearchPageListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVerifySearchPageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeVerifySearchPageListResponseBody) SetCurrentPage(v int32) *DescribeVerifySearchPageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBody) SetItems(v []*DescribeVerifySearchPageListResponseBodyItems) *DescribeVerifySearchPageListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeVerifySearchPageListResponseBody) SetPageSize(v int32) *DescribeVerifySearchPageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBody) SetRequestId(v string) *DescribeVerifySearchPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBody) SetTotalCount(v int32) *DescribeVerifySearchPageListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBody) SetTotalPage(v int32) *DescribeVerifySearchPageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVerifySearchPageListResponseBodyItems struct {
	// The desensitized ID card number.
	//
	// example:
	//
	// 3****************2
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// The certification ID.
	//
	// example:
	//
	// shad861465f2aaeeb805b519e1a93ab2
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// The extended information.
	ExtInfo *DescribeVerifySearchPageListResponseBodyItemsExtInfo `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty" type:"Struct"`
	// The verification time of this authentication record.
	//
	// example:
	//
	// 2025-10-14 15:40:13
	GmtVerify *string `json:"GmtVerify,omitempty" xml:"GmtVerify,omitempty"`
	// The liveness detection scheme.
	//
	// example:
	//
	// MULTI_ACTION
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The unique identifier for the customer request.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// Specifies whether the authentication passed. Valid values:
	//
	// - **T**: Passed.
	//
	// - **F**: Failed.
	//
	// example:
	//
	// T
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// The product code.
	//
	// example:
	//
	// ID_PRO
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The business scenario risk:
	//
	// - **0**: No risk.
	//
	// - **1**: Risk detected.
	//
	// example:
	//
	// 1
	RiskBizScenario *int32 `json:"RiskBizScenario,omitempty" xml:"RiskBizScenario,omitempty"`
	// The device risk:
	//
	// - **0**: No risk.
	//
	// - **1**: Risk detected.
	//
	// example:
	//
	// 1
	RiskDevice *int32 `json:"RiskDevice,omitempty" xml:"RiskDevice,omitempty"`
	// The DeviceToken risk:
	//
	// - **0**: No risk.
	//
	// - **1**: Risk detected.
	//
	// example:
	//
	// 0
	RiskDeviceToken *int32 `json:"RiskDeviceToken,omitempty" xml:"RiskDeviceToken,omitempty"`
	// The generic risk:
	//
	// - **0**: No risk.
	//
	// - **1**: Risk detected.
	//
	// example:
	//
	// 1
	RiskGeneric *int32 `json:"RiskGeneric,omitempty" xml:"RiskGeneric,omitempty"`
	// The large model mining risk:
	//
	// - **0**: No risk.
	//
	// - **1**: Risk detected.
	//
	// example:
	//
	// 1
	RiskModelMining *int32 `json:"RiskModelMining,omitempty" xml:"RiskModelMining,omitempty"`
	// Specifies whether the device is rooted. Set to 1 if selected; otherwise, do not pass this parameter. This parameter corresponds to the identity tag risk type.
	//
	// example:
	//
	// 1
	Root *int32 `json:"Root,omitempty" xml:"Root,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 1000015352
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// Specifies whether the device is a simulator. Set to 1 if selected; otherwise, do not pass this parameter. This parameter corresponds to the device tag risk type.
	//
	// example:
	//
	// 1
	Simulator *int32 `json:"Simulator,omitempty" xml:"Simulator,omitempty"`
	// The error code returned by the system.
	//
	// example:
	//
	// 207
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 198123xxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Specifies whether virtual video is used. Set to 1 if selected; otherwise, do not pass this parameter. This parameter corresponds to the behavior tag risk type.
	//
	// example:
	//
	// 1
	VirtualVideo *int32 `json:"VirtualVideo,omitempty" xml:"VirtualVideo,omitempty"`
}

func (s DescribeVerifySearchPageListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifySearchPageListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetCertNo() *string {
	return s.CertNo
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetCertifyId() *string {
	return s.CertifyId
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetExtInfo() *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	return s.ExtInfo
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetGmtVerify() *string {
	return s.GmtVerify
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetModel() *string {
	return s.Model
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetOuterOrderNo() *string {
	return s.OuterOrderNo
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetPassed() *string {
	return s.Passed
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetRiskBizScenario() *int32 {
	return s.RiskBizScenario
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetRiskDevice() *int32 {
	return s.RiskDevice
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetRiskDeviceToken() *int32 {
	return s.RiskDeviceToken
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetRiskGeneric() *int32 {
	return s.RiskGeneric
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetRiskModelMining() *int32 {
	return s.RiskModelMining
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetRoot() *int32 {
	return s.Root
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetSimulator() *int32 {
	return s.Simulator
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetSubCode() *string {
	return s.SubCode
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetUserId() *string {
	return s.UserId
}

func (s *DescribeVerifySearchPageListResponseBodyItems) GetVirtualVideo() *int32 {
	return s.VirtualVideo
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetCertNo(v string) *DescribeVerifySearchPageListResponseBodyItems {
	s.CertNo = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetCertifyId(v string) *DescribeVerifySearchPageListResponseBodyItems {
	s.CertifyId = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetExtInfo(v *DescribeVerifySearchPageListResponseBodyItemsExtInfo) *DescribeVerifySearchPageListResponseBodyItems {
	s.ExtInfo = v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetGmtVerify(v string) *DescribeVerifySearchPageListResponseBodyItems {
	s.GmtVerify = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetModel(v string) *DescribeVerifySearchPageListResponseBodyItems {
	s.Model = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetOuterOrderNo(v string) *DescribeVerifySearchPageListResponseBodyItems {
	s.OuterOrderNo = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetPassed(v string) *DescribeVerifySearchPageListResponseBodyItems {
	s.Passed = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetProductCode(v string) *DescribeVerifySearchPageListResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetRiskBizScenario(v int32) *DescribeVerifySearchPageListResponseBodyItems {
	s.RiskBizScenario = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetRiskDevice(v int32) *DescribeVerifySearchPageListResponseBodyItems {
	s.RiskDevice = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetRiskDeviceToken(v int32) *DescribeVerifySearchPageListResponseBodyItems {
	s.RiskDeviceToken = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetRiskGeneric(v int32) *DescribeVerifySearchPageListResponseBodyItems {
	s.RiskGeneric = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetRiskModelMining(v int32) *DescribeVerifySearchPageListResponseBodyItems {
	s.RiskModelMining = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetRoot(v int32) *DescribeVerifySearchPageListResponseBodyItems {
	s.Root = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetSceneId(v int64) *DescribeVerifySearchPageListResponseBodyItems {
	s.SceneId = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetSimulator(v int32) *DescribeVerifySearchPageListResponseBodyItems {
	s.Simulator = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetSubCode(v string) *DescribeVerifySearchPageListResponseBodyItems {
	s.SubCode = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetUserId(v string) *DescribeVerifySearchPageListResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) SetVirtualVideo(v int32) *DescribeVerifySearchPageListResponseBodyItems {
	s.VirtualVideo = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItems) Validate() error {
	if s.ExtInfo != nil {
		if err := s.ExtInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVerifySearchPageListResponseBodyItemsExtInfo struct {
	// The desensitized name.
	//
	// example:
	//
	// 何*
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The Face Guard tags.
	//
	// example:
	//
	// HOOK,ROOT
	DeviceRisk *string `json:"DeviceRisk,omitempty" xml:"DeviceRisk,omitempty"`
	// Specifies whether a face attack is detected:
	//
	// - **T**: Yes.
	//
	// - **F**: No.
	//
	// example:
	//
	// T
	FaceAttack *string `json:"FaceAttack,omitempty" xml:"FaceAttack,omitempty"`
	// The face attack score. The value ranges from 0 to 1. A value closer to 1 indicates a higher likelihood of an attack.
	//
	// example:
	//
	// 0.0000445161
	FaceAttackScore *float32 `json:"FaceAttackScore,omitempty" xml:"FaceAttackScore,omitempty"`
	// Specifies whether the face is occluded. T indicates occlusion detected. F indicates no occlusion.
	//
	// example:
	//
	// T
	FaceOcclusion *string `json:"FaceOcclusion,omitempty" xml:"FaceOcclusion,omitempty"`
	// The face-to-ID card comparison score.
	//
	// example:
	//
	// 0.9
	IdCardVerifyScore *float32 `json:"IdCardVerifyScore,omitempty" xml:"IdCardVerifyScore,omitempty"`
	// The OSS bucket for photos.
	//
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-XXX
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The file name of the OCR ID card face image.
	//
	// example:
	//
	// -
	OssIdFaceObjectName *string `json:"OssIdFaceObjectName,omitempty" xml:"OssIdFaceObjectName,omitempty"`
	// The file name of the OCR ID card national emblem image.
	//
	// example:
	//
	// -
	OssIdNationalEmblemObjectName *string `json:"OssIdNationalEmblemObjectName,omitempty" xml:"OssIdNationalEmblemObjectName,omitempty"`
	// The storage object name.
	//
	// example:
	//
	// verify/XXXXX1251634779/sha6a0a0cab01288c7aa8ac3f45220eb_0_normal.jpeg
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// The liveness face quality score.
	//
	// example:
	//
	// 1.0
	QualityScore *float32 `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
	// The face comparison score.
	//
	// example:
	//
	// 0.0
	VerifyScore *float32 `json:"VerifyScore,omitempty" xml:"VerifyScore,omitempty"`
	// The list of ASR texts.
	AsrTexts []*string `json:"asrTexts,omitempty" xml:"asrTexts,omitempty" type:"Repeated"`
	// The list of screen recording file OSS object names.
	//
	// example:
	//
	// -
	ScreenVideoObjectNames []*string `json:"screenVideoObjectNames,omitempty" xml:"screenVideoObjectNames,omitempty" type:"Repeated"`
	// The list of audio file OSS object names.
	//
	// example:
	//
	// -
	VoiceObjectNames []*string `json:"voiceObjectNames,omitempty" xml:"voiceObjectNames,omitempty" type:"Repeated"`
}

func (s DescribeVerifySearchPageListResponseBodyItemsExtInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifySearchPageListResponseBodyItemsExtInfo) GoString() string {
	return s.String()
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetDeviceRisk() *string {
	return s.DeviceRisk
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetFaceAttack() *string {
	return s.FaceAttack
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetFaceAttackScore() *float32 {
	return s.FaceAttackScore
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetFaceOcclusion() *string {
	return s.FaceOcclusion
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetIdCardVerifyScore() *float32 {
	return s.IdCardVerifyScore
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetOssIdFaceObjectName() *string {
	return s.OssIdFaceObjectName
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetOssIdNationalEmblemObjectName() *string {
	return s.OssIdNationalEmblemObjectName
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetQualityScore() *float32 {
	return s.QualityScore
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetVerifyScore() *float32 {
	return s.VerifyScore
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetAsrTexts() []*string {
	return s.AsrTexts
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetScreenVideoObjectNames() []*string {
	return s.ScreenVideoObjectNames
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) GetVoiceObjectNames() []*string {
	return s.VoiceObjectNames
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetCertName(v string) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.CertName = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetDeviceRisk(v string) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.DeviceRisk = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetFaceAttack(v string) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.FaceAttack = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetFaceAttackScore(v float32) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.FaceAttackScore = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetFaceOcclusion(v string) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.FaceOcclusion = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetIdCardVerifyScore(v float32) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.IdCardVerifyScore = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetOssBucketName(v string) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.OssBucketName = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetOssIdFaceObjectName(v string) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.OssIdFaceObjectName = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetOssIdNationalEmblemObjectName(v string) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.OssIdNationalEmblemObjectName = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetOssObjectName(v string) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.OssObjectName = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetQualityScore(v float32) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.QualityScore = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetVerifyScore(v float32) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.VerifyScore = &v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetAsrTexts(v []*string) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.AsrTexts = v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetScreenVideoObjectNames(v []*string) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.ScreenVideoObjectNames = v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) SetVoiceObjectNames(v []*string) *DescribeVerifySearchPageListResponseBodyItemsExtInfo {
	s.VoiceObjectNames = v
	return s
}

func (s *DescribeVerifySearchPageListResponseBodyItemsExtInfo) Validate() error {
	return dara.Validate(s)
}
