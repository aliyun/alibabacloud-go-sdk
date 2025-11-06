// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupAssociateProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *ListResourceGroupAssociateProjectsRequest
	GetResourceGroupId() *string
}

type ListResourceGroupAssociateProjectsRequest struct {
	// The ID of the serverless resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListResourceGroupAssociateProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupAssociateProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupAssociateProjectsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListResourceGroupAssociateProjectsRequest) SetResourceGroupId(v string) *ListResourceGroupAssociateProjectsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourceGroupAssociateProjectsRequest) Validate() error {
	return dara.Validate(s)
}
