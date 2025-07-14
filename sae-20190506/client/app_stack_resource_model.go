// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppStackResource interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *AppStackResource
	GetCreateTime() *int64
	SetInstanceId(v string) *AppStackResource
	GetInstanceId() *string
	SetProductCode(v string) *AppStackResource
	GetProductCode() *string
	SetResourceId(v string) *AppStackResource
	GetResourceId() *string
	SetResourceName(v string) *AppStackResource
	GetResourceName() *string
	SetResourceType(v string) *AppStackResource
	GetResourceType() *string
	SetStackId(v string) *AppStackResource
	GetStackId() *string
	SetStatus(v string) *AppStackResource
	GetStatus() *string
}

type AppStackResource struct {
	// example:
	//
	// 1706518652000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// i-78yt
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// VPC
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// vsw-qwe112233
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// vsw-palworld-a
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// VSWITCH
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// palworld
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AppStackResource) String() string {
	return dara.Prettify(s)
}

func (s AppStackResource) GoString() string {
	return s.String()
}

func (s *AppStackResource) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *AppStackResource) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AppStackResource) GetProductCode() *string {
	return s.ProductCode
}

func (s *AppStackResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *AppStackResource) GetResourceName() *string {
	return s.ResourceName
}

func (s *AppStackResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *AppStackResource) GetStackId() *string {
	return s.StackId
}

func (s *AppStackResource) GetStatus() *string {
	return s.Status
}

func (s *AppStackResource) SetCreateTime(v int64) *AppStackResource {
	s.CreateTime = &v
	return s
}

func (s *AppStackResource) SetInstanceId(v string) *AppStackResource {
	s.InstanceId = &v
	return s
}

func (s *AppStackResource) SetProductCode(v string) *AppStackResource {
	s.ProductCode = &v
	return s
}

func (s *AppStackResource) SetResourceId(v string) *AppStackResource {
	s.ResourceId = &v
	return s
}

func (s *AppStackResource) SetResourceName(v string) *AppStackResource {
	s.ResourceName = &v
	return s
}

func (s *AppStackResource) SetResourceType(v string) *AppStackResource {
	s.ResourceType = &v
	return s
}

func (s *AppStackResource) SetStackId(v string) *AppStackResource {
	s.StackId = &v
	return s
}

func (s *AppStackResource) SetStatus(v string) *AppStackResource {
	s.Status = &v
	return s
}

func (s *AppStackResource) Validate() error {
	return dara.Validate(s)
}
