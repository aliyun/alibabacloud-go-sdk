// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v string) *GetTableColumnsRequest
	GetParams() *string
	SetRegionId(v string) *GetTableColumnsRequest
	GetRegionId() *string
}

type GetTableColumnsRequest struct {
	// The configuration parameters for accessing the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds:
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
	//
	// odps:
	//
	// {
	//
	//     "accessKeySecret": "sk",
	//
	//     "accessKey": "ak",
	//
	//     "projectName": "test_name"
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

func (s GetTableColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnsRequest) GoString() string {
	return s.String()
}

func (s *GetTableColumnsRequest) GetParams() *string {
	return s.Params
}

func (s *GetTableColumnsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTableColumnsRequest) SetParams(v string) *GetTableColumnsRequest {
	s.Params = &v
	return s
}

func (s *GetTableColumnsRequest) SetRegionId(v string) *GetTableColumnsRequest {
	s.RegionId = &v
	return s
}

func (s *GetTableColumnsRequest) Validate() error {
	return dara.Validate(s)
}
