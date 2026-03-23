// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceAutoUpgradeMinorVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceAutoUpgradeMinorVersionResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceAutoUpgradeMinorVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceAutoUpgradeMinorVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceAutoUpgradeMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionResponseBody) SetRequestId(v string) *ModifyDBInstanceAutoUpgradeMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
