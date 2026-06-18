// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotlineInQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOuterGroupId(v string) *QueryHotlineInQueueRequest
	GetOuterGroupId() *string
	SetOuterGroupType(v string) *QueryHotlineInQueueRequest
	GetOuterGroupType() *string
}

type QueryHotlineInQueueRequest struct {
	// External skill group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OuterGroupId *string `json:"OuterGroupId,omitempty" xml:"OuterGroupId,omitempty"`
	// External skill group type.
	//
	// This parameter is required.
	//
	// example:
	//
	// mybank
	OuterGroupType *string `json:"OuterGroupType,omitempty" xml:"OuterGroupType,omitempty"`
}

func (s QueryHotlineInQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHotlineInQueueRequest) GoString() string {
	return s.String()
}

func (s *QueryHotlineInQueueRequest) GetOuterGroupId() *string {
	return s.OuterGroupId
}

func (s *QueryHotlineInQueueRequest) GetOuterGroupType() *string {
	return s.OuterGroupType
}

func (s *QueryHotlineInQueueRequest) SetOuterGroupId(v string) *QueryHotlineInQueueRequest {
	s.OuterGroupId = &v
	return s
}

func (s *QueryHotlineInQueueRequest) SetOuterGroupType(v string) *QueryHotlineInQueueRequest {
	s.OuterGroupType = &v
	return s
}

func (s *QueryHotlineInQueueRequest) Validate() error {
	return dara.Validate(s)
}
