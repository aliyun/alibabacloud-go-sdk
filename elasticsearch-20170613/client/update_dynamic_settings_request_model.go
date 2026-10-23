// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDynamicSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateDynamicSettingsRequest
	GetClientToken() *string
	SetRegionId(v string) *UpdateDynamicSettingsRequest
	GetRegionId() *string
	SetBody(v string) *UpdateDynamicSettingsRequest
	GetBody() *string
	SetMode(v string) *UpdateDynamicSettingsRequest
	GetMode() *string
}

type UpdateDynamicSettingsRequest struct {
	// The client token that is used to ensure idempotence.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the Elasticsearch instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The dynamic configurations of the cluster to be updated. Specify a JSON object where the keys are configuration item names and the values are configuration values.
	//
	// example:
	//
	// {"indices.recovery.max_bytes_per_sec":"100mb"}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The effective mode for dynamic configurations. Valid values:
	//
	// - persistent: persistent configurations.
	//
	// - transient: temporary configurations, which become invalid after the cluster is restarted.
	//
	// example:
	//
	// persistent
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s UpdateDynamicSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDynamicSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateDynamicSettingsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateDynamicSettingsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDynamicSettingsRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateDynamicSettingsRequest) GetMode() *string {
	return s.Mode
}

func (s *UpdateDynamicSettingsRequest) SetClientToken(v string) *UpdateDynamicSettingsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateDynamicSettingsRequest) SetRegionId(v string) *UpdateDynamicSettingsRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDynamicSettingsRequest) SetBody(v string) *UpdateDynamicSettingsRequest {
	s.Body = &v
	return s
}

func (s *UpdateDynamicSettingsRequest) SetMode(v string) *UpdateDynamicSettingsRequest {
	s.Mode = &v
	return s
}

func (s *UpdateDynamicSettingsRequest) Validate() error {
	return dara.Validate(s)
}
