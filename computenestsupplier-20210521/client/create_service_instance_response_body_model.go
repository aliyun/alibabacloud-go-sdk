// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceInstanceResponseBody
	GetRequestId() *string
	SetServiceInstanceId(v string) *CreateServiceInstanceResponseBody
	GetServiceInstanceId() *string
	SetStatus(v string) *CreateServiceInstanceResponseBody
	GetStatus() *string
}

type CreateServiceInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The status of the service instance. Valid values:
	//
	// - Created: The service instance is created.
	//
	// - Deploying: The service instance is being deployed.
	//
	// - DeployedFailed: The service instance failed to be deployed.
	//
	// - Deployed: The service instance is deployed.
	//
	// - Upgrading: The service instance is being upgraded.
	//
	// - Deleting: The service instance is being deleted.
	//
	// - Deleted: The service instance is deleted.
	//
	// - DeletedFailed: The service instance failed to be deleted.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateServiceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceInstanceResponseBody) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *CreateServiceInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateServiceInstanceResponseBody) SetRequestId(v string) *CreateServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetServiceInstanceId(v string) *CreateServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetStatus(v string) *CreateServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
