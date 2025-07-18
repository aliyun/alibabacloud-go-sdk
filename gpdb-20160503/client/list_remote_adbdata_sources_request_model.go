// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemoteADBDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListRemoteADBDataSourcesRequest
	GetDBInstanceId() *string
	SetDataSourceId(v string) *ListRemoteADBDataSourcesRequest
	GetDataSourceId() *string
	SetOwnerId(v int64) *ListRemoteADBDataSourcesRequest
	GetOwnerId() *int64
}

type ListRemoteADBDataSourcesRequest struct {
	// Instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Data source ID.
	//
	// example:
	//
	// 119
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ListRemoteADBDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRemoteADBDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListRemoteADBDataSourcesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListRemoteADBDataSourcesRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ListRemoteADBDataSourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListRemoteADBDataSourcesRequest) SetDBInstanceId(v string) *ListRemoteADBDataSourcesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListRemoteADBDataSourcesRequest) SetDataSourceId(v string) *ListRemoteADBDataSourcesRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListRemoteADBDataSourcesRequest) SetOwnerId(v int64) *ListRemoteADBDataSourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListRemoteADBDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
