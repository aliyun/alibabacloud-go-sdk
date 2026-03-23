// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBNodeResponseBody
	GetDBInstanceId() *string
	SetOrderId(v int64) *ModifyDBNodeResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyDBNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDBNodeResponseBody
	GetSuccess() *bool
}

type ModifyDBNodeResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId      *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDBNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBNodeResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyDBNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDBNodeResponseBody) SetDBInstanceId(v string) *ModifyDBNodeResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBNodeResponseBody) SetOrderId(v int64) *ModifyDBNodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBNodeResponseBody) SetRequestId(v string) *ModifyDBNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBNodeResponseBody) SetSuccess(v bool) *ModifyDBNodeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDBNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
