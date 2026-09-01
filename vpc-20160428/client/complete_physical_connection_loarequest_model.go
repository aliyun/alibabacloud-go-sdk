// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompletePhysicalConnectionLOARequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CompletePhysicalConnectionLOARequest
	GetClientToken() *string
	SetFinishWork(v bool) *CompletePhysicalConnectionLOARequest
	GetFinishWork() *bool
	SetInstanceId(v string) *CompletePhysicalConnectionLOARequest
	GetInstanceId() *string
	SetLineCode(v string) *CompletePhysicalConnectionLOARequest
	GetLineCode() *string
	SetLineLabel(v string) *CompletePhysicalConnectionLOARequest
	GetLineLabel() *string
	SetLineSPContactInfo(v string) *CompletePhysicalConnectionLOARequest
	GetLineSPContactInfo() *string
	SetLineServiceProvider(v string) *CompletePhysicalConnectionLOARequest
	GetLineServiceProvider() *string
	SetOwnerAccount(v string) *CompletePhysicalConnectionLOARequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CompletePhysicalConnectionLOARequest
	GetOwnerId() *int64
	SetRegionId(v string) *CompletePhysicalConnectionLOARequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CompletePhysicalConnectionLOARequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CompletePhysicalConnectionLOARequest
	GetResourceOwnerId() *int64
}

type CompletePhysicalConnectionLOARequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// The client generates the value. Ensure that the value is unique among different requests.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-230e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether the construction is completed. Valid values:
	//
	// 	- **true**: Construction is completed.
	//
	// 	- **false**: Line O&M.
	//
	// example:
	//
	// true
	FinishWork *bool `json:"FinishWork,omitempty" xml:"FinishWork,omitempty"`
	// The instance ID of the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp10tvlhnwkw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The line code of the carrier.
	//
	// example:
	//
	// aaa111****
	LineCode *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	// The cable label in the data center building.
	//
	// example:
	//
	// bbb222****
	LineLabel *string `json:"LineLabel,omitempty" xml:"LineLabel,omitempty"`
	// The O&M contact information of the line carrier.
	//
	// example:
	//
	// 1388888****
	LineSPContactInfo *string `json:"LineSPContactInfo,omitempty" xml:"LineSPContactInfo,omitempty"`
	// The carrier. Valid values:
	//
	// - **中国电信**.
	//
	// - **中国联通**.
	//
	// - **中国移动**.
	//
	// - **中国其他**.
	//
	// example:
	//
	// 中国其他
	LineServiceProvider *string `json:"LineServiceProvider,omitempty" xml:"LineServiceProvider,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CompletePhysicalConnectionLOARequest) String() string {
	return dara.Prettify(s)
}

func (s CompletePhysicalConnectionLOARequest) GoString() string {
	return s.String()
}

func (s *CompletePhysicalConnectionLOARequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CompletePhysicalConnectionLOARequest) GetFinishWork() *bool {
	return s.FinishWork
}

func (s *CompletePhysicalConnectionLOARequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CompletePhysicalConnectionLOARequest) GetLineCode() *string {
	return s.LineCode
}

func (s *CompletePhysicalConnectionLOARequest) GetLineLabel() *string {
	return s.LineLabel
}

func (s *CompletePhysicalConnectionLOARequest) GetLineSPContactInfo() *string {
	return s.LineSPContactInfo
}

func (s *CompletePhysicalConnectionLOARequest) GetLineServiceProvider() *string {
	return s.LineServiceProvider
}

func (s *CompletePhysicalConnectionLOARequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CompletePhysicalConnectionLOARequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CompletePhysicalConnectionLOARequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CompletePhysicalConnectionLOARequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CompletePhysicalConnectionLOARequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CompletePhysicalConnectionLOARequest) SetClientToken(v string) *CompletePhysicalConnectionLOARequest {
	s.ClientToken = &v
	return s
}

func (s *CompletePhysicalConnectionLOARequest) SetFinishWork(v bool) *CompletePhysicalConnectionLOARequest {
	s.FinishWork = &v
	return s
}

func (s *CompletePhysicalConnectionLOARequest) SetInstanceId(v string) *CompletePhysicalConnectionLOARequest {
	s.InstanceId = &v
	return s
}

func (s *CompletePhysicalConnectionLOARequest) SetLineCode(v string) *CompletePhysicalConnectionLOARequest {
	s.LineCode = &v
	return s
}

func (s *CompletePhysicalConnectionLOARequest) SetLineLabel(v string) *CompletePhysicalConnectionLOARequest {
	s.LineLabel = &v
	return s
}

func (s *CompletePhysicalConnectionLOARequest) SetLineSPContactInfo(v string) *CompletePhysicalConnectionLOARequest {
	s.LineSPContactInfo = &v
	return s
}

func (s *CompletePhysicalConnectionLOARequest) SetLineServiceProvider(v string) *CompletePhysicalConnectionLOARequest {
	s.LineServiceProvider = &v
	return s
}

func (s *CompletePhysicalConnectionLOARequest) SetOwnerAccount(v string) *CompletePhysicalConnectionLOARequest {
	s.OwnerAccount = &v
	return s
}

func (s *CompletePhysicalConnectionLOARequest) SetOwnerId(v int64) *CompletePhysicalConnectionLOARequest {
	s.OwnerId = &v
	return s
}

func (s *CompletePhysicalConnectionLOARequest) SetRegionId(v string) *CompletePhysicalConnectionLOARequest {
	s.RegionId = &v
	return s
}

func (s *CompletePhysicalConnectionLOARequest) SetResourceOwnerAccount(v string) *CompletePhysicalConnectionLOARequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CompletePhysicalConnectionLOARequest) SetResourceOwnerId(v int64) *CompletePhysicalConnectionLOARequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CompletePhysicalConnectionLOARequest) Validate() error {
	return dara.Validate(s)
}
