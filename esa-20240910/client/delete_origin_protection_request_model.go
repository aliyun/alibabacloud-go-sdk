// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *DeleteOriginProtectionRequest
	GetSiteId() *int64
}

type DeleteOriginProtectionRequest struct {
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteOriginProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginProtectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteOriginProtectionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteOriginProtectionRequest) SetSiteId(v int64) *DeleteOriginProtectionRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteOriginProtectionRequest) Validate() error {
	return dara.Validate(s)
}
