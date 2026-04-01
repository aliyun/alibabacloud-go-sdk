// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceEndpointAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyDBInstanceEndpointAddressResponseBodyData) *ModifyDBInstanceEndpointAddressResponseBody
	GetData() *ModifyDBInstanceEndpointAddressResponseBodyData
	SetRequestId(v string) *ModifyDBInstanceEndpointAddressResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceEndpointAddressResponseBody struct {
	// The data returned.
	Data *ModifyDBInstanceEndpointAddressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 90496720-2319-42A8-87CD-FCE4DF95EBED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceEndpointAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceEndpointAddressResponseBody) GetData() *ModifyDBInstanceEndpointAddressResponseBodyData {
	return s.Data
}

func (s *ModifyDBInstanceEndpointAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceEndpointAddressResponseBody) SetData(v *ModifyDBInstanceEndpointAddressResponseBodyData) *ModifyDBInstanceEndpointAddressResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDBInstanceEndpointAddressResponseBody) SetRequestId(v string) *ModifyDBInstanceEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDBInstanceEndpointAddressResponseBodyData struct {
	// The endpoint ID of the instance.
	//
	// example:
	//
	// ep-****
	DBInstanceEndpointId *string `json:"DBInstanceEndpointId,omitempty" xml:"DBInstanceEndpointId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rm-****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s ModifyDBInstanceEndpointAddressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceEndpointAddressResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceEndpointAddressResponseBodyData) GetDBInstanceEndpointId() *string {
	return s.DBInstanceEndpointId
}

func (s *ModifyDBInstanceEndpointAddressResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyDBInstanceEndpointAddressResponseBodyData) SetDBInstanceEndpointId(v string) *ModifyDBInstanceEndpointAddressResponseBodyData {
	s.DBInstanceEndpointId = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressResponseBodyData) SetDBInstanceName(v string) *ModifyDBInstanceEndpointAddressResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressResponseBodyData) Validate() error {
	return dara.Validate(s)
}
