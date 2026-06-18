// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *GetEdgeContainerRequest
	GetSecurityToken() *string
}

type GetEdgeContainerRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetEdgeContainerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetEdgeContainerRequest) SetSecurityToken(v string) *GetEdgeContainerRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetEdgeContainerRequest) Validate() error {
	return dara.Validate(s)
}
