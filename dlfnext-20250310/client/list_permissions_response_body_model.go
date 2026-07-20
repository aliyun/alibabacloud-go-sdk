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
	// The pagination token used to retrieve the next page of data. If null is returned, the current page is the last page of results.
	//
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	// The permission list.
	Permissions []*Permission `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
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
	if s.Permissions != nil {
		for _, item := range s.Permissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
