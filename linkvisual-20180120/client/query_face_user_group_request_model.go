// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolationId(v string) *QueryFaceUserGroupRequest
	GetIsolationId() *string
	SetPageNo(v int32) *QueryFaceUserGroupRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryFaceUserGroupRequest
	GetPageSize() *int32
	SetUserId(v string) *QueryFaceUserGroupRequest
	GetUserId() *string
}

type QueryFaceUserGroupRequest struct {
	// The isolation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestIsolationId
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestUserId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *QueryFaceUserGroupRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryFaceUserGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceUserGroupRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryFaceUserGroupRequest) SetIsolationId(v string) *QueryFaceUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceUserGroupRequest) SetPageNo(v int32) *QueryFaceUserGroupRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceUserGroupRequest) SetPageSize(v int32) *QueryFaceUserGroupRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceUserGroupRequest) SetUserId(v string) *QueryFaceUserGroupRequest {
	s.UserId = &v
	return s
}

func (s *QueryFaceUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
