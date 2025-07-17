// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdatePolicyRequest
	GetConfig() *string
	SetDescription(v string) *UpdatePolicyRequest
	GetDescription() *string
	SetName(v string) *UpdatePolicyRequest
	GetName() *string
}

type UpdatePolicyRequest struct {
	// Policy configuration
	//
	// This parameter is required.
	//
	// example:
	//
	// {"unitNum":1,"timeUnit":"s","enable":true}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// Description
	//
	// example:
	//
	// this is a timeout policy description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Policy name
	//
	// This parameter is required.
	//
	// example:
	//
	// celue-timeout-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdatePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePolicyRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePolicyRequest) SetConfig(v string) *UpdatePolicyRequest {
	s.Config = &v
	return s
}

func (s *UpdatePolicyRequest) SetDescription(v string) *UpdatePolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdatePolicyRequest) SetName(v string) *UpdatePolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdatePolicyRequest) Validate() error {
	return dara.Validate(s)
}
