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
	// example:
	//
	// ALIYUN::FC::FUNCTION
	ResouceType *string `json:"resouceType,omitempty" xml:"resouceType,omitempty"`
	// example:
	//
	// acs:fc:cn-shanghai:****:functions/demo
	ResourceArn *string            `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	Tags        map[string]*string `json:"tags" xml:"tags"`
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
