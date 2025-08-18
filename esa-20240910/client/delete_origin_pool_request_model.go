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
	// The ID of the origin address pool, which can be obtained by calling the [ListOriginPools](https://help.aliyun.com/document_detail/2863947.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 103852052519****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The site ID, which can be obtained by calling the [ListSites](~~ListSites~~) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21655860979****
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
