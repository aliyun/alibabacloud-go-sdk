// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *ListVersionsRequest
	GetSiteId() *int64
}

type ListVersionsRequest struct {
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 33968830844****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListVersionsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListVersionsRequest) SetSiteId(v int64) *ListVersionsRequest {
	s.SiteId = &v
	return s
}

func (s *ListVersionsRequest) Validate() error {
	return dara.Validate(s)
}
