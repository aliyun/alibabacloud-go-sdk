// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v string) *GetTableFieldsRequest
	GetParams() *string
	SetRawType(v bool) *GetTableFieldsRequest
	GetRawType() *bool
	SetRegionId(v string) *GetTableFieldsRequest
	GetRegionId() *string
}

type GetTableFieldsRequest struct {
	// The data source parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	// "accessKeySecret": "sk",
	//
	// "accessKey": "ak",
	//
	// "projectName": "test_name",
	//
	// "tableName": "test_table",
	//
	// "partition": "20240904"
	//
	// }
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// Specifies whether to return the original field types of the data source.
	//
	// example:
	//
	// false
	RawType *bool `json:"rawType,omitempty" xml:"rawType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetTableFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableFieldsRequest) GoString() string {
	return s.String()
}

func (s *GetTableFieldsRequest) GetParams() *string {
	return s.Params
}

func (s *GetTableFieldsRequest) GetRawType() *bool {
	return s.RawType
}

func (s *GetTableFieldsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTableFieldsRequest) SetParams(v string) *GetTableFieldsRequest {
	s.Params = &v
	return s
}

func (s *GetTableFieldsRequest) SetRawType(v bool) *GetTableFieldsRequest {
	s.RawType = &v
	return s
}

func (s *GetTableFieldsRequest) SetRegionId(v string) *GetTableFieldsRequest {
	s.RegionId = &v
	return s
}

func (s *GetTableFieldsRequest) Validate() error {
	return dara.Validate(s)
}
