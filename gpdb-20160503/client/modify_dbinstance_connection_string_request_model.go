// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConnectionStringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBInstanceConnectionStringRequest
	GetClientToken() *string
	SetConnectionStringPrefix(v string) *ModifyDBInstanceConnectionStringRequest
	GetConnectionStringPrefix() *string
	SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest
	GetCurrentConnectionString() *string
	SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest
	GetDBInstanceId() *string
	SetPort(v string) *ModifyDBInstanceConnectionStringRequest
	GetPort() *string
}

type ModifyDBInstanceConnectionStringRequest struct {
	// Idempotence check. For more information, see [How to Ensure Idempotence](https://help.aliyun.com/document_detail/327176.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88**********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The endpoint prefix of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-test
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The current endpoint of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-t4n2qg19bnn98****-master.gpdb.rds.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-t4n2qg19bnn98****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The port number. Example: 5432.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5432
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyDBInstanceConnectionStringRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBInstanceConnectionStringRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *ModifyDBInstanceConnectionStringRequest) GetCurrentConnectionString() *string {
	return s.CurrentConnectionString
}

func (s *ModifyDBInstanceConnectionStringRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceConnectionStringRequest) GetPort() *string {
	return s.Port
}

func (s *ModifyDBInstanceConnectionStringRequest) SetClientToken(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetPort(v string) *ModifyDBInstanceConnectionStringRequest {
	s.Port = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) Validate() error {
	return dara.Validate(s)
}
