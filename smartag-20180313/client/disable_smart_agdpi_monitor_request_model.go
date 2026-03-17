// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSmartAGDpiMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableSmartAGDpiMonitorRequest
	GetClientToken() *string
	SetDryRun(v bool) *DisableSmartAGDpiMonitorRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DisableSmartAGDpiMonitorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisableSmartAGDpiMonitorRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DisableSmartAGDpiMonitorRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DisableSmartAGDpiMonitorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisableSmartAGDpiMonitorRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DisableSmartAGDpiMonitorRequest
	GetSmartAGId() *string
}

type DisableSmartAGDpiMonitorRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 02fb3da4-130****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the validity of the request without actually making the request. Valid values:
	//
	// 	- **true**: checks the validity of the request but does not enable or disables the DPI feature for the SAG instance. Check items include the request format, instance status, and whether the required parameters are set. If the request fails the check, an error message is returned. If the request passes the request, the ID of the request is returned.
	//
	// 	- **false**: makes the request and disables the DPI feature for the SAG instance after the request passes the check. This is the default value.
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
	// sag-vwmylqc9521p5l****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DisableSmartAGDpiMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSmartAGDpiMonitorRequest) GoString() string {
	return s.String()
}

func (s *DisableSmartAGDpiMonitorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableSmartAGDpiMonitorRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DisableSmartAGDpiMonitorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisableSmartAGDpiMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableSmartAGDpiMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableSmartAGDpiMonitorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisableSmartAGDpiMonitorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisableSmartAGDpiMonitorRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DisableSmartAGDpiMonitorRequest) SetClientToken(v string) *DisableSmartAGDpiMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableSmartAGDpiMonitorRequest) SetDryRun(v bool) *DisableSmartAGDpiMonitorRequest {
	s.DryRun = &v
	return s
}

func (s *DisableSmartAGDpiMonitorRequest) SetOwnerAccount(v string) *DisableSmartAGDpiMonitorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableSmartAGDpiMonitorRequest) SetOwnerId(v int64) *DisableSmartAGDpiMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableSmartAGDpiMonitorRequest) SetRegionId(v string) *DisableSmartAGDpiMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DisableSmartAGDpiMonitorRequest) SetResourceOwnerAccount(v string) *DisableSmartAGDpiMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableSmartAGDpiMonitorRequest) SetResourceOwnerId(v int64) *DisableSmartAGDpiMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableSmartAGDpiMonitorRequest) SetSmartAGId(v string) *DisableSmartAGDpiMonitorRequest {
	s.SmartAGId = &v
	return s
}

func (s *DisableSmartAGDpiMonitorRequest) Validate() error {
	return dara.Validate(s)
}
