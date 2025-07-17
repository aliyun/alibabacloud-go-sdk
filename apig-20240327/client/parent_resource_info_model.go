// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParentResourceInfo interface {
	dara.Model
	String() string
	GoString() string
	SetApiInfo(v *HttpApiApiInfo) *ParentResourceInfo
	GetApiInfo() *HttpApiApiInfo
	SetResourceType(v string) *ParentResourceInfo
	GetResourceType() *string
}

type ParentResourceInfo struct {
	ApiInfo      *HttpApiApiInfo `json:"apiInfo,omitempty" xml:"apiInfo,omitempty"`
	ResourceType *string         `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ParentResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ParentResourceInfo) GoString() string {
	return s.String()
}

func (s *ParentResourceInfo) GetApiInfo() *HttpApiApiInfo {
	return s.ApiInfo
}

func (s *ParentResourceInfo) GetResourceType() *string {
	return s.ResourceType
}

func (s *ParentResourceInfo) SetApiInfo(v *HttpApiApiInfo) *ParentResourceInfo {
	s.ApiInfo = v
	return s
}

func (s *ParentResourceInfo) SetResourceType(v string) *ParentResourceInfo {
	s.ResourceType = &v
	return s
}

func (s *ParentResourceInfo) Validate() error {
	return dara.Validate(s)
}
