// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateProjectFromResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *DissociateProjectFromResourceGroupRequest
	GetProjectId() *int64
	SetResourceGroupId(v string) *DissociateProjectFromResourceGroupRequest
	GetResourceGroupId() *string
}

type DissociateProjectFromResourceGroupRequest struct {
	// The ID of the workspace from which you want to disassociate the resource group.
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

func (s DissociateProjectFromResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateProjectFromResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DissociateProjectFromResourceGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DissociateProjectFromResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DissociateProjectFromResourceGroupRequest) SetProjectId(v int64) *DissociateProjectFromResourceGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *DissociateProjectFromResourceGroupRequest) SetResourceGroupId(v string) *DissociateProjectFromResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DissociateProjectFromResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
