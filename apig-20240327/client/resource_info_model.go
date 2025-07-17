// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceInfo interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *ResourceInfo
	GetResourceId() *string
	SetResourceName(v string) *ResourceInfo
	GetResourceName() *string
	SetResourceType(v string) *ResourceInfo
	GetResourceType() *string
	SetResourceVersion(v string) *ResourceInfo
	GetResourceVersion() *string
}

type ResourceInfo struct {
	ResourceId      *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceName    *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	ResourceType    *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
}

func (s ResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ResourceInfo) GoString() string {
	return s.String()
}

func (s *ResourceInfo) GetResourceId() *string {
	return s.ResourceId
}

func (s *ResourceInfo) GetResourceName() *string {
	return s.ResourceName
}

func (s *ResourceInfo) GetResourceType() *string {
	return s.ResourceType
}

func (s *ResourceInfo) GetResourceVersion() *string {
	return s.ResourceVersion
}

func (s *ResourceInfo) SetResourceId(v string) *ResourceInfo {
	s.ResourceId = &v
	return s
}

func (s *ResourceInfo) SetResourceName(v string) *ResourceInfo {
	s.ResourceName = &v
	return s
}

func (s *ResourceInfo) SetResourceType(v string) *ResourceInfo {
	s.ResourceType = &v
	return s
}

func (s *ResourceInfo) SetResourceVersion(v string) *ResourceInfo {
	s.ResourceVersion = &v
	return s
}

func (s *ResourceInfo) Validate() error {
	return dara.Validate(s)
}
