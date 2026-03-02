// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPermissionAllowResource interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PermissionAllowResource
	GetRequestId() *string
	SetResourceList(v []*string) *PermissionAllowResource
	GetResourceList() []*string
}

type PermissionAllowResource struct {
	RequestId    *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResourceList []*string `json:"resourceList,omitempty" xml:"resourceList,omitempty" type:"Repeated"`
}

func (s PermissionAllowResource) String() string {
	return dara.Prettify(s)
}

func (s PermissionAllowResource) GoString() string {
	return s.String()
}

func (s *PermissionAllowResource) GetRequestId() *string {
	return s.RequestId
}

func (s *PermissionAllowResource) GetResourceList() []*string {
	return s.ResourceList
}

func (s *PermissionAllowResource) SetRequestId(v string) *PermissionAllowResource {
	s.RequestId = &v
	return s
}

func (s *PermissionAllowResource) SetResourceList(v []*string) *PermissionAllowResource {
	s.ResourceList = v
	return s
}

func (s *PermissionAllowResource) Validate() error {
	return dara.Validate(s)
}
