// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDBInstanceResponseBodyData) *CreateDBInstanceResponseBody
	GetData() *CreateDBInstanceResponseBodyData
	SetRequestId(v string) *CreateDBInstanceResponseBody
	GetRequestId() *string
}

type CreateDBInstanceResponseBody struct {
	// The returned result.
	Data *CreateDBInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4773E4EC-025D-509F-AEA9-D53123FDFB0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBody) GetData() *CreateDBInstanceResponseBodyData {
	return s.Data
}

func (s *CreateDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBInstanceResponseBody) SetData(v *CreateDBInstanceResponseBodyData) *CreateDBInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateDBInstanceResponseBody) SetRequestId(v string) *CreateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDBInstanceResponseBodyData struct {
	// The instance ID.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 21137950671****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDBInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstanceResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateDBInstanceResponseBodyData) SetDBInstanceId(v string) *CreateDBInstanceResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceResponseBodyData) SetOrderId(v int64) *CreateDBInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateDBInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
