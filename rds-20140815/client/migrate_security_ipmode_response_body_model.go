// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateSecurityIPModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *MigrateSecurityIPModeResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *MigrateSecurityIPModeResponseBody
	GetRequestId() *string
	SetSecurityIPMode(v string) *MigrateSecurityIPModeResponseBody
	GetSecurityIPMode() *string
}

type MigrateSecurityIPModeResponseBody struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityIPMode *string `json:"SecurityIPMode,omitempty" xml:"SecurityIPMode,omitempty"`
}

func (s MigrateSecurityIPModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateSecurityIPModeResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateSecurityIPModeResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *MigrateSecurityIPModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateSecurityIPModeResponseBody) GetSecurityIPMode() *string {
	return s.SecurityIPMode
}

func (s *MigrateSecurityIPModeResponseBody) SetDBInstanceId(v string) *MigrateSecurityIPModeResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateSecurityIPModeResponseBody) SetRequestId(v string) *MigrateSecurityIPModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateSecurityIPModeResponseBody) SetSecurityIPMode(v string) *MigrateSecurityIPModeResponseBody {
	s.SecurityIPMode = &v
	return s
}

func (s *MigrateSecurityIPModeResponseBody) Validate() error {
	return dara.Validate(s)
}
