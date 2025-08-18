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
	// The website ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890***
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
