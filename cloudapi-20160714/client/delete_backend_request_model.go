// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendId(v string) *DeleteBackendRequest
	GetBackendId() *string
	SetSecurityToken(v string) *DeleteBackendRequest
	GetSecurityToken() *string
}

type DeleteBackendRequest struct {
	// The ID of the backend service.
	//
	// example:
	//
	// 27be0dd9ebbc467b9e86c0d250d0b92e
	BackendId     *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteBackendRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackendRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackendRequest) GetBackendId() *string {
	return s.BackendId
}

func (s *DeleteBackendRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteBackendRequest) SetBackendId(v string) *DeleteBackendRequest {
	s.BackendId = &v
	return s
}

func (s *DeleteBackendRequest) SetSecurityToken(v string) *DeleteBackendRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteBackendRequest) Validate() error {
	return dara.Validate(s)
}
