// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreRangeInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DescribeRestoreRangeInfoRequest
	GetBackupPlanId() *string
	SetBeginTimestampForRestore(v int64) *DescribeRestoreRangeInfoRequest
	GetBeginTimestampForRestore() *int64
	SetClientToken(v string) *DescribeRestoreRangeInfoRequest
	GetClientToken() *string
	SetEndTimestampForRestore(v int64) *DescribeRestoreRangeInfoRequest
	GetEndTimestampForRestore() *int64
	SetOwnerId(v string) *DescribeRestoreRangeInfoRequest
	GetOwnerId() *string
	SetRecentlyRestore(v bool) *DescribeRestoreRangeInfoRequest
	GetRecentlyRestore() *bool
}

type DescribeRestoreRangeInfoRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain it.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbssl67c7mx****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The earliest point in time to which you can restore data. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the value of the parameter for each backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1646674092000
	BeginTimestampForRestore *int64 `json:"BeginTimestampForRestore,omitempty" xml:"BeginTimestampForRestore,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The latest point in time to which you can restore data. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the value of the parameter for each backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1646846814000
	EndTimestampForRestore *int64  `json:"EndTimestampForRestore,omitempty" xml:"EndTimestampForRestore,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to query the most recent point in time to which you can restore data.
	//
	// example:
	//
	// true
	RecentlyRestore *bool `json:"RecentlyRestore,omitempty" xml:"RecentlyRestore,omitempty"`
}

func (s DescribeRestoreRangeInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreRangeInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeRestoreRangeInfoRequest) GetBeginTimestampForRestore() *int64 {
	return s.BeginTimestampForRestore
}

func (s *DescribeRestoreRangeInfoRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeRestoreRangeInfoRequest) GetEndTimestampForRestore() *int64 {
	return s.EndTimestampForRestore
}

func (s *DescribeRestoreRangeInfoRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeRestoreRangeInfoRequest) GetRecentlyRestore() *bool {
	return s.RecentlyRestore
}

func (s *DescribeRestoreRangeInfoRequest) SetBackupPlanId(v string) *DescribeRestoreRangeInfoRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeRestoreRangeInfoRequest) SetBeginTimestampForRestore(v int64) *DescribeRestoreRangeInfoRequest {
	s.BeginTimestampForRestore = &v
	return s
}

func (s *DescribeRestoreRangeInfoRequest) SetClientToken(v string) *DescribeRestoreRangeInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeRestoreRangeInfoRequest) SetEndTimestampForRestore(v int64) *DescribeRestoreRangeInfoRequest {
	s.EndTimestampForRestore = &v
	return s
}

func (s *DescribeRestoreRangeInfoRequest) SetOwnerId(v string) *DescribeRestoreRangeInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRestoreRangeInfoRequest) SetRecentlyRestore(v bool) *DescribeRestoreRangeInfoRequest {
	s.RecentlyRestore = &v
	return s
}

func (s *DescribeRestoreRangeInfoRequest) Validate() error {
	return dara.Validate(s)
}
