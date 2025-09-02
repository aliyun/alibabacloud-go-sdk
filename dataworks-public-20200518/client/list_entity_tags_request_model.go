// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntityTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifiedName(v string) *ListEntityTagsRequest
	GetQualifiedName() *string
}

type ListEntityTagsRequest struct {
	// The unique identifier of the entity. Example: maxcompute-table.projectA.tableA.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table.projectA.tableA
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
}

func (s ListEntityTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEntityTagsRequest) GoString() string {
	return s.String()
}

func (s *ListEntityTagsRequest) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *ListEntityTagsRequest) SetQualifiedName(v string) *ListEntityTagsRequest {
	s.QualifiedName = &v
	return s
}

func (s *ListEntityTagsRequest) Validate() error {
	return dara.Validate(s)
}
