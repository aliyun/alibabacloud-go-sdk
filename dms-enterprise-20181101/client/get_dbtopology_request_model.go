// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBTopologyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogicDbId(v int64) *GetDBTopologyRequest
	GetLogicDbId() *int64
	SetTid(v int64) *GetDBTopologyRequest
	GetTid() *int64
}

type GetDBTopologyRequest struct {
	// The ID of the logical database. You can call the [ListLogicDatabases](https://www.alibabacloud.com/help/en/data-management-service/latest/listlogicdatabases) or [SearchDatabase](https://www.alibabacloud.com/help/en/data-management-service/latest/searchdatabase) operation to query the ID of the logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 134***
	LogicDbId *int64 `json:"LogicDbId,omitempty" xml:"LogicDbId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 43***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDBTopologyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDBTopologyRequest) GoString() string {
	return s.String()
}

func (s *GetDBTopologyRequest) GetLogicDbId() *int64 {
	return s.LogicDbId
}

func (s *GetDBTopologyRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDBTopologyRequest) SetLogicDbId(v int64) *GetDBTopologyRequest {
	s.LogicDbId = &v
	return s
}

func (s *GetDBTopologyRequest) SetTid(v int64) *GetDBTopologyRequest {
	s.Tid = &v
	return s
}

func (s *GetDBTopologyRequest) Validate() error {
	return dara.Validate(s)
}
