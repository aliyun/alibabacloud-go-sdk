// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSQLLogFilesRequest
	GetDBInstanceId() *string
	SetFileName(v string) *DescribeSQLLogFilesRequest
	GetFileName() *string
	SetOwnerAccount(v string) *DescribeSQLLogFilesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSQLLogFilesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSQLLogFilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSQLLogFilesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeSQLLogFilesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSQLLogFilesRequest
	GetResourceOwnerId() *int64
}

type DescribeSQLLogFilesRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the audit log file.
	//
	// example:
	//
	// custinsxxxxx.csv
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: **1 to 100000**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30 to 200**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSQLLogFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSQLLogFilesRequest) GetFileName() *string {
	return s.FileName
}

func (s *DescribeSQLLogFilesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSQLLogFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSQLLogFilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSQLLogFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSQLLogFilesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSQLLogFilesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSQLLogFilesRequest) SetDBInstanceId(v string) *DescribeSQLLogFilesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetFileName(v string) *DescribeSQLLogFilesRequest {
	s.FileName = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetOwnerAccount(v string) *DescribeSQLLogFilesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetOwnerId(v int64) *DescribeSQLLogFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetPageNumber(v int32) *DescribeSQLLogFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetPageSize(v int32) *DescribeSQLLogFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetResourceOwnerAccount(v string) *DescribeSQLLogFilesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetResourceOwnerId(v int64) *DescribeSQLLogFilesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) Validate() error {
	return dara.Validate(s)
}
