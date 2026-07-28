// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSecondApplyPhysicalConnectionLOARequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *SecondApplyPhysicalConnectionLOARequest
	GetBandwidth() *int32
	SetClientToken(v string) *SecondApplyPhysicalConnectionLOARequest
	GetClientToken() *string
	SetCompanyName(v string) *SecondApplyPhysicalConnectionLOARequest
	GetCompanyName() *string
	SetConstructionTime(v string) *SecondApplyPhysicalConnectionLOARequest
	GetConstructionTime() *string
	SetInstanceId(v string) *SecondApplyPhysicalConnectionLOARequest
	GetInstanceId() *string
	SetLineType(v string) *SecondApplyPhysicalConnectionLOARequest
	GetLineType() *string
	SetOwnerAccount(v string) *SecondApplyPhysicalConnectionLOARequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SecondApplyPhysicalConnectionLOARequest
	GetOwnerId() *int64
	SetPMInfo(v []*SecondApplyPhysicalConnectionLOARequestPMInfo) *SecondApplyPhysicalConnectionLOARequest
	GetPMInfo() []*SecondApplyPhysicalConnectionLOARequestPMInfo
	SetPeerLocation(v string) *SecondApplyPhysicalConnectionLOARequest
	GetPeerLocation() *string
	SetRegionId(v string) *SecondApplyPhysicalConnectionLOARequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SecondApplyPhysicalConnectionLOARequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SecondApplyPhysicalConnectionLOARequest
	GetResourceOwnerId() *int64
	SetSi(v string) *SecondApplyPhysicalConnectionLOARequest
	GetSi() *string
}

type SecondApplyPhysicalConnectionLOARequest struct {
	// The bandwidth of the Express Connect circuit. Unit: Mbit/s.
	//
	// Valid values: **2*	- to **10240**.
	//
	// example:
	//
	// 3
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// The client generates the value. Ensure that the value is unique among different requests.
	//
	// > If you do not specify this parameter, the system uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the company that accesses the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// company
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// The time when the construction company enters the site. The time is in the ISO 8601 standard and must be in UTC. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-02-28T16:00:00Z
	ConstructionTime *string `json:"ConstructionTime,omitempty" xml:"ConstructionTime,omitempty"`
	// The instance ID of the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1qrb3044eqi****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the Express Connect circuit. Valid values:
	//
	// - **MSTP**
	//
	// - **MPLSVPN**
	//
	// - **FIBRE**
	//
	// - **Other**
	//
	// This parameter is required.
	//
	// example:
	//
	// FIBRE
	LineType     *string `json:"LineType,omitempty" xml:"LineType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The information about the construction engineers.
	PMInfo []*SecondApplyPhysicalConnectionLOARequestPMInfo `json:"PMInfo,omitempty" xml:"PMInfo,omitempty" type:"Repeated"`
	// The geographical location where the Express Connect circuit is deployed.
	//
	// example:
	//
	// 杭州
	PeerLocation *string `json:"PeerLocation,omitempty" xml:"PeerLocation,omitempty"`
	// The region where the Express Connect circuit is deployed. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The construction company of the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里
	Si *string `json:"Si,omitempty" xml:"Si,omitempty"`
}

func (s SecondApplyPhysicalConnectionLOARequest) String() string {
	return dara.Prettify(s)
}

func (s SecondApplyPhysicalConnectionLOARequest) GoString() string {
	return s.String()
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetCompanyName() *string {
	return s.CompanyName
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetConstructionTime() *string {
	return s.ConstructionTime
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetLineType() *string {
	return s.LineType
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetPMInfo() []*SecondApplyPhysicalConnectionLOARequestPMInfo {
	return s.PMInfo
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetPeerLocation() *string {
	return s.PeerLocation
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SecondApplyPhysicalConnectionLOARequest) GetSi() *string {
	return s.Si
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetBandwidth(v int32) *SecondApplyPhysicalConnectionLOARequest {
	s.Bandwidth = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetClientToken(v string) *SecondApplyPhysicalConnectionLOARequest {
	s.ClientToken = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetCompanyName(v string) *SecondApplyPhysicalConnectionLOARequest {
	s.CompanyName = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetConstructionTime(v string) *SecondApplyPhysicalConnectionLOARequest {
	s.ConstructionTime = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetInstanceId(v string) *SecondApplyPhysicalConnectionLOARequest {
	s.InstanceId = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetLineType(v string) *SecondApplyPhysicalConnectionLOARequest {
	s.LineType = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetOwnerAccount(v string) *SecondApplyPhysicalConnectionLOARequest {
	s.OwnerAccount = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetOwnerId(v int64) *SecondApplyPhysicalConnectionLOARequest {
	s.OwnerId = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetPMInfo(v []*SecondApplyPhysicalConnectionLOARequestPMInfo) *SecondApplyPhysicalConnectionLOARequest {
	s.PMInfo = v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetPeerLocation(v string) *SecondApplyPhysicalConnectionLOARequest {
	s.PeerLocation = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetRegionId(v string) *SecondApplyPhysicalConnectionLOARequest {
	s.RegionId = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetResourceOwnerAccount(v string) *SecondApplyPhysicalConnectionLOARequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetResourceOwnerId(v int64) *SecondApplyPhysicalConnectionLOARequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) SetSi(v string) *SecondApplyPhysicalConnectionLOARequest {
	s.Si = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequest) Validate() error {
	if s.PMInfo != nil {
		for _, item := range s.PMInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SecondApplyPhysicalConnectionLOARequestPMInfo struct {
	// The ID number of the construction engineer. You can specify an ID card number or a passport number.
	//
	// You can configure information for up to 16 construction engineers.
	//
	// example:
	//
	// 5****************9
	PMCertificateNo *string `json:"PMCertificateNo,omitempty" xml:"PMCertificateNo,omitempty"`
	// The type of the ID document of the construction engineer. Valid values:
	//
	// - **IDCard**: ID card.
	//
	// - **Passport**: passport.
	//
	// example:
	//
	// IDCard
	PMCertificateType *string `json:"PMCertificateType,omitempty" xml:"PMCertificateType,omitempty"`
	// The contact information of the construction engineer.
	//
	// example:
	//
	// 1390000****
	PMContactInfo *string `json:"PMContactInfo,omitempty" xml:"PMContactInfo,omitempty"`
	// The gender of the construction engineer.
	//
	// example:
	//
	// Male
	PMGender *string `json:"PMGender,omitempty" xml:"PMGender,omitempty"`
	// The name of the construction engineer.
	//
	// example:
	//
	// 张三
	PMName *string `json:"PMName,omitempty" xml:"PMName,omitempty"`
}

func (s SecondApplyPhysicalConnectionLOARequestPMInfo) String() string {
	return dara.Prettify(s)
}

func (s SecondApplyPhysicalConnectionLOARequestPMInfo) GoString() string {
	return s.String()
}

func (s *SecondApplyPhysicalConnectionLOARequestPMInfo) GetPMCertificateNo() *string {
	return s.PMCertificateNo
}

func (s *SecondApplyPhysicalConnectionLOARequestPMInfo) GetPMCertificateType() *string {
	return s.PMCertificateType
}

func (s *SecondApplyPhysicalConnectionLOARequestPMInfo) GetPMContactInfo() *string {
	return s.PMContactInfo
}

func (s *SecondApplyPhysicalConnectionLOARequestPMInfo) GetPMGender() *string {
	return s.PMGender
}

func (s *SecondApplyPhysicalConnectionLOARequestPMInfo) GetPMName() *string {
	return s.PMName
}

func (s *SecondApplyPhysicalConnectionLOARequestPMInfo) SetPMCertificateNo(v string) *SecondApplyPhysicalConnectionLOARequestPMInfo {
	s.PMCertificateNo = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequestPMInfo) SetPMCertificateType(v string) *SecondApplyPhysicalConnectionLOARequestPMInfo {
	s.PMCertificateType = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequestPMInfo) SetPMContactInfo(v string) *SecondApplyPhysicalConnectionLOARequestPMInfo {
	s.PMContactInfo = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequestPMInfo) SetPMGender(v string) *SecondApplyPhysicalConnectionLOARequestPMInfo {
	s.PMGender = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequestPMInfo) SetPMName(v string) *SecondApplyPhysicalConnectionLOARequestPMInfo {
	s.PMName = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOARequestPMInfo) Validate() error {
	return dara.Validate(s)
}
