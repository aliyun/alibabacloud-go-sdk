// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDestroyDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DestroyDBInstanceResponseBody
	GetRequestId() *string
}

type DestroyDBInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DestroyDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DestroyDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DestroyDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DestroyDBInstanceResponseBody) SetRequestId(v string) *DestroyDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DestroyDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
