// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ResumeInstanceRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *ResumeInstanceRequest
	GetOwnerId() *int64
}

type ResumeInstanceRequest struct {
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

func (s ResumeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeInstanceRequest) GoString() string {
	return s.String()
}

func (s *ResumeInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ResumeInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResumeInstanceRequest) SetDBInstanceId(v string) *ResumeInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ResumeInstanceRequest) SetOwnerId(v int64) *ResumeInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ResumeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
