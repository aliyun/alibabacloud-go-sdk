// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAuthorizedResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListUserAuthorizedResourcesRequest
	GetNextToken() *string
	SetPermissionCode(v string) *ListUserAuthorizedResourcesRequest
	GetPermissionCode() *string
	SetResourceType(v string) *ListUserAuthorizedResourcesRequest
	GetResourceType() *string
}

type ListUserAuthorizedResourcesRequest struct {
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PermissionCode *string `json:"PermissionCode,omitempty" xml:"PermissionCode,omitempty"`
	ResourceType   *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListUserAuthorizedResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthorizedResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListUserAuthorizedResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserAuthorizedResourcesRequest) GetPermissionCode() *string {
	return s.PermissionCode
}

func (s *ListUserAuthorizedResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListUserAuthorizedResourcesRequest) SetNextToken(v string) *ListUserAuthorizedResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserAuthorizedResourcesRequest) SetPermissionCode(v string) *ListUserAuthorizedResourcesRequest {
	s.PermissionCode = &v
	return s
}

func (s *ListUserAuthorizedResourcesRequest) SetResourceType(v string) *ListUserAuthorizedResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListUserAuthorizedResourcesRequest) Validate() error {
	return dara.Validate(s)
}
