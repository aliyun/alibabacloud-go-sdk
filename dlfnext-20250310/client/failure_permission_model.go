// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailurePermission interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *FailurePermission
	GetErrorCode() *string
	SetErrorMessage(v string) *FailurePermission
	GetErrorMessage() *string
	SetPermission(v *Permission) *FailurePermission
	GetPermission() *Permission
}

type FailurePermission struct {
	ErrorCode    *string     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string     `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Permission   *Permission `json:"permission,omitempty" xml:"permission,omitempty"`
}

func (s FailurePermission) String() string {
	return dara.Prettify(s)
}

func (s FailurePermission) GoString() string {
	return s.String()
}

func (s *FailurePermission) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *FailurePermission) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *FailurePermission) GetPermission() *Permission {
	return s.Permission
}

func (s *FailurePermission) SetErrorCode(v string) *FailurePermission {
	s.ErrorCode = &v
	return s
}

func (s *FailurePermission) SetErrorMessage(v string) *FailurePermission {
	s.ErrorMessage = &v
	return s
}

func (s *FailurePermission) SetPermission(v *Permission) *FailurePermission {
	s.Permission = v
	return s
}

func (s *FailurePermission) Validate() error {
	if s.Permission != nil {
		if err := s.Permission.Validate(); err != nil {
			return err
		}
	}
	return nil
}
