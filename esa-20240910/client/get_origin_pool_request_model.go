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
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
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
