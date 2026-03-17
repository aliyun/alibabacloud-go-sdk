// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagHaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *ModifySagHaRequest
	GetMode() *string
	SetOwnerAccount(v string) *ModifySagHaRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagHaRequest
	GetOwnerId() *int64
	SetPortName(v string) *ModifySagHaRequest
	GetPortName() *string
	SetRegionId(v string) *ModifySagHaRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySagHaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagHaRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ModifySagHaRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagHaRequest
	GetSmartAGSn() *string
	SetVirtualIp(v string) *ModifySagHaRequest
	GetVirtualIp() *string
}

type ModifySagHaRequest struct {
	// The HA mode. Valid values:
	//
	// 	- **NONE**: disables HA.
	//
	// 	- **STATIC**: enables static HA.
	//
	// 	- **DYNAMIC**: enables dynamic HA.
	//
	// This parameter is required.
	//
	// example:
	//
	// NONE
	Mode         *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port name.
	//
	// >  If Mode is set to STATIC, you must specify a port name.
	//
	// example:
	//
	// 5
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The region ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The SAG instance ID.
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
	// The virtual IP address.
	//
	// >  If Mode is set to STATIC, you must specify a virtual IP address.
	//
	// example:
	//
	// 192.XX.XX.5
	VirtualIp *string `json:"VirtualIp,omitempty" xml:"VirtualIp,omitempty"`
}

func (s ModifySagHaRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagHaRequest) GoString() string {
	return s.String()
}

func (s *ModifySagHaRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifySagHaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagHaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagHaRequest) GetPortName() *string {
	return s.PortName
}

func (s *ModifySagHaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagHaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagHaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagHaRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagHaRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagHaRequest) GetVirtualIp() *string {
	return s.VirtualIp
}

func (s *ModifySagHaRequest) SetMode(v string) *ModifySagHaRequest {
	s.Mode = &v
	return s
}

func (s *ModifySagHaRequest) SetOwnerAccount(v string) *ModifySagHaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagHaRequest) SetOwnerId(v int64) *ModifySagHaRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagHaRequest) SetPortName(v string) *ModifySagHaRequest {
	s.PortName = &v
	return s
}

func (s *ModifySagHaRequest) SetRegionId(v string) *ModifySagHaRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagHaRequest) SetResourceOwnerAccount(v string) *ModifySagHaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagHaRequest) SetResourceOwnerId(v int64) *ModifySagHaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagHaRequest) SetSmartAGId(v string) *ModifySagHaRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagHaRequest) SetSmartAGSn(v string) *ModifySagHaRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagHaRequest) SetVirtualIp(v string) *ModifySagHaRequest {
	s.VirtualIp = &v
	return s
}

func (s *ModifySagHaRequest) Validate() error {
	return dara.Validate(s)
}
