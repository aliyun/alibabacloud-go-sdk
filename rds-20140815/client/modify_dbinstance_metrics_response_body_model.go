// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceMetricsResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *ModifyDBInstanceMetricsResponseBody
	GetRequestId() *string
	SetScope(v string) *ModifyDBInstanceMetricsResponseBody
	GetScope() *string
}

type ModifyDBInstanceMetricsResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scope        *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ModifyDBInstanceMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMetricsResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceMetricsResponseBody) GetScope() *string {
	return s.Scope
}

func (s *ModifyDBInstanceMetricsResponseBody) SetDBInstanceId(v string) *ModifyDBInstanceMetricsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceMetricsResponseBody) SetRequestId(v string) *ModifyDBInstanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceMetricsResponseBody) SetScope(v string) *ModifyDBInstanceMetricsResponseBody {
	s.Scope = &v
	return s
}

func (s *ModifyDBInstanceMetricsResponseBody) Validate() error {
	return dara.Validate(s)
}
