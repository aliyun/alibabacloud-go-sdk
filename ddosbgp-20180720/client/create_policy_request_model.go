// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreatePolicyRequest
	GetName() *string
	SetPortVersion(v string) *CreatePolicyRequest
	GetPortVersion() *string
	SetType(v string) *CreatePolicyRequest
	GetType() *string
}

type CreatePolicyRequest struct {
	// The policy name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the port-specific mitigation policy. Valid values:
	//
	// - **Not specified**: creates a default surf DPI engine policy.
	//
	// - **2**: creates a new stream DPI engine policy.
	//
	// > Only port-specific mitigation policies support this parameter.
	//
	// example:
	//
	// 2
	PortVersion *string `json:"PortVersion,omitempty" xml:"PortVersion,omitempty"`
	// The policy type. Valid values:
	//
	// - **l3**: IP-specific mitigation policy.
	//
	// - **l4**: port-specific mitigation policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// l3
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreatePolicyRequest) GetPortVersion() *string {
	return s.PortVersion
}

func (s *CreatePolicyRequest) GetType() *string {
	return s.Type
}

func (s *CreatePolicyRequest) SetName(v string) *CreatePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreatePolicyRequest) SetPortVersion(v string) *CreatePolicyRequest {
	s.PortVersion = &v
	return s
}

func (s *CreatePolicyRequest) SetType(v string) *CreatePolicyRequest {
	s.Type = &v
	return s
}

func (s *CreatePolicyRequest) Validate() error {
	return dara.Validate(s)
}
