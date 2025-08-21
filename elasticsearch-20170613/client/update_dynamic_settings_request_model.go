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
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Body        *string `json:"body,omitempty" xml:"body,omitempty"`
	Mode        *string `json:"mode,omitempty" xml:"mode,omitempty"`
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
