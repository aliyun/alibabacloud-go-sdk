// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEntityTagsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifiedName(v string) *RemoveEntityTagsShrinkRequest
	GetQualifiedName() *string
	SetTagKeysShrink(v string) *RemoveEntityTagsShrinkRequest
	GetTagKeysShrink() *string
}

type RemoveEntityTagsShrinkRequest struct {
	// The unique identifier of the entity. Example: maxcompute-table.projectA.tableA.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table.projectA.tableA
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	// The tag keys.
	//
	// This parameter is required.
	TagKeysShrink *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s RemoveEntityTagsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveEntityTagsShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveEntityTagsShrinkRequest) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *RemoveEntityTagsShrinkRequest) GetTagKeysShrink() *string {
	return s.TagKeysShrink
}

func (s *RemoveEntityTagsShrinkRequest) SetQualifiedName(v string) *RemoveEntityTagsShrinkRequest {
	s.QualifiedName = &v
	return s
}

func (s *RemoveEntityTagsShrinkRequest) SetTagKeysShrink(v string) *RemoveEntityTagsShrinkRequest {
	s.TagKeysShrink = &v
	return s
}

func (s *RemoveEntityTagsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
