// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *PauseInstanceRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *PauseInstanceRequest
	GetOwnerId() *int64
}

type PauseInstanceRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s PauseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseInstanceRequest) GoString() string {
	return s.String()
}

func (s *PauseInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *PauseInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PauseInstanceRequest) SetDBInstanceId(v string) *PauseInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *PauseInstanceRequest) SetOwnerId(v int64) *PauseInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *PauseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
