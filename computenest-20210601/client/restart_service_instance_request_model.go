// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RestartServiceInstanceRequest
	GetClientToken() *string
	SetRegionId(v string) *RestartServiceInstanceRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *RestartServiceInstanceRequest
	GetServiceInstanceId() *string
}

type RestartServiceInstanceRequest struct {
	// A client token to ensure the idempotence of the request. The token must be unique for each request. The token can contain only ASCII characters and cannot be longer than 64 characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// Call the [ListServiceInstances](https://help.aliyun.com/document_detail/396200.html) operation to obtain the service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s RestartServiceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestartServiceInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartServiceInstanceRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *RestartServiceInstanceRequest) SetClientToken(v string) *RestartServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartServiceInstanceRequest) SetRegionId(v string) *RestartServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestartServiceInstanceRequest) SetServiceInstanceId(v string) *RestartServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *RestartServiceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
