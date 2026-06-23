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
	// The attribute of the whitelist group. The default value is **empty**.
	//
	// example:
	//
	// NULL
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the whitelist group to modify.
	//
	// > If you do not specify this parameter, the whitelist of the default group is modified.
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
	// The modification method. Valid values:
	//
	// - **Cover**: Overwrites the whitelist.
	//
	// - **Append**: Adds IP addresses to the whitelist.
	//
	// - **Delete**: Removes IP addresses from the whitelist.
	//
	// > If you do not specify this parameter, the Cover method is used by default.
	//
	// example:
	//
	// Cover
	ModifyMode           *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP address whitelist. The following formats are supported:
	//
	// - IP address: For example, 192.168.0.1. This allows the specified IP address to access the ApsaraDB for ClickHouse cluster.
	//
	// - CIDR block: For example, 192.168.0.0/24. This allows IP addresses from 192.168.0.1 to 192.168.0.255 to access the ApsaraDB for ClickHouse cluster.
	//
	// > 	- Do not enter 0.0.0.0.
	//
	// - Set this parameter to 127.0.0.1 to block access from all IP addresses.
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
