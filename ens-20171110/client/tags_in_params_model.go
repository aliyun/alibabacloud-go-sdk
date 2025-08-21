// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagsInParams interface {
	dara.Model
	String() string
	GoString() string
	SetTag(v []*TagsInParamsTag) *TagsInParams
	GetTag() []*TagsInParamsTag
}

type TagsInParams struct {
	Tag []*TagsInParamsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagsInParams) String() string {
	return dara.Prettify(s)
}

func (s TagsInParams) GoString() string {
	return s.String()
}

func (s *TagsInParams) GetTag() []*TagsInParamsTag {
	return s.Tag
}

func (s *TagsInParams) SetTag(v []*TagsInParamsTag) *TagsInParams {
	s.Tag = v
	return s
}

func (s *TagsInParams) Validate() error {
	return dara.Validate(s)
}

type TagsInParamsTag struct {
	// This parameter is required.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test-key-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagsInParamsTag) String() string {
	return dara.Prettify(s)
}

func (s TagsInParamsTag) GoString() string {
	return s.String()
}

func (s *TagsInParamsTag) GetKey() *string {
	return s.Key
}

func (s *TagsInParamsTag) GetValue() *string {
	return s.Value
}

func (s *TagsInParamsTag) SetKey(v string) *TagsInParamsTag {
	s.Key = &v
	return s
}

func (s *TagsInParamsTag) SetValue(v string) *TagsInParamsTag {
	s.Value = &v
	return s
}

func (s *TagsInParamsTag) Validate() error {
	return dara.Validate(s)
}
