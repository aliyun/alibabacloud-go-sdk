// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConnectionStringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceConnectionStringResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceConnectionStringResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FF36A84C-0694-42D0-861D-C383E8E4FAAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetRequestId(v string) *ModifyDBInstanceConnectionStringResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBody) Validate() error {
	return dara.Validate(s)
}
