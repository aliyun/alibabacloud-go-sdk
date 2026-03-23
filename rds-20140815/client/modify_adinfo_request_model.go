// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyADInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetADAccountName(v string) *ModifyADInfoRequest
	GetADAccountName() *string
	SetADDNS(v string) *ModifyADInfoRequest
	GetADDNS() *string
	SetADPassword(v string) *ModifyADInfoRequest
	GetADPassword() *string
	SetADServerIpAddress(v string) *ModifyADInfoRequest
	GetADServerIpAddress() *string
	SetClientToken(v string) *ModifyADInfoRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyADInfoRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *ModifyADInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyADInfoRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyADInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyADInfoRequest
	GetResourceOwnerId() *int64
}

type ModifyADInfoRequest struct {
	ADAccountName     *string `json:"ADAccountName,omitempty" xml:"ADAccountName,omitempty"`
	ADDNS             *string `json:"ADDNS,omitempty" xml:"ADDNS,omitempty"`
	ADPassword        *string `json:"ADPassword,omitempty" xml:"ADPassword,omitempty"`
	ADServerIpAddress *string `json:"ADServerIpAddress,omitempty" xml:"ADServerIpAddress,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyADInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyADInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyADInfoRequest) GetADAccountName() *string {
	return s.ADAccountName
}

func (s *ModifyADInfoRequest) GetADDNS() *string {
	return s.ADDNS
}

func (s *ModifyADInfoRequest) GetADPassword() *string {
	return s.ADPassword
}

func (s *ModifyADInfoRequest) GetADServerIpAddress() *string {
	return s.ADServerIpAddress
}

func (s *ModifyADInfoRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyADInfoRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyADInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyADInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyADInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyADInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyADInfoRequest) SetADAccountName(v string) *ModifyADInfoRequest {
	s.ADAccountName = &v
	return s
}

func (s *ModifyADInfoRequest) SetADDNS(v string) *ModifyADInfoRequest {
	s.ADDNS = &v
	return s
}

func (s *ModifyADInfoRequest) SetADPassword(v string) *ModifyADInfoRequest {
	s.ADPassword = &v
	return s
}

func (s *ModifyADInfoRequest) SetADServerIpAddress(v string) *ModifyADInfoRequest {
	s.ADServerIpAddress = &v
	return s
}

func (s *ModifyADInfoRequest) SetClientToken(v string) *ModifyADInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyADInfoRequest) SetDBInstanceId(v string) *ModifyADInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyADInfoRequest) SetOwnerId(v int64) *ModifyADInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyADInfoRequest) SetRegionId(v string) *ModifyADInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyADInfoRequest) SetResourceOwnerAccount(v string) *ModifyADInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyADInfoRequest) SetResourceOwnerId(v int64) *ModifyADInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyADInfoRequest) Validate() error {
	return dara.Validate(s)
}
