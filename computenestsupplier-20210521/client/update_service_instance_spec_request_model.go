// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateServiceInstanceSpecRequest
	GetClientToken() *string
	SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecRequest
	GetEnableUserPrometheus() *bool
	SetOperationName(v string) *UpdateServiceInstanceSpecRequest
	GetOperationName() *string
	SetParameters(v map[string]interface{}) *UpdateServiceInstanceSpecRequest
	GetParameters() map[string]interface{}
	SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecRequest
	GetPredefinedParametersName() *string
	SetServiceInstanceId(v string) *UpdateServiceInstanceSpecRequest
	GetServiceInstanceId() *string
}

type UpdateServiceInstanceSpecRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable Prometheus on the customer side. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The name of the configuration update operation.
	//
	// example:
	//
	// package modify
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The configuration parameters of the service instance.
	//
	// example:
	//
	// {\\"EcsInstanceParameter\\":\\"4vCPU 8GiB\\",\\"ZoneId\\":\\"cn-heyuan-a\\",\\"SystemDiskSize\\":50,\\"DataDiskSize\\":150,\\"InternetMaxBandwidthOut\\":2,\\"RegionId\\":\\"cn-heyuan\\"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The name of the specification package.
	//
	// example:
	//
	// package one
	PredefinedParametersName *string `json:"PredefinedParametersName,omitempty" xml:"PredefinedParametersName,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceInstanceSpecRequest) GetEnableUserPrometheus() *bool {
	return s.EnableUserPrometheus
}

func (s *UpdateServiceInstanceSpecRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *UpdateServiceInstanceSpecRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *UpdateServiceInstanceSpecRequest) GetPredefinedParametersName() *string {
	return s.PredefinedParametersName
}

func (s *UpdateServiceInstanceSpecRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *UpdateServiceInstanceSpecRequest) SetClientToken(v string) *UpdateServiceInstanceSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetOperationName(v string) *UpdateServiceInstanceSpecRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetParameters(v map[string]interface{}) *UpdateServiceInstanceSpecRequest {
	s.Parameters = v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecRequest {
	s.PredefinedParametersName = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceSpecRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) Validate() error {
	return dara.Validate(s)
}
