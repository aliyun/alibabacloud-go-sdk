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
	// The lineage ID. You can refer to the ListLineageRelationships operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4as3dasf654a
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
