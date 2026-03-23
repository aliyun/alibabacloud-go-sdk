// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyDBInstanceEndpointResponseBodyData) *ModifyDBInstanceEndpointResponseBody
	GetData() *ModifyDBInstanceEndpointResponseBodyData
	SetRequestId(v string) *ModifyDBInstanceEndpointResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceEndpointResponseBody struct {
	Data      *ModifyDBInstanceEndpointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceEndpointResponseBody) GetData() *ModifyDBInstanceEndpointResponseBodyData {
	return s.Data
}

func (s *ModifyDBInstanceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceEndpointResponseBody) SetData(v *ModifyDBInstanceEndpointResponseBodyData) *ModifyDBInstanceEndpointResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDBInstanceEndpointResponseBody) SetRequestId(v string) *ModifyDBInstanceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceEndpointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDBInstanceEndpointResponseBodyData struct {
	DBInstanceEndpointId *string `json:"DBInstanceEndpointId,omitempty" xml:"DBInstanceEndpointId,omitempty"`
	DBInstanceName       *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s ModifyDBInstanceEndpointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceEndpointResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceEndpointResponseBodyData) GetDBInstanceEndpointId() *string {
	return s.DBInstanceEndpointId
}

func (s *ModifyDBInstanceEndpointResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyDBInstanceEndpointResponseBodyData) SetDBInstanceEndpointId(v string) *ModifyDBInstanceEndpointResponseBodyData {
	s.DBInstanceEndpointId = &v
	return s
}

func (s *ModifyDBInstanceEndpointResponseBodyData) SetDBInstanceName(v string) *ModifyDBInstanceEndpointResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceEndpointResponseBodyData) Validate() error {
	return dara.Validate(s)
}
