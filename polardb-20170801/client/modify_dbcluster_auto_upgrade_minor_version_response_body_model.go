// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAutoUpgradeMinorVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterAutoUpgradeMinorVersionResponseBody
	GetRequestId() *string
}

type ModifyDBClusterAutoUpgradeMinorVersionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterAutoUpgradeMinorVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAutoUpgradeMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAutoUpgradeMinorVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterAutoUpgradeMinorVersionResponseBody) SetRequestId(v string) *ModifyDBClusterAutoUpgradeMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterAutoUpgradeMinorVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
