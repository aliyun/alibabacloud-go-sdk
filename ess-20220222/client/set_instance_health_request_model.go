// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstanceHealthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHealthStatus(v string) *SetInstanceHealthRequest
	GetHealthStatus() *string
	SetInstanceId(v string) *SetInstanceHealthRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *SetInstanceHealthRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SetInstanceHealthRequest
	GetResourceOwnerAccount() *string
}

type SetInstanceHealthRequest struct {
	// The health status of the instance. Valid values:
	//
	// 	- Healthy: sets the instance as healthy.
	//
	// 	- Unhealthy: sets the instance as unhealthy.
	//
	// This parameter is required.
	//
	// example:
	//
	// Healthy
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1ap6bro51a7fsa****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s SetInstanceHealthRequest) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceHealthRequest) GoString() string {
	return s.String()
}

func (s *SetInstanceHealthRequest) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *SetInstanceHealthRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetInstanceHealthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetInstanceHealthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetInstanceHealthRequest) SetHealthStatus(v string) *SetInstanceHealthRequest {
	s.HealthStatus = &v
	return s
}

func (s *SetInstanceHealthRequest) SetInstanceId(v string) *SetInstanceHealthRequest {
	s.InstanceId = &v
	return s
}

func (s *SetInstanceHealthRequest) SetOwnerId(v int64) *SetInstanceHealthRequest {
	s.OwnerId = &v
	return s
}

func (s *SetInstanceHealthRequest) SetResourceOwnerAccount(v string) *SetInstanceHealthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetInstanceHealthRequest) Validate() error {
	return dara.Validate(s)
}
