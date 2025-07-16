// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRevokePermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *BatchRevokePermissionsResponseBody
	GetErrorMessage() *string
	SetFailurePermissions(v []*FailurePermission) *BatchRevokePermissionsResponseBody
	GetFailurePermissions() []*FailurePermission
	SetSuccess(v bool) *BatchRevokePermissionsResponseBody
	GetSuccess() *bool
}

type BatchRevokePermissionsResponseBody struct {
	ErrorMessage       *string              `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FailurePermissions []*FailurePermission `json:"failurePermissions,omitempty" xml:"failurePermissions,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchRevokePermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchRevokePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRevokePermissionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchRevokePermissionsResponseBody) GetFailurePermissions() []*FailurePermission {
	return s.FailurePermissions
}

func (s *BatchRevokePermissionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchRevokePermissionsResponseBody) SetErrorMessage(v string) *BatchRevokePermissionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchRevokePermissionsResponseBody) SetFailurePermissions(v []*FailurePermission) *BatchRevokePermissionsResponseBody {
	s.FailurePermissions = v
	return s
}

func (s *BatchRevokePermissionsResponseBody) SetSuccess(v bool) *BatchRevokePermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchRevokePermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}
