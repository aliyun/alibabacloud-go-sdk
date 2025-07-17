// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizationResourceInfo interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *AuthorizationResourceInfo
	GetEnvironmentId() *string
	SetParentResourceId(v string) *AuthorizationResourceInfo
	GetParentResourceId() *string
	SetResourceId(v string) *AuthorizationResourceInfo
	GetResourceId() *string
}

type AuthorizationResourceInfo struct {
	EnvironmentId    *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	ParentResourceId *string `json:"parentResourceId,omitempty" xml:"parentResourceId,omitempty"`
	ResourceId       *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
}

func (s AuthorizationResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s AuthorizationResourceInfo) GoString() string {
	return s.String()
}

func (s *AuthorizationResourceInfo) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *AuthorizationResourceInfo) GetParentResourceId() *string {
	return s.ParentResourceId
}

func (s *AuthorizationResourceInfo) GetResourceId() *string {
	return s.ResourceId
}

func (s *AuthorizationResourceInfo) SetEnvironmentId(v string) *AuthorizationResourceInfo {
	s.EnvironmentId = &v
	return s
}

func (s *AuthorizationResourceInfo) SetParentResourceId(v string) *AuthorizationResourceInfo {
	s.ParentResourceId = &v
	return s
}

func (s *AuthorizationResourceInfo) SetResourceId(v string) *AuthorizationResourceInfo {
	s.ResourceId = &v
	return s
}

func (s *AuthorizationResourceInfo) Validate() error {
	return dara.Validate(s)
}
