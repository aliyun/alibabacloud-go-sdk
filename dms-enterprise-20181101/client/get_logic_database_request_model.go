// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogicDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v string) *GetLogicDatabaseRequest
	GetDbId() *string
	SetTid(v int64) *GetLogicDatabaseRequest
	GetTid() *int64
}

type GetLogicDatabaseRequest struct {
	// The ID of the logical database. You can call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to obtain the ID of the logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1***
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetLogicDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogicDatabaseRequest) GoString() string {
	return s.String()
}

func (s *GetLogicDatabaseRequest) GetDbId() *string {
	return s.DbId
}

func (s *GetLogicDatabaseRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetLogicDatabaseRequest) SetDbId(v string) *GetLogicDatabaseRequest {
	s.DbId = &v
	return s
}

func (s *GetLogicDatabaseRequest) SetTid(v int64) *GetLogicDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *GetLogicDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
