// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppCodeWorkspaceDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v string) *GetAppCodeWorkspaceDetailRequest
	GetSiteId() *string
}

type GetAppCodeWorkspaceDetailRequest struct {
	// Site ID
	//
	// example:
	//
	// 1071596645435968
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetAppCodeWorkspaceDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppCodeWorkspaceDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAppCodeWorkspaceDetailRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *GetAppCodeWorkspaceDetailRequest) SetSiteId(v string) *GetAppCodeWorkspaceDetailRequest {
	s.SiteId = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailRequest) Validate() error {
	return dara.Validate(s)
}
