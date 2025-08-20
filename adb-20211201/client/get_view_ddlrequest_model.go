// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetViewDDLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetViewDDLRequest
	GetDBClusterId() *string
	SetRegionId(v string) *GetViewDDLRequest
	GetRegionId() *string
	SetSchemaName(v string) *GetViewDDLRequest
	GetSchemaName() *string
	SetViewName(v string) *GetViewDDLRequest
	GetViewName() *string
}

type GetViewDDLRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the view.
	//
	// example:
	//
	// v_modbus
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s GetViewDDLRequest) String() string {
	return dara.Prettify(s)
}

func (s GetViewDDLRequest) GoString() string {
	return s.String()
}

func (s *GetViewDDLRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetViewDDLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetViewDDLRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetViewDDLRequest) GetViewName() *string {
	return s.ViewName
}

func (s *GetViewDDLRequest) SetDBClusterId(v string) *GetViewDDLRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetViewDDLRequest) SetRegionId(v string) *GetViewDDLRequest {
	s.RegionId = &v
	return s
}

func (s *GetViewDDLRequest) SetSchemaName(v string) *GetViewDDLRequest {
	s.SchemaName = &v
	return s
}

func (s *GetViewDDLRequest) SetViewName(v string) *GetViewDDLRequest {
	s.ViewName = &v
	return s
}

func (s *GetViewDDLRequest) Validate() error {
	return dara.Validate(s)
}
