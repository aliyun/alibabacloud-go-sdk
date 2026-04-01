// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCreateDdrDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsValid(v string) *CheckCreateDdrDBInstanceResponseBody
	GetIsValid() *string
	SetRequestId(v string) *CheckCreateDdrDBInstanceResponseBody
	GetRequestId() *string
}

type CheckCreateDdrDBInstanceResponseBody struct {
	// Indicates whether the data of the source instance can be restored across regions. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsValid *string `json:"IsValid,omitempty" xml:"IsValid,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1E43AAE0-BEE8-43DA-860D-EAF2AA0724DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckCreateDdrDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCreateDdrDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCreateDdrDBInstanceResponseBody) GetIsValid() *string {
	return s.IsValid
}

func (s *CheckCreateDdrDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCreateDdrDBInstanceResponseBody) SetIsValid(v string) *CheckCreateDdrDBInstanceResponseBody {
	s.IsValid = &v
	return s
}

func (s *CheckCreateDdrDBInstanceResponseBody) SetRequestId(v string) *CheckCreateDdrDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCreateDdrDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
