// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResource interface {
	dara.Model
	String() string
	GoString() string
	SetResouceType(v string) *Resource
	GetResouceType() *string
	SetResourceArn(v string) *Resource
	GetResourceArn() *string
	SetTags(v map[string]*string) *Resource
	GetTags() map[string]*string
}

type Resource struct {
	// The name of the resource type. Valid values: ALIYUN::FC::FUNCTION and ALIYUN::FC::SERVICE. The former name is used in Function Compute 3.0, and the latter name is used in earlier versions of Function Compute.
	//
	// example:
	//
	// ALIYUN::FC::FUNCTION
	ResouceType *string `json:"resouceType,omitempty" xml:"resouceType,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the resource.
	//
	// example:
	//
	// acs:fc:cn-shanghai:****:functions/demo
	ResourceArn *string `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	// The tag dictionary.
	Tags map[string]*string `json:"tags" xml:"tags"`
}

func (s Resource) String() string {
	return dara.Prettify(s)
}

func (s Resource) GoString() string {
	return s.String()
}

func (s *Resource) GetResouceType() *string {
	return s.ResouceType
}

func (s *Resource) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *Resource) GetTags() map[string]*string {
	return s.Tags
}

func (s *Resource) SetResouceType(v string) *Resource {
	s.ResouceType = &v
	return s
}

func (s *Resource) SetResourceArn(v string) *Resource {
	s.ResourceArn = &v
	return s
}

func (s *Resource) SetTags(v map[string]*string) *Resource {
	s.Tags = v
	return s
}

func (s *Resource) Validate() error {
	return dara.Validate(s)
}
