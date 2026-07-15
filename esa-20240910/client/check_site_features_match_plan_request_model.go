// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSiteFeaturesMatchPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewInstanceId(v string) *CheckSiteFeaturesMatchPlanRequest
	GetNewInstanceId() *string
	SetSiteId(v int64) *CheckSiteFeaturesMatchPlanRequest
	GetSiteId() *int64
}

type CheckSiteFeaturesMatchPlanRequest struct {
	// The target instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-site-b0bivjxucjk0
	NewInstanceId *string `json:"NewInstanceId,omitempty" xml:"NewInstanceId,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861405331573200
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CheckSiteFeaturesMatchPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckSiteFeaturesMatchPlanRequest) GoString() string {
	return s.String()
}

func (s *CheckSiteFeaturesMatchPlanRequest) GetNewInstanceId() *string {
	return s.NewInstanceId
}

func (s *CheckSiteFeaturesMatchPlanRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CheckSiteFeaturesMatchPlanRequest) SetNewInstanceId(v string) *CheckSiteFeaturesMatchPlanRequest {
	s.NewInstanceId = &v
	return s
}

func (s *CheckSiteFeaturesMatchPlanRequest) SetSiteId(v int64) *CheckSiteFeaturesMatchPlanRequest {
	s.SiteId = &v
	return s
}

func (s *CheckSiteFeaturesMatchPlanRequest) Validate() error {
	return dara.Validate(s)
}
