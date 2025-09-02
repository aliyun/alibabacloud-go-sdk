// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaDBInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGuid(v string) *GetMetaDBInfoRequest
	GetAppGuid() *string
	SetClusterId(v string) *GetMetaDBInfoRequest
	GetClusterId() *string
	SetDataSourceType(v string) *GetMetaDBInfoRequest
	GetDataSourceType() *string
	SetDatabaseName(v string) *GetMetaDBInfoRequest
	GetDatabaseName() *string
}

type GetMetaDBInfoRequest struct {
	// The compute engine instance ID. Specify the ID in the `Engine type.Engine name` format.
	//
	// example:
	//
	// odps.some_engine_name
	AppGuid *string `json:"AppGuid,omitempty" xml:"AppGuid,omitempty"`
	// The E-MapReduce (EMR) cluster ID.
	//
	// example:
	//
	// abc
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the data source. Valid values: odps and emr.
	//
	// example:
	//
	// emr
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The name of the metadatabase of the EMR cluster.
	//
	// example:
	//
	// abc
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
}

func (s GetMetaDBInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaDBInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMetaDBInfoRequest) GetAppGuid() *string {
	return s.AppGuid
}

func (s *GetMetaDBInfoRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMetaDBInfoRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetMetaDBInfoRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaDBInfoRequest) SetAppGuid(v string) *GetMetaDBInfoRequest {
	s.AppGuid = &v
	return s
}

func (s *GetMetaDBInfoRequest) SetClusterId(v string) *GetMetaDBInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMetaDBInfoRequest) SetDataSourceType(v string) *GetMetaDBInfoRequest {
	s.DataSourceType = &v
	return s
}

func (s *GetMetaDBInfoRequest) SetDatabaseName(v string) *GetMetaDBInfoRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaDBInfoRequest) Validate() error {
	return dara.Validate(s)
}
