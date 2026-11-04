// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceAllUserGroupAndDeviceGroupRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolationId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest
	GetIsolationId() *string
	SetPageNo(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest
	GetPageSize() *int32
}

type QueryFaceAllUserGroupAndDeviceGroupRelationRequest struct {
	// The isolation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Isolati***
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

func (s QueryFaceAllUserGroupAndDeviceGroupRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) SetPageNo(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) SetPageSize(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) Validate() error {
	return dara.Validate(s)
}
