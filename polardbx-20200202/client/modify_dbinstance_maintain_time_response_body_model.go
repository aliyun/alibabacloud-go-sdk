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
	// example:
	//
	// 73559800-3c8c-11ec-bd40-99cfcff3fe1e
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
