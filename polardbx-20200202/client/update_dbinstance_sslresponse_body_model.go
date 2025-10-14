// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDBInstanceSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateDBInstanceSSLResponseBodyData) *UpdateDBInstanceSSLResponseBody
	GetData() *UpdateDBInstanceSSLResponseBodyData
	SetRequestId(v string) *UpdateDBInstanceSSLResponseBody
	GetRequestId() *string
}

type UpdateDBInstanceSSLResponseBody struct {
	Data *UpdateDBInstanceSSLResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDBInstanceSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceSSLResponseBody) GetData() *UpdateDBInstanceSSLResponseBodyData {
	return s.Data
}

func (s *UpdateDBInstanceSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDBInstanceSSLResponseBody) SetData(v *UpdateDBInstanceSSLResponseBodyData) *UpdateDBInstanceSSLResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDBInstanceSSLResponseBody) SetRequestId(v string) *UpdateDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDBInstanceSSLResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDBInstanceSSLResponseBodyData struct {
	// example:
	//
	// 2209883
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateDBInstanceSSLResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstanceSSLResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceSSLResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *UpdateDBInstanceSSLResponseBodyData) SetTaskId(v int64) *UpdateDBInstanceSSLResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateDBInstanceSSLResponseBodyData) Validate() error {
	return dara.Validate(s)
}
