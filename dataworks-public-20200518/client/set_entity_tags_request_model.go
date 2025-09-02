// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEntityTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifiedName(v string) *SetEntityTagsRequest
	GetQualifiedName() *string
	SetTags(v []*UserEntityTag) *SetEntityTagsRequest
	GetTags() []*UserEntityTag
}

type SetEntityTagsRequest struct {
	// The unique identifier of the entity. Example: maxcompute-table.projectA.tableA.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table.projectA.tableA
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	// The tags.
	Tags []*UserEntityTag `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s SetEntityTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s SetEntityTagsRequest) GoString() string {
	return s.String()
}

func (s *SetEntityTagsRequest) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *SetEntityTagsRequest) GetTags() []*UserEntityTag {
	return s.Tags
}

func (s *SetEntityTagsRequest) SetQualifiedName(v string) *SetEntityTagsRequest {
	s.QualifiedName = &v
	return s
}

func (s *SetEntityTagsRequest) SetTags(v []*UserEntityTag) *SetEntityTagsRequest {
	s.Tags = v
	return s
}

func (s *SetEntityTagsRequest) Validate() error {
	return dara.Validate(s)
}
