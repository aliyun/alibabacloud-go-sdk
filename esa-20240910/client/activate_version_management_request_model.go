// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateVersionManagementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *ActivateVersionManagementRequest
	GetSiteId() *int64
}

type ActivateVersionManagementRequest struct {
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11223***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ActivateVersionManagementRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateVersionManagementRequest) GoString() string {
	return s.String()
}

func (s *ActivateVersionManagementRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ActivateVersionManagementRequest) SetSiteId(v int64) *ActivateVersionManagementRequest {
	s.SiteId = &v
	return s
}

func (s *ActivateVersionManagementRequest) Validate() error {
	return dara.Validate(s)
}
