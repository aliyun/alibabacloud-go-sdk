// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceEndpointAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDBInstanceEndpointAddressResponseBodyData) *CreateDBInstanceEndpointAddressResponseBody
	GetData() *CreateDBInstanceEndpointAddressResponseBodyData
	SetRequestId(v string) *CreateDBInstanceEndpointAddressResponseBody
	GetRequestId() *string
}

type CreateDBInstanceEndpointAddressResponseBody struct {
	// The data returned.
	Data *CreateDBInstanceEndpointAddressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 08A3B71B-FE08-4B03-974F-CC7EA6DB1828
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceEndpointAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceEndpointAddressResponseBody) GetData() *CreateDBInstanceEndpointAddressResponseBodyData {
	return s.Data
}

func (s *CreateDBInstanceEndpointAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBInstanceEndpointAddressResponseBody) SetData(v *CreateDBInstanceEndpointAddressResponseBodyData) *CreateDBInstanceEndpointAddressResponseBody {
	s.Data = v
	return s
}

func (s *CreateDBInstanceEndpointAddressResponseBody) SetRequestId(v string) *CreateDBInstanceEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceEndpointAddressResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDBInstanceEndpointAddressResponseBodyData struct {
	// The public endpoint.
	//
	// example:
	//
	// rm-******.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
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

func (s CreateDBInstanceEndpointAddressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceEndpointAddressResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceEndpointAddressResponseBodyData) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreateDBInstanceEndpointAddressResponseBodyData) GetDBInstanceEndpointId() *string {
	return s.DBInstanceEndpointId
}

func (s *CreateDBInstanceEndpointAddressResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateDBInstanceEndpointAddressResponseBodyData) SetConnectionString(v string) *CreateDBInstanceEndpointAddressResponseBodyData {
	s.ConnectionString = &v
	return s
}

func (s *CreateDBInstanceEndpointAddressResponseBodyData) SetDBInstanceEndpointId(v string) *CreateDBInstanceEndpointAddressResponseBodyData {
	s.DBInstanceEndpointId = &v
	return s
}

func (s *CreateDBInstanceEndpointAddressResponseBodyData) SetDBInstanceName(v string) *CreateDBInstanceEndpointAddressResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *CreateDBInstanceEndpointAddressResponseBodyData) Validate() error {
	return dara.Validate(s)
}
