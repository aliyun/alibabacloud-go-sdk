// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartServiceInstanceRequest
	GetClientToken() *string
	SetRegionId(v string) *StartServiceInstanceRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *StartServiceInstanceRequest
	GetServiceInstanceId() *string
}

type StartServiceInstanceRequest struct {
	// A client token to ensure the idempotence of the request. Generate a unique value on your client. This value can contain only ASCII characters and must be no more than 64 characters long.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the region where the service instance resides.
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

func (s StartServiceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartServiceInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartServiceInstanceRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *StartServiceInstanceRequest) SetClientToken(v string) *StartServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartServiceInstanceRequest) SetRegionId(v string) *StartServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartServiceInstanceRequest) SetServiceInstanceId(v string) *StartServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *StartServiceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
