// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetKeylessServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SetKeylessServerResponseBody
	GetId() *string
	SetRequestId(v string) *SetKeylessServerResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *SetKeylessServerResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *SetKeylessServerResponseBody
	GetSiteName() *string
}

type SetKeylessServerResponseBody struct {
	// Keyless server ID。
	//
	// example:
	//
	// babaabcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s SetKeylessServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetKeylessServerResponseBody) GoString() string {
	return s.String()
}

func (s *SetKeylessServerResponseBody) GetId() *string {
	return s.Id
}

func (s *SetKeylessServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetKeylessServerResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetKeylessServerResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *SetKeylessServerResponseBody) SetId(v string) *SetKeylessServerResponseBody {
	s.Id = &v
	return s
}

func (s *SetKeylessServerResponseBody) SetRequestId(v string) *SetKeylessServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetKeylessServerResponseBody) SetSiteId(v int64) *SetKeylessServerResponseBody {
	s.SiteId = &v
	return s
}

func (s *SetKeylessServerResponseBody) SetSiteName(v string) *SetKeylessServerResponseBody {
	s.SiteName = &v
	return s
}

func (s *SetKeylessServerResponseBody) Validate() error {
	return dara.Validate(s)
}
