// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetOriginPoolRequest
	GetId() *int64
	SetSiteId(v int64) *GetOriginPoolRequest
	GetSiteId() *int64
}

type GetOriginPoolRequest struct {
	// The ID of the origin pool, which can be obtained by calling the [ListOriginPools](https://help.aliyun.com/document_detail/2863947.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1038520525196928
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the site, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 216558609793952
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetOriginPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOriginPoolRequest) GoString() string {
	return s.String()
}

func (s *GetOriginPoolRequest) GetId() *int64 {
	return s.Id
}

func (s *GetOriginPoolRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetOriginPoolRequest) SetId(v int64) *GetOriginPoolRequest {
	s.Id = &v
	return s
}

func (s *GetOriginPoolRequest) SetSiteId(v int64) *GetOriginPoolRequest {
	s.SiteId = &v
	return s
}

func (s *GetOriginPoolRequest) Validate() error {
	return dara.Validate(s)
}
