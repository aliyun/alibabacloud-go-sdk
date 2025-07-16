// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDBInstanceTDEResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateDBInstanceTDEResponseBodyData) *UpdateDBInstanceTDEResponseBody
	GetData() *UpdateDBInstanceTDEResponseBodyData
	SetRequestId(v string) *UpdateDBInstanceTDEResponseBody
	GetRequestId() *string
}

type UpdateDBInstanceTDEResponseBody struct {
	Data *UpdateDBInstanceTDEResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDBInstanceTDEResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstanceTDEResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceTDEResponseBody) GetData() *UpdateDBInstanceTDEResponseBodyData {
	return s.Data
}

func (s *UpdateDBInstanceTDEResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDBInstanceTDEResponseBody) SetData(v *UpdateDBInstanceTDEResponseBodyData) *UpdateDBInstanceTDEResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDBInstanceTDEResponseBody) SetRequestId(v string) *UpdateDBInstanceTDEResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDBInstanceTDEResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateDBInstanceTDEResponseBodyData struct {
	// example:
	//
	// 42292****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateDBInstanceTDEResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstanceTDEResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceTDEResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateDBInstanceTDEResponseBodyData) SetTaskId(v string) *UpdateDBInstanceTDEResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateDBInstanceTDEResponseBodyData) Validate() error {
	return dara.Validate(s)
}
