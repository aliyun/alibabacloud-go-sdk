// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataServiceAppGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int32) *DeleteDataServiceAppGroupRequest
	GetGroupId() *int32
	SetOpTenantId(v int64) *DeleteDataServiceAppGroupRequest
	GetOpTenantId() *int64
}

type DeleteDataServiceAppGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 550980364236
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteDataServiceAppGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataServiceAppGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataServiceAppGroupRequest) GetGroupId() *int32 {
	return s.GroupId
}

func (s *DeleteDataServiceAppGroupRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteDataServiceAppGroupRequest) SetGroupId(v int32) *DeleteDataServiceAppGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteDataServiceAppGroupRequest) SetOpTenantId(v int64) *DeleteDataServiceAppGroupRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteDataServiceAppGroupRequest) Validate() error {
	return dara.Validate(s)
}
