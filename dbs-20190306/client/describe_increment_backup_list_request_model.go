// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIncrementBackupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DescribeIncrementBackupListRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *DescribeIncrementBackupListRequest
	GetClientToken() *string
	SetEndTimestamp(v int64) *DescribeIncrementBackupListRequest
	GetEndTimestamp() *int64
	SetOwnerId(v string) *DescribeIncrementBackupListRequest
	GetOwnerId() *string
	SetPageNum(v int32) *DescribeIncrementBackupListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeIncrementBackupListRequest
	GetPageSize() *int32
	SetShowStorageType(v bool) *DescribeIncrementBackupListRequest
	GetShowStorageType() *bool
	SetStartTimestamp(v int64) *DescribeIncrementBackupListRequest
	GetStartTimestamp() *int64
}

type DescribeIncrementBackupListRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbsqdss5tmh****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 1570701361600
	EndTimestamp *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be a positive integer. Default value: 0.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Valid values: 30, 50, and 100.
	//
	// > Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to return the storage class. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// > Default value: true.
	//
	// example:
	//
	// true
	ShowStorageType *bool `json:"ShowStorageType,omitempty" xml:"ShowStorageType,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 1570701361528
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeIncrementBackupListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIncrementBackupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeIncrementBackupListRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeIncrementBackupListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeIncrementBackupListRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeIncrementBackupListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeIncrementBackupListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeIncrementBackupListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIncrementBackupListRequest) GetShowStorageType() *bool {
	return s.ShowStorageType
}

func (s *DescribeIncrementBackupListRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeIncrementBackupListRequest) SetBackupPlanId(v string) *DescribeIncrementBackupListRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeIncrementBackupListRequest) SetClientToken(v string) *DescribeIncrementBackupListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeIncrementBackupListRequest) SetEndTimestamp(v int64) *DescribeIncrementBackupListRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeIncrementBackupListRequest) SetOwnerId(v string) *DescribeIncrementBackupListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIncrementBackupListRequest) SetPageNum(v int32) *DescribeIncrementBackupListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeIncrementBackupListRequest) SetPageSize(v int32) *DescribeIncrementBackupListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIncrementBackupListRequest) SetShowStorageType(v bool) *DescribeIncrementBackupListRequest {
	s.ShowStorageType = &v
	return s
}

func (s *DescribeIncrementBackupListRequest) SetStartTimestamp(v int64) *DescribeIncrementBackupListRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeIncrementBackupListRequest) Validate() error {
	return dara.Validate(s)
}
