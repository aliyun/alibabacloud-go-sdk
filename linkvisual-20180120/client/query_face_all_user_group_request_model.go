// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceAllUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolationId(v string) *QueryFaceAllUserGroupRequest
	GetIsolationId() *string
	SetPageNo(v int32) *QueryFaceAllUserGroupRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryFaceAllUserGroupRequest
	GetPageSize() *int32
}

type QueryFaceAllUserGroupRequest struct {
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
}

func (s QueryFaceAllUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *QueryFaceAllUserGroupRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryFaceAllUserGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceAllUserGroupRequest) SetIsolationId(v string) *QueryFaceAllUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceAllUserGroupRequest) SetPageNo(v int32) *QueryFaceAllUserGroupRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceAllUserGroupRequest) SetPageSize(v int32) *QueryFaceAllUserGroupRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
