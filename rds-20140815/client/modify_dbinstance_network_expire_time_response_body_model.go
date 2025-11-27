// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceNetworkExpireTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceNetworkExpireTimeResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceNetworkExpireTimeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceNetworkExpireTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceNetworkExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkExpireTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceNetworkExpireTimeResponseBody) SetRequestId(v string) *ModifyDBInstanceNetworkExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceNetworkExpireTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
