// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshResolveCacheShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainsShrink(v string) *RefreshResolveCacheShrinkRequest
	GetDomainsShrink() *string
}

type RefreshResolveCacheShrinkRequest struct {
	DomainsShrink *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
}

func (s RefreshResolveCacheShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshResolveCacheShrinkRequest) GoString() string {
	return s.String()
}

func (s *RefreshResolveCacheShrinkRequest) GetDomainsShrink() *string {
	return s.DomainsShrink
}

func (s *RefreshResolveCacheShrinkRequest) SetDomainsShrink(v string) *RefreshResolveCacheShrinkRequest {
	s.DomainsShrink = &v
	return s
}

func (s *RefreshResolveCacheShrinkRequest) Validate() error {
	return dara.Validate(s)
}
