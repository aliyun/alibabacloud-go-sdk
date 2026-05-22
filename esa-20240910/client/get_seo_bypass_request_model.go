// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSeoBypassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetSeoBypassRequest
	GetSiteId() *int64
}

type GetSeoBypassRequest struct {
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSeoBypassRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSeoBypassRequest) GoString() string {
	return s.String()
}

func (s *GetSeoBypassRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSeoBypassRequest) SetSiteId(v int64) *GetSeoBypassRequest {
	s.SiteId = &v
	return s
}

func (s *GetSeoBypassRequest) Validate() error {
	return dara.Validate(s)
}
