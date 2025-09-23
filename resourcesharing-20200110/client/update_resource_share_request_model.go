// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceShareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowExternalTargets(v bool) *UpdateResourceShareRequest
	GetAllowExternalTargets() *bool
	SetResourceShareId(v string) *UpdateResourceShareRequest
	GetResourceShareId() *string
	SetResourceShareName(v string) *UpdateResourceShareRequest
	GetResourceShareName() *string
}

type UpdateResourceShareRequest struct {
	// Specifies whether resources in the resource share can be shared with accounts outside the resource directory. Valid values:
	//
	// 	- false: Resources in the resource share can be shared only with accounts in the resource directory.
	//
	// 	- true: Resources in the resource share can be shared with both accounts in the resource directory and accounts outside the resource directory.
	//
	// example:
	//
	// false
	AllowExternalTargets *bool `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	// The ID of the resource share.
	//
	// This parameter is required.
	//
	// example:
	//
	// rs-qSkW1HBY****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The new name of the resource share.
	//
	// The name must be 1 to 50 characters in length.
	//
	// The name can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// new
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
}

func (s UpdateResourceShareRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceShareRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceShareRequest) GetAllowExternalTargets() *bool {
	return s.AllowExternalTargets
}

func (s *UpdateResourceShareRequest) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *UpdateResourceShareRequest) GetResourceShareName() *string {
	return s.ResourceShareName
}

func (s *UpdateResourceShareRequest) SetAllowExternalTargets(v bool) *UpdateResourceShareRequest {
	s.AllowExternalTargets = &v
	return s
}

func (s *UpdateResourceShareRequest) SetResourceShareId(v string) *UpdateResourceShareRequest {
	s.ResourceShareId = &v
	return s
}

func (s *UpdateResourceShareRequest) SetResourceShareName(v string) *UpdateResourceShareRequest {
	s.ResourceShareName = &v
	return s
}

func (s *UpdateResourceShareRequest) Validate() error {
	return dara.Validate(s)
}
