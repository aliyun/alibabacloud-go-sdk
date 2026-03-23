// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTempDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTempDBInstanceResponseBody
	GetRequestId() *string
	SetTempDBInstanceId(v string) *CreateTempDBInstanceResponseBody
	GetTempDBInstanceId() *string
}

type CreateTempDBInstanceResponseBody struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TempDBInstanceId *string `json:"TempDBInstanceId,omitempty" xml:"TempDBInstanceId,omitempty"`
}

func (s CreateTempDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTempDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTempDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTempDBInstanceResponseBody) GetTempDBInstanceId() *string {
	return s.TempDBInstanceId
}

func (s *CreateTempDBInstanceResponseBody) SetRequestId(v string) *CreateTempDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTempDBInstanceResponseBody) SetTempDBInstanceId(v string) *CreateTempDBInstanceResponseBody {
	s.TempDBInstanceId = &v
	return s
}

func (s *CreateTempDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
