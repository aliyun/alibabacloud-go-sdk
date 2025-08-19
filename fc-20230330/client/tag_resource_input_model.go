// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourceInput interface {
	dara.Model
	String() string
	GoString() string
	SetResourceArn(v string) *TagResourceInput
	GetResourceArn() *string
	SetTags(v map[string]*string) *TagResourceInput
	GetTags() map[string]*string
}

type TagResourceInput struct {
	// This parameter is required.
	//
	// example:
	//
	// acs:fc:cn-shanghai:xxx:functions/f1
	ResourceArn *string `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	// This parameter is required.
	Tags map[string]*string `json:"tags" xml:"tags"`
}

func (s TagResourceInput) String() string {
	return dara.Prettify(s)
}

func (s TagResourceInput) GoString() string {
	return s.String()
}

func (s *TagResourceInput) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *TagResourceInput) GetTags() map[string]*string {
	return s.Tags
}

func (s *TagResourceInput) SetResourceArn(v string) *TagResourceInput {
	s.ResourceArn = &v
	return s
}

func (s *TagResourceInput) SetTags(v map[string]*string) *TagResourceInput {
	s.Tags = v
	return s
}

func (s *TagResourceInput) Validate() error {
	return dara.Validate(s)
}
