// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneDefensePoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeSceneDefensePoliciesRequest
	GetResourceGroupId() *string
	SetStatus(v string) *DescribeSceneDefensePoliciesRequest
	GetStatus() *string
	SetTemplate(v string) *DescribeSceneDefensePoliciesRequest
	GetTemplate() *string
}

type DescribeSceneDefensePoliciesRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the policy. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: pending enabling
	//
	// 	- **2**: enabled
	//
	// 	- **3**: expired
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the template that is used to create the policy. Valid values:
	//
	// 	- **promotion**: the Important Activity template
	//
	// 	- **bypass**: the Forward All template
	//
	// example:
	//
	// promotion
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s DescribeSceneDefensePoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneDefensePoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefensePoliciesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSceneDefensePoliciesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSceneDefensePoliciesRequest) GetTemplate() *string {
	return s.Template
}

func (s *DescribeSceneDefensePoliciesRequest) SetResourceGroupId(v string) *DescribeSceneDefensePoliciesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSceneDefensePoliciesRequest) SetStatus(v string) *DescribeSceneDefensePoliciesRequest {
	s.Status = &v
	return s
}

func (s *DescribeSceneDefensePoliciesRequest) SetTemplate(v string) *DescribeSceneDefensePoliciesRequest {
	s.Template = &v
	return s
}

func (s *DescribeSceneDefensePoliciesRequest) Validate() error {
	return dara.Validate(s)
}
