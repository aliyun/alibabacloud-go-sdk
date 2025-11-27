// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceReplicationSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceReplicationSwitchResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceReplicationSwitchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9F8C06AD-3F37-57A0-ABBF-ABD7824F55CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceReplicationSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceReplicationSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceReplicationSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceReplicationSwitchResponseBody) SetRequestId(v string) *ModifyDBInstanceReplicationSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceReplicationSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
