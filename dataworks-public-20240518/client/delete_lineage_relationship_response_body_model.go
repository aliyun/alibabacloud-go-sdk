// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLineageRelationshipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLineageRelationshipResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteLineageRelationshipResponseBody
	GetSuccess() *bool
}

type DeleteLineageRelationshipResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 952795279527ab****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: success.
	//
	// 	- false: failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteLineageRelationshipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLineageRelationshipResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLineageRelationshipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLineageRelationshipResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteLineageRelationshipResponseBody) SetRequestId(v string) *DeleteLineageRelationshipResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLineageRelationshipResponseBody) SetSuccess(v bool) *DeleteLineageRelationshipResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteLineageRelationshipResponseBody) Validate() error {
	return dara.Validate(s)
}
