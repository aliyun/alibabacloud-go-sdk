// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublishAutoUpgradeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUpgrade(v int32) *UpdatePublishAutoUpgradeRequest
	GetAutoUpgrade() *int32
}

type UpdatePublishAutoUpgradeRequest struct {
	// Specifies whether to enable automatic upgrade. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	AutoUpgrade *int32 `json:"AutoUpgrade,omitempty" xml:"AutoUpgrade,omitempty"`
}

func (s UpdatePublishAutoUpgradeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublishAutoUpgradeRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublishAutoUpgradeRequest) GetAutoUpgrade() *int32 {
	return s.AutoUpgrade
}

func (s *UpdatePublishAutoUpgradeRequest) SetAutoUpgrade(v int32) *UpdatePublishAutoUpgradeRequest {
	s.AutoUpgrade = &v
	return s
}

func (s *UpdatePublishAutoUpgradeRequest) Validate() error {
	return dara.Validate(s)
}
