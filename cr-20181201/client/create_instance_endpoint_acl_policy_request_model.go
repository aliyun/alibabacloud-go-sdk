// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceEndpointAclPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateInstanceEndpointAclPolicyRequest
	GetComment() *string
	SetEndpointType(v string) *CreateInstanceEndpointAclPolicyRequest
	GetEndpointType() *string
	SetEntry(v string) *CreateInstanceEndpointAclPolicyRequest
	GetEntry() *string
	SetInstanceId(v string) *CreateInstanceEndpointAclPolicyRequest
	GetInstanceId() *string
	SetModuleName(v string) *CreateInstanceEndpointAclPolicyRequest
	GetModuleName() *string
}

type CreateInstanceEndpointAclPolicyRequest struct {
	// The description.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The type of the endpoint. Set the value to Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The CIDR block that is accessible.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.1/32
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
	// Registry
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s CreateInstanceEndpointAclPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceEndpointAclPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceEndpointAclPolicyRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateInstanceEndpointAclPolicyRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateInstanceEndpointAclPolicyRequest) GetEntry() *string {
	return s.Entry
}

func (s *CreateInstanceEndpointAclPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceEndpointAclPolicyRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetComment(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.Comment = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetEndpointType(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetEntry(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.Entry = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetInstanceId(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetModuleName(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.ModuleName = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) Validate() error {
	return dara.Validate(s)
}
