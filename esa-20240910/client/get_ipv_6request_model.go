// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIPv6Request interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetIPv6Request
	GetSiteId() *int64
}

type GetIPv6Request struct {
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetIPv6Request) String() string {
	return dara.Prettify(s)
}

func (s GetIPv6Request) GoString() string {
	return s.String()
}

func (s *GetIPv6Request) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetIPv6Request) SetSiteId(v int64) *GetIPv6Request {
	s.SiteId = &v
	return s
}

func (s *GetIPv6Request) Validate() error {
	return dara.Validate(s)
}
