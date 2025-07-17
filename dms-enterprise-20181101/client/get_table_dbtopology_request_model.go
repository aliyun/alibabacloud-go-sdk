// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableDBTopologyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTableGuid(v string) *GetTableDBTopologyRequest
	GetTableGuid() *string
	SetTid(v int64) *GetTableDBTopologyRequest
	GetTid() *int64
}

type GetTableDBTopologyRequest struct {
	// The GUID of the table in DMS.
	//
	// >
	//
	// 	- If the database to which the table belongs is a logical database, you can call the [ListLogicTables](https://help.aliyun.com/document_detail/141875.html) operation to obtain the GUID. The value of the ReturnGuid parameter must be set to true.
	//
	// 	- If the database to which the table belongs is a physical database, you can call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to obtain the GUID. The value of the ReturnGuid parameter must be set to true.
	//
	// This parameter is required.
	//
	// example:
	//
	// IDB_L_9032.db-test.yuyang_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetTableDBTopologyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableDBTopologyRequest) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetTableDBTopologyRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetTableDBTopologyRequest) SetTableGuid(v string) *GetTableDBTopologyRequest {
	s.TableGuid = &v
	return s
}

func (s *GetTableDBTopologyRequest) SetTid(v int64) *GetTableDBTopologyRequest {
	s.Tid = &v
	return s
}

func (s *GetTableDBTopologyRequest) Validate() error {
	return dara.Validate(s)
}
