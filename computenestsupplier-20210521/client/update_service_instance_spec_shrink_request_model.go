// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateServiceInstanceSpecShrinkRequest
	GetClientToken() *string
	SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecShrinkRequest
	GetEnableUserPrometheus() *bool
	SetOperationName(v string) *UpdateServiceInstanceSpecShrinkRequest
	GetOperationName() *string
	SetParametersShrink(v string) *UpdateServiceInstanceSpecShrinkRequest
	GetParametersShrink() *string
	SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecShrinkRequest
	GetPredefinedParametersName() *string
	SetServiceInstanceId(v string) *UpdateServiceInstanceSpecShrinkRequest
	GetServiceInstanceId() *string
}

type UpdateServiceInstanceSpecShrinkRequest struct {
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
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s UpdateServiceInstanceSpecShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetEnableUserPrometheus() *bool {
	return s.EnableUserPrometheus
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetPredefinedParametersName() *string {
	return s.PredefinedParametersName
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetClientToken(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecShrinkRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetOperationName(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetParametersShrink(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.PredefinedParametersName = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) Validate() error {
	return dara.Validate(s)
}
