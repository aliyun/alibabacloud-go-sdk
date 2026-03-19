// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSetDownloadTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DescribeBackupSetDownloadTaskListRequest
	GetBackupPlanId() *string
	SetBackupSetDownloadTaskId(v string) *DescribeBackupSetDownloadTaskListRequest
	GetBackupSetDownloadTaskId() *string
	SetClientToken(v string) *DescribeBackupSetDownloadTaskListRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeBackupSetDownloadTaskListRequest
	GetOwnerId() *string
	SetPageNum(v int32) *DescribeBackupSetDownloadTaskListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeBackupSetDownloadTaskListRequest
	GetPageSize() *int32
}

type DescribeBackupSetDownloadTaskListRequest struct {
	// The ID of the backup plan. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the value of this parameter.
	//
	// > You must specify either **BackupPlanId*	- or **BackupSetDownloadTaskId**.
	//
	// example:
	//
	// dbsqhnuhyw3****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the backup set download task.
	//
	// - For a full backup, call the [CreateFullBackupSetDownload](https://help.aliyun.com/document_detail/2869842.html) operation to obtain the value of this parameter.
	//
	// - For an incremental backup, call the [CreateIncrementBackupSetDownload](https://help.aliyun.com/document_detail/2869843.html) operation to obtain the value of this parameter.
	//
	// example:
	//
	// urxgrxt7****
	BackupSetDownloadTaskId *string `json:"BackupSetDownloadTaskId,omitempty" xml:"BackupSetDownloadTaskId,omitempty"`
	// A client token that is used to ensure the idempotence of the request. This prevents duplicate requests.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzXXXXXX
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be a non-negative integer that does not exceed the maximum value of the integer type. The default value is 0.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Valid values: 30, 50, and 100.
	//
	// > The default value is 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeBackupSetDownloadTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetDownloadTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetDownloadTaskListRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeBackupSetDownloadTaskListRequest) GetBackupSetDownloadTaskId() *string {
	return s.BackupSetDownloadTaskId
}

func (s *DescribeBackupSetDownloadTaskListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeBackupSetDownloadTaskListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeBackupSetDownloadTaskListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeBackupSetDownloadTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupSetDownloadTaskListRequest) SetBackupPlanId(v string) *DescribeBackupSetDownloadTaskListRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListRequest) SetBackupSetDownloadTaskId(v string) *DescribeBackupSetDownloadTaskListRequest {
	s.BackupSetDownloadTaskId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListRequest) SetClientToken(v string) *DescribeBackupSetDownloadTaskListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListRequest) SetOwnerId(v string) *DescribeBackupSetDownloadTaskListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListRequest) SetPageNum(v int32) *DescribeBackupSetDownloadTaskListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListRequest) SetPageSize(v int32) *DescribeBackupSetDownloadTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListRequest) Validate() error {
	return dara.Validate(s)
}
