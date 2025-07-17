// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableTopologyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTableGuid(v string) *GetTableTopologyRequest
	GetTableGuid() *string
	SetTid(v int64) *GetTableTopologyRequest
	GetTid() *int64
}

type GetTableTopologyRequest struct {
	// The GUID of the table in Data Management (DMS).
	//
	// >
	//
	// > - You can call the [ListLogicTables](https://help.aliyun.com/document_detail/141875.html) operation with ReturnGuid set to true to query the GUIDs of logical tables in a specific logical database.
	//
	// > - You can call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation with ReturnGuid set to true to query the GUIDs of tables in a specific physical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// IDB_L_308302.yuyang_test.test_ch
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the DMS console. For more information, see the "View information about the current tenant" section of the [Tenant information](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetTableTopologyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableTopologyRequest) GoString() string {
	return s.String()
}

func (s *GetTableTopologyRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetTableTopologyRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetTableTopologyRequest) SetTableGuid(v string) *GetTableTopologyRequest {
	s.TableGuid = &v
	return s
}

func (s *GetTableTopologyRequest) SetTid(v int64) *GetTableTopologyRequest {
	s.Tid = &v
	return s
}

func (s *GetTableTopologyRequest) Validate() error {
	return dara.Validate(s)
}
