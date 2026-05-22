// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *CreateSiteRequest
	GetAccessType() *string
	SetCoverage(v string) *CreateSiteRequest
	GetCoverage() *string
	SetInstanceId(v string) *CreateSiteRequest
	GetInstanceId() *string
	SetResourceGroupId(v string) *CreateSiteRequest
	GetResourceGroupId() *string
	SetSiteName(v string) *CreateSiteRequest
	GetSiteName() *string
}

type CreateSiteRequest struct {
	// This parameter is required.
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// This parameter is required.
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// This parameter is required.
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 记录名称
	//
	// This parameter is required.
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s CreateSiteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteRequest) GoString() string {
	return s.String()
}

func (s *CreateSiteRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *CreateSiteRequest) GetCoverage() *string {
	return s.Coverage
}

func (s *CreateSiteRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSiteRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSiteRequest) GetSiteName() *string {
	return s.SiteName
}

func (s *CreateSiteRequest) SetAccessType(v string) *CreateSiteRequest {
	s.AccessType = &v
	return s
}

func (s *CreateSiteRequest) SetCoverage(v string) *CreateSiteRequest {
	s.Coverage = &v
	return s
}

func (s *CreateSiteRequest) SetInstanceId(v string) *CreateSiteRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSiteRequest) SetResourceGroupId(v string) *CreateSiteRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSiteRequest) SetSiteName(v string) *CreateSiteRequest {
	s.SiteName = &v
	return s
}

func (s *CreateSiteRequest) Validate() error {
	return dara.Validate(s)
}
