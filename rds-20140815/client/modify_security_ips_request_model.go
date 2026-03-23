// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceIPArrayAttribute(v string) *ModifySecurityIpsRequest
	GetDBInstanceIPArrayAttribute() *string
	SetDBInstanceIPArrayName(v string) *ModifySecurityIpsRequest
	GetDBInstanceIPArrayName() *string
	SetDBInstanceId(v string) *ModifySecurityIpsRequest
	GetDBInstanceId() *string
	SetFreshWhiteListReadins(v string) *ModifySecurityIpsRequest
	GetFreshWhiteListReadins() *string
	SetModifyMode(v string) *ModifySecurityIpsRequest
	GetModifyMode() *string
	SetResourceOwnerId(v int64) *ModifySecurityIpsRequest
	GetResourceOwnerId() *int64
	SetSecurityIPType(v string) *ModifySecurityIpsRequest
	GetSecurityIPType() *string
	SetSecurityIps(v string) *ModifySecurityIpsRequest
	GetSecurityIps() *string
	SetWhitelistNetworkType(v string) *ModifySecurityIpsRequest
	GetWhitelistNetworkType() *string
}

type ModifySecurityIpsRequest struct {
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
	DBInstanceIPArrayName      *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	// This parameter is required.
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	FreshWhiteListReadins *string `json:"FreshWhiteListReadins,omitempty" xml:"FreshWhiteListReadins,omitempty"`
	ModifyMode            *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityIPType        *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
	// This parameter is required.
	SecurityIps          *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
	WhitelistNetworkType *string `json:"WhitelistNetworkType,omitempty" xml:"WhitelistNetworkType,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) GetDBInstanceIPArrayAttribute() *string {
	return s.DBInstanceIPArrayAttribute
}

func (s *ModifySecurityIpsRequest) GetDBInstanceIPArrayName() *string {
	return s.DBInstanceIPArrayName
}

func (s *ModifySecurityIpsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifySecurityIpsRequest) GetFreshWhiteListReadins() *string {
	return s.FreshWhiteListReadins
}

func (s *ModifySecurityIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifySecurityIpsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySecurityIpsRequest) GetSecurityIPType() *string {
	return s.SecurityIPType
}

func (s *ModifySecurityIpsRequest) GetSecurityIps() *string {
	return s.SecurityIps
}

func (s *ModifySecurityIpsRequest) GetWhitelistNetworkType() *string {
	return s.WhitelistNetworkType
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayAttribute(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayAttribute = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayName(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceId(v string) *ModifySecurityIpsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetFreshWhiteListReadins(v string) *ModifySecurityIpsRequest {
	s.FreshWhiteListReadins = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetModifyMode(v string) *ModifySecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetResourceOwnerId(v int64) *ModifySecurityIpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIPType(v string) *ModifySecurityIpsRequest {
	s.SecurityIPType = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIps(v string) *ModifySecurityIpsRequest {
	s.SecurityIps = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetWhitelistNetworkType(v string) *ModifySecurityIpsRequest {
	s.WhitelistNetworkType = &v
	return s
}

func (s *ModifySecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
