// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIAssetSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeAIAssetSummaryResponseBodyData) *DescribeAIAssetSummaryResponseBody
	GetData() *DescribeAIAssetSummaryResponseBodyData
	SetRequestId(v string) *DescribeAIAssetSummaryResponseBody
	GetRequestId() *string
}

type DescribeAIAssetSummaryResponseBody struct {
	// The returned data.
	Data *DescribeAIAssetSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAIAssetSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIAssetSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIAssetSummaryResponseBody) GetData() *DescribeAIAssetSummaryResponseBodyData {
	return s.Data
}

func (s *DescribeAIAssetSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIAssetSummaryResponseBody) SetData(v *DescribeAIAssetSummaryResponseBodyData) *DescribeAIAssetSummaryResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAIAssetSummaryResponseBody) SetRequestId(v string) *DescribeAIAssetSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAIAssetSummaryResponseBodyData struct {
	// The number of cloud assets with AI security posture management risks.
	//
	// example:
	//
	// 15
	AispmRiskAssetCnt *int32 `json:"AispmRiskAssetCnt,omitempty" xml:"AispmRiskAssetCnt,omitempty"`
	// The number of servers on which AI components are installed.
	//
	// example:
	//
	// 2
	EcsAssetCnt *int32 `json:"EcsAssetCnt,omitempty" xml:"EcsAssetCnt,omitempty"`
	// The number of servers that have exposed AI components.
	//
	// example:
	//
	// 0
	ExposedRiskAssetCnt *int32 `json:"ExposedRiskAssetCnt,omitempty" xml:"ExposedRiskAssetCnt,omitempty"`
	// The number of AI images.
	//
	// example:
	//
	// 2
	ImageAssetCnt *int32 `json:"ImageAssetCnt,omitempty" xml:"ImageAssetCnt,omitempty"`
	// The number of LINGJUN assets.
	//
	// example:
	//
	// 0
	LingjunAssetCnt *int32 `json:"LingjunAssetCnt,omitempty" xml:"LingjunAssetCnt,omitempty"`
	// The number of container image assets in PAI.
	//
	// example:
	//
	// 2
	PaiContainerCnt *int32 `json:"PaiContainerCnt,omitempty" xml:"PaiContainerCnt,omitempty"`
	// The total number of cloud asset instances in Platform for AI (PAI).
	//
	// example:
	//
	// 17
	PaiInstanceCnt *int32 `json:"PaiInstanceCnt,omitempty" xml:"PaiInstanceCnt,omitempty"`
	// The number of serverless assets in PAI.
	//
	// example:
	//
	// 0
	PaiServerlessAssetCnt *int32 `json:"PaiServerlessAssetCnt,omitempty" xml:"PaiServerlessAssetCnt,omitempty"`
	// The statistics on assets that have AI-related keys are stored in plaintext.
	SensitiveSummary *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary `json:"SensitiveSummary,omitempty" xml:"SensitiveSummary,omitempty" type:"Struct"`
	// The number of AI snapshots.
	//
	// example:
	//
	// 2
	SnapshotAssetCnt *int32 `json:"SnapshotAssetCnt,omitempty" xml:"SnapshotAssetCnt,omitempty"`
	// The total number of AI assets.
	//
	// example:
	//
	// 25
	TotalAssetCnt *int32 `json:"TotalAssetCnt,omitempty" xml:"TotalAssetCnt,omitempty"`
	// The total number of assets with AI risks.
	//
	// example:
	//
	// 26
	TotalRiskCnt *int32 `json:"TotalRiskCnt,omitempty" xml:"TotalRiskCnt,omitempty"`
	// The number of servers with AI application vulnerabilities.
	//
	// example:
	//
	// 2
	VulRiskAssetCnt *int32 `json:"VulRiskAssetCnt,omitempty" xml:"VulRiskAssetCnt,omitempty"`
}

func (s DescribeAIAssetSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIAssetSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAIAssetSummaryResponseBodyData) GetAispmRiskAssetCnt() *int32 {
	return s.AispmRiskAssetCnt
}

func (s *DescribeAIAssetSummaryResponseBodyData) GetEcsAssetCnt() *int32 {
	return s.EcsAssetCnt
}

func (s *DescribeAIAssetSummaryResponseBodyData) GetExposedRiskAssetCnt() *int32 {
	return s.ExposedRiskAssetCnt
}

func (s *DescribeAIAssetSummaryResponseBodyData) GetImageAssetCnt() *int32 {
	return s.ImageAssetCnt
}

func (s *DescribeAIAssetSummaryResponseBodyData) GetLingjunAssetCnt() *int32 {
	return s.LingjunAssetCnt
}

func (s *DescribeAIAssetSummaryResponseBodyData) GetPaiContainerCnt() *int32 {
	return s.PaiContainerCnt
}

func (s *DescribeAIAssetSummaryResponseBodyData) GetPaiInstanceCnt() *int32 {
	return s.PaiInstanceCnt
}

func (s *DescribeAIAssetSummaryResponseBodyData) GetPaiServerlessAssetCnt() *int32 {
	return s.PaiServerlessAssetCnt
}

func (s *DescribeAIAssetSummaryResponseBodyData) GetSensitiveSummary() *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary {
	return s.SensitiveSummary
}

func (s *DescribeAIAssetSummaryResponseBodyData) GetSnapshotAssetCnt() *int32 {
	return s.SnapshotAssetCnt
}

func (s *DescribeAIAssetSummaryResponseBodyData) GetTotalAssetCnt() *int32 {
	return s.TotalAssetCnt
}

func (s *DescribeAIAssetSummaryResponseBodyData) GetTotalRiskCnt() *int32 {
	return s.TotalRiskCnt
}

func (s *DescribeAIAssetSummaryResponseBodyData) GetVulRiskAssetCnt() *int32 {
	return s.VulRiskAssetCnt
}

func (s *DescribeAIAssetSummaryResponseBodyData) SetAispmRiskAssetCnt(v int32) *DescribeAIAssetSummaryResponseBodyData {
	s.AispmRiskAssetCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyData) SetEcsAssetCnt(v int32) *DescribeAIAssetSummaryResponseBodyData {
	s.EcsAssetCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyData) SetExposedRiskAssetCnt(v int32) *DescribeAIAssetSummaryResponseBodyData {
	s.ExposedRiskAssetCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyData) SetImageAssetCnt(v int32) *DescribeAIAssetSummaryResponseBodyData {
	s.ImageAssetCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyData) SetLingjunAssetCnt(v int32) *DescribeAIAssetSummaryResponseBodyData {
	s.LingjunAssetCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyData) SetPaiContainerCnt(v int32) *DescribeAIAssetSummaryResponseBodyData {
	s.PaiContainerCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyData) SetPaiInstanceCnt(v int32) *DescribeAIAssetSummaryResponseBodyData {
	s.PaiInstanceCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyData) SetPaiServerlessAssetCnt(v int32) *DescribeAIAssetSummaryResponseBodyData {
	s.PaiServerlessAssetCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyData) SetSensitiveSummary(v *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) *DescribeAIAssetSummaryResponseBodyData {
	s.SensitiveSummary = v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyData) SetSnapshotAssetCnt(v int32) *DescribeAIAssetSummaryResponseBodyData {
	s.SnapshotAssetCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyData) SetTotalAssetCnt(v int32) *DescribeAIAssetSummaryResponseBodyData {
	s.TotalAssetCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyData) SetTotalRiskCnt(v int32) *DescribeAIAssetSummaryResponseBodyData {
	s.TotalRiskCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyData) SetVulRiskAssetCnt(v int32) *DescribeAIAssetSummaryResponseBodyData {
	s.VulRiskAssetCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeAIAssetSummaryResponseBodyDataSensitiveSummary struct {
	// The number of images that have AI-related keys are stored in plaintext detected by image scan.
	//
	// example:
	//
	// 1
	ContainerImageCnt *int32 `json:"ContainerImageCnt,omitempty" xml:"ContainerImageCnt,omitempty"`
	// The number of servers that have AI-related keys are stored in plaintext detected by agentless scan.
	//
	// example:
	//
	// 1
	EcsCnt *int32 `json:"EcsCnt,omitempty" xml:"EcsCnt,omitempty"`
	// The number of images that have AI-related keys are stored in plaintext detected by agentless scan.
	//
	// example:
	//
	// 3
	ImageCnt *int32 `json:"ImageCnt,omitempty" xml:"ImageCnt,omitempty"`
	// The number of snapshots that have AI-related keys are stored in plaintext detected by agentless scan.
	//
	// example:
	//
	// 4
	SnapshotCnt *int32 `json:"SnapshotCnt,omitempty" xml:"SnapshotCnt,omitempty"`
	// The total number of assets that have AI-related keys are stored in plaintext.
	//
	// example:
	//
	// 9
	TotalCnt *int32 `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) GoString() string {
	return s.String()
}

func (s *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) GetContainerImageCnt() *int32 {
	return s.ContainerImageCnt
}

func (s *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) GetEcsCnt() *int32 {
	return s.EcsCnt
}

func (s *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) GetImageCnt() *int32 {
	return s.ImageCnt
}

func (s *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) GetSnapshotCnt() *int32 {
	return s.SnapshotCnt
}

func (s *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) GetTotalCnt() *int32 {
	return s.TotalCnt
}

func (s *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) SetContainerImageCnt(v int32) *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary {
	s.ContainerImageCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) SetEcsCnt(v int32) *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary {
	s.EcsCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) SetImageCnt(v int32) *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary {
	s.ImageCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) SetSnapshotCnt(v int32) *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary {
	s.SnapshotCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) SetTotalCnt(v int32) *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary {
	s.TotalCnt = &v
	return s
}

func (s *DescribeAIAssetSummaryResponseBodyDataSensitiveSummary) Validate() error {
	return dara.Validate(s)
}
