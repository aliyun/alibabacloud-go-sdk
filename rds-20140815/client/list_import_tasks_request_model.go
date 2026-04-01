// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImportTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListImportTasksRequest
	GetDBInstanceId() *string
	SetMaxResults(v int32) *ListImportTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListImportTasksRequest
	GetNextToken() *string
	SetOwnerId(v int64) *ListImportTasksRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListImportTasksRequest
	GetRegionId() *string
}

type ListImportTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListImportTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImportTasksRequest) GoString() string {
	return s.String()
}

func (s *ListImportTasksRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListImportTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListImportTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListImportTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListImportTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImportTasksRequest) SetDBInstanceId(v string) *ListImportTasksRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListImportTasksRequest) SetMaxResults(v int32) *ListImportTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListImportTasksRequest) SetNextToken(v string) *ListImportTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListImportTasksRequest) SetOwnerId(v int64) *ListImportTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *ListImportTasksRequest) SetRegionId(v string) *ListImportTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ListImportTasksRequest) Validate() error {
	return dara.Validate(s)
}
