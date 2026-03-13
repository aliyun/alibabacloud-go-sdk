// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPermission interface {
	dara.Model
	String() string
	GoString() string
	SetIsEnabled(v bool) *Permission
	GetIsEnabled() *bool
	SetResourceType(v string) *Permission
	GetResourceType() *string
}

type Permission struct {
	IsEnabled    *bool   `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s Permission) String() string {
	return dara.Prettify(s)
}

func (s Permission) GoString() string {
	return s.String()
}

func (s *Permission) GetIsEnabled() *bool {
	return s.IsEnabled
}

func (s *Permission) GetResourceType() *string {
	return s.ResourceType
}

func (s *Permission) SetIsEnabled(v bool) *Permission {
	s.IsEnabled = &v
	return s
}

func (s *Permission) SetResourceType(v string) *Permission {
	s.ResourceType = &v
	return s
}

func (s *Permission) Validate() error {
	return dara.Validate(s)
}
