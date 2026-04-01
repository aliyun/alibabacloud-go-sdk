// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDBInstanceReplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDBInstanceReplicationResponseBody
	GetRequestId() *string
}

type UpdateDBInstanceReplicationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 17F57FEE-EA4F-4337-8D2E-9C23CAA63D74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDBInstanceReplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstanceReplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceReplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDBInstanceReplicationResponseBody) SetRequestId(v string) *UpdateDBInstanceReplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDBInstanceReplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
