// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSiteNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteName(v string) *CheckSiteNameRequest
	GetSiteName() *string
}

type CheckSiteNameRequest struct {
	// The website name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s CheckSiteNameRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckSiteNameRequest) GoString() string {
	return s.String()
}

func (s *CheckSiteNameRequest) GetSiteName() *string {
	return s.SiteName
}

func (s *CheckSiteNameRequest) SetSiteName(v string) *CheckSiteNameRequest {
	s.SiteName = &v
	return s
}

func (s *CheckSiteNameRequest) Validate() error {
	return dara.Validate(s)
}
