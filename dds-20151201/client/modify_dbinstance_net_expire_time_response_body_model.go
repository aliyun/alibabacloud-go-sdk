// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceNetExpireTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceNetExpireTimeResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceNetExpireTimeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 459E7D5C-38DA-4E14-9C82-5B5AF693DBAB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceNetExpireTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceNetExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetExpireTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceNetExpireTimeResponseBody) SetRequestId(v string) *ModifyDBInstanceNetExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
