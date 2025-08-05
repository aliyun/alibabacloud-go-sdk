// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOtsTableSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeOtsTableSnapshotsResponseBody
	GetCode() *string
	SetLimit(v int32) *DescribeOtsTableSnapshotsResponseBody
	GetLimit() *int32
	SetMessage(v string) *DescribeOtsTableSnapshotsResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeOtsTableSnapshotsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeOtsTableSnapshotsResponseBody
	GetRequestId() *string
	SetSnapshots(v []*DescribeOtsTableSnapshotsResponseBodySnapshots) *DescribeOtsTableSnapshotsResponseBody
	GetSnapshots() []*DescribeOtsTableSnapshotsResponseBodySnapshots
	SetSuccess(v bool) *DescribeOtsTableSnapshotsResponseBody
	GetSuccess() *bool
}

type DescribeOtsTableSnapshotsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of backup snapshots that are displayed on the current page.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token that is required to obtain the next page of backup snapshots.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 09376812-6290-5E36-B504-E8010D72D1F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The backup snapshots.
	Snapshots []*DescribeOtsTableSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeOtsTableSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOtsTableSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOtsTableSnapshotsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeOtsTableSnapshotsResponseBody) GetLimit() *int32 {
	return s.Limit
}

func (s *DescribeOtsTableSnapshotsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeOtsTableSnapshotsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeOtsTableSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOtsTableSnapshotsResponseBody) GetSnapshots() []*DescribeOtsTableSnapshotsResponseBodySnapshots {
	return s.Snapshots
}

func (s *DescribeOtsTableSnapshotsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetCode(v string) *DescribeOtsTableSnapshotsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetLimit(v int32) *DescribeOtsTableSnapshotsResponseBody {
	s.Limit = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetMessage(v string) *DescribeOtsTableSnapshotsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetNextToken(v string) *DescribeOtsTableSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetRequestId(v string) *DescribeOtsTableSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetSnapshots(v []*DescribeOtsTableSnapshotsResponseBodySnapshots) *DescribeOtsTableSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetSuccess(v bool) *DescribeOtsTableSnapshotsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeOtsTableSnapshotsResponseBodySnapshots struct {
	// The actual data amount of backup snapshots after duplicates are removed. Unit: bytes.
	//
	// example:
	//
	// 0
	ActualBytes *string `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	// The backup type. Valid value: **COMPLETE**, which indicates full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The total amount of data. Unit: bytes.
	//
	// example:
	//
	// 0
	BytesTotal *int64 `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	// The time when the backup snapshot was completed. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1642496679
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the Tablestore instance was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1607436917
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the backup snapshot was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1642496679
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The name of the Tablestore instance.
	//
	// example:
	//
	// instancename
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the backup job.
	//
	// example:
	//
	// job-00030j3chkt******2
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The hash value of the parent backup snapshot.
	//
	// example:
	//
	// f2fe..
	ParentSnapshotHash *string `json:"ParentSnapshotHash,omitempty" xml:"ParentSnapshotHash,omitempty"`
	// The time when the backup job ended. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1642521709966
	RangeEnd *int64 `json:"RangeEnd,omitempty" xml:"RangeEnd,omitempty"`
	// The time when the backup job started. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1642492553038
	RangeStart *int64 `json:"RangeStart,omitempty" xml:"RangeStart,omitempty"`
	// The retention period of the backup snapshot. Unit: days.
	//
	// example:
	//
	// 730
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The hash value of the backup snapshot.
	//
	// example:
	//
	// f2fe...
	SnapshotHash *string `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	// The ID of the backup snapshot.
	//
	// example:
	//
	// s-00047mxg17p26*****b
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: backup snapshots for Elastic Compute Service (ECS) files
	//
	// 	- **OSS**: backup snapshots for Object Storage Service (OSS) buckets
	//
	// 	- **NAS**: backup snapshots for Apsara File Storage NAS file systems
	//
	// 	- **OTS_TABLE**: backup snapshots for Tablestore instances
	//
	// example:
	//
	// OTS_TABLE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time when the backup snapshot started. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1642496543
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the backup job. Valid values:
	//
	// 	- **COMPLETE**: The backup job is completed.
	//
	// 	- **PARTIAL_COMPLETE**: The backup job is partially completed.
	//
	// 	- **FAILED**: The backup job has failed.
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the table in the Tablestore instance.
	//
	// example:
	//
	// table2
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The time when the backup snapshot was updated. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1642496679
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The ID of the backup vault that stores the backup snapshot.
	//
	// example:
	//
	// v-00030j*******sn
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeOtsTableSnapshotsResponseBodySnapshots) String() string {
	return dara.Prettify(s)
}

func (s DescribeOtsTableSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetActualBytes() *string {
	return s.ActualBytes
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetBytesTotal() *int64 {
	return s.BytesTotal
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetCompleteTime() *int64 {
	return s.CompleteTime
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetJobId() *string {
	return s.JobId
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetParentSnapshotHash() *string {
	return s.ParentSnapshotHash
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetRangeEnd() *int64 {
	return s.RangeEnd
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetRangeStart() *int64 {
	return s.RangeStart
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetRetention() *int64 {
	return s.Retention
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetSnapshotHash() *string {
	return s.SnapshotHash
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetStatus() *string {
	return s.Status
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetTableName() *string {
	return s.TableName
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetActualBytes(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.ActualBytes = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetBackupType(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.BackupType = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetBytesTotal(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.BytesTotal = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetCompleteTime(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.CompleteTime = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetCreateTime(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.CreateTime = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetCreatedTime(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.CreatedTime = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetInstanceName(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.InstanceName = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetJobId(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.JobId = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetParentSnapshotHash(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.ParentSnapshotHash = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetRangeEnd(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.RangeEnd = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetRangeStart(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.RangeStart = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetRetention(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.Retention = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetSnapshotHash(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.SnapshotHash = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetSourceType(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.SourceType = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetStartTime(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.StartTime = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetStatus(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetTableName(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.TableName = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetUpdatedTime(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetVaultId(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.VaultId = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) Validate() error {
	return dara.Validate(s)
}
