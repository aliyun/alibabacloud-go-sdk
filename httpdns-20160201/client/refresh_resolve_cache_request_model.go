// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshResolveCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*string) *RefreshResolveCacheRequest
	GetDomains() []*string
}

type RefreshResolveCacheRequest struct {
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s RefreshResolveCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshResolveCacheRequest) GoString() string {
	return s.String()
}

func (s *RefreshResolveCacheRequest) GetDomains() []*string {
	return s.Domains
}

func (s *RefreshResolveCacheRequest) SetDomains(v []*string) *RefreshResolveCacheRequest {
	s.Domains = v
	return s
}

func (s *RefreshResolveCacheRequest) Validate() error {
	return dara.Validate(s)
}
