// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOption(v string) *CreateUploadPolicyRequest
	GetOption() *string
	SetType(v string) *CreateUploadPolicyRequest
	GetType() *string
}

type CreateUploadPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {}
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// scene
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateUploadPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadPolicyRequest) GetOption() *string {
	return s.Option
}

func (s *CreateUploadPolicyRequest) GetType() *string {
	return s.Type
}

func (s *CreateUploadPolicyRequest) SetOption(v string) *CreateUploadPolicyRequest {
	s.Option = &v
	return s
}

func (s *CreateUploadPolicyRequest) SetType(v string) *CreateUploadPolicyRequest {
	s.Type = &v
	return s
}

func (s *CreateUploadPolicyRequest) Validate() error {
	return dara.Validate(s)
}
