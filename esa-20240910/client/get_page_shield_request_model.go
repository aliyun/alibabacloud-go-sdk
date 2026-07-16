// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageShieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetPageShieldRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *GetPageShieldRequest
	GetSiteVersion() *int32
}

type GetPageShieldRequest struct {
	// The site ID. You can obtain the ID by calling [ListSites](~~ListSites~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site. For sites with version management enabled, you can use this parameter to specify the site version on which the configuration takes effect. The default value is 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetPageShieldRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPageShieldRequest) GoString() string {
	return s.String()
}

func (s *GetPageShieldRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetPageShieldRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetPageShieldRequest) SetSiteId(v int64) *GetPageShieldRequest {
	s.SiteId = &v
	return s
}

func (s *GetPageShieldRequest) SetSiteVersion(v int32) *GetPageShieldRequest {
	s.SiteVersion = &v
	return s
}

func (s *GetPageShieldRequest) Validate() error {
	return dara.Validate(s)
}
