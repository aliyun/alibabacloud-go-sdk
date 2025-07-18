// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeDataRedistributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ResumeDataRedistributeRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *ResumeDataRedistributeRequest
	GetOwnerId() *int64
}

type ResumeDataRedistributeRequest struct {
	// The instance ID.
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

func (s ResumeDataRedistributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeDataRedistributeRequest) GoString() string {
	return s.String()
}

func (s *ResumeDataRedistributeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ResumeDataRedistributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResumeDataRedistributeRequest) SetDBInstanceId(v string) *ResumeDataRedistributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ResumeDataRedistributeRequest) SetOwnerId(v int64) *ResumeDataRedistributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ResumeDataRedistributeRequest) Validate() error {
	return dara.Validate(s)
}
