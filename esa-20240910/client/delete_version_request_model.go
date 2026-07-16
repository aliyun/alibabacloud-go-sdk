// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *DeleteVersionRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *DeleteVersionRequest
	GetSiteVersion() *int32
}

type DeleteVersionRequest struct {
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The site version number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s DeleteVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteVersionRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *DeleteVersionRequest) SetSiteId(v int64) *DeleteVersionRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteVersionRequest) SetSiteVersion(v int32) *DeleteVersionRequest {
	s.SiteVersion = &v
	return s
}

func (s *DeleteVersionRequest) Validate() error {
	return dara.Validate(s)
}
