// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateProjectToResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *AssociateProjectToResourceGroupRequest
	GetProjectId() *int64
	SetResourceGroupId(v string) *AssociateProjectToResourceGroupRequest
	GetResourceGroupId() *string
}

type AssociateProjectToResourceGroupRequest struct {
	// The ID of the DataWorks workspace with which you want to associate the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AssociateProjectToResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateProjectToResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *AssociateProjectToResourceGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AssociateProjectToResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AssociateProjectToResourceGroupRequest) SetProjectId(v int64) *AssociateProjectToResourceGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *AssociateProjectToResourceGroupRequest) SetResourceGroupId(v string) *AssociateProjectToResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AssociateProjectToResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
