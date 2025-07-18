// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnloadSampleDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UnloadSampleDataRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *UnloadSampleDataRequest
	GetOwnerId() *int64
}

type UnloadSampleDataRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UnloadSampleDataRequest) String() string {
	return dara.Prettify(s)
}

func (s UnloadSampleDataRequest) GoString() string {
	return s.String()
}

func (s *UnloadSampleDataRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UnloadSampleDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnloadSampleDataRequest) SetDBInstanceId(v string) *UnloadSampleDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UnloadSampleDataRequest) SetOwnerId(v int64) *UnloadSampleDataRequest {
	s.OwnerId = &v
	return s
}

func (s *UnloadSampleDataRequest) Validate() error {
	return dara.Validate(s)
}
