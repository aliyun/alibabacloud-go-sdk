// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlrRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasPermission(v bool) *CreateSlrRoleResponseBody
	GetHasPermission() *bool
	SetRequestId(v string) *CreateSlrRoleResponseBody
	GetRequestId() *string
}

type CreateSlrRoleResponseBody struct {
	// Indicates whether the service-linked role was created. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	HasPermission *bool `json:"HasPermission,omitempty" xml:"HasPermission,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 208B016D-4CB9-4A85-96A5-0B8ED1EBF271
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSlrRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSlrRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSlrRoleResponseBody) GetHasPermission() *bool {
	return s.HasPermission
}

func (s *CreateSlrRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSlrRoleResponseBody) SetHasPermission(v bool) *CreateSlrRoleResponseBody {
	s.HasPermission = &v
	return s
}

func (s *CreateSlrRoleResponseBody) SetRequestId(v string) *CreateSlrRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSlrRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
