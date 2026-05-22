// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSitePauseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPaused(v bool) *UpdateSitePauseRequest
	GetPaused() *bool
	SetSiteId(v int64) *UpdateSitePauseRequest
	GetSiteId() *int64
}

type UpdateSitePauseRequest struct {
	// This parameter is required.
	Paused *bool `json:"Paused,omitempty" xml:"Paused,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateSitePauseRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSitePauseRequest) GoString() string {
	return s.String()
}

func (s *UpdateSitePauseRequest) GetPaused() *bool {
	return s.Paused
}

func (s *UpdateSitePauseRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateSitePauseRequest) SetPaused(v bool) *UpdateSitePauseRequest {
	s.Paused = &v
	return s
}

func (s *UpdateSitePauseRequest) SetSiteId(v int64) *UpdateSitePauseRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSitePauseRequest) Validate() error {
	return dara.Validate(s)
}
