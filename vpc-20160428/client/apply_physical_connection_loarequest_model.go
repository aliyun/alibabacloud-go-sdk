// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyPhysicalConnectionLOARequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *ApplyPhysicalConnectionLOARequest
	GetBandwidth() *int32
	SetClientToken(v string) *ApplyPhysicalConnectionLOARequest
	GetClientToken() *string
	SetCompanyName(v string) *ApplyPhysicalConnectionLOARequest
	GetCompanyName() *string
	SetConstructionTime(v string) *ApplyPhysicalConnectionLOARequest
	GetConstructionTime() *string
	SetInstanceId(v string) *ApplyPhysicalConnectionLOARequest
	GetInstanceId() *string
	SetLineType(v string) *ApplyPhysicalConnectionLOARequest
	GetLineType() *string
	SetOwnerAccount(v string) *ApplyPhysicalConnectionLOARequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ApplyPhysicalConnectionLOARequest
	GetOwnerId() *int64
	SetPMInfo(v []*ApplyPhysicalConnectionLOARequestPMInfo) *ApplyPhysicalConnectionLOARequest
	GetPMInfo() []*ApplyPhysicalConnectionLOARequestPMInfo
	SetPeerLocation(v string) *ApplyPhysicalConnectionLOARequest
	GetPeerLocation() *string
	SetRegionId(v string) *ApplyPhysicalConnectionLOARequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ApplyPhysicalConnectionLOARequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ApplyPhysicalConnectionLOARequest
	GetResourceOwnerId() *int64
	SetSi(v string) *ApplyPhysicalConnectionLOARequest
	GetSi() *string
}

type ApplyPhysicalConnectionLOARequest struct {
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
	// Generate a unique value from your client to ensure that different requests have unique ClientToken values. ClientToken supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
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
	// - **MSTP**: MSTP line.
	//
	// - **MPLSVPN**: MPLSVPN line.
	//
	// - **FIBRE**: fiber optic direct connection.
	//
	// - **Other**: other type of line.
	//
	// This parameter is required.
	//
	// example:
	//
	// FIBRE
	LineType     *string `json:"LineType,omitempty" xml:"LineType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The information about the construction engineer.
	//
	// > This parameter is required. Specify the relevant information.
	PMInfo []*ApplyPhysicalConnectionLOARequestPMInfo `json:"PMInfo,omitempty" xml:"PMInfo,omitempty" type:"Repeated"`
	// The geographical location where the Express Connect circuit is deployed.
	//
	// example:
	//
	// 杭州
	PeerLocation *string `json:"PeerLocation,omitempty" xml:"PeerLocation,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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

func (s ApplyPhysicalConnectionLOARequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyPhysicalConnectionLOARequest) GoString() string {
	return s.String()
}

func (s *ApplyPhysicalConnectionLOARequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ApplyPhysicalConnectionLOARequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ApplyPhysicalConnectionLOARequest) GetCompanyName() *string {
	return s.CompanyName
}

func (s *ApplyPhysicalConnectionLOARequest) GetConstructionTime() *string {
	return s.ConstructionTime
}

func (s *ApplyPhysicalConnectionLOARequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ApplyPhysicalConnectionLOARequest) GetLineType() *string {
	return s.LineType
}

func (s *ApplyPhysicalConnectionLOARequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ApplyPhysicalConnectionLOARequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ApplyPhysicalConnectionLOARequest) GetPMInfo() []*ApplyPhysicalConnectionLOARequestPMInfo {
	return s.PMInfo
}

func (s *ApplyPhysicalConnectionLOARequest) GetPeerLocation() *string {
	return s.PeerLocation
}

func (s *ApplyPhysicalConnectionLOARequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApplyPhysicalConnectionLOARequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ApplyPhysicalConnectionLOARequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ApplyPhysicalConnectionLOARequest) GetSi() *string {
	return s.Si
}

func (s *ApplyPhysicalConnectionLOARequest) SetBandwidth(v int32) *ApplyPhysicalConnectionLOARequest {
	s.Bandwidth = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) SetClientToken(v string) *ApplyPhysicalConnectionLOARequest {
	s.ClientToken = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) SetCompanyName(v string) *ApplyPhysicalConnectionLOARequest {
	s.CompanyName = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) SetConstructionTime(v string) *ApplyPhysicalConnectionLOARequest {
	s.ConstructionTime = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) SetInstanceId(v string) *ApplyPhysicalConnectionLOARequest {
	s.InstanceId = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) SetLineType(v string) *ApplyPhysicalConnectionLOARequest {
	s.LineType = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) SetOwnerAccount(v string) *ApplyPhysicalConnectionLOARequest {
	s.OwnerAccount = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) SetOwnerId(v int64) *ApplyPhysicalConnectionLOARequest {
	s.OwnerId = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) SetPMInfo(v []*ApplyPhysicalConnectionLOARequestPMInfo) *ApplyPhysicalConnectionLOARequest {
	s.PMInfo = v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) SetPeerLocation(v string) *ApplyPhysicalConnectionLOARequest {
	s.PeerLocation = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) SetRegionId(v string) *ApplyPhysicalConnectionLOARequest {
	s.RegionId = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) SetResourceOwnerAccount(v string) *ApplyPhysicalConnectionLOARequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) SetResourceOwnerId(v int64) *ApplyPhysicalConnectionLOARequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) SetSi(v string) *ApplyPhysicalConnectionLOARequest {
	s.Si = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequest) Validate() error {
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

type ApplyPhysicalConnectionLOARequestPMInfo struct {
	// The ID number of the construction engineer. You can specify an ID card number or an international passport number.
	//
	// You can specify information about up to 16 construction engineers.
	//
	// example:
	//
	// 5****************9
	PMCertificateNo *string `json:"PMCertificateNo,omitempty" xml:"PMCertificateNo,omitempty"`
	// The type of the ID document of the construction engineer. Valid values:
	//
	// - **IDCard**: ID card.
	//
	// - **Passport**: international passport.
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

func (s ApplyPhysicalConnectionLOARequestPMInfo) String() string {
	return dara.Prettify(s)
}

func (s ApplyPhysicalConnectionLOARequestPMInfo) GoString() string {
	return s.String()
}

func (s *ApplyPhysicalConnectionLOARequestPMInfo) GetPMCertificateNo() *string {
	return s.PMCertificateNo
}

func (s *ApplyPhysicalConnectionLOARequestPMInfo) GetPMCertificateType() *string {
	return s.PMCertificateType
}

func (s *ApplyPhysicalConnectionLOARequestPMInfo) GetPMContactInfo() *string {
	return s.PMContactInfo
}

func (s *ApplyPhysicalConnectionLOARequestPMInfo) GetPMGender() *string {
	return s.PMGender
}

func (s *ApplyPhysicalConnectionLOARequestPMInfo) GetPMName() *string {
	return s.PMName
}

func (s *ApplyPhysicalConnectionLOARequestPMInfo) SetPMCertificateNo(v string) *ApplyPhysicalConnectionLOARequestPMInfo {
	s.PMCertificateNo = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequestPMInfo) SetPMCertificateType(v string) *ApplyPhysicalConnectionLOARequestPMInfo {
	s.PMCertificateType = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequestPMInfo) SetPMContactInfo(v string) *ApplyPhysicalConnectionLOARequestPMInfo {
	s.PMContactInfo = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequestPMInfo) SetPMGender(v string) *ApplyPhysicalConnectionLOARequestPMInfo {
	s.PMGender = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequestPMInfo) SetPMName(v string) *ApplyPhysicalConnectionLOARequestPMInfo {
	s.PMName = &v
	return s
}

func (s *ApplyPhysicalConnectionLOARequestPMInfo) Validate() error {
	return dara.Validate(s)
}
