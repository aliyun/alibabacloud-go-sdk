// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBatchTableAccessPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int64) *CheckBatchTableAccessPermissionRequest
	GetDbId() *int64
	SetLogic(v bool) *CheckBatchTableAccessPermissionRequest
	GetLogic() *bool
	SetPermissionType(v string) *CheckBatchTableAccessPermissionRequest
	GetPermissionType() *string
	SetTableNameList(v []*string) *CheckBatchTableAccessPermissionRequest
	GetTableNameList() []*string
	SetTid(v int64) *CheckBatchTableAccessPermissionRequest
	GetTid() *int64
}

type CheckBatchTableAccessPermissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 43153
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// QUERY
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
	// This parameter is required.
	TableNameList []*string `json:"TableNameList,omitempty" xml:"TableNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CheckBatchTableAccessPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckBatchTableAccessPermissionRequest) GoString() string {
	return s.String()
}

func (s *CheckBatchTableAccessPermissionRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *CheckBatchTableAccessPermissionRequest) GetLogic() *bool {
	return s.Logic
}

func (s *CheckBatchTableAccessPermissionRequest) GetPermissionType() *string {
	return s.PermissionType
}

func (s *CheckBatchTableAccessPermissionRequest) GetTableNameList() []*string {
	return s.TableNameList
}

func (s *CheckBatchTableAccessPermissionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CheckBatchTableAccessPermissionRequest) SetDbId(v int64) *CheckBatchTableAccessPermissionRequest {
	s.DbId = &v
	return s
}

func (s *CheckBatchTableAccessPermissionRequest) SetLogic(v bool) *CheckBatchTableAccessPermissionRequest {
	s.Logic = &v
	return s
}

func (s *CheckBatchTableAccessPermissionRequest) SetPermissionType(v string) *CheckBatchTableAccessPermissionRequest {
	s.PermissionType = &v
	return s
}

func (s *CheckBatchTableAccessPermissionRequest) SetTableNameList(v []*string) *CheckBatchTableAccessPermissionRequest {
	s.TableNameList = v
	return s
}

func (s *CheckBatchTableAccessPermissionRequest) SetTid(v int64) *CheckBatchTableAccessPermissionRequest {
	s.Tid = &v
	return s
}

func (s *CheckBatchTableAccessPermissionRequest) Validate() error {
	return dara.Validate(s)
}
