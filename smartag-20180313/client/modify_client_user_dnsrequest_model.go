// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClientUserDNSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppDNS(v []*string) *ModifyClientUserDNSRequest
	GetAppDNS() []*string
	SetOwnerAccount(v string) *ModifyClientUserDNSRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyClientUserDNSRequest
	GetOwnerId() *int64
	SetRecoveredDNS(v []*string) *ModifyClientUserDNSRequest
	GetRecoveredDNS() []*string
	SetRegionId(v string) *ModifyClientUserDNSRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyClientUserDNSRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyClientUserDNSRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ModifyClientUserDNSRequest
	GetSmartAGId() *string
}

type ModifyClientUserDNSRequest struct {
	// example:
	//
	// 100.XX.XX.100
	AppDNS       []*string `json:"AppDNS,omitempty" xml:"AppDNS,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 100.XX.XX.110
	RecoveredDNS []*string `json:"RecoveredDNS,omitempty" xml:"RecoveredDNS,omitempty" type:"Repeated"`
	// The ID of the region where the SAG app instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-9uyg53s6juhpxv****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s ModifyClientUserDNSRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClientUserDNSRequest) GoString() string {
	return s.String()
}

func (s *ModifyClientUserDNSRequest) GetAppDNS() []*string {
	return s.AppDNS
}

func (s *ModifyClientUserDNSRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyClientUserDNSRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyClientUserDNSRequest) GetRecoveredDNS() []*string {
	return s.RecoveredDNS
}

func (s *ModifyClientUserDNSRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyClientUserDNSRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyClientUserDNSRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyClientUserDNSRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifyClientUserDNSRequest) SetAppDNS(v []*string) *ModifyClientUserDNSRequest {
	s.AppDNS = v
	return s
}

func (s *ModifyClientUserDNSRequest) SetOwnerAccount(v string) *ModifyClientUserDNSRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyClientUserDNSRequest) SetOwnerId(v int64) *ModifyClientUserDNSRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyClientUserDNSRequest) SetRecoveredDNS(v []*string) *ModifyClientUserDNSRequest {
	s.RecoveredDNS = v
	return s
}

func (s *ModifyClientUserDNSRequest) SetRegionId(v string) *ModifyClientUserDNSRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyClientUserDNSRequest) SetResourceOwnerAccount(v string) *ModifyClientUserDNSRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyClientUserDNSRequest) SetResourceOwnerId(v int64) *ModifyClientUserDNSRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyClientUserDNSRequest) SetSmartAGId(v string) *ModifyClientUserDNSRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifyClientUserDNSRequest) Validate() error {
	return dara.Validate(s)
}
