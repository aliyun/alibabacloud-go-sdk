// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagUserDnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMasterDns(v string) *ModifySagUserDnsRequest
	GetMasterDns() *string
	SetOwnerAccount(v string) *ModifySagUserDnsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagUserDnsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySagUserDnsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySagUserDnsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagUserDnsRequest
	GetResourceOwnerId() *int64
	SetSlaveDns(v string) *ModifySagUserDnsRequest
	GetSlaveDns() *string
	SetSmartAGId(v string) *ModifySagUserDnsRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagUserDnsRequest
	GetSmartAGSn() *string
}

type ModifySagUserDnsRequest struct {
	// The IP address of the primary DNS server.
	//
	// >  If you do not set this parameter, the IP address of the existing DNS server will be cleared.
	//
	// example:
	//
	// 192.XX.XX.1
	MasterDns    *string `json:"MasterDns,omitempty" xml:"MasterDns,omitempty"`
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
	// The IP address of the secondary DNS server.
	//
	// >  If you do not set this parameter, the IP address of the existing DNS server will be cleared.
	//
	// example:
	//
	// 192.XX.XX.2
	SlaveDns *string `json:"SlaveDns,omitempty" xml:"SlaveDns,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-whfn****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number of the SAG device associated with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s ModifySagUserDnsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagUserDnsRequest) GoString() string {
	return s.String()
}

func (s *ModifySagUserDnsRequest) GetMasterDns() *string {
	return s.MasterDns
}

func (s *ModifySagUserDnsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagUserDnsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagUserDnsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagUserDnsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagUserDnsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagUserDnsRequest) GetSlaveDns() *string {
	return s.SlaveDns
}

func (s *ModifySagUserDnsRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagUserDnsRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagUserDnsRequest) SetMasterDns(v string) *ModifySagUserDnsRequest {
	s.MasterDns = &v
	return s
}

func (s *ModifySagUserDnsRequest) SetOwnerAccount(v string) *ModifySagUserDnsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagUserDnsRequest) SetOwnerId(v int64) *ModifySagUserDnsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagUserDnsRequest) SetRegionId(v string) *ModifySagUserDnsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagUserDnsRequest) SetResourceOwnerAccount(v string) *ModifySagUserDnsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagUserDnsRequest) SetResourceOwnerId(v int64) *ModifySagUserDnsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagUserDnsRequest) SetSlaveDns(v string) *ModifySagUserDnsRequest {
	s.SlaveDns = &v
	return s
}

func (s *ModifySagUserDnsRequest) SetSmartAGId(v string) *ModifySagUserDnsRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagUserDnsRequest) SetSmartAGSn(v string) *ModifySagUserDnsRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagUserDnsRequest) Validate() error {
	return dara.Validate(s)
}
