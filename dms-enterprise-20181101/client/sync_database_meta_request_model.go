// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDatabaseMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v string) *SyncDatabaseMetaRequest
	GetDbId() *string
	SetLogic(v bool) *SyncDatabaseMetaRequest
	GetLogic() *bool
	SetTid(v int64) *SyncDatabaseMetaRequest
	GetTid() *int64
}

type SyncDatabaseMetaRequest struct {
	// The ID of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Specifies whether the database is a logical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SyncDatabaseMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncDatabaseMetaRequest) GoString() string {
	return s.String()
}

func (s *SyncDatabaseMetaRequest) GetDbId() *string {
	return s.DbId
}

func (s *SyncDatabaseMetaRequest) GetLogic() *bool {
	return s.Logic
}

func (s *SyncDatabaseMetaRequest) GetTid() *int64 {
	return s.Tid
}

func (s *SyncDatabaseMetaRequest) SetDbId(v string) *SyncDatabaseMetaRequest {
	s.DbId = &v
	return s
}

func (s *SyncDatabaseMetaRequest) SetLogic(v bool) *SyncDatabaseMetaRequest {
	s.Logic = &v
	return s
}

func (s *SyncDatabaseMetaRequest) SetTid(v int64) *SyncDatabaseMetaRequest {
	s.Tid = &v
	return s
}

func (s *SyncDatabaseMetaRequest) Validate() error {
	return dara.Validate(s)
}
