// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateVersionManagementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *DeactivateVersionManagementRequest
	GetSiteId() *int64
}

type DeactivateVersionManagementRequest struct {
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeactivateVersionManagementRequest) String() string {
	return dara.Prettify(s)
}

func (s DeactivateVersionManagementRequest) GoString() string {
	return s.String()
}

func (s *DeactivateVersionManagementRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeactivateVersionManagementRequest) SetSiteId(v int64) *DeactivateVersionManagementRequest {
	s.SiteId = &v
	return s
}

func (s *DeactivateVersionManagementRequest) Validate() error {
	return dara.Validate(s)
}
