// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseDataRedistributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *PauseDataRedistributeRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *PauseDataRedistributeRequest
	GetOwnerId() *int64
}

type PauseDataRedistributeRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s PauseDataRedistributeRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseDataRedistributeRequest) GoString() string {
	return s.String()
}

func (s *PauseDataRedistributeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *PauseDataRedistributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PauseDataRedistributeRequest) SetDBInstanceId(v string) *PauseDataRedistributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *PauseDataRedistributeRequest) SetOwnerId(v int64) *PauseDataRedistributeRequest {
	s.OwnerId = &v
	return s
}

func (s *PauseDataRedistributeRequest) Validate() error {
	return dara.Validate(s)
}
