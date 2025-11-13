// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDBClusterResponseBodyData) *CreateDBClusterResponseBody
	GetData() *CreateDBClusterResponseBodyData
	SetRequestId(v string) *CreateDBClusterResponseBody
	GetRequestId() *string
}

type CreateDBClusterResponseBody struct {
	Data *CreateDBClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// F8900A96-67F7-5274-A41B-7722E1ECF8C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponseBody) GetData() *CreateDBClusterResponseBodyData {
	return s.Data
}

func (s *CreateDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBClusterResponseBody) SetData(v *CreateDBClusterResponseBodyData) *CreateDBClusterResponseBody {
	s.Data = v
	return s
}

func (s *CreateDBClusterResponseBody) SetRequestId(v string) *CreateDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDBClusterResponseBodyData struct {
	// example:
	//
	// selectdb-cn-1ls3sg0po0****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 219543646290345
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDBClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateDBClusterResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBClusterResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateDBClusterResponseBodyData) SetClusterId(v string) *CreateDBClusterResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *CreateDBClusterResponseBodyData) SetDBInstanceId(v string) *CreateDBClusterResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBClusterResponseBodyData) SetOrderId(v int64) *CreateDBClusterResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateDBClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
