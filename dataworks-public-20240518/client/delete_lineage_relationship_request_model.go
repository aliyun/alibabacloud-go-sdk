// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLineageRelationshipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteLineageRelationshipRequest
	GetId() *string
}

type DeleteLineageRelationshipRequest struct {
	// The lineage ID. For more information, see the response returned by the ListLineageRelationships operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 110xxxx:custom-table.xxxxx:maxcompute-table.project.test_big_lineage_080901:custom-sqlxx.00001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteLineageRelationshipRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLineageRelationshipRequest) GoString() string {
	return s.String()
}

func (s *DeleteLineageRelationshipRequest) GetId() *string {
	return s.Id
}

func (s *DeleteLineageRelationshipRequest) SetId(v string) *DeleteLineageRelationshipRequest {
	s.Id = &v
	return s
}

func (s *DeleteLineageRelationshipRequest) Validate() error {
	return dara.Validate(s)
}
