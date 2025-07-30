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
	// The ID of the request.
	//
	// example:
	//
	// 1790D68A-465C-44E3-BC24-9732652961F9
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
