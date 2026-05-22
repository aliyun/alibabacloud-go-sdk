// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteAccessTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *UpdateSiteAccessTypeRequest
	GetAccessType() *string
	SetSiteId(v int64) *UpdateSiteAccessTypeRequest
	GetSiteId() *int64
}

type UpdateSiteAccessTypeRequest struct {
	// This parameter is required.
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateSiteAccessTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteAccessTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteAccessTypeRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *UpdateSiteAccessTypeRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateSiteAccessTypeRequest) SetAccessType(v string) *UpdateSiteAccessTypeRequest {
	s.AccessType = &v
	return s
}

func (s *UpdateSiteAccessTypeRequest) SetSiteId(v int64) *UpdateSiteAccessTypeRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSiteAccessTypeRequest) Validate() error {
	return dara.Validate(s)
}
