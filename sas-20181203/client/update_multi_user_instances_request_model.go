// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiUserInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberInstances(v []*UpdateMultiUserInstancesRequestMemberInstances) *UpdateMultiUserInstancesRequest
	GetMemberInstances() []*UpdateMultiUserInstancesRequestMemberInstances
}

type UpdateMultiUserInstancesRequest struct {
	// Member instances.
	MemberInstances []*UpdateMultiUserInstancesRequestMemberInstances `json:"MemberInstances,omitempty" xml:"MemberInstances,omitempty" type:"Repeated"`
}

func (s UpdateMultiUserInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiUserInstancesRequest) GoString() string {
	return s.String()
}

func (s *UpdateMultiUserInstancesRequest) GetMemberInstances() []*UpdateMultiUserInstancesRequestMemberInstances {
	return s.MemberInstances
}

func (s *UpdateMultiUserInstancesRequest) SetMemberInstances(v []*UpdateMultiUserInstancesRequestMemberInstances) *UpdateMultiUserInstancesRequest {
	s.MemberInstances = v
	return s
}

func (s *UpdateMultiUserInstancesRequest) Validate() error {
	if s.MemberInstances != nil {
		for _, item := range s.MemberInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMultiUserInstancesRequestMemberInstances struct {
	// The Alibaba Cloud account UID of the member.
	//
	// example:
	//
	// 1766185894104675
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// Anti-ransomware capacity allocated to the member, in GB.
	//
	// example:
	//
	// 10
	AntiRansomwareCapacity *int64 `json:"AntiRansomwareCapacity,omitempty" xml:"AntiRansomwareCapacity,omitempty"`
	// Charge type, values:
	//
	// 	- **PREPAID**: Prepaid.
	//
	// 	- **POSTPAID*	- (default): Postpaid.
	//
	// example:
	//
	// PREPAID
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Cloud platform configuration check scan count allocated to the member. Unit: times per month.
	//
	// example:
	//
	// 0
	CspmCapacity *int64 `json:"CspmCapacity,omitempty" xml:"CspmCapacity,omitempty"`
	// Honeypot authorization count allocated to the member.
	//
	// example:
	//
	// 0
	HoneypotCapacity *int64 `json:"HoneypotCapacity,omitempty" xml:"HoneypotCapacity,omitempty"`
	// Image scan authorization count allocated to the member.
	//
	// example:
	//
	// 1
	ImageScanCapacity *int64 `json:"ImageScanCapacity,omitempty" xml:"ImageScanCapacity,omitempty"`
	// The Cloud Security Center instance ID purchased by the member account.
	//
	// example:
	//
	// sas-p0anpb26my69
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Operation type. Values:
	//
	// - **ADD**: Add
	//
	// - **CHANGE**: Change
	//
	// - **DEL**: Delete
	//
	// example:
	//
	// CHANGE
	OptType *string `json:"OptType,omitempty" xml:"OptType,omitempty"`
	// Application protection count allocated to the member. Unit: per month.
	//
	// example:
	//
	// 0
	RaspCapacity *int64 `json:"RaspCapacity,omitempty" xml:"RaspCapacity,omitempty"`
	// Malicious file detection SDK authorization count allocated to the member.
	//
	// example:
	//
	// 10
	SdkCapacity *int64 `json:"SdkCapacity,omitempty" xml:"SdkCapacity,omitempty"`
	// Log storage capacity allocated to the member, in GB.
	//
	// example:
	//
	// 10
	SlsCapacity *int64 `json:"SlsCapacity,omitempty" xml:"SlsCapacity,omitempty"`
	// Status of the member account instance. Values:
	//
	// - **1**: Valid.
	//
	// - **2**: Invalid.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Threat analysis capacity allocated to the member. Unit: GB.
	//
	// example:
	//
	// 10
	ThreatAnalysisCapacity *int64 `json:"ThreatAnalysisCapacity,omitempty" xml:"ThreatAnalysisCapacity,omitempty"`
	// Threat analysis and response log access traffic allocated to the member. Unit: GB/day.
	//
	// example:
	//
	// 0
	ThreatAnalysisFlow *int64 `json:"ThreatAnalysisFlow,omitempty" xml:"ThreatAnalysisFlow,omitempty"`
	// The version of Cloud Security Center protection to be bound. Values:
	//
	// - **1**: Free Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Antivirus Edition
	//
	// - **7**: Flagship Edition
	//
	// example:
	//
	// 7
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// Member account authorization usage information.
	VersionSummary []*UpdateMultiUserInstancesRequestMemberInstancesVersionSummary `json:"VersionSummary,omitempty" xml:"VersionSummary,omitempty" type:"Repeated"`
	// Web tamper-proof authorization count allocated to the member.
	//
	// example:
	//
	// 0
	WebLockCapacity *int64 `json:"WebLockCapacity,omitempty" xml:"WebLockCapacity,omitempty"`
}

func (s UpdateMultiUserInstancesRequestMemberInstances) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiUserInstancesRequestMemberInstances) GoString() string {
	return s.String()
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetAliUid() *int64 {
	return s.AliUid
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetAntiRansomwareCapacity() *int64 {
	return s.AntiRansomwareCapacity
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetChargeType() *string {
	return s.ChargeType
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetCspmCapacity() *int64 {
	return s.CspmCapacity
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetHoneypotCapacity() *int64 {
	return s.HoneypotCapacity
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetImageScanCapacity() *int64 {
	return s.ImageScanCapacity
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetOptType() *string {
	return s.OptType
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetRaspCapacity() *int64 {
	return s.RaspCapacity
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetSdkCapacity() *int64 {
	return s.SdkCapacity
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetSlsCapacity() *int64 {
	return s.SlsCapacity
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetThreatAnalysisCapacity() *int64 {
	return s.ThreatAnalysisCapacity
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetThreatAnalysisFlow() *int64 {
	return s.ThreatAnalysisFlow
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetVersion() *string {
	return s.Version
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetVersionSummary() []*UpdateMultiUserInstancesRequestMemberInstancesVersionSummary {
	return s.VersionSummary
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) GetWebLockCapacity() *int64 {
	return s.WebLockCapacity
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetAliUid(v int64) *UpdateMultiUserInstancesRequestMemberInstances {
	s.AliUid = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetAntiRansomwareCapacity(v int64) *UpdateMultiUserInstancesRequestMemberInstances {
	s.AntiRansomwareCapacity = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetChargeType(v string) *UpdateMultiUserInstancesRequestMemberInstances {
	s.ChargeType = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetCspmCapacity(v int64) *UpdateMultiUserInstancesRequestMemberInstances {
	s.CspmCapacity = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetHoneypotCapacity(v int64) *UpdateMultiUserInstancesRequestMemberInstances {
	s.HoneypotCapacity = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetImageScanCapacity(v int64) *UpdateMultiUserInstancesRequestMemberInstances {
	s.ImageScanCapacity = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetInstanceId(v string) *UpdateMultiUserInstancesRequestMemberInstances {
	s.InstanceId = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetOptType(v string) *UpdateMultiUserInstancesRequestMemberInstances {
	s.OptType = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetRaspCapacity(v int64) *UpdateMultiUserInstancesRequestMemberInstances {
	s.RaspCapacity = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetSdkCapacity(v int64) *UpdateMultiUserInstancesRequestMemberInstances {
	s.SdkCapacity = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetSlsCapacity(v int64) *UpdateMultiUserInstancesRequestMemberInstances {
	s.SlsCapacity = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetStatus(v int32) *UpdateMultiUserInstancesRequestMemberInstances {
	s.Status = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetThreatAnalysisCapacity(v int64) *UpdateMultiUserInstancesRequestMemberInstances {
	s.ThreatAnalysisCapacity = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetThreatAnalysisFlow(v int64) *UpdateMultiUserInstancesRequestMemberInstances {
	s.ThreatAnalysisFlow = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetVersion(v string) *UpdateMultiUserInstancesRequestMemberInstances {
	s.Version = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetVersionSummary(v []*UpdateMultiUserInstancesRequestMemberInstancesVersionSummary) *UpdateMultiUserInstancesRequestMemberInstances {
	s.VersionSummary = v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) SetWebLockCapacity(v int64) *UpdateMultiUserInstancesRequestMemberInstances {
	s.WebLockCapacity = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstances) Validate() error {
	if s.VersionSummary != nil {
		for _, item := range s.VersionSummary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMultiUserInstancesRequestMemberInstancesVersionSummary struct {
	// Number of cores authorized for the member.
	//
	// example:
	//
	// 6
	CoreCount *int64 `json:"CoreCount,omitempty" xml:"CoreCount,omitempty"`
	// Number of authorizations allocated to the member.
	//
	// example:
	//
	// 3
	EcsCount *int64 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	// Version of the Cloud Security Center for the member account. Values:
	//
	// - **1**: Free Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Antivirus Edition
	//
	// - **7**: Flagship Edition
	//
	// - **8**: Multiple Versions
	//
	// - **10**: Only Purchase Value-Added Services
	//
	// example:
	//
	// 5
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateMultiUserInstancesRequestMemberInstancesVersionSummary) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiUserInstancesRequestMemberInstancesVersionSummary) GoString() string {
	return s.String()
}

func (s *UpdateMultiUserInstancesRequestMemberInstancesVersionSummary) GetCoreCount() *int64 {
	return s.CoreCount
}

func (s *UpdateMultiUserInstancesRequestMemberInstancesVersionSummary) GetEcsCount() *int64 {
	return s.EcsCount
}

func (s *UpdateMultiUserInstancesRequestMemberInstancesVersionSummary) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateMultiUserInstancesRequestMemberInstancesVersionSummary) SetCoreCount(v int64) *UpdateMultiUserInstancesRequestMemberInstancesVersionSummary {
	s.CoreCount = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstancesVersionSummary) SetEcsCount(v int64) *UpdateMultiUserInstancesRequestMemberInstancesVersionSummary {
	s.EcsCount = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstancesVersionSummary) SetVersion(v int32) *UpdateMultiUserInstancesRequestMemberInstancesVersionSummary {
	s.Version = &v
	return s
}

func (s *UpdateMultiUserInstancesRequestMemberInstancesVersionSummary) Validate() error {
	return dara.Validate(s)
}
