// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteTrafficSequenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetSiteTrafficSequenceRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *GetSiteTrafficSequenceRequest
	GetSiteVersion() *int32
}

type GetSiteTrafficSequenceRequest struct {
	// The site ID. You can obtain the site ID by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site. After version management is enabled for the site, you can specify a site version number to obtain the traffic sequence information of the corresponding version. The default version is 0. If version management is not enabled for the site, you do not need to specify this parameter.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetSiteTrafficSequenceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSiteTrafficSequenceRequest) GoString() string {
	return s.String()
}

func (s *GetSiteTrafficSequenceRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteTrafficSequenceRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetSiteTrafficSequenceRequest) SetSiteId(v int64) *GetSiteTrafficSequenceRequest {
	s.SiteId = &v
	return s
}

func (s *GetSiteTrafficSequenceRequest) SetSiteVersion(v int32) *GetSiteTrafficSequenceRequest {
	s.SiteVersion = &v
	return s
}

func (s *GetSiteTrafficSequenceRequest) Validate() error {
	return dara.Validate(s)
}
