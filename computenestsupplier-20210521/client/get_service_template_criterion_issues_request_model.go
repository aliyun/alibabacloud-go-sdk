// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTemplateCriterionIssuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetServiceTemplateCriterionIssuesRequest
	GetRegionId() *string
	SetServiceId(v string) *GetServiceTemplateCriterionIssuesRequest
	GetServiceId() *string
	SetServiceVersion(v string) *GetServiceTemplateCriterionIssuesRequest
	GetServiceVersion() *string
}

type GetServiceTemplateCriterionIssuesRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-c2d118c9193e49xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s GetServiceTemplateCriterionIssuesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateCriterionIssuesRequest) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateCriterionIssuesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceTemplateCriterionIssuesRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceTemplateCriterionIssuesRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetServiceTemplateCriterionIssuesRequest) SetRegionId(v string) *GetServiceTemplateCriterionIssuesRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesRequest) SetServiceId(v string) *GetServiceTemplateCriterionIssuesRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesRequest) SetServiceVersion(v string) *GetServiceTemplateCriterionIssuesRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesRequest) Validate() error {
	return dara.Validate(s)
}
