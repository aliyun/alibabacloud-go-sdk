// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeylessServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteKeylessServerResponseBody
	GetId() *string
	SetRequestId(v string) *DeleteKeylessServerResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *DeleteKeylessServerResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *DeleteKeylessServerResponseBody
	GetSiteName() *string
}

type DeleteKeylessServerResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId    *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName  *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s DeleteKeylessServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeylessServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeylessServerResponseBody) GetId() *string {
	return s.Id
}

func (s *DeleteKeylessServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKeylessServerResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteKeylessServerResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *DeleteKeylessServerResponseBody) SetId(v string) *DeleteKeylessServerResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteKeylessServerResponseBody) SetRequestId(v string) *DeleteKeylessServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKeylessServerResponseBody) SetSiteId(v int64) *DeleteKeylessServerResponseBody {
	s.SiteId = &v
	return s
}

func (s *DeleteKeylessServerResponseBody) SetSiteName(v string) *DeleteKeylessServerResponseBody {
	s.SiteName = &v
	return s
}

func (s *DeleteKeylessServerResponseBody) Validate() error {
	return dara.Validate(s)
}
