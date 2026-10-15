// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceConfigShrinkRequest
	GetInstanceId() *string
	SetReason(v string) *ModifyInstanceConfigShrinkRequest
	GetReason() *string
	SetUserConfig(v string) *ModifyInstanceConfigShrinkRequest
	GetUserConfig() *string
	SetBodyShrink(v string) *ModifyInstanceConfigShrinkRequest
	GetBodyShrink() *string
}

type ModifyInstanceConfigShrinkRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The reason for the update.
	//
	// This parameter is required.
	//
	// example:
	//
	// for test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The custom configuration.
	//
	// example:
	//
	// dataCoord:\\n  segment:\\n    maxSize: 1024
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
	// Submits the instance configurations to be updated as a JSON object.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceConfigShrinkRequest) GetReason() *string {
	return s.Reason
}

func (s *ModifyInstanceConfigShrinkRequest) GetUserConfig() *string {
	return s.UserConfig
}

func (s *ModifyInstanceConfigShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *ModifyInstanceConfigShrinkRequest) SetInstanceId(v string) *ModifyInstanceConfigShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceConfigShrinkRequest) SetReason(v string) *ModifyInstanceConfigShrinkRequest {
	s.Reason = &v
	return s
}

func (s *ModifyInstanceConfigShrinkRequest) SetUserConfig(v string) *ModifyInstanceConfigShrinkRequest {
	s.UserConfig = &v
	return s
}

func (s *ModifyInstanceConfigShrinkRequest) SetBodyShrink(v string) *ModifyInstanceConfigShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *ModifyInstanceConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
