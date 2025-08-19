// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTagsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetResouceType(v string) *GetResourceTagsOutput
	GetResouceType() *string
	SetResourceArn(v string) *GetResourceTagsOutput
	GetResourceArn() *string
	SetTags(v map[string]*string) *GetResourceTagsOutput
	GetTags() map[string]*string
}

type GetResourceTagsOutput struct {
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

func (s GetResourceTagsOutput) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTagsOutput) GoString() string {
	return s.String()
}

func (s *GetResourceTagsOutput) GetResouceType() *string {
	return s.ResouceType
}

func (s *GetResourceTagsOutput) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *GetResourceTagsOutput) GetTags() map[string]*string {
	return s.Tags
}

func (s *GetResourceTagsOutput) SetResouceType(v string) *GetResourceTagsOutput {
	s.ResouceType = &v
	return s
}

func (s *GetResourceTagsOutput) SetResourceArn(v string) *GetResourceTagsOutput {
	s.ResourceArn = &v
	return s
}

func (s *GetResourceTagsOutput) SetTags(v map[string]*string) *GetResourceTagsOutput {
	s.Tags = v
	return s
}

func (s *GetResourceTagsOutput) Validate() error {
	return dara.Validate(s)
}
