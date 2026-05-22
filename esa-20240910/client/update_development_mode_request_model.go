// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDevelopmentModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *UpdateDevelopmentModeRequest
	GetEnable() *string
	SetSiteId(v int64) *UpdateDevelopmentModeRequest
	GetSiteId() *int64
}

type UpdateDevelopmentModeRequest struct {
	// This parameter is required.
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateDevelopmentModeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDevelopmentModeRequest) GoString() string {
	return s.String()
}

func (s *UpdateDevelopmentModeRequest) GetEnable() *string {
	return s.Enable
}

func (s *UpdateDevelopmentModeRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateDevelopmentModeRequest) SetEnable(v string) *UpdateDevelopmentModeRequest {
	s.Enable = &v
	return s
}

func (s *UpdateDevelopmentModeRequest) SetSiteId(v int64) *UpdateDevelopmentModeRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateDevelopmentModeRequest) Validate() error {
	return dara.Validate(s)
}
