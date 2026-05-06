// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *GetErServiceRequest
	GetSecurityToken() *string
}

type GetErServiceRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetErServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetErServiceRequest) GoString() string {
	return s.String()
}

func (s *GetErServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetErServiceRequest) SetSecurityToken(v string) *GetErServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetErServiceRequest) Validate() error {
	return dara.Validate(s)
}
