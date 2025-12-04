// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAccessWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterIPArrayAttribute(v string) *ModifyDBClusterAccessWhiteListRequest
	GetDBClusterIPArrayAttribute() *string
	SetDBClusterIPArrayName(v string) *ModifyDBClusterAccessWhiteListRequest
	GetDBClusterIPArrayName() *string
	SetDBClusterId(v string) *ModifyDBClusterAccessWhiteListRequest
	GetDBClusterId() *string
	SetModifyMode(v string) *ModifyDBClusterAccessWhiteListRequest
	GetModifyMode() *string
	SetOwnerAccount(v string) *ModifyDBClusterAccessWhiteListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterAccessWhiteListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBClusterAccessWhiteListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterAccessWhiteListRequest
	GetResourceOwnerId() *int64
	SetSecurityIps(v string) *ModifyDBClusterAccessWhiteListRequest
	GetSecurityIps() *string
}

type ModifyDBClusterAccessWhiteListRequest struct {
	// The attribute of the IP address whitelist. By default, this parameter is **empty**.
	//
	// example:
	//
	// NULL
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist that you want to modify.
	//
	// >  If you do not specify this parameter, the default IP address whitelist is modified.
	//
	// example:
	//
	// default
	DBClusterIPArrayName *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1qx68m06981****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The method that is used to modify the IP address whitelist. Valid values:
	//
	// 	- **Cover**: overwrites the original IP address whitelist.
	//
	// 	- **Append**: appends the specified IP addresses to the original IP address whitelist.
	//
	// 	- **Delete**: deletes the original IP address whitelist.
	//
	// >  If you do not specify this parameter, the default value of Cover is used.
	//
	// example:
	//
	// Cover
	ModifyMode           *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP addresses in the IP address whitelist. You can specify IP addresses in the following formats:
	//
	// 	- IP address. For example, you can set SecurityIps to 192.168.0.1. This allows you to use this IP address to access your ApsaraDB for ClickHouse cluster.
	//
	// 	- CIDR block. For example, you can set SecurityIps to 192.168.0.0/24. This allows you to use the IP addresses from 192.168.0.1 to 192.168.0.255 to access your ApsaraDB for ClickHouse cluster.
	//
	// >
	//
	// 	- Do not set SecurityIps to 0.0.0.0.
	//
	// 	- If you set SecurityIps to 127.0.0.1, all IP addresses are blocked from accessing your ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.xx.xx
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s ModifyDBClusterAccessWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAccessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhiteListRequest) GetDBClusterIPArrayAttribute() *string {
	return s.DBClusterIPArrayAttribute
}

func (s *ModifyDBClusterAccessWhiteListRequest) GetDBClusterIPArrayName() *string {
	return s.DBClusterIPArrayName
}

func (s *ModifyDBClusterAccessWhiteListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterAccessWhiteListRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyDBClusterAccessWhiteListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterAccessWhiteListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterAccessWhiteListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterAccessWhiteListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterAccessWhiteListRequest) GetSecurityIps() *string {
	return s.SecurityIps
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetDBClusterIPArrayAttribute(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetDBClusterIPArrayName(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetDBClusterId(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetModifyMode(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetOwnerAccount(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetOwnerId(v int64) *ModifyDBClusterAccessWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetResourceOwnerId(v int64) *ModifyDBClusterAccessWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetSecurityIps(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.SecurityIps = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
