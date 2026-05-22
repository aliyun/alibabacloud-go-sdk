// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteOriginPoolRequest
	GetId() *int64
	SetSiteId(v int64) *DeleteOriginPoolRequest
	GetSiteId() *int64
}

type DeleteOriginPoolRequest struct {
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteOriginPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginPoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteOriginPoolRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteOriginPoolRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteOriginPoolRequest) SetId(v int64) *DeleteOriginPoolRequest {
	s.Id = &v
	return s
}

func (s *DeleteOriginPoolRequest) SetSiteId(v int64) *DeleteOriginPoolRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteOriginPoolRequest) Validate() error {
	return dara.Validate(s)
}
