// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRemoteADBDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v string) *DeleteRemoteADBDataSourceRequest
	GetDataSourceId() *string
	SetLocalDBInstanceId(v string) *DeleteRemoteADBDataSourceRequest
	GetLocalDBInstanceId() *string
	SetOwnerId(v int64) *DeleteRemoteADBDataSourceRequest
	GetOwnerId() *int64
}

type DeleteRemoteADBDataSourceRequest struct {
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The ID of the instance that uses the data provided by another instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-test
	LocalDBInstanceId *string `json:"LocalDBInstanceId,omitempty" xml:"LocalDBInstanceId,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteRemoteADBDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRemoteADBDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteRemoteADBDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DeleteRemoteADBDataSourceRequest) GetLocalDBInstanceId() *string {
	return s.LocalDBInstanceId
}

func (s *DeleteRemoteADBDataSourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRemoteADBDataSourceRequest) SetDataSourceId(v string) *DeleteRemoteADBDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *DeleteRemoteADBDataSourceRequest) SetLocalDBInstanceId(v string) *DeleteRemoteADBDataSourceRequest {
	s.LocalDBInstanceId = &v
	return s
}

func (s *DeleteRemoteADBDataSourceRequest) SetOwnerId(v int64) *DeleteRemoteADBDataSourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRemoteADBDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
