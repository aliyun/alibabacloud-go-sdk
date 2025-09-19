// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublishGraySwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGraySwitchStatus(v int32) *UpdatePublishGraySwitchRequest
	GetGraySwitchStatus() *int32
}

type UpdatePublishGraySwitchRequest struct {
	// Specifies whether to enable the canary release feature. Valid values:
	//
	// 	- **1**: enabled.
	//
	// 	- **0**: disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	GraySwitchStatus *int32 `json:"GraySwitchStatus,omitempty" xml:"GraySwitchStatus,omitempty"`
}

func (s UpdatePublishGraySwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublishGraySwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublishGraySwitchRequest) GetGraySwitchStatus() *int32 {
	return s.GraySwitchStatus
}

func (s *UpdatePublishGraySwitchRequest) SetGraySwitchStatus(v int32) *UpdatePublishGraySwitchRequest {
	s.GraySwitchStatus = &v
	return s
}

func (s *UpdatePublishGraySwitchRequest) Validate() error {
	return dara.Validate(s)
}
