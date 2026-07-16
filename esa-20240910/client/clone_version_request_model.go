// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *CloneVersionRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CloneVersionRequest
	GetSiteVersion() *int32
}

type CloneVersionRequest struct {
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15846237886****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The site version number to be cloned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CloneVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneVersionRequest) GoString() string {
	return s.String()
}

func (s *CloneVersionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CloneVersionRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CloneVersionRequest) SetSiteId(v int64) *CloneVersionRequest {
	s.SiteId = &v
	return s
}

func (s *CloneVersionRequest) SetSiteVersion(v int32) *CloneVersionRequest {
	s.SiteVersion = &v
	return s
}

func (s *CloneVersionRequest) Validate() error {
	return dara.Validate(s)
}
