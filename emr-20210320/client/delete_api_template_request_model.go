// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *DeleteApiTemplateRequest
	GetApiName() *string
	SetRegionId(v string) *DeleteApiTemplateRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteApiTemplateRequest
	GetResourceGroupId() *string
	SetTemplateId(v string) *DeleteApiTemplateRequest
	GetTemplateId() *string
}

type DeleteApiTemplateRequest struct {
	// Interface name.
	//
	// This parameter is required.
	//
	// example:
	//
	// CreateCluster
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// Region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzabjyop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Cluster template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// at-****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteApiTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiTemplateRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DeleteApiTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteApiTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteApiTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteApiTemplateRequest) SetApiName(v string) *DeleteApiTemplateRequest {
	s.ApiName = &v
	return s
}

func (s *DeleteApiTemplateRequest) SetRegionId(v string) *DeleteApiTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApiTemplateRequest) SetResourceGroupId(v string) *DeleteApiTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteApiTemplateRequest) SetTemplateId(v string) *DeleteApiTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteApiTemplateRequest) Validate() error {
	return dara.Validate(s)
}
