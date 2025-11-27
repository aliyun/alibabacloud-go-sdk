// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDBInstanceEndpointResponseBodyData) *CreateDBInstanceEndpointResponseBody
	GetData() *CreateDBInstanceEndpointResponseBodyData
	SetRequestId(v string) *CreateDBInstanceEndpointResponseBody
	GetRequestId() *string
}

type CreateDBInstanceEndpointResponseBody struct {
	// The data returned.
	Data *CreateDBInstanceEndpointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C8E88DED-533F-4B3C-9207-731FBF394CCA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceEndpointResponseBody) GetData() *CreateDBInstanceEndpointResponseBodyData {
	return s.Data
}

func (s *CreateDBInstanceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBInstanceEndpointResponseBody) SetData(v *CreateDBInstanceEndpointResponseBodyData) *CreateDBInstanceEndpointResponseBody {
	s.Data = v
	return s
}

func (s *CreateDBInstanceEndpointResponseBody) SetRequestId(v string) *CreateDBInstanceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceEndpointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDBInstanceEndpointResponseBodyData struct {
	// The internal endpoint.
	//
	// example:
	//
	// rm-****.mysql.rds.aliyuncs.com
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

func (s CreateDBInstanceEndpointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceEndpointResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceEndpointResponseBodyData) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreateDBInstanceEndpointResponseBodyData) GetDBInstanceEndpointId() *string {
	return s.DBInstanceEndpointId
}

func (s *CreateDBInstanceEndpointResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateDBInstanceEndpointResponseBodyData) SetConnectionString(v string) *CreateDBInstanceEndpointResponseBodyData {
	s.ConnectionString = &v
	return s
}

func (s *CreateDBInstanceEndpointResponseBodyData) SetDBInstanceEndpointId(v string) *CreateDBInstanceEndpointResponseBodyData {
	s.DBInstanceEndpointId = &v
	return s
}

func (s *CreateDBInstanceEndpointResponseBodyData) SetDBInstanceName(v string) *CreateDBInstanceEndpointResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *CreateDBInstanceEndpointResponseBodyData) Validate() error {
	return dara.Validate(s)
}
