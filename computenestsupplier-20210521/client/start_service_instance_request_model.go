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
	// A client token that is used to ensure the idempotence of the request. Generate a value for this parameter from your client and make sure that the value is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9****
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
	// You can call the [ListServiceInstances](https://help.aliyun.com/document_detail/396200.html) operation to query the IDs of service instances.
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
