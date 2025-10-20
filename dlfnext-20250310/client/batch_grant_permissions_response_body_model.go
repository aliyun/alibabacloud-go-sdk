// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGrantPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *BatchGrantPermissionsResponseBody
	GetErrorMessage() *string
	SetFailurePermissions(v []*FailurePermission) *BatchGrantPermissionsResponseBody
	GetFailurePermissions() []*FailurePermission
	SetSuccess(v bool) *BatchGrantPermissionsResponseBody
	GetSuccess() *bool
}

type BatchGrantPermissionsResponseBody struct {
	ErrorMessage       *string              `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FailurePermissions []*FailurePermission `json:"failurePermissions,omitempty" xml:"failurePermissions,omitempty" type:"Repeated"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchGrantPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGrantPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGrantPermissionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchGrantPermissionsResponseBody) GetFailurePermissions() []*FailurePermission {
	return s.FailurePermissions
}

func (s *BatchGrantPermissionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchGrantPermissionsResponseBody) SetErrorMessage(v string) *BatchGrantPermissionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchGrantPermissionsResponseBody) SetFailurePermissions(v []*FailurePermission) *BatchGrantPermissionsResponseBody {
	s.FailurePermissions = v
	return s
}

func (s *BatchGrantPermissionsResponseBody) SetSuccess(v bool) *BatchGrantPermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchGrantPermissionsResponseBody) Validate() error {
	if s.FailurePermissions != nil {
		for _, item := range s.FailurePermissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
