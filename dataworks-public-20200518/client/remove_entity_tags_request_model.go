// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEntityTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifiedName(v string) *RemoveEntityTagsRequest
	GetQualifiedName() *string
	SetTagKeys(v []*string) *RemoveEntityTagsRequest
	GetTagKeys() []*string
}

type RemoveEntityTagsRequest struct {
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
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s RemoveEntityTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveEntityTagsRequest) GoString() string {
	return s.String()
}

func (s *RemoveEntityTagsRequest) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *RemoveEntityTagsRequest) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *RemoveEntityTagsRequest) SetQualifiedName(v string) *RemoveEntityTagsRequest {
	s.QualifiedName = &v
	return s
}

func (s *RemoveEntityTagsRequest) SetTagKeys(v []*string) *RemoveEntityTagsRequest {
	s.TagKeys = v
	return s
}

func (s *RemoveEntityTagsRequest) Validate() error {
	return dara.Validate(s)
}
