// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateMonitorGroupInstancesRequest
	GetClientToken() *string
	SetDryRun(v string) *CreateMonitorGroupInstancesRequest
	GetDryRun() *string
	SetGroupId(v string) *CreateMonitorGroupInstancesRequest
	GetGroupId() *string
	SetInstances(v []*CreateMonitorGroupInstancesRequestInstances) *CreateMonitorGroupInstancesRequest
	GetInstances() []*CreateMonitorGroupInstancesRequestInstances
	SetRegionId(v string) *CreateMonitorGroupInstancesRequest
	GetRegionId() *string
}

type CreateMonitorGroupInstancesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// 6868E9A1-56E6-5FD7-9684-E2963BBAFB5F
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only perform a dry run on the request. Valid values:
	//
	// - **true**: sends a check request and does not bind security groups to instances. The check items include whether required parameters are specified, the request format, and service limits. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): sends a regular request. If the check succeeds, an HTTP 2xx status code is returned and the operation is directly performed.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3607****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The list of instances.
	//
	// This parameter is required.
	Instances []*CreateMonitorGroupInstancesRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	RegionId  *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateMonitorGroupInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateMonitorGroupInstancesRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *CreateMonitorGroupInstancesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateMonitorGroupInstancesRequest) GetInstances() []*CreateMonitorGroupInstancesRequestInstances {
	return s.Instances
}

func (s *CreateMonitorGroupInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMonitorGroupInstancesRequest) SetClientToken(v string) *CreateMonitorGroupInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequest) SetDryRun(v string) *CreateMonitorGroupInstancesRequest {
	s.DryRun = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequest) SetGroupId(v string) *CreateMonitorGroupInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequest) SetInstances(v []*CreateMonitorGroupInstancesRequestInstances) *CreateMonitorGroupInstancesRequest {
	s.Instances = v
	return s
}

func (s *CreateMonitorGroupInstancesRequest) SetRegionId(v string) *CreateMonitorGroupInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequest) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateMonitorGroupInstancesRequestInstances struct {
	// The abbreviation of the Alibaba Cloud service name.
	//
	// For more information about how to obtain the abbreviation of the Alibaba Cloud service name, see the `metricCategory` field in the `Labels` response parameter of the [DescribeProjectMeta](https://help.aliyun.com/document_detail/114916.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2ze26xj5wwy12****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-instance-ecs
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateMonitorGroupInstancesRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupInstancesRequestInstances) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupInstancesRequestInstances) GetCategory() *string {
	return s.Category
}

func (s *CreateMonitorGroupInstancesRequestInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateMonitorGroupInstancesRequestInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateMonitorGroupInstancesRequestInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMonitorGroupInstancesRequestInstances) SetCategory(v string) *CreateMonitorGroupInstancesRequestInstances {
	s.Category = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequestInstances) SetInstanceId(v string) *CreateMonitorGroupInstancesRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequestInstances) SetInstanceName(v string) *CreateMonitorGroupInstancesRequestInstances {
	s.InstanceName = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequestInstances) SetRegionId(v string) *CreateMonitorGroupInstancesRequestInstances {
	s.RegionId = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequestInstances) Validate() error {
	return dara.Validate(s)
}
