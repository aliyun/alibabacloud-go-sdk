// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceHAConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceHAConfigResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceHAConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D4D4BE8A-DD46-440A-BFCD-EE31DA81C9DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceHAConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceHAConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceHAConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceHAConfigResponseBody) SetRequestId(v string) *ModifyDBInstanceHAConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceHAConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
