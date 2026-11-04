// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceAllUserIdsByGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolationId(v string) *QueryFaceAllUserIdsByGroupIdRequest
	GetIsolationId() *string
	SetPageNo(v int32) *QueryFaceAllUserIdsByGroupIdRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryFaceAllUserIdsByGroupIdRequest
	GetPageSize() *int32
	SetUserGroupId(v string) *QueryFaceAllUserIdsByGroupIdRequest
	GetUserGroupId() *string
}

type QueryFaceAllUserIdsByGroupIdRequest struct {
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
	// The page size.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The user group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestUserGroupId
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s QueryFaceAllUserIdsByGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserIdsByGroupIdRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetIsolationId(v string) *QueryFaceAllUserIdsByGroupIdRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetPageNo(v int32) *QueryFaceAllUserIdsByGroupIdRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetPageSize(v int32) *QueryFaceAllUserIdsByGroupIdRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetUserGroupId(v string) *QueryFaceAllUserIdsByGroupIdRequest {
	s.UserGroupId = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) Validate() error {
	return dara.Validate(s)
}
