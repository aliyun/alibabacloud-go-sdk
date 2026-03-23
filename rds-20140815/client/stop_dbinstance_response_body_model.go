// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopDBInstanceResponseBody
	GetRequestId() *string
}

type StopDBInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDBInstanceResponseBody) SetRequestId(v string) *StopDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
