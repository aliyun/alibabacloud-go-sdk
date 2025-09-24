// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRelationListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *QueryRelationListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryRelationListRequest
	GetPageSize() *int32
	SetStatusList(v []*string) *QueryRelationListRequest
	GetStatusList() []*string
	SetUserId(v int64) *QueryRelationListRequest
	GetUserId() *int64
}

type QueryRelationListRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The states of the relationships between the management account and its members. The valid values of this parameter are the enumeration members of the RelationshipStatusEnum data type. If you do not specify this parameter, valid relationship states are queried by default.
	//
	// example:
	//
	// RELATED
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1990699401005016
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryRelationListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRelationListRequest) GoString() string {
	return s.String()
}

func (s *QueryRelationListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryRelationListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRelationListRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *QueryRelationListRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryRelationListRequest) SetPageNum(v int32) *QueryRelationListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryRelationListRequest) SetPageSize(v int32) *QueryRelationListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRelationListRequest) SetStatusList(v []*string) *QueryRelationListRequest {
	s.StatusList = v
	return s
}

func (s *QueryRelationListRequest) SetUserId(v int64) *QueryRelationListRequest {
	s.UserId = &v
	return s
}

func (s *QueryRelationListRequest) Validate() error {
	return dara.Validate(s)
}
