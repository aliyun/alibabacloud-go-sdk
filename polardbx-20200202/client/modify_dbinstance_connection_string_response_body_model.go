// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConnectionStringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ModifyDBInstanceConnectionStringResponseBody
	GetCode() *int64
	SetData(v *ModifyDBInstanceConnectionStringResponseBodyData) *ModifyDBInstanceConnectionStringResponseBody
	GetData() *ModifyDBInstanceConnectionStringResponseBodyData
	SetMessage(v string) *ModifyDBInstanceConnectionStringResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyDBInstanceConnectionStringResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceConnectionStringResponseBody struct {
	// example:
	//
	// 200
	Code *int64                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ModifyDBInstanceConnectionStringResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// c3cf535c-a585-11ea-8263-00163e04d3a7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ModifyDBInstanceConnectionStringResponseBody) GetData() *ModifyDBInstanceConnectionStringResponseBodyData {
	return s.Data
}

func (s *ModifyDBInstanceConnectionStringResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyDBInstanceConnectionStringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetCode(v int64) *ModifyDBInstanceConnectionStringResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetData(v *ModifyDBInstanceConnectionStringResponseBodyData) *ModifyDBInstanceConnectionStringResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetMessage(v string) *ModifyDBInstanceConnectionStringResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetRequestId(v string) *ModifyDBInstanceConnectionStringResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDBInstanceConnectionStringResponseBodyData struct {
	// example:
	//
	// test2.polarx.huhehaote.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// pxc-unrf5ssig0ecg8
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 1
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// example:
	//
	// 3300
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyDBInstanceConnectionStringResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) GetPort() *string {
	return s.Port
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetConnectionString(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetDBInstanceName(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetDBInstanceNetType(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.DBInstanceNetType = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetPort(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.Port = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) Validate() error {
	return dara.Validate(s)
}
