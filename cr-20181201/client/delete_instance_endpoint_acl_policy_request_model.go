// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceEndpointAclPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointType(v string) *DeleteInstanceEndpointAclPolicyRequest
	GetEndpointType() *string
	SetEntry(v string) *DeleteInstanceEndpointAclPolicyRequest
	GetEntry() *string
	SetInstanceId(v string) *DeleteInstanceEndpointAclPolicyRequest
	GetInstanceId() *string
	SetModuleName(v string) *DeleteInstanceEndpointAclPolicyRequest
	GetModuleName() *string
}

type DeleteInstanceEndpointAclPolicyRequest struct {
	// The type of the endpoint. Set the value to Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: the image repository.
	//
	// 	- `Chart`: a Helm chart.
	//
	// example:
	//
	// Chart
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DeleteInstanceEndpointAclPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceEndpointAclPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceEndpointAclPolicyRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DeleteInstanceEndpointAclPolicyRequest) GetEntry() *string {
	return s.Entry
}

func (s *DeleteInstanceEndpointAclPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceEndpointAclPolicyRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *DeleteInstanceEndpointAclPolicyRequest) SetEndpointType(v string) *DeleteInstanceEndpointAclPolicyRequest {
	s.EndpointType = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyRequest) SetEntry(v string) *DeleteInstanceEndpointAclPolicyRequest {
	s.Entry = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyRequest) SetInstanceId(v string) *DeleteInstanceEndpointAclPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyRequest) SetModuleName(v string) *DeleteInstanceEndpointAclPolicyRequest {
	s.ModuleName = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyRequest) Validate() error {
	return dara.Validate(s)
}
