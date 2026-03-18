// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeletePolicyRequest
	GetId() *string
}

type DeletePolicyRequest struct {
	// The ID of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 90300b1a-ced8-4437-b4bf-f9a5*******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyRequest) GetId() *string {
	return s.Id
}

func (s *DeletePolicyRequest) SetId(v string) *DeletePolicyRequest {
	s.Id = &v
	return s
}

func (s *DeletePolicyRequest) Validate() error {
	return dara.Validate(s)
}
