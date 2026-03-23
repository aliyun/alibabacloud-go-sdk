// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceEndpointAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteDBInstanceEndpointAddressResponseBodyData) *DeleteDBInstanceEndpointAddressResponseBody
	GetData() *DeleteDBInstanceEndpointAddressResponseBodyData
	SetRequestId(v string) *DeleteDBInstanceEndpointAddressResponseBody
	GetRequestId() *string
}

type DeleteDBInstanceEndpointAddressResponseBody struct {
	Data      *DeleteDBInstanceEndpointAddressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBInstanceEndpointAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceEndpointAddressResponseBody) GetData() *DeleteDBInstanceEndpointAddressResponseBodyData {
	return s.Data
}

func (s *DeleteDBInstanceEndpointAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBInstanceEndpointAddressResponseBody) SetData(v *DeleteDBInstanceEndpointAddressResponseBodyData) *DeleteDBInstanceEndpointAddressResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDBInstanceEndpointAddressResponseBody) SetRequestId(v string) *DeleteDBInstanceEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBInstanceEndpointAddressResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDBInstanceEndpointAddressResponseBodyData struct {
	DBInstanceEndpointId *string `json:"DBInstanceEndpointId,omitempty" xml:"DBInstanceEndpointId,omitempty"`
	DBInstanceName       *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s DeleteDBInstanceEndpointAddressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceEndpointAddressResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceEndpointAddressResponseBodyData) GetDBInstanceEndpointId() *string {
	return s.DBInstanceEndpointId
}

func (s *DeleteDBInstanceEndpointAddressResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteDBInstanceEndpointAddressResponseBodyData) SetDBInstanceEndpointId(v string) *DeleteDBInstanceEndpointAddressResponseBodyData {
	s.DBInstanceEndpointId = &v
	return s
}

func (s *DeleteDBInstanceEndpointAddressResponseBodyData) SetDBInstanceName(v string) *DeleteDBInstanceEndpointAddressResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteDBInstanceEndpointAddressResponseBodyData) Validate() error {
	return dara.Validate(s)
}
