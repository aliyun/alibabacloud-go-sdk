// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RollbackServiceInstanceRequest
	GetClientToken() *string
	SetRegionId(v string) *RollbackServiceInstanceRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *RollbackServiceInstanceRequest
	GetServiceInstanceId() *string
}

type RollbackServiceInstanceRequest struct {
	// A client token to ensure request idempotence. Generate a unique value for this parameter from your client. **ClientToken*	- supports only ASCII characters and cannot be more than 64 characters long.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-d6ab3a63ccbb4bxxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s RollbackServiceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *RollbackServiceInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RollbackServiceInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RollbackServiceInstanceRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *RollbackServiceInstanceRequest) SetClientToken(v string) *RollbackServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RollbackServiceInstanceRequest) SetRegionId(v string) *RollbackServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RollbackServiceInstanceRequest) SetServiceInstanceId(v string) *RollbackServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *RollbackServiceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
