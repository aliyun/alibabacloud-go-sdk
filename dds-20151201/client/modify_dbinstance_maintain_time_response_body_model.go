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
	// The request ID.
	//
	// example:
	//
	// A9310426-C763-42A2-A3AD-70A8DA204531
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
