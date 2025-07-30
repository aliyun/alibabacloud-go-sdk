// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCreateDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckCreateDBInstanceResponseBody
	GetRequestId() *string
}

type CheckCreateDBInstanceResponseBody struct {
	// example:
	//
	// ADF42B18-43FD-5100-83A9-BE81AB70C863
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckCreateDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCreateDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCreateDBInstanceResponseBody) SetRequestId(v string) *CheckCreateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCreateDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
