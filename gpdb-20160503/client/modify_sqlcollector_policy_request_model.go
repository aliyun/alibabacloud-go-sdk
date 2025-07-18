// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySQLCollectorPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifySQLCollectorPolicyRequest
	GetDBInstanceId() *string
	SetSQLCollectorStatus(v string) *ModifySQLCollectorPolicyRequest
	GetSQLCollectorStatus() *string
}

type ModifySQLCollectorPolicyRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to enable or disable SQL collection.
	//
	// 	- Enable: enables SQL collection.
	//
	// 	- Disabled: disables SQL collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	SQLCollectorStatus *string `json:"SQLCollectorStatus,omitempty" xml:"SQLCollectorStatus,omitempty"`
}

func (s ModifySQLCollectorPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySQLCollectorPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifySQLCollectorPolicyRequest) GetSQLCollectorStatus() *string {
	return s.SQLCollectorStatus
}

func (s *ModifySQLCollectorPolicyRequest) SetDBInstanceId(v string) *ModifySQLCollectorPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySQLCollectorPolicyRequest) SetSQLCollectorStatus(v string) *ModifySQLCollectorPolicyRequest {
	s.SQLCollectorStatus = &v
	return s
}

func (s *ModifySQLCollectorPolicyRequest) Validate() error {
	return dara.Validate(s)
}
