// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiveDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGuardDBInstanceId(v string) *ReceiveDBInstanceResponseBody
	GetGuardDBInstanceId() *string
	SetRequestId(v string) *ReceiveDBInstanceResponseBody
	GetRequestId() *string
}

type ReceiveDBInstanceResponseBody struct {
	GuardDBInstanceId *string `json:"GuardDBInstanceId,omitempty" xml:"GuardDBInstanceId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReceiveDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReceiveDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReceiveDBInstanceResponseBody) GetGuardDBInstanceId() *string {
	return s.GuardDBInstanceId
}

func (s *ReceiveDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReceiveDBInstanceResponseBody) SetGuardDBInstanceId(v string) *ReceiveDBInstanceResponseBody {
	s.GuardDBInstanceId = &v
	return s
}

func (s *ReceiveDBInstanceResponseBody) SetRequestId(v string) *ReceiveDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReceiveDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
