// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMaintainTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceMaintainTimeResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceMaintainTimeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CA9A34C8-AC95-413B-AC6A-CE**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceMaintainTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceMaintainTimeResponseBody) SetRequestId(v string) *ModifyDBInstanceMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
