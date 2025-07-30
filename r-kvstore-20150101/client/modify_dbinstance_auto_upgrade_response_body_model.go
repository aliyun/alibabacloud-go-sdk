// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceAutoUpgradeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceAutoUpgradeResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceAutoUpgradeResponseBody struct {
	// ID of the request.
	//
	// example:
	//
	// 2FF6158E-3394-4A90-B634-79C49184****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceAutoUpgradeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceAutoUpgradeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAutoUpgradeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceAutoUpgradeResponseBody) SetRequestId(v string) *ModifyDBInstanceAutoUpgradeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeResponseBody) Validate() error {
	return dara.Validate(s)
}
