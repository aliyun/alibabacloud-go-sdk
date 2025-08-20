// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstanceEndpointRequest
	GetInstanceId() *string
	SetModuleName(v string) *ListInstanceEndpointRequest
	GetModuleName() *string
	SetSummary(v bool) *ListInstanceEndpointRequest
	GetSummary() *bool
}

type ListInstanceEndpointRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: image repositories.
	//
	// 	- `Chart`: Helm charts.
	//
	// example:
	//
	// Chart
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// Specify whether to return endpoints in simple mode. If yes, no access control information about the endpoint is returned.
	//
	// example:
	//
	// false
	Summary *bool `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListInstanceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceEndpointRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceEndpointRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListInstanceEndpointRequest) GetSummary() *bool {
	return s.Summary
}

func (s *ListInstanceEndpointRequest) SetInstanceId(v string) *ListInstanceEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceEndpointRequest) SetModuleName(v string) *ListInstanceEndpointRequest {
	s.ModuleName = &v
	return s
}

func (s *ListInstanceEndpointRequest) SetSummary(v bool) *ListInstanceEndpointRequest {
	s.Summary = &v
	return s
}

func (s *ListInstanceEndpointRequest) Validate() error {
	return dara.Validate(s)
}
