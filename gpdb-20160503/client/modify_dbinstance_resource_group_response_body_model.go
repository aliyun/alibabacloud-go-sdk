// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceResourceGroupResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceResourceGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 65BDA532-28AF-4122-AA39-B382721EEE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceResourceGroupResponseBody) SetRequestId(v string) *ModifyDBInstanceResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
