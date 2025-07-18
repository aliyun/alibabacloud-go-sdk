// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateSampleDataRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *CreateSampleDataRequest
	GetOwnerId() *int64
}

type CreateSampleDataRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/2361776.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateSampleDataRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleDataRequest) GoString() string {
	return s.String()
}

func (s *CreateSampleDataRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateSampleDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSampleDataRequest) SetDBInstanceId(v string) *CreateSampleDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateSampleDataRequest) SetOwnerId(v int64) *CreateSampleDataRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSampleDataRequest) Validate() error {
	return dara.Validate(s)
}
