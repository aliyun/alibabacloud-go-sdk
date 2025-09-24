// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPermissionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRelationId(v int64) *QueryPermissionListRequest
	GetRelationId() *int64
}

type QueryPermissionListRequest struct {
	// The ID of the relationship. Set this parameter to the value of the RelationId response parameter returned by calling the QueryRelationList operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51463
	RelationId *int64 `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
}

func (s QueryPermissionListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPermissionListRequest) GoString() string {
	return s.String()
}

func (s *QueryPermissionListRequest) GetRelationId() *int64 {
	return s.RelationId
}

func (s *QueryPermissionListRequest) SetRelationId(v int64) *QueryPermissionListRequest {
	s.RelationId = &v
	return s
}

func (s *QueryPermissionListRequest) Validate() error {
	return dara.Validate(s)
}
