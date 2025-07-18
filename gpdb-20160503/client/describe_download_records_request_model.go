// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDownloadRecordsRequest
	GetDBInstanceId() *string
	SetDownloadTaskType(v string) *DescribeDownloadRecordsRequest
	GetDownloadTaskType() *string
}

type DescribeDownloadRecordsRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DownloadTaskType *string `json:"DownloadTaskType,omitempty" xml:"DownloadTaskType,omitempty"`
}

func (s DescribeDownloadRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDownloadRecordsRequest) GetDownloadTaskType() *string {
	return s.DownloadTaskType
}

func (s *DescribeDownloadRecordsRequest) SetDBInstanceId(v string) *DescribeDownloadRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDownloadRecordsRequest) SetDownloadTaskType(v string) *DescribeDownloadRecordsRequest {
	s.DownloadTaskType = &v
	return s
}

func (s *DescribeDownloadRecordsRequest) Validate() error {
	return dara.Validate(s)
}
