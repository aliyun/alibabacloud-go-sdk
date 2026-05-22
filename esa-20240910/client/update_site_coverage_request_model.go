// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteCoverageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverage(v string) *UpdateSiteCoverageRequest
	GetCoverage() *string
	SetSiteId(v int64) *UpdateSiteCoverageRequest
	GetSiteId() *int64
}

type UpdateSiteCoverageRequest struct {
	// This parameter is required.
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateSiteCoverageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteCoverageRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteCoverageRequest) GetCoverage() *string {
	return s.Coverage
}

func (s *UpdateSiteCoverageRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateSiteCoverageRequest) SetCoverage(v string) *UpdateSiteCoverageRequest {
	s.Coverage = &v
	return s
}

func (s *UpdateSiteCoverageRequest) SetSiteId(v int64) *UpdateSiteCoverageRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSiteCoverageRequest) Validate() error {
	return dara.Validate(s)
}
