// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMatchSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMatchSiteResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *GetMatchSiteResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetMatchSiteResponseBody
	GetSiteName() *string
}

type GetMatchSiteResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-3C82-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The site ID.
	//
	// example:
	//
	// 1234567890****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The site name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s GetMatchSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMatchSiteResponseBody) GoString() string {
	return s.String()
}

func (s *GetMatchSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMatchSiteResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetMatchSiteResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetMatchSiteResponseBody) SetRequestId(v string) *GetMatchSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMatchSiteResponseBody) SetSiteId(v int64) *GetMatchSiteResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetMatchSiteResponseBody) SetSiteName(v string) *GetMatchSiteResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetMatchSiteResponseBody) Validate() error {
	return dara.Validate(s)
}
