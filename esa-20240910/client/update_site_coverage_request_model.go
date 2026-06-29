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
	// The target acceleration region to change to. Valid values:
	//
	// - **domestic**: the Chinese mainland only.
	//
	// - **global**: global.
	//
	// - **overseas**: global (excluding the Chinese mainland).
	//
	// <props="china">If the value is set to domestic or global, the site corresponding to SiteId must have a valid ICP filing..
	//
	// This parameter is required.
	//
	// example:
	//
	// overseas
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to query the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
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
