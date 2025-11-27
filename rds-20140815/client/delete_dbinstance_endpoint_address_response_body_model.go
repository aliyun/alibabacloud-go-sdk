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
	// The returned data.
	Data *DeleteDBInstanceEndpointAddressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F2911788-25E8-42E5-A3A3-1B38D263F01E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
