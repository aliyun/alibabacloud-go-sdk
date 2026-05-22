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
	// The DNS setup. Valid values:
	//
	// 	- **NS**
	//
	// 	- **CNAME**
	//
	// This parameter is required.
	//
	// example:
	//
	// NS
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The service location. Valid values:
	//
	// 	- **domestic**: the Chinese mainland
	//
	// 	- **global**: global
	//
	// 	- **overseas**: outside the Chinese mainland
	//
	// This parameter is required.
	//
	// example:
	//
	// domestic
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The instance ID, which can be obtained by calling the [ListUserRatePlanInstances](https://help.aliyun.com/document_detail/2852398.html) operation. Specify at least one of the instance ID and website ID. If you specify both of them, the instance ID is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbaudit-cn-nwy349jdb03
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the resource group. If you leave this parameter empty, the system uses the default resource group ID.
	//
	// example:
	//
	// rg-acfmw4znnok****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The website name.
	//
	// This parameter is required.
	//
	// example:
	//
	// CreateSite
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
