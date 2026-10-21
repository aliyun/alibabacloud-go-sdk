// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccessKeyPolicyForAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyPolicy(v string) *SetAccessKeyPolicyForAccountRequest
	GetAccessKeyPolicy() *string
}

type SetAccessKeyPolicyForAccountRequest struct {
	// The network access restriction policy. The value is a JSON-formatted string. For more information, see the AccessKeyPolicy structure description in the SetAccessKeyPolicy documentation.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Status":"Inactive","Statements":[{"Value":"AllowAllVPC","Type":"VPCWhiteList","IPList":["::/0","0.0.0.0/0"]}]}
	AccessKeyPolicy *string `json:"AccessKeyPolicy,omitempty" xml:"AccessKeyPolicy,omitempty"`
}

func (s SetAccessKeyPolicyForAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAccessKeyPolicyForAccountRequest) GoString() string {
	return s.String()
}

func (s *SetAccessKeyPolicyForAccountRequest) GetAccessKeyPolicy() *string {
	return s.AccessKeyPolicy
}

func (s *SetAccessKeyPolicyForAccountRequest) SetAccessKeyPolicy(v string) *SetAccessKeyPolicyForAccountRequest {
	s.AccessKeyPolicy = &v
	return s
}

func (s *SetAccessKeyPolicyForAccountRequest) Validate() error {
	return dara.Validate(s)
}
