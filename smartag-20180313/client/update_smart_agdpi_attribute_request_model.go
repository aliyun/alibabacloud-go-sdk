// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAGDpiAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateSmartAGDpiAttributeRequest
	GetClientToken() *string
	SetDpiEnabled(v bool) *UpdateSmartAGDpiAttributeRequest
	GetDpiEnabled() *bool
	SetDryRun(v bool) *UpdateSmartAGDpiAttributeRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *UpdateSmartAGDpiAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateSmartAGDpiAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateSmartAGDpiAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateSmartAGDpiAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateSmartAGDpiAttributeRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *UpdateSmartAGDpiAttributeRequest
	GetSmartAGId() *string
}

type UpdateSmartAGDpiAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 02fb3da4-130e****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable the DPI feature for the SAG instance. Valid values:
	//
	// 	- **true**: enables DPI
	//
	// 	- **false**: disables DPI
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DpiEnabled *bool `json:"DpiEnabled,omitempty" xml:"DpiEnabled,omitempty"`
	// Specifies whether to check the validity of the request without actually making the request. Valid values:
	//
	// 	- **true**: checks the validity of the request but does not enable or disable the DPI feature. Check items include the request format, instance status, and whether the required parameters are set. If the request fails the check, an error message is returned. If the request passes the check, the ID of the request is returned.
	//
	// 	- **false**: makes the request and enables or disables the DPI feature after the request passes the check. This is the default value.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-tq3sazs17smldn****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s UpdateSmartAGDpiAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAGDpiAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmartAGDpiAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateSmartAGDpiAttributeRequest) GetDpiEnabled() *bool {
	return s.DpiEnabled
}

func (s *UpdateSmartAGDpiAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateSmartAGDpiAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateSmartAGDpiAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateSmartAGDpiAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSmartAGDpiAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateSmartAGDpiAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateSmartAGDpiAttributeRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *UpdateSmartAGDpiAttributeRequest) SetClientToken(v string) *UpdateSmartAGDpiAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSmartAGDpiAttributeRequest) SetDpiEnabled(v bool) *UpdateSmartAGDpiAttributeRequest {
	s.DpiEnabled = &v
	return s
}

func (s *UpdateSmartAGDpiAttributeRequest) SetDryRun(v bool) *UpdateSmartAGDpiAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateSmartAGDpiAttributeRequest) SetOwnerAccount(v string) *UpdateSmartAGDpiAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateSmartAGDpiAttributeRequest) SetOwnerId(v int64) *UpdateSmartAGDpiAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmartAGDpiAttributeRequest) SetRegionId(v string) *UpdateSmartAGDpiAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSmartAGDpiAttributeRequest) SetResourceOwnerAccount(v string) *UpdateSmartAGDpiAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmartAGDpiAttributeRequest) SetResourceOwnerId(v int64) *UpdateSmartAGDpiAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmartAGDpiAttributeRequest) SetSmartAGId(v string) *UpdateSmartAGDpiAttributeRequest {
	s.SmartAGId = &v
	return s
}

func (s *UpdateSmartAGDpiAttributeRequest) Validate() error {
	return dara.Validate(s)
}
