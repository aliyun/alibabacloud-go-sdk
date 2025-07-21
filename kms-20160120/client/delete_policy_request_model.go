// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeletePolicyRequest
	GetName() *string
}

type DeletePolicyRequest struct {
	// The name of the permission policy that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyRequest) GetName() *string {
	return s.Name
}

func (s *DeletePolicyRequest) SetName(v string) *DeletePolicyRequest {
	s.Name = &v
	return s
}

func (s *DeletePolicyRequest) Validate() error {
	return dara.Validate(s)
}
