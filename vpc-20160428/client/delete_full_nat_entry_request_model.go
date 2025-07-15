// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFullNatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteFullNatEntryRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteFullNatEntryRequest
	GetDryRun() *bool
	SetFullNatEntryId(v string) *DeleteFullNatEntryRequest
	GetFullNatEntryId() *string
	SetFullNatTableId(v string) *DeleteFullNatEntryRequest
	GetFullNatTableId() *string
	SetOwnerAccount(v string) *DeleteFullNatEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteFullNatEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteFullNatEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteFullNatEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteFullNatEntryRequest
	GetResourceOwnerId() *int64
}

type DeleteFullNatEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, the related error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the FULLNAT entry that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// fullnat-gw8fz23jezpbblf1j****
	FullNatEntryId *string `json:"FullNatEntryId,omitempty" xml:"FullNatEntryId,omitempty"`
	// The ID of the FULLNAT table to which the FULLNAT entry to be deleted belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// fulltb-gw88z7hhlv43rmb26****
	FullNatTableId *string `json:"FullNatTableId,omitempty" xml:"FullNatTableId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VPC NAT gateway to which the FULLNAT entry to be deleted belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent list of regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-central-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteFullNatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFullNatEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteFullNatEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteFullNatEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteFullNatEntryRequest) GetFullNatEntryId() *string {
	return s.FullNatEntryId
}

func (s *DeleteFullNatEntryRequest) GetFullNatTableId() *string {
	return s.FullNatTableId
}

func (s *DeleteFullNatEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteFullNatEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteFullNatEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteFullNatEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteFullNatEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteFullNatEntryRequest) SetClientToken(v string) *DeleteFullNatEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFullNatEntryRequest) SetDryRun(v bool) *DeleteFullNatEntryRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteFullNatEntryRequest) SetFullNatEntryId(v string) *DeleteFullNatEntryRequest {
	s.FullNatEntryId = &v
	return s
}

func (s *DeleteFullNatEntryRequest) SetFullNatTableId(v string) *DeleteFullNatEntryRequest {
	s.FullNatTableId = &v
	return s
}

func (s *DeleteFullNatEntryRequest) SetOwnerAccount(v string) *DeleteFullNatEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteFullNatEntryRequest) SetOwnerId(v int64) *DeleteFullNatEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteFullNatEntryRequest) SetRegionId(v string) *DeleteFullNatEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFullNatEntryRequest) SetResourceOwnerAccount(v string) *DeleteFullNatEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteFullNatEntryRequest) SetResourceOwnerId(v int64) *DeleteFullNatEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteFullNatEntryRequest) Validate() error {
	return dara.Validate(s)
}
