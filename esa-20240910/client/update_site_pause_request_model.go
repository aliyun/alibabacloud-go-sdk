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
	// Specifies whether to temporarily pause the proxy acceleration feature for the entire site. After the feature is paused, all DNS records directly return record values to the client. Valid values:
	//
	// - true: Pauses site acceleration.
	//
	// - false: Resumes normal site acceleration.
	//
	// When site acceleration is paused, only activated sites with NS access mode are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Paused *bool `json:"Paused,omitempty" xml:"Paused,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID. Check the Status field to confirm the site status and the AccessType field to confirm the access mode of the site.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
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
