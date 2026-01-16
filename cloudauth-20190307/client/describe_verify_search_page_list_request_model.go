// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifySearchPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertNo(v string) *DescribeVerifySearchPageListRequest
	GetCertNo() *string
	SetCertifyId(v string) *DescribeVerifySearchPageListRequest
	GetCertifyId() *string
	SetCurrentPage(v string) *DescribeVerifySearchPageListRequest
	GetCurrentPage() *string
	SetEndDate(v string) *DescribeVerifySearchPageListRequest
	GetEndDate() *string
	SetHasDeviceRisk(v bool) *DescribeVerifySearchPageListRequest
	GetHasDeviceRisk() *bool
	SetModel(v string) *DescribeVerifySearchPageListRequest
	GetModel() *string
	SetOuterOrderNo(v string) *DescribeVerifySearchPageListRequest
	GetOuterOrderNo() *string
	SetPageSize(v string) *DescribeVerifySearchPageListRequest
	GetPageSize() *string
	SetPassed(v string) *DescribeVerifySearchPageListRequest
	GetPassed() *string
	SetProductCode(v string) *DescribeVerifySearchPageListRequest
	GetProductCode() *string
	SetRiskBizScenario(v int32) *DescribeVerifySearchPageListRequest
	GetRiskBizScenario() *int32
	SetRiskDevice(v int32) *DescribeVerifySearchPageListRequest
	GetRiskDevice() *int32
	SetRiskDeviceToken(v int32) *DescribeVerifySearchPageListRequest
	GetRiskDeviceToken() *int32
	SetRiskGeneric(v int32) *DescribeVerifySearchPageListRequest
	GetRiskGeneric() *int32
	SetRiskModelMining(v int32) *DescribeVerifySearchPageListRequest
	GetRiskModelMining() *int32
	SetRoot(v int32) *DescribeVerifySearchPageListRequest
	GetRoot() *int32
	SetSceneId(v string) *DescribeVerifySearchPageListRequest
	GetSceneId() *string
	SetSimulator(v int32) *DescribeVerifySearchPageListRequest
	GetSimulator() *int32
	SetStartDate(v string) *DescribeVerifySearchPageListRequest
	GetStartDate() *string
	SetSubCode(v string) *DescribeVerifySearchPageListRequest
	GetSubCode() *string
	SetSubCodes(v string) *DescribeVerifySearchPageListRequest
	GetSubCodes() *string
	SetVirtualVideo(v int32) *DescribeVerifySearchPageListRequest
	GetVirtualVideo() *int32
}

type DescribeVerifySearchPageListRequest struct {
	// ID number.
	//
	// example:
	//
	// 3203212000XXXX701X
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// Authentication ID.
	//
	// example:
	//
	// shadbdd3dbacd001cfa892a5e2b98dxx
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// Current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// End date of the query. The format is a Unix timestamp, in milliseconds.
	//
	// example:
	//
	// 2025-10-16 23:59:59 +0800
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Whether there is device risk (pass true if root = 1 or simulator = 1 or virtual_video = 1).
	//
	// example:
	//
	// true
	HasDeviceRisk *bool `json:"HasDeviceRisk,omitempty" xml:"HasDeviceRisk,omitempty"`
	// Liveness detection model.
	//
	// example:
	//
	// LIVENESS
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Unique identifier for the customer request.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// Number of items per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Whether the authentication passed:
	//
	// - **T**: Passed
	//
	// - **F**: Not passed
	//
	// example:
	//
	// F
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// Product code.
	//
	// example:
	//
	// ID_PRO
	ProductCode     *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RiskBizScenario *int32  `json:"RiskBizScenario,omitempty" xml:"RiskBizScenario,omitempty"`
	RiskDevice      *int32  `json:"RiskDevice,omitempty" xml:"RiskDevice,omitempty"`
	RiskDeviceToken *int32  `json:"RiskDeviceToken,omitempty" xml:"RiskDeviceToken,omitempty"`
	RiskGeneric     *int32  `json:"RiskGeneric,omitempty" xml:"RiskGeneric,omitempty"`
	RiskModelMining *int32  `json:"RiskModelMining,omitempty" xml:"RiskModelMining,omitempty"`
	// Whether it is rooted (pass 1 if selected, otherwise do not pass; corresponds to identity label risk type).
	//
	// example:
	//
	// 1
	Root *int32 `json:"Root,omitempty" xml:"Root,omitempty"`
	// Scene ID.
	//
	// example:
	//
	// 10000072xx
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// Whether it is a simulator (pass 1 if selected, otherwise do not pass; corresponds to device label risk type).
	//
	// example:
	//
	// 1
	Simulator *int32 `json:"Simulator,omitempty" xml:"Simulator,omitempty"`
	// Start date of the query.
	//
	// example:
	//
	// 2025-10-10 00:00:00 +0800
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Result Code. For detailed values, please refer to: [SubCode Explanation](https://help.aliyun.com/zh/id-verification/financial-grade-id-verification/error-code-person-verify?spm=a2c4g.11186623.0.0.6015566ebArcFw#d88910e172fgg).
	//
	// example:
	//
	// 201
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// Comma-separated Result Codes. For detailed values, please refer to: [SubCode Explanation](https://help.aliyun.com/zh/id-verification/financial-grade-id-verification/error-code-person-verify?spm=a2c4g.11186623.0.0.6015566ebArcFw#d88910e172fgg).
	//
	// example:
	//
	// 201,202
	SubCodes *string `json:"SubCodes,omitempty" xml:"SubCodes,omitempty"`
	// Whether it is a virtual adaptation (pass 1 if selected, otherwise do not pass; corresponds to behavior label risk type).
	//
	// example:
	//
	// 1
	VirtualVideo *int32 `json:"VirtualVideo,omitempty" xml:"VirtualVideo,omitempty"`
}

func (s DescribeVerifySearchPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifySearchPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifySearchPageListRequest) GetCertNo() *string {
	return s.CertNo
}

func (s *DescribeVerifySearchPageListRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *DescribeVerifySearchPageListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeVerifySearchPageListRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeVerifySearchPageListRequest) GetHasDeviceRisk() *bool {
	return s.HasDeviceRisk
}

func (s *DescribeVerifySearchPageListRequest) GetModel() *string {
	return s.Model
}

func (s *DescribeVerifySearchPageListRequest) GetOuterOrderNo() *string {
	return s.OuterOrderNo
}

func (s *DescribeVerifySearchPageListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVerifySearchPageListRequest) GetPassed() *string {
	return s.Passed
}

func (s *DescribeVerifySearchPageListRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeVerifySearchPageListRequest) GetRiskBizScenario() *int32 {
	return s.RiskBizScenario
}

func (s *DescribeVerifySearchPageListRequest) GetRiskDevice() *int32 {
	return s.RiskDevice
}

func (s *DescribeVerifySearchPageListRequest) GetRiskDeviceToken() *int32 {
	return s.RiskDeviceToken
}

func (s *DescribeVerifySearchPageListRequest) GetRiskGeneric() *int32 {
	return s.RiskGeneric
}

func (s *DescribeVerifySearchPageListRequest) GetRiskModelMining() *int32 {
	return s.RiskModelMining
}

func (s *DescribeVerifySearchPageListRequest) GetRoot() *int32 {
	return s.Root
}

func (s *DescribeVerifySearchPageListRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DescribeVerifySearchPageListRequest) GetSimulator() *int32 {
	return s.Simulator
}

func (s *DescribeVerifySearchPageListRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeVerifySearchPageListRequest) GetSubCode() *string {
	return s.SubCode
}

func (s *DescribeVerifySearchPageListRequest) GetSubCodes() *string {
	return s.SubCodes
}

func (s *DescribeVerifySearchPageListRequest) GetVirtualVideo() *int32 {
	return s.VirtualVideo
}

func (s *DescribeVerifySearchPageListRequest) SetCertNo(v string) *DescribeVerifySearchPageListRequest {
	s.CertNo = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetCertifyId(v string) *DescribeVerifySearchPageListRequest {
	s.CertifyId = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetCurrentPage(v string) *DescribeVerifySearchPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetEndDate(v string) *DescribeVerifySearchPageListRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetHasDeviceRisk(v bool) *DescribeVerifySearchPageListRequest {
	s.HasDeviceRisk = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetModel(v string) *DescribeVerifySearchPageListRequest {
	s.Model = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetOuterOrderNo(v string) *DescribeVerifySearchPageListRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetPageSize(v string) *DescribeVerifySearchPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetPassed(v string) *DescribeVerifySearchPageListRequest {
	s.Passed = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetProductCode(v string) *DescribeVerifySearchPageListRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetRiskBizScenario(v int32) *DescribeVerifySearchPageListRequest {
	s.RiskBizScenario = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetRiskDevice(v int32) *DescribeVerifySearchPageListRequest {
	s.RiskDevice = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetRiskDeviceToken(v int32) *DescribeVerifySearchPageListRequest {
	s.RiskDeviceToken = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetRiskGeneric(v int32) *DescribeVerifySearchPageListRequest {
	s.RiskGeneric = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetRiskModelMining(v int32) *DescribeVerifySearchPageListRequest {
	s.RiskModelMining = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetRoot(v int32) *DescribeVerifySearchPageListRequest {
	s.Root = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetSceneId(v string) *DescribeVerifySearchPageListRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetSimulator(v int32) *DescribeVerifySearchPageListRequest {
	s.Simulator = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetStartDate(v string) *DescribeVerifySearchPageListRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetSubCode(v string) *DescribeVerifySearchPageListRequest {
	s.SubCode = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetSubCodes(v string) *DescribeVerifySearchPageListRequest {
	s.SubCodes = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) SetVirtualVideo(v int32) *DescribeVerifySearchPageListRequest {
	s.VirtualVideo = &v
	return s
}

func (s *DescribeVerifySearchPageListRequest) Validate() error {
	return dara.Validate(s)
}
