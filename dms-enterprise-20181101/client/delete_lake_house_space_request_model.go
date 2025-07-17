// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLakeHouseSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSpaceId(v int64) *DeleteLakeHouseSpaceRequest
	GetSpaceId() *int64
	SetTid(v int64) *DeleteLakeHouseSpaceRequest
	GetTid() *int64
}

type DeleteLakeHouseSpaceRequest struct {
	// The ID of the workspace. You can call the [GetLhSpaceByName](https://help.aliyun.com/document_detail/424379.html) operation to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 24
	SpaceId *int64 `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteLakeHouseSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLakeHouseSpaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteLakeHouseSpaceRequest) GetSpaceId() *int64 {
	return s.SpaceId
}

func (s *DeleteLakeHouseSpaceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteLakeHouseSpaceRequest) SetSpaceId(v int64) *DeleteLakeHouseSpaceRequest {
	s.SpaceId = &v
	return s
}

func (s *DeleteLakeHouseSpaceRequest) SetTid(v int64) *DeleteLakeHouseSpaceRequest {
	s.Tid = &v
	return s
}

func (s *DeleteLakeHouseSpaceRequest) Validate() error {
	return dara.Validate(s)
}
