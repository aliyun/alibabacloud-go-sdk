// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListPermissionsResponseBody
	GetNextPageToken() *string
	SetPermissions(v []*Permission) *ListPermissionsResponseBody
	GetPermissions() []*Permission
}

type ListPermissionsResponseBody struct {
	// example:
	//
	// token!
	NextPageToken *string       `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Permissions   []*Permission `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s ListPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListPermissionsResponseBody) GetPermissions() []*Permission {
	return s.Permissions
}

func (s *ListPermissionsResponseBody) SetNextPageToken(v string) *ListPermissionsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListPermissionsResponseBody) SetPermissions(v []*Permission) *ListPermissionsResponseBody {
	s.Permissions = v
	return s
}

func (s *ListPermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}
