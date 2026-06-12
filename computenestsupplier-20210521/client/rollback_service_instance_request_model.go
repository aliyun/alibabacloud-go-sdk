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
	// The client token that is used to ensure the idempotence of the request. Generate a unique value on your client for each request. The **ClientToken*	- can contain only ASCII characters and must be no more than 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service instance ID.
	//
	// Call [ListServiceInstances](https://help.aliyun.com/document_detail/396200.html) to obtain the service instance ID.
	//
	// example:
	//
	// si-3a8f9a75da074f52b969
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
