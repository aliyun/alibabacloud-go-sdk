// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLineageRelationshipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateLineageRelationshipResponseBody
	GetId() *string
	SetRequestId(v string) *CreateLineageRelationshipResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateLineageRelationshipResponseBody
	GetSuccess() *bool
}

type CreateLineageRelationshipResponseBody struct {
	// The lineage ID.
	//
	// example:
	//
	// 110xxxx:custom-table.xxxxx:maxcompute-table.project.test_big_lineage_080901:custom-sqlxx.00001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// C99E2BE6-9DEA-5C2E-8F51-1DDCFEADE490
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLineageRelationshipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLineageRelationshipResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLineageRelationshipResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateLineageRelationshipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLineageRelationshipResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateLineageRelationshipResponseBody) SetId(v string) *CreateLineageRelationshipResponseBody {
	s.Id = &v
	return s
}

func (s *CreateLineageRelationshipResponseBody) SetRequestId(v string) *CreateLineageRelationshipResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLineageRelationshipResponseBody) SetSuccess(v bool) *CreateLineageRelationshipResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLineageRelationshipResponseBody) Validate() error {
	return dara.Validate(s)
}
