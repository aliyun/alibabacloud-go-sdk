// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoUpgradeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUpgrade(v string) *UpdateAutoUpgradeRequest
	GetAutoUpgrade() *string
}

type UpdateAutoUpgradeRequest struct {
	// example:
	//
	// auto
	AutoUpgrade *string `json:"autoUpgrade,omitempty" xml:"autoUpgrade,omitempty"`
}

func (s UpdateAutoUpgradeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoUpgradeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoUpgradeRequest) GetAutoUpgrade() *string {
	return s.AutoUpgrade
}

func (s *UpdateAutoUpgradeRequest) SetAutoUpgrade(v string) *UpdateAutoUpgradeRequest {
	s.AutoUpgrade = &v
	return s
}

func (s *UpdateAutoUpgradeRequest) Validate() error {
	return dara.Validate(s)
}
