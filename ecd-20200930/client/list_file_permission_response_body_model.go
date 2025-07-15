// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilePermissions(v []*FilePermissionMember) *ListFilePermissionResponseBody
	GetFilePermissions() []*FilePermissionMember
	SetRequestId(v string) *ListFilePermissionResponseBody
	GetRequestId() *string
}

type ListFilePermissionResponseBody struct {
	// The permissions on the shared file.
	FilePermissions []*FilePermissionMember `json:"FilePermissions,omitempty" xml:"FilePermissions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFilePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFilePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ListFilePermissionResponseBody) GetFilePermissions() []*FilePermissionMember {
	return s.FilePermissions
}

func (s *ListFilePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFilePermissionResponseBody) SetFilePermissions(v []*FilePermissionMember) *ListFilePermissionResponseBody {
	s.FilePermissions = v
	return s
}

func (s *ListFilePermissionResponseBody) SetRequestId(v string) *ListFilePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFilePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
