// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceNetworkTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceNetworkTypeResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceNetworkTypeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2d0c35a9-f5da-44ba-852d-741e27b7eb0b
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceNetworkTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceNetworkTypeResponseBody) SetRequestId(v string) *ModifyDBInstanceNetworkTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
