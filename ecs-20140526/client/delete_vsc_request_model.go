// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVscRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVscRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteVscRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteVscRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVscRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVscRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVscRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVscRequest
	GetResourceOwnerId() *int64
	SetVscId(v string) *DeleteVscRequest
	GetVscId() *string
}

type DeleteVscRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-1**3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values: true: sends a check request without querying resource status. The check items include whether your AccessKey pair is valid, whether Resource Access Management (RAM) user authorization is complete, and whether the required parameters are specified. If the check fails, the corresponding error is returned. If the check succeeds, the DryRunOperation error code is returned. false: sends a Normal request. After the check succeeds, a 2xx HTTP status code is returned and the resource status is queried. Default value: false.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VSC that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsc-bp1j8y**etwq1ow3jal
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DeleteVscRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVscRequest) GoString() string {
	return s.String()
}

func (s *DeleteVscRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVscRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteVscRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVscRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVscRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVscRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVscRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVscRequest) GetVscId() *string {
	return s.VscId
}

func (s *DeleteVscRequest) SetClientToken(v string) *DeleteVscRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVscRequest) SetDryRun(v bool) *DeleteVscRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteVscRequest) SetOwnerAccount(v string) *DeleteVscRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVscRequest) SetOwnerId(v int64) *DeleteVscRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVscRequest) SetRegionId(v string) *DeleteVscRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVscRequest) SetResourceOwnerAccount(v string) *DeleteVscRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVscRequest) SetResourceOwnerId(v int64) *DeleteVscRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVscRequest) SetVscId(v string) *DeleteVscRequest {
	s.VscId = &v
	return s
}

func (s *DeleteVscRequest) Validate() error {
	return dara.Validate(s)
}
