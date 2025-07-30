// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBClusterResponseBody
	GetDBInstanceId() *string
	SetData(v *ModifyDBClusterResponseBodyData) *ModifyDBClusterResponseBody
	GetData() *ModifyDBClusterResponseBodyData
	SetRequestId(v string) *ModifyDBClusterResponseBody
	GetRequestId() *string
}

type ModifyDBClusterResponseBody struct {
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string                          `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Data         *ModifyDBClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0BF61F90-ACED-55DF-A6FE-56714B2BFCF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBClusterResponseBody) GetData() *ModifyDBClusterResponseBodyData {
	return s.Data
}

func (s *ModifyDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterResponseBody) SetDBInstanceId(v string) *ModifyDBClusterResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) SetData(v *ModifyDBClusterResponseBodyData) *ModifyDBClusterResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDBClusterResponseBody) SetRequestId(v string) *ModifyDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyDBClusterResponseBodyData struct {
	// example:
	//
	// selectdb-cn-7213cjv****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 219396937240838
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyDBClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBClusterResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyDBClusterResponseBodyData) SetDBClusterId(v string) *ModifyDBClusterResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterResponseBodyData) SetDBInstanceId(v string) *ModifyDBClusterResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBClusterResponseBodyData) SetOrderId(v int64) *ModifyDBClusterResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ModifyDBClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
