// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *ModifyDBInstanceSSLRequest
	GetConnectionString() *string
	SetDBInstanceId(v string) *ModifyDBInstanceSSLRequest
	GetDBInstanceId() *string
	SetSSLEnabled(v int32) *ModifyDBInstanceSSLRequest
	GetSSLEnabled() *int32
}

type ModifyDBInstanceSSLRequest struct {
	// The encrypted endpoint. By default, the wildcards are used for instances that are hosted on ECS instances. This way, the endpoints that can be resolved to the same IP address are encrypted.
	//
	// example:
	//
	// gp-xxxxxxxxxxx-master.gpdbmaster.singapore.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The status of SSL encryption. Valid values:
	//
	// 	- 0: disables SSL encryption.
	//
	// 	- 1: enables SSL encryption.
	//
	// 	- 2: updates SSL encryption.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SSLEnabled *int32 `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
}

func (s ModifyDBInstanceSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyDBInstanceSSLRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceSSLRequest) GetSSLEnabled() *int32 {
	return s.SSLEnabled
}

func (s *ModifyDBInstanceSSLRequest) SetConnectionString(v string) *ModifyDBInstanceSSLRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetDBInstanceId(v string) *ModifyDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetSSLEnabled(v int32) *ModifyDBInstanceSSLRequest {
	s.SSLEnabled = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) Validate() error {
	return dara.Validate(s)
}
