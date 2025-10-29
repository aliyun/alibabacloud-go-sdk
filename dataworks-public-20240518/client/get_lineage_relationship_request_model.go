// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLineageRelationshipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetLineageRelationshipRequest
	GetId() *string
}

type GetLineageRelationshipRequest struct {
	// The lineage ID. You can refer to the return result of the ListLineageRelationships operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 110xxxx:custom-table.xxxxx:maxcompute-table.project.test_big_lineage_080901:custom-sqlxx.00001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetLineageRelationshipRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLineageRelationshipRequest) GoString() string {
	return s.String()
}

func (s *GetLineageRelationshipRequest) GetId() *string {
	return s.Id
}

func (s *GetLineageRelationshipRequest) SetId(v string) *GetLineageRelationshipRequest {
	s.Id = &v
	return s
}

func (s *GetLineageRelationshipRequest) Validate() error {
	return dara.Validate(s)
}
