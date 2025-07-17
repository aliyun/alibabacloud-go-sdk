// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int64) *GetPhysicalDatabaseRequest
	GetDbId() *int64
	SetTid(v int64) *GetPhysicalDatabaseRequest
	GetTid() *int64
}

type GetPhysicalDatabaseRequest struct {
	// The ID of the physical database. You can call the [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 43153
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the ID of the tenant.
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetPhysicalDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalDatabaseRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalDatabaseRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *GetPhysicalDatabaseRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetPhysicalDatabaseRequest) SetDbId(v int64) *GetPhysicalDatabaseRequest {
	s.DbId = &v
	return s
}

func (s *GetPhysicalDatabaseRequest) SetTid(v int64) *GetPhysicalDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *GetPhysicalDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
