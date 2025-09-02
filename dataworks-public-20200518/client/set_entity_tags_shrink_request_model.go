// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEntityTagsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifiedName(v string) *SetEntityTagsShrinkRequest
	GetQualifiedName() *string
	SetTagsShrink(v string) *SetEntityTagsShrinkRequest
	GetTagsShrink() *string
}

type SetEntityTagsShrinkRequest struct {
	// The unique identifier of the entity. Example: maxcompute-table.projectA.tableA.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table.projectA.tableA
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	// The tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s SetEntityTagsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetEntityTagsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetEntityTagsShrinkRequest) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *SetEntityTagsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *SetEntityTagsShrinkRequest) SetQualifiedName(v string) *SetEntityTagsShrinkRequest {
	s.QualifiedName = &v
	return s
}

func (s *SetEntityTagsShrinkRequest) SetTagsShrink(v string) *SetEntityTagsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *SetEntityTagsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
