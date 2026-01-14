// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceManagedControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateServiceManagedControlRequest
	GetClientToken() *string
	SetRegionId(v string) *UpdateServiceManagedControlRequest
	GetRegionId() *string
	SetResourceId(v string) *UpdateServiceManagedControlRequest
	GetResourceId() *string
	SetResourceType(v string) *UpdateServiceManagedControlRequest
	GetResourceType() *string
	SetServiceManaged(v bool) *UpdateServiceManagedControlRequest
	GetServiceManaged() *bool
}

type UpdateServiceManagedControlRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the value of **RequestId*	- as the value of **ClientToken**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the GA instance. Set the value to cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource whose control mode you want to change.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp149u6o36qt1as9b****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource whose control mode you want to change.
	//
	// Set the value to **Accelerator**, which specifies a standard GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// Accelerator
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The control mode of the resource.
	//
	// Set the value to **false**, which changes the control mode of the resource from managed mode to unmanaged mode.
	//
	// >  You can change the control mode only from managed mode to unmanaged mode.
	//
	// example:
	//
	// false
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
}

func (s UpdateServiceManagedControlRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceManagedControlRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceManagedControlRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceManagedControlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateServiceManagedControlRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateServiceManagedControlRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateServiceManagedControlRequest) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *UpdateServiceManagedControlRequest) SetClientToken(v string) *UpdateServiceManagedControlRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceManagedControlRequest) SetRegionId(v string) *UpdateServiceManagedControlRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceManagedControlRequest) SetResourceId(v string) *UpdateServiceManagedControlRequest {
	s.ResourceId = &v
	return s
}

func (s *UpdateServiceManagedControlRequest) SetResourceType(v string) *UpdateServiceManagedControlRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateServiceManagedControlRequest) SetServiceManaged(v bool) *UpdateServiceManagedControlRequest {
	s.ServiceManaged = &v
	return s
}

func (s *UpdateServiceManagedControlRequest) Validate() error {
	return dara.Validate(s)
}
