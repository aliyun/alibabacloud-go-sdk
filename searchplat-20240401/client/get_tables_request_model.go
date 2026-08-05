// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v string) *GetTablesRequest
	GetParams() *string
	SetRegionId(v string) *GetTablesRequest
	GetRegionId() *string
}

type GetTablesRequest struct {
	// The data source parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "instanceId": "instance_id",
	//
	//     "dbName": "db_name",
	//
	//     "dbUser": "db_user",
	//
	//     "dbPassword": "passwoed"
	//
	// }
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTablesRequest) GoString() string {
	return s.String()
}

func (s *GetTablesRequest) GetParams() *string {
	return s.Params
}

func (s *GetTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTablesRequest) SetParams(v string) *GetTablesRequest {
	s.Params = &v
	return s
}

func (s *GetTablesRequest) SetRegionId(v string) *GetTablesRequest {
	s.RegionId = &v
	return s
}

func (s *GetTablesRequest) Validate() error {
	return dara.Validate(s)
}
