// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *ListEnvironmentsRequest
	GetSiteId() *int64
}

type ListEnvironmentsRequest struct {
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListEnvironmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListEnvironmentsRequest) SetSiteId(v int64) *ListEnvironmentsRequest {
	s.SiteId = &v
	return s
}

func (s *ListEnvironmentsRequest) Validate() error {
	return dara.Validate(s)
}
