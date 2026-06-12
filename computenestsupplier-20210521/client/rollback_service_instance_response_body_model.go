// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackServiceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RollbackServiceInstanceResponseBody
	GetRequestId() *string
	SetServiceInstanceId(v string) *RollbackServiceInstanceResponseBody
	GetServiceInstanceId() *string
	SetStatus(v string) *RollbackServiceInstanceResponseBody
	GetStatus() *string
}

type RollbackServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EE9EF87D-46F8-5AF6-9A65-6B034E204136
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-5289e1d6d0c14397881d
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The deployment status of the service instance. Valid values:
	//
	// - Created: The service instance is created.
	//
	// - Deploying: The service instance is being deployed.
	//
	// - DeployedFailed: The deployment of the service instance failed.
	//
	// - Deployed: The service instance is deployed.
	//
	// - Upgrading: The service instance is being upgraded.
	//
	// - UpgradeRollbacking: The service instance is being rolled back.
	//
	// - Deleting: The service instance is being deleted.
	//
	// - Deleted: The service instance is deleted.
	//
	// - DeletedFailed: The deletion of the service instance failed.
	//
	// example:
	//
	// UpgradeRollbacking
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RollbackServiceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackServiceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackServiceInstanceResponseBody) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *RollbackServiceInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *RollbackServiceInstanceResponseBody) SetRequestId(v string) *RollbackServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackServiceInstanceResponseBody) SetServiceInstanceId(v string) *RollbackServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *RollbackServiceInstanceResponseBody) SetStatus(v string) *RollbackServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *RollbackServiceInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
