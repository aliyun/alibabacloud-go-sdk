// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenEdgeContainerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *OpenEdgeContainerRequest
	GetSecurityToken() *string
}

type OpenEdgeContainerRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s OpenEdgeContainerRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenEdgeContainerRequest) GoString() string {
	return s.String()
}

func (s *OpenEdgeContainerRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *OpenEdgeContainerRequest) SetSecurityToken(v string) *OpenEdgeContainerRequest {
	s.SecurityToken = &v
	return s
}

func (s *OpenEdgeContainerRequest) Validate() error {
	return dara.Validate(s)
}
