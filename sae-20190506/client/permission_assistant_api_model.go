// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPermissionAssistantApi interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *PermissionAssistantApi
	GetCreateTime() *string
	SetId(v int64) *PermissionAssistantApi
	GetId() *int64
	SetName(v string) *PermissionAssistantApi
	GetName() *string
	SetResourceType(v string) *PermissionAssistantApi
	GetResourceType() *string
	SetUpdateTime(v string) *PermissionAssistantApi
	GetUpdateTime() *string
}

type PermissionAssistantApi struct {
	CreateTime   *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Id           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	UpdateTime   *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s PermissionAssistantApi) String() string {
	return dara.Prettify(s)
}

func (s PermissionAssistantApi) GoString() string {
	return s.String()
}

func (s *PermissionAssistantApi) GetCreateTime() *string {
	return s.CreateTime
}

func (s *PermissionAssistantApi) GetId() *int64 {
	return s.Id
}

func (s *PermissionAssistantApi) GetName() *string {
	return s.Name
}

func (s *PermissionAssistantApi) GetResourceType() *string {
	return s.ResourceType
}

func (s *PermissionAssistantApi) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *PermissionAssistantApi) SetCreateTime(v string) *PermissionAssistantApi {
	s.CreateTime = &v
	return s
}

func (s *PermissionAssistantApi) SetId(v int64) *PermissionAssistantApi {
	s.Id = &v
	return s
}

func (s *PermissionAssistantApi) SetName(v string) *PermissionAssistantApi {
	s.Name = &v
	return s
}

func (s *PermissionAssistantApi) SetResourceType(v string) *PermissionAssistantApi {
	s.ResourceType = &v
	return s
}

func (s *PermissionAssistantApi) SetUpdateTime(v string) *PermissionAssistantApi {
	s.UpdateTime = &v
	return s
}

func (s *PermissionAssistantApi) Validate() error {
	return dara.Validate(s)
}
