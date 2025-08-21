// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStsGrantStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeStsGrantStatusRequest
	GetResourceGroupId() *string
	SetRole(v string) *DescribeStsGrantStatusRequest
	GetRole() *string
}

type DescribeStsGrantStatusRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the RAM role to query. Set the value to **AliyunDDoSCOODefaultRole**, which indicates the default role of Anti-DDoS Pro or Anti-DDoS Premium.
	//
	// > Anti-DDoS Pro or Anti-DDoS Premium uses the default role to access other cloud services.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunDDoSCOODefaultRole
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeStsGrantStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStsGrantStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeStsGrantStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeStsGrantStatusRequest) GetRole() *string {
	return s.Role
}

func (s *DescribeStsGrantStatusRequest) SetResourceGroupId(v string) *DescribeStsGrantStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeStsGrantStatusRequest) SetRole(v string) *DescribeStsGrantStatusRequest {
	s.Role = &v
	return s
}

func (s *DescribeStsGrantStatusRequest) Validate() error {
	return dara.Validate(s)
}
