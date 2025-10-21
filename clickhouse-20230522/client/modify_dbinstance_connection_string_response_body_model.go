// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConnectionStringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyDBInstanceConnectionStringResponseBodyData) *ModifyDBInstanceConnectionStringResponseBody
	GetData() *ModifyDBInstanceConnectionStringResponseBodyData
	SetRequestId(v string) *ModifyDBInstanceConnectionStringResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceConnectionStringResponseBody struct {
	// The data returned.
	Data *ModifyDBInstanceConnectionStringResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponseBody) GetData() *ModifyDBInstanceConnectionStringResponseBodyData {
	return s.Data
}

func (s *ModifyDBInstanceConnectionStringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetData(v *ModifyDBInstanceConnectionStringResponseBodyData) *ModifyDBInstanceConnectionStringResponseBody {
	s.Data = v
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
	// The endpoint of the cluster.
	//
	// example:
	//
	// cc-2ze34****-clickhouse..clickhouseserver.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The disabled database ports.
	//
	// example:
	//
	// 9001,8123
	DisabledPorts *string `json:"DisabledPorts,omitempty" xml:"DisabledPorts,omitempty"`
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

func (s *ModifyDBInstanceConnectionStringResponseBodyData) GetDBInstanceID() *int32 {
	return s.DBInstanceID
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) GetDisabledPorts() *string {
	return s.DisabledPorts
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetConnectionString(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetDBInstanceID(v int32) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetDBInstanceName(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetDisabledPorts(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.DisabledPorts = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) Validate() error {
	return dara.Validate(s)
}
