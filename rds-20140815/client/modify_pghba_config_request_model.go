// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPGHbaConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyPGHbaConfigRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyPGHbaConfigRequest
	GetDBInstanceId() *string
	SetHbaItem(v []*ModifyPGHbaConfigRequestHbaItem) *ModifyPGHbaConfigRequest
	GetHbaItem() []*ModifyPGHbaConfigRequestHbaItem
	SetOpsType(v string) *ModifyPGHbaConfigRequest
	GetOpsType() *string
	SetOwnerAccount(v string) *ModifyPGHbaConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyPGHbaConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyPGHbaConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyPGHbaConfigRequest
	GetResourceOwnerId() *int64
}

type ModifyPGHbaConfigRequest struct {
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// 1
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp1lymyn1v3i****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// An array that consists of the details of the AD domain services.
	//
	// This parameter is required.
	HbaItem []*ModifyPGHbaConfigRequestHbaItem `json:"HbaItem,omitempty" xml:"HbaItem,omitempty" type:"Repeated"`
	// The method that you use to modify the pg_hba.conf file. Valid values:
	//
	// 	- **add**: adds one or more records. If you use this method, make sure that the value of the PriorityId parameter for each new record is different from the value of the PriorityId parameter for any existing record.
	//
	// 	- **delete**: deletes one or more records. If you use this method, the record that corresponds to the specified value of the **PriorityId*	- parameter is deleted from the pg_hba.conf file.
	//
	// 	- **modify**: modifies one or more records. If you use this method, the record that corresponds to the specified value of the **PriorityId*	- parameter is modified.
	//
	// 	- **update**: overwrites the existing configuration in the pg_hba.conf file by using the new configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	OpsType              *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyPGHbaConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPGHbaConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyPGHbaConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyPGHbaConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyPGHbaConfigRequest) GetHbaItem() []*ModifyPGHbaConfigRequestHbaItem {
	return s.HbaItem
}

func (s *ModifyPGHbaConfigRequest) GetOpsType() *string {
	return s.OpsType
}

func (s *ModifyPGHbaConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyPGHbaConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPGHbaConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPGHbaConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyPGHbaConfigRequest) SetClientToken(v string) *ModifyPGHbaConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetDBInstanceId(v string) *ModifyPGHbaConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetHbaItem(v []*ModifyPGHbaConfigRequestHbaItem) *ModifyPGHbaConfigRequest {
	s.HbaItem = v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetOpsType(v string) *ModifyPGHbaConfigRequest {
	s.OpsType = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetOwnerAccount(v string) *ModifyPGHbaConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetOwnerId(v int64) *ModifyPGHbaConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetResourceOwnerAccount(v string) *ModifyPGHbaConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetResourceOwnerId(v int64) *ModifyPGHbaConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) Validate() error {
	if s.HbaItem != nil {
		for _, item := range s.HbaItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyPGHbaConfigRequestHbaItem struct {
	// The IP addresses from which the specified users can access the specified databases. If you set this parameter to 0.0.0.0/0, the specified users are allowed to access the specified databases from all IP addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.0.0/0
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The name of the database. If you set this parameter to all, the specified users are allowed to access all databases on the instance.
	//
	// If you specify multiple entries, separate the entries with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The mask of the IP address. If the value of the **Address*	- parameter is an IP address, you can use this parameter to specify the mask of the IP address.
	//
	// example:
	//
	// 0
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The authentication method. Valid values:
	//
	// 	- **trust**
	//
	// 	- **reject**
	//
	// 	- **scram-sha-256**
	//
	// 	- **md5**
	//
	// 	- **password**
	//
	// 	- **gss**
	//
	// 	- **sspi**
	//
	// 	- **ldap**
	//
	// 	- **radius**
	//
	// 	- **cert**
	//
	// 	- **pam**
	//
	// This parameter is required.
	//
	// example:
	//
	// ldap
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The options of the authentication method. In this topic, LDAP is used as an example. You must configure this parameter. For more information, see [Authentication Methods](https://www.postgresql.org/docs/11/auth-methods.html).
	//
	// example:
	//
	// ldapserver=The private IP address of the ECS instance ldapbasedn="CN=Users,DC=pgsqldomain,DC=net" ldapbinddn="CN=The username of the administrator user of the AD domain controller, CN=Users,DC=pgsqldomain,DC=net" ldapbindpasswd="The password of the administrator user of the AD domain controller" ldapsearchattribute="sAMAccountName"
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// The priority of the record. If you set this parameter to 0, the record has the highest priority. Valid values: 0 to 10000.
	//
	// This parameter is used to identify each record. When you add a record, the value of the PriorityId parameter for the new record must be different from the value of the PriorityId parameter of any existing record. When you modify or delete a record, you must also modify or delete the value of the PriorityId parameter for this record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	PriorityId *int32 `json:"PriorityId,omitempty" xml:"PriorityId,omitempty"`
	// The connection type.
	//
	// Valid values:
	//
	// 	- **host**: The record matches TCP/IP connections, including SSL connections and non-SSL connections.
	//
	// 	- **hostssl**: The record matches only TCP/IP connections that are established over SSL.
	//
	// 	- **hostnossl**: The record matches only TCP/IP connections that are not established over SSL connections.
	//
	// >  You can set this parameter to hostssl only when SSL encryption is enabled for the instance. For more information, see Configure SSL encryption for an ApsaraDB RDS for PostgreSQL instance.[](~~229518~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// host
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who is allowed to access the specified databases. You must specify the user that is used to log on to the RDS instance. If you specify multiple entries, separate the entries with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ldapuser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ModifyPGHbaConfigRequestHbaItem) String() string {
	return dara.Prettify(s)
}

func (s ModifyPGHbaConfigRequestHbaItem) GoString() string {
	return s.String()
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetAddress() *string {
	return s.Address
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetDatabase() *string {
	return s.Database
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetMask() *string {
	return s.Mask
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetMethod() *string {
	return s.Method
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetOption() *string {
	return s.Option
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetPriorityId() *int32 {
	return s.PriorityId
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetType() *string {
	return s.Type
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetUser() *string {
	return s.User
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetAddress(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.Address = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetDatabase(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.Database = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetMask(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.Mask = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetMethod(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.Method = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetOption(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.Option = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetPriorityId(v int32) *ModifyPGHbaConfigRequestHbaItem {
	s.PriorityId = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetType(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.Type = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetUser(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.User = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) Validate() error {
	return dara.Validate(s)
}
