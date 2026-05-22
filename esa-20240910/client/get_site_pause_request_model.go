// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSitePauseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetSitePauseRequest
	GetSiteId() *int64
}

type GetSitePauseRequest struct {
	// The website ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSitePauseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSitePauseRequest) GoString() string {
	return s.String()
}

func (s *GetSitePauseRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSitePauseRequest) SetSiteId(v int64) *GetSitePauseRequest {
	s.SiteId = &v
	return s
}

func (s *GetSitePauseRequest) Validate() error {
	return dara.Validate(s)
}
