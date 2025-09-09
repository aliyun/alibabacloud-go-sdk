// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyServiceInstanceResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResources(v string) *ModifyServiceInstanceResourcesRequest
	GetResources() *string
	SetServiceInstanceId(v string) *ModifyServiceInstanceResourcesRequest
	GetServiceInstanceId() *string
	SetServiceInstanceResourcesAction(v string) *ModifyServiceInstanceResourcesRequest
	GetServiceInstanceResourcesAction() *string
}

type ModifyServiceInstanceResourcesRequest struct {
	// The imported resources.
	//
	// example:
	//
	// {
	//
	//   "RegionId": "cn-hangzhou",
	//
	//   "Type": "ResourceIds",
	//
	//   "ResourceIds": {
	//
	//     "ALIYUN::ECS::INSTANCE": ["i-xxx", "i-yyy"],
	//
	//     "ALIYUN::RDS::INSTANCE": ["rm-xxx", "rm-yyy"],
	//
	//     "ALIYUN::VPC::VPC": ["vpc-xxx", "vpc-yyy"],
	//
	//     "ALIYUN::SLB::INSTANCE": ["lb-xxx", "lb-yyy"]
	//
	//   }
	//
	// }
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d8a0cc2a1ee04dce****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The type of operation performed on the service instance resource. Valid values:
	//
	// 	- Import: The resource is imported.
	//
	// 	- UnImport: The resource import is canceled.
	//
	// example:
	//
	// Import
	ServiceInstanceResourcesAction *string `json:"ServiceInstanceResourcesAction,omitempty" xml:"ServiceInstanceResourcesAction,omitempty"`
}

func (s ModifyServiceInstanceResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyServiceInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *ModifyServiceInstanceResourcesRequest) GetResources() *string {
	return s.Resources
}

func (s *ModifyServiceInstanceResourcesRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *ModifyServiceInstanceResourcesRequest) GetServiceInstanceResourcesAction() *string {
	return s.ServiceInstanceResourcesAction
}

func (s *ModifyServiceInstanceResourcesRequest) SetResources(v string) *ModifyServiceInstanceResourcesRequest {
	s.Resources = &v
	return s
}

func (s *ModifyServiceInstanceResourcesRequest) SetServiceInstanceId(v string) *ModifyServiceInstanceResourcesRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ModifyServiceInstanceResourcesRequest) SetServiceInstanceResourcesAction(v string) *ModifyServiceInstanceResourcesRequest {
	s.ServiceInstanceResourcesAction = &v
	return s
}

func (s *ModifyServiceInstanceResourcesRequest) Validate() error {
	return dara.Validate(s)
}
