// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLineageRelationshipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLineageRelationship(v *LineageRelationship) *GetLineageRelationshipResponseBody
	GetLineageRelationship() *LineageRelationship
	SetRequestId(v string) *GetLineageRelationshipResponseBody
	GetRequestId() *string
}

type GetLineageRelationshipResponseBody struct {
	LineageRelationship *LineageRelationship `json:"LineageRelationship,omitempty" xml:"LineageRelationship,omitempty"`
	// example:
	//
	// 58D5334A-B013-430E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLineageRelationshipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLineageRelationshipResponseBody) GoString() string {
	return s.String()
}

func (s *GetLineageRelationshipResponseBody) GetLineageRelationship() *LineageRelationship {
	return s.LineageRelationship
}

func (s *GetLineageRelationshipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLineageRelationshipResponseBody) SetLineageRelationship(v *LineageRelationship) *GetLineageRelationshipResponseBody {
	s.LineageRelationship = v
	return s
}

func (s *GetLineageRelationshipResponseBody) SetRequestId(v string) *GetLineageRelationshipResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLineageRelationshipResponseBody) Validate() error {
	return dara.Validate(s)
}
