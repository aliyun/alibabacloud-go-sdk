// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceStop(v bool) *RedeployInstanceRequest
	GetForceStop() *bool
	SetInstanceId(v string) *RedeployInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *RedeployInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RedeployInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RedeployInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RedeployInstanceRequest
	GetResourceOwnerId() *int64
}

type RedeployInstanceRequest struct {
	// Specifies whether to forcefully stop the instance that is in the Running state.
	//
	// Default value: false.
	//
	// > A forced stop is equivalent to a typical server power-off. Data in the instance operating system that has not been written to storage devices is lost. We recommend that you redeploy instances that are already stopped.
	//
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The instance ID of the instance that is in the Running or Stopped state.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1azkttqpldxgted****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RedeployInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RedeployInstanceRequest) GoString() string {
	return s.String()
}

func (s *RedeployInstanceRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *RedeployInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RedeployInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RedeployInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RedeployInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RedeployInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RedeployInstanceRequest) SetForceStop(v bool) *RedeployInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *RedeployInstanceRequest) SetInstanceId(v string) *RedeployInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RedeployInstanceRequest) SetOwnerAccount(v string) *RedeployInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RedeployInstanceRequest) SetOwnerId(v int64) *RedeployInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RedeployInstanceRequest) SetResourceOwnerAccount(v string) *RedeployInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RedeployInstanceRequest) SetResourceOwnerId(v int64) *RedeployInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RedeployInstanceRequest) Validate() error {
	return dara.Validate(s)
}
