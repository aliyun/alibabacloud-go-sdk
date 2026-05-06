// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenErServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *OpenErServiceRequest
	GetSecurityToken() *string
}

type OpenErServiceRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s OpenErServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenErServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenErServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *OpenErServiceRequest) SetSecurityToken(v string) *OpenErServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *OpenErServiceRequest) Validate() error {
	return dara.Validate(s)
}
